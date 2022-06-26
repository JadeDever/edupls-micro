// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/course/service/v1/course.proto

package v1

import (
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

type CreateCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateCourseRequest) Reset() {
	*x = CreateCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_course_service_v1_course_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCourseRequest) ProtoMessage() {}

func (x *CreateCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_course_service_v1_course_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCourseRequest.ProtoReflect.Descriptor instead.
func (*CreateCourseRequest) Descriptor() ([]byte, []int) {
	return file_api_course_service_v1_course_proto_rawDescGZIP(), []int{0}
}

type CreateCourseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateCourseReply) Reset() {
	*x = CreateCourseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_course_service_v1_course_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCourseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCourseReply) ProtoMessage() {}

func (x *CreateCourseReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_course_service_v1_course_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCourseReply.ProtoReflect.Descriptor instead.
func (*CreateCourseReply) Descriptor() ([]byte, []int) {
	return file_api_course_service_v1_course_proto_rawDescGZIP(), []int{1}
}

type UpdateCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCourseRequest) Reset() {
	*x = UpdateCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_course_service_v1_course_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCourseRequest) ProtoMessage() {}

func (x *UpdateCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_course_service_v1_course_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCourseRequest.ProtoReflect.Descriptor instead.
func (*UpdateCourseRequest) Descriptor() ([]byte, []int) {
	return file_api_course_service_v1_course_proto_rawDescGZIP(), []int{2}
}

type UpdateCourseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCourseReply) Reset() {
	*x = UpdateCourseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_course_service_v1_course_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCourseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCourseReply) ProtoMessage() {}

func (x *UpdateCourseReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_course_service_v1_course_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCourseReply.ProtoReflect.Descriptor instead.
func (*UpdateCourseReply) Descriptor() ([]byte, []int) {
	return file_api_course_service_v1_course_proto_rawDescGZIP(), []int{3}
}

type DeleteCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCourseRequest) Reset() {
	*x = DeleteCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_course_service_v1_course_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCourseRequest) ProtoMessage() {}

func (x *DeleteCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_course_service_v1_course_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCourseRequest.ProtoReflect.Descriptor instead.
func (*DeleteCourseRequest) Descriptor() ([]byte, []int) {
	return file_api_course_service_v1_course_proto_rawDescGZIP(), []int{4}
}

type DeleteCourseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCourseReply) Reset() {
	*x = DeleteCourseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_course_service_v1_course_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCourseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCourseReply) ProtoMessage() {}

func (x *DeleteCourseReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_course_service_v1_course_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCourseReply.ProtoReflect.Descriptor instead.
func (*DeleteCourseReply) Descriptor() ([]byte, []int) {
	return file_api_course_service_v1_course_proto_rawDescGZIP(), []int{5}
}

type GetCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCourseRequest) Reset() {
	*x = GetCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_course_service_v1_course_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCourseRequest) ProtoMessage() {}

func (x *GetCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_course_service_v1_course_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCourseRequest.ProtoReflect.Descriptor instead.
func (*GetCourseRequest) Descriptor() ([]byte, []int) {
	return file_api_course_service_v1_course_proto_rawDescGZIP(), []int{6}
}

type GetCourseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCourseReply) Reset() {
	*x = GetCourseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_course_service_v1_course_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCourseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCourseReply) ProtoMessage() {}

func (x *GetCourseReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_course_service_v1_course_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCourseReply.ProtoReflect.Descriptor instead.
func (*GetCourseReply) Descriptor() ([]byte, []int) {
	return file_api_course_service_v1_course_proto_rawDescGZIP(), []int{7}
}

type ListCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCourseRequest) Reset() {
	*x = ListCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_course_service_v1_course_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCourseRequest) ProtoMessage() {}

func (x *ListCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_course_service_v1_course_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCourseRequest.ProtoReflect.Descriptor instead.
func (*ListCourseRequest) Descriptor() ([]byte, []int) {
	return file_api_course_service_v1_course_proto_rawDescGZIP(), []int{8}
}

type ListCourseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCourseReply) Reset() {
	*x = ListCourseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_course_service_v1_course_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCourseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCourseReply) ProtoMessage() {}

