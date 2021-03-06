// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: yb/tserver/tserver_service.proto

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

// TabletServerServiceClient is the client API for TabletServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TabletServerServiceClient interface {
	Write(ctx context.Context, in *WriteRequestPB, opts ...grpc.CallOption) (*WriteResponsePB, error)
	Read(ctx context.Context, in *ReadRequestPB, opts ...grpc.CallOption) (*ReadResponsePB, error)
	VerifyTableRowRange(ctx context.Context, in *VerifyTableRowRangeRequestPB, opts ...grpc.CallOption) (*VerifyTableRowRangeResponsePB, error)
	NoOp(ctx context.Context, in *NoOpRequestPB, opts ...grpc.CallOption) (*NoOpResponsePB, error)
	ListTablets(ctx context.Context, in *ListTabletsRequestPB, opts ...grpc.CallOption) (*ListTabletsResponsePB, error)
	GetLogLocation(ctx context.Context, in *GetLogLocationRequestPB, opts ...grpc.CallOption) (*GetLogLocationResponsePB, error)
	// Run full-scan data checksum on a tablet to verify data integrity.
	//
	// TODO: Consider refactoring this as a scan that runs a checksum aggregation
	// function.
	Checksum(ctx context.Context, in *ChecksumRequestPB, opts ...grpc.CallOption) (*ChecksumResponsePB, error)
	ListTabletsForTabletServer(ctx context.Context, in *ListTabletsForTabletServerRequestPB, opts ...grpc.CallOption) (*ListTabletsForTabletServerResponsePB, error)
	ImportData(ctx context.Context, in *ImportDataRequestPB, opts ...grpc.CallOption) (*ImportDataResponsePB, error)
	UpdateTransaction(ctx context.Context, in *UpdateTransactionRequestPB, opts ...grpc.CallOption) (*UpdateTransactionResponsePB, error)
	// Returns transaction status at coordinator, i.e. PENDING, ABORTED, COMMITTED etc.
	GetTransactionStatus(ctx context.Context, in *GetTransactionStatusRequestPB, opts ...grpc.CallOption) (*GetTransactionStatusResponsePB, error)
	// Returns transaction status at participant, i.e. number of replicated batches or whether it was
	// aborted.
	GetTransactionStatusAtParticipant(ctx context.Context, in *GetTransactionStatusAtParticipantRequestPB, opts ...grpc.CallOption) (*GetTransactionStatusAtParticipantResponsePB, error)
	AbortTransaction(ctx context.Context, in *AbortTransactionRequestPB, opts ...grpc.CallOption) (*AbortTransactionResponsePB, error)
	Truncate(ctx context.Context, in *TruncateRequestPB, opts ...grpc.CallOption) (*TruncateResponsePB, error)
	GetTabletStatus(ctx context.Context, in *GetTabletStatusRequestPB, opts ...grpc.CallOption) (*GetTabletStatusResponsePB, error)
	GetMasterAddresses(ctx context.Context, in *GetMasterAddressesRequestPB, opts ...grpc.CallOption) (*GetMasterAddressesResponsePB, error)
	Publish(ctx context.Context, in *PublishRequestPB, opts ...grpc.CallOption) (*PublishResponsePB, error)
	IsTabletServerReady(ctx context.Context, in *IsTabletServerReadyRequestPB, opts ...grpc.CallOption) (*IsTabletServerReadyResponsePB, error)
	// Takes precreated transaction from this tserver.
	TakeTransaction(ctx context.Context, in *TakeTransactionRequestPB, opts ...grpc.CallOption) (*TakeTransactionResponsePB, error)
	GetSplitKey(ctx context.Context, in *GetSplitKeyRequestPB, opts ...grpc.CallOption) (*GetSplitKeyResponsePB, error)
	GetSharedData(ctx context.Context, in *GetSharedDataRequestPB, opts ...grpc.CallOption) (*GetSharedDataResponsePB, error)
}

type tabletServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTabletServerServiceClient(cc grpc.ClientConnInterface) TabletServerServiceClient {
	return &tabletServerServiceClient{cc}
}

