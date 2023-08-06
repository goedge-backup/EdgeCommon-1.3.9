// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: service_acme_provider.proto

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

// 查找所有的服务商
type FindAllACMEProvidersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllACMEProvidersRequest) Reset() {
	*x = FindAllACMEProvidersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_acme_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllACMEProvidersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllACMEProvidersRequest) ProtoMessage() {}

func (x *FindAllACMEProvidersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_acme_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllACMEProvidersRequest.ProtoReflect.Descriptor instead.
func (*FindAllACMEProvidersRequest) Descriptor() ([]byte, []int) {
	return file_service_acme_provider_proto_rawDescGZIP(), []int{0}
}

type FindAllACMEProvidersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AcmeProviders []*ACMEProvider `protobuf:"bytes,1,rep,name=acmeProviders,proto3" json:"acmeProviders,omitempty"`
}

func (x *FindAllACMEProvidersResponse) Reset() {
	*x = FindAllACMEProvidersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_acme_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllACMEProvidersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllACMEProvidersResponse) ProtoMessage() {}

func (x *FindAllACMEProvidersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_acme_provider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllACMEProvidersResponse.ProtoReflect.Descriptor instead.
func (*FindAllACMEProvidersResponse) Descriptor() ([]byte, []int) {
	return file_service_acme_provider_proto_rawDescGZIP(), []int{1}
}

func (x *FindAllACMEProvidersResponse) GetAcmeProviders() []*ACMEProvider {
	if x != nil {
		return x.AcmeProviders
	}
	return nil
}

// 根据代号查找服务商
type FindACMEProviderWithCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AcmeProviderCode string `protobuf:"bytes,1,opt,name=acmeProviderCode,proto3" json:"acmeProviderCode,omitempty"`
}

func (x *FindACMEProviderWithCodeRequest) Reset() {
	*x = FindACMEProviderWithCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_acme_provider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindACMEProviderWithCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindACMEProviderWithCodeRequest) ProtoMessage() {}

func (x *FindACMEProviderWithCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_acme_provider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindACMEProviderWithCodeRequest.ProtoReflect.Descriptor instead.
func (*FindACMEProviderWithCodeRequest) Descriptor() ([]byte, []int) {
	return file_service_acme_provider_proto_rawDescGZIP(), []int{2}
}

func (x *FindACMEProviderWithCodeRequest) GetAcmeProviderCode() string {
	if x != nil {
		return x.AcmeProviderCode
	}
	return ""
}

type FindACMEProviderWithCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AcmeProvider *ACMEProvider `protobuf:"bytes,1,opt,name=acmeProvider,proto3" json:"acmeProvider,omitempty"`
}

func (x *FindACMEProviderWithCodeResponse) Reset() {
	*x = FindACMEProviderWithCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_acme_provider_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindACMEProviderWithCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindACMEProviderWithCodeResponse) ProtoMessage() {}

