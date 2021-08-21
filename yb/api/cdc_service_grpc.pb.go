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

// CDCServiceClient is the client API for CDCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CDCServiceClient interface {
	CreateCDCStream(ctx context.Context, in *CreateCDCStreamRequestPB, opts ...grpc.CallOption) (*CreateCDCStreamResponsePB, error)
	DeleteCDCStream(ctx context.Context, in *DeleteCDCStreamRequestPB, opts ...grpc.CallOption) (*DeleteCDCStreamResponsePB, error)
	ListTablets(ctx context.Context, in *CDCListTabletsRequestPB, opts ...grpc.CallOption) (*CDCListTabletsResponsePB, error)
	GetChanges(ctx context.Context, in *GetChangesRequestPB, opts ...grpc.CallOption) (*GetChangesResponsePB, error)
	GetCheckpoint(ctx context.Context, in *GetCheckpointRequestPB, opts ...grpc.CallOption) (*GetCheckpointResponsePB, error)
	UpdateCdcReplicatedIndex(ctx context.Context, in *UpdateCdcReplicatedIndexRequestPB, opts ...grpc.CallOption) (*UpdateCdcReplicatedIndexResponsePB, error)
	BootstrapProducer(ctx context.Context, in *BootstrapProducerRequestPB, opts ...grpc.CallOption) (*BootstrapProducerResponsePB, error)
	GetLatestEntryOpId(ctx context.Context, in *GetLatestEntryOpIdRequestPB, opts ...grpc.CallOption) (*GetLatestEntryOpIdResponsePB, error)
}

type cDCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCDCServiceClient(cc grpc.ClientConnInterface) CDCServiceClient {
	return &cDCServiceClient{cc}
}

