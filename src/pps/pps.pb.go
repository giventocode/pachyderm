// Code generated by protoc-gen-go.
// source: pps/pps.proto
// DO NOT EDIT!

/*
Package pps is a generated protocol buffer package.

It is generated from these files:
	pps/pps.proto

It has these top-level messages:
	Transform
	Job
	JobStatus
	JobInfo
	JobInfos
	Pipeline
	PipelineInfo
	PipelineInfos
	CreateJobRequest
	InspectJobRequest
	ListJobRequest
	GetJobLogsRequest
	CreatePipelineRequest
	InspectPipelineRequest
	ListPipelineRequest
	DeletePipelineRequest
*/
package pps

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/google-protobuf"
import google_protobuf1 "go.pedge.io/google-protobuf"
import google_protobuf2 "go.pedge.io/google-protobuf"
import pfs "go.pachyderm.com/pachyderm/src/pfs"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type JobStatusType int32

const (
	JobStatusType_JOB_STATUS_TYPE_NONE    JobStatusType = 0
	JobStatusType_JOB_STATUS_TYPE_STARTED JobStatusType = 1
	JobStatusType_JOB_STATUS_TYPE_ERROR   JobStatusType = 2
	JobStatusType_JOB_STATUS_TYPE_SUCCESS JobStatusType = 3
)

var JobStatusType_name = map[int32]string{
	0: "JOB_STATUS_TYPE_NONE",
	1: "JOB_STATUS_TYPE_STARTED",
	2: "JOB_STATUS_TYPE_ERROR",
	3: "JOB_STATUS_TYPE_SUCCESS",
}
var JobStatusType_value = map[string]int32{
	"JOB_STATUS_TYPE_NONE":    0,
	"JOB_STATUS_TYPE_STARTED": 1,
	"JOB_STATUS_TYPE_ERROR":   2,
	"JOB_STATUS_TYPE_SUCCESS": 3,
}

func (x JobStatusType) String() string {
	return proto.EnumName(JobStatusType_name, int32(x))
}

type OutputStream int32

const (
	OutputStream_OUTPUT_STREAM_NONE   OutputStream = 0
	OutputStream_OUTPUT_STREAM_STDOUT OutputStream = 1
	OutputStream_OUTPUT_STREAM_STDERR OutputStream = 2
)

var OutputStream_name = map[int32]string{
	0: "OUTPUT_STREAM_NONE",
	1: "OUTPUT_STREAM_STDOUT",
	2: "OUTPUT_STREAM_STDERR",
}
var OutputStream_value = map[string]int32{
	"OUTPUT_STREAM_NONE":   0,
	"OUTPUT_STREAM_STDOUT": 1,
	"OUTPUT_STREAM_STDERR": 2,
}

func (x OutputStream) String() string {
	return proto.EnumName(OutputStream_name, int32(x))
}

type Transform struct {
	Image string   `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	Cmd   []string `protobuf:"bytes,2,rep,name=cmd" json:"cmd,omitempty"`
}

func (m *Transform) Reset()         { *m = Transform{} }
func (m *Transform) String() string { return proto.CompactTextString(m) }
func (*Transform) ProtoMessage()    {}

type Job struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}

type JobStatus struct {
	Type      JobStatusType               `protobuf:"varint,1,opt,name=type,enum=pachyderm.pps.JobStatusType" json:"type,omitempty"`
	Timestamp *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Message   string                      `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *JobStatus) Reset()         { *m = JobStatus{} }
func (m *JobStatus) String() string { return proto.CompactTextString(m) }
func (*JobStatus) ProtoMessage()    {}

func (m *JobStatus) GetTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

// TODO(pedge): add created at?
type JobInfo struct {
	Job *Job `protobuf:"bytes,1,opt,name=job" json:"job,omitempty"`
	// Types that are valid to be assigned to Spec:
	//	*JobInfo_Transform
	//	*JobInfo_Pipeline
	Spec   isJobInfo_Spec `protobuf_oneof:"spec"`
	Input  *pfs.Commit    `protobuf:"bytes,4,opt,name=input" json:"input,omitempty"`
	Output *pfs.Commit    `protobuf:"bytes,5,opt,name=output" json:"output,omitempty"`
	// latest to earliest
	JobStatus []*JobStatus `protobuf:"bytes,6,rep,name=job_status" json:"job_status,omitempty"`
}

func (m *JobInfo) Reset()         { *m = JobInfo{} }
func (m *JobInfo) String() string { return proto.CompactTextString(m) }
func (*JobInfo) ProtoMessage()    {}

