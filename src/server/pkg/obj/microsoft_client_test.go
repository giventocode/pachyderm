package obj

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

const data = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const dataSize = 1024 * 1024 * 1024

func getDataBlock(size int) []byte {

	b := make([]byte, size)

	bdata := []byte(data)
	ln := len(bdata)

	for i := 0; i < len(b); i++ {

		b[i] = bdata[rand.Intn(ln)]

	}

	return b

}
func Test5GFile(t *testing.T) {
	oneMiB := 1024 * 1024
	fileName := fmt.Sprintf("5gb%v", time.Now().UTC().UnixNano())
	accName := os.Getenv("ACCOUNT_NAME")
	accKey := os.Getenv("ACCOUNT_KEY")
	container := os.Getenv("CONTAINER_NAME")

	msClient, err := newMicrosoftClient(container, accName, accKey)
	msWriter, err := newMicrosoftWriter(msClient, fileName)

	if err != nil {
		t.Fatal(err)
	}

	data := getDataBlock(oneMiB)

	for b := 0; b < 5*1024; b++ {
		n, err := msWriter.Write(data)

		if n != oneMiB {
			t.Fatal("bytes written don't match expected value")
		}

		if err != nil {
			t.Fatal(err)
		}
	}

	err = msWriter.Close()

	if err != nil {
		t.Fatal(err)
	}

}
func TestBasicTracking(t *testing.T) {
	accName := os.Getenv("ACCOUNT_NAME")
	accKey := os.Getenv("ACCOUNT_KEY")
	container := os.Getenv("CONTAINER_NAME")

	msClient, err := newMicrosoftClient(container, accName, accKey)

	if err != nil {
		t.Fatal(err)
	}
	blobName := fmt.Sprintf("blob%v", time.Now().UTC().UnixNano())
	msWriter, err := newMicrosoftWriter(msClient, blobName)

	if err != nil {
		t.Fatal(err)
	}

	data := getDataBlock(dataSize)
	n, err := msWriter.Write(data)

	if err != nil {
		t.Fatal(err)
	}

	if n != dataSize {
		t.Fatal("bytes written don't match expected value")
	}

	err = msWriter.Close()

	if err != nil {
		t.Fatal(err)
	}
}
