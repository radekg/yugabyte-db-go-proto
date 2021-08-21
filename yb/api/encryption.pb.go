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
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: yb/util/encryption.proto

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

type EncryptionParamsPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataKey []byte `protobuf:"bytes,1,opt,name=data_key,json=dataKey,proto3" json:"data_key,omitempty"`
	Nonce   []byte `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Counter int32  `protobuf:"varint,3,opt,name=counter,proto3" json:"counter,omitempty"`
	// When computing counter increment, do we want to overflow the counter into the rest of the
	// initialization vector as part of the new format.
	OpensslCompatibleCounterOverflow bool `protobuf:"varint,4,opt,name=openssl_compatible_counter_overflow,json=opensslCompatibleCounterOverflow,proto3" json:"openssl_compatible_counter_overflow,omitempty"`
}

func (x *EncryptionParamsPB) Reset() {
	*x = EncryptionParamsPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yb_util_encryption_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptionParamsPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptionParamsPB) ProtoMessage() {}

func (x *EncryptionParamsPB) ProtoReflect() protoreflect.Message {
	mi := &file_yb_util_encryption_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptionParamsPB.ProtoReflect.Descriptor instead.
func (*EncryptionParamsPB) Descriptor() ([]byte, []int) {
	return file_yb_util_encryption_proto_rawDescGZIP(), []int{0}
}

func (x *EncryptionParamsPB) GetDataKey() []byte {
	if x != nil {
		return x.DataKey
	}
	return nil
}

func (x *EncryptionParamsPB) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *EncryptionParamsPB) GetCounter() int32 {
	if x != nil {
		return x.Counter
	}
	return 0
}

func (x *EncryptionParamsPB) GetOpensslCompatibleCounterOverflow() bool {
	if x != nil {
		return x.OpensslCompatibleCounterOverflow
	}
	return false
}

type UniverseKeysPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Map map[string][]byte `protobuf:"bytes,1,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UniverseKeysPB) Reset() {
	*x = UniverseKeysPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yb_util_encryption_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniverseKeysPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniverseKeysPB) ProtoMessage() {}

