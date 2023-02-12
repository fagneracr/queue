// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: messages/messages.proto

package grpcserve

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

// QServiceClient is the client API for QService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QServiceClient interface {
	NewQ(ctx context.Context, in *QueueMSG, opts ...grpc.CallOption) (*NewQResponse, error)
}

type qServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQServiceClient(cc grpc.ClientConnInterface) QServiceClient {
	return &qServiceClient{cc}
}

func (c *qServiceClient) NewQ(ctx context.Context, in *QueueMSG, opts ...grpc.CallOption) (*NewQResponse, error) {
	out := new(NewQResponse)
	err := c.cc.Invoke(ctx, "/QService/NewQ", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


// QServiceServer is the server API for QService service.
// All implementations must embed UnimplementedQServiceServer
// for forward compatibility
type QServiceServer interface {
	NewQ(context.Context, *QueueMSG) (*NewQResponse, error)
	mustEmbedUnimplementedQServiceServer()
}

// UnimplementedQServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQServiceServer struct {
}

func (UnimplementedQServiceServer) NewQ(context.Context, *QueueMSG) (*NewQResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewQ not implemented")
}
func (UnimplementedQServiceServer) mustEmbedUnimplementedQServiceServer() {}

// UnsafeQServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QServiceServer will
// result in compilation errors.
type UnsafeQServiceServer interface {
	mustEmbedUnimplementedQServiceServer()
}

func RegisterQServiceServer(s grpc.ServiceRegistrar, srv QServiceServer) {
	s.RegisterService(&QService_ServiceDesc, srv)
}

func _QService_NewQ_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueueMSG)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QServiceServer).NewQ(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QService/NewQ",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QServiceServer).NewQ(ctx, req.(*QueueMSG))
	}
	return interceptor(ctx, in, info, handler)
}

// QService_ServiceDesc is the grpc.ServiceDesc for QService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "QService",
	HandlerType: (*QServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewQ",
			Handler:    _QService_NewQ_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messages/messages.proto",
}