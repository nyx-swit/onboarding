// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: channel_message.proto

package v1

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

// ChannelMessageServiceClient is the client API for ChannelMessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChannelMessageServiceClient interface {
	GetListChannelMessage(ctx context.Context, in *GetListChannelMessageRequest, opts ...grpc.CallOption) (*GetListChannelMessageResponse, error)
	SetChannelMessage(ctx context.Context, in *SetChannelMessageRequest, opts ...grpc.CallOption) (*SetChannelMessageResponse, error)
}

type channelMessageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelMessageServiceClient(cc grpc.ClientConnInterface) ChannelMessageServiceClient {
	return &channelMessageServiceClient{cc}
}

func (c *channelMessageServiceClient) GetListChannelMessage(ctx context.Context, in *GetListChannelMessageRequest, opts ...grpc.CallOption) (*GetListChannelMessageResponse, error) {
	out := new(GetListChannelMessageResponse)
	err := c.cc.Invoke(ctx, "/v1.channel_message.ChannelMessageService/GetListChannelMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelMessageServiceClient) SetChannelMessage(ctx context.Context, in *SetChannelMessageRequest, opts ...grpc.CallOption) (*SetChannelMessageResponse, error) {
	out := new(SetChannelMessageResponse)
	err := c.cc.Invoke(ctx, "/v1.channel_message.ChannelMessageService/SetChannelMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChannelMessageServiceServer is the server API for ChannelMessageService service.
// All implementations must embed UnimplementedChannelMessageServiceServer
// for forward compatibility
type ChannelMessageServiceServer interface {
	GetListChannelMessage(context.Context, *GetListChannelMessageRequest) (*GetListChannelMessageResponse, error)
	SetChannelMessage(context.Context, *SetChannelMessageRequest) (*SetChannelMessageResponse, error)
	mustEmbedUnimplementedChannelMessageServiceServer()
}

// UnimplementedChannelMessageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChannelMessageServiceServer struct {
}

func (UnimplementedChannelMessageServiceServer) GetListChannelMessage(context.Context, *GetListChannelMessageRequest) (*GetListChannelMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListChannelMessage not implemented")
}
func (UnimplementedChannelMessageServiceServer) SetChannelMessage(context.Context, *SetChannelMessageRequest) (*SetChannelMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetChannelMessage not implemented")
}
func (UnimplementedChannelMessageServiceServer) mustEmbedUnimplementedChannelMessageServiceServer() {}

// UnsafeChannelMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChannelMessageServiceServer will
// result in compilation errors.
type UnsafeChannelMessageServiceServer interface {
	mustEmbedUnimplementedChannelMessageServiceServer()
}

func RegisterChannelMessageServiceServer(s grpc.ServiceRegistrar, srv ChannelMessageServiceServer) {
	s.RegisterService(&ChannelMessageService_ServiceDesc, srv)
}

func _ChannelMessageService_GetListChannelMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListChannelMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelMessageServiceServer).GetListChannelMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.channel_message.ChannelMessageService/GetListChannelMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelMessageServiceServer).GetListChannelMessage(ctx, req.(*GetListChannelMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelMessageService_SetChannelMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetChannelMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelMessageServiceServer).SetChannelMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.channel_message.ChannelMessageService/SetChannelMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelMessageServiceServer).SetChannelMessage(ctx, req.(*SetChannelMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChannelMessageService_ServiceDesc is the grpc.ServiceDesc for ChannelMessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChannelMessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.channel_message.ChannelMessageService",
	HandlerType: (*ChannelMessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetListChannelMessage",
			Handler:    _ChannelMessageService_GetListChannelMessage_Handler,
		},
		{
			MethodName: "SetChannelMessage",
			Handler:    _ChannelMessageService_SetChannelMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "channel_message.proto",
}
