package obj

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"sync"

	"github.com/pachyderm/pachyderm/src/client/pkg/grpcutil"

	"github.com/Azure/azure-sdk-for-go/storage"
)

type microsoftClient struct {
	blobClient storage.BlobStorageClient
	container  string
}

func newMicrosoftClient(container string, accountName string, accountKey string) (*microsoftClient, error) {
	client, err := storage.NewBasicClient(
		accountName,
		accountKey,
	)
	if err != nil {
		return nil, err
	}

	return &microsoftClient{
		blobClient: client.GetBlobService(),
		container:  container,
	}, nil
}

func (c *microsoftClient) Writer(name string) (io.WriteCloser, error) {
	writer, err := newMicrosoftWriter(c, name)
	if err != nil {
		return nil, err
	}
	return newBackoffWriteCloser(c, writer), nil
}

func (c *microsoftClient) Reader(name string, offset uint64, size uint64) (io.ReadCloser, error) {
	byteRange := byteRange(offset, size)
	var reader io.ReadCloser
	var err error
	if byteRange == "" {
		reader, err = c.blobClient.GetBlob(c.container, name)
	} else {
		reader, err = c.blobClient.GetBlobRange(c.container, name, byteRange, nil)
	}

	if err != nil {
		return nil, err
	}
	return newBackoffReadCloser(c, reader), nil
}

func (c *microsoftClient) Delete(name string) error {
	return c.blobClient.DeleteBlob(c.container, name, nil)
}

func (c *microsoftClient) Walk(name string, fn func(name string) error) error {
	// See Azure docs for what `marker` does:
	// https://docs.microsoft.com/en-us/rest/api/storageservices/List-Blobs?redirectedfrom=MSDN
	var marker string
	for {
		blobList, err := c.blobClient.ListBlobs(c.container, storage.ListBlobsParameters{
			Prefix: name,
			Marker: marker,
		})
		if err != nil {
			return err
		}

		for _, file := range blobList.Blobs {
			if err := fn(file.Name); err != nil {
				return err
			}
		}

		// NextMarker is empty when all results have been returned
		if blobList.NextMarker == "" {
			break
		}
		marker = blobList.NextMarker
	}
	return nil
}

func (c *microsoftClient) Exists(name string) bool {
	exists, _ := c.blobClient.BlobExists(c.container, name)
	return exists
}

func (c *microsoftClient) IsRetryable(err error) (ret bool) {
	microsoftErr, ok := err.(storage.AzureStorageServiceError)
	if !ok {
		return false
	}
	return microsoftErr.StatusCode >= 500
}

func (c *microsoftClient) IsNotExist(err error) bool {
	microsoftErr, ok := err.(storage.AzureStorageServiceError)
	if !ok {
		return false
	}
	return microsoftErr.StatusCode == 404
}

func (c *microsoftClient) IsIgnorable(err error) bool {
	return false
}

type microsoftWriter struct {
	container            string
	blob                 string
	blobClient           storage.BlobStorageClient
	blocksChan           chan blockInfo
	blockWriterWaitGroup *sync.WaitGroup
	blockCount           int
	blockWriterErrorChan chan error
}

type blockInfo struct {
	id               string
	name             string
	container        string
	data             []byte
	blockWriterError chan error
}

const numberOfWorkers = 2
const writeBlockChannelLen = 2

func newMicrosoftWriter(client *microsoftClient, name string) (*microsoftWriter, error) {
	// create container
	_, err := client.blobClient.CreateContainerIfNotExists(client.container, storage.ContainerAccessTypePrivate)
	if err != nil {
		return nil, err
	}

	exists, err := client.blobClient.BlobExists(client.container, name)


	if err != nil {
		return nil, err
	}

	blockCount := 0
	if exists {
		blockList, err := client.blobClient.GetBlockList(client.container, name, storage.BlockListTypeAll)
		if err != nil {
			return nil, err
		}
		blockCount = len(blockList.CommittedBlocks)
	} else {
		err = client.blobClient.CreateBlockBlob(client.container, name)
		if err != nil {
			return nil, err
		}
	}

	msWriter := &microsoftWriter{
		container:            client.container,
		blob:                 name,
		blobClient:           client.blobClient,
		blockWriterErrorChan: make(chan error, writeBlockChannelLen),
		blockCount:           blockCount,
	}

	msWriter.startWorkers(numberOfWorkers)

	return msWriter, nil
}

func (w *microsoftWriter) newBlockInfo(blockOrdinal int, data []byte, length int) blockInfo {

	return blockInfo{
		id:               base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%011d\n", blockOrdinal))),
		data:             data[:length],
		name:             w.blob,
		container:        w.container,
		blockWriterError: w.blockWriterErrorChan,
	}
}

func (w *microsoftWriter) startWorkers(numOfWorkers int) {
	wg := sync.WaitGroup{}
	blockChannel := make(chan blockInfo, writeBlockChannelLen)

	for writers := 0; writers < numOfWorkers; writers++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			for {
				blockInfo, ok := <-blockChannel

				if !ok {
					break
				}

				err := w.blobClient.PutBlock(blockInfo.container, blockInfo.name, blockInfo.id, blockInfo.data)
				grpcutil.PutBuffer(blockInfo.data)
				if err != nil {
					select {
					case blockInfo.blockWriterError <- err:
					default:
					}
				}
			}
		}()
	}

	w.blocksChan = blockChannel
	w.blockWriterWaitGroup = &wg
}

func (w *microsoftWriter) Write(b []byte) (int, error) {

	inputSourceReader := bytes.NewReader(b)
	for {
		chunk := grpcutil.GetBuffer()
		n, err := inputSourceReader.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
		select {
		case w.blocksChan <- w.newBlockInfo(w.blockCount, chunk, n):
			w.blockCount++
		case err := <-w.blockWriterErrorChan:
			return n, err
		}
	}
	return len(b), nil
}

func (w *microsoftWriter) Close() error {
	close(w.blocksChan)
	w.blockWriterWaitGroup.Wait()
	select {
	case err := <-w.blockWriterErrorChan:
		return err
	default:
	}

	blockList := make([]storage.Block, w.blockCount)

	for blockOrdinal := 0; blockOrdinal < w.blockCount; blockOrdinal++ {
		blockID := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%011d\n", blockOrdinal)))
		blockList[blockOrdinal] = storage.Block{ID: blockID, Status: storage.BlockStatusUncommitted}
	}
	err := w.blobClient.PutBlockList(w.container, w.blob, blockList)
	if err != nil {
		return fmt.Errorf("BlobStorageClient.PutBlockList: %v", err)
	}
	return nil
}
