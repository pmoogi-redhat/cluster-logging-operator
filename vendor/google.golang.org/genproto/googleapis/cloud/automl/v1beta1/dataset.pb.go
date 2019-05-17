// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/dataset.proto

package automl // import "google.golang.org/genproto/googleapis/cloud/automl/v1beta1"

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

// A workspace for solving a single, particular machine learning (ML) problem.
// A workspace contains examples that may be annotated.
type Dataset struct {
	// Required.
	// The dataset metadata that is specific to the problem type.
	//
	// Types that are valid to be assigned to DatasetMetadata:
	//	*Dataset_TranslationDatasetMetadata
	//	*Dataset_ImageClassificationDatasetMetadata
	//	*Dataset_TextClassificationDatasetMetadata
	//	*Dataset_ImageObjectDetectionDatasetMetadata
	//	*Dataset_VideoClassificationDatasetMetadata
	//	*Dataset_TextExtractionDatasetMetadata
	//	*Dataset_TextSentimentDatasetMetadata
	//	*Dataset_TablesDatasetMetadata
	DatasetMetadata isDataset_DatasetMetadata `protobuf_oneof:"dataset_metadata"`
	// Output only. The resource name of the dataset.
	// Form: `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The name of the dataset to show in the interface. The name can be
	// up to 32 characters long and can consist only of ASCII Latin letters A-Z
	// and a-z, underscores
	// (_), and ASCII digits 0-9.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// User-provided description of the dataset. The description can be up to
	// 25000 characters long.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. The number of examples in the dataset.
	ExampleCount int32 `protobuf:"varint,21,opt,name=example_count,json=exampleCount,proto3" json:"example_count,omitempty"`
	// Output only. Timestamp when this dataset was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,14,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Used to perform consistent read-modify-write updates. If not set, a blind
	// "overwrite" update happens.
	Etag                 string   `protobuf:"bytes,17,opt,name=etag,proto3" json:"etag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dataset) Reset()         { *m = Dataset{} }
func (m *Dataset) String() string { return proto.CompactTextString(m) }
func (*Dataset) ProtoMessage()    {}
func (*Dataset) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataset_954237b3f3904ab8, []int{0}
}
func (m *Dataset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dataset.Unmarshal(m, b)
}
func (m *Dataset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dataset.Marshal(b, m, deterministic)
}
func (dst *Dataset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dataset.Merge(dst, src)
}
func (m *Dataset) XXX_Size() int {
	return xxx_messageInfo_Dataset.Size(m)
}
func (m *Dataset) XXX_DiscardUnknown() {
	xxx_messageInfo_Dataset.DiscardUnknown(m)
}

var xxx_messageInfo_Dataset proto.InternalMessageInfo

type isDataset_DatasetMetadata interface {
	isDataset_DatasetMetadata()
}

type Dataset_TranslationDatasetMetadata struct {
	TranslationDatasetMetadata *TranslationDatasetMetadata `protobuf:"bytes,23,opt,name=translation_dataset_metadata,json=translationDatasetMetadata,proto3,oneof"`
}

type Dataset_ImageClassificationDatasetMetadata struct {
	ImageClassificationDatasetMetadata *ImageClassificationDatasetMetadata `protobuf:"bytes,24,opt,name=image_classification_dataset_metadata,json=imageClassificationDatasetMetadata,proto3,oneof"`
}

type Dataset_TextClassificationDatasetMetadata struct {
	TextClassificationDatasetMetadata *TextClassificationDatasetMetadata `protobuf:"bytes,25,opt,name=text_classification_dataset_metadata,json=textClassificationDatasetMetadata,proto3,oneof"`
}

type Dataset_ImageObjectDetectionDatasetMetadata struct {
	ImageObjectDetectionDatasetMetadata *ImageObjectDetectionDatasetMetadata `protobuf:"bytes,26,opt,name=image_object_detection_dataset_metadata,json=imageObjectDetectionDatasetMetadata,proto3,oneof"`
}

type Dataset_VideoClassificationDatasetMetadata struct {
	VideoClassificationDatasetMetadata *VideoClassificationDatasetMetadata `protobuf:"bytes,31,opt,name=video_classification_dataset_metadata,json=videoClassificationDatasetMetadata,proto3,oneof"`
}

type Dataset_TextExtractionDatasetMetadata struct {
	TextExtractionDatasetMetadata *TextExtractionDatasetMetadata `protobuf:"bytes,28,opt,name=text_extraction_dataset_metadata,json=textExtractionDatasetMetadata,proto3,oneof"`
}

type Dataset_TextSentimentDatasetMetadata struct {
	TextSentimentDatasetMetadata *TextSentimentDatasetMetadata `protobuf:"bytes,30,opt,name=text_sentiment_dataset_metadata,json=textSentimentDatasetMetadata,proto3,oneof"`
}

type Dataset_TablesDatasetMetadata struct {
	TablesDatasetMetadata *TablesDatasetMetadata `protobuf:"bytes,33,opt,name=tables_dataset_metadata,json=tablesDatasetMetadata,proto3,oneof"`
}

func (*Dataset_TranslationDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_ImageClassificationDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_TextClassificationDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_ImageObjectDetectionDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_VideoClassificationDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_TextExtractionDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_TextSentimentDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_TablesDatasetMetadata) isDataset_DatasetMetadata() {}

func (m *Dataset) GetDatasetMetadata() isDataset_DatasetMetadata {
	if m != nil {
		return m.DatasetMetadata
	}
	return nil
}

func (m *Dataset) GetTranslationDatasetMetadata() *TranslationDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_TranslationDatasetMetadata); ok {
		return x.TranslationDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetImageClassificationDatasetMetadata() *ImageClassificationDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_ImageClassificationDatasetMetadata); ok {
		return x.ImageClassificationDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetTextClassificationDatasetMetadata() *TextClassificationDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_TextClassificationDatasetMetadata); ok {
		return x.TextClassificationDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetImageObjectDetectionDatasetMetadata() *ImageObjectDetectionDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_ImageObjectDetectionDatasetMetadata); ok {
		return x.ImageObjectDetectionDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetVideoClassificationDatasetMetadata() *VideoClassificationDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_VideoClassificationDatasetMetadata); ok {
		return x.VideoClassificationDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetTextExtractionDatasetMetadata() *TextExtractionDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_TextExtractionDatasetMetadata); ok {
		return x.TextExtractionDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetTextSentimentDatasetMetadata() *TextSentimentDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_TextSentimentDatasetMetadata); ok {
		return x.TextSentimentDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetTablesDatasetMetadata() *TablesDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_TablesDatasetMetadata); ok {
		return x.TablesDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Dataset) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Dataset) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Dataset) GetExampleCount() int32 {
	if m != nil {
		return m.ExampleCount
	}
	return 0
}

func (m *Dataset) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Dataset) GetEtag() string {
	if m != nil {
		return m.Etag
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Dataset) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Dataset_OneofMarshaler, _Dataset_OneofUnmarshaler, _Dataset_OneofSizer, []interface{}{
		(*Dataset_TranslationDatasetMetadata)(nil),
		(*Dataset_ImageClassificationDatasetMetadata)(nil),
		(*Dataset_TextClassificationDatasetMetadata)(nil),
		(*Dataset_ImageObjectDetectionDatasetMetadata)(nil),
		(*Dataset_VideoClassificationDatasetMetadata)(nil),
		(*Dataset_TextExtractionDatasetMetadata)(nil),
		(*Dataset_TextSentimentDatasetMetadata)(nil),
		(*Dataset_TablesDatasetMetadata)(nil),
	}
}

func _Dataset_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Dataset)
	// dataset_metadata
	switch x := m.DatasetMetadata.(type) {
	case *Dataset_TranslationDatasetMetadata:
		b.EncodeVarint(23<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TranslationDatasetMetadata); err != nil {
			return err
		}
	case *Dataset_ImageClassificationDatasetMetadata:
		b.EncodeVarint(24<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ImageClassificationDatasetMetadata); err != nil {
			return err
		}
	case *Dataset_TextClassificationDatasetMetadata:
		b.EncodeVarint(25<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TextClassificationDatasetMetadata); err != nil {
			return err
		}
	case *Dataset_ImageObjectDetectionDatasetMetadata:
		b.EncodeVarint(26<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ImageObjectDetectionDatasetMetadata); err != nil {
			return err
		}
	case *Dataset_VideoClassificationDatasetMetadata:
		b.EncodeVarint(31<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.VideoClassificationDatasetMetadata); err != nil {
			return err
		}
	case *Dataset_TextExtractionDatasetMetadata:
		b.EncodeVarint(28<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TextExtractionDatasetMetadata); err != nil {
			return err
		}
	case *Dataset_TextSentimentDatasetMetadata:
		b.EncodeVarint(30<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TextSentimentDatasetMetadata); err != nil {
			return err
		}
	case *Dataset_TablesDatasetMetadata:
		b.EncodeVarint(33<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TablesDatasetMetadata); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Dataset.DatasetMetadata has unexpected type %T", x)
	}
	return nil
}

func _Dataset_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Dataset)
	switch tag {
	case 23: // dataset_metadata.translation_dataset_metadata
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TranslationDatasetMetadata)
		err := b.DecodeMessage(msg)
		m.DatasetMetadata = &Dataset_TranslationDatasetMetadata{msg}
		return true, err
	case 24: // dataset_metadata.image_classification_dataset_metadata
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ImageClassificationDatasetMetadata)
		err := b.DecodeMessage(msg)
		m.DatasetMetadata = &Dataset_ImageClassificationDatasetMetadata{msg}
		return true, err
	case 25: // dataset_metadata.text_classification_dataset_metadata
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TextClassificationDatasetMetadata)
		err := b.DecodeMessage(msg)
		m.DatasetMetadata = &Dataset_TextClassificationDatasetMetadata{msg}
		return true, err
	case 26: // dataset_metadata.image_object_detection_dataset_metadata
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ImageObjectDetectionDatasetMetadata)
		err := b.DecodeMessage(msg)
		m.DatasetMetadata = &Dataset_ImageObjectDetectionDatasetMetadata{msg}
		return true, err
	case 31: // dataset_metadata.video_classification_dataset_metadata
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(VideoClassificationDatasetMetadata)
		err := b.DecodeMessage(msg)
		m.DatasetMetadata = &Dataset_VideoClassificationDatasetMetadata{msg}
		return true, err
	case 28: // dataset_metadata.text_extraction_dataset_metadata
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TextExtractionDatasetMetadata)
		err := b.DecodeMessage(msg)
		m.DatasetMetadata = &Dataset_TextExtractionDatasetMetadata{msg}
		return true, err
	case 30: // dataset_metadata.text_sentiment_dataset_metadata
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TextSentimentDatasetMetadata)
		err := b.DecodeMessage(msg)
		m.DatasetMetadata = &Dataset_TextSentimentDatasetMetadata{msg}
		return true, err
	case 33: // dataset_metadata.tables_dataset_metadata
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TablesDatasetMetadata)
		err := b.DecodeMessage(msg)
		m.DatasetMetadata = &Dataset_TablesDatasetMetadata{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Dataset_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Dataset)
	// dataset_metadata
	switch x := m.DatasetMetadata.(type) {
	case *Dataset_TranslationDatasetMetadata:
		s := proto.Size(x.TranslationDatasetMetadata)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Dataset_ImageClassificationDatasetMetadata:
		s := proto.Size(x.ImageClassificationDatasetMetadata)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Dataset_TextClassificationDatasetMetadata:
		s := proto.Size(x.TextClassificationDatasetMetadata)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Dataset_ImageObjectDetectionDatasetMetadata:
		s := proto.Size(x.ImageObjectDetectionDatasetMetadata)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Dataset_VideoClassificationDatasetMetadata:
		s := proto.Size(x.VideoClassificationDatasetMetadata)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Dataset_TextExtractionDatasetMetadata:
		s := proto.Size(x.TextExtractionDatasetMetadata)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Dataset_TextSentimentDatasetMetadata:
		s := proto.Size(x.TextSentimentDatasetMetadata)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Dataset_TablesDatasetMetadata:
		s := proto.Size(x.TablesDatasetMetadata)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A definition of an annotation.
type AnnotationSpec struct {
	// Output only. Resource name of the annotation spec.
	// Form:
	//
	// 'projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/annotationSpecs/{annotation_spec_id}'
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required.
	// The name of the annotation spec to show in the interface. The name can be
	// up to 32 characters long and can consist only of ASCII Latin letters A-Z
	// and a-z, underscores
	// (_), and ASCII digits 0-9.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Output only. The number of examples in the parent dataset
	// labeled by the annotation spec.
	ExampleCount         int32    `protobuf:"varint,9,opt,name=example_count,json=exampleCount,proto3" json:"example_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnnotationSpec) Reset()         { *m = AnnotationSpec{} }
func (m *AnnotationSpec) String() string { return proto.CompactTextString(m) }
func (*AnnotationSpec) ProtoMessage()    {}
func (*AnnotationSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataset_954237b3f3904ab8, []int{1}
}
func (m *AnnotationSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnnotationSpec.Unmarshal(m, b)
}
func (m *AnnotationSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnnotationSpec.Marshal(b, m, deterministic)
}
func (dst *AnnotationSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnnotationSpec.Merge(dst, src)
}
func (m *AnnotationSpec) XXX_Size() int {
	return xxx_messageInfo_AnnotationSpec.Size(m)
}
func (m *AnnotationSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_AnnotationSpec.DiscardUnknown(m)
}

var xxx_messageInfo_AnnotationSpec proto.InternalMessageInfo

func (m *AnnotationSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AnnotationSpec) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *AnnotationSpec) GetExampleCount() int32 {
	if m != nil {
		return m.ExampleCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Dataset)(nil), "google.cloud.automl.v1beta1.Dataset")
	proto.RegisterType((*AnnotationSpec)(nil), "google.cloud.automl.v1beta1.AnnotationSpec")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/dataset.proto", fileDescriptor_dataset_954237b3f3904ab8)
}

