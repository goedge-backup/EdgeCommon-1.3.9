// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: service_user_traffic_bill.proto

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

// 列出某个账单下的流量带宽子账单
type FindUserTrafficBillsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserBillId int64 `protobuf:"varint,1,opt,name=userBillId,proto3" json:"userBillId,omitempty"`
}

func (x *FindUserTrafficBillsRequest) Reset() {
	*x = FindUserTrafficBillsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_traffic_bill_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserTrafficBillsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserTrafficBillsRequest) ProtoMessage() {}

func (x *FindUserTrafficBillsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_traffic_bill_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserTrafficBillsRequest.ProtoReflect.Descriptor instead.
func (*FindUserTrafficBillsRequest) Descriptor() ([]byte, []int) {
	return file_service_user_traffic_bill_proto_rawDescGZIP(), []int{0}
}

func (x *FindUserTrafficBillsRequest) GetUserBillId() int64 {
	if x != nil {
		return x.UserBillId
	}
	return 0
}

type FindUserTrafficBillsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTrafficBills []*UserTrafficBill `protobuf:"bytes,1,rep,name=userTrafficBills,proto3" json:"userTrafficBills,omitempty"`
}

func (x *FindUserTrafficBillsResponse) Reset() {
	*x = FindUserTrafficBillsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_traffic_bill_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserTrafficBillsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserTrafficBillsResponse) ProtoMessage() {}

func (x *FindUserTrafficBillsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_traffic_bill_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserTrafficBillsResponse.ProtoReflect.Descriptor instead.
func (*FindUserTrafficBillsResponse) Descriptor() ([]byte, []int) {
	return file_service_user_traffic_bill_proto_rawDescGZIP(), []int{1}
}

func (x *FindUserTrafficBillsResponse) GetUserTrafficBills() []*UserTrafficBill {
	if x != nil {
		return x.UserTrafficBills
	}
	return nil
}

var File_service_user_traffic_bill_proto protoreflect.FileDescriptor

var file_service_user_traffic_bill_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x62, 0x69, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x24, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x5f, 0x62, 0x69, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x1b, 0x46,
	0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x42, 0x69,
	0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x75, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x22, 0x5f, 0x0a, 0x1c, 0x46, 0x69,
	0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x42, 0x69, 0x6c,
	0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x10, 0x75, 0x73,
	0x65, 0x72, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x10, 0x75, 0x73, 0x65, 0x72, 0x54,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x32, 0x73, 0x0a, 0x16, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x42, 0x69, 0x6c, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x14, 0x66, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x1f, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66, 0x66,
	0x69, 0x63, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x42, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_user_traffic_bill_proto_rawDescOnce sync.Once
	file_service_user_traffic_bill_proto_rawDescData = file_service_user_traffic_bill_proto_rawDesc
)

func file_service_user_traffic_bill_proto_rawDescGZIP() []byte {
	file_service_user_traffic_bill_proto_rawDescOnce.Do(func() {
		file_service_user_traffic_bill_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_user_traffic_bill_proto_rawDescData)
	})
	return file_service_user_traffic_bill_proto_rawDescData
}

var file_service_user_traffic_bill_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service_user_traffic_bill_proto_goTypes = []interface{}{
	(*FindUserTrafficBillsRequest)(nil),  // 0: pb.FindUserTrafficBillsRequest
	(*FindUserTrafficBillsResponse)(nil), // 1: pb.FindUserTrafficBillsResponse
	(*UserTrafficBill)(nil),              // 2: pb.UserTrafficBill
}
var file_service_user_traffic_bill_proto_depIdxs = []int32{
	2, // 0: pb.FindUserTrafficBillsResponse.userTrafficBills:type_name -> pb.UserTrafficBill
	0, // 1: pb.UserTrafficBillService.findUserTrafficBills:input_type -> pb.FindUserTrafficBillsRequest
	1, // 2: pb.UserTrafficBillService.findUserTrafficBills:output_type -> pb.FindUserTrafficBillsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_user_traffic_bill_proto_init() }
func file_service_user_traffic_bill_proto_init() {
	if File_service_user_traffic_bill_proto != nil {
		return
	}
	file_models_model_user_traffic_bill_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_user_traffic_bill_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserTrafficBillsRequest); i {
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
		file_service_user_traffic_bill_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserTrafficBillsResponse); i {
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
			RawDescriptor: file_service_user_traffic_bill_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_user_traffic_bill_proto_goTypes,
		DependencyIndexes: file_service_user_traffic_bill_proto_depIdxs,
		MessageInfos:      file_service_user_traffic_bill_proto_msgTypes,
	}.Build()
	File_service_user_traffic_bill_proto = out.File
	file_service_user_traffic_bill_proto_rawDesc = nil
	file_service_user_traffic_bill_proto_goTypes = nil
	file_service_user_traffic_bill_proto_depIdxs = nil
}