func (c *cDCServiceClient) CreateCDCStream(ctx context.Context, in *CreateCDCStreamRequestPB, opts ...grpc.CallOption) (*CreateCDCStreamResponsePB, error) {
	out := new(CreateCDCStreamResponsePB)
	err := c.cc.Invoke(ctx, "/yb.cdc.CDCService/CreateCDCStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cDCServiceClient) DeleteCDCStream(ctx context.Context, in *DeleteCDCStreamRequestPB, opts ...grpc.CallOption) (*DeleteCDCStreamResponsePB, error) {
	out := new(DeleteCDCStreamResponsePB)
	err := c.cc.Invoke(ctx, "/yb.cdc.CDCService/DeleteCDCStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cDCServiceClient) ListTablets(ctx context.Context, in *CDCListTabletsRequestPB, opts ...grpc.CallOption) (*CDCListTabletsResponsePB, error) {
	out := new(CDCListTabletsResponsePB)
	err := c.cc.Invoke(ctx, "/yb.cdc.CDCService/ListTablets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cDCServiceClient) GetChanges(ctx context.Context, in *GetChangesRequestPB, opts ...grpc.CallOption) (*GetChangesResponsePB, error) {
	out := new(GetChangesResponsePB)
	err := c.cc.Invoke(ctx, "/yb.cdc.CDCService/GetChanges", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cDCServiceClient) GetCheckpoint(ctx context.Context, in *GetCheckpointRequestPB, opts ...grpc.CallOption) (*GetCheckpointResponsePB, error) {
	out := new(GetCheckpointResponsePB)
	err := c.cc.Invoke(ctx, "/yb.cdc.CDCService/GetCheckpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cDCServiceClient) UpdateCdcReplicatedIndex(ctx context.Context, in *UpdateCdcReplicatedIndexRequestPB, opts ...grpc.CallOption) (*UpdateCdcReplicatedIndexResponsePB, error) {
	out := new(UpdateCdcReplicatedIndexResponsePB)
	err := c.cc.Invoke(ctx, "/yb.cdc.CDCService/UpdateCdcReplicatedIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cDCServiceClient) BootstrapProducer(ctx context.Context, in *BootstrapProducerRequestPB, opts ...grpc.CallOption) (*BootstrapProducerResponsePB, error) {
	out := new(BootstrapProducerResponsePB)
	err := c.cc.Invoke(ctx, "/yb.cdc.CDCService/BootstrapProducer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cDCServiceClient) GetLatestEntryOpId(ctx context.Context, in *GetLatestEntryOpIdRequestPB, opts ...grpc.CallOption) (*GetLatestEntryOpIdResponsePB, error) {
	out := new(GetLatestEntryOpIdResponsePB)
	err := c.cc.Invoke(ctx, "/yb.cdc.CDCService/GetLatestEntryOpId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CDCServiceServer is the server API for CDCService service.
// All implementations should embed UnimplementedCDCServiceServer
// for forward compatibility
type CDCServiceServer interface {
	CreateCDCStream(context.Context, *CreateCDCStreamRequestPB) (*CreateCDCStreamResponsePB, error)
	DeleteCDCStream(context.Context, *DeleteCDCStreamRequestPB) (*DeleteCDCStreamResponsePB, error)
	ListTablets(context.Context, *CDCListTabletsRequestPB) (*CDCListTabletsResponsePB, error)
	GetChanges(context.Context, *GetChangesRequestPB) (*GetChangesResponsePB, error)
	GetCheckpoint(context.Context, *GetCheckpointRequestPB) (*GetCheckpointResponsePB, error)
	UpdateCdcReplicatedIndex(context.Context, *UpdateCdcReplicatedIndexRequestPB) (*UpdateCdcReplicatedIndexResponsePB, error)
	BootstrapProducer(context.Context, *BootstrapProducerRequestPB) (*BootstrapProducerResponsePB, error)
	GetLatestEntryOpId(context.Context, *GetLatestEntryOpIdRequestPB) (*GetLatestEntryOpIdResponsePB, error)
}

// UnimplementedCDCServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCDCServiceServer struct {
}

func (UnimplementedCDCServiceServer) CreateCDCStream(context.Context, *CreateCDCStreamRequestPB) (*CreateCDCStreamResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCDCStream not implemented")
}
func (UnimplementedCDCServiceServer) DeleteCDCStream(context.Context, *DeleteCDCStreamRequestPB) (*DeleteCDCStreamResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCDCStream not implemented")
}
func (UnimplementedCDCServiceServer) ListTablets(context.Context, *CDCListTabletsRequestPB) (*CDCListTabletsResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTablets not implemented")
}
func (UnimplementedCDCServiceServer) GetChanges(context.Context, *GetChangesRequestPB) (*GetChangesResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChanges not implemented")
}
func (UnimplementedCDCServiceServer) GetCheckpoint(context.Context, *GetCheckpointRequestPB) (*GetCheckpointResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCheckpoint not implemented")
}
func (UnimplementedCDCServiceServer) UpdateCdcReplicatedIndex(context.Context, *UpdateCdcReplicatedIndexRequestPB) (*UpdateCdcReplicatedIndexResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCdcReplicatedIndex not implemented")
}
func (UnimplementedCDCServiceServer) BootstrapProducer(context.Context, *BootstrapProducerRequestPB) (*BootstrapProducerResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BootstrapProducer not implemented")
}
func (UnimplementedCDCServiceServer) GetLatestEntryOpId(context.Context, *GetLatestEntryOpIdRequestPB) (*GetLatestEntryOpIdResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestEntryOpId not implemented")
}

// UnsafeCDCServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CDCServiceServer will
// result in compilation errors.
type UnsafeCDCServiceServer interface {
	mustEmbedUnimplementedCDCServiceServer()
}

func RegisterCDCServiceServer(s grpc.ServiceRegistrar, srv CDCServiceServer) {
	s.RegisterService(&CDCService_ServiceDesc, srv)
}

func _CDCService_CreateCDCStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCDCStreamRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CDCServiceServer).CreateCDCStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.cdc.CDCService/CreateCDCStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CDCServiceServer).CreateCDCStream(ctx, req.(*CreateCDCStreamRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _CDCService_DeleteCDCStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCDCStreamRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CDCServiceServer).DeleteCDCStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.cdc.CDCService/DeleteCDCStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CDCServiceServer).DeleteCDCStream(ctx, req.(*DeleteCDCStreamRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _CDCService_ListTablets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CDCListTabletsRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CDCServiceServer).ListTablets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.cdc.CDCService/ListTablets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CDCServiceServer).ListTablets(ctx, req.(*CDCListTabletsRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _CDCService_GetChanges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChangesRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CDCServiceServer).GetChanges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.cdc.CDCService/GetChanges",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CDCServiceServer).GetChanges(ctx, req.(*GetChangesRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _CDCService_GetCheckpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCheckpointRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CDCServiceServer).GetCheckpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.cdc.CDCService/GetCheckpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CDCServiceServer).GetCheckpoint(ctx, req.(*GetCheckpointRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _CDCService_UpdateCdcReplicatedIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCdcReplicatedIndexRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CDCServiceServer).UpdateCdcReplicatedIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.cdc.CDCService/UpdateCdcReplicatedIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CDCServiceServer).UpdateCdcReplicatedIndex(ctx, req.(*UpdateCdcReplicatedIndexRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _CDCService_BootstrapProducer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BootstrapProducerRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CDCServiceServer).BootstrapProducer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.cdc.CDCService/BootstrapProducer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CDCServiceServer).BootstrapProducer(ctx, req.(*BootstrapProducerRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _CDCService_GetLatestEntryOpId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestEntryOpIdRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CDCServiceServer).GetLatestEntryOpId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.cdc.CDCService/GetLatestEntryOpId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CDCServiceServer).GetLatestEntryOpId(ctx, req.(*GetLatestEntryOpIdRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

// CDCService_ServiceDesc is the grpc.ServiceDesc for CDCService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CDCService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yb.cdc.CDCService",
	HandlerType: (*CDCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCDCStream",
			Handler:    _CDCService_CreateCDCStream_Handler,
		},
		{
			MethodName: "DeleteCDCStream",
			Handler:    _CDCService_DeleteCDCStream_Handler,
		},
		{
			MethodName: "ListTablets",
			Handler:    _CDCService_ListTablets_Handler,
		},
		{
			MethodName: "GetChanges",
			Handler:    _CDCService_GetChanges_Handler,
		},
		{
			MethodName: "GetCheckpoint",
			Handler:    _CDCService_GetCheckpoint_Handler,
		},
		{
			MethodName: "UpdateCdcReplicatedIndex",
			Handler:    _CDCService_UpdateCdcReplicatedIndex_Handler,
		},
		{
			MethodName: "BootstrapProducer",
			Handler:    _CDCService_BootstrapProducer_Handler,
		},
		{
			MethodName: "GetLatestEntryOpId",
			Handler:    _CDCService_GetLatestEntryOpId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yb/cdc/cdc_service.proto",
}
