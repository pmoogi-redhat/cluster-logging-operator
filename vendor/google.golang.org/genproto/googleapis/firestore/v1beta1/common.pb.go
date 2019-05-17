// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/firestore/v1beta1/common.proto

package firestore // import "google.golang.org/genproto/googleapis/firestore/v1beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// A set of field paths on a document.
// Used to restrict a get or update operation on a document to a subset of its
// fields.
// This is different from standard field masks, as this is always scoped to a
// [Document][google.firestore.v1beta1.Document], and takes in account the dynamic nature of [Value][google.firestore.v1beta1.Value].
type DocumentMask struct {
	// The list of field paths in the mask. See [Document.fields][google.firestore.v1beta1.Document.fields] for a field
	// path syntax reference.
	FieldPaths           []string `protobuf:"bytes,1,rep,name=field_paths,json=fieldPaths,proto3" json:"field_paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DocumentMask) Reset()         { *m = DocumentMask{} }
func (m *DocumentMask) String() string { return proto.CompactTextString(m) }
func (*DocumentMask) ProtoMessage()    {}
func (*DocumentMask) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_33e7705eedc91e86, []int{0}
}
func (m *DocumentMask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocumentMask.Unmarshal(m, b)
}
func (m *DocumentMask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocumentMask.Marshal(b, m, deterministic)
}
func (dst *DocumentMask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocumentMask.Merge(dst, src)
}
func (m *DocumentMask) XXX_Size() int {
	return xxx_messageInfo_DocumentMask.Size(m)
}
func (m *DocumentMask) XXX_DiscardUnknown() {
	xxx_messageInfo_DocumentMask.DiscardUnknown(m)
}

var xxx_messageInfo_DocumentMask proto.InternalMessageInfo

func (m *DocumentMask) GetFieldPaths() []string {
	if m != nil {
		return m.FieldPaths
	}
	return nil
}

// A precondition on a document, used for conditional operations.
type Precondition struct {
	// The type of precondition.
	//
	// Types that are valid to be assigned to ConditionType:
	//	*Precondition_Exists
	//	*Precondition_UpdateTime
	ConditionType        isPrecondition_ConditionType `protobuf_oneof:"condition_type"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *Precondition) Reset()         { *m = Precondition{} }
func (m *Precondition) String() string { return proto.CompactTextString(m) }
func (*Precondition) ProtoMessage()    {}
func (*Precondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_33e7705eedc91e86, []int{1}
}
func (m *Precondition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Precondition.Unmarshal(m, b)
}
func (m *Precondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Precondition.Marshal(b, m, deterministic)
}
func (dst *Precondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Precondition.Merge(dst, src)
}
func (m *Precondition) XXX_Size() int {
	return xxx_messageInfo_Precondition.Size(m)
}
func (m *Precondition) XXX_DiscardUnknown() {
	xxx_messageInfo_Precondition.DiscardUnknown(m)
}

var xxx_messageInfo_Precondition proto.InternalMessageInfo

type isPrecondition_ConditionType interface {
	isPrecondition_ConditionType()
}

type Precondition_Exists struct {
	Exists bool `protobuf:"varint,1,opt,name=exists,proto3,oneof"`
}

type Precondition_UpdateTime struct {
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=update_time,json=updateTime,proto3,oneof"`
}

func (*Precondition_Exists) isPrecondition_ConditionType() {}

func (*Precondition_UpdateTime) isPrecondition_ConditionType() {}

func (m *Precondition) GetConditionType() isPrecondition_ConditionType {
	if m != nil {
		return m.ConditionType
	}
	return nil
}

func (m *Precondition) GetExists() bool {
	if x, ok := m.GetConditionType().(*Precondition_Exists); ok {
		return x.Exists
	}
	return false
}

func (m *Precondition) GetUpdateTime() *timestamp.Timestamp {
	if x, ok := m.GetConditionType().(*Precondition_UpdateTime); ok {
		return x.UpdateTime
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Precondition) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Precondition_OneofMarshaler, _Precondition_OneofUnmarshaler, _Precondition_OneofSizer, []interface{}{
		(*Precondition_Exists)(nil),
		(*Precondition_UpdateTime)(nil),
	}
}