func (c *tabletServerServiceClient) Write(ctx context.Context, in *WriteRequestPB, opts ...grpc.CallOption) (*WriteResponsePB, error) {
	out := new(WriteResponsePB)
	err := c.cc.Invoke(ctx, "/yb.tserver.TabletServerService/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tabletServerServiceClient) Read(ctx context.Context, in *ReadRequestPB, opts ...grpc.CallOption) (*ReadResponsePB, error) {
	out := new(ReadResponsePB)
	err := c.cc.Invoke(ctx, "/yb.tserver.TabletServerService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tabletServerServiceClient) VerifyTableRowRange(ctx context.Context, in *VerifyTableRowRangeRequestPB, opts ...grpc.CallOption) (*VerifyTableRowRangeResponsePB, error) {
	out := new(VerifyTableRowRangeResponsePB)
	err := c.cc.Invoke(ctx, "/yb.tserver.TabletServerService/VerifyTableRowRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tabletServerServiceClient) NoOp(ctx context.Context, in *NoOpRequestPB, opts ...grpc.CallOption) (*NoOpResponsePB, error) {
	out := new(NoOpResponsePB)
	err := c.cc.Invoke(ctx, "/yb.tserver.TabletServerService/NoOp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tabletServerServiceClient) ListTablets(ctx context.Context, in *ListTabletsRequestPB, opts ...grpc.CallOption) (*ListTabletsResponsePB, error) {
	out := new(ListTabletsResponsePB)
	err := c.cc.Invoke(ctx, "/yb.tserver.TabletServerService/ListTablets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tabletServerServiceClient) GetLogLocation(ctx context.Context, in *GetLogLocationRequestPB, opts ...grpc.CallOption) (*GetLogLocationResponsePB, error) {
	out := new(GetLogLocationResponsePB)
	err := c.cc.Invoke(ctx, "/yb.tserver.TabletServerService/GetLogLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tabletServerServiceClient) Checksum(ctx context.Context, in *ChecksumRequestPB, opts ...grpc.CallOption) (*ChecksumResponsePB, error) {
	out := new(ChecksumResponsePB)
	err := c.cc.Invoke(ctx, "/yb.tserver.TabletServerService/Checksum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tabletServerServiceClient) ListTabletsForTabletServer(ctx context.Context, in *ListTabletsForTabletServerRequestPB, opts ...grpc.CallOption) (*ListTabletsForTabletServerResponsePB, error) {
	out := new(ListTabletsForTabletServerResponsePB)
	err := c.cc.Invoke(ctx, "/yb.tserver.TabletServerService/ListTabletsForTabletServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tabletServerServiceClient) ImportData(ctx context.Context, in *ImportDataRequestPB, opts ...grpc.CallOption) (*ImportDataResponsePB, error) {
	out := new(ImportDataResponsePB)
	err := c.cc.Invoke(ctx, "/yb.tserver.TabletServerService/ImportData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tabletServerServiceClient) UpdateTransaction(ctx context.Context, in *UpdateTransactionRequestPB, opts ...grpc.CallOption) (*UpdateTransactionResponsePB, error) {
	out := new(UpdateTransactionResponsePB)
	err := c.cc.Invoke(ctx, "/yb.tserver.TabletServerService/UpdateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tabletServerServiceClient) GetTransactionStatus(ctx context.Context, in *GetTransactionStatusRequestPB, opts ...grpc.CallOption) (*GetTransactionStatusResponsePB, error) {
	out := new(GetTransactionStatusResponsePB)
	err := c.cc.Invoke(ctx, "/yb.tserver.TabletServerService/GetTransactionStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tabletServerServiceClient) GetTransactionStatusAtParticipant(ctx context.Context, in *GetTransactionStatusAtParticipantRequestPB, opts ...grpc.CallOption) (*GetTransactionStatusAtParticipantResponsePB, error) {
	out := new(GetTransactionStatusAtParticipantResponsePB)
	err := c.cc.Invoke(ctx, "/yb.tserver.TabletServerService/GetTransactionStatusAtParticipant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tabletServerServiceClient) AbortTransaction(ctx context.Context, in *AbortTransactionRequestPB, opts ...grpc.CallOption) (*AbortTransactionResponsePB, error) {
	out := new(AbortTransactionResponsePB)
	err := c.cc.Invoke(ctx, "/yb.tserver.TabletServerService/AbortTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tabletServerServiceClient) Truncate(ctx context.Context, in *TruncateRequestPB, opts ...grpc.CallOption) (*TruncateResponsePB, error) {
	out := new(TruncateResponsePB)
	err := c.cc.Invoke(ctx, "/yb.tserver.TabletServerService/Truncate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tabletServerServiceClient) GetTabletStatus(ctx context.Context, in *GetTabletStatusRequestPB, opts ...grpc.CallOption) (*GetTabletStatusResponsePB, error) {
	out := new(GetTabletStatusResponsePB)
	err := c.cc.Invoke(ctx, "/yb.tserver.TabletServerService/GetTabletStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tabletServerServiceClient) GetMasterAddresses(ctx context.Context, in *GetMasterAddressesRequestPB, opts ...grpc.CallOption) (*GetMasterAddressesResponsePB, error) {
	out := new(GetMasterAddressesResponsePB)
	err := c.cc.Invoke(ctx, "/yb.tserver.TabletServerService/GetMasterAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tabletServerServiceClient) Publish(ctx context.Context, in *PublishRequestPB, opts ...grpc.CallOption) (*PublishResponsePB, error) {
	out := new(PublishResponsePB)
	err := c.cc.Invoke(ctx, "/yb.tserver.TabletServerService/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tabletServerServiceClient) IsTabletServerReady(ctx context.Context, in *IsTabletServerReadyRequestPB, opts ...grpc.CallOption) (*IsTabletServerReadyResponsePB, error) {
	out := new(IsTabletServerReadyResponsePB)
	err := c.cc.Invoke(ctx, "/yb.tserver.TabletServerService/IsTabletServerReady", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tabletServerServiceClient) TakeTransaction(ctx context.Context, in *TakeTransactionRequestPB, opts ...grpc.CallOption) (*TakeTransactionResponsePB, error) {
	out := new(TakeTransactionResponsePB)
	err := c.cc.Invoke(ctx, "/yb.tserver.TabletServerService/TakeTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tabletServerServiceClient) GetSplitKey(ctx context.Context, in *GetSplitKeyRequestPB, opts ...grpc.CallOption) (*GetSplitKeyResponsePB, error) {
	out := new(GetSplitKeyResponsePB)
	err := c.cc.Invoke(ctx, "/yb.tserver.TabletServerService/GetSplitKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tabletServerServiceClient) GetSharedData(ctx context.Context, in *GetSharedDataRequestPB, opts ...grpc.CallOption) (*GetSharedDataResponsePB, error) {
	out := new(GetSharedDataResponsePB)
	err := c.cc.Invoke(ctx, "/yb.tserver.TabletServerService/GetSharedData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TabletServerServiceServer is the server API for TabletServerService service.
// All implementations should embed UnimplementedTabletServerServiceServer
// for forward compatibility
type TabletServerServiceServer interface {
	Write(context.Context, *WriteRequestPB) (*WriteResponsePB, error)
	Read(context.Context, *ReadRequestPB) (*ReadResponsePB, error)
	VerifyTableRowRange(context.Context, *VerifyTableRowRangeRequestPB) (*VerifyTableRowRangeResponsePB, error)
	NoOp(context.Context, *NoOpRequestPB) (*NoOpResponsePB, error)
	ListTablets(context.Context, *ListTabletsRequestPB) (*ListTabletsResponsePB, error)
	GetLogLocation(context.Context, *GetLogLocationRequestPB) (*GetLogLocationResponsePB, error)
	// Run full-scan data checksum on a tablet to verify data integrity.
	//
	// TODO: Consider refactoring this as a scan that runs a checksum aggregation
	// function.
	Checksum(context.Context, *ChecksumRequestPB) (*ChecksumResponsePB, error)
	ListTabletsForTabletServer(context.Context, *ListTabletsForTabletServerRequestPB) (*ListTabletsForTabletServerResponsePB, error)
	ImportData(context.Context, *ImportDataRequestPB) (*ImportDataResponsePB, error)
	UpdateTransaction(context.Context, *UpdateTransactionRequestPB) (*UpdateTransactionResponsePB, error)
	// Returns transaction status at coordinator, i.e. PENDING, ABORTED, COMMITTED etc.
	GetTransactionStatus(context.Context, *GetTransactionStatusRequestPB) (*GetTransactionStatusResponsePB, error)
	// Returns transaction status at participant, i.e. number of replicated batches or whether it was
	// aborted.
	GetTransactionStatusAtParticipant(context.Context, *GetTransactionStatusAtParticipantRequestPB) (*GetTransactionStatusAtParticipantResponsePB, error)
	AbortTransaction(context.Context, *AbortTransactionRequestPB) (*AbortTransactionResponsePB, error)
	Truncate(context.Context, *TruncateRequestPB) (*TruncateResponsePB, error)
	GetTabletStatus(context.Context, *GetTabletStatusRequestPB) (*GetTabletStatusResponsePB, error)
	GetMasterAddresses(context.Context, *GetMasterAddressesRequestPB) (*GetMasterAddressesResponsePB, error)
	Publish(context.Context, *PublishRequestPB) (*PublishResponsePB, error)
	IsTabletServerReady(context.Context, *IsTabletServerReadyRequestPB) (*IsTabletServerReadyResponsePB, error)
	// Takes precreated transaction from this tserver.
	TakeTransaction(context.Context, *TakeTransactionRequestPB) (*TakeTransactionResponsePB, error)
	GetSplitKey(context.Context, *GetSplitKeyRequestPB) (*GetSplitKeyResponsePB, error)
	GetSharedData(context.Context, *GetSharedDataRequestPB) (*GetSharedDataResponsePB, error)
}

// UnimplementedTabletServerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTabletServerServiceServer struct {
}

func (UnimplementedTabletServerServiceServer) Write(context.Context, *WriteRequestPB) (*WriteResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedTabletServerServiceServer) Read(context.Context, *ReadRequestPB) (*ReadResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedTabletServerServiceServer) VerifyTableRowRange(context.Context, *VerifyTableRowRangeRequestPB) (*VerifyTableRowRangeResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyTableRowRange not implemented")
}
func (UnimplementedTabletServerServiceServer) NoOp(context.Context, *NoOpRequestPB) (*NoOpResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NoOp not implemented")
}
func (UnimplementedTabletServerServiceServer) ListTablets(context.Context, *ListTabletsRequestPB) (*ListTabletsResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTablets not implemented")
}
func (UnimplementedTabletServerServiceServer) GetLogLocation(context.Context, *GetLogLocationRequestPB) (*GetLogLocationResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogLocation not implemented")
}
func (UnimplementedTabletServerServiceServer) Checksum(context.Context, *ChecksumRequestPB) (*ChecksumResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Checksum not implemented")
}
func (UnimplementedTabletServerServiceServer) ListTabletsForTabletServer(context.Context, *ListTabletsForTabletServerRequestPB) (*ListTabletsForTabletServerResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTabletsForTabletServer not implemented")
}
func (UnimplementedTabletServerServiceServer) ImportData(context.Context, *ImportDataRequestPB) (*ImportDataResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportData not implemented")
}
func (UnimplementedTabletServerServiceServer) UpdateTransaction(context.Context, *UpdateTransactionRequestPB) (*UpdateTransactionResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTransaction not implemented")
}
func (UnimplementedTabletServerServiceServer) GetTransactionStatus(context.Context, *GetTransactionStatusRequestPB) (*GetTransactionStatusResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionStatus not implemented")
}
func (UnimplementedTabletServerServiceServer) GetTransactionStatusAtParticipant(context.Context, *GetTransactionStatusAtParticipantRequestPB) (*GetTransactionStatusAtParticipantResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionStatusAtParticipant not implemented")
}
func (UnimplementedTabletServerServiceServer) AbortTransaction(context.Context, *AbortTransactionRequestPB) (*AbortTransactionResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AbortTransaction not implemented")
}
func (UnimplementedTabletServerServiceServer) Truncate(context.Context, *TruncateRequestPB) (*TruncateResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Truncate not implemented")
}
func (UnimplementedTabletServerServiceServer) GetTabletStatus(context.Context, *GetTabletStatusRequestPB) (*GetTabletStatusResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTabletStatus not implemented")
}
func (UnimplementedTabletServerServiceServer) GetMasterAddresses(context.Context, *GetMasterAddressesRequestPB) (*GetMasterAddressesResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMasterAddresses not implemented")
}
func (UnimplementedTabletServerServiceServer) Publish(context.Context, *PublishRequestPB) (*PublishResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedTabletServerServiceServer) IsTabletServerReady(context.Context, *IsTabletServerReadyRequestPB) (*IsTabletServerReadyResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsTabletServerReady not implemented")
}
func (UnimplementedTabletServerServiceServer) TakeTransaction(context.Context, *TakeTransactionRequestPB) (*TakeTransactionResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TakeTransaction not implemented")
}
func (UnimplementedTabletServerServiceServer) GetSplitKey(context.Context, *GetSplitKeyRequestPB) (*GetSplitKeyResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSplitKey not implemented")
}
func (UnimplementedTabletServerServiceServer) GetSharedData(context.Context, *GetSharedDataRequestPB) (*GetSharedDataResponsePB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSharedData not implemented")
}

// UnsafeTabletServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TabletServerServiceServer will
// result in compilation errors.
type UnsafeTabletServerServiceServer interface {
	mustEmbedUnimplementedTabletServerServiceServer()
}

func RegisterTabletServerServiceServer(s grpc.ServiceRegistrar, srv TabletServerServiceServer) {
	s.RegisterService(&TabletServerService_ServiceDesc, srv)
}

func _TabletServerService_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TabletServerServiceServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.tserver.TabletServerService/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TabletServerServiceServer).Write(ctx, req.(*WriteRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _TabletServerService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TabletServerServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.tserver.TabletServerService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TabletServerServiceServer).Read(ctx, req.(*ReadRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _TabletServerService_VerifyTableRowRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTableRowRangeRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TabletServerServiceServer).VerifyTableRowRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.tserver.TabletServerService/VerifyTableRowRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TabletServerServiceServer).VerifyTableRowRange(ctx, req.(*VerifyTableRowRangeRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _TabletServerService_NoOp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoOpRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TabletServerServiceServer).NoOp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.tserver.TabletServerService/NoOp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TabletServerServiceServer).NoOp(ctx, req.(*NoOpRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _TabletServerService_ListTablets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTabletsRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TabletServerServiceServer).ListTablets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.tserver.TabletServerService/ListTablets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TabletServerServiceServer).ListTablets(ctx, req.(*ListTabletsRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _TabletServerService_GetLogLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogLocationRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TabletServerServiceServer).GetLogLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.tserver.TabletServerService/GetLogLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TabletServerServiceServer).GetLogLocation(ctx, req.(*GetLogLocationRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _TabletServerService_Checksum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChecksumRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TabletServerServiceServer).Checksum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.tserver.TabletServerService/Checksum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TabletServerServiceServer).Checksum(ctx, req.(*ChecksumRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _TabletServerService_ListTabletsForTabletServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTabletsForTabletServerRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TabletServerServiceServer).ListTabletsForTabletServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.tserver.TabletServerService/ListTabletsForTabletServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TabletServerServiceServer).ListTabletsForTabletServer(ctx, req.(*ListTabletsForTabletServerRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _TabletServerService_ImportData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportDataRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TabletServerServiceServer).ImportData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.tserver.TabletServerService/ImportData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TabletServerServiceServer).ImportData(ctx, req.(*ImportDataRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _TabletServerService_UpdateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTransactionRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TabletServerServiceServer).UpdateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.tserver.TabletServerService/UpdateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TabletServerServiceServer).UpdateTransaction(ctx, req.(*UpdateTransactionRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _TabletServerService_GetTransactionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionStatusRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TabletServerServiceServer).GetTransactionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.tserver.TabletServerService/GetTransactionStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TabletServerServiceServer).GetTransactionStatus(ctx, req.(*GetTransactionStatusRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _TabletServerService_GetTransactionStatusAtParticipant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionStatusAtParticipantRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TabletServerServiceServer).GetTransactionStatusAtParticipant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.tserver.TabletServerService/GetTransactionStatusAtParticipant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TabletServerServiceServer).GetTransactionStatusAtParticipant(ctx, req.(*GetTransactionStatusAtParticipantRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _TabletServerService_AbortTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AbortTransactionRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TabletServerServiceServer).AbortTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.tserver.TabletServerService/AbortTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TabletServerServiceServer).AbortTransaction(ctx, req.(*AbortTransactionRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _TabletServerService_Truncate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TruncateRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TabletServerServiceServer).Truncate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.tserver.TabletServerService/Truncate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TabletServerServiceServer).Truncate(ctx, req.(*TruncateRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _TabletServerService_GetTabletStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTabletStatusRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TabletServerServiceServer).GetTabletStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.tserver.TabletServerService/GetTabletStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TabletServerServiceServer).GetTabletStatus(ctx, req.(*GetTabletStatusRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _TabletServerService_GetMasterAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMasterAddressesRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TabletServerServiceServer).GetMasterAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.tserver.TabletServerService/GetMasterAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TabletServerServiceServer).GetMasterAddresses(ctx, req.(*GetMasterAddressesRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _TabletServerService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TabletServerServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.tserver.TabletServerService/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TabletServerServiceServer).Publish(ctx, req.(*PublishRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _TabletServerService_IsTabletServerReady_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsTabletServerReadyRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TabletServerServiceServer).IsTabletServerReady(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.tserver.TabletServerService/IsTabletServerReady",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TabletServerServiceServer).IsTabletServerReady(ctx, req.(*IsTabletServerReadyRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _TabletServerService_TakeTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TakeTransactionRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TabletServerServiceServer).TakeTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.tserver.TabletServerService/TakeTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TabletServerServiceServer).TakeTransaction(ctx, req.(*TakeTransactionRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _TabletServerService_GetSplitKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSplitKeyRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TabletServerServiceServer).GetSplitKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.tserver.TabletServerService/GetSplitKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TabletServerServiceServer).GetSplitKey(ctx, req.(*GetSplitKeyRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

func _TabletServerService_GetSharedData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSharedDataRequestPB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TabletServerServiceServer).GetSharedData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yb.tserver.TabletServerService/GetSharedData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TabletServerServiceServer).GetSharedData(ctx, req.(*GetSharedDataRequestPB))
	}
	return interceptor(ctx, in, info, handler)
}

// TabletServerService_ServiceDesc is the grpc.ServiceDesc for TabletServerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TabletServerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yb.tserver.TabletServerService",
	HandlerType: (*TabletServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _TabletServerService_Write_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _TabletServerService_Read_Handler,
		},
		{
			MethodName: "VerifyTableRowRange",
			Handler:    _TabletServerService_VerifyTableRowRange_Handler,
		},
		{
			MethodName: "NoOp",
			Handler:    _TabletServerService_NoOp_Handler,
		},
		{
			MethodName: "ListTablets",
			Handler:    _TabletServerService_ListTablets_Handler,
		},
		{
			MethodName: "GetLogLocation",
			Handler:    _TabletServerService_GetLogLocation_Handler,
		},
		{
			MethodName: "Checksum",
			Handler:    _TabletServerService_Checksum_Handler,
		},
		{
			MethodName: "ListTabletsForTabletServer",
			Handler:    _TabletServerService_ListTabletsForTabletServer_Handler,
		},
		{
			MethodName: "ImportData",
			Handler:    _TabletServerService_ImportData_Handler,
		},
		{
			MethodName: "UpdateTransaction",
			Handler:    _TabletServerService_UpdateTransaction_Handler,
		},
		{
			MethodName: "GetTransactionStatus",
			Handler:    _TabletServerService_GetTransactionStatus_Handler,
		},
		{
			MethodName: "GetTransactionStatusAtParticipant",
			Handler:    _TabletServerService_GetTransactionStatusAtParticipant_Handler,
		},
		{
			MethodName: "AbortTransaction",
			Handler:    _TabletServerService_AbortTransaction_Handler,
		},
		{
			MethodName: "Truncate",
			Handler:    _TabletServerService_Truncate_Handler,
		},
		{
			MethodName: "GetTabletStatus",
			Handler:    _TabletServerService_GetTabletStatus_Handler,
		},
		{
			MethodName: "GetMasterAddresses",
			Handler:    _TabletServerService_GetMasterAddresses_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _TabletServerService_Publish_Handler,
		},
		{
			MethodName: "IsTabletServerReady",
			Handler:    _TabletServerService_IsTabletServerReady_Handler,
		},
		{
			MethodName: "TakeTransaction",
			Handler:    _TabletServerService_TakeTransaction_Handler,
		},
		{
			MethodName: "GetSplitKey",
			Handler:    _TabletServerService_GetSplitKey_Handler,
		},
		{
			MethodName: "GetSharedData",
			Handler:    _TabletServerService_GetSharedData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yb/tserver/tserver_service.proto",
}
