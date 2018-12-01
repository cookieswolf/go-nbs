// Code generated by protoc-gen-go. DO NOT EDIT.
// source: natManager.proto

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

type NatType int32

const (
	NatType_UnknownRES     NatType = 0
	NatType_NoNatDevice    NatType = 1
	NatType_BehindNat      NatType = 2
	NatType_CanBeNatServer NatType = 3
	NatType_ToBeChecked    NatType = 4
)

var NatType_name = map[int32]string{
	0: "UnknownRES",
	1: "NoNatDevice",
	2: "BehindNat",
	3: "CanBeNatServer",
	4: "ToBeChecked",
}

var NatType_value = map[string]int32{
	"UnknownRES":     0,
	"NoNatDevice":    1,
	"BehindNat":      2,
	"CanBeNatServer": 3,
	"ToBeChecked":    4,
}

func (x NatType) String() string {
	return proto.EnumName(NatType_name, int32(x))
}

func (NatType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{0}
}

type BootReg struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	PrivatePort          string   `protobuf:"bytes,3,opt,name=privatePort,proto3" json:"privatePort,omitempty"`
	PrivateIp            string   `protobuf:"bytes,2,opt,name=privateIp,proto3" json:"privateIp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BootReg) Reset()         { *m = BootReg{} }
func (m *BootReg) String() string { return proto.CompactTextString(m) }
func (*BootReg) ProtoMessage()    {}
func (*BootReg) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{0}
}

func (m *BootReg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BootReg.Unmarshal(m, b)
}
func (m *BootReg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BootReg.Marshal(b, m, deterministic)
}
func (m *BootReg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootReg.Merge(m, src)
}
func (m *BootReg) XXX_Size() int {
	return xxx_messageInfo_BootReg.Size(m)
}
func (m *BootReg) XXX_DiscardUnknown() {
	xxx_messageInfo_BootReg.DiscardUnknown(m)
}

var xxx_messageInfo_BootReg proto.InternalMessageInfo

func (m *BootReg) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *BootReg) GetPrivatePort() string {
	if m != nil {
		return m.PrivatePort
	}
	return ""
}

func (m *BootReg) GetPrivateIp() string {
	if m != nil {
		return m.PrivateIp
	}
	return ""
}

type BootRegAck struct {
	NatType              NatType  `protobuf:"varint,1,opt,name=natType,proto3,enum=net.pb.NatType" json:"natType,omitempty"`
	NodeId               string   `protobuf:"bytes,2,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	PublicIp             string   `protobuf:"bytes,3,opt,name=publicIp,proto3" json:"publicIp,omitempty"`
	PublicPort           string   `protobuf:"bytes,4,opt,name=publicPort,proto3" json:"publicPort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BootRegAck) Reset()         { *m = BootRegAck{} }
func (m *BootRegAck) String() string { return proto.CompactTextString(m) }
func (*BootRegAck) ProtoMessage()    {}
func (*BootRegAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{1}
}

func (m *BootRegAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BootRegAck.Unmarshal(m, b)
}
func (m *BootRegAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BootRegAck.Marshal(b, m, deterministic)
}
func (m *BootRegAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootRegAck.Merge(m, src)
}
func (m *BootRegAck) XXX_Size() int {
	return xxx_messageInfo_BootRegAck.Size(m)
}
func (m *BootRegAck) XXX_DiscardUnknown() {
	xxx_messageInfo_BootRegAck.DiscardUnknown(m)
}

var xxx_messageInfo_BootRegAck proto.InternalMessageInfo

func (m *BootRegAck) GetNatType() NatType {
	if m != nil {
		return m.NatType
	}
	return NatType_UnknownRES
}

func (m *BootRegAck) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *BootRegAck) GetPublicIp() string {
	if m != nil {
		return m.PublicIp
	}
	return ""
}

func (m *BootRegAck) GetPublicPort() string {
	if m != nil {
		return m.PublicPort
	}
	return ""
}

type NatConnect struct {
	FromAddr             *NbsAddr `protobuf:"bytes,1,opt,name=fromAddr,proto3" json:"fromAddr,omitempty"`
	ToAddr               *NbsAddr `protobuf:"bytes,2,opt,name=toAddr,proto3" json:"toAddr,omitempty"`
	SessionId            string   `protobuf:"bytes,3,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	ToPort               int32    `protobuf:"varint,4,opt,name=toPort,proto3" json:"toPort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NatConnect) Reset()         { *m = NatConnect{} }
