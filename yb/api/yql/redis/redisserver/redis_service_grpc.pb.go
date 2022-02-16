// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: yb/yql/redis/redisserver/redis_service.proto

package redisserver

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RedisServerServiceClient is the client API for RedisServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RedisServerServiceClient interface {
}

type redisServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRedisServerServiceClient(cc grpc.ClientConnInterface) RedisServerServiceClient {
	return &redisServerServiceClient{cc}
}

// RedisServerServiceServer is the server API for RedisServerService service.
// All implementations should embed UnimplementedRedisServerServiceServer
// for forward compatibility
type RedisServerServiceServer interface {
}

// UnimplementedRedisServerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRedisServerServiceServer struct {
}

// UnsafeRedisServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RedisServerServiceServer will
// result in compilation errors.
type UnsafeRedisServerServiceServer interface {
	mustEmbedUnimplementedRedisServerServiceServer()
}

func RegisterRedisServerServiceServer(s grpc.ServiceRegistrar, srv RedisServerServiceServer) {
	s.RegisterService(&RedisServerService_ServiceDesc, srv)
}

// RedisServerService_ServiceDesc is the grpc.ServiceDesc for RedisServerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RedisServerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yb.redisserver.RedisServerService",
	HandlerType: (*RedisServerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "yb/yql/redis/redisserver/redis_service.proto",
}