func (x *ListCourseReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_course_service_v1_course_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCourseReply.ProtoReflect.Descriptor instead.
func (*ListCourseReply) Descriptor() ([]byte, []int) {
	return file_api_course_service_v1_course_proto_rawDescGZIP(), []int{9}
}

var File_api_course_service_v1_course_proto protoreflect.FileDescriptor

var file_api_course_service_v1_course_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x15, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xf7, 0x03,
	0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x64,
	0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x2a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x64, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5b, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5e, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x3a, 0x0a, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x1f, 0x65, 0x64, 0x75, 0x70, 0x6c, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_course_service_v1_course_proto_rawDescOnce sync.Once
	file_api_course_service_v1_course_proto_rawDescData = file_api_course_service_v1_course_proto_rawDesc
)

func file_api_course_service_v1_course_proto_rawDescGZIP() []byte {
	file_api_course_service_v1_course_proto_rawDescOnce.Do(func() {
		file_api_course_service_v1_course_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_course_service_v1_course_proto_rawDescData)
	})
	return file_api_course_service_v1_course_proto_rawDescData
}

var file_api_course_service_v1_course_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_course_service_v1_course_proto_goTypes = []interface{}{
	(*CreateCourseRequest)(nil), // 0: api.course.service.v1.CreateCourseRequest
	(*CreateCourseReply)(nil),   // 1: api.course.service.v1.CreateCourseReply
	(*UpdateCourseRequest)(nil), // 2: api.course.service.v1.UpdateCourseRequest
	(*UpdateCourseReply)(nil),   // 3: api.course.service.v1.UpdateCourseReply
	(*DeleteCourseRequest)(nil), // 4: api.course.service.v1.DeleteCourseRequest
	(*DeleteCourseReply)(nil),   // 5: api.course.service.v1.DeleteCourseReply
	(*GetCourseRequest)(nil),    // 6: api.course.service.v1.GetCourseRequest
	(*GetCourseReply)(nil),      // 7: api.course.service.v1.GetCourseReply
	(*ListCourseRequest)(nil),   // 8: api.course.service.v1.ListCourseRequest
	(*ListCourseReply)(nil),     // 9: api.course.service.v1.ListCourseReply
}
var file_api_course_service_v1_course_proto_depIdxs = []int32{
	0, // 0: api.course.service.v1.Course.CreateCourse:input_type -> api.course.service.v1.CreateCourseRequest
	2, // 1: api.course.service.v1.Course.UpdateCourse:input_type -> api.course.service.v1.UpdateCourseRequest
	4, // 2: api.course.service.v1.Course.DeleteCourse:input_type -> api.course.service.v1.DeleteCourseRequest
	6, // 3: api.course.service.v1.Course.GetCourse:input_type -> api.course.service.v1.GetCourseRequest
	8, // 4: api.course.service.v1.Course.ListCourse:input_type -> api.course.service.v1.ListCourseRequest
	1, // 5: api.course.service.v1.Course.CreateCourse:output_type -> api.course.service.v1.CreateCourseReply
	3, // 6: api.course.service.v1.Course.UpdateCourse:output_type -> api.course.service.v1.UpdateCourseReply
	5, // 7: api.course.service.v1.Course.DeleteCourse:output_type -> api.course.service.v1.DeleteCourseReply
	7, // 8: api.course.service.v1.Course.GetCourse:output_type -> api.course.service.v1.GetCourseReply
	9, // 9: api.course.service.v1.Course.ListCourse:output_type -> api.course.service.v1.ListCourseReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_course_service_v1_course_proto_init() }
func file_api_course_service_v1_course_proto_init() {
	if File_api_course_service_v1_course_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_course_service_v1_course_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCourseRequest); i {
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
		file_api_course_service_v1_course_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCourseReply); i {
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
		file_api_course_service_v1_course_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCourseRequest); i {
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
		file_api_course_service_v1_course_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCourseReply); i {
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
		file_api_course_service_v1_course_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCourseRequest); i {
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
		file_api_course_service_v1_course_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCourseReply); i {
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
		file_api_course_service_v1_course_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCourseRequest); i {
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
		file_api_course_service_v1_course_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCourseReply); i {
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
		file_api_course_service_v1_course_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCourseRequest); i {
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
		file_api_course_service_v1_course_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCourseReply); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_course_service_v1_course_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_course_service_v1_course_proto_goTypes,
		DependencyIndexes: file_api_course_service_v1_course_proto_depIdxs,
		MessageInfos:      file_api_course_service_v1_course_proto_msgTypes,
	}.Build()
	File_api_course_service_v1_course_proto = out.File
	file_api_course_service_v1_course_proto_rawDesc = nil
	file_api_course_service_v1_course_proto_goTypes = nil
	file_api_course_service_v1_course_proto_depIdxs = nil
}