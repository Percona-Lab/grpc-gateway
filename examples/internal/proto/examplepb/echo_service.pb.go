// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: examples/internal/proto/examplepb/echo_service.proto

// Echo Service
//
// Echo Service API consists of a single service which returns
// a message.

package examplepb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Embedded represents a message embedded in SimpleMessage.
type Embedded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Mark:
	//	*Embedded_Progress
	//	*Embedded_Note
	Mark isEmbedded_Mark `protobuf_oneof:"mark"`
}

func (x *Embedded) Reset() {
	*x = Embedded{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_echo_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Embedded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Embedded) ProtoMessage() {}

func (x *Embedded) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_echo_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Embedded.ProtoReflect.Descriptor instead.
func (*Embedded) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_echo_service_proto_rawDescGZIP(), []int{0}
}

func (m *Embedded) GetMark() isEmbedded_Mark {
	if m != nil {
		return m.Mark
	}
	return nil
}

func (x *Embedded) GetProgress() int64 {
	if x, ok := x.GetMark().(*Embedded_Progress); ok {
		return x.Progress
	}
	return 0
}

func (x *Embedded) GetNote() string {
	if x, ok := x.GetMark().(*Embedded_Note); ok {
		return x.Note
	}
	return ""
}

type isEmbedded_Mark interface {
	isEmbedded_Mark()
}

type Embedded_Progress struct {
	Progress int64 `protobuf:"varint,1,opt,name=progress,proto3,oneof"`
}

type Embedded_Note struct {
	Note string `protobuf:"bytes,2,opt,name=note,proto3,oneof"`
}

func (*Embedded_Progress) isEmbedded_Mark() {}

func (*Embedded_Note) isEmbedded_Mark() {}

// SimpleMessage represents a simple message sent to the Echo service.
type SimpleMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id represents the message identifier.
	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Num int64  `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	// Types that are assignable to Code:
	//	*SimpleMessage_LineNum
	//	*SimpleMessage_Lang
	Code   isSimpleMessage_Code `protobuf_oneof:"code"`
	Status *Embedded            `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	// Types that are assignable to Ext:
	//	*SimpleMessage_En
	//	*SimpleMessage_No
	Ext isSimpleMessage_Ext `protobuf_oneof:"ext"`
}

func (x *SimpleMessage) Reset() {
	*x = SimpleMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_echo_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleMessage) ProtoMessage() {}

func (x *SimpleMessage) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_echo_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleMessage.ProtoReflect.Descriptor instead.
func (*SimpleMessage) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_echo_service_proto_rawDescGZIP(), []int{1}
}

func (x *SimpleMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SimpleMessage) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (m *SimpleMessage) GetCode() isSimpleMessage_Code {
	if m != nil {
		return m.Code
	}
	return nil
}

func (x *SimpleMessage) GetLineNum() int64 {
	if x, ok := x.GetCode().(*SimpleMessage_LineNum); ok {
		return x.LineNum
	}
	return 0
}

func (x *SimpleMessage) GetLang() string {
	if x, ok := x.GetCode().(*SimpleMessage_Lang); ok {
		return x.Lang
	}
	return ""
}

func (x *SimpleMessage) GetStatus() *Embedded {
	if x != nil {
		return x.Status
	}
	return nil
}

func (m *SimpleMessage) GetExt() isSimpleMessage_Ext {
	if m != nil {
		return m.Ext
	}
	return nil
}

func (x *SimpleMessage) GetEn() int64 {
	if x, ok := x.GetExt().(*SimpleMessage_En); ok {
		return x.En
	}
	return 0
}

func (x *SimpleMessage) GetNo() *Embedded {
	if x, ok := x.GetExt().(*SimpleMessage_No); ok {
		return x.No
	}
	return nil
}

type isSimpleMessage_Code interface {
	isSimpleMessage_Code()
}

