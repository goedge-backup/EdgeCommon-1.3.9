// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: service_ns_question_option.proto

package pb

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

// 创建选项
type CreateNSQuestionOptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ValuesJSON []byte `protobuf:"bytes,2,opt,name=valuesJSON,proto3" json:"valuesJSON,omitempty"`
}

func (x *CreateNSQuestionOptionRequest) Reset() {
	*x = CreateNSQuestionOptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_question_option_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNSQuestionOptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNSQuestionOptionRequest) ProtoMessage() {}

func (x *CreateNSQuestionOptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_question_option_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNSQuestionOptionRequest.ProtoReflect.Descriptor instead.
func (*CreateNSQuestionOptionRequest) Descriptor() ([]byte, []int) {
	return file_service_ns_question_option_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNSQuestionOptionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateNSQuestionOptionRequest) GetValuesJSON() []byte {
	if x != nil {
		return x.ValuesJSON
	}
	return nil
}

type CreateNSQuestionOptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NsQuestionOptionId int64 `protobuf:"varint,1,opt,name=nsQuestionOptionId,proto3" json:"nsQuestionOptionId,omitempty"`
}

func (x *CreateNSQuestionOptionResponse) Reset() {
	*x = CreateNSQuestionOptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_question_option_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNSQuestionOptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNSQuestionOptionResponse) ProtoMessage() {}

func (x *CreateNSQuestionOptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_question_option_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNSQuestionOptionResponse.ProtoReflect.Descriptor instead.
func (*CreateNSQuestionOptionResponse) Descriptor() ([]byte, []int) {
	return file_service_ns_question_option_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNSQuestionOptionResponse) GetNsQuestionOptionId() int64 {
	if x != nil {
		return x.NsQuestionOptionId
	}
	return 0
}

// 读取选项
type FindNSQuestionOptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NsQuestionOptionId int64 `protobuf:"varint,1,opt,name=nsQuestionOptionId,proto3" json:"nsQuestionOptionId,omitempty"`
}

func (x *FindNSQuestionOptionRequest) Reset() {
	*x = FindNSQuestionOptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_question_option_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindNSQuestionOptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindNSQuestionOptionRequest) ProtoMessage() {}

func (x *FindNSQuestionOptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_question_option_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindNSQuestionOptionRequest.ProtoReflect.Descriptor instead.
func (*FindNSQuestionOptionRequest) Descriptor() ([]byte, []int) {
	return file_service_ns_question_option_proto_rawDescGZIP(), []int{2}
}

func (x *FindNSQuestionOptionRequest) GetNsQuestionOptionId() int64 {
	if x != nil {
		return x.NsQuestionOptionId
	}
	return 0
}

type FindNSQuestionOptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NsQuestionOption *NSQuestionOption `protobuf:"bytes,1,opt,name=nsQuestionOption,proto3" json:"nsQuestionOption,omitempty"`
}

func (x *FindNSQuestionOptionResponse) Reset() {
	*x = FindNSQuestionOptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_question_option_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindNSQuestionOptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindNSQuestionOptionResponse) ProtoMessage() {}

func (x *FindNSQuestionOptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_question_option_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindNSQuestionOptionResponse.ProtoReflect.Descriptor instead.
func (*FindNSQuestionOptionResponse) Descriptor() ([]byte, []int) {
	return file_service_ns_question_option_proto_rawDescGZIP(), []int{3}
}

func (x *FindNSQuestionOptionResponse) GetNsQuestionOption() *NSQuestionOption {
	if x != nil {
		return x.NsQuestionOption
	}
	return nil
}

// 删除选项
type DeleteNSQuestionOptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NsQuestionOptionId int64 `protobuf:"varint,1,opt,name=nsQuestionOptionId,proto3" json:"nsQuestionOptionId,omitempty"`
}

func (x *DeleteNSQuestionOptionRequest) Reset() {
	*x = DeleteNSQuestionOptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_question_option_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNSQuestionOptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNSQuestionOptionRequest) ProtoMessage() {}

