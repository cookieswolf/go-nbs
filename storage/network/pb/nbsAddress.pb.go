// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nbsAddress.proto

package net_pb

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NbsAddress struct {
	CanBeService         bool     `protobuf:"varint,10,opt,name=canBeService,proto3" json:"canBeService,omitempty"`
	PublicAddr           string   `protobuf:"bytes,20,opt,name=publicAddr,proto3" json:"publicAddr,omitempty"`
	PrivateAddr          string   `protobuf:"bytes,30,opt,name=privateAddr,proto3" json:"privateAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NbsAddress) Reset()         { *m = NbsAddress{} }
func (m *NbsAddress) String() string { return proto.CompactTextString(m) }
func (*NbsAddress) ProtoMessage()    {}
func (*NbsAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_50d06d17426aade4, []int{0}
}

func (m *NbsAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NbsAddress.Unmarshal(m, b)
}
func (m *NbsAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NbsAddress.Marshal(b, m, deterministic)
}
func (m *NbsAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NbsAddress.Merge(m, src)
}
func (m *NbsAddress) XXX_Size() int {
	return xxx_messageInfo_NbsAddress.Size(m)
}
func (m *NbsAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_NbsAddress.DiscardUnknown(m)
}

var xxx_messageInfo_NbsAddress proto.InternalMessageInfo

func (m *NbsAddress) GetCanBeService() bool {
	if m != nil {
		return m.CanBeService
	}
	return false
}

func (m *NbsAddress) GetPublicAddr() string {
	if m != nil {
		return m.PublicAddr
	}
	return ""
}

func (m *NbsAddress) GetPrivateAddr() string {
	if m != nil {
		return m.PrivateAddr
	}
	return ""
}

func init() {
	proto.RegisterType((*NbsAddress)(nil), "net.pb.NbsAddress")
}

func init() { proto.RegisterFile("nbsAddress.proto", fileDescriptor_50d06d17426aade4) }

var fileDescriptor_50d06d17426aade4 = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x4b, 0x2a, 0x76,
	0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcb, 0x4b,
	0x2d, 0xd1, 0x2b, 0x48, 0x52, 0x2a, 0xe2, 0xe2, 0xf2, 0x83, 0xcb, 0x09, 0x29, 0x71, 0xf1, 0x24,
	0x27, 0xe6, 0x39, 0xa5, 0x06, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x4a, 0x70, 0x29, 0x30, 0x6a,
	0x70, 0x04, 0xa1, 0x88, 0x09, 0xc9, 0x71, 0x71, 0x15, 0x94, 0x26, 0xe5, 0x64, 0x26, 0x83, 0x34,
	0x49, 0x88, 0x28, 0x30, 0x6a, 0x70, 0x06, 0x21, 0x89, 0x08, 0x29, 0x70, 0x71, 0x17, 0x14, 0x65,
	0x96, 0x25, 0x96, 0xa4, 0x82, 0x15, 0xc8, 0x81, 0x15, 0x20, 0x0b, 0x25, 0xb1, 0x81, 0x9d, 0x60,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x71, 0xa5, 0xd3, 0xba, 0x96, 0x00, 0x00, 0x00,
}
