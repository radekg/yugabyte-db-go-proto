// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: yb/tserver/backup.proto

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

// TabletServerBackupServiceClient is the client API for TabletServerBackupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TabletServerBackupServiceClient interface {
	TabletSnapshotOp(ctx context.Context, in *TabletSnapshotOpRequestPB, opts ...grpc.CallOption) (*TabletSnapshotOpResponsePB, error)
}

type tabletServerBackupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTabletServerBackupServiceClient(cc grpc.ClientConnInterface) TabletServerBackupServiceClient {
	return &tabletServerBackupServiceClient{cc}
}

func (c *tabletServerBackupServiceClient) TabletSnapshotOp(ctx context.Context, in *TabletSnapshotOpRequestPB, opts ...grpc.CallOption) (*TabletSnapshotOpResponsePB, error) {
	out := new(TabletSnapshotOpResponsePB)
	err := c.cc.Invoke(ctx, "/yb.tserver.TabletServerBackupService/TabletSnapshotOp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TabletServerBackupServiceServer is the server API for TabletServerBackupService service.
// All implementations should embed UnimplementedTabletServerBackupServiceServer
// for forward compatibility
type TabletServerBackupServiceServer interface {
	TabletSnapshotOp(context.Context, *TabletSnapshotOpRequestPB) (*TabletSnapshotOpResponsePB, error)
}

// UnimplementedTabletServerBackupServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTabletServerBackupServiceServer struct {
}

func (UnimplementedTabletServerBackupServiceServer) TabletSnapshotOp(context.Context, *TabletSnapshotOpRequestPB) (*TabletSnapshotOpResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TabletSnapshotOp not implemented")
}

// UnsafeTabletServerBackupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TabletServerBackupServiceServer will
// result in compilation errors.
type UnsafeTabletServerBackupServiceServer interface {
	mustEmbedUnimplementedTabletServerBackupServiceServer()
}

func RegisterTabletServerBackupServiceServer(s grpc.ServiceRegistrar, srv TabletServerBackupServiceServer) {
	s.RegisterService(&TabletServerBackupService_ServiceDesc, srv)
}

func _TabletServerBackupService_TabletSnapshotOp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TabletSnapshotOpRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TabletServerBackupServiceServer).TabletSnapshotOp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.tserver.TabletServerBackupService/TabletSnapshotOp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TabletServerBackupServiceServer).TabletSnapshotOp(ctx, req.(*TabletSnapshotOpRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

// TabletServerBackupService_ServiceDesc is the grpc.ServiceDesc for TabletServerBackupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TabletServerBackupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yb.tserver.TabletServerBackupService",
	HandlerType: (*TabletServerBackupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TabletSnapshotOp",
			Handler:    _TabletServerBackupService_TabletSnapshotOp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yb/tserver/backup.proto",
}
