// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/bigquery/storage/v1beta1/table_reference.proto

package storage // import "google.golang.org/genproto/googleapis/cloud/bigquery/storage/v1beta1"

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

// Table reference that includes just the 3 strings needed to identify a table.
type TableReference struct {
	// The assigned project ID of the project.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// The ID of the dataset in the above project.
	DatasetId string `protobuf:"bytes,2,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	// The ID of the table in the above dataset.
	TableId              string   `protobuf:"bytes,3,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TableReference) Reset()         { *m = TableReference{} }
func (m *TableReference) String() string { return proto.CompactTextString(m) }
func (*TableReference) ProtoMessage()    {}
func (*TableReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_table_reference_d7a619381c8a4dbe, []int{0}
}
func (m *TableReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableReference.Unmarshal(m, b)
}
func (m *TableReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableReference.Marshal(b, m, deterministic)
}
func (dst *TableReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableReference.Merge(dst, src)
}
func (m *TableReference) XXX_Size() int {
	return xxx_messageInfo_TableReference.Size(m)
}
func (m *TableReference) XXX_DiscardUnknown() {
	xxx_messageInfo_TableReference.DiscardUnknown(m)
}

var xxx_messageInfo_TableReference proto.InternalMessageInfo

func (m *TableReference) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *TableReference) GetDatasetId() string {
	if m != nil {
		return m.DatasetId
	}
	return ""
}

func (m *TableReference) GetTableId() string {
	if m != nil {
		return m.TableId
	}
	return ""
}

// All fields in this message optional.
type TableModifiers struct {
	// The snapshot time of the table. If not set, interpreted as now.
	SnapshotTime         *timestamp.Timestamp `protobuf:"bytes,1,opt,name=snapshot_time,json=snapshotTime,proto3" json:"snapshot_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TableModifiers) Reset()         { *m = TableModifiers{} }
func (m *TableModifiers) String() string { return proto.CompactTextString(m) }
func (*TableModifiers) ProtoMessage()    {}
func (*TableModifiers) Descriptor() ([]byte, []int) {
	return fileDescriptor_table_reference_d7a619381c8a4dbe, []int{1}
}
func (m *TableModifiers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableModifiers.Unmarshal(m, b)
}
func (m *TableModifiers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableModifiers.Marshal(b, m, deterministic)
}
func (dst *TableModifiers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableModifiers.Merge(dst, src)
}
func (m *TableModifiers) XXX_Size() int {
	return xxx_messageInfo_TableModifiers.Size(m)
}
func (m *TableModifiers) XXX_DiscardUnknown() {
	xxx_messageInfo_TableModifiers.DiscardUnknown(m)
}

var xxx_messageInfo_TableModifiers proto.InternalMessageInfo

func (m *TableModifiers) GetSnapshotTime() *timestamp.Timestamp {
	if m != nil {
		return m.SnapshotTime
	}
	return nil
}

func init() {
	proto.RegisterType((*TableReference)(nil), "google.cloud.bigquery.storage.v1beta1.TableReference")
	proto.RegisterType((*TableModifiers)(nil), "google.cloud.bigquery.storage.v1beta1.TableModifiers")
}

func init() {
	proto.RegisterFile("google/cloud/bigquery/storage/v1beta1/table_reference.proto", fileDescriptor_table_reference_d7a619381c8a4dbe)
}

var fileDescriptor_table_reference_d7a619381c8a4dbe = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0x86, 0x59, 0x05, 0x75, 0xe3, 0xc7, 0xa1, 0x5e, 0x76, 0x0b, 0xa2, 0x2c, 0x08, 0x7a, 0x49,
	0x58, 0x3d, 0xee, 0x41, 0xd8, 0x5b, 0x41, 0x41, 0xcb, 0x9e, 0xbc, 0x94, 0xb4, 0x99, 0xc6, 0x48,
	0xdb, 0x89, 0x49, 0x2a, 0xf8, 0x27, 0xfc, 0xcd, 0xd2, 0x7c, 0x1c, 0x3c, 0xe9, 0x71, 0xe6, 0xed,
	0x33, 0xcf, 0x4c, 0x43, 0x36, 0x12, 0x51, 0x76, 0xc0, 0x9a, 0x0e, 0x47, 0xc1, 0x6a, 0x25, 0x3f,
	0x46, 0x30, 0x5f, 0xcc, 0x3a, 0x34, 0x5c, 0x02, 0xfb, 0x5c, 0xd7, 0xe0, 0xf8, 0x9a, 0x39, 0x5e,
	0x77, 0x50, 0x19, 0x68, 0xc1, 0xc0, 0xd0, 0x00, 0xd5, 0x06, 0x1d, 0x66, 0xd7, 0x01, 0xa6, 0x1e,
	0xa6, 0x09, 0xa6, 0x11, 0xa6, 0x11, 0xce, 0x97, 0xd1, 0xc1, 0xb5, 0x62, 0x06, 0x2c, 0x8e, 0x26,
	0x4d, 0xc8, 0x2f, 0x63, 0xe4, 0xab, 0x7a, 0x6c, 0x99, 0x53, 0x3d, 0x58, 0xc7, 0x7b, 0x1d, 0x3e,
	0x58, 0x29, 0x72, 0xb6, 0x9b, 0xdc, 0x65, 0x52, 0x67, 0x17, 0x84, 0x68, 0x83, 0xef, 0xd0, 0xb8,
	0x4a, 0x89, 0xc5, 0xec, 0x6a, 0x76, 0x33, 0x2f, 0xe7, 0xb1, 0x53, 0x88, 0x29, 0x16, 0xdc, 0x71,
	0x0b, 0x3e, 0xde, 0x0b, 0x71, 0xec, 0x14, 0x22, 0x5b, 0x92, 0xa3, 0x70, 0x8b, 0x12, 0x8b, 0x7d,
	0x1f, 0x1e, 0xfa, 0xba, 0x10, 0xab, 0x97, 0xa8, 0x7a, 0x42, 0xa1, 0x5a, 0x05, 0xc6, 0x66, 0x0f,
	0xe4, 0xd4, 0x0e, 0x5c, 0xdb, 0x37, 0x74, 0xd5, 0xb4, 0x98, 0xb7, 0x1d, 0xdf, 0xe5, 0x34, 0xde,
	0x9d, 0xb6, 0xa6, 0xbb, 0xb4, 0x75, 0x79, 0x92, 0x80, 0xa9, 0xb5, 0xfd, 0x9e, 0x91, 0xdb, 0x06,
	0x7b, 0xfa, 0xaf, 0xff, 0xb4, 0x3d, 0xff, 0x7d, 0xe9, 0xf3, 0x34, 0xfd, 0xf5, 0x31, 0xb2, 0x12,
	0x3b, 0x3e, 0x48, 0x8a, 0x46, 0x32, 0x09, 0x83, 0x37, 0xb3, 0x10, 0x71, 0xad, 0xec, 0x1f, 0xef,
	0xb7, 0x89, 0x75, 0x7d, 0xe0, 0xc1, 0xfb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0x1c, 0xaf,
	0x7a, 0xf7, 0x01, 0x00, 0x00,
}