type isJobInfo_Spec interface {
	isJobInfo_Spec()
}

type JobInfo_Transform struct {
	Transform *Transform `protobuf:"bytes,2,opt,name=transform,oneof"`
}
type JobInfo_Pipeline struct {
	Pipeline *Pipeline `protobuf:"bytes,3,opt,name=pipeline,oneof"`
}

func (*JobInfo_Transform) isJobInfo_Spec() {}
func (*JobInfo_Pipeline) isJobInfo_Spec()  {}

func (m *JobInfo) GetSpec() isJobInfo_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *JobInfo) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *JobInfo) GetTransform() *Transform {
	if x, ok := m.GetSpec().(*JobInfo_Transform); ok {
		return x.Transform
	}
	return nil
}

func (m *JobInfo) GetPipeline() *Pipeline {
	if x, ok := m.GetSpec().(*JobInfo_Pipeline); ok {
		return x.Pipeline
	}
	return nil
}

func (m *JobInfo) GetInput() *pfs.Commit {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *JobInfo) GetOutput() *pfs.Commit {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *JobInfo) GetJobStatus() []*JobStatus {
	if m != nil {
		return m.JobStatus
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*JobInfo) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), []interface{}) {
	return _JobInfo_OneofMarshaler, _JobInfo_OneofUnmarshaler, []interface{}{
		(*JobInfo_Transform)(nil),
		(*JobInfo_Pipeline)(nil),
	}
}

func _JobInfo_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*JobInfo)
	// spec
	switch x := m.Spec.(type) {
	case *JobInfo_Transform:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Transform); err != nil {
			return err
		}
	case *JobInfo_Pipeline:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Pipeline); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("JobInfo.Spec has unexpected type %T", x)
	}
	return nil
}

func _JobInfo_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*JobInfo)
	switch tag {
	case 2: // spec.transform
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Transform)
		err := b.DecodeMessage(msg)
		m.Spec = &JobInfo_Transform{msg}
		return true, err
	case 3: // spec.pipeline
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Pipeline)
		err := b.DecodeMessage(msg)
		m.Spec = &JobInfo_Pipeline{msg}
		return true, err
	default:
		return false, nil
	}
}

type JobInfos struct {
	JobInfo []*JobInfo `protobuf:"bytes,1,rep,name=job_info" json:"job_info,omitempty"`
}

func (m *JobInfos) Reset()         { *m = JobInfos{} }
func (m *JobInfos) String() string { return proto.CompactTextString(m) }
func (*JobInfos) ProtoMessage()    {}

func (m *JobInfos) GetJobInfo() []*JobInfo {
	if m != nil {
		return m.JobInfo
	}
	return nil
}

type Pipeline struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Pipeline) Reset()         { *m = Pipeline{} }
func (m *Pipeline) String() string { return proto.CompactTextString(m) }
func (*Pipeline) ProtoMessage()    {}

// TODO(pedge): add created at?
type PipelineInfo struct {
	Pipeline  *Pipeline  `protobuf:"bytes,1,opt,name=pipeline" json:"pipeline,omitempty"`
	Transform *Transform `protobuf:"bytes,2,opt,name=transform" json:"transform,omitempty"`
	Input     *pfs.Repo  `protobuf:"bytes,3,opt,name=input" json:"input,omitempty"`
	Output    *pfs.Repo  `protobuf:"bytes,4,opt,name=output" json:"output,omitempty"`
}

func (m *PipelineInfo) Reset()         { *m = PipelineInfo{} }
func (m *PipelineInfo) String() string { return proto.CompactTextString(m) }
func (*PipelineInfo) ProtoMessage()    {}

func (m *PipelineInfo) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *PipelineInfo) GetTransform() *Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *PipelineInfo) GetInput() *pfs.Repo {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *PipelineInfo) GetOutput() *pfs.Repo {
	if m != nil {
		return m.Output
	}
	return nil
}

type PipelineInfos struct {
	PipelineInfo []*PipelineInfo `protobuf:"bytes,1,rep,name=pipeline_info" json:"pipeline_info,omitempty"`
}

func (m *PipelineInfos) Reset()         { *m = PipelineInfos{} }
func (m *PipelineInfos) String() string { return proto.CompactTextString(m) }
func (*PipelineInfos) ProtoMessage()    {}

func (m *PipelineInfos) GetPipelineInfo() []*PipelineInfo {
	if m != nil {
		return m.PipelineInfo
	}
	return nil
}

