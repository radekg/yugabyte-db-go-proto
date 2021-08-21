// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
// The following only applies to changes made to this file as part of YugaByte development.
//
// Portions Copyright (c) YugaByte, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
// in compliance with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied.  See the License for the specific language governing permissions and limitations
// under the License.
//
// Protobufs used by both client-server and server-server traffic
// for user data transfer. This file should only contain protobufs
// which are exclusively used on the wire. If a protobuf is persisted on
// disk and not used as part of the wire protocol, it belongs in another
// place such as common/common.proto or within server/, etc.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: yb/common/wire_protocol.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AppStatusPB_ErrorCode int32

const (
	AppStatusPB_UNKNOWN_ERROR             AppStatusPB_ErrorCode = 999
	AppStatusPB_OK                        AppStatusPB_ErrorCode = 0
	AppStatusPB_NOT_FOUND                 AppStatusPB_ErrorCode = 1
	AppStatusPB_CORRUPTION                AppStatusPB_ErrorCode = 2
	AppStatusPB_NOT_SUPPORTED             AppStatusPB_ErrorCode = 3
	AppStatusPB_INVALID_ARGUMENT          AppStatusPB_ErrorCode = 4
	AppStatusPB_IO_ERROR                  AppStatusPB_ErrorCode = 5
	AppStatusPB_ALREADY_PRESENT           AppStatusPB_ErrorCode = 6
	AppStatusPB_RUNTIME_ERROR             AppStatusPB_ErrorCode = 7
	AppStatusPB_NETWORK_ERROR             AppStatusPB_ErrorCode = 8
	AppStatusPB_ILLEGAL_STATE             AppStatusPB_ErrorCode = 9
	AppStatusPB_NOT_AUTHORIZED            AppStatusPB_ErrorCode = 10
	AppStatusPB_ABORTED                   AppStatusPB_ErrorCode = 11
	AppStatusPB_REMOTE_ERROR              AppStatusPB_ErrorCode = 12
	AppStatusPB_SERVICE_UNAVAILABLE       AppStatusPB_ErrorCode = 13
	AppStatusPB_TIMED_OUT                 AppStatusPB_ErrorCode = 14
	AppStatusPB_UNINITIALIZED             AppStatusPB_ErrorCode = 15
	AppStatusPB_CONFIGURATION_ERROR       AppStatusPB_ErrorCode = 16
	AppStatusPB_INCOMPLETE                AppStatusPB_ErrorCode = 17
	AppStatusPB_END_OF_FILE               AppStatusPB_ErrorCode = 18
	AppStatusPB_INVALID_COMMAND           AppStatusPB_ErrorCode = 19
	AppStatusPB_QL_ERROR                  AppStatusPB_ErrorCode = 20
	AppStatusPB_INTERNAL_ERROR            AppStatusPB_ErrorCode = 21
	AppStatusPB_EXPIRED                   AppStatusPB_ErrorCode = 22
	AppStatusPB_LEADER_NOT_READY_TO_SERVE AppStatusPB_ErrorCode = 23
	AppStatusPB_LEADER_HAS_NO_LEASE       AppStatusPB_ErrorCode = 24
	AppStatusPB_TRY_AGAIN_CODE            AppStatusPB_ErrorCode = 25
	AppStatusPB_BUSY                      AppStatusPB_ErrorCode = 26
	AppStatusPB_SHUTDOWN_IN_PROGRESS      AppStatusPB_ErrorCode = 27
	AppStatusPB_MERGE_IN_PROGRESS         AppStatusPB_ErrorCode = 28
	AppStatusPB_COMBINED_ERROR            AppStatusPB_ErrorCode = 29
	AppStatusPB_SNAPSHOT_TOO_OLD          AppStatusPB_ErrorCode = 30
)

