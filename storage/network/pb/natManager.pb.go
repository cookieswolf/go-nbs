// Code generated by protoc-gen-go. DO NOT EDIT.
// source: natManager.proto

package nat_pb

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

type NatMsgType int32

const (
	NatMsgType_BootStrapReg NatMsgType = 0
	NatMsgType_KeepAlive    NatMsgType = 1
	NatMsgType_Connect      NatMsgType = 2
	NatMsgType_Probe        NatMsgType = 3
)

var NatMsgType_name = map[int32]string{
	0: "BootStrapReg",
	1: "KeepAlive",
	2: "Connect",
	3: "Probe",
}

var NatMsgType_value = map[string]int32{
	"BootStrapReg": 0,
	"KeepAlive":    1,
	"Connect":      2,
	"Probe":        3,
}

func (x NatMsgType) String() string {
	return proto.EnumName(NatMsgType_name, int32(x))
}

func (NatMsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{0}
}

type NatType int32

const (
	NatType_NoNatDevice    NatType = 0
	NatType_BehindNat      NatType = 1
	NatType_CanBeNatServer NatType = 2
	NatType_ToBeChecked    NatType = 3
)

var NatType_name = map[int32]string{
	0: "NoNatDevice",
	1: "BehindNat",
	2: "CanBeNatServer",
	3: "ToBeChecked",
}

var NatType_value = map[string]int32{
	"NoNatDevice":    0,
	"BehindNat":      1,
	"CanBeNatServer": 2,
	"ToBeChecked":    3,
}

func (x NatType) String() string {
	return proto.EnumName(NatType_name, int32(x))
}

func (NatType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{1}
}

type BootNatRegReq struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	PrivatePort          string   `protobuf:"bytes,3,opt,name=privatePort,proto3" json:"privatePort,omitempty"`
	Zone                 string   `protobuf:"bytes,4,opt,name=zone,proto3" json:"zone,omitempty"`
	PrivateIp            string   `protobuf:"bytes,2,opt,name=privateIp,proto3" json:"privateIp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BootNatRegReq) Reset()         { *m = BootNatRegReq{} }
func (m *BootNatRegReq) String() string { return proto.CompactTextString(m) }
func (*BootNatRegReq) ProtoMessage()    {}
func (*BootNatRegReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{0}
}

func (m *BootNatRegReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BootNatRegReq.Unmarshal(m, b)
}
func (m *BootNatRegReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BootNatRegReq.Marshal(b, m, deterministic)
}
func (m *BootNatRegReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootNatRegReq.Merge(m, src)
}
func (m *BootNatRegReq) XXX_Size() int {
	return xxx_messageInfo_BootNatRegReq.Size(m)
}
func (m *BootNatRegReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BootNatRegReq.DiscardUnknown(m)
}

var xxx_messageInfo_BootNatRegReq proto.InternalMessageInfo

func (m *BootNatRegReq) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *BootNatRegReq) GetPrivatePort() string {
	if m != nil {
		return m.PrivatePort
	}
	return ""
}

func (m *BootNatRegReq) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *BootNatRegReq) GetPrivateIp() string {
	if m != nil {
		return m.PrivateIp
	}
	return ""
}

type BootNatRegRes struct {
	NatType              NatType  `protobuf:"varint,1,opt,name=natType,proto3,enum=nat.pb.NatType" json:"natType,omitempty"`
	NodeId               string   `protobuf:"bytes,2,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	PublicIp             string   `protobuf:"bytes,3,opt,name=publicIp,proto3" json:"publicIp,omitempty"`
	PublicPort           string   `protobuf:"bytes,4,opt,name=publicPort,proto3" json:"publicPort,omitempty"`
	Zone                 string   `protobuf:"bytes,5,opt,name=zone,proto3" json:"zone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BootNatRegRes) Reset()         { *m = BootNatRegRes{} }
func (m *BootNatRegRes) String() string { return proto.CompactTextString(m) }
func (*BootNatRegRes) ProtoMessage()    {}
func (*BootNatRegRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{1}
}

func (m *BootNatRegRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BootNatRegRes.Unmarshal(m, b)
}
func (m *BootNatRegRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BootNatRegRes.Marshal(b, m, deterministic)
}
func (m *BootNatRegRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootNatRegRes.Merge(m, src)
}
func (m *BootNatRegRes) XXX_Size() int {
	return xxx_messageInfo_BootNatRegRes.Size(m)
}
func (m *BootNatRegRes) XXX_DiscardUnknown() {
	xxx_messageInfo_BootNatRegRes.DiscardUnknown(m)
}

var xxx_messageInfo_BootNatRegRes proto.InternalMessageInfo

func (m *BootNatRegRes) GetNatType() NatType {
	if m != nil {
		return m.NatType
	}
	return NatType_NoNatDevice
}

func (m *BootNatRegRes) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *BootNatRegRes) GetPublicIp() string {
	if m != nil {
		return m.PublicIp
	}
	return ""
}

func (m *BootNatRegRes) GetPublicPort() string {
	if m != nil {
		return m.PublicPort
	}
	return ""
}

func (m *BootNatRegRes) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

type NatConReq struct {
	FromPeerId           string   `protobuf:"bytes,1,opt,name=FromPeerId,proto3" json:"FromPeerId,omitempty"`
	ToPeerId             string   `protobuf:"bytes,2,opt,name=ToPeerId,proto3" json:"ToPeerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NatConReq) Reset()         { *m = NatConReq{} }