func _Precondition_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Precondition)
	// condition_type
	switch x := m.ConditionType.(type) {
	case *Precondition_Exists:
		t := uint64(0)
		if x.Exists {
			t = 1
		}
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Precondition_UpdateTime:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UpdateTime); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Precondition.ConditionType has unexpected type %T", x)
	}
	return nil
}

func _Precondition_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Precondition)
	switch tag {
	case 1: // condition_type.exists
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ConditionType = &Precondition_Exists{x != 0}
		return true, err
	case 2: // condition_type.update_time
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(timestamp.Timestamp)
		err := b.DecodeMessage(msg)
		m.ConditionType = &Precondition_UpdateTime{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Precondition_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Precondition)
	// condition_type
	switch x := m.ConditionType.(type) {
	case *Precondition_Exists:
		n += 1 // tag and wire
		n += 1
	case *Precondition_UpdateTime:
		s := proto.Size(x.UpdateTime)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Options for creating a new transaction.
type TransactionOptions struct {
	// The mode of the transaction.
	//
	// Types that are valid to be assigned to Mode:
	//	*TransactionOptions_ReadOnly_
	//	*TransactionOptions_ReadWrite_
	Mode                 isTransactionOptions_Mode `protobuf_oneof:"mode"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *TransactionOptions) Reset()         { *m = TransactionOptions{} }
func (m *TransactionOptions) String() string { return proto.CompactTextString(m) }
func (*TransactionOptions) ProtoMessage()    {}
func (*TransactionOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_33e7705eedc91e86, []int{2}
}
func (m *TransactionOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionOptions.Unmarshal(m, b)
}
func (m *TransactionOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionOptions.Marshal(b, m, deterministic)
}
func (dst *TransactionOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionOptions.Merge(dst, src)
}
func (m *TransactionOptions) XXX_Size() int {
	return xxx_messageInfo_TransactionOptions.Size(m)
}
func (m *TransactionOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionOptions.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionOptions proto.InternalMessageInfo

type isTransactionOptions_Mode interface {
	isTransactionOptions_Mode()
}

type TransactionOptions_ReadOnly_ struct {
	ReadOnly *TransactionOptions_ReadOnly `protobuf:"bytes,2,opt,name=read_only,json=readOnly,proto3,oneof"`
}

type TransactionOptions_ReadWrite_ struct {
	ReadWrite *TransactionOptions_ReadWrite `protobuf:"bytes,3,opt,name=read_write,json=readWrite,proto3,oneof"`
}

func (*TransactionOptions_ReadOnly_) isTransactionOptions_Mode() {}

func (*TransactionOptions_ReadWrite_) isTransactionOptions_Mode() {}

func (m *TransactionOptions) GetMode() isTransactionOptions_Mode {
	if m != nil {
		return m.Mode
	}
	return nil
}

func (m *TransactionOptions) GetReadOnly() *TransactionOptions_ReadOnly {
	if x, ok := m.GetMode().(*TransactionOptions_ReadOnly_); ok {
		return x.ReadOnly
	}
	return nil
}

func (m *TransactionOptions) GetReadWrite() *TransactionOptions_ReadWrite {
	if x, ok := m.GetMode().(*TransactionOptions_ReadWrite_); ok {
		return x.ReadWrite
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TransactionOptions) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TransactionOptions_OneofMarshaler, _TransactionOptions_OneofUnmarshaler, _TransactionOptions_OneofSizer, []interface{}{
		(*TransactionOptions_ReadOnly_)(nil),
		(*TransactionOptions_ReadWrite_)(nil),
	}
}

func _TransactionOptions_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TransactionOptions)
	// mode
	switch x := m.Mode.(type) {
	case *TransactionOptions_ReadOnly_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ReadOnly); err != nil {
			return err
		}
	case *TransactionOptions_ReadWrite_:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ReadWrite); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TransactionOptions.Mode has unexpected type %T", x)
	}
	return nil
}

func _TransactionOptions_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TransactionOptions)
	switch tag {
	case 2: // mode.read_only
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TransactionOptions_ReadOnly)
		err := b.DecodeMessage(msg)
		m.Mode = &TransactionOptions_ReadOnly_{msg}
		return true, err
	case 3: // mode.read_write
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TransactionOptions_ReadWrite)
		err := b.DecodeMessage(msg)
		m.Mode = &TransactionOptions_ReadWrite_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TransactionOptions_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TransactionOptions)
	// mode
	switch x := m.Mode.(type) {
	case *TransactionOptions_ReadOnly_:
		s := proto.Size(x.ReadOnly)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TransactionOptions_ReadWrite_:
		s := proto.Size(x.ReadWrite)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Options for a transaction that can be used to read and write documents.
type TransactionOptions_ReadWrite struct {
	// An optional transaction to retry.
	RetryTransaction     []byte   `protobuf:"bytes,1,opt,name=retry_transaction,json=retryTransaction,proto3" json:"retry_transaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionOptions_ReadWrite) Reset()         { *m = TransactionOptions_ReadWrite{} }
func (m *TransactionOptions_ReadWrite) String() string { return proto.CompactTextString(m) }
func (*TransactionOptions_ReadWrite) ProtoMessage()    {}
func (*TransactionOptions_ReadWrite) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_33e7705eedc91e86, []int{2, 0}
}
func (m *TransactionOptions_ReadWrite) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionOptions_ReadWrite.Unmarshal(m, b)
}
func (m *TransactionOptions_ReadWrite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionOptions_ReadWrite.Marshal(b, m, deterministic)
}
func (dst *TransactionOptions_ReadWrite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionOptions_ReadWrite.Merge(dst, src)
}
func (m *TransactionOptions_ReadWrite) XXX_Size() int {
	return xxx_messageInfo_TransactionOptions_ReadWrite.Size(m)
}
func (m *TransactionOptions_ReadWrite) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionOptions_ReadWrite.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionOptions_ReadWrite proto.InternalMessageInfo

func (m *TransactionOptions_ReadWrite) GetRetryTransaction() []byte {
	if m != nil {
		return m.RetryTransaction
	}
	return nil
}

// Options for a transaction that can only be used to read documents.
type TransactionOptions_ReadOnly struct {
	// The consistency mode for this transaction. If not set, defaults to strong
	// consistency.
	//
	// Types that are valid to be assigned to ConsistencySelector:
	//	*TransactionOptions_ReadOnly_ReadTime
	ConsistencySelector  isTransactionOptions_ReadOnly_ConsistencySelector `protobuf_oneof:"consistency_selector"`
	XXX_NoUnkeyedLiteral struct{}                                          `json:"-"`
	XXX_unrecognized     []byte                                            `json:"-"`
	XXX_sizecache        int32                                             `json:"-"`
}

func (m *TransactionOptions_ReadOnly) Reset()         { *m = TransactionOptions_ReadOnly{} }
func (m *TransactionOptions_ReadOnly) String() string { return proto.CompactTextString(m) }
func (*TransactionOptions_ReadOnly) ProtoMessage()    {}
func (*TransactionOptions_ReadOnly) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_33e7705eedc91e86, []int{2, 1}
}
func (m *TransactionOptions_ReadOnly) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionOptions_ReadOnly.Unmarshal(m, b)
}
func (m *TransactionOptions_ReadOnly) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionOptions_ReadOnly.Marshal(b, m, deterministic)
}
func (dst *TransactionOptions_ReadOnly) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionOptions_ReadOnly.Merge(dst, src)
}
func (m *TransactionOptions_ReadOnly) XXX_Size() int {
	return xxx_messageInfo_TransactionOptions_ReadOnly.Size(m)
}
func (m *TransactionOptions_ReadOnly) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionOptions_ReadOnly.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionOptions_ReadOnly proto.InternalMessageInfo

