// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_ssl_cert.proto

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
	SSLCertService_CreateSSLCert_FullMethodName                 = "/pb.SSLCertService/createSSLCert"
	SSLCertService_CreateSSLCerts_FullMethodName                = "/pb.SSLCertService/createSSLCerts"
	SSLCertService_UpdateSSLCert_FullMethodName                 = "/pb.SSLCertService/updateSSLCert"
	SSLCertService_DeleteSSLCert_FullMethodName                 = "/pb.SSLCertService/deleteSSLCert"
	SSLCertService_FindEnabledSSLCertConfig_FullMethodName      = "/pb.SSLCertService/findEnabledSSLCertConfig"
	SSLCertService_CountSSLCerts_FullMethodName                 = "/pb.SSLCertService/countSSLCerts"
	SSLCertService_ListSSLCerts_FullMethodName                  = "/pb.SSLCertService/listSSLCerts"
	SSLCertService_CountAllSSLCertsWithOCSPError_FullMethodName = "/pb.SSLCertService/countAllSSLCertsWithOCSPError"
	SSLCertService_ListSSLCertsWithOCSPError_FullMethodName     = "/pb.SSLCertService/listSSLCertsWithOCSPError"
	SSLCertService_IgnoreSSLCertsWithOCSPError_FullMethodName   = "/pb.SSLCertService/ignoreSSLCertsWithOCSPError"
	SSLCertService_ResetSSLCertsWithOCSPError_FullMethodName    = "/pb.SSLCertService/resetSSLCertsWithOCSPError"
	SSLCertService_ResetAllSSLCertsWithOCSPError_FullMethodName = "/pb.SSLCertService/resetAllSSLCertsWithOCSPError"
	SSLCertService_ListUpdatedSSLCertOCSP_FullMethodName        = "/pb.SSLCertService/listUpdatedSSLCertOCSP"
)