func (m *NatConnect) String() string { return proto.CompactTextString(m) }
func (*NatConnect) ProtoMessage()    {}
func (*NatConnect) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{2}
}

func (m *NatConnect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NatConnect.Unmarshal(m, b)
}
func (m *NatConnect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NatConnect.Marshal(b, m, deterministic)
}
func (m *NatConnect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NatConnect.Merge(m, src)
}
func (m *NatConnect) XXX_Size() int {
	return xxx_messageInfo_NatConnect.Size(m)
}
func (m *NatConnect) XXX_DiscardUnknown() {
	xxx_messageInfo_NatConnect.DiscardUnknown(m)
}

var xxx_messageInfo_NatConnect proto.InternalMessageInfo

func (m *NatConnect) GetFromAddr() *NbsAddr {
	if m != nil {
		return m.FromAddr
	}
	return nil
}

func (m *NatConnect) GetToAddr() *NbsAddr {
	if m != nil {
		return m.ToAddr
	}
	return nil
}

func (m *NatConnect) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *NatConnect) GetToPort() int32 {
	if m != nil {
		return m.ToPort
	}
	return 0
}

type NatKeepAlive struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	LAddr                string   `protobuf:"bytes,2,opt,name=lAddr,proto3" json:"lAddr,omitempty"`
	PubIP                string   `protobuf:"bytes,3,opt,name=PubIP,proto3" json:"PubIP,omitempty"`
	PubPort              int32    `protobuf:"varint,4,opt,name=PubPort,proto3" json:"PubPort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NatKeepAlive) Reset()         { *m = NatKeepAlive{} }
func (m *NatKeepAlive) String() string { return proto.CompactTextString(m) }
func (*NatKeepAlive) ProtoMessage()    {}
func (*NatKeepAlive) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{3}
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

func (m *NatKeepAlive) GetLAddr() string {
	if m != nil {
		return m.LAddr
	}
	return ""
}

func (m *NatKeepAlive) GetPubIP() string {
	if m != nil {
		return m.PubIP
	}
	return ""
}

func (m *NatKeepAlive) GetPubPort() int32 {
	if m != nil {
		return m.PubPort
	}
	return 0
}

type NatPing struct {
	Ping                 string   `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
	Pong                 string   `protobuf:"bytes,2,opt,name=pong,proto3" json:"pong,omitempty"`
	Nonce                string   `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	TTL                  int32    `protobuf:"varint,4,opt,name=TTL,proto3" json:"TTL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NatPing) Reset()         { *m = NatPing{} }
func (m *NatPing) String() string { return proto.CompactTextString(m) }
func (*NatPing) ProtoMessage()    {}
func (*NatPing) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{4}
}

func (m *NatPing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NatPing.Unmarshal(m, b)
}
func (m *NatPing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NatPing.Marshal(b, m, deterministic)
}
func (m *NatPing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NatPing.Merge(m, src)
}
func (m *NatPing) XXX_Size() int {
	return xxx_messageInfo_NatPing.Size(m)
}
func (m *NatPing) XXX_DiscardUnknown() {
	xxx_messageInfo_NatPing.DiscardUnknown(m)
}

var xxx_messageInfo_NatPing proto.InternalMessageInfo

func (m *NatPing) GetPing() string {
	if m != nil {
		return m.Ping
	}
	return ""
}

func (m *NatPing) GetPong() string {
	if m != nil {
		return m.Pong
	}
	return ""
}

func (m *NatPing) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *NatPing) GetTTL() int32 {
	if m != nil {
		return m.TTL
	}
	return 0
}

