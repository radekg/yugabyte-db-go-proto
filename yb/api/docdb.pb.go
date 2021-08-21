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
// source: yb/docdb/docdb.proto

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

type KeyValuePairPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key                []byte  `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value              []byte  `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	ExternalHybridTime *uint64 `protobuf:"fixed64,3,opt,name=external_hybrid_time,json=externalHybridTime" json:"external_hybrid_time,omitempty"`
}

func (x *KeyValuePairPB) Reset() {
	*x = KeyValuePairPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yb_docdb_docdb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValuePairPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValuePairPB) ProtoMessage() {}

func (x *KeyValuePairPB) ProtoReflect() protoreflect.Message {
	mi := &file_yb_docdb_docdb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValuePairPB.ProtoReflect.Descriptor instead.
func (*KeyValuePairPB) Descriptor() ([]byte, []int) {
	return file_yb_docdb_docdb_proto_rawDescGZIP(), []int{0}
}

func (x *KeyValuePairPB) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *KeyValuePairPB) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *KeyValuePairPB) GetExternalHybridTime() uint64 {
	if x != nil && x.ExternalHybridTime != nil {
		return *x.ExternalHybridTime
	}
	return 0
}

type ApplyExternalTransactionPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId    []byte  `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId" json:"transaction_id,omitempty"`
	CommitHybridTime *uint64 `protobuf:"fixed64,2,opt,name=commit_hybrid_time,json=commitHybridTime" json:"commit_hybrid_time,omitempty"`
}

func (x *ApplyExternalTransactionPB) Reset() {
	*x = ApplyExternalTransactionPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yb_docdb_docdb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyExternalTransactionPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyExternalTransactionPB) ProtoMessage() {}

func (x *ApplyExternalTransactionPB) ProtoReflect() protoreflect.Message {
	mi := &file_yb_docdb_docdb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyExternalTransactionPB.ProtoReflect.Descriptor instead.
func (*ApplyExternalTransactionPB) Descriptor() ([]byte, []int) {
	return file_yb_docdb_docdb_proto_rawDescGZIP(), []int{1}
}

func (x *ApplyExternalTransactionPB) GetTransactionId() []byte {
	if x != nil {
		return x.TransactionId
	}
	return nil
}

func (x *ApplyExternalTransactionPB) GetCommitHybridTime() uint64 {
	if x != nil && x.CommitHybridTime != nil {
		return *x.CommitHybridTime
	}
	return 0
}

// A set of key/value pairs to be written into RocksDB.
type KeyValueWriteBatchPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WritePairs                []*KeyValuePairPB      `protobuf:"bytes,1,rep,name=write_pairs,json=writePairs" json:"write_pairs,omitempty"`
	Transaction               *TransactionMetadataPB `protobuf:"bytes,2,opt,name=transaction" json:"transaction,omitempty"`
	DEPRECATEDMayHaveMetadata *bool                  `protobuf:"varint,3,opt,name=DEPRECATED_may_have_metadata,json=DEPRECATEDMayHaveMetadata" json:"DEPRECATED_may_have_metadata,omitempty"`
	// Used by serializable isolation transactions and row locking statements to store read intents.
	// In case of read-modify-write operation both read_pairs and write_pairs could present.
	ReadPairs                 []*KeyValuePairPB             `protobuf:"bytes,5,rep,name=read_pairs,json=readPairs" json:"read_pairs,omitempty"`
	RowMarkType               *RowMarkType                  `protobuf:"varint,6,opt,name=row_mark_type,json=rowMarkType,enum=yb.RowMarkType" json:"row_mark_type,omitempty"`
	ApplyExternalTransactions []*ApplyExternalTransactionPB `protobuf:"bytes,7,rep,name=apply_external_transactions,json=applyExternalTransactions" json:"apply_external_transactions,omitempty"`
}

func (x *KeyValueWriteBatchPB) Reset() {
	*x = KeyValueWriteBatchPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yb_docdb_docdb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValueWriteBatchPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValueWriteBatchPB) ProtoMessage() {}

