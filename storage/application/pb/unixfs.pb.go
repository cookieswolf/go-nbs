// Code generated by protoc-gen-go. DO NOT EDIT.
// source: unixfs.proto

package unixfs_pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Data_DataType int32

const (
	Data_Raw       Data_DataType = 0
	Data_Directory Data_DataType = 1
	Data_File      Data_DataType = 2
	Data_Metadata  Data_DataType = 3
	Data_Symlink   Data_DataType = 4
	Data_HAMTShard Data_DataType = 5
)

var Data_DataType_name = map[int32]string{
	0: "Raw",
	1: "Directory",
	2: "File",
	3: "Metadata",
	4: "Symlink",
	5: "HAMTShard",
}

var Data_DataType_value = map[string]int32{
	"Raw":       0,
	"Directory": 1,
	"File":      2,
	"Metadata":  3,
	"Symlink":   4,
	"HAMTShard": 5,
}

func (x Data_DataType) Enum() *Data_DataType {
	p := new(Data_DataType)
	*p = x
	return p
}

func (x Data_DataType) String() string {
	return proto.EnumName(Data_DataType_name, int32(x))
}

func (x *Data_DataType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Data_DataType_value, data, "Data_DataType")
	if err != nil {
		return err
	}
	*x = Data_DataType(value)
	return nil
}

func (Data_DataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e2fd76cc44dfc7c3, []int{0, 0}
}

type Data struct {
	Type                 *Data_DataType `protobuf:"varint,1,req,name=Type,enum=unixfs.pb.Data_DataType" json:"Type,omitempty"`
	Data                 []byte         `protobuf:"bytes,2,opt,name=Data" json:"Data,omitempty"`
	Filesize             *uint64        `protobuf:"varint,3,opt,name=filesize" json:"filesize,omitempty"`
	Blocksizes           []uint64       `protobuf:"varint,4,rep,name=blocksizes" json:"blocksizes,omitempty"`
	HashType             *uint64        `protobuf:"varint,5,opt,name=hashType" json:"hashType,omitempty"`
	Fanout               *uint64        `protobuf:"varint,6,opt,name=fanout" json:"fanout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2fd76cc44dfc7c3, []int{0}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetType() Data_DataType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Data_Raw
}

func (m *Data) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Data) GetFilesize() uint64 {
	if m != nil && m.Filesize != nil {
		return *m.Filesize
	}
	return 0
}

func (m *Data) GetBlocksizes() []uint64 {
	if m != nil {
		return m.Blocksizes
	}
	return nil
}

func (m *Data) GetHashType() uint64 {
	if m != nil && m.HashType != nil {
		return *m.HashType
	}
	return 0
}

func (m *Data) GetFanout() uint64 {
	if m != nil && m.Fanout != nil {
		return *m.Fanout
	}
	return 0
}

type Metadata struct {
	MimeType             *string  `protobuf:"bytes,1,opt,name=MimeType" json:"MimeType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2fd76cc44dfc7c3, []int{1}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetMimeType() string {
	if m != nil && m.MimeType != nil {
		return *m.MimeType
	}
	return ""
}

func init() {
	proto.RegisterEnum("unixfs.pb.Data_DataType", Data_DataType_name, Data_DataType_value)
	proto.RegisterType((*Data)(nil), "unixfs.pb.Data")
	proto.RegisterType((*Metadata)(nil), "unixfs.pb.Metadata")
}

func init() { proto.RegisterFile("unixfs.proto", fileDescriptor_e2fd76cc44dfc7c3) }

var fileDescriptor_e2fd76cc44dfc7c3 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x6a, 0xeb, 0x30,
	0x18, 0x85, 0xaf, 0x6c, 0x25, 0xb1, 0xff, 0xeb, 0x16, 0xf1, 0x0f, 0x45, 0x74, 0x28, 0xc6, 0x43,
	0xd1, 0x50, 0x3c, 0xf4, 0x0d, 0x0a, 0xa1, 0x74, 0xf1, 0xa2, 0x84, 0xee, 0x4a, 0x22, 0x63, 0x11,
	0xc7, 0x0a, 0xb6, 0x42, 0xeb, 0x3e, 0x45, 0x1f, 0xb9, 0xc8, 0x8e, 0xdd, 0x2e, 0x82, 0x4f, 0xe7,
	0x7c, 0xe2, 0x20, 0x48, 0x2e, 0x8d, 0xf9, 0x2c, 0xbb, 0xfc, 0xdc, 0x5a, 0x67, 0x31, 0x9e, 0x68,
	0x97, 0x7d, 0x07, 0x40, 0xd7, 0xca, 0x29, 0x7c, 0x02, 0xba, 0xed, 0xcf, 0x9a, 0x93, 0x34, 0x10,
	0xb7, 0xcf, 0x3c, 0x9f, 0x2b, 0xb9, 0x8f, 0x87, 0xc3, 0xe7, 0x72, 0x68, 0x21, 0x8e, 0x16, 0x0f,
	0x52, 0x22, 0x12, 0x39, 0xbe, 0x70, 0x0f, 0x51, 0x69, 0x6a, 0xdd, 0x99, 0x2f, 0xcd, 0xc3, 0x94,
	0x08, 0x2a, 0x67, 0xc6, 0x07, 0x80, 0x5d, 0x6d, 0xf7, 0x47, 0x0f, 0x1d, 0xa7, 0x69, 0x28, 0xa8,
	0xfc, 0x73, 0xe3, 0xdd, 0x4a, 0x75, 0xd5, 0xb0, 0x60, 0x31, 0xba, 0x13, 0xe3, 0x1d, 0x2c, 0x4b,
	0xd5, 0xd8, 0x8b, 0xe3, 0xcb, 0x21, 0xb9, 0x52, 0xf6, 0x0e, 0xd1, 0xb4, 0x0a, 0x57, 0x10, 0x4a,
	0xf5, 0xc1, 0xfe, 0xe1, 0x0d, 0xc4, 0x6b, 0xd3, 0xea, 0xbd, 0xb3, 0x6d, 0xcf, 0x08, 0x46, 0x40,
	0x5f, 0x4d, 0xad, 0x59, 0x80, 0x09, 0x44, 0x85, 0x76, 0xea, 0xa0, 0x9c, 0x62, 0x21, 0xfe, 0x87,
	0xd5, 0xa6, 0x3f, 0xd5, 0xa6, 0x39, 0x32, 0xea, 0x9d, 0xb7, 0x97, 0x62, 0xbb, 0xa9, 0x54, 0x7b,
	0x60, 0x8b, 0xec, 0xf1, 0xb7, 0xe9, 0x77, 0x15, 0xe6, 0xa4, 0xaf, 0x3f, 0x43, 0x44, 0x2c, 0x67,
	0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xe9, 0xa0, 0x51, 0x10, 0x54, 0x01, 0x00, 0x00,
}
