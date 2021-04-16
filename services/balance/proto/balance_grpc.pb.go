// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// BalanceServiceClient is the client API for BalanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BalanceServiceClient interface {
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
}

type balanceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBalanceServiceClient(cc grpc.ClientConnInterface) BalanceServiceClient {
	return &balanceServiceClient{cc}
}

func (c *balanceServiceClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, "/mercari.go_conference_2021_spring_office_hour.balance.BalanceService/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BalanceServiceServer is the server API for BalanceService service.
// All implementations must embed UnimplementedBalanceServiceServer
// for forward compatibility
type BalanceServiceServer interface {
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	mustEmbedUnimplementedBalanceServiceServer()
}

// UnimplementedBalanceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBalanceServiceServer struct {
}

func (UnimplementedBalanceServiceServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedBalanceServiceServer) mustEmbedUnimplementedBalanceServiceServer() {}

// UnsafeBalanceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BalanceServiceServer will
// result in compilation errors.
type UnsafeBalanceServiceServer interface {
	mustEmbedUnimplementedBalanceServiceServer()
}

func RegisterBalanceServiceServer(s grpc.ServiceRegistrar, srv BalanceServiceServer) {
	s.RegisterService(&BalanceService_ServiceDesc, srv)
}

func _BalanceService_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceServiceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mercari.go_conference_2021_spring_office_hour.balance.BalanceService/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceServiceServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BalanceService_ServiceDesc is the grpc.ServiceDesc for BalanceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BalanceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mercari.go_conference_2021_spring_office_hour.balance.BalanceService",
	HandlerType: (*BalanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBalance",
			Handler:    _BalanceService_GetBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/balance/proto/balance.proto",
}