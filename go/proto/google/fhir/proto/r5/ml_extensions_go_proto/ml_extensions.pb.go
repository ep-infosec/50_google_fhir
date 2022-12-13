//    Copyright 2019 Google Inc.
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        https://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: proto/google/fhir/proto/r5/ml_extensions.proto

package ml_extensions_go_proto

import (
	_ "github.com/google/fhir/go/proto/google/fhir/proto/annotations_go_proto"
	datatypes_go_proto "github.com/google/fhir/go/proto/google/fhir/proto/r5/core/datatypes_go_proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// EventLabels define labels used for TensorFlow model training and evaluation.
// See https://g.co/fhir/StructureDefinition/eventLabel
type EventLabel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique id for inter-element referencing
	Id *datatypes_go_proto.String `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []*datatypes_go_proto.Extension `protobuf:"bytes,2,rep,name=extension,proto3" json:"extension,omitempty"`
	// The patient associated with this label
	Patient *datatypes_go_proto.Reference `protobuf:"bytes,4,opt,name=patient,proto3" json:"patient,omitempty"`
	// The type of label, part of the prediction task definition
	Type *datatypes_go_proto.Coding `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	// Time associated with the label event, if available
	EventTime *datatypes_go_proto.DateTime `protobuf:"bytes,6,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	// Resource that owns this label
	Source *datatypes_go_proto.Reference `protobuf:"bytes,7,opt,name=source,proto3" json:"source,omitempty"`
	Label  []*EventLabel_Label           `protobuf:"bytes,8,rep,name=label,proto3" json:"label,omitempty"`
}

func (x *EventLabel) Reset() {
	*x = EventLabel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_google_fhir_proto_r5_ml_extensions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventLabel) ProtoMessage() {}

func (x *EventLabel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_google_fhir_proto_r5_ml_extensions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventLabel.ProtoReflect.Descriptor instead.
func (*EventLabel) Descriptor() ([]byte, []int) {
	return file_proto_google_fhir_proto_r5_ml_extensions_proto_rawDescGZIP(), []int{0}
}

func (x *EventLabel) GetId() *datatypes_go_proto.String {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *EventLabel) GetExtension() []*datatypes_go_proto.Extension {
	if x != nil {
		return x.Extension
	}
	return nil
}

func (x *EventLabel) GetPatient() *datatypes_go_proto.Reference {
	if x != nil {
		return x.Patient
	}
	return nil
}

func (x *EventLabel) GetType() *datatypes_go_proto.Coding {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *EventLabel) GetEventTime() *datatypes_go_proto.DateTime {
	if x != nil {
		return x.EventTime
	}
	return nil
}

func (x *EventLabel) GetSource() *datatypes_go_proto.Reference {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *EventLabel) GetLabel() []*EventLabel_Label {
	if x != nil {
		return x.Label
	}
	return nil
}

// EventTriggers specify cutoff times for generated TensorFlow examples.
// See https://g.co/fhir/StructureDefinition/eventTrigger
type EventTrigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique id for inter-element referencing
	Id *datatypes_go_proto.String `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The type of trigger, part of the prediction task definition.
	Type *datatypes_go_proto.Coding `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	// Prediction event time, more recent data will be elided.
	EventTime *datatypes_go_proto.DateTime `protobuf:"bytes,5,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	// Resource that owns this trigger.
	Source *datatypes_go_proto.Reference `protobuf:"bytes,6,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *EventTrigger) Reset() {
	*x = EventTrigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_google_fhir_proto_r5_ml_extensions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventTrigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventTrigger) ProtoMessage() {}

func (x *EventTrigger) ProtoReflect() protoreflect.Message {
	mi := &file_proto_google_fhir_proto_r5_ml_extensions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventTrigger.ProtoReflect.Descriptor instead.
func (*EventTrigger) Descriptor() ([]byte, []int) {
	return file_proto_google_fhir_proto_r5_ml_extensions_proto_rawDescGZIP(), []int{1}
}

func (x *EventTrigger) GetId() *datatypes_go_proto.String {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *EventTrigger) GetType() *datatypes_go_proto.Coding {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *EventTrigger) GetEventTime() *datatypes_go_proto.DateTime {
	if x != nil {
		return x.EventTime
	}
	return nil
}

func (x *EventTrigger) GetSource() *datatypes_go_proto.Reference {
	if x != nil {
		return x.Source
	}
	return nil
}

// List of labels associated with this event
type EventLabel_Label struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique id for inter-element referencing
	Id *datatypes_go_proto.String `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Class name in multi-class prediction tasks, e.g. code "780.60" for icd9
	ClassName  *datatypes_go_proto.Coding    `protobuf:"bytes,4,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`
	ClassValue *EventLabel_Label_ClassValueX `protobuf:"bytes,5,opt,name=class_value,json=classValue,proto3" json:"class_value,omitempty"`
}

