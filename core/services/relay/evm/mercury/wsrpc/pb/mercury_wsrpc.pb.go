// Code generated by protoc-gen-go-wsrpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-wsrpc v0.0.1
// - protoc             v4.25.1

package pb

import (
	context "context"

	wsrpc "github.com/smartcontractkit/wsrpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MercuryClient is the client API for Mercury service.
//
type MercuryClient interface {
	Transmit(ctx context.Context, in *TransmitRequest) (*TransmitResponse, error)
	LatestReport(ctx context.Context, in *LatestReportRequest) (*LatestReportResponse, error)
}

type mercuryClient struct {
	cc wsrpc.ClientInterface
}

type mercuryGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewMercuryClient(cc wsrpc.ClientInterface) MercuryClient {
	return &mercuryClient{cc}
}

func NewMercuryGrpcClient(cc grpc.ClientConnInterface) MercuryClient {
	return &mercuryGrpcClient{cc}
}

func (c *mercuryClient) Transmit(ctx context.Context, in *TransmitRequest) (*TransmitResponse, error) {
	out := new(TransmitResponse)
	err := c.cc.Invoke(ctx, "Transmit", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mercuryGrpcClient) Transmit(ctx context.Context, in *TransmitRequest) (*TransmitResponse, error) {
	out := new(TransmitResponse)
	err := c.cc.Invoke(ctx, "/pb.Mercury/Transmit", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mercuryClient) LatestReport(ctx context.Context, in *LatestReportRequest) (*LatestReportResponse, error) {
	out := new(LatestReportResponse)
	err := c.cc.Invoke(ctx, "LatestReport", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mercuryGrpcClient) LatestReport(ctx context.Context, in *LatestReportRequest) (*LatestReportResponse, error) {
	out := new(LatestReportResponse)
	err := c.cc.Invoke(ctx, "/pb.Mercury/LatestReport", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MercuryServer is the server API for Mercury service.
type MercuryServer interface {
	Transmit(context.Context, *TransmitRequest) (*TransmitResponse, error)
	LatestReport(context.Context, *LatestReportRequest) (*LatestReportResponse, error)
	mustEmbedUnimplementedMercuryServer()
}

// UnimplementedMercuryServer must be embedded to have forward compatible implementations.
type UnimplementedMercuryServer struct {
}

func (UnimplementedMercuryServer) Transmit(context.Context, *TransmitRequest) (*TransmitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transmit not implemented")
}
func (UnimplementedMercuryServer) LatestReport(context.Context, *LatestReportRequest) (*LatestReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LatestReport not implemented")
}
func (UnimplementedMercuryServer) mustEmbedUnimplementedMercuryServer() {}

// UnsafeMercuryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MercuryServer will
// result in compilation errors.
type UnsafeMercuryServer interface {
	mustEmbedUnimplementedMercuryServer()
}

func RegisterMercuryServer(s wsrpc.ServiceRegistrar, srv MercuryServer) {
	s.RegisterService(&Mercury_ServiceDesc, srv)
}

func RegisterGrpcMercuryServer(s grpc.ServiceRegistrar, srv MercuryServer) {
	s.RegisterService(&Mercury_ServiceDesc_Grpc, srv)
}

func _Mercury_Transmit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(TransmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	return srv.(MercuryServer).Transmit(ctx, in)
}

func _Mercury_Transmit_Grpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MercuryServer).Transmit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mercury/Transmit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MercuryServer).Transmit(ctx, req.(*TransmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mercury_LatestReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(LatestReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	return srv.(MercuryServer).LatestReport(ctx, in)
}

func _Mercury_LatestReport_Grpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LatestReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MercuryServer).LatestReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mercury/LatestReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MercuryServer).LatestReport(ctx, req.(*LatestReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mercury_ServiceDesc is the wsrpc.ServiceDesc for Mercury service.
// It's only intended for direct use with wsrpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mercury_ServiceDesc = wsrpc.ServiceDesc{
	ServiceName: "pb.Mercury",
	HandlerType: (*MercuryServer)(nil),
	Methods: []wsrpc.MethodDesc{
		{
			MethodName: "Transmit",
			Handler:    _Mercury_Transmit_Handler,
		},
		{
			MethodName: "LatestReport",
			Handler:    _Mercury_LatestReport_Handler,
		},
	},
}

// Mercury_ServiceDesc_Grpc is the grpc.ServiceDesc for Mercury service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mercury_ServiceDesc_Grpc = grpc.ServiceDesc{
	ServiceName: "pb.Mercury",
	HandlerType: (*MercuryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transmit",
			Handler:    _Mercury_Transmit_Grpc_Handler,
		},
		{
			MethodName: "LatestReport",
			Handler:    _Mercury_LatestReport_Grpc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mercury.proto",
}