// Enum value maps for AppStatusPB_ErrorCode.
var (
	AppStatusPB_ErrorCode_name = map[int32]string{
		999: "UNKNOWN_ERROR",
		0:   "OK",
		1:   "NOT_FOUND",
		2:   "CORRUPTION",
		3:   "NOT_SUPPORTED",
		4:   "INVALID_ARGUMENT",
		5:   "IO_ERROR",
		6:   "ALREADY_PRESENT",
		7:   "RUNTIME_ERROR",
		8:   "NETWORK_ERROR",
		9:   "ILLEGAL_STATE",
		10:  "NOT_AUTHORIZED",
		11:  "ABORTED",
		12:  "REMOTE_ERROR",
		13:  "SERVICE_UNAVAILABLE",
		14:  "TIMED_OUT",
		15:  "UNINITIALIZED",
		16:  "CONFIGURATION_ERROR",
		17:  "INCOMPLETE",
		18:  "END_OF_FILE",
		19:  "INVALID_COMMAND",
		20:  "QL_ERROR",
		21:  "INTERNAL_ERROR",
		22:  "EXPIRED",
		23:  "LEADER_NOT_READY_TO_SERVE",
		24:  "LEADER_HAS_NO_LEASE",
		25:  "TRY_AGAIN_CODE",
		26:  "BUSY",
		27:  "SHUTDOWN_IN_PROGRESS",
		28:  "MERGE_IN_PROGRESS",
		29:  "COMBINED_ERROR",
		30:  "SNAPSHOT_TOO_OLD",
	}
	AppStatusPB_ErrorCode_value = map[string]int32{
		"UNKNOWN_ERROR":             999,
		"OK":                        0,
		"NOT_FOUND":                 1,
		"CORRUPTION":                2,
		"NOT_SUPPORTED":             3,
		"INVALID_ARGUMENT":          4,
		"IO_ERROR":                  5,
		"ALREADY_PRESENT":           6,
		"RUNTIME_ERROR":             7,
		"NETWORK_ERROR":             8,
		"ILLEGAL_STATE":             9,
		"NOT_AUTHORIZED":            10,
		"ABORTED":                   11,
		"REMOTE_ERROR":              12,
		"SERVICE_UNAVAILABLE":       13,
		"TIMED_OUT":                 14,
		"UNINITIALIZED":             15,
		"CONFIGURATION_ERROR":       16,
		"INCOMPLETE":                17,
		"END_OF_FILE":               18,
		"INVALID_COMMAND":           19,
		"QL_ERROR":                  20,
		"INTERNAL_ERROR":            21,
		"EXPIRED":                   22,
		"LEADER_NOT_READY_TO_SERVE": 23,
		"LEADER_HAS_NO_LEASE":       24,
		"TRY_AGAIN_CODE":            25,
		"BUSY":                      26,
		"SHUTDOWN_IN_PROGRESS":      27,
		"MERGE_IN_PROGRESS":         28,
		"COMBINED_ERROR":            29,
		"SNAPSHOT_TOO_OLD":          30,
	}
)

func (x AppStatusPB_ErrorCode) Enum() *AppStatusPB_ErrorCode {
	p := new(AppStatusPB_ErrorCode)
	*p = x
	return p
}

func (x AppStatusPB_ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppStatusPB_ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_yb_common_wire_protocol_proto_enumTypes[0].Descriptor()
}

func (AppStatusPB_ErrorCode) Type() protoreflect.EnumType {
	return &file_yb_common_wire_protocol_proto_enumTypes[0]
}

func (x AppStatusPB_ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *AppStatusPB_ErrorCode) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = AppStatusPB_ErrorCode(num)
	return nil
}

// Deprecated: Use AppStatusPB_ErrorCode.Descriptor instead.
func (AppStatusPB_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_yb_common_wire_protocol_proto_rawDescGZIP(), []int{0, 0}
}