func (x *UniverseKeysPB) ProtoReflect() protoreflect.Message {
	mi := &file_yb_util_encryption_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniverseKeysPB.ProtoReflect.Descriptor instead.
func (*UniverseKeysPB) Descriptor() ([]byte, []int) {
	return file_yb_util_encryption_proto_rawDescGZIP(), []int{1}
}

func (x *UniverseKeysPB) GetMap() map[string][]byte {
	if x != nil {
		return x.Map
	}
	return nil
}

type UniverseKeyRegistryPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EncryptionEnabled bool                           `protobuf:"varint,1,opt,name=encryption_enabled,json=encryptionEnabled,proto3" json:"encryption_enabled,omitempty"`
	UniverseKeys      map[string]*EncryptionParamsPB `protobuf:"bytes,2,rep,name=universe_keys,json=universeKeys,proto3" json:"universe_keys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	LatestVersionId   string                         `protobuf:"bytes,3,opt,name=latest_version_id,json=latestVersionId,proto3" json:"latest_version_id,omitempty"`
}

func (x *UniverseKeyRegistryPB) Reset() {
	*x = UniverseKeyRegistryPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yb_util_encryption_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniverseKeyRegistryPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniverseKeyRegistryPB) ProtoMessage() {}

func (x *UniverseKeyRegistryPB) ProtoReflect() protoreflect.Message {
	mi := &file_yb_util_encryption_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniverseKeyRegistryPB.ProtoReflect.Descriptor instead.
func (*UniverseKeyRegistryPB) Descriptor() ([]byte, []int) {
	return file_yb_util_encryption_proto_rawDescGZIP(), []int{2}
}

func (x *UniverseKeyRegistryPB) GetEncryptionEnabled() bool {
	if x != nil {
		return x.EncryptionEnabled
	}
	return false
}

func (x *UniverseKeyRegistryPB) GetUniverseKeys() map[string]*EncryptionParamsPB {
	if x != nil {
		return x.UniverseKeys
	}
	return nil
}

func (x *UniverseKeyRegistryPB) GetLatestVersionId() string {
	if x != nil {
		return x.LatestVersionId
	}
	return ""
}

var File_yb_util_encryption_proto protoreflect.FileDescriptor

var file_yb_util_encryption_proto_rawDesc = []byte{
	0x0a, 0x18, 0x79, 0x62, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x79, 0x62, 0x22, 0xae,
	0x01, 0x0a, 0x12, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x50, 0x42, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x4b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x12, 0x4d, 0x0a, 0x23, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x73, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x6f,
	0x76, 0x65, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x20, 0x6f,
	0x70, 0x65, 0x6e, 0x73, 0x73, 0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x6c, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x4f, 0x76, 0x65, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x22,
	0x77, 0x0a, 0x0e, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x50,
	0x42, 0x12, 0x2d, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x79, 0x62, 0x2e, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x73,
	0x50, 0x42, 0x2e, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x6d, 0x61, 0x70,
	0x1a, 0x36, 0x0a, 0x08, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9d, 0x02, 0x0a, 0x15, 0x55, 0x6e, 0x69,
	0x76, 0x65, 0x72, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x50, 0x42, 0x12, 0x2d, 0x0a, 0x12, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11,
	0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x12, 0x50, 0x0a, 0x0d, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x6b, 0x65,
	0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x79, 0x62, 0x2e, 0x55, 0x6e,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x50, 0x42, 0x2e, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x4b,
	0x65, 0x79, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a,
	0x57, 0x0a, 0x11, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x79, 0x62, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x50, 0x42, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x2e,
	0x79, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yb_util_encryption_proto_rawDescOnce sync.Once
	file_yb_util_encryption_proto_rawDescData = file_yb_util_encryption_proto_rawDesc
)

func file_yb_util_encryption_proto_rawDescGZIP() []byte {
	file_yb_util_encryption_proto_rawDescOnce.Do(func() {
		file_yb_util_encryption_proto_rawDescData = protoimpl.X.CompressGZIP(file_yb_util_encryption_proto_rawDescData)
	})
	return file_yb_util_encryption_proto_rawDescData
}

var file_yb_util_encryption_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yb_util_encryption_proto_goTypes = []interface{}{
	(*EncryptionParamsPB)(nil),    // 0: yb.EncryptionParamsPB
	(*UniverseKeysPB)(nil),        // 1: yb.UniverseKeysPB
	(*UniverseKeyRegistryPB)(nil), // 2: yb.UniverseKeyRegistryPB
	nil,                           // 3: yb.UniverseKeysPB.MapEntry
	nil,                           // 4: yb.UniverseKeyRegistryPB.UniverseKeysEntry
}
var file_yb_util_encryption_proto_depIdxs = []int32{
	3, // 0: yb.UniverseKeysPB.map:type_name -> yb.UniverseKeysPB.MapEntry
	4, // 1: yb.UniverseKeyRegistryPB.universe_keys:type_name -> yb.UniverseKeyRegistryPB.UniverseKeysEntry
	0, // 2: yb.UniverseKeyRegistryPB.UniverseKeysEntry.value:type_name -> yb.EncryptionParamsPB
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yb_util_encryption_proto_init() }
func file_yb_util_encryption_proto_init() {
	if File_yb_util_encryption_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yb_util_encryption_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptionParamsPB); i {
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
		file_yb_util_encryption_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniverseKeysPB); i {
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
		file_yb_util_encryption_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniverseKeyRegistryPB); i {
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
			RawDescriptor: file_yb_util_encryption_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yb_util_encryption_proto_goTypes,
		DependencyIndexes: file_yb_util_encryption_proto_depIdxs,
		MessageInfos:      file_yb_util_encryption_proto_msgTypes,
	}.Build()
	File_yb_util_encryption_proto = out.File
	file_yb_util_encryption_proto_rawDesc = nil
	file_yb_util_encryption_proto_goTypes = nil
	file_yb_util_encryption_proto_depIdxs = nil
}