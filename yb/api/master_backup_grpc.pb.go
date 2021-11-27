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

// MasterBackupServiceClient is the client API for MasterBackupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterBackupServiceClient interface {
	// Client->Master RPCs
	CreateSnapshot(ctx context.Context, in *CreateSnapshotRequestPB, opts ...grpc.CallOption) (*CreateSnapshotResponsePB, error)
	ListSnapshots(ctx context.Context, in *ListSnapshotsRequestPB, opts ...grpc.CallOption) (*ListSnapshotsResponsePB, error)
	ListSnapshotRestorations(ctx context.Context, in *ListSnapshotRestorationsRequestPB, opts ...grpc.CallOption) (*ListSnapshotRestorationsResponsePB, error)
	RestoreSnapshot(ctx context.Context, in *RestoreSnapshotRequestPB, opts ...grpc.CallOption) (*RestoreSnapshotResponsePB, error)
	DeleteSnapshot(ctx context.Context, in *DeleteSnapshotRequestPB, opts ...grpc.CallOption) (*DeleteSnapshotResponsePB, error)
	ImportSnapshotMeta(ctx context.Context, in *ImportSnapshotMetaRequestPB, opts ...grpc.CallOption) (*ImportSnapshotMetaResponsePB, error)
	CreateSnapshotSchedule(ctx context.Context, in *CreateSnapshotScheduleRequestPB, opts ...grpc.CallOption) (*CreateSnapshotScheduleResponsePB, error)
	ListSnapshotSchedules(ctx context.Context, in *ListSnapshotSchedulesRequestPB, opts ...grpc.CallOption) (*ListSnapshotSchedulesResponsePB, error)
	DeleteSnapshotSchedule(ctx context.Context, in *DeleteSnapshotScheduleRequestPB, opts ...grpc.CallOption) (*DeleteSnapshotScheduleResponsePB, error)
}

type masterBackupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterBackupServiceClient(cc grpc.ClientConnInterface) MasterBackupServiceClient {
	return &masterBackupServiceClient{cc}
}

