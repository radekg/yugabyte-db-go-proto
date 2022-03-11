// Copyright (c) YugaByte, Inc.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: yb/tserver/tserver_types.proto

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

// Enum of the server's Tablet Manager state: currently this is only
// used for assertions, but this can also be sent to the master.
type TSTabletManagerStatePB int32

const (
	TSTabletManagerStatePB_UNKNOWN TSTabletManagerStatePB = 999
	// Indicates that Tablet Manager is initializing.
	TSTabletManagerStatePB_MANAGER_INITIALIZING TSTabletManagerStatePB = 0
	// Indicates that Tablet Manager is running and can create new
	// tablets.
	TSTabletManagerStatePB_MANAGER_RUNNING TSTabletManagerStatePB = 1
	// Indicates that tablet manager is shutting down and no new tablets
	// can be created.
	TSTabletManagerStatePB_MANAGER_QUIESCING TSTabletManagerStatePB = 2
	// Tablet Manager has shutdown.
	TSTabletManagerStatePB_MANAGER_SHUTDOWN TSTabletManagerStatePB = 3
)

// Enum value maps for TSTabletManagerStatePB.
var (
	TSTabletManagerStatePB_name = map[int32]string{
		999: "UNKNOWN",
		0:   "MANAGER_INITIALIZING",
		1:   "MANAGER_RUNNING",
		2:   "MANAGER_QUIESCING",
		3:   "MANAGER_SHUTDOWN",
	}
	TSTabletManagerStatePB_value = map[string]int32{
		"UNKNOWN":              999,
		"MANAGER_INITIALIZING": 0,
		"MANAGER_RUNNING":      1,
		"MANAGER_QUIESCING":    2,
		"MANAGER_SHUTDOWN":     3,
	}
)

func (x TSTabletManagerStatePB) Enum() *TSTabletManagerStatePB {
	p := new(TSTabletManagerStatePB)
	*p = x
	return p
}

func (x TSTabletManagerStatePB) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TSTabletManagerStatePB) Descriptor() protoreflect.EnumDescriptor {
	return file_yb_tserver_tserver_types_proto_enumTypes[0].Descriptor()
}

func (TSTabletManagerStatePB) Type() protoreflect.EnumType {
	return &file_yb_tserver_tserver_types_proto_enumTypes[0]
}

func (x TSTabletManagerStatePB) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *TSTabletManagerStatePB) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = TSTabletManagerStatePB(num)
	return nil
}

// Deprecated: Use TSTabletManagerStatePB.Descriptor instead.
func (TSTabletManagerStatePB) EnumDescriptor() ([]byte, []int) {
	return file_yb_tserver_tserver_types_proto_rawDescGZIP(), []int{0}
}

type TabletServerErrorPB_Code int32