type HoleDig struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	NetworkType          int32    `protobuf:"varint,2,opt,name=networkType,proto3" json:"networkType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HoleDig) Reset()         { *m = HoleDig{} }
func (m *HoleDig) String() string { return proto.CompactTextString(m) }
func (*HoleDig) ProtoMessage()    {}
func (*HoleDig) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{5}
}

func (m *HoleDig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HoleDig.Unmarshal(m, b)
}
func (m *HoleDig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HoleDig.Marshal(b, m, deterministic)
}
func (m *HoleDig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HoleDig.Merge(m, src)
}
func (m *HoleDig) XXX_Size() int {
	return xxx_messageInfo_HoleDig.Size(m)
}
func (m *HoleDig) XXX_DiscardUnknown() {
	xxx_messageInfo_HoleDig.DiscardUnknown(m)
}

var xxx_messageInfo_HoleDig proto.InternalMessageInfo

func (m *HoleDig) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *HoleDig) GetNetworkType() int32 {
	if m != nil {
		return m.NetworkType
	}
	return 0
}

type HolePayLoad struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	TargetPort           int32    `protobuf:"varint,2,opt,name=targetPort,proto3" json:"targetPort,omitempty"`
	PayLoad              []byte   `protobuf:"bytes,3,opt,name=payLoad,proto3" json:"payLoad,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HolePayLoad) Reset()         { *m = HolePayLoad{} }
func (m *HolePayLoad) String() string { return proto.CompactTextString(m) }
func (*HolePayLoad) ProtoMessage()    {}
func (*HolePayLoad) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{6}
}

func (m *HolePayLoad) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HolePayLoad.Unmarshal(m, b)
}
func (m *HolePayLoad) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HolePayLoad.Marshal(b, m, deterministic)
}
func (m *HolePayLoad) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HolePayLoad.Merge(m, src)
}
func (m *HolePayLoad) XXX_Size() int {
	return xxx_messageInfo_HolePayLoad.Size(m)
}
func (m *HolePayLoad) XXX_DiscardUnknown() {
	xxx_messageInfo_HolePayLoad.DiscardUnknown(m)
}

var xxx_messageInfo_HolePayLoad proto.InternalMessageInfo

func (m *HolePayLoad) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *HolePayLoad) GetTargetPort() int32 {
	if m != nil {
		return m.TargetPort
	}
	return 0
}

func (m *HolePayLoad) GetPayLoad() []byte {
	if m != nil {
		return m.PayLoad
	}
	return nil
}

type ErrorNotify struct {
	ErrMsg               string   `protobuf:"bytes,1,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorNotify) Reset()         { *m = ErrorNotify{} }
func (m *ErrorNotify) String() string { return proto.CompactTextString(m) }
func (*ErrorNotify) ProtoMessage()    {}
func (*ErrorNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{7}
}

func (m *ErrorNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorNotify.Unmarshal(m, b)
}
func (m *ErrorNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorNotify.Marshal(b, m, deterministic)
}
func (m *ErrorNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorNotify.Merge(m, src)
}
func (m *ErrorNotify) XXX_Size() int {
	return xxx_messageInfo_ErrorNotify.Size(m)
}
func (m *ErrorNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorNotify.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorNotify proto.InternalMessageInfo

func (m *ErrorNotify) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type ReverseInvite struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	PubIp                string   `protobuf:"bytes,2,opt,name=pubIp,proto3" json:"pubIp,omitempty"`
	ToPort               int32    `protobuf:"varint,3,opt,name=toPort,proto3" json:"toPort,omitempty"`
	PeerId               string   `protobuf:"bytes,4,opt,name=peerId,proto3" json:"peerId,omitempty"`
	FromPort             string   `protobuf:"bytes,5,opt,name=fromPort,proto3" json:"fromPort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReverseInvite) Reset()         { *m = ReverseInvite{} }
func (m *ReverseInvite) String() string { return proto.CompactTextString(m) }
func (*ReverseInvite) ProtoMessage()    {}
func (*ReverseInvite) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{8}
}

func (m *ReverseInvite) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReverseInvite.Unmarshal(m, b)
}
func (m *ReverseInvite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReverseInvite.Marshal(b, m, deterministic)
}
func (m *ReverseInvite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReverseInvite.Merge(m, src)
}
func (m *ReverseInvite) XXX_Size() int {
	return xxx_messageInfo_ReverseInvite.Size(m)
}
func (m *ReverseInvite) XXX_DiscardUnknown() {
	xxx_messageInfo_ReverseInvite.DiscardUnknown(m)
}

var xxx_messageInfo_ReverseInvite proto.InternalMessageInfo

func (m *ReverseInvite) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *ReverseInvite) GetPubIp() string {
	if m != nil {
		return m.PubIp
	}
	return ""
}

func (m *ReverseInvite) GetToPort() int32 {
	if m != nil {
		return m.ToPort
	}
	return 0
}

