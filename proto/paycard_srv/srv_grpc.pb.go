// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: paygw/proto/paycard_srv/srv.proto

package paycard_srv

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	paycard "paygw/proto/paycard"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PayCardServiceClient is the client API for PayCardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PayCardServiceClient interface {
	// POST /authorize
	Authorize(ctx context.Context, in *paycard.Authorization, opts ...grpc.CallOption) (*paycard.AuthorizationState, error)
	// POST /process/capture|void|refund|+state
	Process(ctx context.Context, in *paycard.Transaction, opts ...grpc.CallOption) (*paycard.TransactionState, error)
}

type payCardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPayCardServiceClient(cc grpc.ClientConnInterface) PayCardServiceClient {
	return &payCardServiceClient{cc}
}

func (c *payCardServiceClient) Authorize(ctx context.Context, in *paycard.Authorization, opts ...grpc.CallOption) (*paycard.AuthorizationState, error) {
	out := new(paycard.AuthorizationState)
	err := c.cc.Invoke(ctx, "/proto.paycard_srv.PayCardService/Authorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payCardServiceClient) Process(ctx context.Context, in *paycard.Transaction, opts ...grpc.CallOption) (*paycard.TransactionState, error) {
	out := new(paycard.TransactionState)
	err := c.cc.Invoke(ctx, "/proto.paycard_srv.PayCardService/Process", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayCardServiceServer is the server API for PayCardService service.
// All implementations must embed UnimplementedPayCardServiceServer
// for forward compatibility
type PayCardServiceServer interface {
	// POST /authorize
	Authorize(context.Context, *paycard.Authorization) (*paycard.AuthorizationState, error)
	// POST /process/capture|void|refund|+state
	Process(context.Context, *paycard.Transaction) (*paycard.TransactionState, error)
	mustEmbedUnimplementedPayCardServiceServer()
}

// UnimplementedPayCardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPayCardServiceServer struct {
}

func (UnimplementedPayCardServiceServer) Authorize(context.Context, *paycard.Authorization) (*paycard.AuthorizationState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}
func (UnimplementedPayCardServiceServer) Process(context.Context, *paycard.Transaction) (*paycard.TransactionState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Process not implemented")
}
func (UnimplementedPayCardServiceServer) mustEmbedUnimplementedPayCardServiceServer() {}

// UnsafePayCardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PayCardServiceServer will
// result in compilation errors.
type UnsafePayCardServiceServer interface {
	mustEmbedUnimplementedPayCardServiceServer()
}

func RegisterPayCardServiceServer(s grpc.ServiceRegistrar, srv PayCardServiceServer) {
	s.RegisterService(&PayCardService_ServiceDesc, srv)
}

func _PayCardService_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(paycard.Authorization)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayCardServiceServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.paycard_srv.PayCardService/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayCardServiceServer).Authorize(ctx, req.(*paycard.Authorization))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayCardService_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(paycard.Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayCardServiceServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.paycard_srv.PayCardService/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayCardServiceServer).Process(ctx, req.(*paycard.Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

// PayCardService_ServiceDesc is the grpc.ServiceDesc for PayCardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PayCardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.paycard_srv.PayCardService",
	HandlerType: (*PayCardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authorize",
			Handler:    _PayCardService_Authorize_Handler,
		},
		{
			MethodName: "Process",
			Handler:    _PayCardService_Process_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "paygw/proto/paycard_srv/srv.proto",
}