// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ledgerData.proto

package bitswap_pb

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

type Ledger struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Received             float64  `protobuf:"fixed64,2,opt,name=received,proto3" json:"received,omitempty"`
	Score                float64  `protobuf:"fixed64,3,opt,name=score,proto3" json:"score,omitempty"`
	Sent                 float64  `protobuf:"fixed64,4,opt,name=sent,proto3" json:"sent,omitempty"`
	UpdateTime           string   `protobuf:"bytes,5,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ledger) Reset()         { *m = Ledger{} }
func (m *Ledger) String() string { return proto.CompactTextString(m) }
func (*Ledger) ProtoMessage()    {}
func (*Ledger) Descriptor() ([]byte, []int) {
	return fileDescriptor_d937399e46bdf5d2, []int{0}
}

func (m *Ledger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ledger.Unmarshal(m, b)
}
func (m *Ledger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ledger.Marshal(b, m, deterministic)
}
func (m *Ledger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ledger.Merge(m, src)
}
func (m *Ledger) XXX_Size() int {
	return xxx_messageInfo_Ledger.Size(m)
}
func (m *Ledger) XXX_DiscardUnknown() {
	xxx_messageInfo_Ledger.DiscardUnknown(m)
}

var xxx_messageInfo_Ledger proto.InternalMessageInfo

func (m *Ledger) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Ledger) GetReceived() float64 {
	if m != nil {
		return m.Received
	}
	return 0
}

func (m *Ledger) GetScore() float64 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Ledger) GetSent() float64 {
	if m != nil {
		return m.Sent
	}
	return 0
}

func (m *Ledger) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

type LedgerSet struct {
	Ledgers              []*Ledger `protobuf:"bytes,1,rep,name=ledgers,proto3" json:"ledgers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *LedgerSet) Reset()         { *m = LedgerSet{} }
func (m *LedgerSet) String() string { return proto.CompactTextString(m) }
func (*LedgerSet) ProtoMessage()    {}
func (*LedgerSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_d937399e46bdf5d2, []int{1}
}

func (m *LedgerSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LedgerSet.Unmarshal(m, b)
}
func (m *LedgerSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LedgerSet.Marshal(b, m, deterministic)
}
func (m *LedgerSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LedgerSet.Merge(m, src)
}
func (m *LedgerSet) XXX_Size() int {
	return xxx_messageInfo_LedgerSet.Size(m)
}
func (m *LedgerSet) XXX_DiscardUnknown() {
	xxx_messageInfo_LedgerSet.DiscardUnknown(m)
}

var xxx_messageInfo_LedgerSet proto.InternalMessageInfo

func (m *LedgerSet) GetLedgers() []*Ledger {
	if m != nil {
		return m.Ledgers
	}
	return nil
}

func init() {
	proto.RegisterType((*Ledger)(nil), "bitswap.pb.ledger")
	proto.RegisterType((*LedgerSet)(nil), "bitswap.pb.ledgerSet")
}

func init() { proto.RegisterFile("ledgerData.proto", fileDescriptor_d937399e46bdf5d2) }

var fileDescriptor_d937399e46bdf5d2 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x3b, 0x0a, 0x02, 0x31,
	0x10, 0x86, 0xc9, 0xbe, 0x74, 0x47, 0x10, 0x19, 0x2c, 0x82, 0x85, 0x2c, 0x5b, 0x6d, 0x21, 0x29,
	0xb4, 0xb2, 0xf7, 0x04, 0xab, 0x17, 0xc8, 0x6e, 0x06, 0x09, 0xa8, 0x09, 0x49, 0xd4, 0xce, 0xb3,
	0x0b, 0x09, 0x3e, 0xba, 0xf9, 0xfe, 0x6f, 0x98, 0x07, 0x2c, 0x2e, 0xa4, 0xce, 0xe4, 0x0e, 0x32,
	0x48, 0x61, 0x9d, 0x09, 0x06, 0x61, 0xd0, 0xc1, 0x3f, 0xa5, 0x15, 0x76, 0x68, 0x5f, 0x50, 0x25,
	0x8f, 0x73, 0xc8, 0xb4, 0xe2, 0xac, 0x61, 0x5d, 0xdd, 0x67, 0x5a, 0xe1, 0x0a, 0xa6, 0x8e, 0x46,
	0xd2, 0x0f, 0x52, 0x3c, 0x6b, 0x58, 0xc7, 0xfa, 0x2f, 0xe3, 0x12, 0x4a, 0x3f, 0x1a, 0x47, 0x3c,
	0x8f, 0x22, 0x01, 0x22, 0x14, 0x9e, 0x6e, 0x81, 0x17, 0x31, 0x8c, 0x35, 0xae, 0x01, 0xee, 0x56,
	0xc9, 0x40, 0x27, 0x7d, 0x25, 0x5e, 0xc6, 0xe9, 0x7f, 0x49, 0xbb, 0x87, 0x3a, 0xed, 0x3f, 0x52,
	0xc0, 0x0d, 0x4c, 0x12, 0x78, 0xce, 0x9a, 0xbc, 0x9b, 0x6d, 0x51, 0xfc, 0x4e, 0x15, 0x49, 0xf5,
	0x9f, 0x96, 0xa1, 0x8a, 0xdf, 0xec, 0xde, 0x01, 0x00, 0x00, 0xff, 0xff, 0x07, 0xa2, 0xd5, 0x36,
	0xe1, 0x00, 0x00, 0x00,
}
