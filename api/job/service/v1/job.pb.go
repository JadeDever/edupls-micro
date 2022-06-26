// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/job/service/v1/job.proto

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

type CreateJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateJobRequest) Reset() {
	*x = CreateJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_job_service_v1_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobRequest) ProtoMessage() {}

func (x *CreateJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_job_service_v1_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobRequest.ProtoReflect.Descriptor instead.
func (*CreateJobRequest) Descriptor() ([]byte, []int) {
	return file_api_job_service_v1_job_proto_rawDescGZIP(), []int{0}
}

type CreateJobReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateJobReply) Reset() {
	*x = CreateJobReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_job_service_v1_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobReply) ProtoMessage() {}

func (x *CreateJobReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_job_service_v1_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobReply.ProtoReflect.Descriptor instead.
func (*CreateJobReply) Descriptor() ([]byte, []int) {
	return file_api_job_service_v1_job_proto_rawDescGZIP(), []int{1}
}

type UpdateJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateJobRequest) Reset() {
	*x = UpdateJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_job_service_v1_job_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobRequest) ProtoMessage() {}

func (x *UpdateJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_job_service_v1_job_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobRequest.ProtoReflect.Descriptor instead.
func (*UpdateJobRequest) Descriptor() ([]byte, []int) {
	return file_api_job_service_v1_job_proto_rawDescGZIP(), []int{2}
}

type UpdateJobReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateJobReply) Reset() {
	*x = UpdateJobReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_job_service_v1_job_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJobReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobReply) ProtoMessage() {}

func (x *UpdateJobReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_job_service_v1_job_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobReply.ProtoReflect.Descriptor instead.
func (*UpdateJobReply) Descriptor() ([]byte, []int) {
	return file_api_job_service_v1_job_proto_rawDescGZIP(), []int{3}
}

type DeleteJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteJobRequest) Reset() {
	*x = DeleteJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_job_service_v1_job_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJobRequest) ProtoMessage() {}

func (x *DeleteJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_job_service_v1_job_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJobRequest.ProtoReflect.Descriptor instead.
func (*DeleteJobRequest) Descriptor() ([]byte, []int) {
	return file_api_job_service_v1_job_proto_rawDescGZIP(), []int{4}
}

type DeleteJobReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteJobReply) Reset() {
	*x = DeleteJobReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_job_service_v1_job_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJobReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJobReply) ProtoMessage() {}

func (x *DeleteJobReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_job_service_v1_job_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJobReply.ProtoReflect.Descriptor instead.
func (*DeleteJobReply) Descriptor() ([]byte, []int) {
	return file_api_job_service_v1_job_proto_rawDescGZIP(), []int{5}
}

type GetJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetJobRequest) Reset() {
	*x = GetJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_job_service_v1_job_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobRequest) ProtoMessage() {}

func (x *GetJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_job_service_v1_job_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobRequest.ProtoReflect.Descriptor instead.
func (*GetJobRequest) Descriptor() ([]byte, []int) {
	return file_api_job_service_v1_job_proto_rawDescGZIP(), []int{6}
}

type GetJobReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetJobReply) Reset() {
	*x = GetJobReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_job_service_v1_job_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJobReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobReply) ProtoMessage() {}

func (x *GetJobReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_job_service_v1_job_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobReply.ProtoReflect.Descriptor instead.
func (*GetJobReply) Descriptor() ([]byte, []int) {
	return file_api_job_service_v1_job_proto_rawDescGZIP(), []int{7}
}

type ListJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListJobRequest) Reset() {
	*x = ListJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_job_service_v1_job_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobRequest) ProtoMessage() {}

func (x *ListJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_job_service_v1_job_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobRequest.ProtoReflect.Descriptor instead.
func (*ListJobRequest) Descriptor() ([]byte, []int) {
	return file_api_job_service_v1_job_proto_rawDescGZIP(), []int{8}
}

type ListJobReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListJobReply) Reset() {
	*x = ListJobReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_job_service_v1_job_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListJobReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobReply) ProtoMessage() {}

