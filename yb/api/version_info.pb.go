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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: yb/util/version_info.proto

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

// Information about the build environment, configuration, etc.
type VersionInfoPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GitHash        *string `protobuf:"bytes,1,opt,name=git_hash,json=gitHash" json:"git_hash,omitempty"`
	BuildHostname  *string `protobuf:"bytes,2,opt,name=build_hostname,json=buildHostname" json:"build_hostname,omitempty"`
	BuildTimestamp *string `protobuf:"bytes,3,opt,name=build_timestamp,json=buildTimestamp" json:"build_timestamp,omitempty"`
	BuildUsername  *string `protobuf:"bytes,4,opt,name=build_username,json=buildUsername" json:"build_username,omitempty"`
	BuildCleanRepo *bool   `protobuf:"varint,5,opt,name=build_clean_repo,json=buildCleanRepo" json:"build_clean_repo,omitempty"`
	BuildId        *string `protobuf:"bytes,6,opt,name=build_id,json=buildId" json:"build_id,omitempty"`
	BuildType      *string `protobuf:"bytes,7,opt,name=build_type,json=buildType" json:"build_type,omitempty"`
	VersionNumber  *string `protobuf:"bytes,9,opt,name=version_number,json=versionNumber" json:"version_number,omitempty"`
	BuildNumber    *string `protobuf:"bytes,10,opt,name=build_number,json=buildNumber" json:"build_number,omitempty"`
}

func (x *VersionInfoPB) Reset() {
	*x = VersionInfoPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yb_util_version_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionInfoPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionInfoPB) ProtoMessage() {}

func (x *VersionInfoPB) ProtoReflect() protoreflect.Message {
	mi := &file_yb_util_version_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionInfoPB.ProtoReflect.Descriptor instead.
func (*VersionInfoPB) Descriptor() ([]byte, []int) {
	return file_yb_util_version_info_proto_rawDescGZIP(), []int{0}
}

func (x *VersionInfoPB) GetGitHash() string {
	if x != nil && x.GitHash != nil {
		return *x.GitHash
	}
	return ""
}

func (x *VersionInfoPB) GetBuildHostname() string {
	if x != nil && x.BuildHostname != nil {
		return *x.BuildHostname
	}
	return ""
}

func (x *VersionInfoPB) GetBuildTimestamp() string {
	if x != nil && x.BuildTimestamp != nil {
		return *x.BuildTimestamp
	}
	return ""
}

func (x *VersionInfoPB) GetBuildUsername() string {
	if x != nil && x.BuildUsername != nil {
		return *x.BuildUsername
	}
	return ""
}

func (x *VersionInfoPB) GetBuildCleanRepo() bool {
	if x != nil && x.BuildCleanRepo != nil {
		return *x.BuildCleanRepo
	}
	return false
}

func (x *VersionInfoPB) GetBuildId() string {
	if x != nil && x.BuildId != nil {
		return *x.BuildId
	}
	return ""
}

func (x *VersionInfoPB) GetBuildType() string {
	if x != nil && x.BuildType != nil {
		return *x.BuildType
	}
	return ""
}

func (x *VersionInfoPB) GetVersionNumber() string {
	if x != nil && x.VersionNumber != nil {
		return *x.VersionNumber
	}
	return ""
}

func (x *VersionInfoPB) GetBuildNumber() string {
	if x != nil && x.BuildNumber != nil {
		return *x.BuildNumber
	}
	return ""
}

var File_yb_util_version_info_proto protoreflect.FileDescriptor

var file_yb_util_version_info_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x79, 0x62, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x79, 0x62,
	0x22, 0xcf, 0x02, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x50, 0x42, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x69, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x69, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x25, 0x0a,
	0x0e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x48, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x25, 0x0a,
	0x0e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x63, 0x6c,
	0x65, 0x61, 0x6e, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x19,
	0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x21, 0x0a, 0x0c, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x2e, 0x79, 0x62,
}

var (
	file_yb_util_version_info_proto_rawDescOnce sync.Once
	file_yb_util_version_info_proto_rawDescData = file_yb_util_version_info_proto_rawDesc
)

func file_yb_util_version_info_proto_rawDescGZIP() []byte {
	file_yb_util_version_info_proto_rawDescOnce.Do(func() {
		file_yb_util_version_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_yb_util_version_info_proto_rawDescData)
	})
	return file_yb_util_version_info_proto_rawDescData
}

var file_yb_util_version_info_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yb_util_version_info_proto_goTypes = []interface{}{
	(*VersionInfoPB)(nil), // 0: yb.VersionInfoPB
}
var file_yb_util_version_info_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yb_util_version_info_proto_init() }
func file_yb_util_version_info_proto_init() {
	if File_yb_util_version_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yb_util_version_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionInfoPB); i {
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
			RawDescriptor: file_yb_util_version_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yb_util_version_info_proto_goTypes,
		DependencyIndexes: file_yb_util_version_info_proto_depIdxs,
		MessageInfos:      file_yb_util_version_info_proto_msgTypes,
	}.Build()
	File_yb_util_version_info_proto = out.File
	file_yb_util_version_info_proto_rawDesc = nil
	file_yb_util_version_info_proto_goTypes = nil
	file_yb_util_version_info_proto_depIdxs = nil
}
