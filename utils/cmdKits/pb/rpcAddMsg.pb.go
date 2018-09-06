// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpcAddMsg.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FileType int32

const (
	FileType_FILE       FileType = 0
	FileType_LARGEFILE  FileType = 1
	FileType_DIRECTORY  FileType = 2
	FileType_SYSTEMLINK FileType = 3
)

var FileType_name = map[int32]string{
	0: "FILE",
	1: "LARGEFILE",
	2: "DIRECTORY",
	3: "SYSTEMLINK",
}
var FileType_value = map[string]int32{
	"FILE":       0,
	"LARGEFILE":  1,
	"DIRECTORY":  2,
	"SYSTEMLINK": 3,
}

func (x FileType) String() string {
	return proto.EnumName(FileType_name, int32(x))
}
func (FileType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_rpcAddMsg_6fccfbf37a812ef5, []int{0}
}

type FileChunk struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileChunk) Reset()         { *m = FileChunk{} }
func (m *FileChunk) String() string { return proto.CompactTextString(m) }
func (*FileChunk) ProtoMessage()    {}
func (*FileChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpcAddMsg_6fccfbf37a812ef5, []int{0}
}
func (m *FileChunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileChunk.Unmarshal(m, b)
}
func (m *FileChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileChunk.Marshal(b, m, deterministic)
}
func (dst *FileChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileChunk.Merge(dst, src)
}
func (m *FileChunk) XXX_Size() int {
	return xxx_messageInfo_FileChunk.Size(m)
}
func (m *FileChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_FileChunk.DiscardUnknown(m)
}

var xxx_messageInfo_FileChunk proto.InternalMessageInfo

func (m *FileChunk) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type AddRequest struct {
	FileName             string   `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	FullPath             string   `protobuf:"bytes,2,opt,name=fullPath,proto3" json:"fullPath,omitempty"`
	FileSize             int64    `protobuf:"varint,3,opt,name=fileSize,proto3" json:"fileSize,omitempty"`
	FileType             FileType `protobuf:"varint,4,opt,name=fileType,proto3,enum=pb.FileType" json:"fileType,omitempty"`
	FileData             []byte   `protobuf:"bytes,5,opt,name=fileData,proto3" json:"fileData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpcAddMsg_6fccfbf37a812ef5, []int{1}
}
func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (dst *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(dst, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *AddRequest) GetFullPath() string {
	if m != nil {
		return m.FullPath
	}
	return ""
}

func (m *AddRequest) GetFileSize() int64 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

func (m *AddRequest) GetFileType() FileType {
	if m != nil {
		return m.FileType
	}
	return FileType_FILE
}

func (m *AddRequest) GetFileData() []byte {
	if m != nil {
		return m.FileData
	}
	return nil
}

type AddResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	SessionId            string   `protobuf:"bytes,2,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	DataReceive          int64    `protobuf:"varint,3,opt,name=dataReceive,proto3" json:"dataReceive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpcAddMsg_6fccfbf37a812ef5, []int{2}
}
func (m *AddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResponse.Unmarshal(m, b)
}
func (m *AddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResponse.Marshal(b, m, deterministic)
}
func (dst *AddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResponse.Merge(dst, src)
}
func (m *AddResponse) XXX_Size() int {
	return xxx_messageInfo_AddResponse.Size(m)
}
func (m *AddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddResponse proto.InternalMessageInfo

func (m *AddResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AddResponse) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *AddResponse) GetDataReceive() int64 {
	if m != nil {
		return m.DataReceive
	}
	return 0
}

