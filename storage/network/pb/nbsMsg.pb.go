// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nbsMsg.proto

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MsgType int32

const (
	MsgType_Error              MsgType = 0
	MsgType_NatBootReg         MsgType = 1001
	MsgType_NatKeepAlive       MsgType = 1002
	MsgType_NatDigApply        MsgType = 1003
	MsgType_NatPingPong        MsgType = 1004
	MsgType_NatDigOut          MsgType = 1006
	MsgType_NatReversInvite    MsgType = 1008
	MsgType_NatReversInviteAck MsgType = 1009
	MsgType_NatPriDigSyn       MsgType = 1010
	MsgType_NatPriDigAck       MsgType = 1011
	MsgType_NatDigConfirm      MsgType = 1012
	MsgType_NatBootAnswer      MsgType = 1013
	MsgType_GspSub             MsgType = 2001
	MsgType_GspSubACK          MsgType = 2002
	MsgType_GspVoteContact     MsgType = 2003
	MsgType_GspVoteResult      MsgType = 2004
	MsgType_GspVoteResAck      MsgType = 2005
	MsgType_GspIntroduce       MsgType = 2006
	MsgType_GspWelcome         MsgType = 2007
	MsgType_GspHeartBeat       MsgType = 2008
)

var MsgType_name = map[int32]string{
	0:    "Error",
	1001: "NatBootReg",
	1002: "NatKeepAlive",
	1003: "NatDigApply",
	1004: "NatPingPong",
	1006: "NatDigOut",
	1008: "NatReversInvite",
	1009: "NatReversInviteAck",
	1010: "NatPriDigSyn",
	1011: "NatPriDigAck",
	1012: "NatDigConfirm",
	1013: "NatBootAnswer",
	2001: "GspSub",
	2002: "GspSubACK",
	2003: "GspVoteContact",
	2004: "GspVoteResult",
	2005: "GspVoteResAck",
	2006: "GspIntroduce",
	2007: "GspWelcome",
	2008: "GspHeartBeat",
}

var MsgType_value = map[string]int32{
	"Error":              0,
	"NatBootReg":         1001,
	"NatKeepAlive":       1002,
	"NatDigApply":        1003,
	"NatPingPong":        1004,
	"NatDigOut":          1006,
	"NatReversInvite":    1008,
	"NatReversInviteAck": 1009,
	"NatPriDigSyn":       1010,
	"NatPriDigAck":       1011,
	"NatDigConfirm":      1012,
	"NatBootAnswer":      1013,
	"GspSub":             2001,
	"GspSubACK":          2002,
	"GspVoteContact":     2003,
	"GspVoteResult":      2004,
	"GspVoteResAck":      2005,
	"GspIntroduce":       2006,
	"GspWelcome":         2007,
	"GspHeartBeat":       2008,
}

func (x MsgType) String() string {
	return proto.EnumName(MsgType_name, int32(x))
}

func (MsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1ac27572d1ea9f39, []int{0}
}

func init() {
	proto.RegisterEnum("net.pb.MsgType", MsgType_name, MsgType_value)
}

func init() { proto.RegisterFile("nbsMsg.proto", fileDescriptor_1ac27572d1ea9f39) }

var fileDescriptor_1ac27572d1ea9f39 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x4b, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x11, 0x12, 0xb5, 0xea, 0x96, 0x7a, 0x30, 0x48, 0xdc, 0x81, 0x45, 0x37, 0x9c, 0x20,
	0x6d, 0x51, 0xa8, 0xaa, 0x86, 0xaa, 0x45, 0xb0, 0x4e, 0xc2, 0x10, 0x45, 0xa4, 0xb6, 0x65, 0x4f,
	0x8a, 0x7a, 0x4c, 0xde, 0x5c, 0x81, 0x87, 0x10, 0xcf, 0x0d, 0x2b, 0x64, 0x4a, 0x16, 0x74, 0xe9,
	0x4f, 0xbf, 0xfc, 0x7f, 0x33, 0xc3, 0x9b, 0x2a, 0x71, 0x43, 0x97, 0xb5, 0x8d, 0xd5, 0xa4, 0x65,
	0x4d, 0x21, 0xb5, 0x4d, 0xb2, 0xf3, 0xbd, 0xca, 0xd9, 0xd0, 0x65, 0x87, 0x73, 0x83, 0xb2, 0xce,
	0xd7, 0xf6, 0xac, 0xd5, 0x16, 0x56, 0xa4, 0xe0, 0x3c, 0x8a, 0xa9, 0xa3, 0x35, 0x8d, 0x31, 0x83,
	0x07, 0x26, 0x37, 0x78, 0x33, 0x8a, 0x69, 0x80, 0x68, 0x82, 0x22, 0x9f, 0x21, 0x3c, 0x32, 0x09,
	0xbc, 0x11, 0xc5, 0xd4, 0xcb, 0xb3, 0xc0, 0x98, 0x62, 0x0e, 0x4f, 0x15, 0x19, 0xe5, 0x2a, 0x1b,
	0x69, 0x95, 0xc1, 0x33, 0x93, 0x2d, 0x5e, 0x5f, 0x64, 0x0e, 0x4a, 0x82, 0x17, 0x26, 0xb7, 0xb8,
	0x88, 0x62, 0x1a, 0xe3, 0x0c, 0xad, 0xeb, 0xab, 0x59, 0x4e, 0x08, 0xaf, 0x4c, 0x6e, 0x73, 0xb9,
	0x44, 0x83, 0xf4, 0x0c, 0xde, 0xaa, 0xd6, 0x91, 0xcd, 0x7b, 0x79, 0x36, 0x99, 0x2b, 0x78, 0xff,
	0x8f, 0x7c, 0xea, 0x83, 0x49, 0xc9, 0xd7, 0x17, 0x25, 0x5d, 0xad, 0x4e, 0x73, 0x3b, 0x85, 0xcf,
	0x8a, 0xf9, 0x01, 0x02, 0xe5, 0xce, 0xd1, 0xc2, 0x17, 0x93, 0x0d, 0x5e, 0x0b, 0x9d, 0x99, 0x94,
	0x09, 0x5c, 0x08, 0x6f, 0xb6, 0x78, 0x04, 0xdd, 0x01, 0x5c, 0x0a, 0xb9, 0xc9, 0x5b, 0xa1, 0x33,
	0x47, 0x9a, 0xb0, 0xab, 0x15, 0xc5, 0x29, 0xc1, 0x95, 0xf0, 0xbf, 0xfc, 0xc1, 0x31, 0xba, 0xb2,
	0x20, 0xb8, 0x5e, 0x62, 0xde, 0xe0, 0x46, 0x78, 0xa9, 0xd0, 0x99, 0xbe, 0x22, 0xab, 0x4f, 0xca,
	0x14, 0xe1, 0x56, 0xf8, 0x0d, 0x86, 0xce, 0x1c, 0x63, 0x91, 0xea, 0x29, 0xc2, 0x5d, 0x95, 0xd9,
	0xc7, 0xd8, 0x52, 0x07, 0x63, 0x82, 0x7b, 0x91, 0xd4, 0x7e, 0x6f, 0xb1, 0xfb, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x70, 0x83, 0x84, 0x51, 0x9b, 0x01, 0x00, 0x00,
}