type isTransactionOptions_ReadOnly_ConsistencySelector interface {
	isTransactionOptions_ReadOnly_ConsistencySelector()
}

type TransactionOptions_ReadOnly_ReadTime struct {
	ReadTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=read_time,json=readTime,proto3,oneof"`
}

func (*TransactionOptions_ReadOnly_ReadTime) isTransactionOptions_ReadOnly_ConsistencySelector() {}

func (m *TransactionOptions_ReadOnly) GetConsistencySelector() isTransactionOptions_ReadOnly_ConsistencySelector {
	if m != nil {
		return m.ConsistencySelector
	}
	return nil
}

func (m *TransactionOptions_ReadOnly) GetReadTime() *timestamp.Timestamp {
	if x, ok := m.GetConsistencySelector().(*TransactionOptions_ReadOnly_ReadTime); ok {
		return x.ReadTime
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TransactionOptions_ReadOnly) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TransactionOptions_ReadOnly_OneofMarshaler, _TransactionOptions_ReadOnly_OneofUnmarshaler, _TransactionOptions_ReadOnly_OneofSizer, []interface{}{
		(*TransactionOptions_ReadOnly_ReadTime)(nil),
	}
}

func _TransactionOptions_ReadOnly_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TransactionOptions_ReadOnly)
	// consistency_selector
	switch x := m.ConsistencySelector.(type) {
	case *TransactionOptions_ReadOnly_ReadTime:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ReadTime); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TransactionOptions_ReadOnly.ConsistencySelector has unexpected type %T", x)
	}
	return nil
}