const (
	// An error which has no more specific error code.
	// The code and message in 'status' may reveal more details.
	//
	// RPCs should avoid returning this, since callers will not be
	// able to easily parse the error.
	TabletServerErrorPB_UNKNOWN_ERROR TabletServerErrorPB_Code = 1
	// The schema provided for a request was not well-formed.
	TabletServerErrorPB_INVALID_SCHEMA TabletServerErrorPB_Code = 2
	// The row data provided for a request was not well-formed.
	TabletServerErrorPB_INVALID_ROW_BLOCK TabletServerErrorPB_Code = 3
	// The mutations or mutation keys provided for a request were
	// not well formed.
	TabletServerErrorPB_INVALID_MUTATION TabletServerErrorPB_Code = 4
	// The schema provided for a request didn't match the actual
	// schema of the tablet.
	TabletServerErrorPB_MISMATCHED_SCHEMA TabletServerErrorPB_Code = 5
	// The requested tablet_id is not currently hosted on this server.
	TabletServerErrorPB_TABLET_NOT_FOUND TabletServerErrorPB_Code = 6
	// A request was made against a scanner ID that was either never
	// created or has expired.
	TabletServerErrorPB_SCANNER_EXPIRED TabletServerErrorPB_Code = 7
	// An invalid scan was specified -- e.g the values passed for
	// predicates were incorrect sizes.
	TabletServerErrorPB_INVALID_SCAN_SPEC TabletServerErrorPB_Code = 8
	// The provided configuration was not well-formed and/or
	// had a sequence number that was below the current config.
	TabletServerErrorPB_INVALID_CONFIG TabletServerErrorPB_Code = 9
	// On a create tablet request, signals that the tablet already exists.
	TabletServerErrorPB_TABLET_ALREADY_EXISTS TabletServerErrorPB_Code = 10
	// If the tablet has a newer schema than the requested one the "alter"
	// request will be rejected with this error.
	TabletServerErrorPB_TABLET_HAS_A_NEWER_SCHEMA TabletServerErrorPB_Code = 11
	// The tablet is hosted on this server, but not in RUNNING state.
	TabletServerErrorPB_TABLET_NOT_RUNNING TabletServerErrorPB_Code = 12
	// Client requested a snapshot read but the snapshot was invalid.
	TabletServerErrorPB_INVALID_SNAPSHOT TabletServerErrorPB_Code = 13
	// An invalid scan call sequence ID was specified.
	TabletServerErrorPB_INVALID_SCAN_CALL_SEQ_ID TabletServerErrorPB_Code = 14
	// This tserver is not the leader of the consensus configuration.
	TabletServerErrorPB_NOT_THE_LEADER TabletServerErrorPB_Code = 15
	// The destination UUID in the request does not match this server.
	TabletServerErrorPB_WRONG_SERVER_UUID TabletServerErrorPB_Code = 16
	// The compare-and-swap specified by an atomic RPC operation failed.
	TabletServerErrorPB_CAS_FAILED TabletServerErrorPB_Code = 17
	// This server leader is not ready for the change configuration operation.
	TabletServerErrorPB_LEADER_NOT_READY_CHANGE_CONFIG TabletServerErrorPB_Code = 18
	// This server leader is not ready to step down.
	TabletServerErrorPB_LEADER_NOT_READY_TO_STEP_DOWN TabletServerErrorPB_Code = 19
	// Adding a peer which is already present in the current raft config.
	TabletServerErrorPB_ADD_CHANGE_CONFIG_ALREADY_PRESENT TabletServerErrorPB_Code = 20
	// Removing a peer which is not present in the current raft config.
	TabletServerErrorPB_REMOVE_CHANGE_CONFIG_NOT_PRESENT TabletServerErrorPB_Code = 21
	// Leader needs to be stepped down before calling change config. This happens
	// if the server we are trying to remove from the config is currently the leader.
	TabletServerErrorPB_LEADER_NEEDS_STEP_DOWN TabletServerErrorPB_Code = 22
	// The operation is not supported.
	TabletServerErrorPB_OPERATION_NOT_SUPPORTED TabletServerErrorPB_Code = 23
	// This tserver is the leader of the consensus configuration, but it's not ready to serve
	// requests. (That means in fact that the elected leader has not yet commited NoOp request.
	// The client must wait a bit for the end of this replica-operation.)
	TabletServerErrorPB_LEADER_NOT_READY_TO_SERVE TabletServerErrorPB_Code = 24
	// This follower hasn't heard from the leader for a specified amount of time.
	TabletServerErrorPB_STALE_FOLLOWER TabletServerErrorPB_Code = 25
	// The operation is already in progress. Used for remote bootstrap requests for now.
	TabletServerErrorPB_ALREADY_IN_PROGRESS TabletServerErrorPB_Code = 26
	// Tablet server has some tablets pending local bootstraps.
	TabletServerErrorPB_PENDING_LOCAL_BOOTSTRAPS TabletServerErrorPB_Code = 27
	// Tablet splitting has been started (after split is completed - tablet stays in this state
	// until it is deleted).
	TabletServerErrorPB_TABLET_SPLIT TabletServerErrorPB_Code = 28
	// Tablet splitting has not been started on this peer yet.
	TabletServerErrorPB_TABLET_SPLIT_PARENT_STILL_LIVE TabletServerErrorPB_Code = 29
)