type SimpleMessage_LineNum struct {
	LineNum int64 `protobuf:"varint,3,opt,name=line_num,json=lineNum,proto3,oneof"`
}

type SimpleMessage_Lang struct {
	Lang string `protobuf:"bytes,4,opt,name=lang,proto3,oneof"`
}

func (*SimpleMessage_LineNum) isSimpleMessage_Code() {}

func (*SimpleMessage_Lang) isSimpleMessage_Code() {}

type isSimpleMessage_Ext interface {
	isSimpleMessage_Ext()
}

type SimpleMessage_En struct {
	En int64 `protobuf:"varint,6,opt,name=en,proto3,oneof"`
}

type SimpleMessage_No struct {
	No *Embedded `protobuf:"bytes,7,opt,name=no,proto3,oneof"`
}

func (*SimpleMessage_En) isSimpleMessage_Ext() {}

func (*SimpleMessage_No) isSimpleMessage_Ext() {}

// DynamicMessage represents a message which can have its structure
// built dynamically using Struct and Values.
type DynamicMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StructField *structpb.Struct `protobuf:"bytes,1,opt,name=struct_field,json=structField,proto3" json:"struct_field,omitempty"`
	ValueField  *structpb.Value  `protobuf:"bytes,2,opt,name=value_field,json=valueField,proto3" json:"value_field,omitempty"`
}

func (x *DynamicMessage) Reset() {
	*x = DynamicMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_echo_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DynamicMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynamicMessage) ProtoMessage() {}

func (x *DynamicMessage) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_echo_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynamicMessage.ProtoReflect.Descriptor instead.
func (*DynamicMessage) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_echo_service_proto_rawDescGZIP(), []int{2}
}

func (x *DynamicMessage) GetStructField() *structpb.Struct {
	if x != nil {
		return x.StructField
	}
	return nil
}

func (x *DynamicMessage) GetValueField() *structpb.Value {
	if x != nil {
		return x.ValueField
	}
	return nil
}

type DynamicMessageUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body       *DynamicMessage        `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *DynamicMessageUpdate) Reset() {
	*x = DynamicMessageUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_echo_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DynamicMessageUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynamicMessageUpdate) ProtoMessage() {}

func (x *DynamicMessageUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_echo_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynamicMessageUpdate.ProtoReflect.Descriptor instead.
func (*DynamicMessageUpdate) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_echo_service_proto_rawDescGZIP(), []int{3}
}

func (x *DynamicMessageUpdate) GetBody() *DynamicMessage {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *DynamicMessageUpdate) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

var File_examples_internal_proto_examplepb_echo_service_proto protoreflect.FileDescriptor

var file_examples_internal_proto_examplepb_echo_service_proto_rawDesc = []byte{
	0x0a, 0x34, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x70, 0x62, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x08, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64,
	0x12, 0x1c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14,
	0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04,
	0x6e, 0x6f, 0x74, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0xa3, 0x02, 0x0a,
	0x0d, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d,
	0x12, 0x1b, 0x0a, 0x08, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x00, 0x52, 0x07, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x14, 0x0a,
	0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6c,
	0x61, 0x6e, 0x67, 0x12, 0x50, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x02, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x01, 0x52, 0x02, 0x65, 0x6e, 0x12, 0x4a, 0x0a, 0x02, 0x6e, 0x6f, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x48, 0x01, 0x52,
	0x02, 0x6e, 0x6f, 0x42, 0x06, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x65,
	0x78, 0x74, 0x22, 0x85, 0x01, 0x0a, 0x0e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x37, 0x0a, 0x0b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0xa7, 0x01, 0x0a, 0x14, 0x44,
	0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x52, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x70, 0x62, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x73, 0x6b, 0x32, 0xa1, 0x08, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0xbc, 0x02, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x3d, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x3d, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xb5, 0x01, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0xae, 0x01, 0x5a, 0x1d, 0x12, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x7b,
	0x6e, 0x75, 0x6d, 0x7d, 0x5a, 0x24, 0x12, 0x22, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x7b, 0x6e,
	0x75, 0x6d, 0x7d, 0x2f, 0x7b, 0x6c, 0x61, 0x6e, 0x67, 0x7d, 0x5a, 0x31, 0x12, 0x2f, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x31, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x7b, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x7d, 0x2f,
	0x7b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x7d, 0x5a, 0x1d, 0x12,
	0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x63, 0x68,
	0x6f, 0x32, 0x2f, 0x7b, 0x6e, 0x6f, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x7d, 0x22, 0x15, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0xaa, 0x01, 0x0a, 0x08, 0x45, 0x63, 0x68, 0x6f, 0x42, 0x6f, 0x64, 0x79,
	0x12, 0x3d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70,
	0x62, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x3d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62,
	0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x20,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x5f, 0x62, 0x6f, 0x64, 0x79,
	0x12, 0xab, 0x01, 0x0a, 0x0a, 0x45, 0x63, 0x68, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x3d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62,
	0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x3d,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x19, 0x2a, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0xbd,
	0x01, 0x0a, 0x09, 0x45, 0x63, 0x68, 0x6f, 0x50, 0x61, 0x74, 0x63, 0x68, 0x12, 0x44, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x44, 0x79,
	0x6e, 0x61, 0x6d, 0x69, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x1a, 0x44, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x70, 0x62, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e,
	0x3a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x32, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x5f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0xb7,
	0x01, 0x0a, 0x10, 0x45, 0x63, 0x68, 0x6f, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x64, 0x12, 0x3d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x3d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x5f, 0x75, 0x6e, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2d, 0x4c,
	0x61, 0x62, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x76, 0x32, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_examples_internal_proto_examplepb_echo_service_proto_rawDescOnce sync.Once
	file_examples_internal_proto_examplepb_echo_service_proto_rawDescData = file_examples_internal_proto_examplepb_echo_service_proto_rawDesc
)

func file_examples_internal_proto_examplepb_echo_service_proto_rawDescGZIP() []byte {
	file_examples_internal_proto_examplepb_echo_service_proto_rawDescOnce.Do(func() {
		file_examples_internal_proto_examplepb_echo_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_examples_internal_proto_examplepb_echo_service_proto_rawDescData)
	})
	return file_examples_internal_proto_examplepb_echo_service_proto_rawDescData
}

var file_examples_internal_proto_examplepb_echo_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_examples_internal_proto_examplepb_echo_service_proto_goTypes = []interface{}{
	(*Embedded)(nil),              // 0: grpc.gateway.examples.internal.proto.examplepb.Embedded
	(*SimpleMessage)(nil),         // 1: grpc.gateway.examples.internal.proto.examplepb.SimpleMessage
	(*DynamicMessage)(nil),        // 2: grpc.gateway.examples.internal.proto.examplepb.DynamicMessage
	(*DynamicMessageUpdate)(nil),  // 3: grpc.gateway.examples.internal.proto.examplepb.DynamicMessageUpdate
	(*structpb.Struct)(nil),       // 4: google.protobuf.Struct
	(*structpb.Value)(nil),        // 5: google.protobuf.Value
	(*fieldmaskpb.FieldMask)(nil), // 6: google.protobuf.FieldMask
}
var file_examples_internal_proto_examplepb_echo_service_proto_depIdxs = []int32{
	0,  // 0: grpc.gateway.examples.internal.proto.examplepb.SimpleMessage.status:type_name -> grpc.gateway.examples.internal.proto.examplepb.Embedded
	0,  // 1: grpc.gateway.examples.internal.proto.examplepb.SimpleMessage.no:type_name -> grpc.gateway.examples.internal.proto.examplepb.Embedded
	4,  // 2: grpc.gateway.examples.internal.proto.examplepb.DynamicMessage.struct_field:type_name -> google.protobuf.Struct
	5,  // 3: grpc.gateway.examples.internal.proto.examplepb.DynamicMessage.value_field:type_name -> google.protobuf.Value
	2,  // 4: grpc.gateway.examples.internal.proto.examplepb.DynamicMessageUpdate.body:type_name -> grpc.gateway.examples.internal.proto.examplepb.DynamicMessage
	6,  // 5: grpc.gateway.examples.internal.proto.examplepb.DynamicMessageUpdate.update_mask:type_name -> google.protobuf.FieldMask
	1,  // 6: grpc.gateway.examples.internal.proto.examplepb.EchoService.Echo:input_type -> grpc.gateway.examples.internal.proto.examplepb.SimpleMessage
	1,  // 7: grpc.gateway.examples.internal.proto.examplepb.EchoService.EchoBody:input_type -> grpc.gateway.examples.internal.proto.examplepb.SimpleMessage
	1,  // 8: grpc.gateway.examples.internal.proto.examplepb.EchoService.EchoDelete:input_type -> grpc.gateway.examples.internal.proto.examplepb.SimpleMessage
	3,  // 9: grpc.gateway.examples.internal.proto.examplepb.EchoService.EchoPatch:input_type -> grpc.gateway.examples.internal.proto.examplepb.DynamicMessageUpdate
	1,  // 10: grpc.gateway.examples.internal.proto.examplepb.EchoService.EchoUnauthorized:input_type -> grpc.gateway.examples.internal.proto.examplepb.SimpleMessage
	1,  // 11: grpc.gateway.examples.internal.proto.examplepb.EchoService.Echo:output_type -> grpc.gateway.examples.internal.proto.examplepb.SimpleMessage
	1,  // 12: grpc.gateway.examples.internal.proto.examplepb.EchoService.EchoBody:output_type -> grpc.gateway.examples.internal.proto.examplepb.SimpleMessage
	1,  // 13: grpc.gateway.examples.internal.proto.examplepb.EchoService.EchoDelete:output_type -> grpc.gateway.examples.internal.proto.examplepb.SimpleMessage
	3,  // 14: grpc.gateway.examples.internal.proto.examplepb.EchoService.EchoPatch:output_type -> grpc.gateway.examples.internal.proto.examplepb.DynamicMessageUpdate
	1,  // 15: grpc.gateway.examples.internal.proto.examplepb.EchoService.EchoUnauthorized:output_type -> grpc.gateway.examples.internal.proto.examplepb.SimpleMessage
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_examples_internal_proto_examplepb_echo_service_proto_init() }
func file_examples_internal_proto_examplepb_echo_service_proto_init() {
	if File_examples_internal_proto_examplepb_echo_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_examples_internal_proto_examplepb_echo_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Embedded); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_examples_internal_proto_examplepb_echo_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_examples_internal_proto_examplepb_echo_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DynamicMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_examples_internal_proto_examplepb_echo_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DynamicMessageUpdate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_examples_internal_proto_examplepb_echo_service_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Embedded_Progress)(nil),
		(*Embedded_Note)(nil),
	}
	file_examples_internal_proto_examplepb_echo_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*SimpleMessage_LineNum)(nil),
		(*SimpleMessage_Lang)(nil),
		(*SimpleMessage_En)(nil),
		(*SimpleMessage_No)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_examples_internal_proto_examplepb_echo_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_examples_internal_proto_examplepb_echo_service_proto_goTypes,
		DependencyIndexes: file_examples_internal_proto_examplepb_echo_service_proto_depIdxs,
		MessageInfos:      file_examples_internal_proto_examplepb_echo_service_proto_msgTypes,
	}.Build()
	File_examples_internal_proto_examplepb_echo_service_proto = out.File
	file_examples_internal_proto_examplepb_echo_service_proto_rawDesc = nil
	file_examples_internal_proto_examplepb_echo_service_proto_goTypes = nil
	file_examples_internal_proto_examplepb_echo_service_proto_depIdxs = nil
}