func (m *NatConReq) String() string { return proto.CompactTextString(m) }
func (*NatConReq) ProtoMessage()    {}
func (*NatConReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{2}
}

func (m *NatConReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NatConReq.Unmarshal(m, b)
}
func (m *NatConReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NatConReq.Marshal(b, m, deterministic)
}
func (m *NatConReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NatConReq.Merge(m, src)
}
func (m *NatConReq) XXX_Size() int {
	return xxx_messageInfo_NatConReq.Size(m)
}
func (m *NatConReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NatConReq.DiscardUnknown(m)
}

var xxx_messageInfo_NatConReq proto.InternalMessageInfo

func (m *NatConReq) GetFromPeerId() string {
	if m != nil {
		return m.FromPeerId
	}
	return ""
}

func (m *NatConReq) GetToPeerId() string {
	if m != nil {
		return m.ToPeerId
	}
	return ""
}

type NatConRes struct {
	PeerId               string   `protobuf:"bytes,2,opt,name=peerId,proto3" json:"peerId,omitempty"`
	PublicIp             string   `protobuf:"bytes,3,opt,name=publicIp,proto3" json:"publicIp,omitempty"`
	PublicPort           string   `protobuf:"bytes,4,opt,name=publicPort,proto3" json:"publicPort,omitempty"`
	PrivateIp            string   `protobuf:"bytes,5,opt,name=privateIp,proto3" json:"privateIp,omitempty"`
	PrivatePort          string   `protobuf:"bytes,6,opt,name=privatePort,proto3" json:"privatePort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NatConRes) Reset()         { *m = NatConRes{} }
func (m *NatConRes) String() string { return proto.CompactTextString(m) }
func (*NatConRes) ProtoMessage()    {}
func (*NatConRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{3}
}

func (m *NatConRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NatConRes.Unmarshal(m, b)
}
func (m *NatConRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NatConRes.Marshal(b, m, deterministic)
}
func (m *NatConRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NatConRes.Merge(m, src)
}
func (m *NatConRes) XXX_Size() int {
	return xxx_messageInfo_NatConRes.Size(m)
}
func (m *NatConRes) XXX_DiscardUnknown() {
	xxx_messageInfo_NatConRes.DiscardUnknown(m)
}

var xxx_messageInfo_NatConRes proto.InternalMessageInfo

func (m *NatConRes) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *NatConRes) GetPublicIp() string {
	if m != nil {
		return m.PublicIp
	}
	return ""
}

func (m *NatConRes) GetPublicPort() string {
	if m != nil {
		return m.PublicPort
	}
	return ""
}

func (m *NatConRes) GetPrivateIp() string {
	if m != nil {
		return m.PrivateIp
	}
	return ""
}

func (m *NatConRes) GetPrivatePort() string {
	if m != nil {
		return m.PrivatePort
	}
	return ""
}

type NatKeepAlive struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NatKeepAlive) Reset()         { *m = NatKeepAlive{} }
func (m *NatKeepAlive) String() string { return proto.CompactTextString(m) }
func (*NatKeepAlive) ProtoMessage()    {}
func (*NatKeepAlive) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{4}
}

func (m *NatKeepAlive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NatKeepAlive.Unmarshal(m, b)
}
func (m *NatKeepAlive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NatKeepAlive.Marshal(b, m, deterministic)
}
func (m *NatKeepAlive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NatKeepAlive.Merge(m, src)
}
func (m *NatKeepAlive) XXX_Size() int {
	return xxx_messageInfo_NatKeepAlive.Size(m)
}
func (m *NatKeepAlive) XXX_DiscardUnknown() {
	xxx_messageInfo_NatKeepAlive.DiscardUnknown(m)
}

var xxx_messageInfo_NatKeepAlive proto.InternalMessageInfo

func (m *NatKeepAlive) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type NatProbe struct {
	Ping                 string   `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
	Pong                 string   `protobuf:"bytes,2,opt,name=pong,proto3" json:"pong,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NatProbe) Reset()         { *m = NatProbe{} }
func (m *NatProbe) String() string { return proto.CompactTextString(m) }
func (*NatProbe) ProtoMessage()    {}
func (*NatProbe) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{5}
}

func (m *NatProbe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NatProbe.Unmarshal(m, b)
}
func (m *NatProbe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NatProbe.Marshal(b, m, deterministic)
}
func (m *NatProbe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NatProbe.Merge(m, src)
}
func (m *NatProbe) XXX_Size() int {
	return xxx_messageInfo_NatProbe.Size(m)
}
func (m *NatProbe) XXX_DiscardUnknown() {
	xxx_messageInfo_NatProbe.DiscardUnknown(m)
}

var xxx_messageInfo_NatProbe proto.InternalMessageInfo

func (m *NatProbe) GetPing() string {
	if m != nil {
		return m.Ping
	}
	return ""
}

func (m *NatProbe) GetPong() string {
	if m != nil {
		return m.Pong
	}
	return ""
}

type NatRequest struct {
	MsgType              NatMsgType     `protobuf:"varint,1,opt,name=msgType,proto3,enum=nat.pb.NatMsgType" json:"msgType,omitempty"`
	BootRegReq           *BootNatRegReq `protobuf:"bytes,2,opt,name=BootRegReq,proto3" json:"BootRegReq,omitempty"`
	ConnReq              *NatConReq     `protobuf:"bytes,3,opt,name=connReq,proto3" json:"connReq,omitempty"`
	KeepAlive            *NatKeepAlive  `protobuf:"bytes,4,opt,name=keepAlive,proto3" json:"keepAlive,omitempty"`
	ProbeReq             *NatProbe      `protobuf:"bytes,5,opt,name=probeReq,proto3" json:"probeReq,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *NatRequest) Reset()         { *m = NatRequest{} }
func (m *NatRequest) String() string { return proto.CompactTextString(m) }
func (*NatRequest) ProtoMessage()    {}
func (*NatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{6}
}

func (m *NatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NatRequest.Unmarshal(m, b)
}
func (m *NatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NatRequest.Marshal(b, m, deterministic)
}
func (m *NatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NatRequest.Merge(m, src)
}
func (m *NatRequest) XXX_Size() int {
	return xxx_messageInfo_NatRequest.Size(m)
}
func (m *NatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NatRequest proto.InternalMessageInfo

func (m *NatRequest) GetMsgType() NatMsgType {
	if m != nil {
		return m.MsgType
	}
	return NatMsgType_BootStrapReg
}

func (m *NatRequest) GetBootRegReq() *BootNatRegReq {
	if m != nil {
		return m.BootRegReq
	}
	return nil
}

func (m *NatRequest) GetConnReq() *NatConReq {
	if m != nil {
		return m.ConnReq
	}
	return nil
}

func (m *NatRequest) GetKeepAlive() *NatKeepAlive {
	if m != nil {
		return m.KeepAlive
	}
	return nil
}

func (m *NatRequest) GetProbeReq() *NatProbe {
	if m != nil {
		return m.ProbeReq
	}
	return nil
}

type Response struct {
	MsgType              NatMsgType     `protobuf:"varint,1,opt,name=msgType,proto3,enum=nat.pb.NatMsgType" json:"msgType,omitempty"`
	BootRegRes           *BootNatRegRes `protobuf:"bytes,2,opt,name=BootRegRes,proto3" json:"BootRegRes,omitempty"`
	ConnRes              *NatConRes     `protobuf:"bytes,3,opt,name=connRes,proto3" json:"connRes,omitempty"`
	KeepAlive            *NatKeepAlive  `protobuf:"bytes,4,opt,name=keepAlive,proto3" json:"keepAlive,omitempty"`
	ProbeRes             *NatProbe      `protobuf:"bytes,5,opt,name=probeRes,proto3" json:"probeRes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{7}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsgType() NatMsgType {
	if m != nil {
		return m.MsgType
	}
	return NatMsgType_BootStrapReg
}

func (m *Response) GetBootRegRes() *BootNatRegRes {
	if m != nil {
		return m.BootRegRes
	}
	return nil
}

func (m *Response) GetConnRes() *NatConRes {
	if m != nil {
		return m.ConnRes
	}
	return nil
}

func (m *Response) GetKeepAlive() *NatKeepAlive {
	if m != nil {
		return m.KeepAlive
	}
	return nil
}

func (m *Response) GetProbeRes() *NatProbe {
	if m != nil {
		return m.ProbeRes
	}
	return nil
}

func init() {
	proto.RegisterEnum("nat.pb.NatMsgType", NatMsgType_name, NatMsgType_value)
	proto.RegisterEnum("nat.pb.NatType", NatType_name, NatType_value)
	proto.RegisterType((*BootNatRegReq)(nil), "nat.pb.BootNatRegReq")
	proto.RegisterType((*BootNatRegRes)(nil), "nat.pb.BootNatRegRes")
	proto.RegisterType((*NatConReq)(nil), "nat.pb.NatConReq")
	proto.RegisterType((*NatConRes)(nil), "nat.pb.NatConRes")
	proto.RegisterType((*NatKeepAlive)(nil), "nat.pb.NatKeepAlive")
	proto.RegisterType((*NatProbe)(nil), "nat.pb.NatProbe")
	proto.RegisterType((*NatRequest)(nil), "nat.pb.NatRequest")
	proto.RegisterType((*Response)(nil), "nat.pb.response")
}

func init() { proto.RegisterFile("natManager.proto", fileDescriptor_2f5cc5a30f7d4657) }

var fileDescriptor_2f5cc5a30f7d4657 = []byte{
	// 546 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xdb, 0x6e, 0x13, 0x31,
	0x10, 0x86, 0xd9, 0xa4, 0x39, 0xec, 0x6c, 0x0f, 0xcb, 0x08, 0x50, 0x84, 0x50, 0x55, 0xed, 0x05,
	0x82, 0x52, 0xe5, 0x62, 0x11, 0x0f, 0x40, 0xc2, 0x41, 0x11, 0xea, 0x2a, 0xda, 0xe6, 0x05, 0x9c,
	0x64, 0xb4, 0x5d, 0xb5, 0xb5, 0x5d, 0xdb, 0x8d, 0x04, 0xdc, 0xf0, 0x26, 0x48, 0x3c, 0x28, 0x42,
	0xeb, 0x3d, 0x39, 0x45, 0x85, 0x8b, 0xdc, 0xd9, 0x33, 0xbf, 0x3d, 0xff, 0x7c, 0x63, 0x19, 0x42,
	0xce, 0xcc, 0x39, 0xe3, 0x2c, 0x23, 0x35, 0x96, 0x4a, 0x18, 0x81, 0x7d, 0xce, 0xcc, 0x58, 0x2e,
	0xa3, 0xef, 0x70, 0x30, 0x11, 0xc2, 0x24, 0xcc, 0xa4, 0x94, 0xa5, 0x74, 0x8b, 0xcf, 0xa0, 0xcf,
	0xc5, 0x9a, 0x66, 0xeb, 0x91, 0x77, 0xe2, 0xbd, 0xf2, 0xd3, 0x6a, 0x87, 0x27, 0x10, 0x48, 0x95,
	0x6f, 0x98, 0xa1, 0xb9, 0x50, 0x66, 0xd4, 0xb5, 0x49, 0x37, 0x84, 0x08, 0x7b, 0xdf, 0x04, 0xa7,
	0xd1, 0x9e, 0x4d, 0xd9, 0x35, 0xbe, 0x00, 0xbf, 0x92, 0xcc, 0xe4, 0xa8, 0x63, 0x13, 0x6d, 0x20,
	0xfa, 0xe5, 0x6d, 0x57, 0xd7, 0xf8, 0x1a, 0x06, 0x9c, 0x99, 0xc5, 0x57, 0x49, 0xb6, 0xfc, 0x61,
	0x7c, 0x34, 0x2e, 0x8d, 0x8e, 0x93, 0x32, 0x9c, 0xd6, 0x79, 0xc7, 0x68, 0x67, 0xcb, 0xe8, 0x73,
	0x18, 0xca, 0xbb, 0xe5, 0x75, 0xbe, 0x9a, 0xc9, 0xca, 0x65, 0xb3, 0xc7, 0x63, 0x80, 0x72, 0x6d,
	0x7b, 0x28, 0x8d, 0x3a, 0x91, 0xa6, 0x85, 0x5e, 0xdb, 0x42, 0xf4, 0x19, 0xfc, 0x84, 0x99, 0xa9,
	0xe0, 0x05, 0x9d, 0x63, 0x80, 0x4f, 0x4a, 0xdc, 0xcc, 0x89, 0x54, 0x43, 0xc8, 0x89, 0x14, 0xc5,
	0x17, 0xa2, 0xca, 0x96, 0xb6, 0x9a, 0x7d, 0xf4, 0xd3, 0x6b, 0x6f, 0xd2, 0x85, 0x7d, 0xe9, 0xea,
	0xaa, 0xdd, 0x4e, 0xf6, 0xb7, 0x68, 0xf7, 0xee, 0xd1, 0xbe, 0x3f, 0xc1, 0xfe, 0x5f, 0x13, 0x8c,
	0x5e, 0xc2, 0x7e, 0xc2, 0xcc, 0x17, 0x22, 0xf9, 0xfe, 0x3a, 0xdf, 0xd0, 0x43, 0x6f, 0x21, 0x8a,
	0x61, 0x98, 0x30, 0x33, 0x57, 0x62, 0x49, 0x05, 0x32, 0x99, 0xf3, 0xac, 0x52, 0xd8, 0xb5, 0x8d,
	0x09, 0x9e, 0x55, 0x9d, 0xd9, 0x75, 0xf4, 0xa3, 0x03, 0x60, 0xe7, 0x7c, 0x7b, 0x47, 0xda, 0xe0,
	0x19, 0x0c, 0x6e, 0x74, 0xe6, 0x0c, 0x1a, 0x9d, 0x41, 0x9f, 0x97, 0x99, 0xb4, 0x96, 0xe0, 0x3b,
	0x80, 0xe2, 0x9d, 0x94, 0x4f, 0xd4, 0x5e, 0x1b, 0xc4, 0x4f, 0xeb, 0x03, 0x5b, 0xef, 0x37, 0x75,
	0x84, 0xf8, 0x06, 0x06, 0x2b, 0xc1, 0x8b, 0xc1, 0x59, 0x94, 0x41, 0xfc, 0xd8, 0x29, 0x52, 0x4e,
	0x34, 0xad, 0x15, 0x18, 0x83, 0x7f, 0x55, 0x77, 0x6e, 0xd9, 0x06, 0xf1, 0x13, 0x47, 0xde, 0x50,
	0x49, 0x5b, 0x19, 0x9e, 0xc1, 0x50, 0x16, 0x14, 0x8a, 0x0a, 0x3d, 0x7b, 0x24, 0x74, 0x8e, 0x58,
	0x40, 0x69, 0xa3, 0x88, 0x7e, 0x7b, 0x30, 0x54, 0xa4, 0xa5, 0xe0, 0x9a, 0x76, 0x00, 0xa0, 0xff,
	0x05, 0x40, 0x3b, 0x00, 0x74, 0x0b, 0x40, 0x3f, 0x04, 0x40, 0xd7, 0x00, 0xf4, 0x8e, 0x00, 0xf4,
	0x7f, 0x01, 0xe8, 0xd3, 0x8f, 0xf6, 0x09, 0x54, 0xcd, 0x61, 0x08, 0xfb, 0x85, 0xd5, 0x0b, 0xa3,
	0x98, 0x4c, 0x29, 0x0b, 0x1f, 0xe1, 0x01, 0xf8, 0x4d, 0x95, 0xd0, 0xc3, 0x00, 0x06, 0x53, 0xc1,
	0x39, 0xad, 0x4c, 0xd8, 0x41, 0x1f, 0x7a, 0xf6, 0xba, 0xb0, 0x7b, 0x9a, 0xc0, 0xa0, 0xfa, 0x0d,
	0xf0, 0x08, 0x82, 0x44, 0x24, 0xcc, 0x7c, 0xa0, 0x4d, 0xbe, 0xa2, 0xf2, 0x8a, 0x09, 0x5d, 0xe6,
	0x7c, 0x9d, 0x30, 0x13, 0x7a, 0x88, 0x70, 0x38, 0x65, 0x7c, 0x42, 0x09, 0x33, 0x17, 0xa4, 0x36,
	0xa4, 0xc2, 0x4e, 0x71, 0x66, 0x21, 0x26, 0x34, 0xbd, 0xa4, 0xd5, 0x15, 0xad, 0xc3, 0xee, 0xb2,
	0x6f, 0xbf, 0xc4, 0xb7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x3a, 0x36, 0x78, 0x26, 0x05,
	0x00, 0x00,
}