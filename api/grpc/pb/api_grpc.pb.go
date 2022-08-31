// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api.proto

package pb

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

// Payment_APIClient is the client API for Payment_API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Payment_APIClient interface {
	GetConfig(ctx context.Context, in *GetConfigReq, opts ...grpc.CallOption) (*GetConfigResp, error)
	OpenSession(ctx context.Context, in *OpenSessionReq, opts ...grpc.CallOption) (*OpenSessionResp, error)
	Time(ctx context.Context, in *TimeReq, opts ...grpc.CallOption) (*TimeResp, error)
	RegisterCurrency(ctx context.Context, in *RegisterCurrencyReq, opts ...grpc.CallOption) (*RegisterCurrencyResp, error)
	Help(ctx context.Context, in *HelpReq, opts ...grpc.CallOption) (*HelpResp, error)
	AddPeerID(ctx context.Context, in *AddPeerIDReq, opts ...grpc.CallOption) (*AddPeerIDResp, error)
	GetPeerID(ctx context.Context, in *GetPeerIDReq, opts ...grpc.CallOption) (*GetPeerIDResp, error)
	OpenPayCh(ctx context.Context, in *OpenPayChReq, opts ...grpc.CallOption) (*OpenPayChResp, error)
	GetPayChsInfo(ctx context.Context, in *GetPayChsInfoReq, opts ...grpc.CallOption) (*GetPayChsInfoResp, error)
	SubPayChProposals(ctx context.Context, in *SubPayChProposalsReq, opts ...grpc.CallOption) (Payment_API_SubPayChProposalsClient, error)
	UnsubPayChProposals(ctx context.Context, in *UnsubPayChProposalsReq, opts ...grpc.CallOption) (*UnsubPayChProposalsResp, error)
	RespondPayChProposal(ctx context.Context, in *RespondPayChProposalReq, opts ...grpc.CallOption) (*RespondPayChProposalResp, error)
	CloseSession(ctx context.Context, in *CloseSessionReq, opts ...grpc.CallOption) (*CloseSessionResp, error)
	DeployAssetERC20(ctx context.Context, in *DeployAssetERC20Req, opts ...grpc.CallOption) (*DeployAssetERC20Resp, error)
	SendPayChUpdate(ctx context.Context, in *SendPayChUpdateReq, opts ...grpc.CallOption) (*SendPayChUpdateResp, error)
	SubPayChUpdates(ctx context.Context, in *SubpayChUpdatesReq, opts ...grpc.CallOption) (Payment_API_SubPayChUpdatesClient, error)
	UnsubPayChUpdates(ctx context.Context, in *UnsubPayChUpdatesReq, opts ...grpc.CallOption) (*UnsubPayChUpdatesResp, error)
	RespondPayChUpdate(ctx context.Context, in *RespondPayChUpdateReq, opts ...grpc.CallOption) (*RespondPayChUpdateResp, error)
	GetPayChInfo(ctx context.Context, in *GetPayChInfoReq, opts ...grpc.CallOption) (*GetPayChInfoResp, error)
	ClosePayCh(ctx context.Context, in *ClosePayChReq, opts ...grpc.CallOption) (*ClosePayChResp, error)
}

type payment_APIClient struct {
	cc grpc.ClientConnInterface
}

func NewPayment_APIClient(cc grpc.ClientConnInterface) Payment_APIClient {
	return &payment_APIClient{cc}
}