// Enum value maps for TabletServerErrorPB_Code.
var (
	TabletServerErrorPB_Code_name = map[int32]string{
		1:  "UNKNOWN_ERROR",
		2:  "INVALID_SCHEMA",
		3:  "INVALID_ROW_BLOCK",
		4:  "INVALID_MUTATION",
		5:  "MISMATCHED_SCHEMA",
		6:  "TABLET_NOT_FOUND",
		7:  "SCANNER_EXPIRED",
		8:  "INVALID_SCAN_SPEC",
		9:  "INVALID_CONFIG",
		10: "TABLET_ALREADY_EXISTS",
		11: "TABLET_HAS_A_NEWER_SCHEMA",
		12: "TABLET_NOT_RUNNING",
		13: "INVALID_SNAPSHOT",
		14: "INVALID_SCAN_CALL_SEQ_ID",
		15: "NOT_THE_LEADER",
		16: "WRONG_SERVER_UUID",
		17: "CAS_FAILED",
		18: "LEADER_NOT_READY_CHANGE_CONFIG",
		19: "LEADER_NOT_READY_TO_STEP_DOWN",
		20: "ADD_CHANGE_CONFIG_ALREADY_PRESENT",
		21: "REMOVE_CHANGE_CONFIG_NOT_PRESENT",
		22: "LEADER_NEEDS_STEP_DOWN",
		23: "OPERATION_NOT_SUPPORTED",
		24: "LEADER_NOT_READY_TO_SERVE",
		25: "STALE_FOLLOWER",
		26: "ALREADY_IN_PROGRESS",
		27: "PENDING_LOCAL_BOOTSTRAPS",
		28: "TABLET_SPLIT",
		29: "TABLET_SPLIT_PARENT_STILL_LIVE",
	}
	TabletServerErrorPB_Code_value = map[string]int32{
		"UNKNOWN_ERROR":                     1,
		"INVALID_SCHEMA":                    2,
		"INVALID_ROW_BLOCK":                 3,
		"INVALID_MUTATION":                  4,
		"MISMATCHED_SCHEMA":                 5,
		"TABLET_NOT_FOUND":                  6,
		"SCANNER_EXPIRED":                   7,
		"INVALID_SCAN_SPEC":                 8,
		"INVALID_CONFIG":                    9,
		"TABLET_ALREADY_EXISTS":             10,
		"TABLET_HAS_A_NEWER_SCHEMA":         11,
		"TABLET_NOT_RUNNING":                12,
		"INVALID_SNAPSHOT":                  13,
		"INVALID_SCAN_CALL_SEQ_ID":          14,
		"NOT_THE_LEADER":                    15,
		"WRONG_SERVER_UUID":                 16,
		"CAS_FAILED":                        17,
		"LEADER_NOT_READY_CHANGE_CONFIG":    18,
		"LEADER_NOT_READY_TO_STEP_DOWN":     19,
		"ADD_CHANGE_CONFIG_ALREADY_PRESENT": 20,
		"REMOVE_CHANGE_CONFIG_NOT_PRESENT":  21,
		"LEADER_NEEDS_STEP_DOWN":            22,
		"OPERATION_NOT_SUPPORTED":           23,
		"LEADER_NOT_READY_TO_SERVE":         24,
		"STALE_FOLLOWER":                    25,
		"ALREADY_IN_PROGRESS":               26,
		"PENDING_LOCAL_BOOTSTRAPS":          27,
		"TABLET_SPLIT":                      28,
		"TABLET_SPLIT_PARENT_STILL_LIVE":    29,
	}
)

func (x TabletServerErrorPB_Code) Enum() *TabletServerErrorPB_Code {
	p := new(TabletServerErrorPB_Code)
	*p = x
	return p
}

func (x TabletServerErrorPB_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TabletServerErrorPB_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_yb_tserver_tserver_types_proto_enumTypes[1].Descriptor()
}

func (TabletServerErrorPB_Code) Type() protoreflect.EnumType {
	return &file_yb_tserver_tserver_types_proto_enumTypes[1]
}

func (x TabletServerErrorPB_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *TabletServerErrorPB_Code) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = TabletServerErrorPB_Code(num)
	return nil
}

// Deprecated: Use TabletServerErrorPB_Code.Descriptor instead.
func (TabletServerErrorPB_Code) EnumDescriptor() ([]byte, []int) {
	return file_yb_tserver_tserver_types_proto_rawDescGZIP(), []int{0, 0}
}

// Tablet-server specific errors use this protobuf.
type TabletServerErrorPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The error code.
	Code *TabletServerErrorPB_Code `protobuf:"varint,1,req,name=code,enum=yb.tserver.TabletServerErrorPB_Code,def=1" json:"code,omitempty"`
	// The Status object for the error. This will include a textual
	// message that may be more useful to present in log messages, etc,
	// though its error code is less specific.
	Status *AppStatusPB `protobuf:"bytes,2,req,name=status" json:"status,omitempty"`
}

// Default values for TabletServerErrorPB fields.
const (
	Default_TabletServerErrorPB_Code = TabletServerErrorPB_UNKNOWN_ERROR
)

func (x *TabletServerErrorPB) Reset() {
	*x = TabletServerErrorPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yb_tserver_tserver_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TabletServerErrorPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TabletServerErrorPB) ProtoMessage() {}

