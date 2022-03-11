// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: yb/master/master_dcl.proto

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

// MasterDclClient is the client API for MasterDcl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterDclClient interface {
	//  Authentication and Authorization.
	CreateRole(ctx context.Context, in *CreateRoleRequestPB, opts ...grpc.CallOption) (*CreateRoleResponsePB, error)
	AlterRole(ctx context.Context, in *AlterRoleRequestPB, opts ...grpc.CallOption) (*AlterRoleResponsePB, error)
	DeleteRole(ctx context.Context, in *DeleteRoleRequestPB, opts ...grpc.CallOption) (*DeleteRoleResponsePB, error)
	GrantRevokeRole(ctx context.Context, in *GrantRevokeRoleRequestPB, opts ...grpc.CallOption) (*GrantRevokeRoleResponsePB, error)
	GrantRevokePermission(ctx context.Context, in *GrantRevokePermissionRequestPB, opts ...grpc.CallOption) (*GrantRevokePermissionResponsePB, error)
	GetPermissions(ctx context.Context, in *GetPermissionsRequestPB, opts ...grpc.CallOption) (*GetPermissionsResponsePB, error)
}

type masterDclClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterDclClient(cc grpc.ClientConnInterface) MasterDclClient {
	return &masterDclClient{cc}
}

func (c *masterDclClient) CreateRole(ctx context.Context, in *CreateRoleRequestPB, opts ...grpc.CallOption) (*CreateRoleResponsePB, error) {
	out := new(CreateRoleResponsePB)
	err := c.cc.Invoke(ctx, "/yb.master.MasterDcl/CreateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterDclClient) AlterRole(ctx context.Context, in *AlterRoleRequestPB, opts ...grpc.CallOption) (*AlterRoleResponsePB, error) {
	out := new(AlterRoleResponsePB)
	err := c.cc.Invoke(ctx, "/yb.master.MasterDcl/AlterRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterDclClient) DeleteRole(ctx context.Context, in *DeleteRoleRequestPB, opts ...grpc.CallOption) (*DeleteRoleResponsePB, error) {
	out := new(DeleteRoleResponsePB)
	err := c.cc.Invoke(ctx, "/yb.master.MasterDcl/DeleteRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterDclClient) GrantRevokeRole(ctx context.Context, in *GrantRevokeRoleRequestPB, opts ...grpc.CallOption) (*GrantRevokeRoleResponsePB, error) {
	out := new(GrantRevokeRoleResponsePB)
	err := c.cc.Invoke(ctx, "/yb.master.MasterDcl/GrantRevokeRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterDclClient) GrantRevokePermission(ctx context.Context, in *GrantRevokePermissionRequestPB, opts ...grpc.CallOption) (*GrantRevokePermissionResponsePB, error) {
	out := new(GrantRevokePermissionResponsePB)
	err := c.cc.Invoke(ctx, "/yb.master.MasterDcl/GrantRevokePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterDclClient) GetPermissions(ctx context.Context, in *GetPermissionsRequestPB, opts ...grpc.CallOption) (*GetPermissionsResponsePB, error) {
	out := new(GetPermissionsResponsePB)
	err := c.cc.Invoke(ctx, "/yb.master.MasterDcl/GetPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterDclServer is the server API for MasterDcl service.
// All implementations should embed UnimplementedMasterDclServer
// for forward compatibility
type MasterDclServer interface {
	//  Authentication and Authorization.
	CreateRole(context.Context, *CreateRoleRequestPB) (*CreateRoleResponsePB, error)
	AlterRole(context.Context, *AlterRoleRequestPB) (*AlterRoleResponsePB, error)
	DeleteRole(context.Context, *DeleteRoleRequestPB) (*DeleteRoleResponsePB, error)
	GrantRevokeRole(context.Context, *GrantRevokeRoleRequestPB) (*GrantRevokeRoleResponsePB, error)
	GrantRevokePermission(context.Context, *GrantRevokePermissionRequestPB) (*GrantRevokePermissionResponsePB, error)
	GetPermissions(context.Context, *GetPermissionsRequestPB) (*GetPermissionsResponsePB, error)
}

// UnimplementedMasterDclServer should be embedded to have forward compatible implementations.
type UnimplementedMasterDclServer struct {
}

func (UnimplementedMasterDclServer) CreateRole(context.Context, *CreateRoleRequestPB) (*CreateRoleResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedMasterDclServer) AlterRole(context.Context, *AlterRoleRequestPB) (*AlterRoleResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlterRole not implemented")
}
func (UnimplementedMasterDclServer) DeleteRole(context.Context, *DeleteRoleRequestPB) (*DeleteRoleResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}
func (UnimplementedMasterDclServer) GrantRevokeRole(context.Context, *GrantRevokeRoleRequestPB) (*GrantRevokeRoleResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrantRevokeRole not implemented")
}
func (UnimplementedMasterDclServer) GrantRevokePermission(context.Context, *GrantRevokePermissionRequestPB) (*GrantRevokePermissionResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrantRevokePermission not implemented")
}
func (UnimplementedMasterDclServer) GetPermissions(context.Context, *GetPermissionsRequestPB) (*GetPermissionsResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissions not implemented")
}

// UnsafeMasterDclServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterDclServer will
// result in compilation errors.
type UnsafeMasterDclServer interface {
	mustEmbedUnimplementedMasterDclServer()
}

func RegisterMasterDclServer(s grpc.ServiceRegistrar, srv MasterDclServer) {
	s.RegisterService(&MasterDcl_ServiceDesc, srv)
}

func _MasterDcl_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterDclServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.master.MasterDcl/CreateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterDclServer).CreateRole(ctx, req.(*CreateRoleRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterDcl_AlterRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlterRoleRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterDclServer).AlterRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.master.MasterDcl/AlterRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterDclServer).AlterRole(ctx, req.(*AlterRoleRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterDcl_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterDclServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.master.MasterDcl/DeleteRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterDclServer).DeleteRole(ctx, req.(*DeleteRoleRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterDcl_GrantRevokeRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrantRevokeRoleRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterDclServer).GrantRevokeRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.master.MasterDcl/GrantRevokeRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterDclServer).GrantRevokeRole(ctx, req.(*GrantRevokeRoleRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterDcl_GrantRevokePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrantRevokePermissionRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterDclServer).GrantRevokePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.master.MasterDcl/GrantRevokePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterDclServer).GrantRevokePermission(ctx, req.(*GrantRevokePermissionRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterDcl_GetPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPermissionsRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterDclServer).GetPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.master.MasterDcl/GetPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterDclServer).GetPermissions(ctx, req.(*GetPermissionsRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

// MasterDcl_ServiceDesc is the grpc.ServiceDesc for MasterDcl service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MasterDcl_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yb.master.MasterDcl",
	HandlerType: (*MasterDclServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRole",
			Handler:    _MasterDcl_CreateRole_Handler,
		},
		{
			MethodName: "AlterRole",
			Handler:    _MasterDcl_AlterRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _MasterDcl_DeleteRole_Handler,
		},
		{
			MethodName: "GrantRevokeRole",
			Handler:    _MasterDcl_GrantRevokeRole_Handler,
		},
		{
			MethodName: "GrantRevokePermission",
			Handler:    _MasterDcl_GrantRevokePermission_Handler,
		},
		{
			MethodName: "GetPermissions",
			Handler:    _MasterDcl_GetPermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yb/master/master_dcl.proto",
}