func (m *ReverseInvite) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *ReverseInvite) GetFromPort() string {
	if m != nil {
		return m.FromPort
	}
	return ""
}

type ReverseInviteAck struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReverseInviteAck) Reset()         { *m = ReverseInviteAck{} }
func (m *ReverseInviteAck) String() string { return proto.CompactTextString(m) }
func (*ReverseInviteAck) ProtoMessage()    {}
func (*ReverseInviteAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{9}
}

func (m *ReverseInviteAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReverseInviteAck.Unmarshal(m, b)
}
func (m *ReverseInviteAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReverseInviteAck.Marshal(b, m, deterministic)
}
func (m *ReverseInviteAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReverseInviteAck.Merge(m, src)
}
func (m *ReverseInviteAck) XXX_Size() int {
	return xxx_messageInfo_ReverseInviteAck.Size(m)
}
func (m *ReverseInviteAck) XXX_DiscardUnknown() {
	xxx_messageInfo_ReverseInviteAck.DiscardUnknown(m)
}

var xxx_messageInfo_ReverseInviteAck proto.InternalMessageInfo

func (m *ReverseInviteAck) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type PriNetDig struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PriNetDig) Reset()         { *m = PriNetDig{} }
func (m *PriNetDig) String() string { return proto.CompactTextString(m) }
func (*PriNetDig) ProtoMessage()    {}
func (*PriNetDig) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{10}
}

func (m *PriNetDig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PriNetDig.Unmarshal(m, b)
}
func (m *PriNetDig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PriNetDig.Marshal(b, m, deterministic)
}
func (m *PriNetDig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriNetDig.Merge(m, src)
}
func (m *PriNetDig) XXX_Size() int {
	return xxx_messageInfo_PriNetDig.Size(m)
}
func (m *PriNetDig) XXX_DiscardUnknown() {
	xxx_messageInfo_PriNetDig.DiscardUnknown(m)
}

var xxx_messageInfo_PriNetDig proto.InternalMessageInfo

