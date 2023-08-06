// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_dns.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DNSService_FindAllDNSIssues_FullMethodName = "/pb.DNSService/findAllDNSIssues"
)

// DNSServiceClient is the client API for DNSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DNSServiceClient interface {
	// 查找问题
	FindAllDNSIssues(ctx context.Context, in *FindAllDNSIssuesRequest, opts ...grpc.CallOption) (*FindAllDNSIssuesResponse, error)
}

type dNSServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDNSServiceClient(cc grpc.ClientConnInterface) DNSServiceClient {
	return &dNSServiceClient{cc}
}

func (c *dNSServiceClient) FindAllDNSIssues(ctx context.Context, in *FindAllDNSIssuesRequest, opts ...grpc.CallOption) (*FindAllDNSIssuesResponse, error) {
	out := new(FindAllDNSIssuesResponse)
	err := c.cc.Invoke(ctx, DNSService_FindAllDNSIssues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DNSServiceServer is the server API for DNSService service.
// All implementations should embed UnimplementedDNSServiceServer
// for forward compatibility
type DNSServiceServer interface {
	// 查找问题
	FindAllDNSIssues(context.Context, *FindAllDNSIssuesRequest) (*FindAllDNSIssuesResponse, error)
}

// UnimplementedDNSServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDNSServiceServer struct {
}

func (UnimplementedDNSServiceServer) FindAllDNSIssues(context.Context, *FindAllDNSIssuesRequest) (*FindAllDNSIssuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllDNSIssues not implemented")
}

// UnsafeDNSServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DNSServiceServer will
// result in compilation errors.
type UnsafeDNSServiceServer interface {
	mustEmbedUnimplementedDNSServiceServer()
}

func RegisterDNSServiceServer(s grpc.ServiceRegistrar, srv DNSServiceServer) {
	s.RegisterService(&DNSService_ServiceDesc, srv)
}

func _DNSService_FindAllDNSIssues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllDNSIssuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSServiceServer).FindAllDNSIssues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DNSService_FindAllDNSIssues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSServiceServer).FindAllDNSIssues(ctx, req.(*FindAllDNSIssuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DNSService_ServiceDesc is the grpc.ServiceDesc for DNSService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DNSService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DNSService",
	HandlerType: (*DNSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findAllDNSIssues",
			Handler:    _DNSService_FindAllDNSIssues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_dns.proto",
}
