// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// BotServiceClient is the client API for BotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BotServiceClient interface {
	CreateBot(ctx context.Context, in *CreateBotRequest, opts ...grpc.CallOption) (*CreateBotResponse, error)
	FindBots(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FindBotsResponse, error)
}

type botServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBotServiceClient(cc grpc.ClientConnInterface) BotServiceClient {
	return &botServiceClient{cc}
}

func (c *botServiceClient) CreateBot(ctx context.Context, in *CreateBotRequest, opts ...grpc.CallOption) (*CreateBotResponse, error) {
	out := new(CreateBotResponse)
	err := c.cc.Invoke(ctx, "/BotService/CreateBot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServiceClient) FindBots(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FindBotsResponse, error) {
	out := new(FindBotsResponse)
	err := c.cc.Invoke(ctx, "/BotService/FindBots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BotServiceServer is the server API for BotService service.
// All implementations must embed UnimplementedBotServiceServer
// for forward compatibility
type BotServiceServer interface {
	CreateBot(context.Context, *CreateBotRequest) (*CreateBotResponse, error)
	FindBots(context.Context, *Empty) (*FindBotsResponse, error)
	mustEmbedUnimplementedBotServiceServer()
}

// UnimplementedBotServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBotServiceServer struct {
}

func (UnimplementedBotServiceServer) CreateBot(context.Context, *CreateBotRequest) (*CreateBotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBot not implemented")
}
func (UnimplementedBotServiceServer) FindBots(context.Context, *Empty) (*FindBotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindBots not implemented")
}
func (UnimplementedBotServiceServer) mustEmbedUnimplementedBotServiceServer() {}

// UnsafeBotServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BotServiceServer will
// result in compilation errors.
type UnsafeBotServiceServer interface {
	mustEmbedUnimplementedBotServiceServer()
}

func RegisterBotServiceServer(s grpc.ServiceRegistrar, srv BotServiceServer) {
	s.RegisterService(&BotService_ServiceDesc, srv)
}

func _BotService_CreateBot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServiceServer).CreateBot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BotService/CreateBot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServiceServer).CreateBot(ctx, req.(*CreateBotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotService_FindBots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServiceServer).FindBots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BotService/FindBots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServiceServer).FindBots(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// BotService_ServiceDesc is the grpc.ServiceDesc for BotService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BotService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BotService",
	HandlerType: (*BotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBot",
			Handler:    _BotService_CreateBot_Handler,
		},
		{
			MethodName: "FindBots",
			Handler:    _BotService_FindBots_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}