type CreateJobRequest struct {
	// Types that are valid to be assigned to Spec:
	//	*CreateJobRequest_Transform
	//	*CreateJobRequest_Pipeline
	Spec         isCreateJobRequest_Spec `protobuf_oneof:"spec"`
	Input        *pfs.Commit             `protobuf:"bytes,3,opt,name=input" json:"input,omitempty"`
	OutputParent *pfs.Commit             `protobuf:"bytes,4,opt,name=output_parent" json:"output_parent,omitempty"`
}

func (m *CreateJobRequest) Reset()         { *m = CreateJobRequest{} }
func (m *CreateJobRequest) String() string { return proto.CompactTextString(m) }
func (*CreateJobRequest) ProtoMessage()    {}

type isCreateJobRequest_Spec interface {
	isCreateJobRequest_Spec()
}

type CreateJobRequest_Transform struct {
	Transform *Transform `protobuf:"bytes,1,opt,name=transform,oneof"`
}
type CreateJobRequest_Pipeline struct {
	Pipeline *Pipeline `protobuf:"bytes,2,opt,name=pipeline,oneof"`
}

func (*CreateJobRequest_Transform) isCreateJobRequest_Spec() {}
func (*CreateJobRequest_Pipeline) isCreateJobRequest_Spec()  {}

func (m *CreateJobRequest) GetSpec() isCreateJobRequest_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *CreateJobRequest) GetTransform() *Transform {
	if x, ok := m.GetSpec().(*CreateJobRequest_Transform); ok {
		return x.Transform
	}
	return nil
}

func (m *CreateJobRequest) GetPipeline() *Pipeline {
	if x, ok := m.GetSpec().(*CreateJobRequest_Pipeline); ok {
		return x.Pipeline
	}
	return nil
}

func (m *CreateJobRequest) GetInput() *pfs.Commit {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *CreateJobRequest) GetOutputParent() *pfs.Commit {
	if m != nil {
		return m.OutputParent
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CreateJobRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), []interface{}) {
	return _CreateJobRequest_OneofMarshaler, _CreateJobRequest_OneofUnmarshaler, []interface{}{
		(*CreateJobRequest_Transform)(nil),
		(*CreateJobRequest_Pipeline)(nil),
	}
}

func _CreateJobRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CreateJobRequest)
	// spec
	switch x := m.Spec.(type) {
	case *CreateJobRequest_Transform:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Transform); err != nil {
			return err
		}
	case *CreateJobRequest_Pipeline:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Pipeline); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CreateJobRequest.Spec has unexpected type %T", x)
	}
	return nil
}

func _CreateJobRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CreateJobRequest)
	switch tag {
	case 1: // spec.transform
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Transform)
		err := b.DecodeMessage(msg)
		m.Spec = &CreateJobRequest_Transform{msg}
		return true, err
	case 2: // spec.pipeline
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Pipeline)
		err := b.DecodeMessage(msg)
		m.Spec = &CreateJobRequest_Pipeline{msg}
		return true, err
	default:
		return false, nil
	}
}

type InspectJobRequest struct {
	Job *Job `protobuf:"bytes,1,opt,name=job" json:"job,omitempty"`
}

func (m *InspectJobRequest) Reset()         { *m = InspectJobRequest{} }
func (m *InspectJobRequest) String() string { return proto.CompactTextString(m) }
func (*InspectJobRequest) ProtoMessage()    {}

func (m *InspectJobRequest) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

type ListJobRequest struct {
	Pipeline *Pipeline `protobuf:"bytes,1,opt,name=pipeline" json:"pipeline,omitempty"`
}

func (m *ListJobRequest) Reset()         { *m = ListJobRequest{} }
func (m *ListJobRequest) String() string { return proto.CompactTextString(m) }
func (*ListJobRequest) ProtoMessage()    {}

func (m *ListJobRequest) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

type GetJobLogsRequest struct {
	Job          *Job         `protobuf:"bytes,1,opt,name=job" json:"job,omitempty"`
	OutputStream OutputStream `protobuf:"varint,2,opt,name=output_stream,enum=pachyderm.pps.OutputStream" json:"output_stream,omitempty"`
}

func (m *GetJobLogsRequest) Reset()         { *m = GetJobLogsRequest{} }
func (m *GetJobLogsRequest) String() string { return proto.CompactTextString(m) }
func (*GetJobLogsRequest) ProtoMessage()    {}

func (m *GetJobLogsRequest) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