func (x *DeleteNSQuestionOptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_question_option_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNSQuestionOptionRequest.ProtoReflect.Descriptor instead.
func (*DeleteNSQuestionOptionRequest) Descriptor() ([]byte, []int) {
	return file_service_ns_question_option_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteNSQuestionOptionRequest) GetNsQuestionOptionId() int64 {
	if x != nil {
		return x.NsQuestionOptionId
	}
	return 0
}

var File_service_ns_question_option_proto protoreflect.FileDescriptor

var file_service_ns_question_option_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x73, 0x5f, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72,
	0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x25, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x6e, 0x73, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x53, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x22, 0x50, 0x0a,
	0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x53, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2e, 0x0a, 0x12, 0x6e, 0x73, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6e, 0x73, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x4d, 0x0a, 0x1b, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x53, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e,
	0x0a, 0x12, 0x6e, 0x73, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6e, 0x73, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x60,
	0x0a, 0x1c, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x53, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40,
	0x0a, 0x10, 0x6e, 0x73, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x53,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10,
	0x6e, 0x73, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x4f, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x53, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2e, 0x0a, 0x12, 0x6e, 0x73, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6e,
	0x73, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x32, 0xa2, 0x02, 0x0a, 0x17, 0x4e, 0x53, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a,
	0x16, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x53, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4e, 0x53, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x53, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59,
	0x0a, 0x14, 0x66, 0x69, 0x6e, 0x64, 0x4e, 0x53, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x4e, 0x53, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x4e, 0x53, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x16, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4e, 0x53, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e,
	0x53, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_ns_question_option_proto_rawDescOnce sync.Once
	file_service_ns_question_option_proto_rawDescData = file_service_ns_question_option_proto_rawDesc
)

func file_service_ns_question_option_proto_rawDescGZIP() []byte {
	file_service_ns_question_option_proto_rawDescOnce.Do(func() {
		file_service_ns_question_option_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_ns_question_option_proto_rawDescData)
	})
	return file_service_ns_question_option_proto_rawDescData
}

var file_service_ns_question_option_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_service_ns_question_option_proto_goTypes = []interface{}{
	(*CreateNSQuestionOptionRequest)(nil),  // 0: pb.CreateNSQuestionOptionRequest
	(*CreateNSQuestionOptionResponse)(nil), // 1: pb.CreateNSQuestionOptionResponse
	(*FindNSQuestionOptionRequest)(nil),    // 2: pb.FindNSQuestionOptionRequest
	(*FindNSQuestionOptionResponse)(nil),   // 3: pb.FindNSQuestionOptionResponse
	(*DeleteNSQuestionOptionRequest)(nil),  // 4: pb.DeleteNSQuestionOptionRequest
	(*NSQuestionOption)(nil),               // 5: pb.NSQuestionOption
	(*RPCSuccess)(nil),                     // 6: pb.RPCSuccess
}
var file_service_ns_question_option_proto_depIdxs = []int32{
	5, // 0: pb.FindNSQuestionOptionResponse.nsQuestionOption:type_name -> pb.NSQuestionOption
	0, // 1: pb.NSQuestionOptionService.createNSQuestionOption:input_type -> pb.CreateNSQuestionOptionRequest
	2, // 2: pb.NSQuestionOptionService.findNSQuestionOption:input_type -> pb.FindNSQuestionOptionRequest
	4, // 3: pb.NSQuestionOptionService.deleteNSQuestionOption:input_type -> pb.DeleteNSQuestionOptionRequest
	1, // 4: pb.NSQuestionOptionService.createNSQuestionOption:output_type -> pb.CreateNSQuestionOptionResponse
	3, // 5: pb.NSQuestionOptionService.findNSQuestionOption:output_type -> pb.FindNSQuestionOptionResponse
	6, // 6: pb.NSQuestionOptionService.deleteNSQuestionOption:output_type -> pb.RPCSuccess
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_ns_question_option_proto_init() }
func file_service_ns_question_option_proto_init() {
	if File_service_ns_question_option_proto != nil {
		return
	}
	file_models_rpc_messages_proto_init()
	file_models_model_ns_question_option_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_ns_question_option_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNSQuestionOptionRequest); i {
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
		file_service_ns_question_option_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNSQuestionOptionResponse); i {
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
		file_service_ns_question_option_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindNSQuestionOptionRequest); i {
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
		file_service_ns_question_option_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindNSQuestionOptionResponse); i {
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
		file_service_ns_question_option_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNSQuestionOptionRequest); i {
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
			RawDescriptor: file_service_ns_question_option_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_ns_question_option_proto_goTypes,
		DependencyIndexes: file_service_ns_question_option_proto_depIdxs,
		MessageInfos:      file_service_ns_question_option_proto_msgTypes,
	}.Build()
	File_service_ns_question_option_proto = out.File
	file_service_ns_question_option_proto_rawDesc = nil
	file_service_ns_question_option_proto_goTypes = nil
	file_service_ns_question_option_proto_depIdxs = nil
}