var fileDescriptor_dataset_954237b3f3904ab8 = []byte{
	// 647 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0x5d, 0x6b, 0xd4, 0x4c,
	0x14, 0xc7, 0x9f, 0xf4, 0xf1, 0x85, 0xce, 0xd6, 0xa2, 0x03, 0xa5, 0x71, 0xbb, 0xba, 0xdb, 0x56,
	0xed, 0x0a, 0x9a, 0xd0, 0x2a, 0x88, 0x16, 0xd4, 0xbe, 0x88, 0x7a, 0x51, 0x95, 0x6d, 0xe9, 0x85,
	0x14, 0xc2, 0x6c, 0x72, 0x1a, 0x46, 0x26, 0x99, 0x90, 0x9c, 0x2d, 0x5b, 0xbc, 0x13, 0xfd, 0x00,
	0x82, 0x17, 0x7e, 0x1d, 0x6f, 0xfd, 0x54, 0x32, 0x93, 0x49, 0xab, 0xa6, 0x9b, 0x29, 0xde, 0x25,
	0xe7, 0xfc, 0xe6, 0x7f, 0xfe, 0xe7, 0x9c, 0xdd, 0x09, 0xb9, 0x1b, 0x4b, 0x19, 0x0b, 0xf0, 0x43,
	0x21, 0x47, 0x91, 0xcf, 0x46, 0x28, 0x13, 0xe1, 0x1f, 0xad, 0x0e, 0x01, 0xd9, 0xaa, 0x1f, 0x31,
	0x64, 0x05, 0xa0, 0x97, 0xe5, 0x12, 0x25, 0x5d, 0x28, 0x51, 0x4f, 0xa3, 0x5e, 0x89, 0x7a, 0x06,
	0x6d, 0x77, 0x8c, 0x0e, 0xcb, 0xb8, 0xcf, 0xd2, 0x54, 0x22, 0x43, 0x2e, 0xd3, 0xa2, 0x3c, 0xda,
	0x7e, 0xd8, 0x54, 0xe5, 0x14, 0x0f, 0x32, 0x76, 0x2c, 0x24, 0x8b, 0xcc, 0xa9, 0x7b, 0x36, 0x6f,
	0x01, 0x47, 0x48, 0xaa, 0x1a, 0x2b, 0x4d, 0x34, 0x4f, 0x58, 0x0c, 0x06, 0xec, 0x37, 0x81, 0xc8,
	0x86, 0x02, 0x2a, 0xc9, 0x3b, 0x8d, 0x24, 0x8c, 0xcd, 0x64, 0xda, 0xf7, 0x1b, 0xb9, 0x9c, 0xa5,
	0x85, 0xd0, 0xfd, 0x9d, 0xc7, 0xe9, 0x11, 0x8f, 0x40, 0x1a, 0xb0, 0x6b, 0x40, 0xfd, 0x36, 0x1c,
	0x1d, 0xfa, 0xc8, 0x13, 0x28, 0x90, 0x25, 0x59, 0x09, 0x2c, 0xfd, 0x20, 0xe4, 0xf2, 0x76, 0xb9,
	0x24, 0xfa, 0x91, 0x74, 0x7e, 0x2b, 0x15, 0x98, 0xdd, 0x05, 0x09, 0x20, 0x53, 0xcf, 0xee, 0x7c,
	0xcf, 0xe9, 0xb7, 0xd6, 0x1e, 0x79, 0x0d, 0x5b, 0xf4, 0xf6, 0x4e, 0x05, 0x8c, 0xec, 0x8e, 0x39,
	0xfe, 0xea, 0xbf, 0x41, 0x1b, 0x27, 0x66, 0xe9, 0x37, 0x87, 0xdc, 0xd6, 0x33, 0x0e, 0x42, 0xc1,
	0x8a, 0x82, 0x1f, 0xf2, 0x70, 0x82, 0x0d, 0x57, 0xdb, 0x78, 0xd6, 0x68, 0xe3, 0xb5, 0x52, 0xda,
	0xfa, 0x43, 0xa8, 0x6e, 0x67, 0x89, 0x5b, 0x29, 0xfa, 0xd5, 0x21, 0xb7, 0xd4, 0x9e, 0xac, 0xae,
	0xae, 0x6b, 0x57, 0x4f, 0x9b, 0x87, 0x03, 0x63, 0xb4, 0x99, 0x5a, 0x44, 0x1b, 0x44, 0xbf, 0x3b,
	0x64, 0xa5, 0x1c, 0x95, 0x1c, 0x7e, 0x80, 0x10, 0x83, 0x08, 0x10, 0xc2, 0xb3, 0x6d, 0xb5, 0xb5,
	0xad, 0xe7, 0xf6, 0x61, 0xbd, 0xd5, 0x52, 0xdb, 0x95, 0x52, 0xdd, 0xd8, 0x32, 0xb7, 0x63, 0x7a,
	0x8b, 0xfa, 0xf7, 0x67, 0x9d, 0x57, 0xf7, 0x1c, 0x5b, 0xdc, 0x57, 0x4a, 0xd6, 0x2d, 0x1e, 0x59,
	0x29, 0xfa, 0xc5, 0x21, 0x3d, 0xbd, 0x45, 0x18, 0x63, 0xce, 0x26, 0x8c, 0xaa, 0xa3, 0x1d, 0x3d,
	0xb1, 0x6e, 0xf0, 0xc5, 0x89, 0x46, 0xdd, 0xcc, 0x0d, 0x6c, 0x02, 0xe8, 0x27, 0x87, 0x74, 0xb5,
	0x8f, 0x02, 0x52, 0xf5, 0x4f, 0x4c, 0xb1, 0x6e, 0xe3, 0xa6, 0xb6, 0xf1, 0xd8, 0x6a, 0x63, 0xb7,
	0x92, 0xa8, 0xbb, 0xe8, 0x60, 0x43, 0x9e, 0x0a, 0x32, 0x5f, 0xde, 0x51, 0xf5, 0xda, 0x8b, 0xba,
	0xf6, 0x5a, 0x73, 0x6d, 0x7d, 0xb6, 0x5e, 0x74, 0x0e, 0xcf, 0x4a, 0x50, 0x4a, 0x2e, 0xa4, 0x2c,
	0x01, 0xd7, 0xe9, 0x39, 0xfd, 0xe9, 0x81, 0x7e, 0xa6, 0x8b, 0x64, 0x26, 0xe2, 0x45, 0x26, 0xd8,
	0x71, 0xa0, 0x73, 0x53, 0x3a, 0xd7, 0x32, 0xb1, 0x37, 0x0a, 0xe9, 0x91, 0x56, 0x04, 0x45, 0x98,
	0xf3, 0x4c, 0xcd, 0xd1, 0xfd, 0xdf, 0x10, 0xa7, 0x21, 0xba, 0x4c, 0xae, 0xc0, 0x98, 0x25, 0x99,
	0x80, 0x20, 0x94, 0xa3, 0x14, 0xdd, 0xb9, 0x9e, 0xd3, 0xbf, 0x38, 0x98, 0x31, 0xc1, 0x2d, 0x15,
	0xa3, 0xeb, 0xa4, 0x15, 0xe6, 0xc0, 0x10, 0x02, 0x35, 0x0b, 0x77, 0x56, 0xf7, 0xd7, 0xae, 0xfa,
	0xab, 0x6e, 0x45, 0x6f, 0xaf, 0xba, 0x15, 0x07, 0xa4, 0xc4, 0x55, 0x40, 0x59, 0x07, 0x64, 0xb1,
	0x7b, 0xad, 0xb4, 0xae, 0x9e, 0x37, 0x29, 0xb9, 0xfa, 0xf7, 0xd4, 0x96, 0x04, 0x99, 0xdd, 0x38,
	0xf9, 0x02, 0xed, 0x66, 0x10, 0xfe, 0x6b, 0xd3, 0xb5, 0x96, 0xa6, 0xeb, 0x2d, 0x6d, 0x7e, 0x76,
	0x48, 0x37, 0x94, 0x49, 0xd3, 0x8e, 0xde, 0x39, 0xef, 0x37, 0x4c, 0x3a, 0x96, 0x82, 0xa5, 0xb1,
	0x27, 0xf3, 0xd8, 0x8f, 0x21, 0xd5, 0x0d, 0xfb, 0x65, 0x8a, 0x65, 0xbc, 0x38, 0xf3, 0x03, 0xb2,
	0x5e, 0xbe, 0xfe, 0x9c, 0x5a, 0x78, 0xa9, 0xc1, 0x83, 0x2d, 0x05, 0x1d, 0x6c, 0x8c, 0x50, 0xee,
	0x88, 0x83, 0xfd, 0x12, 0x1a, 0x5e, 0xd2, 0x5a, 0x0f, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xa6,
	0x6b, 0xbf, 0xc3, 0xff, 0x07, 0x00, 0x00,
}