func (x *FindACMEProviderWithCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_acme_provider_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindACMEProviderWithCodeResponse.ProtoReflect.Descriptor instead.
func (*FindACMEProviderWithCodeResponse) Descriptor() ([]byte, []int) {
	return file_service_acme_provider_proto_rawDescGZIP(), []int{3}
}

func (x *FindACMEProviderWithCodeResponse) GetAcmeProvider() *ACMEProvider {
	if x != nil {
		return x.AcmeProvider
	}
	return nil
}

var File_service_acme_provider_proto protoreflect.FileDescriptor

var file_service_acme_provider_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x6d, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x20, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x61, 0x63, 0x6d, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x1b, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x41, 0x43,
	0x4d, 0x45, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x56, 0x0a, 0x1c, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x41, 0x43, 0x4d,
	0x45, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x36, 0x0a, 0x0d, 0x61, 0x63, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x41,
	0x43, 0x4d, 0x45, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x0d, 0x61, 0x63, 0x6d,
	0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x22, 0x4d, 0x0a, 0x1f, 0x46, 0x69,
	0x6e, 0x64, 0x41, 0x43, 0x4d, 0x45, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x57, 0x69,
	0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a,
	0x10, 0x61, 0x63, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x63, 0x6d, 0x65, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x58, 0x0a, 0x20, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x43, 0x4d, 0x45, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x57, 0x69, 0x74,
	0x68, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a,
	0x0c, 0x61, 0x63, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x43, 0x4d, 0x45, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x0c, 0x61, 0x63, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x32, 0xd7, 0x01, 0x0a, 0x13, 0x41, 0x43, 0x4d, 0x45, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x14, 0x66,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x41, 0x43, 0x4d, 0x45, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c,
	0x41, 0x43, 0x4d, 0x45, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c,
	0x6c, 0x41, 0x43, 0x4d, 0x45, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x18, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x43,
	0x4d, 0x45, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x23, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x43, 0x4d, 0x45,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x43, 0x4d, 0x45, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x57, 0x69, 0x74,
	0x68, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_acme_provider_proto_rawDescOnce sync.Once
	file_service_acme_provider_proto_rawDescData = file_service_acme_provider_proto_rawDesc
)

func file_service_acme_provider_proto_rawDescGZIP() []byte {
	file_service_acme_provider_proto_rawDescOnce.Do(func() {
		file_service_acme_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_acme_provider_proto_rawDescData)
	})
	return file_service_acme_provider_proto_rawDescData
}

var file_service_acme_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_service_acme_provider_proto_goTypes = []interface{}{
	(*FindAllACMEProvidersRequest)(nil),      // 0: pb.FindAllACMEProvidersRequest
	(*FindAllACMEProvidersResponse)(nil),     // 1: pb.FindAllACMEProvidersResponse
	(*FindACMEProviderWithCodeRequest)(nil),  // 2: pb.FindACMEProviderWithCodeRequest
	(*FindACMEProviderWithCodeResponse)(nil), // 3: pb.FindACMEProviderWithCodeResponse
	(*ACMEProvider)(nil),                     // 4: pb.ACMEProvider
}
var file_service_acme_provider_proto_depIdxs = []int32{
	4, // 0: pb.FindAllACMEProvidersResponse.acmeProviders:type_name -> pb.ACMEProvider
	4, // 1: pb.FindACMEProviderWithCodeResponse.acmeProvider:type_name -> pb.ACMEProvider
	0, // 2: pb.ACMEProviderService.findAllACMEProviders:input_type -> pb.FindAllACMEProvidersRequest
	2, // 3: pb.ACMEProviderService.findACMEProviderWithCode:input_type -> pb.FindACMEProviderWithCodeRequest
	1, // 4: pb.ACMEProviderService.findAllACMEProviders:output_type -> pb.FindAllACMEProvidersResponse
	3, // 5: pb.ACMEProviderService.findACMEProviderWithCode:output_type -> pb.FindACMEProviderWithCodeResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_service_acme_provider_proto_init() }
func file_service_acme_provider_proto_init() {
	if File_service_acme_provider_proto != nil {
		return
	}
	file_models_model_acme_provider_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_acme_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllACMEProvidersRequest); i {
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
		file_service_acme_provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllACMEProvidersResponse); i {
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
		file_service_acme_provider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindACMEProviderWithCodeRequest); i {
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
		file_service_acme_provider_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindACMEProviderWithCodeResponse); i {
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
			RawDescriptor: file_service_acme_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_acme_provider_proto_goTypes,
		DependencyIndexes: file_service_acme_provider_proto_depIdxs,
		MessageInfos:      file_service_acme_provider_proto_msgTypes,
	}.Build()
	File_service_acme_provider_proto = out.File
	file_service_acme_provider_proto_rawDesc = nil
	file_service_acme_provider_proto_goTypes = nil
	file_service_acme_provider_proto_depIdxs = nil
}