type CreatePipelineRequest struct {
	Pipeline  *Pipeline  `protobuf:"bytes,1,opt,name=pipeline" json:"pipeline,omitempty"`
	Transform *Transform `protobuf:"bytes,2,opt,name=transform" json:"transform,omitempty"`
	Input     *pfs.Repo  `protobuf:"bytes,3,opt,name=input" json:"input,omitempty"`
	Output    *pfs.Repo  `protobuf:"bytes,4,opt,name=output" json:"output,omitempty"`
}

func (m *CreatePipelineRequest) Reset()         { *m = CreatePipelineRequest{} }
func (m *CreatePipelineRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePipelineRequest) ProtoMessage()    {}

func (m *CreatePipelineRequest) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *CreatePipelineRequest) GetTransform() *Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *CreatePipelineRequest) GetInput() *pfs.Repo {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *CreatePipelineRequest) GetOutput() *pfs.Repo {
	if m != nil {
		return m.Output
	}
	return nil
}

type InspectPipelineRequest struct {
	Pipeline *Pipeline `protobuf:"bytes,1,opt,name=pipeline" json:"pipeline,omitempty"`
}

func (m *InspectPipelineRequest) Reset()         { *m = InspectPipelineRequest{} }
func (m *InspectPipelineRequest) String() string { return proto.CompactTextString(m) }
func (*InspectPipelineRequest) ProtoMessage()    {}

func (m *InspectPipelineRequest) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

type ListPipelineRequest struct {
}

func (m *ListPipelineRequest) Reset()         { *m = ListPipelineRequest{} }
func (m *ListPipelineRequest) String() string { return proto.CompactTextString(m) }
func (*ListPipelineRequest) ProtoMessage()    {}

type DeletePipelineRequest struct {
	Pipeline *Pipeline `protobuf:"bytes,1,opt,name=pipeline" json:"pipeline,omitempty"`
}

func (m *DeletePipelineRequest) Reset()         { *m = DeletePipelineRequest{} }
func (m *DeletePipelineRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePipelineRequest) ProtoMessage()    {}