func init() {
	proto.RegisterType((*FileChunk)(nil), "pb.FileChunk")
	proto.RegisterType((*AddRequest)(nil), "pb.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "pb.AddResponse")
	proto.RegisterEnum("pb.FileType", FileType_name, FileType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AddTaskClient is the client API for AddTask service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddTaskClient interface {
	AddFile(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	TransLargeFile(ctx context.Context, opts ...grpc.CallOption) (AddTask_TransLargeFileClient, error)
}

type addTaskClient struct {
	cc *grpc.ClientConn
}

func NewAddTaskClient(cc *grpc.ClientConn) AddTaskClient {
	return &addTaskClient{cc}
}

func (c *addTaskClient) AddFile(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/pb.AddTask/AddFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addTaskClient) TransLargeFile(ctx context.Context, opts ...grpc.CallOption) (AddTask_TransLargeFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AddTask_serviceDesc.Streams[0], "/pb.AddTask/TransLargeFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &addTaskTransLargeFileClient{stream}
	return x, nil
}

type AddTask_TransLargeFileClient interface {
	Send(*FileChunk) error
	CloseAndRecv() (*AddResponse, error)
	grpc.ClientStream
}

type addTaskTransLargeFileClient struct {
	grpc.ClientStream
}

func (x *addTaskTransLargeFileClient) Send(m *FileChunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *addTaskTransLargeFileClient) CloseAndRecv() (*AddResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AddResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AddTaskServer is the server API for AddTask service.
type AddTaskServer interface {
	AddFile(context.Context, *AddRequest) (*AddResponse, error)
	TransLargeFile(AddTask_TransLargeFileServer) error
}

func RegisterAddTaskServer(s *grpc.Server, srv AddTaskServer) {
	s.RegisterService(&_AddTask_serviceDesc, srv)
}

func _AddTask_AddFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddTaskServer).AddFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AddTask/AddFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddTaskServer).AddFile(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddTask_TransLargeFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AddTaskServer).TransLargeFile(&addTaskTransLargeFileServer{stream})
}

type AddTask_TransLargeFileServer interface {
	SendAndClose(*AddResponse) error
	Recv() (*FileChunk, error)
	grpc.ServerStream
}

type addTaskTransLargeFileServer struct {
	grpc.ServerStream
}

func (x *addTaskTransLargeFileServer) SendAndClose(m *AddResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *addTaskTransLargeFileServer) Recv() (*FileChunk, error) {
	m := new(FileChunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AddTask_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AddTask",
	HandlerType: (*AddTaskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFile",
			Handler:    _AddTask_AddFile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TransLargeFile",
			Handler:       _AddTask_TransLargeFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "rpcAddMsg.proto",
}

func init() { proto.RegisterFile("rpcAddMsg.proto", fileDescriptor_rpcAddMsg_6fccfbf37a812ef5) }

var fileDescriptor_rpcAddMsg_6fccfbf37a812ef5 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x6e, 0x9c, 0x30,
	0x10, 0x86, 0xe3, 0xdd, 0xb4, 0xd9, 0x9d, 0x64, 0xc9, 0xca, 0x27, 0x14, 0xb5, 0x12, 0x42, 0xaa,
	0x84, 0xaa, 0xca, 0xaa, 0xd2, 0xbe, 0x00, 0x9b, 0xb0, 0x15, 0x0a, 0x49, 0x23, 0x2f, 0x97, 0x1c,
	0x0d, 0x76, 0x09, 0x0a, 0x18, 0x8a, 0x4d, 0xa5, 0xf6, 0x81, 0xfa, 0x9c, 0x15, 0x06, 0xb3, 0x3d,
	0xe4, 0xc6, 0x37, 0x3f, 0x33, 0xfa, 0xe7, 0xf7, 0xc0, 0x65, 0xd7, 0xe6, 0x21, 0xe7, 0xf7, 0xaa,
	0x20, 0x6d, 0xd7, 0xe8, 0x06, 0x2f, 0xda, 0xcc, 0xff, 0x00, 0xeb, 0x7d, 0x59, 0x89, 0x9b, 0xe7,
	0x5e, 0xbe, 0x60, 0x17, 0xce, 0xf2, 0x46, 0x6a, 0x21, 0xb5, 0x8b, 0x3c, 0x14, 0x5c, 0x50, 0x8b,
	0xfe, 0x5f, 0x04, 0x10, 0x72, 0x4e, 0xc5, 0xcf, 0x5e, 0x28, 0x8d, 0xaf, 0x60, 0xf5, 0xa3, 0xac,
	0xc4, 0x03, 0xab, 0x85, 0xf9, 0x73, 0x4d, 0x67, 0x36, 0x5a, 0x5f, 0x55, 0x8f, 0x4c, 0x3f, 0xbb,
	0x8b, 0x49, 0x9b, 0xd8, 0xf6, 0x1d, 0xca, 0x3f, 0xc2, 0x5d, 0x7a, 0x28, 0x58, 0xd2, 0x99, 0x71,
	0x30, 0x6a, 0xe9, 0xef, 0x56, 0xb8, 0xa7, 0x1e, 0x0a, 0x9c, 0xeb, 0x0b, 0xd2, 0x66, 0x64, 0x3f,
	0xd5, 0xe8, 0xac, 0xda, 0x29, 0xb7, 0x4c, 0x33, 0xf7, 0x8d, 0xf1, 0x39, 0xb3, 0x5f, 0xc0, 0xb9,
	0xf1, 0xa9, 0xda, 0x46, 0x2a, 0x31, 0x6c, 0x54, 0x0b, 0xa5, 0x58, 0x61, 0x7d, 0x5a, 0xc4, 0xef,
	0x60, 0xad, 0x84, 0x52, 0x65, 0x23, 0x63, 0x3e, 0xf9, 0x3c, 0x16, 0xb0, 0x07, 0xe7, 0x9c, 0x69,
	0x46, 0x45, 0x2e, 0xca, 0x5f, 0xd6, 0xeb, 0xff, 0xa5, 0x8f, 0x3b, 0x58, 0x59, 0x6b, 0x78, 0x05,
	0xa7, 0xfb, 0x38, 0x89, 0xb6, 0x27, 0x78, 0x03, 0xeb, 0x24, 0xa4, 0xdf, 0x22, 0x83, 0x68, 0xc0,
	0xdb, 0x98, 0x46, 0x37, 0xe9, 0x77, 0xfa, 0xb4, 0x5d, 0x60, 0x07, 0xe0, 0xf0, 0x74, 0x48, 0xa3,
	0xfb, 0x24, 0x7e, 0xb8, 0xdb, 0x2e, 0xaf, 0x6b, 0x38, 0x0b, 0x39, 0x4f, 0x99, 0x7a, 0xc1, 0x9f,
	0xcc, 0xe7, 0x30, 0x11, 0x3b, 0xc3, 0xda, 0xc7, 0xb0, 0xaf, 0x2e, 0x67, 0x1e, 0x97, 0xf2, 0x4f,
	0xf0, 0x57, 0x70, 0xd2, 0x8e, 0x49, 0x95, 0xb0, 0xae, 0x10, 0xa6, 0x69, 0x63, 0xb3, 0x32, 0x2f,
	0xf9, 0x4a, 0x4f, 0x80, 0x76, 0x9f, 0xe1, 0x7d, 0xde, 0xd4, 0x44, 0x66, 0x8a, 0xc8, 0x86, 0x0b,
	0xd2, 0xeb, 0xb2, 0x52, 0x24, 0xaf, 0xf9, 0x5d, 0xa9, 0x15, 0x69, 0xb3, 0xdd, 0x86, 0x8e, 0x17,
	0x32, 0x46, 0xf4, 0x88, 0xb2, 0xb7, 0xe6, 0x50, 0xbe, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xe4,
	0x68, 0x5b, 0x04, 0x3b, 0x02, 0x00, 0x00,
}
