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
// source: yb/tablet/tablet_types.proto

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

// The enum of Raft group states.
// Raft group states are sent in TabletReports and kept in TabletPeer.
type RaftGroupStatePB int32

const (
	RaftGroupStatePB_UNKNOWN RaftGroupStatePB = 999
	// Raft group has not yet started.
	RaftGroupStatePB_NOT_STARTED RaftGroupStatePB = 5
	// Indicates the Raft group is bootstrapping, i.e. that the Raft group is not available for RPC.
	RaftGroupStatePB_BOOTSTRAPPING RaftGroupStatePB = 0
	// Once the configuration phase is over Peers are in RUNNING state. In this state Peers are
	// available for client RPCs.
	RaftGroupStatePB_RUNNING RaftGroupStatePB = 1
	// The Raft group failed to for some reason. TabletPeer::error() will return the reason for the
	// failure.
	RaftGroupStatePB_FAILED RaftGroupStatePB = 2
	// The Raft group is shutting down, and will not accept further requests.
	RaftGroupStatePB_QUIESCING RaftGroupStatePB = 3
	// The Raft group has been stopped.
	RaftGroupStatePB_SHUTDOWN RaftGroupStatePB = 4
)

// Enum value maps for RaftGroupStatePB.
var (
	RaftGroupStatePB_name = map[int32]string{
		999: "UNKNOWN",
		5:   "NOT_STARTED",
		0:   "BOOTSTRAPPING",
		1:   "RUNNING",
		2:   "FAILED",
		3:   "QUIESCING",
		4:   "SHUTDOWN",
	}
	RaftGroupStatePB_value = map[string]int32{
		"UNKNOWN":       999,
		"NOT_STARTED":   5,
		"BOOTSTRAPPING": 0,
		"RUNNING":       1,
		"FAILED":        2,
		"QUIESCING":     3,
		"SHUTDOWN":      4,
	}
)

func (x RaftGroupStatePB) Enum() *RaftGroupStatePB {
	p := new(RaftGroupStatePB)
	*p = x
	return p
}

func (x RaftGroupStatePB) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RaftGroupStatePB) Descriptor() protoreflect.EnumDescriptor {
	return file_yb_tablet_tablet_types_proto_enumTypes[0].Descriptor()
}

func (RaftGroupStatePB) Type() protoreflect.EnumType {
	return &file_yb_tablet_tablet_types_proto_enumTypes[0]
}

func (x RaftGroupStatePB) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RaftGroupStatePB) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RaftGroupStatePB(num)
	return nil
}

// Deprecated: Use RaftGroupStatePB.Descriptor instead.
func (RaftGroupStatePB) EnumDescriptor() ([]byte, []int) {
	return file_yb_tablet_tablet_types_proto_rawDescGZIP(), []int{0}
}

// State flags indicating whether the tablet is in the middle of being copied
// and is therefore not possible to bring up, whether it has been deleted, or
// whether the data is in a usable state.
type TabletDataState int32

const (
	TabletDataState_TABLET_DATA_UNKNOWN TabletDataState = 999
	// The tablet is set to TABLET_DATA_COPYING state when in the middle of
	// remote bootstrap while copying data files from a remote peer. If a tablet
	// server crashes with a tablet in this state, the tablet must be deleted and
	// the remote bootstrap process must be restarted for that tablet.
	TabletDataState_TABLET_DATA_COPYING TabletDataState = 0
	// Fresh empty tablets and successfully copied tablets are set to the
	// TABLET_DATA_READY state.
	TabletDataState_TABLET_DATA_READY TabletDataState = 1
	// This tablet is in the process of being deleted.
	// The tablet server should "roll forward" the deletion during boot,
	// rather than trying to load the tablet.
	TabletDataState_TABLET_DATA_DELETED TabletDataState = 2
	// The tablet has been deleted, and now just consists of a "tombstone".
	TabletDataState_TABLET_DATA_TOMBSTONED TabletDataState = 3
	// This tablet split has been completed. In this state tablet could only be used as a remote
	// bootstrap source for some time (follower_unavailable_considered_failed_sec) and then will be
	// shutdown.
	// TODO(tsplit): After https://github.com/yugabyte/yugabyte-db/issues/1461:
	// - implement deleting tablet in state TABLET_DATA_SPLIT_COMPLETED after
	//   follower_unavailable_considered_failed_sec.
	// - add integration test for remote bootstrap from source split tablet before/after
	// follower_unavailable_considered_failed_sec.
	TabletDataState_TABLET_DATA_SPLIT_COMPLETED TabletDataState = 4
	// This tablet has been initialized as a subtablet of another tablet undergoing a split. Once this
	// tablet is ready to assume responsibility for it's portion of the split keyspace, it will be
	// moved to the TABLET_DATA_READY state.
	TabletDataState_TABLET_DATA_INIT_STARTED TabletDataState = 5
)