func (m *DeletePipelineRequest) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func init() {
	proto.RegisterEnum("pachyderm.pps.JobStatusType", JobStatusType_name, JobStatusType_value)
	proto.RegisterEnum("pachyderm.pps.OutputStream", OutputStream_name, OutputStream_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for JobAPI service

type JobAPIClient interface {
	CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*Job, error)
	InspectJob(ctx context.Context, in *InspectJobRequest, opts ...grpc.CallOption) (*JobInfo, error)
	ListJob(ctx context.Context, in *ListJobRequest, opts ...grpc.CallOption) (*JobInfos, error)
	GetJobLogs(ctx context.Context, in *GetJobLogsRequest, opts ...grpc.CallOption) (JobAPI_GetJobLogsClient, error)
}

type jobAPIClient struct {
	cc *grpc.ClientConn
}

func NewJobAPIClient(cc *grpc.ClientConn) JobAPIClient {
	return &jobAPIClient{cc}
}

func (c *jobAPIClient) CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := grpc.Invoke(ctx, "/pachyderm.pps.JobAPI/CreateJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobAPIClient) InspectJob(ctx context.Context, in *InspectJobRequest, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.JobAPI/InspectJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobAPIClient) ListJob(ctx context.Context, in *ListJobRequest, opts ...grpc.CallOption) (*JobInfos, error) {
	out := new(JobInfos)
	err := grpc.Invoke(ctx, "/pachyderm.pps.JobAPI/ListJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobAPIClient) GetJobLogs(ctx context.Context, in *GetJobLogsRequest, opts ...grpc.CallOption) (JobAPI_GetJobLogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_JobAPI_serviceDesc.Streams[0], c.cc, "/pachyderm.pps.JobAPI/GetJobLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &jobAPIGetJobLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type JobAPI_GetJobLogsClient interface {
	Recv() (*google_protobuf2.BytesValue, error)
	grpc.ClientStream
}

type jobAPIGetJobLogsClient struct {
	grpc.ClientStream
}

func (x *jobAPIGetJobLogsClient) Recv() (*google_protobuf2.BytesValue, error) {
	m := new(google_protobuf2.BytesValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for JobAPI service

type JobAPIServer interface {
	CreateJob(context.Context, *CreateJobRequest) (*Job, error)
	InspectJob(context.Context, *InspectJobRequest) (*JobInfo, error)
	ListJob(context.Context, *ListJobRequest) (*JobInfos, error)
	GetJobLogs(*GetJobLogsRequest, JobAPI_GetJobLogsServer) error
}

func RegisterJobAPIServer(s *grpc.Server, srv JobAPIServer) {
	s.RegisterService(&_JobAPI_serviceDesc, srv)
}

func _JobAPI_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(JobAPIServer).CreateJob(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _JobAPI_InspectJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(InspectJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(JobAPIServer).InspectJob(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _JobAPI_ListJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(JobAPIServer).ListJob(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _JobAPI_GetJobLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetJobLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(JobAPIServer).GetJobLogs(m, &jobAPIGetJobLogsServer{stream})
}

type JobAPI_GetJobLogsServer interface {
	Send(*google_protobuf2.BytesValue) error
	grpc.ServerStream
}

type jobAPIGetJobLogsServer struct {
	grpc.ServerStream
}

func (x *jobAPIGetJobLogsServer) Send(m *google_protobuf2.BytesValue) error {
	return x.ServerStream.SendMsg(m)
}

var _JobAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pachyderm.pps.JobAPI",
	HandlerType: (*JobAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJob",
			Handler:    _JobAPI_CreateJob_Handler,
		},
		{
			MethodName: "InspectJob",
			Handler:    _JobAPI_InspectJob_Handler,
		},
		{
			MethodName: "ListJob",
			Handler:    _JobAPI_ListJob_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetJobLogs",
			Handler:       _JobAPI_GetJobLogs_Handler,
			ServerStreams: true,
		},
	},
}

// Client API for PipelineAPI service

type PipelineAPIClient interface {
	CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	InspectPipeline(ctx context.Context, in *InspectPipelineRequest, opts ...grpc.CallOption) (*PipelineInfo, error)
	ListPipeline(ctx context.Context, in *ListPipelineRequest, opts ...grpc.CallOption) (*PipelineInfos, error)
	DeletePipeline(ctx context.Context, in *DeletePipelineRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type pipelineAPIClient struct {
	cc *grpc.ClientConn
}

func NewPipelineAPIClient(cc *grpc.ClientConn) PipelineAPIClient {
	return &pipelineAPIClient{cc}
}

func (c *pipelineAPIClient) CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.PipelineAPI/CreatePipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineAPIClient) InspectPipeline(ctx context.Context, in *InspectPipelineRequest, opts ...grpc.CallOption) (*PipelineInfo, error) {
	out := new(PipelineInfo)
	err := grpc.Invoke(ctx, "/pachyderm.pps.PipelineAPI/InspectPipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineAPIClient) ListPipeline(ctx context.Context, in *ListPipelineRequest, opts ...grpc.CallOption) (*PipelineInfos, error) {
	out := new(PipelineInfos)
	err := grpc.Invoke(ctx, "/pachyderm.pps.PipelineAPI/ListPipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineAPIClient) DeletePipeline(ctx context.Context, in *DeletePipelineRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pachyderm.pps.PipelineAPI/DeletePipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PipelineAPI service

type PipelineAPIServer interface {
	CreatePipeline(context.Context, *CreatePipelineRequest) (*google_protobuf.Empty, error)
	InspectPipeline(context.Context, *InspectPipelineRequest) (*PipelineInfo, error)
	ListPipeline(context.Context, *ListPipelineRequest) (*PipelineInfos, error)
	DeletePipeline(context.Context, *DeletePipelineRequest) (*google_protobuf.Empty, error)
}

func RegisterPipelineAPIServer(s *grpc.Server, srv PipelineAPIServer) {
	s.RegisterService(&_PipelineAPI_serviceDesc, srv)
}

func _PipelineAPI_CreatePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreatePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PipelineAPIServer).CreatePipeline(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _PipelineAPI_InspectPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(InspectPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PipelineAPIServer).InspectPipeline(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _PipelineAPI_ListPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PipelineAPIServer).ListPipeline(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _PipelineAPI_DeletePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DeletePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PipelineAPIServer).DeletePipeline(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _PipelineAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pachyderm.pps.PipelineAPI",
	HandlerType: (*PipelineAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePipeline",
			Handler:    _PipelineAPI_CreatePipeline_Handler,
		},
		{
			MethodName: "InspectPipeline",
			Handler:    _PipelineAPI_InspectPipeline_Handler,
		},
		{
			MethodName: "ListPipeline",
			Handler:    _PipelineAPI_ListPipeline_Handler,
		},
		{
			MethodName: "DeletePipeline",
			Handler:    _PipelineAPI_DeletePipeline_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
