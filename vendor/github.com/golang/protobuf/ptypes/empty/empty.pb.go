// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/empty.proto

package empty

import proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A generic empty message that you can re-use to avoid defining duplicated
// empty messages in your APIs. A typical example is to use it as the request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_900544acb223d5b8, []int{0}
}
func (*Empty) XXX_WellKnownType() string { return "Empty" }
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Empty)(nil), "google.protobuf.Empty")
}

func init() { proto.RegisterFile("google/protobuf/empty.proto", fileDescriptor_900544acb223d5b8) }

var fileDescriptor_900544acb223d5b8 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x4f, 0xcd, 0x2d, 0x28,
	0xa9, 0xd4, 0x03, 0x73, 0x85, 0xf8, 0x21, 0x92, 0x7a, 0x30, 0x49, 0x25, 0x76, 0x2e, 0x56, 0x57,
	0x90, 0xbc, 0x53, 0x19, 0x97, 0x70, 0x72, 0x7e, 0xae, 0x1e, 0x9a, 0xbc, 0x13, 0x17, 0x58, 0x36,
	0x00, 0xc4, 0x0d, 0x60, 0x8c, 0x52, 0x4f, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf,
	0xd5, 0x4f, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0x47, 0x58, 0x53, 0x50, 0x52, 0x59, 0x90, 0x5a, 0x0c,
	0xb1, 0xed, 0x07, 0x23, 0xe3, 0x22, 0x26, 0x66, 0xf7, 0x00, 0xa7, 0x55, 0x4c, 0x72, 0xee, 0x10,
	0x13, 0x03, 0xa0, 0xea, 0xf4, 0xc2, 0x53, 0x73, 0x72, 0xbc, 0xf3, 0xf2, 0xcb, 0xf3, 0x42, 0x40,
	0xea, 0x93, 0xd8, 0xc0, 0x06, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x64, 0xd4, 0xb3, 0xa6,
	0xb7, 0x00, 0x00, 0x00,
}