func (x *KeyValueWriteBatchPB) ProtoReflect() protoreflect.Message {
	mi := &file_yb_docdb_docdb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValueWriteBatchPB.ProtoReflect.Descriptor instead.
func (*KeyValueWriteBatchPB) Descriptor() ([]byte, []int) {
	return file_yb_docdb_docdb_proto_rawDescGZIP(), []int{2}
}

func (x *KeyValueWriteBatchPB) GetWritePairs() []*KeyValuePairPB {
	if x != nil {
		return x.WritePairs
	}
	return nil
}

func (x *KeyValueWriteBatchPB) GetTransaction() *TransactionMetadataPB {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *KeyValueWriteBatchPB) GetDEPRECATEDMayHaveMetadata() bool {
	if x != nil && x.DEPRECATEDMayHaveMetadata != nil {
		return *x.DEPRECATEDMayHaveMetadata
	}
	return false
}

func (x *KeyValueWriteBatchPB) GetReadPairs() []*KeyValuePairPB {
	if x != nil {
		return x.ReadPairs
	}
	return nil
}

func (x *KeyValueWriteBatchPB) GetRowMarkType() RowMarkType {
	if x != nil && x.RowMarkType != nil {
		return *x.RowMarkType
	}
	return RowMarkType_ROW_MARK_EXCLUSIVE
}

func (x *KeyValueWriteBatchPB) GetApplyExternalTransactions() []*ApplyExternalTransactionPB {
	if x != nil {
		return x.ApplyExternalTransactions
	}
	return nil
}

type ConsensusFrontierPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpId             *OpIdPB `protobuf:"bytes,1,opt,name=op_id,json=opId" json:"op_id,omitempty"`
	HybridTime       *uint64 `protobuf:"fixed64,2,opt,name=hybrid_time,json=hybridTime" json:"hybrid_time,omitempty"`
	HistoryCutoff    *uint64 `protobuf:"fixed64,3,opt,name=history_cutoff,json=historyCutoff" json:"history_cutoff,omitempty"`
	HybridTimeFilter *uint64 `protobuf:"fixed64,4,opt,name=hybrid_time_filter,json=hybridTimeFilter" json:"hybrid_time_filter,omitempty"`
}

func (x *ConsensusFrontierPB) Reset() {
	*x = ConsensusFrontierPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yb_docdb_docdb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsensusFrontierPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsensusFrontierPB) ProtoMessage() {}

