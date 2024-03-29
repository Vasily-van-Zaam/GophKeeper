// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: internal/transport/grpc/server.proto

package server

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
	Grpc_GetAccess_FullMethodName        = "/grpc.server.Grpc/GetAccess"
	Grpc_ConfirmAccess_FullMethodName    = "/grpc.server.Grpc/ConfirmAccess"
	Grpc_Ping_FullMethodName             = "/grpc.server.Grpc/Ping"
	Grpc_CheckChangesData_FullMethodName = "/grpc.server.Grpc/CheckChangesData"
	Grpc_AddData_FullMethodName          = "/grpc.server.Grpc/AddData"
	Grpc_ChangeData_FullMethodName       = "/grpc.server.Grpc/ChangeData"
	Grpc_GetData_FullMethodName          = "/grpc.server.Grpc/GetData"
)

// GrpcClient is the client API for Grpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GrpcClient interface {
	GetAccess(ctx context.Context, in *GetAccessRequest, opts ...grpc.CallOption) (*GetAccessResponse, error)
	ConfirmAccess(ctx context.Context, in *ConfirmAccessRequest, opts ...grpc.CallOption) (*ConfirmAccessResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	CheckChangesData(ctx context.Context, in *CheckChangesDataRequest, opts ...grpc.CallOption) (*CheckChangesDataResponse, error)
	AddData(ctx context.Context, in *AddDataRequest, opts ...grpc.CallOption) (*AddDataResponse, error)
	ChangeData(ctx context.Context, in *ChangeDataRequest, opts ...grpc.CallOption) (*ChangeDataResponse, error)
	GetData(ctx context.Context, in *GetDataRequest, opts ...grpc.CallOption) (*GetDataResponse, error)
}

type grpcClient struct {
	cc grpc.ClientConnInterface
}

func NewGrpcClient(cc grpc.ClientConnInterface) GrpcClient {
	return &grpcClient{cc}
}

func (c *grpcClient) GetAccess(ctx context.Context, in *GetAccessRequest, opts ...grpc.CallOption) (*GetAccessResponse, error) {
	out := new(GetAccessResponse)
	err := c.cc.Invoke(ctx, Grpc_GetAccess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) ConfirmAccess(ctx context.Context, in *ConfirmAccessRequest, opts ...grpc.CallOption) (*ConfirmAccessResponse, error) {
	out := new(ConfirmAccessResponse)
	err := c.cc.Invoke(ctx, Grpc_ConfirmAccess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, Grpc_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) CheckChangesData(ctx context.Context, in *CheckChangesDataRequest, opts ...grpc.CallOption) (*CheckChangesDataResponse, error) {
	out := new(CheckChangesDataResponse)
	err := c.cc.Invoke(ctx, Grpc_CheckChangesData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) AddData(ctx context.Context, in *AddDataRequest, opts ...grpc.CallOption) (*AddDataResponse, error) {
	out := new(AddDataResponse)
	err := c.cc.Invoke(ctx, Grpc_AddData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) ChangeData(ctx context.Context, in *ChangeDataRequest, opts ...grpc.CallOption) (*ChangeDataResponse, error) {
	out := new(ChangeDataResponse)
	err := c.cc.Invoke(ctx, Grpc_ChangeData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) GetData(ctx context.Context, in *GetDataRequest, opts ...grpc.CallOption) (*GetDataResponse, error) {
	out := new(GetDataResponse)
	err := c.cc.Invoke(ctx, Grpc_GetData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcServer is the server API for Grpc service.
// All implementations must embed UnimplementedGrpcServer
// for forward compatibility
type GrpcServer interface {
	GetAccess(context.Context, *GetAccessRequest) (*GetAccessResponse, error)
	ConfirmAccess(context.Context, *ConfirmAccessRequest) (*ConfirmAccessResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	CheckChangesData(context.Context, *CheckChangesDataRequest) (*CheckChangesDataResponse, error)
	AddData(context.Context, *AddDataRequest) (*AddDataResponse, error)
	ChangeData(context.Context, *ChangeDataRequest) (*ChangeDataResponse, error)
	GetData(context.Context, *GetDataRequest) (*GetDataResponse, error)
	mustEmbedUnimplementedGrpcServer()
}

// UnimplementedGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedGrpcServer struct {
}

func (UnimplementedGrpcServer) GetAccess(context.Context, *GetAccessRequest) (*GetAccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccess not implemented")
}
func (UnimplementedGrpcServer) ConfirmAccess(context.Context, *ConfirmAccessRequest) (*ConfirmAccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmAccess not implemented")
}
func (UnimplementedGrpcServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedGrpcServer) CheckChangesData(context.Context, *CheckChangesDataRequest) (*CheckChangesDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckChangesData not implemented")
}
func (UnimplementedGrpcServer) AddData(context.Context, *AddDataRequest) (*AddDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddData not implemented")
}
func (UnimplementedGrpcServer) ChangeData(context.Context, *ChangeDataRequest) (*ChangeDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeData not implemented")
}
func (UnimplementedGrpcServer) GetData(context.Context, *GetDataRequest) (*GetDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (UnimplementedGrpcServer) mustEmbedUnimplementedGrpcServer() {}

// UnsafeGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GrpcServer will
// result in compilation errors.
type UnsafeGrpcServer interface {
	mustEmbedUnimplementedGrpcServer()
}

func RegisterGrpcServer(s grpc.ServiceRegistrar, srv GrpcServer) {
	s.RegisterService(&Grpc_ServiceDesc, srv)
}

func _Grpc_GetAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).GetAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Grpc_GetAccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).GetAccess(ctx, req.(*GetAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpc_ConfirmAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).ConfirmAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Grpc_ConfirmAccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).ConfirmAccess(ctx, req.(*ConfirmAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpc_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Grpc_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpc_CheckChangesData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckChangesDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).CheckChangesData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Grpc_CheckChangesData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).CheckChangesData(ctx, req.(*CheckChangesDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpc_AddData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).AddData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Grpc_AddData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).AddData(ctx, req.(*AddDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpc_ChangeData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).ChangeData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Grpc_ChangeData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).ChangeData(ctx, req.(*ChangeDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpc_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Grpc_GetData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).GetData(ctx, req.(*GetDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Grpc_ServiceDesc is the grpc.ServiceDesc for Grpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Grpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.server.Grpc",
	HandlerType: (*GrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccess",
			Handler:    _Grpc_GetAccess_Handler,
		},
		{
			MethodName: "ConfirmAccess",
			Handler:    _Grpc_ConfirmAccess_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Grpc_Ping_Handler,
		},
		{
			MethodName: "CheckChangesData",
			Handler:    _Grpc_CheckChangesData_Handler,
		},
		{
			MethodName: "AddData",
			Handler:    _Grpc_AddData_Handler,
		},
		{
			MethodName: "ChangeData",
			Handler:    _Grpc_ChangeData_Handler,
		},
		{
			MethodName: "GetData",
			Handler:    _Grpc_GetData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/transport/grpc/server.proto",
}