func (x *TabletServerErrorPB) ProtoReflect() protoreflect.Message {
	mi := &file_yb_tserver_tserver_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TabletServerErrorPB.ProtoReflect.Descriptor instead.
func (*TabletServerErrorPB) Descriptor() ([]byte, []int) {
	return file_yb_tserver_tserver_types_proto_rawDescGZIP(), []int{0}
}

func (x *TabletServerErrorPB) GetCode() TabletServerErrorPB_Code {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return Default_TabletServerErrorPB_Code
}

func (x *TabletServerErrorPB) GetStatus() *AppStatusPB {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_yb_tserver_tserver_types_proto protoreflect.FileDescriptor

var file_yb_tserver_tserver_types_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x79, 0x62, 0x2f, 0x74, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x74, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x79, 0x62, 0x2e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x1d, 0x79, 0x62,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x77, 0x69, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xff, 0x06, 0x0a, 0x13,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x50, 0x42, 0x12, 0x47, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x0e, 0x32, 0x24, 0x2e, 0x79, 0x62, 0x2e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x50, 0x42, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x3a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x27, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x79,
	0x62, 0x2e, 0x41, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x42, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xf5, 0x05, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x11,
	0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x01, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53, 0x43, 0x48,
	0x45, 0x4d, 0x41, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x52, 0x4f, 0x57, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4d, 0x55, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x49, 0x53, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x45, 0x44,
	0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x41, 0x42,
	0x4c, 0x45, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x06, 0x12,
	0x13, 0x0a, 0x0f, 0x53, 0x43, 0x41, 0x4e, 0x4e, 0x45, 0x52, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52,
	0x45, 0x44, 0x10, 0x07, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x53, 0x43, 0x41, 0x4e, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x10, 0x08, 0x12, 0x12, 0x0a, 0x0e, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x10, 0x09, 0x12,
	0x19, 0x0a, 0x15, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x54, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44,
	0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x0a, 0x12, 0x1d, 0x0a, 0x19, 0x54, 0x41,
	0x42, 0x4c, 0x45, 0x54, 0x5f, 0x48, 0x41, 0x53, 0x5f, 0x41, 0x5f, 0x4e, 0x45, 0x57, 0x45, 0x52,
	0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x10, 0x0b, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x41, 0x42,
	0x4c, 0x45, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10,
	0x0c, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53, 0x4e, 0x41,
	0x50, 0x53, 0x48, 0x4f, 0x54, 0x10, 0x0d, 0x12, 0x1c, 0x0a, 0x18, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x53, 0x43, 0x41, 0x4e, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x53, 0x45, 0x51,
	0x5f, 0x49, 0x44, 0x10, 0x0e, 0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x4f, 0x54, 0x5f, 0x54, 0x48, 0x45,
	0x5f, 0x4c, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0x0f, 0x12, 0x15, 0x0a, 0x11, 0x57, 0x52, 0x4f,
	0x4e, 0x47, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x55, 0x55, 0x49, 0x44, 0x10, 0x10,
	0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x41, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x11,
	0x12, 0x22, 0x0a, 0x1e, 0x4c, 0x45, 0x41, 0x44, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x52,
	0x45, 0x41, 0x44, 0x59, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x46,
	0x49, 0x47, 0x10, 0x12, 0x12, 0x21, 0x0a, 0x1d, 0x4c, 0x45, 0x41, 0x44, 0x45, 0x52, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x54, 0x4f, 0x5f, 0x53, 0x54, 0x45, 0x50,
	0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x13, 0x12, 0x25, 0x0a, 0x21, 0x41, 0x44, 0x44, 0x5f, 0x43,
	0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x41, 0x4c, 0x52,
	0x45, 0x41, 0x44, 0x59, 0x5f, 0x50, 0x52, 0x45, 0x53, 0x45, 0x4e, 0x54, 0x10, 0x14, 0x12, 0x24,
	0x0a, 0x20, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f,
	0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x52, 0x45, 0x53, 0x45,
	0x4e, 0x54, 0x10, 0x15, 0x12, 0x1a, 0x0a, 0x16, 0x4c, 0x45, 0x41, 0x44, 0x45, 0x52, 0x5f, 0x4e,
	0x45, 0x45, 0x44, 0x53, 0x5f, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x16,
	0x12, 0x1b, 0x0a, 0x17, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x17, 0x12, 0x1d, 0x0a,
	0x19, 0x4c, 0x45, 0x41, 0x44, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x44,
	0x59, 0x5f, 0x54, 0x4f, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x10, 0x18, 0x12, 0x12, 0x0a, 0x0e,
	0x53, 0x54, 0x41, 0x4c, 0x45, 0x5f, 0x46, 0x4f, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x52, 0x10, 0x19,
	0x12, 0x17, 0x0a, 0x13, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x49, 0x4e, 0x5f, 0x50,
	0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x1a, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x45, 0x4e,
	0x44, 0x49, 0x4e, 0x47, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x5f, 0x42, 0x4f, 0x4f, 0x54, 0x53,
	0x54, 0x52, 0x41, 0x50, 0x53, 0x10, 0x1b, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x41, 0x42, 0x4c, 0x45,
	0x54, 0x5f, 0x53, 0x50, 0x4c, 0x49, 0x54, 0x10, 0x1c, 0x12, 0x22, 0x0a, 0x1e, 0x54, 0x41, 0x42,
	0x4c, 0x45, 0x54, 0x5f, 0x53, 0x50, 0x4c, 0x49, 0x54, 0x5f, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x54,
	0x5f, 0x53, 0x54, 0x49, 0x4c, 0x4c, 0x5f, 0x4c, 0x49, 0x56, 0x45, 0x10, 0x1d, 0x2a, 0x82, 0x01,
	0x0a, 0x16, 0x54, 0x53, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x42, 0x12, 0x0c, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0xe7, 0x07, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45,
	0x52, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x49, 0x5a, 0x49, 0x4e, 0x47, 0x10, 0x00,
	0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x5f, 0x52, 0x55, 0x4e, 0x4e,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52,
	0x5f, 0x51, 0x55, 0x49, 0x45, 0x53, 0x43, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10,
	0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x5f, 0x53, 0x48, 0x55, 0x54, 0x44, 0x4f, 0x57, 0x4e,
	0x10, 0x03, 0x42, 0x10, 0x0a, 0x0e, 0x6f, 0x72, 0x67, 0x2e, 0x79, 0x62, 0x2e, 0x74, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72,
}

var (
	file_yb_tserver_tserver_types_proto_rawDescOnce sync.Once
	file_yb_tserver_tserver_types_proto_rawDescData = file_yb_tserver_tserver_types_proto_rawDesc
)

func file_yb_tserver_tserver_types_proto_rawDescGZIP() []byte {
	file_yb_tserver_tserver_types_proto_rawDescOnce.Do(func() {
		file_yb_tserver_tserver_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_yb_tserver_tserver_types_proto_rawDescData)
	})
	return file_yb_tserver_tserver_types_proto_rawDescData
}

