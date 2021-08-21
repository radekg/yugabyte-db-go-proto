// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// GenericServiceClient is the client API for GenericService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GenericServiceClient interface {
	SetFlag(ctx context.Context, in *SetFlagRequestPB, opts ...grpc.CallOption) (*SetFlagResponsePB, error)
	GetFlag(ctx context.Context, in *GetFlagRequestPB, opts ...grpc.CallOption) (*GetFlagResponsePB, error)
	RefreshFlags(ctx context.Context, in *RefreshFlagsRequestPB, opts ...grpc.CallOption) (*RefreshFlagsResponsePB, error)
	FlushCoverage(ctx context.Context, in *FlushCoverageRequestPB, opts ...grpc.CallOption) (*FlushCoverageResponsePB, error)
	ServerClock(ctx context.Context, in *ServerClockRequestPB, opts ...grpc.CallOption) (*ServerClockResponsePB, error)
	GetStatus(ctx context.Context, in *GetStatusRequestPB, opts ...grpc.CallOption) (*GetStatusResponsePB, error)
	Ping(ctx context.Context, in *PingRequestPB, opts ...grpc.CallOption) (*PingResponsePB, error)
}

type genericServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGenericServiceClient(cc grpc.ClientConnInterface) GenericServiceClient {
	return &genericServiceClient{cc}
}

func (c *genericServiceClient) SetFlag(ctx context.Context, in *SetFlagRequestPB, opts ...grpc.CallOption) (*SetFlagResponsePB, error) {
	out := new(SetFlagResponsePB)
	err := c.cc.Invoke(ctx, "/yb.server.GenericService/SetFlag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genericServiceClient) GetFlag(ctx context.Context, in *GetFlagRequestPB, opts ...grpc.CallOption) (*GetFlagResponsePB, error) {
	out := new(GetFlagResponsePB)
	err := c.cc.Invoke(ctx, "/yb.server.GenericService/GetFlag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genericServiceClient) RefreshFlags(ctx context.Context, in *RefreshFlagsRequestPB, opts ...grpc.CallOption) (*RefreshFlagsResponsePB, error) {
	out := new(RefreshFlagsResponsePB)
	err := c.cc.Invoke(ctx, "/yb.server.GenericService/RefreshFlags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genericServiceClient) FlushCoverage(ctx context.Context, in *FlushCoverageRequestPB, opts ...grpc.CallOption) (*FlushCoverageResponsePB, error) {
	out := new(FlushCoverageResponsePB)
	err := c.cc.Invoke(ctx, "/yb.server.GenericService/FlushCoverage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genericServiceClient) ServerClock(ctx context.Context, in *ServerClockRequestPB, opts ...grpc.CallOption) (*ServerClockResponsePB, error) {
	out := new(ServerClockResponsePB)
	err := c.cc.Invoke(ctx, "/yb.server.GenericService/ServerClock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genericServiceClient) GetStatus(ctx context.Context, in *GetStatusRequestPB, opts ...grpc.CallOption) (*GetStatusResponsePB, error) {
	out := new(GetStatusResponsePB)
	err := c.cc.Invoke(ctx, "/yb.server.GenericService/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genericServiceClient) Ping(ctx context.Context, in *PingRequestPB, opts ...grpc.CallOption) (*PingResponsePB, error) {
	out := new(PingResponsePB)
	err := c.cc.Invoke(ctx, "/yb.server.GenericService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenericServiceServer is the server API for GenericService service.
// All implementations should embed UnimplementedGenericServiceServer
// for forward compatibility
type GenericServiceServer interface {
	SetFlag(context.Context, *SetFlagRequestPB) (*SetFlagResponsePB, error)
	GetFlag(context.Context, *GetFlagRequestPB) (*GetFlagResponsePB, error)
	RefreshFlags(context.Context, *RefreshFlagsRequestPB) (*RefreshFlagsResponsePB, error)
	FlushCoverage(context.Context, *FlushCoverageRequestPB) (*FlushCoverageResponsePB, error)
	ServerClock(context.Context, *ServerClockRequestPB) (*ServerClockResponsePB, error)
	GetStatus(context.Context, *GetStatusRequestPB) (*GetStatusResponsePB, error)
	Ping(context.Context, *PingRequestPB) (*PingResponsePB, error)
}

// UnimplementedGenericServiceServer should be embedded to have forward compatible implementations.
type UnimplementedGenericServiceServer struct {
}

func (UnimplementedGenericServiceServer) SetFlag(context.Context, *SetFlagRequestPB) (*SetFlagResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFlag not implemented")
}
func (UnimplementedGenericServiceServer) GetFlag(context.Context, *GetFlagRequestPB) (*GetFlagResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlag not implemented")
}
func (UnimplementedGenericServiceServer) RefreshFlags(context.Context, *RefreshFlagsRequestPB) (*RefreshFlagsResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshFlags not implemented")
}
func (UnimplementedGenericServiceServer) FlushCoverage(context.Context, *FlushCoverageRequestPB) (*FlushCoverageResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlushCoverage not implemented")
}
func (UnimplementedGenericServiceServer) ServerClock(context.Context, *ServerClockRequestPB) (*ServerClockResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerClock not implemented")
}
func (UnimplementedGenericServiceServer) GetStatus(context.Context, *GetStatusRequestPB) (*GetStatusResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedGenericServiceServer) Ping(context.Context, *PingRequestPB) (*PingResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

// UnsafeGenericServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GenericServiceServer will
// result in compilation errors.
type UnsafeGenericServiceServer interface {
	mustEmbedUnimplementedGenericServiceServer()
}

func RegisterGenericServiceServer(s grpc.ServiceRegistrar, srv GenericServiceServer) {
	s.RegisterService(&GenericService_ServiceDesc, srv)
}

func _GenericService_SetFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFlagRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenericServiceServer).SetFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.server.GenericService/SetFlag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenericServiceServer).SetFlag(ctx, req.(*SetFlagRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenericService_GetFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFlagRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenericServiceServer).GetFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.server.GenericService/GetFlag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenericServiceServer).GetFlag(ctx, req.(*GetFlagRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenericService_RefreshFlags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshFlagsRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenericServiceServer).RefreshFlags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.server.GenericService/RefreshFlags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenericServiceServer).RefreshFlags(ctx, req.(*RefreshFlagsRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenericService_FlushCoverage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushCoverageRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenericServiceServer).FlushCoverage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.server.GenericService/FlushCoverage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenericServiceServer).FlushCoverage(ctx, req.(*FlushCoverageRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenericService_ServerClock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerClockRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenericServiceServer).ServerClock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.server.GenericService/ServerClock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenericServiceServer).ServerClock(ctx, req.(*ServerClockRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenericService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenericServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.server.GenericService/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenericServiceServer).GetStatus(ctx, req.(*GetStatusRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenericService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenericServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.server.GenericService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenericServiceServer).Ping(ctx, req.(*PingRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

// GenericService_ServiceDesc is the grpc.ServiceDesc for GenericService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GenericService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yb.server.GenericService",
	HandlerType: (*GenericServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetFlag",
			Handler:    _GenericService_SetFlag_Handler,
		},
		{
			MethodName: "GetFlag",
			Handler:    _GenericService_GetFlag_Handler,
		},
		{
			MethodName: "RefreshFlags",
			Handler:    _GenericService_RefreshFlags_Handler,
		},
		{
			MethodName: "FlushCoverage",
			Handler:    _GenericService_FlushCoverage_Handler,
		},
		{
			MethodName: "ServerClock",
			Handler:    _GenericService_ServerClock_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _GenericService_GetStatus_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _GenericService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yb/server/server_base.proto",
}