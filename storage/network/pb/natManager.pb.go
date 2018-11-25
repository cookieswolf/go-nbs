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

type NatMsgType int32

const (
	NatMsgType_UnknownReq    NatMsgType = 0
	NatMsgType_BootStrapReg  NatMsgType = 1
	NatMsgType_KeepAlive     NatMsgType = 2
	NatMsgType_Connect       NatMsgType = 3
	NatMsgType_Ping          NatMsgType = 4
	NatMsgType_DigIn         NatMsgType = 5
	NatMsgType_DigOut        NatMsgType = 6
	NatMsgType_ReverseDig    NatMsgType = 7
	NatMsgType_ReverseDigACK NatMsgType = 8
	NatMsgType_error         NatMsgType = 20
)

var NatMsgType_name = map[int32]string{
	0:  "UnknownReq",
	1:  "BootStrapReg",
	2:  "KeepAlive",
	3:  "Connect",
	4:  "Ping",
	5:  "DigIn",
	6:  "DigOut",
	7:  "ReverseDig",
	8:  "ReverseDigACK",
	20: "error",
}

var NatMsgType_value = map[string]int32{
	"UnknownReq":    0,
	"BootStrapReg":  1,
	"KeepAlive":     2,
	"Connect":       3,
	"Ping":          4,
	"DigIn":         5,
	"DigOut":        6,
	"ReverseDig":    7,
	"ReverseDigACK": 8,
	"error":         20,
}

func (x NatMsgType) String() string {
	return proto.EnumName(NatMsgType_name, int32(x))
}

func (NatMsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{0}
}

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
	return fileDescriptor_2f5cc5a30f7d4657, []int{1}
}

type BootNatRegReq struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	PrivatePort          string   `protobuf:"bytes,3,opt,name=privatePort,proto3" json:"privatePort,omitempty"`
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

func (m *BootNatRegReq) GetPrivateIp() string {
	if m != nil {
		return m.PrivateIp
	}
	return ""
}

