// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: yb/master/master_admin.proto

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

// MasterAdminClient is the client API for MasterAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterAdminClient interface {
	FlushTables(ctx context.Context, in *FlushTablesRequestPB, opts ...grpc.CallOption) (*FlushTablesResponsePB, error)
	IsFlushTablesDone(ctx context.Context, in *IsFlushTablesDoneRequestPB, opts ...grpc.CallOption) (*IsFlushTablesDoneResponsePB, error)
	FlushSysCatalog(ctx context.Context, in *FlushSysCatalogRequestPB, opts ...grpc.CallOption) (*FlushSysCatalogResponsePB, error)
	CompactSysCatalog(ctx context.Context, in *CompactSysCatalogRequestPB, opts ...grpc.CallOption) (*CompactSysCatalogResponsePB, error)
	IsInitDbDone(ctx context.Context, in *IsInitDbDoneRequestPB, opts ...grpc.CallOption) (*IsInitDbDoneResponsePB, error)
	SplitTablet(ctx context.Context, in *SplitTabletRequestPB, opts ...grpc.CallOption) (*SplitTabletResponsePB, error)
	CreateTransactionStatusTable(ctx context.Context, in *CreateTransactionStatusTableRequestPB, opts ...grpc.CallOption) (*CreateTransactionStatusTableResponsePB, error)
	DeleteNotServingTablet(ctx context.Context, in *DeleteNotServingTabletRequestPB, opts ...grpc.CallOption) (*DeleteNotServingTabletResponsePB, error)
	DdlLog(ctx context.Context, in *DdlLogRequestPB, opts ...grpc.CallOption) (*DdlLogResponsePB, error)
}

type masterAdminClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterAdminClient(cc grpc.ClientConnInterface) MasterAdminClient {
	return &masterAdminClient{cc}
}