// Error status returned by any RPC method.
// Every RPC method which could generate an application-level error
// should have this (or a more complex error result) as an optional field
// in its response.
//
// This maps to yb::Status in C++.
type AppStatusPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    *AppStatusPB_ErrorCode `protobuf:"varint,1,req,name=code,enum=yb.AppStatusPB_ErrorCode" json:"code,omitempty"`
	Message *string                `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	// For backward compatibility error_codes would also be filled.
	// But errors would have priority if both are set.
	//
	// Types that are assignable to ErrorCodes:
	//	*AppStatusPB_PosixCode
	//	*AppStatusPB_QlErrorCode
	ErrorCodes isAppStatusPB_ErrorCodes `protobuf_oneof:"error_codes"`
	SourceFile *string                  `protobuf:"bytes,6,opt,name=source_file,json=sourceFile" json:"source_file,omitempty"`
	SourceLine *int32                   `protobuf:"varint,7,opt,name=source_line,json=sourceLine" json:"source_line,omitempty"`
	// Stores encoded error codes of Status. See Status for details.
	Errors []byte `protobuf:"bytes,9,opt,name=errors" json:"errors,omitempty"`
}

func (x *AppStatusPB) Reset() {
	*x = AppStatusPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yb_common_wire_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppStatusPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppStatusPB) ProtoMessage() {}

func (x *AppStatusPB) ProtoReflect() protoreflect.Message {
	mi := &file_yb_common_wire_protocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppStatusPB.ProtoReflect.Descriptor instead.
func (*AppStatusPB) Descriptor() ([]byte, []int) {
	return file_yb_common_wire_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *AppStatusPB) GetCode() AppStatusPB_ErrorCode {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return AppStatusPB_UNKNOWN_ERROR
}

func (x *AppStatusPB) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (m *AppStatusPB) GetErrorCodes() isAppStatusPB_ErrorCodes {
	if m != nil {
		return m.ErrorCodes
	}
	return nil
}

func (x *AppStatusPB) GetPosixCode() int32 {
	if x, ok := x.GetErrorCodes().(*AppStatusPB_PosixCode); ok {
		return x.PosixCode
	}
	return 0
}

func (x *AppStatusPB) GetQlErrorCode() int64 {
	if x, ok := x.GetErrorCodes().(*AppStatusPB_QlErrorCode); ok {
		return x.QlErrorCode
	}
	return 0
}

func (x *AppStatusPB) GetSourceFile() string {
	if x != nil && x.SourceFile != nil {
		return *x.SourceFile
	}
	return ""
}

func (x *AppStatusPB) GetSourceLine() int32 {
	if x != nil && x.SourceLine != nil {
		return *x.SourceLine
	}
	return 0
}

func (x *AppStatusPB) GetErrors() []byte {
	if x != nil {
		return x.Errors
	}
	return nil
}

type isAppStatusPB_ErrorCodes interface {
	isAppStatusPB_ErrorCodes()
}

type AppStatusPB_PosixCode struct {
	PosixCode int32 `protobuf:"varint,4,opt,name=posix_code,json=posixCode,oneof"`
}

type AppStatusPB_QlErrorCode struct {
	QlErrorCode int64 `protobuf:"varint,5,opt,name=ql_error_code,json=qlErrorCode,oneof"`
}

func (*AppStatusPB_PosixCode) isAppStatusPB_ErrorCodes() {}

func (*AppStatusPB_QlErrorCode) isAppStatusPB_ErrorCodes() {}

// Uniquely identify a particular instance of a particular server in the
// cluster.
type NodeInstancePB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique ID which is created when the server is first started
	// up. This is stored persistently on disk.
	PermanentUuid []byte `protobuf:"bytes,1,req,name=permanent_uuid,json=permanentUuid" json:"permanent_uuid,omitempty"`
	// Sequence number incremented on every start-up of the server.
	// This makes it easy to detect when an instance has restarted (and
	// thus can be assumed to have forgotten any soft state it had in
	// memory).
	//
	// On a freshly initialized server, the first sequence number
	// should be 0.
	InstanceSeqno *int64 `protobuf:"varint,2,req,name=instance_seqno,json=instanceSeqno" json:"instance_seqno,omitempty"`
	// The time point when tserver/master starts up, it is derived from Env::NowMicros()
	StartTimeUs *uint64 `protobuf:"varint,3,opt,name=start_time_us,json=startTimeUs" json:"start_time_us,omitempty"`
}

func (x *NodeInstancePB) Reset() {
	*x = NodeInstancePB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yb_common_wire_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInstancePB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInstancePB) ProtoMessage() {}

func (x *NodeInstancePB) ProtoReflect() protoreflect.Message {
	mi := &file_yb_common_wire_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInstancePB.ProtoReflect.Descriptor instead.
func (*NodeInstancePB) Descriptor() ([]byte, []int) {
	return file_yb_common_wire_protocol_proto_rawDescGZIP(), []int{1}
}

func (x *NodeInstancePB) GetPermanentUuid() []byte {
	if x != nil {
		return x.PermanentUuid
	}
	return nil
}

func (x *NodeInstancePB) GetInstanceSeqno() int64 {
	if x != nil && x.InstanceSeqno != nil {
		return *x.InstanceSeqno
	}
	return 0
}

func (x *NodeInstancePB) GetStartTimeUs() uint64 {
	if x != nil && x.StartTimeUs != nil {
		return *x.StartTimeUs
	}
	return 0
}

// RPC and HTTP addresses for each server, as well as cloud related information.
type ServerRegistrationPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrivateRpcAddresses []*HostPortPB `protobuf:"bytes,1,rep,name=private_rpc_addresses,json=privateRpcAddresses" json:"private_rpc_addresses,omitempty"`
	BroadcastAddresses  []*HostPortPB `protobuf:"bytes,5,rep,name=broadcast_addresses,json=broadcastAddresses" json:"broadcast_addresses,omitempty"`
	HttpAddresses       []*HostPortPB `protobuf:"bytes,2,rep,name=http_addresses,json=httpAddresses" json:"http_addresses,omitempty"`
	PgPort              *uint32       `protobuf:"varint,6,opt,name=pg_port,json=pgPort,def=5433" json:"pg_port,omitempty"`
	CloudInfo           *CloudInfoPB  `protobuf:"bytes,3,opt,name=cloud_info,json=cloudInfo" json:"cloud_info,omitempty"`
	// Placement uuid of the tserver's cluster.
	PlacementUuid []byte `protobuf:"bytes,4,opt,name=placement_uuid,json=placementUuid" json:"placement_uuid,omitempty"`
}

// Default values for ServerRegistrationPB fields.
const (
	Default_ServerRegistrationPB_PgPort = uint32(5433)
)

func (x *ServerRegistrationPB) Reset() {
	*x = ServerRegistrationPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yb_common_wire_protocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerRegistrationPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerRegistrationPB) ProtoMessage() {}

func (x *ServerRegistrationPB) ProtoReflect() protoreflect.Message {
	mi := &file_yb_common_wire_protocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerRegistrationPB.ProtoReflect.Descriptor instead.
func (*ServerRegistrationPB) Descriptor() ([]byte, []int) {
	return file_yb_common_wire_protocol_proto_rawDescGZIP(), []int{2}
}

func (x *ServerRegistrationPB) GetPrivateRpcAddresses() []*HostPortPB {
	if x != nil {
		return x.PrivateRpcAddresses
	}
	return nil
}

func (x *ServerRegistrationPB) GetBroadcastAddresses() []*HostPortPB {
	if x != nil {
		return x.BroadcastAddresses
	}
	return nil
}

func (x *ServerRegistrationPB) GetHttpAddresses() []*HostPortPB {
	if x != nil {
		return x.HttpAddresses
	}
	return nil
}

func (x *ServerRegistrationPB) GetPgPort() uint32 {
	if x != nil && x.PgPort != nil {
		return *x.PgPort
	}
	return Default_ServerRegistrationPB_PgPort
}

func (x *ServerRegistrationPB) GetCloudInfo() *CloudInfoPB {
	if x != nil {
		return x.CloudInfo
	}
	return nil
}

func (x *ServerRegistrationPB) GetPlacementUuid() []byte {
	if x != nil {
		return x.PlacementUuid
	}
	return nil
}

type ServerEntryPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If there is an error communicating with the server (or retrieving
	// the server registration on the server itself), this field will be
	// set to contain the error.
	//
	// All subsequent fields are optional, as they may not be set if
	// an error is encountered communicating with the individual server.
	Error        *AppStatusPB          `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	InstanceId   *NodeInstancePB       `protobuf:"bytes,2,opt,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
	Registration *ServerRegistrationPB `protobuf:"bytes,3,opt,name=registration" json:"registration,omitempty"`
	// If an error has occurred earlier in the RPC call, the role
	// may be not be set.
	Role *RaftPeerPB_Role `protobuf:"varint,4,opt,name=role,enum=yb.consensus.RaftPeerPB_Role" json:"role,omitempty"`
}

func (x *ServerEntryPB) Reset() {
	*x = ServerEntryPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yb_common_wire_protocol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerEntryPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerEntryPB) ProtoMessage() {}

func (x *ServerEntryPB) ProtoReflect() protoreflect.Message {
	mi := &file_yb_common_wire_protocol_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerEntryPB.ProtoReflect.Descriptor instead.
func (*ServerEntryPB) Descriptor() ([]byte, []int) {
	return file_yb_common_wire_protocol_proto_rawDescGZIP(), []int{3}
}

func (x *ServerEntryPB) GetError() *AppStatusPB {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ServerEntryPB) GetInstanceId() *NodeInstancePB {
	if x != nil {
		return x.InstanceId
	}
	return nil
}

func (x *ServerEntryPB) GetRegistration() *ServerRegistrationPB {
	if x != nil {
		return x.Registration
	}
	return nil
}

func (x *ServerEntryPB) GetRole() RaftPeerPB_Role {
	if x != nil && x.Role != nil {
		return *x.Role
	}
	return RaftPeerPB_FOLLOWER
}

var File_yb_common_wire_protocol_proto protoreflect.FileDescriptor

var file_yb_common_wire_protocol_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x79, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x77, 0x69, 0x72, 0x65,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x79, 0x62, 0x1a, 0x16, 0x79, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x79, 0x62, 0x2f,
	0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x73, 0x75, 0x73, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf1, 0x06, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x50, 0x42, 0x12, 0x2d, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e,
	0x32, 0x19, 0x2e, 0x79, 0x62, 0x2e, 0x41, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50,
	0x42, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x70,
	0x6f, 0x73, 0x69, 0x78, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x00, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x78, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x71, 0x6c, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0b, 0x71, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xe8, 0x04, 0x0a,
	0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x0d, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xe7, 0x07, 0x12, 0x06,
	0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x52, 0x52, 0x55, 0x50, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50,
	0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x04, 0x12, 0x0c,
	0x0a, 0x08, 0x49, 0x4f, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f,
	0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x50, 0x52, 0x45, 0x53, 0x45, 0x4e, 0x54, 0x10,
	0x06, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x55, 0x4e, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x07, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x08, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x4c, 0x4c, 0x45, 0x47,
	0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0x09, 0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x4f,
	0x54, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x0b,
	0x0a, 0x07, 0x41, 0x42, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x52,
	0x45, 0x4d, 0x4f, 0x54, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x0c, 0x12, 0x17, 0x0a,
	0x13, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c,
	0x41, 0x42, 0x4c, 0x45, 0x10, 0x0d, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x49, 0x4d, 0x45, 0x44, 0x5f,
	0x4f, 0x55, 0x54, 0x10, 0x0e, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x49, 0x4e, 0x49, 0x54, 0x49,
	0x41, 0x4c, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x0f, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4e, 0x46,
	0x49, 0x47, 0x55, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x10, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10,
	0x11, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x4e, 0x44, 0x5f, 0x4f, 0x46, 0x5f, 0x46, 0x49, 0x4c, 0x45,
	0x10, 0x12, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x4f,
	0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x10, 0x13, 0x12, 0x0c, 0x0a, 0x08, 0x51, 0x4c, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x14, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41,
	0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x15, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58, 0x50,
	0x49, 0x52, 0x45, 0x44, 0x10, 0x16, 0x12, 0x1d, 0x0a, 0x19, 0x4c, 0x45, 0x41, 0x44, 0x45, 0x52,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x54, 0x4f, 0x5f, 0x53, 0x45,
	0x52, 0x56, 0x45, 0x10, 0x17, 0x12, 0x17, 0x0a, 0x13, 0x4c, 0x45, 0x41, 0x44, 0x45, 0x52, 0x5f,
	0x48, 0x41, 0x53, 0x5f, 0x4e, 0x4f, 0x5f, 0x4c, 0x45, 0x41, 0x53, 0x45, 0x10, 0x18, 0x12, 0x12,
	0x0a, 0x0e, 0x54, 0x52, 0x59, 0x5f, 0x41, 0x47, 0x41, 0x49, 0x4e, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x10, 0x19, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x55, 0x53, 0x59, 0x10, 0x1a, 0x12, 0x18, 0x0a, 0x14,
	0x53, 0x48, 0x55, 0x54, 0x44, 0x4f, 0x57, 0x4e, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47,
	0x52, 0x45, 0x53, 0x53, 0x10, 0x1b, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x45, 0x52, 0x47, 0x45, 0x5f,
	0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x1c, 0x12, 0x12, 0x0a,
	0x0e, 0x43, 0x4f, 0x4d, 0x42, 0x49, 0x4e, 0x45, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x1d, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x4e, 0x41, 0x50, 0x53, 0x48, 0x4f, 0x54, 0x5f, 0x54, 0x4f,
	0x4f, 0x5f, 0x4f, 0x4c, 0x44, 0x10, 0x1e, 0x42, 0x0d, 0x0a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x0e, 0x4e, 0x6f, 0x64, 0x65, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x42, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x65, 0x72,
	0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x0c, 0x52, 0x0d, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x55, 0x75, 0x69, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x71,
	0x6e, 0x6f, 0x18, 0x02, 0x20, 0x02, 0x28, 0x03, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x53, 0x65, 0x71, 0x6e, 0x6f, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x22, 0xc8, 0x02, 0x0a, 0x14,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x42, 0x12, 0x42, 0x0a, 0x15, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f,
	0x72, 0x70, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x79, 0x62, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x72,
	0x74, 0x50, 0x42, 0x52, 0x13, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x70, 0x63, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x13, 0x62, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x79, 0x62, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x50,
	0x6f, 0x72, 0x74, 0x50, 0x42, 0x52, 0x12, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x0e, 0x68, 0x74, 0x74,
	0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x79, 0x62, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x50,
	0x42, 0x52, 0x0d, 0x68, 0x74, 0x74, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x12, 0x1d, 0x0a, 0x07, 0x70, 0x67, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x3a, 0x04, 0x35, 0x34, 0x33, 0x33, 0x52, 0x06, 0x70, 0x67, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x2e, 0x0a, 0x0a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x79, 0x62, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x50, 0x42, 0x52, 0x09, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x25, 0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x55, 0x75, 0x69, 0x64, 0x22, 0xdc, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x42, 0x12, 0x25, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x79, 0x62, 0x2e, 0x41, 0x70, 0x70,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x42, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x33, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x79, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x42, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x79, 0x62, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x42, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1d, 0x2e, 0x79, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e,
	0x52, 0x61, 0x66, 0x74, 0x50, 0x65, 0x65, 0x72, 0x50, 0x42, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x2e, 0x79, 0x62,
}

var (
	file_yb_common_wire_protocol_proto_rawDescOnce sync.Once
	file_yb_common_wire_protocol_proto_rawDescData = file_yb_common_wire_protocol_proto_rawDesc
)

func file_yb_common_wire_protocol_proto_rawDescGZIP() []byte {
	file_yb_common_wire_protocol_proto_rawDescOnce.Do(func() {
		file_yb_common_wire_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_yb_common_wire_protocol_proto_rawDescData)
	})
	return file_yb_common_wire_protocol_proto_rawDescData
}

var file_yb_common_wire_protocol_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yb_common_wire_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_yb_common_wire_protocol_proto_goTypes = []interface{}{
	(AppStatusPB_ErrorCode)(0),   // 0: yb.AppStatusPB.ErrorCode
	(*AppStatusPB)(nil),          // 1: yb.AppStatusPB
	(*NodeInstancePB)(nil),       // 2: yb.NodeInstancePB
	(*ServerRegistrationPB)(nil), // 3: yb.ServerRegistrationPB
	(*ServerEntryPB)(nil),        // 4: yb.ServerEntryPB
	(*HostPortPB)(nil),           // 5: yb.HostPortPB
	(*CloudInfoPB)(nil),          // 6: yb.CloudInfoPB
	(RaftPeerPB_Role)(0),         // 7: yb.consensus.RaftPeerPB.Role
}
var file_yb_common_wire_protocol_proto_depIdxs = []int32{
	0, // 0: yb.AppStatusPB.code:type_name -> yb.AppStatusPB.ErrorCode
	5, // 1: yb.ServerRegistrationPB.private_rpc_addresses:type_name -> yb.HostPortPB
	5, // 2: yb.ServerRegistrationPB.broadcast_addresses:type_name -> yb.HostPortPB
	5, // 3: yb.ServerRegistrationPB.http_addresses:type_name -> yb.HostPortPB
	6, // 4: yb.ServerRegistrationPB.cloud_info:type_name -> yb.CloudInfoPB
	1, // 5: yb.ServerEntryPB.error:type_name -> yb.AppStatusPB
	2, // 6: yb.ServerEntryPB.instance_id:type_name -> yb.NodeInstancePB
	3, // 7: yb.ServerEntryPB.registration:type_name -> yb.ServerRegistrationPB
	7, // 8: yb.ServerEntryPB.role:type_name -> yb.consensus.RaftPeerPB.Role
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_yb_common_wire_protocol_proto_init() }
func file_yb_common_wire_protocol_proto_init() {
	if File_yb_common_wire_protocol_proto != nil {
		return
	}
	file_yb_common_common_proto_init()
	file_yb_consensus_consensus_metadata_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yb_common_wire_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppStatusPB); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yb_common_wire_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInstancePB); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yb_common_wire_protocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerRegistrationPB); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yb_common_wire_protocol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerEntryPB); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_yb_common_wire_protocol_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AppStatusPB_PosixCode)(nil),
		(*AppStatusPB_QlErrorCode)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yb_common_wire_protocol_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yb_common_wire_protocol_proto_goTypes,
		DependencyIndexes: file_yb_common_wire_protocol_proto_depIdxs,
		EnumInfos:         file_yb_common_wire_protocol_proto_enumTypes,
		MessageInfos:      file_yb_common_wire_protocol_proto_msgTypes,
	}.Build()
	File_yb_common_wire_protocol_proto = out.File
	file_yb_common_wire_protocol_proto_rawDesc = nil
	file_yb_common_wire_protocol_proto_goTypes = nil
	file_yb_common_wire_protocol_proto_depIdxs = nil
}