type BootNatRegRes struct {
	NatType              NatType  `protobuf:"varint,1,opt,name=natType,proto3,enum=net.pb.NatType" json:"natType,omitempty"`
	NodeId               string   `protobuf:"bytes,2,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	PublicIp             string   `protobuf:"bytes,3,opt,name=publicIp,proto3" json:"publicIp,omitempty"`
	PublicPort           string   `protobuf:"bytes,4,opt,name=publicPort,proto3" json:"publicPort,omitempty"`
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
	return NatType_UnknownRES
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
	return fileDescriptor_2f5cc5a30f7d4657, []int{6}
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
	return fileDescriptor_2f5cc5a30f7d4657, []int{7}
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
	return fileDescriptor_2f5cc5a30f7d4657, []int{8}
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

type NatRequest struct {
	MsgType              NatMsgType        `protobuf:"varint,1,opt,name=msgType,proto3,enum=net.pb.NatMsgType" json:"msgType,omitempty"`
	BootRegReq           *BootNatRegReq    `protobuf:"bytes,2,opt,name=BootRegReq,proto3" json:"BootRegReq,omitempty"`
	ConnReq              *NatConnect       `protobuf:"bytes,3,opt,name=connReq,proto3" json:"connReq,omitempty"`
	KeepAlive            *NatKeepAlive     `protobuf:"bytes,4,opt,name=keepAlive,proto3" json:"keepAlive,omitempty"`
	Ping                 *NatPing          `protobuf:"bytes,5,opt,name=ping,proto3" json:"ping,omitempty"`
	HoleMsg              *HoleDig          `protobuf:"bytes,6,opt,name=holeMsg,proto3" json:"holeMsg,omitempty"`
	Invite               *ReverseInvite    `protobuf:"bytes,7,opt,name=invite,proto3" json:"invite,omitempty"`
	InviteAck            *ReverseInviteAck `protobuf:"bytes,8,opt,name=inviteAck,proto3" json:"inviteAck,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *NatRequest) Reset()         { *m = NatRequest{} }
func (m *NatRequest) String() string { return proto.CompactTextString(m) }
func (*NatRequest) ProtoMessage()    {}
func (*NatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{9}
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
	return NatMsgType_UnknownReq
}

func (m *NatRequest) GetBootRegReq() *BootNatRegReq {
	if m != nil {
		return m.BootRegReq
	}
	return nil
}

func (m *NatRequest) GetConnReq() *NatConnect {
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

func (m *NatRequest) GetPing() *NatPing {
	if m != nil {
		return m.Ping
	}
	return nil
}

func (m *NatRequest) GetHoleMsg() *HoleDig {
	if m != nil {
		return m.HoleMsg
	}
	return nil
}

func (m *NatRequest) GetInvite() *ReverseInvite {
	if m != nil {
		return m.Invite
	}
	return nil
}

func (m *NatRequest) GetInviteAck() *ReverseInviteAck {
	if m != nil {
		return m.InviteAck
	}
	return nil
}

type NatResponse struct {
	MsgType              NatMsgType     `protobuf:"varint,1,opt,name=msgType,proto3,enum=net.pb.NatMsgType" json:"msgType,omitempty"`
	BootRegRes           *BootNatRegRes `protobuf:"bytes,2,opt,name=BootRegRes,proto3" json:"BootRegRes,omitempty"`
	ConnRes              *NatConnect    `protobuf:"bytes,3,opt,name=connRes,proto3" json:"connRes,omitempty"`
	KeepAlive            *NatKeepAlive  `protobuf:"bytes,4,opt,name=keepAlive,proto3" json:"keepAlive,omitempty"`
	Pong                 *NatPing       `protobuf:"bytes,5,opt,name=pong,proto3" json:"pong,omitempty"`
	HoleMsg              *HoleDig       `protobuf:"bytes,6,opt,name=holeMsg,proto3" json:"holeMsg,omitempty"`
	Invite               *ReverseInvite `protobuf:"bytes,7,opt,name=invite,proto3" json:"invite,omitempty"`
	Error                *ErrorNotify   `protobuf:"bytes,50,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *NatResponse) Reset()         { *m = NatResponse{} }
func (m *NatResponse) String() string { return proto.CompactTextString(m) }
func (*NatResponse) ProtoMessage()    {}
func (*NatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f5cc5a30f7d4657, []int{10}
}

func (m *NatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NatResponse.Unmarshal(m, b)
}
func (m *NatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NatResponse.Marshal(b, m, deterministic)
}
func (m *NatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NatResponse.Merge(m, src)
}
func (m *NatResponse) XXX_Size() int {
	return xxx_messageInfo_NatResponse.Size(m)
}
func (m *NatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NatResponse proto.InternalMessageInfo

func (m *NatResponse) GetMsgType() NatMsgType {
	if m != nil {
		return m.MsgType
	}
	return NatMsgType_UnknownReq
}

func (m *NatResponse) GetBootRegRes() *BootNatRegRes {
	if m != nil {
		return m.BootRegRes
	}
	return nil
}

func (m *NatResponse) GetConnRes() *NatConnect {
	if m != nil {
		return m.ConnRes
	}
	return nil
}

func (m *NatResponse) GetKeepAlive() *NatKeepAlive {
	if m != nil {
		return m.KeepAlive
	}
	return nil
}

func (m *NatResponse) GetPong() *NatPing {
	if m != nil {
		return m.Pong
	}
	return nil
}

func (m *NatResponse) GetHoleMsg() *HoleDig {
	if m != nil {
		return m.HoleMsg
	}
	return nil
}

func (m *NatResponse) GetInvite() *ReverseInvite {
	if m != nil {
		return m.Invite
	}
	return nil
}

func (m *NatResponse) GetError() *ErrorNotify {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterEnum("net.pb.NatMsgType", NatMsgType_name, NatMsgType_value)
	proto.RegisterEnum("net.pb.NatType", NatType_name, NatType_value)
	proto.RegisterType((*BootNatRegReq)(nil), "net.pb.BootNatRegReq")
	proto.RegisterType((*BootNatRegRes)(nil), "net.pb.BootNatRegRes")
	proto.RegisterType((*NatConnect)(nil), "net.pb.NatConnect")
	proto.RegisterType((*NatKeepAlive)(nil), "net.pb.NatKeepAlive")
	proto.RegisterType((*NatPing)(nil), "net.pb.NatPing")
	proto.RegisterType((*HoleDig)(nil), "net.pb.HoleDig")
	proto.RegisterType((*ErrorNotify)(nil), "net.pb.ErrorNotify")
	proto.RegisterType((*ReverseInvite)(nil), "net.pb.ReverseInvite")
	proto.RegisterType((*ReverseInviteAck)(nil), "net.pb.ReverseInviteAck")
	proto.RegisterType((*NatRequest)(nil), "net.pb.natRequest")
	proto.RegisterType((*NatResponse)(nil), "net.pb.natResponse")
}

func init() { proto.RegisterFile("natManager.proto", fileDescriptor_2f5cc5a30f7d4657) }

var fileDescriptor_2f5cc5a30f7d4657 = []byte{
	// 800 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdf, 0x8f, 0xdb, 0x44,
	0x10, 0x6e, 0x7e, 0xd9, 0xc9, 0xa4, 0x77, 0x5d, 0x96, 0x03, 0x59, 0x15, 0x42, 0x95, 0x11, 0x2a,
	0x57, 0xe0, 0x84, 0x82, 0xe0, 0x3d, 0x3f, 0x2a, 0x61, 0x95, 0x33, 0xd1, 0x5e, 0x78, 0xe0, 0x05,
	0xc9, 0xb1, 0xa7, 0xbe, 0x55, 0xd2, 0x5d, 0xd7, 0xbb, 0x09, 0xea, 0xbf, 0xc0, 0x03, 0xbc, 0xc2,
	0x3f, 0x8b, 0xd0, 0xae, 0xd7, 0x8e, 0x7d, 0x2a, 0x57, 0x09, 0xc4, 0xdb, 0xce, 0xcc, 0x37, 0x9e,
	0x6f, 0xbf, 0x99, 0x59, 0x03, 0x11, 0x89, 0xbe, 0x4e, 0x44, 0x92, 0x63, 0x79, 0x55, 0x94, 0x52,
	0x4b, 0xea, 0x09, 0xd4, 0x57, 0xc5, 0xf6, 0xf1, 0x99, 0xd8, 0xaa, 0x79, 0x96, 0x39, 0x77, 0x98,
	0xc3, 0xd9, 0x42, 0x4a, 0x1d, 0x27, 0x9a, 0x61, 0xce, 0xf0, 0x35, 0xfd, 0x10, 0x3c, 0x21, 0x33,
	0x8c, 0xb2, 0xa0, 0xf7, 0xa4, 0xf7, 0xd9, 0x84, 0x39, 0x8b, 0x3e, 0x81, 0x69, 0x51, 0xf2, 0x63,
	0xa2, 0x71, 0x2d, 0x4b, 0x1d, 0x0c, 0x6c, 0xb0, 0xed, 0xa2, 0x1f, 0xc1, 0xc4, 0x99, 0x51, 0x11,
	0xf4, 0x6d, 0xfc, 0xe4, 0x08, 0x7f, 0xeb, 0x75, 0x2b, 0x29, 0x7a, 0x09, 0xbe, 0x48, 0xf4, 0xe6,
	0x4d, 0x81, 0xb6, 0xd4, 0xf9, 0xec, 0xd1, 0x55, 0xc5, 0xf1, 0x2a, 0xae, 0xdc, 0xac, 0x8e, 0xb7,
	0x48, 0xf5, 0x3b, 0xa4, 0x1e, 0xc3, 0xb8, 0x38, 0x6c, 0xf7, 0x3c, 0x8d, 0x0a, 0xc7, 0xa8, 0xb1,
	0xe9, 0xc7, 0x00, 0xd5, 0xd9, 0xf2, 0x1d, 0xda, 0x68, 0xcb, 0x13, 0xfe, 0xd1, 0x03, 0x88, 0x13,
	0xbd, 0x94, 0x42, 0x60, 0xaa, 0xe9, 0xe7, 0x30, 0x7e, 0x59, 0xca, 0x57, 0x46, 0x1a, 0x4b, 0x67,
	0xda, 0xa2, 0x53, 0x29, 0xc6, 0x1a, 0x00, 0x7d, 0x0a, 0x9e, 0x96, 0x16, 0xda, 0x7f, 0x3b, 0xd4,
	0x85, 0x8d, 0x26, 0x0a, 0x95, 0xe2, 0x52, 0x44, 0x99, 0x63, 0x78, 0x72, 0x98, 0x6b, 0x69, 0xd9,
	0xd0, 0x1b, 0x31, 0x67, 0x85, 0x7b, 0x78, 0x18, 0x27, 0xfa, 0x05, 0x62, 0x31, 0xdf, 0xf3, 0x23,
	0xfe, 0x63, 0x4f, 0x2e, 0x60, 0xb4, 0x6f, 0x58, 0x4c, 0x58, 0x65, 0x18, 0xef, 0xfa, 0xb0, 0x8d,
	0xd6, 0xae, 0x5e, 0x65, 0xd0, 0x00, 0xfc, 0xf5, 0x61, 0xdb, 0x2a, 0x56, 0x9b, 0xe1, 0x4f, 0xe0,
	0xc7, 0x89, 0x5e, 0x73, 0x91, 0x53, 0x0a, 0xc3, 0x82, 0x8b, 0xdc, 0x95, 0xb1, 0x67, 0xeb, 0x93,
	0x22, 0x77, 0x35, 0xec, 0xd9, 0x94, 0x10, 0x52, 0xa4, 0x58, 0x97, 0xb0, 0x06, 0x25, 0x30, 0xd8,
	0x6c, 0xbe, 0x77, 0x9f, 0x37, 0xc7, 0xf0, 0x29, 0xf8, 0xdf, 0xc9, 0x3d, 0xae, 0x78, 0xde, 0x55,
	0xa2, 0x77, 0x47, 0x89, 0xf0, 0x53, 0x98, 0x3e, 0x2f, 0x4b, 0x59, 0xc6, 0x52, 0xf3, 0x97, 0x6f,
	0xcc, 0x85, 0xb1, 0x2c, 0xaf, 0x55, 0xcd, 0xc4, 0x59, 0xe1, 0xef, 0x3d, 0x38, 0x63, 0x78, 0xc4,
	0x52, 0x61, 0x24, 0x8e, 0x5c, 0xe3, 0xfd, 0x9f, 0x35, 0x3c, 0x8b, 0xc3, 0xb6, 0x19, 0xc7, 0xca,
	0x68, 0xc9, 0x3e, 0x68, 0xcb, 0x6e, 0xfc, 0x05, 0x62, 0x19, 0x65, 0x6e, 0x5a, 0x9c, 0x65, 0xa6,
	0xcc, 0x74, 0xde, 0x66, 0x8c, 0xaa, 0x29, 0xab, 0xed, 0xf0, 0x2b, 0x20, 0x1d, 0x42, 0xf3, 0x74,
	0xf7, 0x8e, 0xab, 0xfe, 0x3a, 0x00, 0x10, 0x66, 0x09, 0x5e, 0x1f, 0x50, 0x69, 0xfa, 0x05, 0xf8,
	0xaf, 0x54, 0xde, 0xda, 0x02, 0xda, 0xda, 0x82, 0xeb, 0x2a, 0xc2, 0x6a, 0x08, 0xfd, 0x06, 0xc0,
	0x2c, 0x51, 0xb5, 0xab, 0x6e, 0xf8, 0x3e, 0xa8, 0x13, 0x3a, 0x8b, 0xcc, 0x5a, 0x40, 0x53, 0x24,
	0x95, 0x42, 0x98, 0x9c, 0x81, 0xcd, 0x69, 0x17, 0x71, 0x1b, 0xc0, 0x6a, 0x08, 0x9d, 0xc1, 0x64,
	0x57, 0xcf, 0x9e, 0x95, 0x62, 0x3a, 0xbb, 0x68, 0xe1, 0x9b, 0xb9, 0x64, 0x27, 0x18, 0xfd, 0xc4,
	0x4d, 0xce, 0xe8, 0xce, 0x3e, 0x54, 0x83, 0xe5, 0x46, 0xe9, 0x12, 0xfc, 0x5b, 0xb9, 0x47, 0xd3,
	0x57, 0xaf, 0x8b, 0x73, 0x53, 0xc2, 0xea, 0x38, 0xfd, 0x12, 0x3c, 0x6e, 0x05, 0x0d, 0xfc, 0xee,
	0x25, 0x3b, 0x6a, 0x33, 0x07, 0xa2, 0xdf, 0xc2, 0x84, 0xd7, 0xfa, 0x07, 0x63, 0x9b, 0x11, 0xbc,
	0x35, 0x63, 0x9e, 0xee, 0xd8, 0x09, 0x1a, 0xfe, 0xd5, 0x87, 0xa9, 0x6d, 0x86, 0x2a, 0xa4, 0x50,
	0xf8, 0x1f, 0xba, 0xa1, 0xee, 0xeb, 0x86, 0x6a, 0x75, 0x43, 0x9d, 0xba, 0xa1, 0xde, 0xdd, 0x0d,
	0xf5, 0xaf, 0xbb, 0x21, 0xef, 0xeb, 0x86, 0xfc, 0x5f, 0xbb, 0x71, 0x09, 0x23, 0x34, 0xdb, 0x1c,
	0xcc, 0x2c, 0xfa, 0xfd, 0x1a, 0xdd, 0x5a, 0x71, 0x56, 0x21, 0x9e, 0xfd, 0x59, 0xbd, 0xc2, 0x4e,
	0x5a, 0x7a, 0x0e, 0xf0, 0xa3, 0xd8, 0x09, 0xf9, 0x8b, 0x19, 0x44, 0xf2, 0x80, 0x12, 0x78, 0x68,
	0x84, 0xbb, 0xd1, 0x65, 0x52, 0x30, 0xcc, 0x49, 0x8f, 0x9e, 0xc1, 0xa4, 0xb9, 0x32, 0xe9, 0xd3,
	0x29, 0xf8, 0x4e, 0x31, 0x32, 0xa0, 0x63, 0x18, 0x9a, 0xfb, 0x91, 0x21, 0x9d, 0xc0, 0x68, 0xc5,
	0xf3, 0x48, 0x90, 0x11, 0x05, 0xf0, 0x56, 0x3c, 0xff, 0xe1, 0xa0, 0x89, 0x67, 0x3e, 0xef, 0x18,
	0xaf, 0x78, 0x4e, 0x7c, 0xfa, 0x5e, 0xf3, 0x9c, 0xac, 0x78, 0x3e, 0x5f, 0xbe, 0x20, 0x63, 0x93,
	0x69, 0x99, 0x91, 0x8b, 0x67, 0x3f, 0xdb, 0x87, 0xf1, 0x2e, 0xaf, 0xe7, 0x37, 0xe4, 0x01, 0x7d,
	0x04, 0xd3, 0x58, 0xc6, 0x89, 0x5e, 0xe1, 0x91, 0xa7, 0x58, 0xd1, 0x5a, 0xe0, 0x2d, 0x17, 0x59,
	0x9c, 0x68, 0xd2, 0xa7, 0x14, 0xce, 0x97, 0x89, 0x58, 0x60, 0x9c, 0xe8, 0x1b, 0x2c, 0x8f, 0x58,
	0x92, 0x81, 0xc9, 0xd9, 0xc8, 0x05, 0x2e, 0x6f, 0x31, 0xdd, 0x61, 0x46, 0x86, 0x5b, 0xcf, 0xfe,
	0x82, 0xbf, 0xfe, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x92, 0xa3, 0x01, 0xad, 0x07, 0x00, 0x00,
}