func (x *ListJobReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_job_service_v1_job_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobReply.ProtoReflect.Descriptor instead.
func (*ListJobReply) Descriptor() ([]byte, []int) {
	return file_api_job_service_v1_job_proto_rawDescGZIP(), []int{9}
}

var File_api_job_service_v1_job_proto protoreflect.FileDescriptor

var file_api_job_service_v1_job_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x61, 0x70, 0x69, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12,
	0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xa9, 0x03, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x55,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x24, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x55, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a,
	0x6f, 0x62, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6a,
	0x6f, 0x62, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x55, 0x0a, 0x09,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6a, 0x6f, 0x62, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x4c, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x12, 0x21, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x4f, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x12, 0x22, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x34, 0x0a, 0x12, 0x61, 0x70, 0x69, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1c, 0x65, 0x64, 0x75, 0x70,
	0x6c, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_job_service_v1_job_proto_rawDescOnce sync.Once
	file_api_job_service_v1_job_proto_rawDescData = file_api_job_service_v1_job_proto_rawDesc
)

func file_api_job_service_v1_job_proto_rawDescGZIP() []byte {
	file_api_job_service_v1_job_proto_rawDescOnce.Do(func() {
		file_api_job_service_v1_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_job_service_v1_job_proto_rawDescData)
	})
	return file_api_job_service_v1_job_proto_rawDescData
}

var file_api_job_service_v1_job_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_job_service_v1_job_proto_goTypes = []interface{}{
	(*CreateJobRequest)(nil), // 0: api.job.service.v1.CreateJobRequest
	(*CreateJobReply)(nil),   // 1: api.job.service.v1.CreateJobReply
	(*UpdateJobRequest)(nil), // 2: api.job.service.v1.UpdateJobRequest
	(*UpdateJobReply)(nil),   // 3: api.job.service.v1.UpdateJobReply
	(*DeleteJobRequest)(nil), // 4: api.job.service.v1.DeleteJobRequest
	(*DeleteJobReply)(nil),   // 5: api.job.service.v1.DeleteJobReply
	(*GetJobRequest)(nil),    // 6: api.job.service.v1.GetJobRequest
	(*GetJobReply)(nil),      // 7: api.job.service.v1.GetJobReply
	(*ListJobRequest)(nil),   // 8: api.job.service.v1.ListJobRequest
	(*ListJobReply)(nil),     // 9: api.job.service.v1.ListJobReply
}
var file_api_job_service_v1_job_proto_depIdxs = []int32{
	0, // 0: api.job.service.v1.Job.CreateJob:input_type -> api.job.service.v1.CreateJobRequest
	2, // 1: api.job.service.v1.Job.UpdateJob:input_type -> api.job.service.v1.UpdateJobRequest
	4, // 2: api.job.service.v1.Job.DeleteJob:input_type -> api.job.service.v1.DeleteJobRequest
	6, // 3: api.job.service.v1.Job.GetJob:input_type -> api.job.service.v1.GetJobRequest
	8, // 4: api.job.service.v1.Job.ListJob:input_type -> api.job.service.v1.ListJobRequest
	1, // 5: api.job.service.v1.Job.CreateJob:output_type -> api.job.service.v1.CreateJobReply
	3, // 6: api.job.service.v1.Job.UpdateJob:output_type -> api.job.service.v1.UpdateJobReply
	5, // 7: api.job.service.v1.Job.DeleteJob:output_type -> api.job.service.v1.DeleteJobReply
	7, // 8: api.job.service.v1.Job.GetJob:output_type -> api.job.service.v1.GetJobReply
	9, // 9: api.job.service.v1.Job.ListJob:output_type -> api.job.service.v1.ListJobReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_job_service_v1_job_proto_init() }
func file_api_job_service_v1_job_proto_init() {
	if File_api_job_service_v1_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_job_service_v1_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJobRequest); i {
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
		file_api_job_service_v1_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJobReply); i {
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
		file_api_job_service_v1_job_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateJobRequest); i {
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
		file_api_job_service_v1_job_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateJobReply); i {
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
		file_api_job_service_v1_job_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteJobRequest); i {
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
		file_api_job_service_v1_job_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteJobReply); i {
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
		file_api_job_service_v1_job_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJobRequest); i {
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
		file_api_job_service_v1_job_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJobReply); i {
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
		file_api_job_service_v1_job_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListJobRequest); i {
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
		file_api_job_service_v1_job_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListJobReply); i {
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
			RawDescriptor: file_api_job_service_v1_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_job_service_v1_job_proto_goTypes,
		DependencyIndexes: file_api_job_service_v1_job_proto_depIdxs,
		MessageInfos:      file_api_job_service_v1_job_proto_msgTypes,
	}.Build()
	File_api_job_service_v1_job_proto = out.File
	file_api_job_service_v1_job_proto_rawDesc = nil
	file_api_job_service_v1_job_proto_goTypes = nil
	file_api_job_service_v1_job_proto_depIdxs = nil
}