func (m *PriNetDig) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type NatMsg struct {
	T                    int32    `protobuf:"varint,1,opt,name=T,proto3" json:"T,omitempty"`
	L                    int32    `protobuf:"varint,2,opt,name=L,proto3" json:"L,omitempty"`
	V                    []byte   `protobuf:"bytes,3,opt,name=V,proto3" json:"V,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NatMsg) Reset()         { *m = NatMsg{} }
func (m *NatMsg) String() string { return proto.CompactTextString(m) }
func (*NatMsg) ProtoMessage()    {}
func (*NatMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{11}
}

func (m *NatMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NatMsg.Unmarshal(m, b)
}
func (m *NatMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NatMsg.Marshal(b, m, deterministic)
}
func (m *NatMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NatMsg.Merge(m, src)
}
func (m *NatMsg) XXX_Size() int {
	return xxx_messageInfo_NatMsg.Size(m)
}
func (m *NatMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_NatMsg.DiscardUnknown(m)
}

var xxx_messageInfo_NatMsg proto.InternalMessageInfo

func (m *NatMsg) GetT() int32 {
	if m != nil {
		return m.T
	}
	return 0
}

func (m *NatMsg) GetL() int32 {
	if m != nil {
		return m.L
	}
	return 0
}

func (m *NatMsg) GetV() []byte {
	if m != nil {
		return m.V
	}
	return nil
}

func init() {
	proto.RegisterEnum("net.pb.NatType", NatType_name, NatType_value)
	proto.RegisterType((*BootReg)(nil), "net.pb.BootReg")
	proto.RegisterType((*BootRegAck)(nil), "net.pb.BootRegAck")
	proto.RegisterType((*NatConnect)(nil), "net.pb.NatConnect")
	proto.RegisterType((*NatKeepAlive)(nil), "net.pb.NatKeepAlive")
	proto.RegisterType((*NatPing)(nil), "net.pb.NatPing")
	proto.RegisterType((*HoleDig)(nil), "net.pb.HoleDig")
	proto.RegisterType((*HolePayLoad)(nil), "net.pb.HolePayLoad")
	proto.RegisterType((*ErrorNotify)(nil), "net.pb.ErrorNotify")
	proto.RegisterType((*ReverseInvite)(nil), "net.pb.ReverseInvite")
	proto.RegisterType((*ReverseInviteAck)(nil), "net.pb.ReverseInviteAck")
	proto.RegisterType((*PriNetDig)(nil), "net.pb.PriNetDig")
	proto.RegisterType((*NatMsg)(nil), "net.pb.NatMsg")
}

func init() { proto.RegisterFile("natManager.proto", fileDescriptor_2f5cc5a30f7d4657) }

var fileDescriptor_2f5cc5a30f7d4657 = []byte{
	// 606 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5d, 0x6b, 0xdb, 0x30,
	0x14, 0x9d, 0x93, 0x26, 0x69, 0x6e, 0xfa, 0x61, 0x44, 0x19, 0xa1, 0x8c, 0x52, 0x04, 0x63, 0xeb,
	0x06, 0x61, 0x74, 0xbf, 0xa0, 0x69, 0x0b, 0x33, 0x4b, 0x8d, 0x71, 0xb3, 0xc2, 0x5e, 0x06, 0x72,
	0x7c, 0xeb, 0x8a, 0x64, 0x92, 0x90, 0x15, 0x97, 0xfe, 0x86, 0x3d, 0xec, 0x75, 0x3f, 0x77, 0x48,
	0x96, 0x53, 0x77, 0x6c, 0xeb, 0x9b, 0xce, 0xb9, 0xf2, 0x3d, 0x47, 0x47, 0x57, 0x86, 0x50, 0x30,
	0x73, 0xc5, 0x04, 0x2b, 0x50, 0x4f, 0x94, 0x96, 0x46, 0x92, 0xbe, 0x40, 0x33, 0x51, 0xd9, 0xe1,
	0xae, 0xc8, 0xca, 0xb3, 0x3c, 0xf7, 0x34, 0x65, 0x30, 0x98, 0x4a, 0x69, 0x52, 0x2c, 0xc8, 0x4b,
	0xe8, 0x0b, 0x99, 0x63, 0x94, 0x8f, 0x83, 0xe3, 0xe0, 0xed, 0x30, 0xf5, 0x88, 0x1c, 0xc3, 0x48,
	0x69, 0x5e, 0x31, 0x83, 0x89, 0xd4, 0x66, 0xdc, 0x75, 0xc5, 0x36, 0x45, 0x5e, 0xc1, 0xd0, 0xc3,
	0x48, 0x8d, 0x3b, 0xae, 0xfe, 0x48, 0xd0, 0x1f, 0x01, 0x80, 0xd7, 0x38, 0x5b, 0x2c, 0xc9, 0x09,
	0x0c, 0x04, 0x33, 0xf3, 0x07, 0x85, 0x4e, 0x67, 0xef, 0x74, 0x7f, 0x52, 0x5b, 0x9b, 0xc4, 0x35,
	0x9d, 0x36, 0xf5, 0x96, 0xa3, 0xce, 0x13, 0x47, 0x87, 0xb0, 0xad, 0xd6, 0xd9, 0x8a, 0x2f, 0x22,
	0xe5, 0xed, 0x6c, 0x30, 0x39, 0x02, 0xa8, 0xd7, 0xce, 0xec, 0x96, 0xab, 0xb6, 0x18, 0xfa, 0x2b,
	0x00, 0x88, 0x99, 0x39, 0x97, 0x42, 0xe0, 0xc2, 0x90, 0xf7, 0xb0, 0x7d, 0xab, 0xe5, 0x77, 0x9b,
	0x88, 0xb3, 0x33, 0x6a, 0xd9, 0xa9, 0x83, 0x4a, 0x37, 0x1b, 0xc8, 0x1b, 0xe8, 0x1b, 0xe9, 0xb6,
	0x76, 0xfe, 0xbe, 0xd5, 0x97, 0x6d, 0x20, 0x25, 0x96, 0x25, 0x97, 0x22, 0xca, 0xbd, 0xc3, 0x47,
	0xc2, 0x1e, 0xcb, 0xc8, 0x8d, 0xbd, 0x5e, 0xea, 0x11, 0x5d, 0xc1, 0x4e, 0xcc, 0xcc, 0x67, 0x44,
	0x75, 0xb6, 0xe2, 0x15, 0xfe, 0xf3, 0x42, 0x0e, 0xa0, 0xb7, 0xda, 0xb8, 0x18, 0xa6, 0x35, 0xb0,
	0x6c, 0xb2, 0xce, 0xa2, 0xc4, 0xeb, 0xd5, 0x80, 0x8c, 0x61, 0x90, 0xac, 0xb3, 0x96, 0x58, 0x03,
	0xe9, 0x57, 0x18, 0xc4, 0xcc, 0x24, 0x5c, 0x14, 0x84, 0xc0, 0x96, 0xe2, 0xa2, 0xf0, 0x32, 0x6e,
	0xed, 0x38, 0x29, 0x0a, 0xaf, 0xe1, 0xd6, 0x56, 0x42, 0x48, 0xb1, 0xc0, 0x46, 0xc2, 0x01, 0x12,
	0x42, 0x77, 0x3e, 0x9f, 0xf9, 0xf6, 0x76, 0x49, 0x23, 0x18, 0x7c, 0x92, 0x2b, 0xbc, 0xe0, 0xc5,
	0xd3, 0x24, 0x82, 0x3f, 0x93, 0x38, 0x86, 0x91, 0x40, 0x73, 0x2f, 0xf5, 0xd2, 0xcd, 0x43, 0xc7,
	0xb5, 0x68, 0x53, 0x14, 0x61, 0x64, 0x5b, 0x25, 0xec, 0x61, 0x26, 0x59, 0xfe, 0x4c, 0xbb, 0x23,
	0x00, 0xc3, 0x74, 0x81, 0xc6, 0x9d, 0xb7, 0xee, 0xd6, 0x62, 0x6c, 0x18, 0xaa, 0x6e, 0xe4, 0x4e,
	0xb0, 0x93, 0x36, 0x90, 0xbe, 0x86, 0xd1, 0xa5, 0xd6, 0x52, 0xc7, 0xd2, 0xf0, 0xdb, 0x07, 0x9b,
	0x3c, 0x6a, 0x7d, 0x55, 0x36, 0x91, 0x78, 0x44, 0x7f, 0x06, 0xb0, 0x9b, 0x62, 0x85, 0xba, 0xc4,
	0x48, 0x54, 0xdc, 0xe0, 0x33, 0x86, 0x0e, 0xa0, 0xa7, 0xd6, 0xd9, 0xe6, 0x51, 0xd4, 0xa0, 0x75,
	0xff, 0xdd, 0xf6, 0xfd, 0x5b, 0x5e, 0x21, 0xea, 0x28, 0xf7, 0x63, 0xeb, 0x91, 0x1d, 0x77, 0x3b,
	0x82, 0xee, 0x8b, 0x5e, 0x3d, 0xee, 0x0d, 0xa6, 0x1f, 0x20, 0x7c, 0x62, 0xc8, 0xbe, 0xb0, 0xff,
	0x7a, 0xa2, 0x27, 0x30, 0x4c, 0x34, 0x8f, 0xd1, 0x3c, 0x7b, 0x3d, 0xf4, 0x14, 0xfa, 0x31, 0x33,
	0x57, 0x65, 0x41, 0x76, 0x20, 0x98, 0xbb, 0x7a, 0x2f, 0x0d, 0xe6, 0x16, 0xcd, 0x7c, 0xbc, 0xc1,
	0xcc, 0xa2, 0x1b, 0x9f, 0x67, 0x70, 0xf3, 0xee, 0x9b, 0x1b, 0x2b, 0xf7, 0x7c, 0xf7, 0x00, 0xbe,
	0x88, 0xa5, 0x90, 0xf7, 0x22, 0xbd, 0xbc, 0x0e, 0x5f, 0x90, 0x7d, 0x18, 0xc5, 0x32, 0x66, 0xe6,
	0x02, 0x2b, 0xbe, 0xc0, 0x30, 0x20, 0xbb, 0x30, 0x9c, 0xe2, 0x1d, 0x17, 0x79, 0xcc, 0x4c, 0xd8,
	0x21, 0x04, 0xf6, 0xce, 0x99, 0x98, 0x62, 0xcc, 0xcc, 0x35, 0xea, 0x0a, 0x75, 0xd8, 0xb5, 0xdf,
	0xcc, 0xe5, 0x14, 0xcf, 0xef, 0x70, 0xb1, 0xc4, 0x3c, 0xdc, 0xca, 0xfa, 0xee, 0xbf, 0xf5, 0xf1,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x83, 0x3d, 0x09, 0xe2, 0x04, 0x00, 0x00,
}