func _TransactionOptions_ReadOnly_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TransactionOptions_ReadOnly)
	switch tag {
	case 2: // consistency_selector.read_time
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(timestamp.Timestamp)
		err := b.DecodeMessage(msg)
		m.ConsistencySelector = &TransactionOptions_ReadOnly_ReadTime{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TransactionOptions_ReadOnly_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TransactionOptions_ReadOnly)
	// consistency_selector
	switch x := m.ConsistencySelector.(type) {
	case *TransactionOptions_ReadOnly_ReadTime:
		s := proto.Size(x.ReadTime)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*DocumentMask)(nil), "google.firestore.v1beta1.DocumentMask")
	proto.RegisterType((*Precondition)(nil), "google.firestore.v1beta1.Precondition")
	proto.RegisterType((*TransactionOptions)(nil), "google.firestore.v1beta1.TransactionOptions")
	proto.RegisterType((*TransactionOptions_ReadWrite)(nil), "google.firestore.v1beta1.TransactionOptions.ReadWrite")
	proto.RegisterType((*TransactionOptions_ReadOnly)(nil), "google.firestore.v1beta1.TransactionOptions.ReadOnly")
}

func init() {
	proto.RegisterFile("google/firestore/v1beta1/common.proto", fileDescriptor_common_33e7705eedc91e86)
}