var file_yb_tserver_tserver_types_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_yb_tserver_tserver_types_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yb_tserver_tserver_types_proto_goTypes = []interface{}{
	(TSTabletManagerStatePB)(0),   // 0: yb.tserver.TSTabletManagerStatePB
	(TabletServerErrorPB_Code)(0), // 1: yb.tserver.TabletServerErrorPB.Code
	(*TabletServerErrorPB)(nil),   // 2: yb.tserver.TabletServerErrorPB
	(*AppStatusPB)(nil),           // 3: yb.AppStatusPB
}
var file_yb_tserver_tserver_types_proto_depIdxs = []int32{
	1, // 0: yb.tserver.TabletServerErrorPB.code:type_name -> yb.tserver.TabletServerErrorPB.Code
	3, // 1: yb.tserver.TabletServerErrorPB.status:type_name -> yb.AppStatusPB
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_yb_tserver_tserver_types_proto_init() }
func file_yb_tserver_tserver_types_proto_init() {
	if File_yb_tserver_tserver_types_proto != nil {
		return
	}
	file_yb_common_wire_protocol_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yb_tserver_tserver_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TabletServerErrorPB); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yb_tserver_tserver_types_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yb_tserver_tserver_types_proto_goTypes,
		DependencyIndexes: file_yb_tserver_tserver_types_proto_depIdxs,
		EnumInfos:         file_yb_tserver_tserver_types_proto_enumTypes,
		MessageInfos:      file_yb_tserver_tserver_types_proto_msgTypes,
	}.Build()
	File_yb_tserver_tserver_types_proto = out.File
	file_yb_tserver_tserver_types_proto_rawDesc = nil
	file_yb_tserver_tserver_types_proto_goTypes = nil
	file_yb_tserver_tserver_types_proto_depIdxs = nil
}