// SSLCertServiceClient is the client API for SSLCertService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SSLCertServiceClient interface {
	// 创建证书
	CreateSSLCert(ctx context.Context, in *CreateSSLCertRequest, opts ...grpc.CallOption) (*CreateSSLCertResponse, error)
	// 创建一组证书
	CreateSSLCerts(ctx context.Context, in *CreateSSLCertsRequest, opts ...grpc.CallOption) (*CreateSSLCertsResponse, error)
	// 修改证书
	UpdateSSLCert(ctx context.Context, in *UpdateSSLCertRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 删除证书
	DeleteSSLCert(ctx context.Context, in *DeleteSSLCertRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查找证书配置
	FindEnabledSSLCertConfig(ctx context.Context, in *FindEnabledSSLCertConfigRequest, opts ...grpc.CallOption) (*FindEnabledSSLCertConfigResponse, error)
	// 计算匹配的证书数量
	CountSSLCerts(ctx context.Context, in *CountSSLCertRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页匹配的证书
	ListSSLCerts(ctx context.Context, in *ListSSLCertsRequest, opts ...grpc.CallOption) (*ListSSLCertsResponse, error)
	// 计算有OCSP错误的证书数量
	CountAllSSLCertsWithOCSPError(ctx context.Context, in *CountAllSSLCertsWithOCSPErrorRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出有OCSP错误的证书
	ListSSLCertsWithOCSPError(ctx context.Context, in *ListSSLCertsWithOCSPErrorRequest, opts ...grpc.CallOption) (*ListSSLCertsWithOCSPErrorResponse, error)
	// 忽略一组OCSP证书错误
	IgnoreSSLCertsWithOCSPError(ctx context.Context, in *IgnoreSSLCertsWithOCSPErrorRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 重置一组证书OCSP错误状态
	ResetSSLCertsWithOCSPError(ctx context.Context, in *ResetSSLCertsWithOCSPErrorRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 重置所有证书OCSP错误状态
	ResetAllSSLCertsWithOCSPError(ctx context.Context, in *ResetAllSSLCertsWithOCSPErrorRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 读取证书的OCSP
	ListUpdatedSSLCertOCSP(ctx context.Context, in *ListUpdatedSSLCertOCSPRequest, opts ...grpc.CallOption) (*ListUpdatedSSLCertOCSPResponse, error)
}

type sSLCertServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSSLCertServiceClient(cc grpc.ClientConnInterface) SSLCertServiceClient {
	return &sSLCertServiceClient{cc}
}

func (c *sSLCertServiceClient) CreateSSLCert(ctx context.Context, in *CreateSSLCertRequest, opts ...grpc.CallOption) (*CreateSSLCertResponse, error) {
	out := new(CreateSSLCertResponse)
	err := c.cc.Invoke(ctx, SSLCertService_CreateSSLCert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSLCertServiceClient) CreateSSLCerts(ctx context.Context, in *CreateSSLCertsRequest, opts ...grpc.CallOption) (*CreateSSLCertsResponse, error) {
	out := new(CreateSSLCertsResponse)
	err := c.cc.Invoke(ctx, SSLCertService_CreateSSLCerts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSLCertServiceClient) UpdateSSLCert(ctx context.Context, in *UpdateSSLCertRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, SSLCertService_UpdateSSLCert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSLCertServiceClient) DeleteSSLCert(ctx context.Context, in *DeleteSSLCertRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, SSLCertService_DeleteSSLCert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSLCertServiceClient) FindEnabledSSLCertConfig(ctx context.Context, in *FindEnabledSSLCertConfigRequest, opts ...grpc.CallOption) (*FindEnabledSSLCertConfigResponse, error) {
	out := new(FindEnabledSSLCertConfigResponse)
	err := c.cc.Invoke(ctx, SSLCertService_FindEnabledSSLCertConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSLCertServiceClient) CountSSLCerts(ctx context.Context, in *CountSSLCertRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, SSLCertService_CountSSLCerts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSLCertServiceClient) ListSSLCerts(ctx context.Context, in *ListSSLCertsRequest, opts ...grpc.CallOption) (*ListSSLCertsResponse, error) {
	out := new(ListSSLCertsResponse)
	err := c.cc.Invoke(ctx, SSLCertService_ListSSLCerts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSLCertServiceClient) CountAllSSLCertsWithOCSPError(ctx context.Context, in *CountAllSSLCertsWithOCSPErrorRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, SSLCertService_CountAllSSLCertsWithOCSPError_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSLCertServiceClient) ListSSLCertsWithOCSPError(ctx context.Context, in *ListSSLCertsWithOCSPErrorRequest, opts ...grpc.CallOption) (*ListSSLCertsWithOCSPErrorResponse, error) {
	out := new(ListSSLCertsWithOCSPErrorResponse)
	err := c.cc.Invoke(ctx, SSLCertService_ListSSLCertsWithOCSPError_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSLCertServiceClient) IgnoreSSLCertsWithOCSPError(ctx context.Context, in *IgnoreSSLCertsWithOCSPErrorRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, SSLCertService_IgnoreSSLCertsWithOCSPError_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSLCertServiceClient) ResetSSLCertsWithOCSPError(ctx context.Context, in *ResetSSLCertsWithOCSPErrorRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, SSLCertService_ResetSSLCertsWithOCSPError_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSLCertServiceClient) ResetAllSSLCertsWithOCSPError(ctx context.Context, in *ResetAllSSLCertsWithOCSPErrorRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, SSLCertService_ResetAllSSLCertsWithOCSPError_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSLCertServiceClient) ListUpdatedSSLCertOCSP(ctx context.Context, in *ListUpdatedSSLCertOCSPRequest, opts ...grpc.CallOption) (*ListUpdatedSSLCertOCSPResponse, error) {
	out := new(ListUpdatedSSLCertOCSPResponse)
	err := c.cc.Invoke(ctx, SSLCertService_ListUpdatedSSLCertOCSP_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SSLCertServiceServer is the server API for SSLCertService service.
// All implementations should embed UnimplementedSSLCertServiceServer
// for forward compatibility
type SSLCertServiceServer interface {
	// 创建证书
	CreateSSLCert(context.Context, *CreateSSLCertRequest) (*CreateSSLCertResponse, error)
	// 创建一组证书
	CreateSSLCerts(context.Context, *CreateSSLCertsRequest) (*CreateSSLCertsResponse, error)
	// 修改证书
	UpdateSSLCert(context.Context, *UpdateSSLCertRequest) (*RPCSuccess, error)
	// 删除证书
	DeleteSSLCert(context.Context, *DeleteSSLCertRequest) (*RPCSuccess, error)
	// 查找证书配置
	FindEnabledSSLCertConfig(context.Context, *FindEnabledSSLCertConfigRequest) (*FindEnabledSSLCertConfigResponse, error)
	// 计算匹配的证书数量
	CountSSLCerts(context.Context, *CountSSLCertRequest) (*RPCCountResponse, error)
	// 列出单页匹配的证书
	ListSSLCerts(context.Context, *ListSSLCertsRequest) (*ListSSLCertsResponse, error)
	// 计算有OCSP错误的证书数量
	CountAllSSLCertsWithOCSPError(context.Context, *CountAllSSLCertsWithOCSPErrorRequest) (*RPCCountResponse, error)
	// 列出有OCSP错误的证书
	ListSSLCertsWithOCSPError(context.Context, *ListSSLCertsWithOCSPErrorRequest) (*ListSSLCertsWithOCSPErrorResponse, error)
	// 忽略一组OCSP证书错误
	IgnoreSSLCertsWithOCSPError(context.Context, *IgnoreSSLCertsWithOCSPErrorRequest) (*RPCSuccess, error)
	// 重置一组证书OCSP错误状态
	ResetSSLCertsWithOCSPError(context.Context, *ResetSSLCertsWithOCSPErrorRequest) (*RPCSuccess, error)
	// 重置所有证书OCSP错误状态
	ResetAllSSLCertsWithOCSPError(context.Context, *ResetAllSSLCertsWithOCSPErrorRequest) (*RPCSuccess, error)
	// 读取证书的OCSP
	ListUpdatedSSLCertOCSP(context.Context, *ListUpdatedSSLCertOCSPRequest) (*ListUpdatedSSLCertOCSPResponse, error)
}

// UnimplementedSSLCertServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSSLCertServiceServer struct {
}

func (UnimplementedSSLCertServiceServer) CreateSSLCert(context.Context, *CreateSSLCertRequest) (*CreateSSLCertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSSLCert not implemented")
}
func (UnimplementedSSLCertServiceServer) CreateSSLCerts(context.Context, *CreateSSLCertsRequest) (*CreateSSLCertsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSSLCerts not implemented")
}
func (UnimplementedSSLCertServiceServer) UpdateSSLCert(context.Context, *UpdateSSLCertRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSSLCert not implemented")
}
func (UnimplementedSSLCertServiceServer) DeleteSSLCert(context.Context, *DeleteSSLCertRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSSLCert not implemented")
}
func (UnimplementedSSLCertServiceServer) FindEnabledSSLCertConfig(context.Context, *FindEnabledSSLCertConfigRequest) (*FindEnabledSSLCertConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledSSLCertConfig not implemented")
}
func (UnimplementedSSLCertServiceServer) CountSSLCerts(context.Context, *CountSSLCertRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountSSLCerts not implemented")
}
func (UnimplementedSSLCertServiceServer) ListSSLCerts(context.Context, *ListSSLCertsRequest) (*ListSSLCertsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSSLCerts not implemented")
}
func (UnimplementedSSLCertServiceServer) CountAllSSLCertsWithOCSPError(context.Context, *CountAllSSLCertsWithOCSPErrorRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAllSSLCertsWithOCSPError not implemented")
}
func (UnimplementedSSLCertServiceServer) ListSSLCertsWithOCSPError(context.Context, *ListSSLCertsWithOCSPErrorRequest) (*ListSSLCertsWithOCSPErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSSLCertsWithOCSPError not implemented")
}
func (UnimplementedSSLCertServiceServer) IgnoreSSLCertsWithOCSPError(context.Context, *IgnoreSSLCertsWithOCSPErrorRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IgnoreSSLCertsWithOCSPError not implemented")
}
func (UnimplementedSSLCertServiceServer) ResetSSLCertsWithOCSPError(context.Context, *ResetSSLCertsWithOCSPErrorRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetSSLCertsWithOCSPError not implemented")
}
func (UnimplementedSSLCertServiceServer) ResetAllSSLCertsWithOCSPError(context.Context, *ResetAllSSLCertsWithOCSPErrorRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetAllSSLCertsWithOCSPError not implemented")
}
func (UnimplementedSSLCertServiceServer) ListUpdatedSSLCertOCSP(context.Context, *ListUpdatedSSLCertOCSPRequest) (*ListUpdatedSSLCertOCSPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUpdatedSSLCertOCSP not implemented")
}

// UnsafeSSLCertServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SSLCertServiceServer will
// result in compilation errors.
type UnsafeSSLCertServiceServer interface {
	mustEmbedUnimplementedSSLCertServiceServer()
}

func RegisterSSLCertServiceServer(s grpc.ServiceRegistrar, srv SSLCertServiceServer) {
	s.RegisterService(&SSLCertService_ServiceDesc, srv)
}

func _SSLCertService_CreateSSLCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSSLCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSLCertServiceServer).CreateSSLCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSLCertService_CreateSSLCert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSLCertServiceServer).CreateSSLCert(ctx, req.(*CreateSSLCertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSLCertService_CreateSSLCerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSSLCertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSLCertServiceServer).CreateSSLCerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSLCertService_CreateSSLCerts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSLCertServiceServer).CreateSSLCerts(ctx, req.(*CreateSSLCertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSLCertService_UpdateSSLCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSSLCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSLCertServiceServer).UpdateSSLCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSLCertService_UpdateSSLCert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSLCertServiceServer).UpdateSSLCert(ctx, req.(*UpdateSSLCertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSLCertService_DeleteSSLCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSSLCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSLCertServiceServer).DeleteSSLCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSLCertService_DeleteSSLCert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSLCertServiceServer).DeleteSSLCert(ctx, req.(*DeleteSSLCertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSLCertService_FindEnabledSSLCertConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledSSLCertConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSLCertServiceServer).FindEnabledSSLCertConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSLCertService_FindEnabledSSLCertConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSLCertServiceServer).FindEnabledSSLCertConfig(ctx, req.(*FindEnabledSSLCertConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSLCertService_CountSSLCerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountSSLCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSLCertServiceServer).CountSSLCerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSLCertService_CountSSLCerts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSLCertServiceServer).CountSSLCerts(ctx, req.(*CountSSLCertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSLCertService_ListSSLCerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSSLCertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSLCertServiceServer).ListSSLCerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSLCertService_ListSSLCerts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSLCertServiceServer).ListSSLCerts(ctx, req.(*ListSSLCertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSLCertService_CountAllSSLCertsWithOCSPError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAllSSLCertsWithOCSPErrorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSLCertServiceServer).CountAllSSLCertsWithOCSPError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSLCertService_CountAllSSLCertsWithOCSPError_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSLCertServiceServer).CountAllSSLCertsWithOCSPError(ctx, req.(*CountAllSSLCertsWithOCSPErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSLCertService_ListSSLCertsWithOCSPError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSSLCertsWithOCSPErrorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSLCertServiceServer).ListSSLCertsWithOCSPError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSLCertService_ListSSLCertsWithOCSPError_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSLCertServiceServer).ListSSLCertsWithOCSPError(ctx, req.(*ListSSLCertsWithOCSPErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSLCertService_IgnoreSSLCertsWithOCSPError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IgnoreSSLCertsWithOCSPErrorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSLCertServiceServer).IgnoreSSLCertsWithOCSPError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSLCertService_IgnoreSSLCertsWithOCSPError_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSLCertServiceServer).IgnoreSSLCertsWithOCSPError(ctx, req.(*IgnoreSSLCertsWithOCSPErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSLCertService_ResetSSLCertsWithOCSPError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetSSLCertsWithOCSPErrorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSLCertServiceServer).ResetSSLCertsWithOCSPError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSLCertService_ResetSSLCertsWithOCSPError_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSLCertServiceServer).ResetSSLCertsWithOCSPError(ctx, req.(*ResetSSLCertsWithOCSPErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSLCertService_ResetAllSSLCertsWithOCSPError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetAllSSLCertsWithOCSPErrorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSLCertServiceServer).ResetAllSSLCertsWithOCSPError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSLCertService_ResetAllSSLCertsWithOCSPError_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSLCertServiceServer).ResetAllSSLCertsWithOCSPError(ctx, req.(*ResetAllSSLCertsWithOCSPErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSLCertService_ListUpdatedSSLCertOCSP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUpdatedSSLCertOCSPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSLCertServiceServer).ListUpdatedSSLCertOCSP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSLCertService_ListUpdatedSSLCertOCSP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSLCertServiceServer).ListUpdatedSSLCertOCSP(ctx, req.(*ListUpdatedSSLCertOCSPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SSLCertService_ServiceDesc is the grpc.ServiceDesc for SSLCertService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SSLCertService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SSLCertService",
	HandlerType: (*SSLCertServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createSSLCert",
			Handler:    _SSLCertService_CreateSSLCert_Handler,
		},
		{
			MethodName: "createSSLCerts",
			Handler:    _SSLCertService_CreateSSLCerts_Handler,
		},
		{
			MethodName: "updateSSLCert",
			Handler:    _SSLCertService_UpdateSSLCert_Handler,
		},
		{
			MethodName: "deleteSSLCert",
			Handler:    _SSLCertService_DeleteSSLCert_Handler,
		},
		{
			MethodName: "findEnabledSSLCertConfig",
			Handler:    _SSLCertService_FindEnabledSSLCertConfig_Handler,
		},
		{
			MethodName: "countSSLCerts",
			Handler:    _SSLCertService_CountSSLCerts_Handler,
		},
		{
			MethodName: "listSSLCerts",
			Handler:    _SSLCertService_ListSSLCerts_Handler,
		},
		{
			MethodName: "countAllSSLCertsWithOCSPError",
			Handler:    _SSLCertService_CountAllSSLCertsWithOCSPError_Handler,
		},
		{
			MethodName: "listSSLCertsWithOCSPError",
			Handler:    _SSLCertService_ListSSLCertsWithOCSPError_Handler,
		},
		{
			MethodName: "ignoreSSLCertsWithOCSPError",
			Handler:    _SSLCertService_IgnoreSSLCertsWithOCSPError_Handler,
		},
		{
			MethodName: "resetSSLCertsWithOCSPError",
			Handler:    _SSLCertService_ResetSSLCertsWithOCSPError_Handler,
		},
		{
			MethodName: "resetAllSSLCertsWithOCSPError",
			Handler:    _SSLCertService_ResetAllSSLCertsWithOCSPError_Handler,
		},
		{
			MethodName: "listUpdatedSSLCertOCSP",
			Handler:    _SSLCertService_ListUpdatedSSLCertOCSP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_ssl_cert.proto",
}