// Enum value maps for TabletDataState.
var (
	TabletDataState_name = map[int32]string{
		999: "TABLET_DATA_UNKNOWN",
		0:   "TABLET_DATA_COPYING",
		1:   "TABLET_DATA_READY",
		2:   "TABLET_DATA_DELETED",
		3:   "TABLET_DATA_TOMBSTONED",
		4:   "TABLET_DATA_SPLIT_COMPLETED",
		5:   "TABLET_DATA_INIT_STARTED",
	}
	TabletDataState_value = map[string]int32{
		"TABLET_DATA_UNKNOWN":         999,
		"TABLET_DATA_COPYING":         0,
		"TABLET_DATA_READY":           1,
		"TABLET_DATA_DELETED":         2,
		"TABLET_DATA_TOMBSTONED":      3,
		"TABLET_DATA_SPLIT_COMPLETED": 4,
		"TABLET_DATA_INIT_STARTED":    5,
	}
)

func (x TabletDataState) Enum() *TabletDataState {
	p := new(TabletDataState)
	*p = x
	return p
}

func (x TabletDataState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TabletDataState) Descriptor() protoreflect.EnumDescriptor {
	return file_yb_tablet_tablet_types_proto_enumTypes[1].Descriptor()
}

func (TabletDataState) Type() protoreflect.EnumType {
	return &file_yb_tablet_tablet_types_proto_enumTypes[1]
}

func (x TabletDataState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *TabletDataState) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = TabletDataState(num)
	return nil
}

// Deprecated: Use TabletDataState.Descriptor instead.
func (TabletDataState) EnumDescriptor() ([]byte, []int) {
	return file_yb_tablet_tablet_types_proto_rawDescGZIP(), []int{1}
}

var File_yb_tablet_tablet_types_proto protoreflect.FileDescriptor

var file_yb_tablet_tablet_types_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x79, 0x62, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x2f, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x79, 0x62, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x2a, 0x7a, 0x0a, 0x10, 0x52, 0x61, 0x66,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x42, 0x12, 0x0c, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0xe7, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x4e,
	0x4f, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d,
	0x42, 0x4f, 0x4f, 0x54, 0x53, 0x54, 0x52, 0x41, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x51, 0x55, 0x49, 0x45,
	0x53, 0x43, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x48, 0x55, 0x54, 0x44,
	0x4f, 0x57, 0x4e, 0x10, 0x04, 0x2a, 0xcf, 0x01, 0x0a, 0x0f, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x13, 0x54, 0x41, 0x42,
	0x4c, 0x45, 0x54, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0xe7, 0x07, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x54, 0x5f, 0x44, 0x41,
	0x54, 0x41, 0x5f, 0x43, 0x4f, 0x50, 0x59, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11,
	0x54, 0x41, 0x42, 0x4c, 0x45, 0x54, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x52, 0x45, 0x41, 0x44,
	0x59, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x54, 0x5f, 0x44, 0x41,
	0x54, 0x41, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16,
	0x54, 0x41, 0x42, 0x4c, 0x45, 0x54, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x4f, 0x4d, 0x42,
	0x53, 0x54, 0x4f, 0x4e, 0x45, 0x44, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x54, 0x41, 0x42, 0x4c,
	0x45, 0x54, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x50, 0x4c, 0x49, 0x54, 0x5f, 0x43, 0x4f,
	0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x1c, 0x0a, 0x18, 0x54, 0x41, 0x42,
	0x4c, 0x45, 0x54, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x5f, 0x53, 0x54,
	0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x05, 0x42, 0x0f, 0x0a, 0x0d, 0x6f, 0x72, 0x67, 0x2e, 0x79,
	0x62, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x74,
}

var (
	file_yb_tablet_tablet_types_proto_rawDescOnce sync.Once
	file_yb_tablet_tablet_types_proto_rawDescData = file_yb_tablet_tablet_types_proto_rawDesc
)

func file_yb_tablet_tablet_types_proto_rawDescGZIP() []byte {
	file_yb_tablet_tablet_types_proto_rawDescOnce.Do(func() {
		file_yb_tablet_tablet_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_yb_tablet_tablet_types_proto_rawDescData)
	})
	return file_yb_tablet_tablet_types_proto_rawDescData
}

var file_yb_tablet_tablet_types_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_yb_tablet_tablet_types_proto_goTypes = []interface{}{
	(RaftGroupStatePB)(0), // 0: yb.tablet.RaftGroupStatePB
	(TabletDataState)(0),  // 1: yb.tablet.TabletDataState
}
var file_yb_tablet_tablet_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yb_tablet_tablet_types_proto_init() }
func file_yb_tablet_tablet_types_proto_init() {
	if File_yb_tablet_tablet_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yb_tablet_tablet_types_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yb_tablet_tablet_types_proto_goTypes,
		DependencyIndexes: file_yb_tablet_tablet_types_proto_depIdxs,
		EnumInfos:         file_yb_tablet_tablet_types_proto_enumTypes,
	}.Build()
	File_yb_tablet_tablet_types_proto = out.File
	file_yb_tablet_tablet_types_proto_rawDesc = nil
	file_yb_tablet_tablet_types_proto_goTypes = nil
	file_yb_tablet_tablet_types_proto_depIdxs = nil
}