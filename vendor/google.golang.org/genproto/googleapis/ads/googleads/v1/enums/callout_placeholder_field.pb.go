// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/callout_placeholder_field.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Possible values for Callout placeholder fields.
type CalloutPlaceholderFieldEnum_CalloutPlaceholderField int32

const (
	// Not specified.
	CalloutPlaceholderFieldEnum_UNSPECIFIED CalloutPlaceholderFieldEnum_CalloutPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	CalloutPlaceholderFieldEnum_UNKNOWN CalloutPlaceholderFieldEnum_CalloutPlaceholderField = 1
	// Data Type: STRING. Callout text.
	CalloutPlaceholderFieldEnum_CALLOUT_TEXT CalloutPlaceholderFieldEnum_CalloutPlaceholderField = 2
)

var CalloutPlaceholderFieldEnum_CalloutPlaceholderField_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CALLOUT_TEXT",
}
var CalloutPlaceholderFieldEnum_CalloutPlaceholderField_value = map[string]int32{
	"UNSPECIFIED":  0,
	"UNKNOWN":      1,
	"CALLOUT_TEXT": 2,
}

func (x CalloutPlaceholderFieldEnum_CalloutPlaceholderField) String() string {
	return proto.EnumName(CalloutPlaceholderFieldEnum_CalloutPlaceholderField_name, int32(x))
}
func (CalloutPlaceholderFieldEnum_CalloutPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_callout_placeholder_field_16aa589fff00a281, []int{0, 0}
}

// Values for Callout placeholder fields.
type CalloutPlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalloutPlaceholderFieldEnum) Reset()         { *m = CalloutPlaceholderFieldEnum{} }
func (m *CalloutPlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*CalloutPlaceholderFieldEnum) ProtoMessage()    {}
func (*CalloutPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_callout_placeholder_field_16aa589fff00a281, []int{0}
}
func (m *CalloutPlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalloutPlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *CalloutPlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalloutPlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (dst *CalloutPlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalloutPlaceholderFieldEnum.Merge(dst, src)
}
func (m *CalloutPlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_CalloutPlaceholderFieldEnum.Size(m)
}
func (m *CalloutPlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CalloutPlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CalloutPlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CalloutPlaceholderFieldEnum)(nil), "google.ads.googleads.v1.enums.CalloutPlaceholderFieldEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.CalloutPlaceholderFieldEnum_CalloutPlaceholderField", CalloutPlaceholderFieldEnum_CalloutPlaceholderField_name, CalloutPlaceholderFieldEnum_CalloutPlaceholderField_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/callout_placeholder_field.proto", fileDescriptor_callout_placeholder_field_16aa589fff00a281)
}

var fileDescriptor_callout_placeholder_field_16aa589fff00a281 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x4a, 0xfb, 0x30,
	0x18, 0xc5, 0xff, 0xeb, 0x1f, 0x14, 0x32, 0xc1, 0xd2, 0x1b, 0x41, 0xb7, 0x8b, 0xed, 0x01, 0x12,
	0x8a, 0x77, 0x11, 0x2f, 0xb2, 0xd9, 0x8d, 0xe1, 0xe8, 0x0a, 0x6e, 0x53, 0xa4, 0x30, 0xe2, 0x12,
	0xb3, 0x42, 0x96, 0x94, 0xa5, 0xdd, 0x03, 0x79, 0xe9, 0xa3, 0xf8, 0x28, 0xde, 0xf9, 0x06, 0xd2,
	0xc4, 0xd6, 0xab, 0x7a, 0x53, 0x0e, 0x3d, 0xdf, 0xf9, 0xe5, 0x7c, 0x1f, 0xb8, 0x15, 0x5a, 0x0b,
	0xc9, 0x11, 0x65, 0x06, 0x39, 0x59, 0xa9, 0x63, 0x88, 0xb8, 0x2a, 0xf7, 0x06, 0x6d, 0xa9, 0x94,
	0xba, 0x2c, 0x36, 0xb9, 0xa4, 0x5b, 0xbe, 0xd3, 0x92, 0xf1, 0xc3, 0xe6, 0x35, 0xe3, 0x92, 0xc1,
	0xfc, 0xa0, 0x0b, 0x1d, 0xf4, 0x5d, 0x06, 0x52, 0x66, 0x60, 0x13, 0x87, 0xc7, 0x10, 0xda, 0xf8,
	0x65, 0xaf, 0xa6, 0xe7, 0x19, 0xa2, 0x4a, 0xe9, 0x82, 0x16, 0x99, 0x56, 0xc6, 0x85, 0x87, 0x3b,
	0x70, 0x35, 0x76, 0xfc, 0xe4, 0x17, 0x3f, 0xa9, 0xe8, 0x91, 0x2a, 0xf7, 0xc3, 0x19, 0xb8, 0x68,
	0xb1, 0x83, 0x73, 0xd0, 0x5d, 0xc5, 0x0f, 0x49, 0x34, 0x9e, 0x4d, 0x66, 0xd1, 0x9d, 0xff, 0x2f,
	0xe8, 0x82, 0xd3, 0x55, 0x7c, 0x1f, 0x2f, 0x1e, 0x63, 0xbf, 0x13, 0xf8, 0xe0, 0x6c, 0x4c, 0xe6,
	0xf3, 0xc5, 0x6a, 0xb9, 0x59, 0x46, 0x4f, 0x4b, 0xdf, 0x1b, 0x7d, 0x75, 0xc0, 0x60, 0xab, 0xf7,
	0xf0, 0xcf, 0xb6, 0xa3, 0x5e, 0xcb, 0x73, 0x49, 0xd5, 0x36, 0xe9, 0x3c, 0x8f, 0x7e, 0xe2, 0x42,
	0x4b, 0xaa, 0x04, 0xd4, 0x07, 0x81, 0x04, 0x57, 0x76, 0x97, 0xfa, 0x76, 0x79, 0x66, 0x5a, 0x4e,
	0x79, 0x63, 0xbf, 0x6f, 0xde, 0xff, 0x29, 0x21, 0xef, 0x5e, 0x7f, 0xea, 0x50, 0x84, 0x19, 0xe8,
	0x64, 0xa5, 0xd6, 0x21, 0xac, 0x36, 0x37, 0x1f, 0xb5, 0x9f, 0x12, 0x66, 0xd2, 0xc6, 0x4f, 0xd7,
	0x61, 0x6a, 0xfd, 0x4f, 0x6f, 0xe0, 0x7e, 0x62, 0x4c, 0x98, 0xc1, 0xb8, 0x99, 0xc0, 0x78, 0x1d,
	0x62, 0x6c, 0x67, 0x5e, 0x4e, 0x6c, 0xb1, 0xeb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x0a,
	0xbf, 0xd3, 0xe2, 0x01, 0x00, 0x00,
}