func (c *masterAdminClient) FlushTables(ctx context.Context, in *FlushTablesRequestPB, opts ...grpc.CallOption) (*FlushTablesResponsePB, error) {
	out := new(FlushTablesResponsePB)
	err := c.cc.Invoke(ctx, "/yb.master.MasterAdmin/FlushTables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterAdminClient) IsFlushTablesDone(ctx context.Context, in *IsFlushTablesDoneRequestPB, opts ...grpc.CallOption) (*IsFlushTablesDoneResponsePB, error) {
	out := new(IsFlushTablesDoneResponsePB)
	err := c.cc.Invoke(ctx, "/yb.master.MasterAdmin/IsFlushTablesDone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterAdminClient) FlushSysCatalog(ctx context.Context, in *FlushSysCatalogRequestPB, opts ...grpc.CallOption) (*FlushSysCatalogResponsePB, error) {
	out := new(FlushSysCatalogResponsePB)
	err := c.cc.Invoke(ctx, "/yb.master.MasterAdmin/FlushSysCatalog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterAdminClient) CompactSysCatalog(ctx context.Context, in *CompactSysCatalogRequestPB, opts ...grpc.CallOption) (*CompactSysCatalogResponsePB, error) {
	out := new(CompactSysCatalogResponsePB)
	err := c.cc.Invoke(ctx, "/yb.master.MasterAdmin/CompactSysCatalog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterAdminClient) IsInitDbDone(ctx context.Context, in *IsInitDbDoneRequestPB, opts ...grpc.CallOption) (*IsInitDbDoneResponsePB, error) {
	out := new(IsInitDbDoneResponsePB)
	err := c.cc.Invoke(ctx, "/yb.master.MasterAdmin/IsInitDbDone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterAdminClient) SplitTablet(ctx context.Context, in *SplitTabletRequestPB, opts ...grpc.CallOption) (*SplitTabletResponsePB, error) {
	out := new(SplitTabletResponsePB)
	err := c.cc.Invoke(ctx, "/yb.master.MasterAdmin/SplitTablet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterAdminClient) CreateTransactionStatusTable(ctx context.Context, in *CreateTransactionStatusTableRequestPB, opts ...grpc.CallOption) (*CreateTransactionStatusTableResponsePB, error) {
	out := new(CreateTransactionStatusTableResponsePB)
	err := c.cc.Invoke(ctx, "/yb.master.MasterAdmin/CreateTransactionStatusTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterAdminClient) DeleteNotServingTablet(ctx context.Context, in *DeleteNotServingTabletRequestPB, opts ...grpc.CallOption) (*DeleteNotServingTabletResponsePB, error) {
	out := new(DeleteNotServingTabletResponsePB)
	err := c.cc.Invoke(ctx, "/yb.master.MasterAdmin/DeleteNotServingTablet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterAdminClient) DdlLog(ctx context.Context, in *DdlLogRequestPB, opts ...grpc.CallOption) (*DdlLogResponsePB, error) {
	out := new(DdlLogResponsePB)
	err := c.cc.Invoke(ctx, "/yb.master.MasterAdmin/DdlLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterAdminServer is the server API for MasterAdmin service.
// All implementations should embed UnimplementedMasterAdminServer
// for forward compatibility
type MasterAdminServer interface {
	FlushTables(context.Context, *FlushTablesRequestPB) (*FlushTablesResponsePB, error)
	IsFlushTablesDone(context.Context, *IsFlushTablesDoneRequestPB) (*IsFlushTablesDoneResponsePB, error)
	FlushSysCatalog(context.Context, *FlushSysCatalogRequestPB) (*FlushSysCatalogResponsePB, error)
	CompactSysCatalog(context.Context, *CompactSysCatalogRequestPB) (*CompactSysCatalogResponsePB, error)
	IsInitDbDone(context.Context, *IsInitDbDoneRequestPB) (*IsInitDbDoneResponsePB, error)
	SplitTablet(context.Context, *SplitTabletRequestPB) (*SplitTabletResponsePB, error)
	CreateTransactionStatusTable(context.Context, *CreateTransactionStatusTableRequestPB) (*CreateTransactionStatusTableResponsePB, error)
	DeleteNotServingTablet(context.Context, *DeleteNotServingTabletRequestPB) (*DeleteNotServingTabletResponsePB, error)
	DdlLog(context.Context, *DdlLogRequestPB) (*DdlLogResponsePB, error)
}

// UnimplementedMasterAdminServer should be embedded to have forward compatible implementations.
type UnimplementedMasterAdminServer struct {
}

func (UnimplementedMasterAdminServer) FlushTables(context.Context, *FlushTablesRequestPB) (*FlushTablesResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlushTables not implemented")
}
func (UnimplementedMasterAdminServer) IsFlushTablesDone(context.Context, *IsFlushTablesDoneRequestPB) (*IsFlushTablesDoneResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFlushTablesDone not implemented")
}
func (UnimplementedMasterAdminServer) FlushSysCatalog(context.Context, *FlushSysCatalogRequestPB) (*FlushSysCatalogResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlushSysCatalog not implemented")
}
func (UnimplementedMasterAdminServer) CompactSysCatalog(context.Context, *CompactSysCatalogRequestPB) (*CompactSysCatalogResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompactSysCatalog not implemented")
}
func (UnimplementedMasterAdminServer) IsInitDbDone(context.Context, *IsInitDbDoneRequestPB) (*IsInitDbDoneResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsInitDbDone not implemented")
}
func (UnimplementedMasterAdminServer) SplitTablet(context.Context, *SplitTabletRequestPB) (*SplitTabletResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SplitTablet not implemented")
}
func (UnimplementedMasterAdminServer) CreateTransactionStatusTable(context.Context, *CreateTransactionStatusTableRequestPB) (*CreateTransactionStatusTableResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransactionStatusTable not implemented")
}
func (UnimplementedMasterAdminServer) DeleteNotServingTablet(context.Context, *DeleteNotServingTabletRequestPB) (*DeleteNotServingTabletResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNotServingTablet not implemented")
}
func (UnimplementedMasterAdminServer) DdlLog(context.Context, *DdlLogRequestPB) (*DdlLogResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DdlLog not implemented")
}

// UnsafeMasterAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterAdminServer will
// result in compilation errors.
type UnsafeMasterAdminServer interface {
	mustEmbedUnimplementedMasterAdminServer()
}

func RegisterMasterAdminServer(s grpc.ServiceRegistrar, srv MasterAdminServer) {
	s.RegisterService(&MasterAdmin_ServiceDesc, srv)
}

func _MasterAdmin_FlushTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushTablesRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterAdminServer).FlushTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.master.MasterAdmin/FlushTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterAdminServer).FlushTables(ctx, req.(*FlushTablesRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterAdmin_IsFlushTablesDone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsFlushTablesDoneRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterAdminServer).IsFlushTablesDone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.master.MasterAdmin/IsFlushTablesDone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterAdminServer).IsFlushTablesDone(ctx, req.(*IsFlushTablesDoneRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterAdmin_FlushSysCatalog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushSysCatalogRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterAdminServer).FlushSysCatalog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.master.MasterAdmin/FlushSysCatalog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterAdminServer).FlushSysCatalog(ctx, req.(*FlushSysCatalogRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterAdmin_CompactSysCatalog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompactSysCatalogRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterAdminServer).CompactSysCatalog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.master.MasterAdmin/CompactSysCatalog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterAdminServer).CompactSysCatalog(ctx, req.(*CompactSysCatalogRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterAdmin_IsInitDbDone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsInitDbDoneRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterAdminServer).IsInitDbDone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.master.MasterAdmin/IsInitDbDone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterAdminServer).IsInitDbDone(ctx, req.(*IsInitDbDoneRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterAdmin_SplitTablet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SplitTabletRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterAdminServer).SplitTablet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.master.MasterAdmin/SplitTablet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterAdminServer).SplitTablet(ctx, req.(*SplitTabletRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterAdmin_CreateTransactionStatusTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionStatusTableRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterAdminServer).CreateTransactionStatusTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.master.MasterAdmin/CreateTransactionStatusTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterAdminServer).CreateTransactionStatusTable(ctx, req.(*CreateTransactionStatusTableRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterAdmin_DeleteNotServingTablet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNotServingTabletRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterAdminServer).DeleteNotServingTablet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.master.MasterAdmin/DeleteNotServingTablet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterAdminServer).DeleteNotServingTablet(ctx, req.(*DeleteNotServingTabletRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterAdmin_DdlLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DdlLogRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterAdminServer).DdlLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.master.MasterAdmin/DdlLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterAdminServer).DdlLog(ctx, req.(*DdlLogRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

// MasterAdmin_ServiceDesc is the grpc.ServiceDesc for MasterAdmin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MasterAdmin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yb.master.MasterAdmin",
	HandlerType: (*MasterAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FlushTables",
			Handler:    _MasterAdmin_FlushTables_Handler,
		},
		{
			MethodName: "IsFlushTablesDone",
			Handler:    _MasterAdmin_IsFlushTablesDone_Handler,
		},
		{
			MethodName: "FlushSysCatalog",
			Handler:    _MasterAdmin_FlushSysCatalog_Handler,
		},
		{
			MethodName: "CompactSysCatalog",
			Handler:    _MasterAdmin_CompactSysCatalog_Handler,
		},
		{
			MethodName: "IsInitDbDone",
			Handler:    _MasterAdmin_IsInitDbDone_Handler,
		},
		{
			MethodName: "SplitTablet",
			Handler:    _MasterAdmin_SplitTablet_Handler,
		},
		{
			MethodName: "CreateTransactionStatusTable",
			Handler:    _MasterAdmin_CreateTransactionStatusTable_Handler,
		},
		{
			MethodName: "DeleteNotServingTablet",
			Handler:    _MasterAdmin_DeleteNotServingTablet_Handler,
		},
		{
			MethodName: "DdlLog",
			Handler:    _MasterAdmin_DdlLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yb/master/master_admin.proto",
}