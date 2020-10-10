// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: service_http_firewall_rule_set.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 根据配置创建或修改规则集
type CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirewallRuleSetConfigJSON []byte `protobuf:"bytes,1,opt,name=firewallRuleSetConfigJSON,proto3" json:"firewallRuleSetConfigJSON,omitempty"`
}

func (x *CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest) Reset() {
	*x = CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_firewall_rule_set_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest) ProtoMessage() {}

func (x *CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_firewall_rule_set_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest.ProtoReflect.Descriptor instead.
func (*CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest) Descriptor() ([]byte, []int) {
	return file_service_http_firewall_rule_set_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest) GetFirewallRuleSetConfigJSON() []byte {
	if x != nil {
		return x.FirewallRuleSetConfigJSON
	}
	return nil
}

type CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirewallRuleSetId int64 `protobuf:"varint,1,opt,name=firewallRuleSetId,proto3" json:"firewallRuleSetId,omitempty"`
}

func (x *CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse) Reset() {
	*x = CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_firewall_rule_set_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse) ProtoMessage() {}

func (x *CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_firewall_rule_set_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse.ProtoReflect.Descriptor instead.
func (*CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse) Descriptor() ([]byte, []int) {
	return file_service_http_firewall_rule_set_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse) GetFirewallRuleSetId() int64 {
	if x != nil {
		return x.FirewallRuleSetId
	}
	return 0
}

// 设置开启状态
type UpdateHTTPFirewallRuleSetIsOnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirewallRuleSetId int64 `protobuf:"varint,1,opt,name=firewallRuleSetId,proto3" json:"firewallRuleSetId,omitempty"`
	IsOn              bool  `protobuf:"varint,2,opt,name=isOn,proto3" json:"isOn,omitempty"`
}

func (x *UpdateHTTPFirewallRuleSetIsOnRequest) Reset() {
	*x = UpdateHTTPFirewallRuleSetIsOnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_firewall_rule_set_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHTTPFirewallRuleSetIsOnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHTTPFirewallRuleSetIsOnRequest) ProtoMessage() {}

func (x *UpdateHTTPFirewallRuleSetIsOnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_firewall_rule_set_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHTTPFirewallRuleSetIsOnRequest.ProtoReflect.Descriptor instead.
func (*UpdateHTTPFirewallRuleSetIsOnRequest) Descriptor() ([]byte, []int) {
	return file_service_http_firewall_rule_set_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateHTTPFirewallRuleSetIsOnRequest) GetFirewallRuleSetId() int64 {
	if x != nil {
		return x.FirewallRuleSetId
	}
	return 0
}

func (x *UpdateHTTPFirewallRuleSetIsOnRequest) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

// 查找规则集配置
type FindHTTPFirewallRuleSetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirewallRuleSetId int64 `protobuf:"varint,1,opt,name=firewallRuleSetId,proto3" json:"firewallRuleSetId,omitempty"`
}

func (x *FindHTTPFirewallRuleSetConfigRequest) Reset() {
	*x = FindHTTPFirewallRuleSetConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_firewall_rule_set_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindHTTPFirewallRuleSetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindHTTPFirewallRuleSetConfigRequest) ProtoMessage() {}

func (x *FindHTTPFirewallRuleSetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_firewall_rule_set_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindHTTPFirewallRuleSetConfigRequest.ProtoReflect.Descriptor instead.
func (*FindHTTPFirewallRuleSetConfigRequest) Descriptor() ([]byte, []int) {
	return file_service_http_firewall_rule_set_proto_rawDescGZIP(), []int{3}
}

func (x *FindHTTPFirewallRuleSetConfigRequest) GetFirewallRuleSetId() int64 {
	if x != nil {
		return x.FirewallRuleSetId
	}
	return 0
}

type FindHTTPFirewallRuleSetConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirewallRuleSetJSON []byte `protobuf:"bytes,1,opt,name=firewallRuleSetJSON,proto3" json:"firewallRuleSetJSON,omitempty"`
}

func (x *FindHTTPFirewallRuleSetConfigResponse) Reset() {
	*x = FindHTTPFirewallRuleSetConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_firewall_rule_set_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindHTTPFirewallRuleSetConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindHTTPFirewallRuleSetConfigResponse) ProtoMessage() {}

func (x *FindHTTPFirewallRuleSetConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_firewall_rule_set_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindHTTPFirewallRuleSetConfigResponse.ProtoReflect.Descriptor instead.
func (*FindHTTPFirewallRuleSetConfigResponse) Descriptor() ([]byte, []int) {
	return file_service_http_firewall_rule_set_proto_rawDescGZIP(), []int{4}
}

func (x *FindHTTPFirewallRuleSetConfigResponse) GetFirewallRuleSetJSON() []byte {
	if x != nil {
		return x.FirewallRuleSetJSON
	}
	return nil
}

var File_service_http_firewall_rule_set_proto protoreflect.FileDescriptor

var file_service_http_firewall_rule_set_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x66,
	0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x12, 0x72, 0x70, 0x63, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72,
	0x0a, 0x32, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65,
	0x53, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x19, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c,
	0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4a, 0x53, 0x4f,
	0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x19, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c,
	0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4a, 0x53,
	0x4f, 0x4e, 0x22, 0x63, 0x0a, 0x33, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c,
	0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x66, 0x69, 0x72,
	0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75,
	0x6c, 0x65, 0x53, 0x65, 0x74, 0x49, 0x64, 0x22, 0x68, 0x0a, 0x24, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c,
	0x65, 0x53, 0x65, 0x74, 0x49, 0x73, 0x4f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2c, 0x0a, 0x11, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53,
	0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x66, 0x69, 0x72, 0x65,
	0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f,
	0x6e, 0x22, 0x54, 0x0a, 0x24, 0x46, 0x69, 0x6e, 0x64, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72,
	0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x66, 0x69, 0x72,
	0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75,
	0x6c, 0x65, 0x53, 0x65, 0x74, 0x49, 0x64, 0x22, 0x59, 0x0a, 0x25, 0x46, 0x69, 0x6e, 0x64, 0x48,
	0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x30, 0x0a, 0x13, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65,
	0x53, 0x65, 0x74, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x13, 0x66,
	0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x4a, 0x53,
	0x4f, 0x4e, 0x32, 0x94, 0x03, 0x0a, 0x1a, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77,
	0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x9e, 0x01, 0x0a, 0x2b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c,
	0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x36, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c,
	0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54,
	0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74,
	0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5f, 0x0a, 0x1d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50,
	0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x49,
	0x73, 0x4f, 0x6e, 0x12, 0x28, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48,
	0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53,
	0x65, 0x74, 0x49, 0x73, 0x4f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x74, 0x0a, 0x1d, 0x66, 0x69, 0x6e, 0x64, 0x48, 0x54, 0x54, 0x50, 0x46,
	0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x28, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x48, 0x54,
	0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65,
	0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_http_firewall_rule_set_proto_rawDescOnce sync.Once
	file_service_http_firewall_rule_set_proto_rawDescData = file_service_http_firewall_rule_set_proto_rawDesc
)

func file_service_http_firewall_rule_set_proto_rawDescGZIP() []byte {
	file_service_http_firewall_rule_set_proto_rawDescOnce.Do(func() {
		file_service_http_firewall_rule_set_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_http_firewall_rule_set_proto_rawDescData)
	})
	return file_service_http_firewall_rule_set_proto_rawDescData
}