func (c *payment_APIClient) GetConfig(ctx context.Context, in *GetConfigReq, opts ...grpc.CallOption) (*GetConfigResp, error) {
	out := new(GetConfigResp)
	err := c.cc.Invoke(ctx, "/pb.Payment_API/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payment_APIClient) OpenSession(ctx context.Context, in *OpenSessionReq, opts ...grpc.CallOption) (*OpenSessionResp, error) {
	out := new(OpenSessionResp)
	err := c.cc.Invoke(ctx, "/pb.Payment_API/OpenSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payment_APIClient) Time(ctx context.Context, in *TimeReq, opts ...grpc.CallOption) (*TimeResp, error) {
	out := new(TimeResp)
	err := c.cc.Invoke(ctx, "/pb.Payment_API/Time", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payment_APIClient) RegisterCurrency(ctx context.Context, in *RegisterCurrencyReq, opts ...grpc.CallOption) (*RegisterCurrencyResp, error) {
	out := new(RegisterCurrencyResp)
	err := c.cc.Invoke(ctx, "/pb.Payment_API/RegisterCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payment_APIClient) Help(ctx context.Context, in *HelpReq, opts ...grpc.CallOption) (*HelpResp, error) {
	out := new(HelpResp)
	err := c.cc.Invoke(ctx, "/pb.Payment_API/Help", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payment_APIClient) AddPeerID(ctx context.Context, in *AddPeerIDReq, opts ...grpc.CallOption) (*AddPeerIDResp, error) {
	out := new(AddPeerIDResp)
	err := c.cc.Invoke(ctx, "/pb.Payment_API/AddPeerID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payment_APIClient) GetPeerID(ctx context.Context, in *GetPeerIDReq, opts ...grpc.CallOption) (*GetPeerIDResp, error) {
	out := new(GetPeerIDResp)
	err := c.cc.Invoke(ctx, "/pb.Payment_API/GetPeerID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payment_APIClient) OpenPayCh(ctx context.Context, in *OpenPayChReq, opts ...grpc.CallOption) (*OpenPayChResp, error) {
	out := new(OpenPayChResp)
	err := c.cc.Invoke(ctx, "/pb.Payment_API/OpenPayCh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payment_APIClient) GetPayChsInfo(ctx context.Context, in *GetPayChsInfoReq, opts ...grpc.CallOption) (*GetPayChsInfoResp, error) {
	out := new(GetPayChsInfoResp)
	err := c.cc.Invoke(ctx, "/pb.Payment_API/GetPayChsInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payment_APIClient) SubPayChProposals(ctx context.Context, in *SubPayChProposalsReq, opts ...grpc.CallOption) (Payment_API_SubPayChProposalsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Payment_API_ServiceDesc.Streams[0], "/pb.Payment_API/SubPayChProposals", opts...)
	if err != nil {
		return nil, err
	}
	x := &payment_APISubPayChProposalsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Payment_API_SubPayChProposalsClient interface {
	Recv() (*SubPayChProposalsResp, error)
	grpc.ClientStream
}

type payment_APISubPayChProposalsClient struct {
	grpc.ClientStream
}

func (x *payment_APISubPayChProposalsClient) Recv() (*SubPayChProposalsResp, error) {
	m := new(SubPayChProposalsResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *payment_APIClient) UnsubPayChProposals(ctx context.Context, in *UnsubPayChProposalsReq, opts ...grpc.CallOption) (*UnsubPayChProposalsResp, error) {
	out := new(UnsubPayChProposalsResp)
	err := c.cc.Invoke(ctx, "/pb.Payment_API/UnsubPayChProposals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payment_APIClient) RespondPayChProposal(ctx context.Context, in *RespondPayChProposalReq, opts ...grpc.CallOption) (*RespondPayChProposalResp, error) {
	out := new(RespondPayChProposalResp)
	err := c.cc.Invoke(ctx, "/pb.Payment_API/RespondPayChProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payment_APIClient) CloseSession(ctx context.Context, in *CloseSessionReq, opts ...grpc.CallOption) (*CloseSessionResp, error) {
	out := new(CloseSessionResp)
	err := c.cc.Invoke(ctx, "/pb.Payment_API/CloseSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payment_APIClient) DeployAssetERC20(ctx context.Context, in *DeployAssetERC20Req, opts ...grpc.CallOption) (*DeployAssetERC20Resp, error) {
	out := new(DeployAssetERC20Resp)
	err := c.cc.Invoke(ctx, "/pb.Payment_API/DeployAssetERC20", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payment_APIClient) SendPayChUpdate(ctx context.Context, in *SendPayChUpdateReq, opts ...grpc.CallOption) (*SendPayChUpdateResp, error) {
	out := new(SendPayChUpdateResp)
	err := c.cc.Invoke(ctx, "/pb.Payment_API/SendPayChUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payment_APIClient) SubPayChUpdates(ctx context.Context, in *SubpayChUpdatesReq, opts ...grpc.CallOption) (Payment_API_SubPayChUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Payment_API_ServiceDesc.Streams[1], "/pb.Payment_API/SubPayChUpdates", opts...)
	if err != nil {
		return nil, err
	}
	x := &payment_APISubPayChUpdatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Payment_API_SubPayChUpdatesClient interface {
	Recv() (*SubPayChUpdatesResp, error)
	grpc.ClientStream
}

type payment_APISubPayChUpdatesClient struct {
	grpc.ClientStream
}

func (x *payment_APISubPayChUpdatesClient) Recv() (*SubPayChUpdatesResp, error) {
	m := new(SubPayChUpdatesResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *payment_APIClient) UnsubPayChUpdates(ctx context.Context, in *UnsubPayChUpdatesReq, opts ...grpc.CallOption) (*UnsubPayChUpdatesResp, error) {
	out := new(UnsubPayChUpdatesResp)
	err := c.cc.Invoke(ctx, "/pb.Payment_API/UnsubPayChUpdates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payment_APIClient) RespondPayChUpdate(ctx context.Context, in *RespondPayChUpdateReq, opts ...grpc.CallOption) (*RespondPayChUpdateResp, error) {
	out := new(RespondPayChUpdateResp)
	err := c.cc.Invoke(ctx, "/pb.Payment_API/RespondPayChUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payment_APIClient) GetPayChInfo(ctx context.Context, in *GetPayChInfoReq, opts ...grpc.CallOption) (*GetPayChInfoResp, error) {
	out := new(GetPayChInfoResp)
	err := c.cc.Invoke(ctx, "/pb.Payment_API/GetPayChInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payment_APIClient) ClosePayCh(ctx context.Context, in *ClosePayChReq, opts ...grpc.CallOption) (*ClosePayChResp, error) {
	out := new(ClosePayChResp)
	err := c.cc.Invoke(ctx, "/pb.Payment_API/ClosePayCh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Payment_APIServer is the server API for Payment_API service.
// All implementations must embed UnimplementedPayment_APIServer
// for forward compatibility
type Payment_APIServer interface {
	GetConfig(context.Context, *GetConfigReq) (*GetConfigResp, error)
	OpenSession(context.Context, *OpenSessionReq) (*OpenSessionResp, error)
	Time(context.Context, *TimeReq) (*TimeResp, error)
	RegisterCurrency(context.Context, *RegisterCurrencyReq) (*RegisterCurrencyResp, error)
	Help(context.Context, *HelpReq) (*HelpResp, error)
	AddPeerID(context.Context, *AddPeerIDReq) (*AddPeerIDResp, error)
	GetPeerID(context.Context, *GetPeerIDReq) (*GetPeerIDResp, error)
	OpenPayCh(context.Context, *OpenPayChReq) (*OpenPayChResp, error)
	GetPayChsInfo(context.Context, *GetPayChsInfoReq) (*GetPayChsInfoResp, error)
	SubPayChProposals(*SubPayChProposalsReq, Payment_API_SubPayChProposalsServer) error
	UnsubPayChProposals(context.Context, *UnsubPayChProposalsReq) (*UnsubPayChProposalsResp, error)
	RespondPayChProposal(context.Context, *RespondPayChProposalReq) (*RespondPayChProposalResp, error)
	CloseSession(context.Context, *CloseSessionReq) (*CloseSessionResp, error)
	DeployAssetERC20(context.Context, *DeployAssetERC20Req) (*DeployAssetERC20Resp, error)
	SendPayChUpdate(context.Context, *SendPayChUpdateReq) (*SendPayChUpdateResp, error)
	SubPayChUpdates(*SubpayChUpdatesReq, Payment_API_SubPayChUpdatesServer) error
	UnsubPayChUpdates(context.Context, *UnsubPayChUpdatesReq) (*UnsubPayChUpdatesResp, error)
	RespondPayChUpdate(context.Context, *RespondPayChUpdateReq) (*RespondPayChUpdateResp, error)
	GetPayChInfo(context.Context, *GetPayChInfoReq) (*GetPayChInfoResp, error)
	ClosePayCh(context.Context, *ClosePayChReq) (*ClosePayChResp, error)
	mustEmbedUnimplementedPayment_APIServer()
}

// UnimplementedPayment_APIServer must be embedded to have forward compatible implementations.
type UnimplementedPayment_APIServer struct {
}

func (UnimplementedPayment_APIServer) GetConfig(context.Context, *GetConfigReq) (*GetConfigResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedPayment_APIServer) OpenSession(context.Context, *OpenSessionReq) (*OpenSessionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenSession not implemented")
}
func (UnimplementedPayment_APIServer) Time(context.Context, *TimeReq) (*TimeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Time not implemented")
}
func (UnimplementedPayment_APIServer) RegisterCurrency(context.Context, *RegisterCurrencyReq) (*RegisterCurrencyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterCurrency not implemented")
}
func (UnimplementedPayment_APIServer) Help(context.Context, *HelpReq) (*HelpResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Help not implemented")
}
func (UnimplementedPayment_APIServer) AddPeerID(context.Context, *AddPeerIDReq) (*AddPeerIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPeerID not implemented")
}
func (UnimplementedPayment_APIServer) GetPeerID(context.Context, *GetPeerIDReq) (*GetPeerIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeerID not implemented")
}
func (UnimplementedPayment_APIServer) OpenPayCh(context.Context, *OpenPayChReq) (*OpenPayChResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenPayCh not implemented")
}
func (UnimplementedPayment_APIServer) GetPayChsInfo(context.Context, *GetPayChsInfoReq) (*GetPayChsInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayChsInfo not implemented")
}
func (UnimplementedPayment_APIServer) SubPayChProposals(*SubPayChProposalsReq, Payment_API_SubPayChProposalsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubPayChProposals not implemented")
}
func (UnimplementedPayment_APIServer) UnsubPayChProposals(context.Context, *UnsubPayChProposalsReq) (*UnsubPayChProposalsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsubPayChProposals not implemented")
}
func (UnimplementedPayment_APIServer) RespondPayChProposal(context.Context, *RespondPayChProposalReq) (*RespondPayChProposalResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RespondPayChProposal not implemented")
}
func (UnimplementedPayment_APIServer) CloseSession(context.Context, *CloseSessionReq) (*CloseSessionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseSession not implemented")
}
func (UnimplementedPayment_APIServer) DeployAssetERC20(context.Context, *DeployAssetERC20Req) (*DeployAssetERC20Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployAssetERC20 not implemented")
}
func (UnimplementedPayment_APIServer) SendPayChUpdate(context.Context, *SendPayChUpdateReq) (*SendPayChUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPayChUpdate not implemented")
}
func (UnimplementedPayment_APIServer) SubPayChUpdates(*SubpayChUpdatesReq, Payment_API_SubPayChUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubPayChUpdates not implemented")
}
func (UnimplementedPayment_APIServer) UnsubPayChUpdates(context.Context, *UnsubPayChUpdatesReq) (*UnsubPayChUpdatesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsubPayChUpdates not implemented")
}
func (UnimplementedPayment_APIServer) RespondPayChUpdate(context.Context, *RespondPayChUpdateReq) (*RespondPayChUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RespondPayChUpdate not implemented")
}
func (UnimplementedPayment_APIServer) GetPayChInfo(context.Context, *GetPayChInfoReq) (*GetPayChInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayChInfo not implemented")
}
func (UnimplementedPayment_APIServer) ClosePayCh(context.Context, *ClosePayChReq) (*ClosePayChResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClosePayCh not implemented")
}
func (UnimplementedPayment_APIServer) mustEmbedUnimplementedPayment_APIServer() {}

// UnsafePayment_APIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Payment_APIServer will
// result in compilation errors.
type UnsafePayment_APIServer interface {
	mustEmbedUnimplementedPayment_APIServer()
}

func RegisterPayment_APIServer(s grpc.ServiceRegistrar, srv Payment_APIServer) {
	s.RegisterService(&Payment_API_ServiceDesc, srv)
}

func _Payment_API_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Payment_APIServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Payment_API/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Payment_APIServer).GetConfig(ctx, req.(*GetConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_API_OpenSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenSessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Payment_APIServer).OpenSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Payment_API/OpenSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Payment_APIServer).OpenSession(ctx, req.(*OpenSessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_API_Time_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Payment_APIServer).Time(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Payment_API/Time",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Payment_APIServer).Time(ctx, req.(*TimeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_API_RegisterCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterCurrencyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Payment_APIServer).RegisterCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Payment_API/RegisterCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Payment_APIServer).RegisterCurrency(ctx, req.(*RegisterCurrencyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_API_Help_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelpReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Payment_APIServer).Help(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Payment_API/Help",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Payment_APIServer).Help(ctx, req.(*HelpReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_API_AddPeerID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPeerIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Payment_APIServer).AddPeerID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Payment_API/AddPeerID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Payment_APIServer).AddPeerID(ctx, req.(*AddPeerIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_API_GetPeerID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPeerIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Payment_APIServer).GetPeerID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Payment_API/GetPeerID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Payment_APIServer).GetPeerID(ctx, req.(*GetPeerIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_API_OpenPayCh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenPayChReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Payment_APIServer).OpenPayCh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Payment_API/OpenPayCh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Payment_APIServer).OpenPayCh(ctx, req.(*OpenPayChReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_API_GetPayChsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPayChsInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Payment_APIServer).GetPayChsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Payment_API/GetPayChsInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Payment_APIServer).GetPayChsInfo(ctx, req.(*GetPayChsInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_API_SubPayChProposals_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubPayChProposalsReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(Payment_APIServer).SubPayChProposals(m, &payment_APISubPayChProposalsServer{stream})
}

type Payment_API_SubPayChProposalsServer interface {
	Send(*SubPayChProposalsResp) error
	grpc.ServerStream
}

type payment_APISubPayChProposalsServer struct {
	grpc.ServerStream
}

func (x *payment_APISubPayChProposalsServer) Send(m *SubPayChProposalsResp) error {
	return x.ServerStream.SendMsg(m)
}

func _Payment_API_UnsubPayChProposals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubPayChProposalsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Payment_APIServer).UnsubPayChProposals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Payment_API/UnsubPayChProposals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Payment_APIServer).UnsubPayChProposals(ctx, req.(*UnsubPayChProposalsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_API_RespondPayChProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RespondPayChProposalReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Payment_APIServer).RespondPayChProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Payment_API/RespondPayChProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Payment_APIServer).RespondPayChProposal(ctx, req.(*RespondPayChProposalReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_API_CloseSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseSessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Payment_APIServer).CloseSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Payment_API/CloseSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Payment_APIServer).CloseSession(ctx, req.(*CloseSessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_API_DeployAssetERC20_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployAssetERC20Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Payment_APIServer).DeployAssetERC20(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Payment_API/DeployAssetERC20",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Payment_APIServer).DeployAssetERC20(ctx, req.(*DeployAssetERC20Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_API_SendPayChUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendPayChUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Payment_APIServer).SendPayChUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Payment_API/SendPayChUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Payment_APIServer).SendPayChUpdate(ctx, req.(*SendPayChUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_API_SubPayChUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubpayChUpdatesReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(Payment_APIServer).SubPayChUpdates(m, &payment_APISubPayChUpdatesServer{stream})
}

type Payment_API_SubPayChUpdatesServer interface {
	Send(*SubPayChUpdatesResp) error
	grpc.ServerStream
}

type payment_APISubPayChUpdatesServer struct {
	grpc.ServerStream
}

func (x *payment_APISubPayChUpdatesServer) Send(m *SubPayChUpdatesResp) error {
	return x.ServerStream.SendMsg(m)
}

func _Payment_API_UnsubPayChUpdates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubPayChUpdatesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Payment_APIServer).UnsubPayChUpdates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Payment_API/UnsubPayChUpdates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Payment_APIServer).UnsubPayChUpdates(ctx, req.(*UnsubPayChUpdatesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_API_RespondPayChUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RespondPayChUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Payment_APIServer).RespondPayChUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Payment_API/RespondPayChUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Payment_APIServer).RespondPayChUpdate(ctx, req.(*RespondPayChUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_API_GetPayChInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPayChInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Payment_APIServer).GetPayChInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Payment_API/GetPayChInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Payment_APIServer).GetPayChInfo(ctx, req.(*GetPayChInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_API_ClosePayCh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClosePayChReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Payment_APIServer).ClosePayCh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Payment_API/ClosePayCh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Payment_APIServer).ClosePayCh(ctx, req.(*ClosePayChReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Payment_API_ServiceDesc is the grpc.ServiceDesc for Payment_API service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Payment_API_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Payment_API",
	HandlerType: (*Payment_APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfig",
			Handler:    _Payment_API_GetConfig_Handler,
		},
		{
			MethodName: "OpenSession",
			Handler:    _Payment_API_OpenSession_Handler,
		},
		{
			MethodName: "Time",
			Handler:    _Payment_API_Time_Handler,
		},
		{
			MethodName: "RegisterCurrency",
			Handler:    _Payment_API_RegisterCurrency_Handler,
		},
		{
			MethodName: "Help",
			Handler:    _Payment_API_Help_Handler,
		},
		{
			MethodName: "AddPeerID",
			Handler:    _Payment_API_AddPeerID_Handler,
		},
		{
			MethodName: "GetPeerID",
			Handler:    _Payment_API_GetPeerID_Handler,
		},
		{
			MethodName: "OpenPayCh",
			Handler:    _Payment_API_OpenPayCh_Handler,
		},
		{
			MethodName: "GetPayChsInfo",
			Handler:    _Payment_API_GetPayChsInfo_Handler,
		},
		{
			MethodName: "UnsubPayChProposals",
			Handler:    _Payment_API_UnsubPayChProposals_Handler,
		},
		{
			MethodName: "RespondPayChProposal",
			Handler:    _Payment_API_RespondPayChProposal_Handler,
		},
		{
			MethodName: "CloseSession",
			Handler:    _Payment_API_CloseSession_Handler,
		},
		{
			MethodName: "DeployAssetERC20",
			Handler:    _Payment_API_DeployAssetERC20_Handler,
		},
		{
			MethodName: "SendPayChUpdate",
			Handler:    _Payment_API_SendPayChUpdate_Handler,
		},
		{
			MethodName: "UnsubPayChUpdates",
			Handler:    _Payment_API_UnsubPayChUpdates_Handler,
		},
		{
			MethodName: "RespondPayChUpdate",
			Handler:    _Payment_API_RespondPayChUpdate_Handler,
		},
		{
			MethodName: "GetPayChInfo",
			Handler:    _Payment_API_GetPayChInfo_Handler,
		},
		{
			MethodName: "ClosePayCh",
			Handler:    _Payment_API_ClosePayCh_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubPayChProposals",
			Handler:       _Payment_API_SubPayChProposals_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubPayChUpdates",
			Handler:       _Payment_API_SubPayChUpdates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}