var fileDescriptor_common_33e7705eedc91e86 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xe1, 0x8a, 0xd3, 0x40,
	0x10, 0xc7, 0x9b, 0xde, 0x71, 0xb4, 0xd3, 0x22, 0x67, 0x10, 0x09, 0xe1, 0xf0, 0x8e, 0x82, 0x50,
	0x10, 0x36, 0x54, 0x51, 0x14, 0xf1, 0x83, 0xa9, 0xdc, 0xf5, 0x8b, 0x5c, 0x89, 0xe5, 0x0e, 0xa4,
	0x12, 0xb6, 0xc9, 0x34, 0x2e, 0x26, 0x3b, 0x61, 0x77, 0xab, 0xe6, 0x75, 0xfc, 0xe8, 0x1b, 0xf8,
	0x0a, 0x3e, 0x87, 0x0f, 0x22, 0xd9, 0xa4, 0x51, 0x38, 0x0e, 0xf4, 0x5b, 0x32, 0xf3, 0x9b, 0xff,
	0xfc, 0xff, 0xc3, 0xc2, 0xc3, 0x8c, 0x28, 0xcb, 0x31, 0xd8, 0x0a, 0x85, 0xda, 0x90, 0xc2, 0xe0,
	0xf3, 0x6c, 0x83, 0x86, 0xcf, 0x82, 0x84, 0x8a, 0x82, 0x24, 0x2b, 0x15, 0x19, 0x72, 0xbd, 0x06,
	0x63, 0x1d, 0xc6, 0x5a, 0xcc, 0x3f, 0x6d, 0x05, 0x2c, 0xb7, 0xd9, 0x6d, 0x03, 0x23, 0x0a, 0xd4,
	0x86, 0x17, 0x65, 0x33, 0xea, 0x9f, 0xb4, 0x00, 0x2f, 0x45, 0xc0, 0xa5, 0x24, 0xc3, 0x8d, 0x20,
	0xa9, 0x9b, 0xee, 0x24, 0x80, 0xf1, 0x1b, 0x4a, 0x76, 0x05, 0x4a, 0xf3, 0x96, 0xeb, 0x4f, 0xee,
	0x29, 0x8c, 0xb6, 0x02, 0xf3, 0x34, 0x2e, 0xb9, 0xf9, 0xa8, 0x3d, 0xe7, 0xec, 0x60, 0x3a, 0x8c,
	0xc0, 0x96, 0x96, 0x75, 0x65, 0x52, 0xc1, 0x78, 0xa9, 0x30, 0x21, 0x99, 0x8a, 0x5a, 0xc7, 0xf5,
	0xe0, 0x08, 0xbf, 0x0a, 0x6d, 0x6a, 0xd6, 0x99, 0x0e, 0x16, 0xbd, 0xa8, 0xfd, 0x77, 0x5f, 0xc1,
	0x68, 0x57, 0xa6, 0xdc, 0x60, 0x5c, 0x5b, 0xf2, 0xfa, 0x67, 0xce, 0x74, 0xf4, 0xd8, 0x67, 0x6d,
	0x92, 0xbd, 0x5f, 0xb6, 0xda, 0xfb, 0x5d, 0xf4, 0x22, 0x68, 0x06, 0xea, 0x52, 0x78, 0x0c, 0x77,
	0xba, 0x2d, 0xb1, 0xa9, 0x4a, 0x9c, 0xfc, 0xea, 0x83, 0xbb, 0x52, 0x5c, 0x6a, 0x9e, 0xd4, 0xc5,
	0xcb, 0xd2, 0x06, 0x71, 0x57, 0x30, 0x54, 0xc8, 0xd3, 0x98, 0x64, 0x5e, 0xb5, 0x5b, 0x9e, 0xb2,
	0xdb, 0xee, 0xc5, 0x6e, 0x0a, 0xb0, 0x08, 0x79, 0x7a, 0x29, 0xf3, 0x6a, 0xd1, 0x8b, 0x06, 0xaa,
	0xfd, 0x76, 0xaf, 0x01, 0xac, 0xea, 0x17, 0x25, 0x0c, 0x7a, 0x07, 0x56, 0xf6, 0xd9, 0x7f, 0xcb,
	0x5e, 0xd7, 0xd3, 0x8b, 0x5e, 0x64, 0x1d, 0xda, 0x1f, 0xff, 0x39, 0x0c, 0xbb, 0x8e, 0xfb, 0x08,
	0xee, 0x2a, 0x34, 0xaa, 0x8a, 0xcd, 0x9f, 0x79, 0x7b, 0xc8, 0x71, 0x74, 0x6c, 0x1b, 0x7f, 0xe9,
	0xfa, 0x1f, 0x60, 0xb0, 0xb7, 0xea, 0xbe, 0x68, 0x43, 0xff, 0xf3, 0x69, 0x6d, 0x32, 0x7b, 0xd8,
	0xfb, 0x70, 0x2f, 0x21, 0xa9, 0x85, 0x36, 0x28, 0x93, 0x2a, 0xd6, 0x98, 0x63, 0x62, 0x48, 0x85,
	0x47, 0x70, 0x58, 0x50, 0x8a, 0xe1, 0x0f, 0x07, 0x4e, 0x12, 0x2a, 0x6e, 0xcd, 0x1a, 0x8e, 0xe6,
	0xf6, 0x69, 0x2e, 0xeb, 0x35, 0x4b, 0xe7, 0xfd, 0xeb, 0x16, 0xcc, 0x28, 0xe7, 0x32, 0x63, 0xa4,
	0xb2, 0x20, 0x43, 0x69, 0x4d, 0x04, 0x4d, 0x8b, 0x97, 0x42, 0xdf, 0x7c, 0xe1, 0x2f, 0xbb, 0xca,
	0xb7, 0xfe, 0xe1, 0xc5, 0xfc, 0xfc, 0xdd, 0xf7, 0xfe, 0x83, 0x8b, 0x46, 0x6a, 0x9e, 0xd3, 0x2e,
	0x65, 0xe7, 0xdd, 0xe6, 0xab, 0x59, 0x58, 0x4f, 0xfc, 0xdc, 0x03, 0x6b, 0x0b, 0xac, 0x3b, 0x60,
	0x7d, 0xd5, 0x48, 0x6e, 0x8e, 0xec, 0xda, 0x27, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x64,
	0x00, 0xd8, 0x57, 0x03, 0x00, 0x00,
}