var file_service_http_firewall_rule_set_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_service_http_firewall_rule_set_proto_goTypes = []interface{}{
	(*CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest)(nil),  // 0: pb.CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest
	(*CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse)(nil), // 1: pb.CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse
	(*UpdateHTTPFirewallRuleSetIsOnRequest)(nil),                // 2: pb.UpdateHTTPFirewallRuleSetIsOnRequest
	(*FindHTTPFirewallRuleSetConfigRequest)(nil),                // 3: pb.FindHTTPFirewallRuleSetConfigRequest
	(*FindHTTPFirewallRuleSetConfigResponse)(nil),               // 4: pb.FindHTTPFirewallRuleSetConfigResponse
	(*RPCUpdateSuccess)(nil),                                    // 5: pb.RPCUpdateSuccess
}
var file_service_http_firewall_rule_set_proto_depIdxs = []int32{
	0, // 0: pb.HTTPFirewallRuleSetService.createOrUpdateHTTPFirewallRuleSetFromConfig:input_type -> pb.CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest
	2, // 1: pb.HTTPFirewallRuleSetService.updateHTTPFirewallRuleSetIsOn:input_type -> pb.UpdateHTTPFirewallRuleSetIsOnRequest
	3, // 2: pb.HTTPFirewallRuleSetService.findHTTPFirewallRuleSetConfig:input_type -> pb.FindHTTPFirewallRuleSetConfigRequest
	1, // 3: pb.HTTPFirewallRuleSetService.createOrUpdateHTTPFirewallRuleSetFromConfig:output_type -> pb.CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse
	5, // 4: pb.HTTPFirewallRuleSetService.updateHTTPFirewallRuleSetIsOn:output_type -> pb.RPCUpdateSuccess
	4, // 5: pb.HTTPFirewallRuleSetService.findHTTPFirewallRuleSetConfig:output_type -> pb.FindHTTPFirewallRuleSetConfigResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_http_firewall_rule_set_proto_init() }
func file_service_http_firewall_rule_set_proto_init() {
	if File_service_http_firewall_rule_set_proto != nil {
		return
	}
	file_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_http_firewall_rule_set_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest); i {
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
		file_service_http_firewall_rule_set_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse); i {
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
		file_service_http_firewall_rule_set_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHTTPFirewallRuleSetIsOnRequest); i {
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
		file_service_http_firewall_rule_set_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindHTTPFirewallRuleSetConfigRequest); i {
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
		file_service_http_firewall_rule_set_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindHTTPFirewallRuleSetConfigResponse); i {
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
			RawDescriptor: file_service_http_firewall_rule_set_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_http_firewall_rule_set_proto_goTypes,
		DependencyIndexes: file_service_http_firewall_rule_set_proto_depIdxs,
		MessageInfos:      file_service_http_firewall_rule_set_proto_msgTypes,
	}.Build()
	File_service_http_firewall_rule_set_proto = out.File
	file_service_http_firewall_rule_set_proto_rawDesc = nil
	file_service_http_firewall_rule_set_proto_goTypes = nil
	file_service_http_firewall_rule_set_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HTTPFirewallRuleSetServiceClient is the client API for HTTPFirewallRuleSetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HTTPFirewallRuleSetServiceClient interface {
	// 根据配置创建或修改规则集
	CreateOrUpdateHTTPFirewallRuleSetFromConfig(ctx context.Context, in *CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest, opts ...grpc.CallOption) (*CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse, error)
	// 设置开启状态
	UpdateHTTPFirewallRuleSetIsOn(ctx context.Context, in *UpdateHTTPFirewallRuleSetIsOnRequest, opts ...grpc.CallOption) (*RPCUpdateSuccess, error)
	// 查找规则集配置
	FindHTTPFirewallRuleSetConfig(ctx context.Context, in *FindHTTPFirewallRuleSetConfigRequest, opts ...grpc.CallOption) (*FindHTTPFirewallRuleSetConfigResponse, error)
}

type hTTPFirewallRuleSetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHTTPFirewallRuleSetServiceClient(cc grpc.ClientConnInterface) HTTPFirewallRuleSetServiceClient {
	return &hTTPFirewallRuleSetServiceClient{cc}
}

func (c *hTTPFirewallRuleSetServiceClient) CreateOrUpdateHTTPFirewallRuleSetFromConfig(ctx context.Context, in *CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest, opts ...grpc.CallOption) (*CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse, error) {
	out := new(CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPFirewallRuleSetService/createOrUpdateHTTPFirewallRuleSetFromConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPFirewallRuleSetServiceClient) UpdateHTTPFirewallRuleSetIsOn(ctx context.Context, in *UpdateHTTPFirewallRuleSetIsOnRequest, opts ...grpc.CallOption) (*RPCUpdateSuccess, error) {
	out := new(RPCUpdateSuccess)
	err := c.cc.Invoke(ctx, "/pb.HTTPFirewallRuleSetService/updateHTTPFirewallRuleSetIsOn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPFirewallRuleSetServiceClient) FindHTTPFirewallRuleSetConfig(ctx context.Context, in *FindHTTPFirewallRuleSetConfigRequest, opts ...grpc.CallOption) (*FindHTTPFirewallRuleSetConfigResponse, error) {
	out := new(FindHTTPFirewallRuleSetConfigResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPFirewallRuleSetService/findHTTPFirewallRuleSetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HTTPFirewallRuleSetServiceServer is the server API for HTTPFirewallRuleSetService service.
type HTTPFirewallRuleSetServiceServer interface {
	// 根据配置创建或修改规则集
	CreateOrUpdateHTTPFirewallRuleSetFromConfig(context.Context, *CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest) (*CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse, error)
	// 设置开启状态
	UpdateHTTPFirewallRuleSetIsOn(context.Context, *UpdateHTTPFirewallRuleSetIsOnRequest) (*RPCUpdateSuccess, error)
	// 查找规则集配置
	FindHTTPFirewallRuleSetConfig(context.Context, *FindHTTPFirewallRuleSetConfigRequest) (*FindHTTPFirewallRuleSetConfigResponse, error)
}

// UnimplementedHTTPFirewallRuleSetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHTTPFirewallRuleSetServiceServer struct {
}

func (*UnimplementedHTTPFirewallRuleSetServiceServer) CreateOrUpdateHTTPFirewallRuleSetFromConfig(context.Context, *CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest) (*CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateHTTPFirewallRuleSetFromConfig not implemented")
}
func (*UnimplementedHTTPFirewallRuleSetServiceServer) UpdateHTTPFirewallRuleSetIsOn(context.Context, *UpdateHTTPFirewallRuleSetIsOnRequest) (*RPCUpdateSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHTTPFirewallRuleSetIsOn not implemented")
}
func (*UnimplementedHTTPFirewallRuleSetServiceServer) FindHTTPFirewallRuleSetConfig(context.Context, *FindHTTPFirewallRuleSetConfigRequest) (*FindHTTPFirewallRuleSetConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindHTTPFirewallRuleSetConfig not implemented")
}

func RegisterHTTPFirewallRuleSetServiceServer(s *grpc.Server, srv HTTPFirewallRuleSetServiceServer) {
	s.RegisterService(&_HTTPFirewallRuleSetService_serviceDesc, srv)
}

func _HTTPFirewallRuleSetService_CreateOrUpdateHTTPFirewallRuleSetFromConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFirewallRuleSetServiceServer).CreateOrUpdateHTTPFirewallRuleSetFromConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPFirewallRuleSetService/CreateOrUpdateHTTPFirewallRuleSetFromConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFirewallRuleSetServiceServer).CreateOrUpdateHTTPFirewallRuleSetFromConfig(ctx, req.(*CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPFirewallRuleSetService_UpdateHTTPFirewallRuleSetIsOn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHTTPFirewallRuleSetIsOnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFirewallRuleSetServiceServer).UpdateHTTPFirewallRuleSetIsOn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPFirewallRuleSetService/UpdateHTTPFirewallRuleSetIsOn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFirewallRuleSetServiceServer).UpdateHTTPFirewallRuleSetIsOn(ctx, req.(*UpdateHTTPFirewallRuleSetIsOnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPFirewallRuleSetService_FindHTTPFirewallRuleSetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindHTTPFirewallRuleSetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFirewallRuleSetServiceServer).FindHTTPFirewallRuleSetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPFirewallRuleSetService/FindHTTPFirewallRuleSetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFirewallRuleSetServiceServer).FindHTTPFirewallRuleSetConfig(ctx, req.(*FindHTTPFirewallRuleSetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HTTPFirewallRuleSetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HTTPFirewallRuleSetService",
	HandlerType: (*HTTPFirewallRuleSetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createOrUpdateHTTPFirewallRuleSetFromConfig",
			Handler:    _HTTPFirewallRuleSetService_CreateOrUpdateHTTPFirewallRuleSetFromConfig_Handler,
		},
		{
			MethodName: "updateHTTPFirewallRuleSetIsOn",
			Handler:    _HTTPFirewallRuleSetService_UpdateHTTPFirewallRuleSetIsOn_Handler,
		},
		{
			MethodName: "findHTTPFirewallRuleSetConfig",
			Handler:    _HTTPFirewallRuleSetService_FindHTTPFirewallRuleSetConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_http_firewall_rule_set.proto",
}