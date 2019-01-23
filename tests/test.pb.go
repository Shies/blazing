// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: test.proto

/*
Package tests is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	Time
*/
package tests

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Time struct {
	Now int64 `protobuf:"varint,1,opt,name=Now,proto3" json:"Now,omitempty"`
}

func (m *Time) Reset()                    { *m = Time{} }
func (m *Time) String() string            { return proto.CompactTextString(m) }
func (*Time) ProtoMessage()               {}
func (*Time) Descriptor() ([]byte, []int) { return fileDescriptorTest, []int{0} }

func (m *Time) GetNow() int64 {
	if m != nil {
		return m.Now
	}
	return 0
}

func init() {
	proto.RegisterType((*Time)(nil), "tests.Time")
}

func init() { proto.RegisterFile("test.proto", fileDescriptorTest) }

var fileDescriptorTest = []byte{
	// 70 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0xb1, 0x8b, 0x95, 0x24, 0xb8, 0x58, 0x42,
	0x32, 0x73, 0x53, 0x85, 0x04, 0xb8, 0x98, 0xfd, 0xf2, 0xcb, 0x25, 0x18, 0x15, 0x18, 0x35, 0x98,
	0x83, 0x40, 0xcc, 0x24, 0x36, 0xb0, 0x3a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x42, 0x5c,
	0x67, 0x7c, 0x35, 0x00, 0x00, 0x00,
}