func (c *masterBackupServiceClient) CreateSnapshot(ctx context.Context, in *CreateSnapshotRequestPB, opts ...grpc.CallOption) (*CreateSnapshotResponsePB, error) {
	out := new(CreateSnapshotResponsePB)
	err := c.cc.Invoke(ctx, "/yb.master.MasterBackupService/CreateSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterBackupServiceClient) ListSnapshots(ctx context.Context, in *ListSnapshotsRequestPB, opts ...grpc.CallOption) (*ListSnapshotsResponsePB, error) {
	out := new(ListSnapshotsResponsePB)
	err := c.cc.Invoke(ctx, "/yb.master.MasterBackupService/ListSnapshots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterBackupServiceClient) ListSnapshotRestorations(ctx context.Context, in *ListSnapshotRestorationsRequestPB, opts ...grpc.CallOption) (*ListSnapshotRestorationsResponsePB, error) {
	out := new(ListSnapshotRestorationsResponsePB)
	err := c.cc.Invoke(ctx, "/yb.master.MasterBackupService/ListSnapshotRestorations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterBackupServiceClient) RestoreSnapshot(ctx context.Context, in *RestoreSnapshotRequestPB, opts ...grpc.CallOption) (*RestoreSnapshotResponsePB, error) {
	out := new(RestoreSnapshotResponsePB)
	err := c.cc.Invoke(ctx, "/yb.master.MasterBackupService/RestoreSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterBackupServiceClient) DeleteSnapshot(ctx context.Context, in *DeleteSnapshotRequestPB, opts ...grpc.CallOption) (*DeleteSnapshotResponsePB, error) {
	out := new(DeleteSnapshotResponsePB)
	err := c.cc.Invoke(ctx, "/yb.master.MasterBackupService/DeleteSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterBackupServiceClient) ImportSnapshotMeta(ctx context.Context, in *ImportSnapshotMetaRequestPB, opts ...grpc.CallOption) (*ImportSnapshotMetaResponsePB, error) {
	out := new(ImportSnapshotMetaResponsePB)
	err := c.cc.Invoke(ctx, "/yb.master.MasterBackupService/ImportSnapshotMeta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterBackupServiceClient) CreateSnapshotSchedule(ctx context.Context, in *CreateSnapshotScheduleRequestPB, opts ...grpc.CallOption) (*CreateSnapshotScheduleResponsePB, error) {
	out := new(CreateSnapshotScheduleResponsePB)
	err := c.cc.Invoke(ctx, "/yb.master.MasterBackupService/CreateSnapshotSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterBackupServiceClient) ListSnapshotSchedules(ctx context.Context, in *ListSnapshotSchedulesRequestPB, opts ...grpc.CallOption) (*ListSnapshotSchedulesResponsePB, error) {
	out := new(ListSnapshotSchedulesResponsePB)
	err := c.cc.Invoke(ctx, "/yb.master.MasterBackupService/ListSnapshotSchedules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterBackupServiceClient) DeleteSnapshotSchedule(ctx context.Context, in *DeleteSnapshotScheduleRequestPB, opts ...grpc.CallOption) (*DeleteSnapshotScheduleResponsePB, error) {
	out := new(DeleteSnapshotScheduleResponsePB)
	err := c.cc.Invoke(ctx, "/yb.master.MasterBackupService/DeleteSnapshotSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterBackupServiceServer is the server API for MasterBackupService service.
// All implementations should embed UnimplementedMasterBackupServiceServer
// for forward compatibility
type MasterBackupServiceServer interface {
	// Client->Master RPCs
	CreateSnapshot(context.Context, *CreateSnapshotRequestPB) (*CreateSnapshotResponsePB, error)
	ListSnapshots(context.Context, *ListSnapshotsRequestPB) (*ListSnapshotsResponsePB, error)
	ListSnapshotRestorations(context.Context, *ListSnapshotRestorationsRequestPB) (*ListSnapshotRestorationsResponsePB, error)
	RestoreSnapshot(context.Context, *RestoreSnapshotRequestPB) (*RestoreSnapshotResponsePB, error)
	DeleteSnapshot(context.Context, *DeleteSnapshotRequestPB) (*DeleteSnapshotResponsePB, error)
	ImportSnapshotMeta(context.Context, *ImportSnapshotMetaRequestPB) (*ImportSnapshotMetaResponsePB, error)
	CreateSnapshotSchedule(context.Context, *CreateSnapshotScheduleRequestPB) (*CreateSnapshotScheduleResponsePB, error)
	ListSnapshotSchedules(context.Context, *ListSnapshotSchedulesRequestPB) (*ListSnapshotSchedulesResponsePB, error)
	DeleteSnapshotSchedule(context.Context, *DeleteSnapshotScheduleRequestPB) (*DeleteSnapshotScheduleResponsePB, error)
}

// UnimplementedMasterBackupServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMasterBackupServiceServer struct {
}

func (UnimplementedMasterBackupServiceServer) CreateSnapshot(context.Context, *CreateSnapshotRequestPB) (*CreateSnapshotResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSnapshot not implemented")
}
func (UnimplementedMasterBackupServiceServer) ListSnapshots(context.Context, *ListSnapshotsRequestPB) (*ListSnapshotsResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSnapshots not implemented")
}
func (UnimplementedMasterBackupServiceServer) ListSnapshotRestorations(context.Context, *ListSnapshotRestorationsRequestPB) (*ListSnapshotRestorationsResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSnapshotRestorations not implemented")
}
func (UnimplementedMasterBackupServiceServer) RestoreSnapshot(context.Context, *RestoreSnapshotRequestPB) (*RestoreSnapshotResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestoreSnapshot not implemented")
}
func (UnimplementedMasterBackupServiceServer) DeleteSnapshot(context.Context, *DeleteSnapshotRequestPB) (*DeleteSnapshotResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSnapshot not implemented")
}
func (UnimplementedMasterBackupServiceServer) ImportSnapshotMeta(context.Context, *ImportSnapshotMetaRequestPB) (*ImportSnapshotMetaResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportSnapshotMeta not implemented")
}
func (UnimplementedMasterBackupServiceServer) CreateSnapshotSchedule(context.Context, *CreateSnapshotScheduleRequestPB) (*CreateSnapshotScheduleResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSnapshotSchedule not implemented")
}
func (UnimplementedMasterBackupServiceServer) ListSnapshotSchedules(context.Context, *ListSnapshotSchedulesRequestPB) (*ListSnapshotSchedulesResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSnapshotSchedules not implemented")
}
func (UnimplementedMasterBackupServiceServer) DeleteSnapshotSchedule(context.Context, *DeleteSnapshotScheduleRequestPB) (*DeleteSnapshotScheduleResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSnapshotSchedule not implemented")
}

// UnsafeMasterBackupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterBackupServiceServer will
// result in compilation errors.
type UnsafeMasterBackupServiceServer interface {
	mustEmbedUnimplementedMasterBackupServiceServer()
}

func RegisterMasterBackupServiceServer(s grpc.ServiceRegistrar, srv MasterBackupServiceServer) {
	s.RegisterService(&MasterBackupService_ServiceDesc, srv)
}

func _MasterBackupService_CreateSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSnapshotRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterBackupServiceServer).CreateSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.master.MasterBackupService/CreateSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterBackupServiceServer).CreateSnapshot(ctx, req.(*CreateSnapshotRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterBackupService_ListSnapshots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSnapshotsRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterBackupServiceServer).ListSnapshots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.master.MasterBackupService/ListSnapshots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterBackupServiceServer).ListSnapshots(ctx, req.(*ListSnapshotsRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterBackupService_ListSnapshotRestorations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSnapshotRestorationsRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterBackupServiceServer).ListSnapshotRestorations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.master.MasterBackupService/ListSnapshotRestorations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterBackupServiceServer).ListSnapshotRestorations(ctx, req.(*ListSnapshotRestorationsRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterBackupService_RestoreSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreSnapshotRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterBackupServiceServer).RestoreSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.master.MasterBackupService/RestoreSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterBackupServiceServer).RestoreSnapshot(ctx, req.(*RestoreSnapshotRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterBackupService_DeleteSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSnapshotRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterBackupServiceServer).DeleteSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.master.MasterBackupService/DeleteSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterBackupServiceServer).DeleteSnapshot(ctx, req.(*DeleteSnapshotRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterBackupService_ImportSnapshotMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportSnapshotMetaRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterBackupServiceServer).ImportSnapshotMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.master.MasterBackupService/ImportSnapshotMeta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterBackupServiceServer).ImportSnapshotMeta(ctx, req.(*ImportSnapshotMetaRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterBackupService_CreateSnapshotSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSnapshotScheduleRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterBackupServiceServer).CreateSnapshotSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.master.MasterBackupService/CreateSnapshotSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterBackupServiceServer).CreateSnapshotSchedule(ctx, req.(*CreateSnapshotScheduleRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterBackupService_ListSnapshotSchedules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSnapshotSchedulesRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterBackupServiceServer).ListSnapshotSchedules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.master.MasterBackupService/ListSnapshotSchedules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterBackupServiceServer).ListSnapshotSchedules(ctx, req.(*ListSnapshotSchedulesRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterBackupService_DeleteSnapshotSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSnapshotScheduleRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterBackupServiceServer).DeleteSnapshotSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.master.MasterBackupService/DeleteSnapshotSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterBackupServiceServer).DeleteSnapshotSchedule(ctx, req.(*DeleteSnapshotScheduleRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

// MasterBackupService_ServiceDesc is the grpc.ServiceDesc for MasterBackupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MasterBackupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yb.master.MasterBackupService",
	HandlerType: (*MasterBackupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSnapshot",
			Handler:    _MasterBackupService_CreateSnapshot_Handler,
		},
		{
			MethodName: "ListSnapshots",
			Handler:    _MasterBackupService_ListSnapshots_Handler,
		},
		{
			MethodName: "ListSnapshotRestorations",
			Handler:    _MasterBackupService_ListSnapshotRestorations_Handler,
		},
		{
			MethodName: "RestoreSnapshot",
			Handler:    _MasterBackupService_RestoreSnapshot_Handler,
		},
		{
			MethodName: "DeleteSnapshot",
			Handler:    _MasterBackupService_DeleteSnapshot_Handler,
		},
		{
			MethodName: "ImportSnapshotMeta",
			Handler:    _MasterBackupService_ImportSnapshotMeta_Handler,
		},
		{
			MethodName: "CreateSnapshotSchedule",
			Handler:    _MasterBackupService_CreateSnapshotSchedule_Handler,
		},
		{
			MethodName: "ListSnapshotSchedules",
			Handler:    _MasterBackupService_ListSnapshotSchedules_Handler,
		},
		{
			MethodName: "DeleteSnapshotSchedule",
			Handler:    _MasterBackupService_DeleteSnapshotSchedule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yb/master/master_backup.proto",
}
