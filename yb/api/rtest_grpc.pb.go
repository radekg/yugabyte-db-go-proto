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

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Add(ctx context.Context, in *AddRequestPB, opts ...grpc.CallOption) (*AddResponsePB, error)
	Sleep(ctx context.Context, in *SleepRequestPB, opts ...grpc.CallOption) (*SleepResponsePB, error)
	Echo(ctx context.Context, in *EchoRequestPB, opts ...grpc.CallOption) (*EchoResponsePB, error)
	WhoAmI(ctx context.Context, in *WhoAmIRequestPB, opts ...grpc.CallOption) (*WhoAmIResponsePB, error)
	TestArgumentsInDiffPackage(ctx context.Context, in *ReqDiffPackagePB, opts ...grpc.CallOption) (*RespDiffPackagePB, error)
	Panic(ctx context.Context, in *PanicRequestPB, opts ...grpc.CallOption) (*PanicResponsePB, error)
	Ping(ctx context.Context, in *RTestPingRequestPB, opts ...grpc.CallOption) (*RTestPingResponsePB, error)
	Disconnect(ctx context.Context, in *DisconnectRequestPB, opts ...grpc.CallOption) (*DisconnectResponsePB, error)
	Forward(ctx context.Context, in *ForwardRequestPB, opts ...grpc.CallOption) (*ForwardResponsePB, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Add(ctx context.Context, in *AddRequestPB, opts ...grpc.CallOption) (*AddResponsePB, error) {
	out := new(AddResponsePB)
	err := c.cc.Invoke(ctx, "/yb.rpc_test.CalculatorService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Sleep(ctx context.Context, in *SleepRequestPB, opts ...grpc.CallOption) (*SleepResponsePB, error) {
	out := new(SleepResponsePB)
	err := c.cc.Invoke(ctx, "/yb.rpc_test.CalculatorService/Sleep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Echo(ctx context.Context, in *EchoRequestPB, opts ...grpc.CallOption) (*EchoResponsePB, error) {
	out := new(EchoResponsePB)
	err := c.cc.Invoke(ctx, "/yb.rpc_test.CalculatorService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) WhoAmI(ctx context.Context, in *WhoAmIRequestPB, opts ...grpc.CallOption) (*WhoAmIResponsePB, error) {
	out := new(WhoAmIResponsePB)
	err := c.cc.Invoke(ctx, "/yb.rpc_test.CalculatorService/WhoAmI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) TestArgumentsInDiffPackage(ctx context.Context, in *ReqDiffPackagePB, opts ...grpc.CallOption) (*RespDiffPackagePB, error) {
	out := new(RespDiffPackagePB)
	err := c.cc.Invoke(ctx, "/yb.rpc_test.CalculatorService/TestArgumentsInDiffPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Panic(ctx context.Context, in *PanicRequestPB, opts ...grpc.CallOption) (*PanicResponsePB, error) {
	out := new(PanicResponsePB)
	err := c.cc.Invoke(ctx, "/yb.rpc_test.CalculatorService/Panic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Ping(ctx context.Context, in *RTestPingRequestPB, opts ...grpc.CallOption) (*RTestPingResponsePB, error) {
	out := new(RTestPingResponsePB)
	err := c.cc.Invoke(ctx, "/yb.rpc_test.CalculatorService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Disconnect(ctx context.Context, in *DisconnectRequestPB, opts ...grpc.CallOption) (*DisconnectResponsePB, error) {
	out := new(DisconnectResponsePB)
	err := c.cc.Invoke(ctx, "/yb.rpc_test.CalculatorService/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Forward(ctx context.Context, in *ForwardRequestPB, opts ...grpc.CallOption) (*ForwardResponsePB, error) {
	out := new(ForwardResponsePB)
	err := c.cc.Invoke(ctx, "/yb.rpc_test.CalculatorService/Forward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
// All implementations should embed UnimplementedCalculatorServiceServer
// for forward compatibility
type CalculatorServiceServer interface {
	Add(context.Context, *AddRequestPB) (*AddResponsePB, error)
	Sleep(context.Context, *SleepRequestPB) (*SleepResponsePB, error)
	Echo(context.Context, *EchoRequestPB) (*EchoResponsePB, error)
	WhoAmI(context.Context, *WhoAmIRequestPB) (*WhoAmIResponsePB, error)
	TestArgumentsInDiffPackage(context.Context, *ReqDiffPackagePB) (*RespDiffPackagePB, error)
	Panic(context.Context, *PanicRequestPB) (*PanicResponsePB, error)
	Ping(context.Context, *RTestPingRequestPB) (*RTestPingResponsePB, error)
	Disconnect(context.Context, *DisconnectRequestPB) (*DisconnectResponsePB, error)
	Forward(context.Context, *ForwardRequestPB) (*ForwardResponsePB, error)
}

// UnimplementedCalculatorServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (UnimplementedCalculatorServiceServer) Add(context.Context, *AddRequestPB) (*AddResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCalculatorServiceServer) Sleep(context.Context, *SleepRequestPB) (*SleepResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sleep not implemented")
}
func (UnimplementedCalculatorServiceServer) Echo(context.Context, *EchoRequestPB) (*EchoResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedCalculatorServiceServer) WhoAmI(context.Context, *WhoAmIRequestPB) (*WhoAmIResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhoAmI not implemented")
}
func (UnimplementedCalculatorServiceServer) TestArgumentsInDiffPackage(context.Context, *ReqDiffPackagePB) (*RespDiffPackagePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestArgumentsInDiffPackage not implemented")
}
func (UnimplementedCalculatorServiceServer) Panic(context.Context, *PanicRequestPB) (*PanicResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Panic not implemented")
}
func (UnimplementedCalculatorServiceServer) Ping(context.Context, *RTestPingRequestPB) (*RTestPingResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedCalculatorServiceServer) Disconnect(context.Context, *DisconnectRequestPB) (*DisconnectResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (UnimplementedCalculatorServiceServer) Forward(context.Context, *ForwardRequestPB) (*ForwardResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Forward not implemented")
}

// UnsafeCalculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServiceServer will
// result in compilation errors.
type UnsafeCalculatorServiceServer interface {
	mustEmbedUnimplementedCalculatorServiceServer()
}

func RegisterCalculatorServiceServer(s grpc.ServiceRegistrar, srv CalculatorServiceServer) {
	s.RegisterService(&CalculatorService_ServiceDesc, srv)
}

func _CalculatorService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.rpc_test.CalculatorService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Add(ctx, req.(*AddRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Sleep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SleepRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sleep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.rpc_test.CalculatorService/Sleep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sleep(ctx, req.(*SleepRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.rpc_test.CalculatorService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Echo(ctx, req.(*EchoRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_WhoAmI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhoAmIRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).WhoAmI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.rpc_test.CalculatorService/WhoAmI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).WhoAmI(ctx, req.(*WhoAmIRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_TestArgumentsInDiffPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqDiffPackagePB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).TestArgumentsInDiffPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.rpc_test.CalculatorService/TestArgumentsInDiffPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).TestArgumentsInDiffPackage(ctx, req.(*ReqDiffPackagePB))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Panic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PanicRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Panic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.rpc_test.CalculatorService/Panic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Panic(ctx, req.(*PanicRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RTestPingRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.rpc_test.CalculatorService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Ping(ctx, req.(*RTestPingRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.rpc_test.CalculatorService/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Disconnect(ctx, req.(*DisconnectRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Forward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForwardRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Forward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.rpc_test.CalculatorService/Forward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Forward(ctx, req.(*ForwardRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

// CalculatorService_ServiceDesc is the grpc.ServiceDesc for CalculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yb.rpc_test.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _CalculatorService_Add_Handler,
		},
		{
			MethodName: "Sleep",
			Handler:    _CalculatorService_Sleep_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _CalculatorService_Echo_Handler,
		},
		{
			MethodName: "WhoAmI",
			Handler:    _CalculatorService_WhoAmI_Handler,
		},
		{
			MethodName: "TestArgumentsInDiffPackage",
			Handler:    _CalculatorService_TestArgumentsInDiffPackage_Handler,
		},
		{
			MethodName: "Panic",
			Handler:    _CalculatorService_Panic_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _CalculatorService_Ping_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _CalculatorService_Disconnect_Handler,
		},
		{
			MethodName: "Forward",
			Handler:    _CalculatorService_Forward_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yb/rpc/rtest.proto",
}