func (x *EventLabel_Label) Reset() {
	*x = EventLabel_Label{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_google_fhir_proto_r5_ml_extensions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventLabel_Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventLabel_Label) ProtoMessage() {}

func (x *EventLabel_Label) ProtoReflect() protoreflect.Message {
	mi := &file_proto_google_fhir_proto_r5_ml_extensions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventLabel_Label.ProtoReflect.Descriptor instead.
func (*EventLabel_Label) Descriptor() ([]byte, []int) {
	return file_proto_google_fhir_proto_r5_ml_extensions_proto_rawDescGZIP(), []int{0, 0}
}

func (x *EventLabel_Label) GetId() *datatypes_go_proto.String {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *EventLabel_Label) GetClassName() *datatypes_go_proto.Coding {
	if x != nil {
		return x.ClassName
	}
	return nil
}

func (x *EventLabel_Label) GetClassValue() *EventLabel_Label_ClassValueX {
	if x != nil {
		return x.ClassValue
	}
	return nil
}

// The value of the label
type EventLabel_Label_ClassValueX struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Choice:
	//	*EventLabel_Label_ClassValueX_Boolean
	//	*EventLabel_Label_ClassValueX_Decimal
	//	*EventLabel_Label_ClassValueX_Integer
	//	*EventLabel_Label_ClassValueX_StringValue
	//	*EventLabel_Label_ClassValueX_DateTime
	Choice isEventLabel_Label_ClassValueX_Choice `protobuf_oneof:"choice"`
}

func (x *EventLabel_Label_ClassValueX) Reset() {
	*x = EventLabel_Label_ClassValueX{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_google_fhir_proto_r5_ml_extensions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventLabel_Label_ClassValueX) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventLabel_Label_ClassValueX) ProtoMessage() {}

func (x *EventLabel_Label_ClassValueX) ProtoReflect() protoreflect.Message {
	mi := &file_proto_google_fhir_proto_r5_ml_extensions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventLabel_Label_ClassValueX.ProtoReflect.Descriptor instead.
func (*EventLabel_Label_ClassValueX) Descriptor() ([]byte, []int) {
	return file_proto_google_fhir_proto_r5_ml_extensions_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (m *EventLabel_Label_ClassValueX) GetChoice() isEventLabel_Label_ClassValueX_Choice {
	if m != nil {
		return m.Choice
	}
	return nil
}

func (x *EventLabel_Label_ClassValueX) GetBoolean() *datatypes_go_proto.Boolean {
	if x, ok := x.GetChoice().(*EventLabel_Label_ClassValueX_Boolean); ok {
		return x.Boolean
	}
	return nil
}

func (x *EventLabel_Label_ClassValueX) GetDecimal() *datatypes_go_proto.Decimal {
	if x, ok := x.GetChoice().(*EventLabel_Label_ClassValueX_Decimal); ok {
		return x.Decimal
	}
	return nil
}

func (x *EventLabel_Label_ClassValueX) GetInteger() *datatypes_go_proto.Integer {
	if x, ok := x.GetChoice().(*EventLabel_Label_ClassValueX_Integer); ok {
		return x.Integer
	}
	return nil
}

func (x *EventLabel_Label_ClassValueX) GetStringValue() *datatypes_go_proto.String {
	if x, ok := x.GetChoice().(*EventLabel_Label_ClassValueX_StringValue); ok {
		return x.StringValue
	}
	return nil
}

func (x *EventLabel_Label_ClassValueX) GetDateTime() *datatypes_go_proto.DateTime {
	if x, ok := x.GetChoice().(*EventLabel_Label_ClassValueX_DateTime); ok {
		return x.DateTime
	}
	return nil
}

type isEventLabel_Label_ClassValueX_Choice interface {
	isEventLabel_Label_ClassValueX_Choice()
}

type EventLabel_Label_ClassValueX_Boolean struct {
	Boolean *datatypes_go_proto.Boolean `protobuf:"bytes,1,opt,name=boolean,proto3,oneof"`
}

type EventLabel_Label_ClassValueX_Decimal struct {
	Decimal *datatypes_go_proto.Decimal `protobuf:"bytes,2,opt,name=decimal,proto3,oneof"`
}

type EventLabel_Label_ClassValueX_Integer struct {
	Integer *datatypes_go_proto.Integer `protobuf:"bytes,3,opt,name=integer,proto3,oneof"`
}

type EventLabel_Label_ClassValueX_StringValue struct {
	StringValue *datatypes_go_proto.String `protobuf:"bytes,4,opt,name=string_value,json=string,proto3,oneof"`
}

type EventLabel_Label_ClassValueX_DateTime struct {
	DateTime *datatypes_go_proto.DateTime `protobuf:"bytes,5,opt,name=date_time,json=dateTime,proto3,oneof"`
}

func (*EventLabel_Label_ClassValueX_Boolean) isEventLabel_Label_ClassValueX_Choice() {}

func (*EventLabel_Label_ClassValueX_Decimal) isEventLabel_Label_ClassValueX_Choice() {}

func (*EventLabel_Label_ClassValueX_Integer) isEventLabel_Label_ClassValueX_Choice() {}

func (*EventLabel_Label_ClassValueX_StringValue) isEventLabel_Label_ClassValueX_Choice() {}

func (*EventLabel_Label_ClassValueX_DateTime) isEventLabel_Label_ClassValueX_Choice() {}

var File_proto_google_fhir_proto_r5_ml_extensions_proto protoreflect.FileDescriptor

var file_proto_google_fhir_proto_r5_ml_extensions_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66,
	0x68, 0x69, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x35, 0x2f, 0x6d, 0x6c, 0x5f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35,
	0x2e, 0x6d, 0x6c, 0x1a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x68, 0x69,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x35, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc4, 0x08, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x2b,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3c, 0x0a, 0x09, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x07, 0x70, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x06, 0xf0, 0xd0, 0x87, 0xeb,
	0x04, 0x01, 0x52, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x43, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x06, 0xf0, 0xd0, 0x87, 0xeb, 0x04, 0x01, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x44,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72,
	0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e, 0x6d, 0x6c, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x1a, 0xa9, 0x04, 0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12,
	0x2b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x02, 0x69, 0x64, 0x12, 0x42, 0x0a, 0x0a,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72,
	0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x06, 0xf0,
	0xd0, 0x87, 0xeb, 0x04, 0x01, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x58, 0x0a, 0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66,
	0x68, 0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e, 0x6d, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x58, 0x42, 0x06, 0xf0, 0xd0, 0x87, 0xeb, 0x04, 0x01, 0x52, 0x0a,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0xc8, 0x02, 0x0a, 0x0b, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x58, 0x12, 0x38, 0x0a, 0x07, 0x62, 0x6f,
	0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x48, 0x00, 0x52, 0x07, 0x62, 0x6f, 0x6f,
	0x6c, 0x65, 0x61, 0x6e, 0x12, 0x38, 0x0a, 0x07, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66,
	0x68, 0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x69,
	0x6d, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x07, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x12, 0x38,
	0x0a, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x48, 0x00, 0x52,
	0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x06, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x3c, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x44,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x48, 0x00, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x3a, 0x06, 0xa0, 0x83, 0x83, 0xe8, 0x06, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x63,
	0x68, 0x6f, 0x69, 0x63, 0x65, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x03, 0x10,
	0x04, 0x3a, 0x73, 0xc0, 0x9f, 0xe3, 0xb6, 0x05, 0x02, 0x9a, 0xb5, 0x8e, 0x93, 0x06, 0x31, 0x68,
	0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x68, 0x6c, 0x37, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x68,
	0x69, 0x72, 0x2f, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0xb2, 0xfe, 0xe4, 0x97, 0x06, 0x30, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x2e,
	0x63, 0x6f, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0xe9, 0x02, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69,
	0x72, 0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72,
	0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x42,
	0x06, 0xf0, 0xd0, 0x87, 0xeb, 0x04, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x44, 0x0a,
	0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e,
	0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x42, 0x06, 0xf0, 0xd0, 0x87, 0xeb, 0x04, 0x01, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x68, 0x69,
	0x72, 0x2e, 0x72, 0x35, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a, 0x75, 0xc0, 0x9f, 0xe3,
	0xb6, 0x05, 0x02, 0x9a, 0xb5, 0x8e, 0x93, 0x06, 0x31, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f,
	0x68, 0x6c, 0x37, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0xb2, 0xfe, 0xe4, 0x97, 0x06, 0x32,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x2e, 0x63, 0x6f, 0x2f, 0x66, 0x68, 0x69,
	0x72, 0x2f, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x42, 0x6c, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x66, 0x68, 0x69, 0x72, 0x2e, 0x72, 0x35, 0x2e, 0x6d, 0x6c, 0x50, 0x01, 0x5a, 0x4b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x68, 0x69, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x72, 0x35, 0x2f, 0x6d, 0x6c, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x98, 0xc6, 0xb0, 0xb5, 0x07, 0x05,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_google_fhir_proto_r5_ml_extensions_proto_rawDescOnce sync.Once
	file_proto_google_fhir_proto_r5_ml_extensions_proto_rawDescData = file_proto_google_fhir_proto_r5_ml_extensions_proto_rawDesc
)

func file_proto_google_fhir_proto_r5_ml_extensions_proto_rawDescGZIP() []byte {
	file_proto_google_fhir_proto_r5_ml_extensions_proto_rawDescOnce.Do(func() {
		file_proto_google_fhir_proto_r5_ml_extensions_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_google_fhir_proto_r5_ml_extensions_proto_rawDescData)
	})
	return file_proto_google_fhir_proto_r5_ml_extensions_proto_rawDescData
}

var file_proto_google_fhir_proto_r5_ml_extensions_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_google_fhir_proto_r5_ml_extensions_proto_goTypes = []interface{}{
	(*EventLabel)(nil),                   // 0: google.fhir.r5.ml.EventLabel
	(*EventTrigger)(nil),                 // 1: google.fhir.r5.ml.EventTrigger
	(*EventLabel_Label)(nil),             // 2: google.fhir.r5.ml.EventLabel.Label
	(*EventLabel_Label_ClassValueX)(nil), // 3: google.fhir.r5.ml.EventLabel.Label.ClassValueX
	(*datatypes_go_proto.String)(nil),    // 4: google.fhir.r5.core.String
	(*datatypes_go_proto.Extension)(nil), // 5: google.fhir.r5.core.Extension
	(*datatypes_go_proto.Reference)(nil), // 6: google.fhir.r5.core.Reference
	(*datatypes_go_proto.Coding)(nil),    // 7: google.fhir.r5.core.Coding
	(*datatypes_go_proto.DateTime)(nil),  // 8: google.fhir.r5.core.DateTime
	(*datatypes_go_proto.Boolean)(nil),   // 9: google.fhir.r5.core.Boolean
	(*datatypes_go_proto.Decimal)(nil),   // 10: google.fhir.r5.core.Decimal
	(*datatypes_go_proto.Integer)(nil),   // 11: google.fhir.r5.core.Integer
}
var file_proto_google_fhir_proto_r5_ml_extensions_proto_depIdxs = []int32{
	4,  // 0: google.fhir.r5.ml.EventLabel.id:type_name -> google.fhir.r5.core.String
	5,  // 1: google.fhir.r5.ml.EventLabel.extension:type_name -> google.fhir.r5.core.Extension
	6,  // 2: google.fhir.r5.ml.EventLabel.patient:type_name -> google.fhir.r5.core.Reference
	7,  // 3: google.fhir.r5.ml.EventLabel.type:type_name -> google.fhir.r5.core.Coding
	8,  // 4: google.fhir.r5.ml.EventLabel.event_time:type_name -> google.fhir.r5.core.DateTime
	6,  // 5: google.fhir.r5.ml.EventLabel.source:type_name -> google.fhir.r5.core.Reference
	2,  // 6: google.fhir.r5.ml.EventLabel.label:type_name -> google.fhir.r5.ml.EventLabel.Label
	4,  // 7: google.fhir.r5.ml.EventTrigger.id:type_name -> google.fhir.r5.core.String
	7,  // 8: google.fhir.r5.ml.EventTrigger.type:type_name -> google.fhir.r5.core.Coding
	8,  // 9: google.fhir.r5.ml.EventTrigger.event_time:type_name -> google.fhir.r5.core.DateTime
	6,  // 10: google.fhir.r5.ml.EventTrigger.source:type_name -> google.fhir.r5.core.Reference
	4,  // 11: google.fhir.r5.ml.EventLabel.Label.id:type_name -> google.fhir.r5.core.String
	7,  // 12: google.fhir.r5.ml.EventLabel.Label.class_name:type_name -> google.fhir.r5.core.Coding
	3,  // 13: google.fhir.r5.ml.EventLabel.Label.class_value:type_name -> google.fhir.r5.ml.EventLabel.Label.ClassValueX
	9,  // 14: google.fhir.r5.ml.EventLabel.Label.ClassValueX.boolean:type_name -> google.fhir.r5.core.Boolean
	10, // 15: google.fhir.r5.ml.EventLabel.Label.ClassValueX.decimal:type_name -> google.fhir.r5.core.Decimal
	11, // 16: google.fhir.r5.ml.EventLabel.Label.ClassValueX.integer:type_name -> google.fhir.r5.core.Integer
	4,  // 17: google.fhir.r5.ml.EventLabel.Label.ClassValueX.string_value:type_name -> google.fhir.r5.core.String
	8,  // 18: google.fhir.r5.ml.EventLabel.Label.ClassValueX.date_time:type_name -> google.fhir.r5.core.DateTime
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_proto_google_fhir_proto_r5_ml_extensions_proto_init() }
func file_proto_google_fhir_proto_r5_ml_extensions_proto_init() {
	if File_proto_google_fhir_proto_r5_ml_extensions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_google_fhir_proto_r5_ml_extensions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventLabel); i {
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
		file_proto_google_fhir_proto_r5_ml_extensions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventTrigger); i {
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
		file_proto_google_fhir_proto_r5_ml_extensions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventLabel_Label); i {
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
		file_proto_google_fhir_proto_r5_ml_extensions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventLabel_Label_ClassValueX); i {
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
	file_proto_google_fhir_proto_r5_ml_extensions_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*EventLabel_Label_ClassValueX_Boolean)(nil),
		(*EventLabel_Label_ClassValueX_Decimal)(nil),
		(*EventLabel_Label_ClassValueX_Integer)(nil),
		(*EventLabel_Label_ClassValueX_StringValue)(nil),
		(*EventLabel_Label_ClassValueX_DateTime)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_google_fhir_proto_r5_ml_extensions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_google_fhir_proto_r5_ml_extensions_proto_goTypes,
		DependencyIndexes: file_proto_google_fhir_proto_r5_ml_extensions_proto_depIdxs,
		MessageInfos:      file_proto_google_fhir_proto_r5_ml_extensions_proto_msgTypes,
	}.Build()
	File_proto_google_fhir_proto_r5_ml_extensions_proto = out.File
	file_proto_google_fhir_proto_r5_ml_extensions_proto_rawDesc = nil
	file_proto_google_fhir_proto_r5_ml_extensions_proto_goTypes = nil
	file_proto_google_fhir_proto_r5_ml_extensions_proto_depIdxs = nil
}