func (x *ConsensusFrontierPB) ProtoReflect() protoreflect.Message {
	mi := &file_yb_docdb_docdb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsensusFrontierPB.ProtoReflect.Descriptor instead.
func (*ConsensusFrontierPB) Descriptor() ([]byte, []int) {
	return file_yb_docdb_docdb_proto_rawDescGZIP(), []int{3}
}

func (x *ConsensusFrontierPB) GetOpId() *OpIdPB {
	if x != nil {
		return x.OpId
	}
	return nil
}

func (x *ConsensusFrontierPB) GetHybridTime() uint64 {
	if x != nil && x.HybridTime != nil {
		return *x.HybridTime
	}
	return 0
}

func (x *ConsensusFrontierPB) GetHistoryCutoff() uint64 {
	if x != nil && x.HistoryCutoff != nil {
		return *x.HistoryCutoff
	}
	return 0
}

func (x *ConsensusFrontierPB) GetHybridTimeFilter() uint64 {
	if x != nil && x.HybridTimeFilter != nil {
		return *x.HybridTimeFilter
	}
	return 0
}

type ApplyTransactionStatePB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Apply should be continued from this key in reverse index.
	Key []byte `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// Write id for first continued record.
	WriteId *uint32 `protobuf:"varint,2,opt,name=write_id,json=writeId" json:"write_id,omitempty"`
	// Transaction commit hybrid time.
	CommitHt *uint64 `protobuf:"fixed64,3,opt,name=commit_ht,json=commitHt" json:"commit_ht,omitempty"`
}

func (x *ApplyTransactionStatePB) Reset() {
	*x = ApplyTransactionStatePB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yb_docdb_docdb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyTransactionStatePB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyTransactionStatePB) ProtoMessage() {}

func (x *ApplyTransactionStatePB) ProtoReflect() protoreflect.Message {
	mi := &file_yb_docdb_docdb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyTransactionStatePB.ProtoReflect.Descriptor instead.
func (*ApplyTransactionStatePB) Descriptor() ([]byte, []int) {
	return file_yb_docdb_docdb_proto_rawDescGZIP(), []int{4}
}

func (x *ApplyTransactionStatePB) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *ApplyTransactionStatePB) GetWriteId() uint32 {
	if x != nil && x.WriteId != nil {
		return *x.WriteId
	}
	return 0
}

func (x *ApplyTransactionStatePB) GetCommitHt() uint64 {
	if x != nil && x.CommitHt != nil {
		return *x.CommitHt
	}
	return 0
}

var File_yb_docdb_docdb_proto protoreflect.FileDescriptor

var file_yb_docdb_docdb_proto_rawDesc = []byte{
	0x0a, 0x14, 0x79, 0x62, 0x2f, 0x64, 0x6f, 0x63, 0x64, 0x62, 0x2f, 0x64, 0x6f, 0x63, 0x64, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x79, 0x62, 0x2e, 0x64, 0x6f, 0x63, 0x64, 0x62,
	0x1a, 0x16, 0x79, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x79, 0x62, 0x2f, 0x75, 0x74, 0x69,
	0x6c, 0x2f, 0x6f, 0x70, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x0e,
	0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x50, 0x42, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5f, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x06, 0x52, 0x12, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x48, 0x79,
	0x62, 0x72, 0x69, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x71, 0x0a, 0x1a, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x42, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2c, 0x0a,
	0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x06, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x48, 0x79, 0x62, 0x72, 0x69, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xa3, 0x03, 0x0a, 0x14,
	0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x50, 0x42, 0x12, 0x39, 0x0a, 0x0b, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x70, 0x61,
	0x69, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x79, 0x62, 0x2e, 0x64,
	0x6f, 0x63, 0x64, 0x62, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69,
	0x72, 0x50, 0x42, 0x52, 0x0a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x50, 0x61, 0x69, 0x72, 0x73, 0x12,
	0x3b, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x79, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x50, 0x42, 0x52,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x1c,
	0x44, 0x45, 0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x6d, 0x61, 0x79, 0x5f, 0x68,
	0x61, 0x76, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x19, 0x44, 0x45, 0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x4d, 0x61,
	0x79, 0x48, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x37, 0x0a,
	0x0a, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x79, 0x62, 0x2e, 0x64, 0x6f, 0x63, 0x64, 0x62, 0x2e, 0x4b, 0x65, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x50, 0x42, 0x52, 0x09, 0x72, 0x65, 0x61,
	0x64, 0x50, 0x61, 0x69, 0x72, 0x73, 0x12, 0x33, 0x0a, 0x0d, 0x72, 0x6f, 0x77, 0x5f, 0x6d, 0x61,
	0x72, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e,
	0x79, 0x62, 0x2e, 0x52, 0x6f, 0x77, 0x4d, 0x61, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b,
	0x72, 0x6f, 0x77, 0x4d, 0x61, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x64, 0x0a, 0x1b, 0x61,
	0x70, 0x70, 0x6c, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x79, 0x62, 0x2e, 0x64, 0x6f, 0x63, 0x64, 0x62, 0x2e, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x42, 0x52, 0x19, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0xac, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x46,
	0x72, 0x6f, 0x6e, 0x74, 0x69, 0x65, 0x72, 0x50, 0x42, 0x12, 0x1f, 0x0a, 0x05, 0x6f, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x79, 0x62, 0x2e, 0x4f, 0x70,
	0x49, 0x64, 0x50, 0x42, 0x52, 0x04, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x79,
	0x62, 0x72, 0x69, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x06, 0x52,
	0x0a, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x63, 0x75, 0x74, 0x6f, 0x66, 0x66, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x06, 0x52, 0x0d, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x75, 0x74, 0x6f,
	0x66, 0x66, 0x12, 0x2c, 0x0a, 0x12, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x06, 0x52, 0x10,
	0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x22, 0x63, 0x0a, 0x17, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x42, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x19, 0x0a,
	0x08, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x77, 0x72, 0x69, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x5f, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x06, 0x52, 0x08, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x48, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x2e, 0x79, 0x62, 0x2e,
	0x64, 0x6f, 0x63, 0x64, 0x62,
}

var (
	file_yb_docdb_docdb_proto_rawDescOnce sync.Once
	file_yb_docdb_docdb_proto_rawDescData = file_yb_docdb_docdb_proto_rawDesc
)

func file_yb_docdb_docdb_proto_rawDescGZIP() []byte {
	file_yb_docdb_docdb_proto_rawDescOnce.Do(func() {
		file_yb_docdb_docdb_proto_rawDescData = protoimpl.X.CompressGZIP(file_yb_docdb_docdb_proto_rawDescData)
	})
	return file_yb_docdb_docdb_proto_rawDescData
}

var file_yb_docdb_docdb_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yb_docdb_docdb_proto_goTypes = []interface{}{
	(*KeyValuePairPB)(nil),             // 0: yb.docdb.KeyValuePairPB
	(*ApplyExternalTransactionPB)(nil), // 1: yb.docdb.ApplyExternalTransactionPB
	(*KeyValueWriteBatchPB)(nil),       // 2: yb.docdb.KeyValueWriteBatchPB
	(*ConsensusFrontierPB)(nil),        // 3: yb.docdb.ConsensusFrontierPB
	(*ApplyTransactionStatePB)(nil),    // 4: yb.docdb.ApplyTransactionStatePB
	(*TransactionMetadataPB)(nil),      // 5: yb.TransactionMetadataPB
	(RowMarkType)(0),                   // 6: yb.RowMarkType
	(*OpIdPB)(nil),                     // 7: yb.OpIdPB
}
var file_yb_docdb_docdb_proto_depIdxs = []int32{
	0, // 0: yb.docdb.KeyValueWriteBatchPB.write_pairs:type_name -> yb.docdb.KeyValuePairPB
	5, // 1: yb.docdb.KeyValueWriteBatchPB.transaction:type_name -> yb.TransactionMetadataPB
	0, // 2: yb.docdb.KeyValueWriteBatchPB.read_pairs:type_name -> yb.docdb.KeyValuePairPB
	6, // 3: yb.docdb.KeyValueWriteBatchPB.row_mark_type:type_name -> yb.RowMarkType
	1, // 4: yb.docdb.KeyValueWriteBatchPB.apply_external_transactions:type_name -> yb.docdb.ApplyExternalTransactionPB
	7, // 5: yb.docdb.ConsensusFrontierPB.op_id:type_name -> yb.OpIdPB
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_yb_docdb_docdb_proto_init() }
func file_yb_docdb_docdb_proto_init() {
	if File_yb_docdb_docdb_proto != nil {
		return
	}
	file_yb_common_common_proto_init()
	file_yb_util_opid_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yb_docdb_docdb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValuePairPB); i {
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
		file_yb_docdb_docdb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyExternalTransactionPB); i {
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
		file_yb_docdb_docdb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValueWriteBatchPB); i {
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
		file_yb_docdb_docdb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsensusFrontierPB); i {
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
		file_yb_docdb_docdb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyTransactionStatePB); i {
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
			RawDescriptor: file_yb_docdb_docdb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yb_docdb_docdb_proto_goTypes,
		DependencyIndexes: file_yb_docdb_docdb_proto_depIdxs,
		MessageInfos:      file_yb_docdb_docdb_proto_msgTypes,
	}.Build()
	File_yb_docdb_docdb_proto = out.File
	file_yb_docdb_docdb_proto_rawDesc = nil
	file_yb_docdb_docdb_proto_goTypes = nil
	file_yb_docdb_docdb_proto_depIdxs = nil
}
