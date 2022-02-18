// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: channel_member.proto

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

// ChannelMemberServiceClient is the client API for ChannelMemberService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChannelMemberServiceClient interface {
	GetListChannelMember(ctx context.Context, in *GetListChannelMemberRequest, opts ...grpc.CallOption) (*GetListChannelMemberResponse, error)
}

type channelMemberServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelMemberServiceClient(cc grpc.ClientConnInterface) ChannelMemberServiceClient {
	return &channelMemberServiceClient{cc}
}

func (c *channelMemberServiceClient) GetListChannelMember(ctx context.Context, in *GetListChannelMemberRequest, opts ...grpc.CallOption) (*GetListChannelMemberResponse, error) {
	out := new(GetListChannelMemberResponse)
	err := c.cc.Invoke(ctx, "/v1.channel.ChannelMemberService/getListChannelMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChannelMemberServiceServer is the server API for ChannelMemberService service.
// All implementations must embed UnimplementedChannelMemberServiceServer
// for forward compatibility
type ChannelMemberServiceServer interface {
	GetListChannelMember(context.Context, *GetListChannelMemberRequest) (*GetListChannelMemberResponse, error)
	mustEmbedUnimplementedChannelMemberServiceServer()
}

// UnimplementedChannelMemberServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChannelMemberServiceServer struct {
}

func (UnimplementedChannelMemberServiceServer) GetListChannelMember(context.Context, *GetListChannelMemberRequest) (*GetListChannelMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListChannelMember not implemented")
}
func (UnimplementedChannelMemberServiceServer) mustEmbedUnimplementedChannelMemberServiceServer() {}

// UnsafeChannelMemberServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChannelMemberServiceServer will
// result in compilation errors.
type UnsafeChannelMemberServiceServer interface {
	mustEmbedUnimplementedChannelMemberServiceServer()
}

func RegisterChannelMemberServiceServer(s grpc.ServiceRegistrar, srv ChannelMemberServiceServer) {
	s.RegisterService(&ChannelMemberService_ServiceDesc, srv)
}

func _ChannelMemberService_GetListChannelMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListChannelMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelMemberServiceServer).GetListChannelMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.channel.ChannelMemberService/getListChannelMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelMemberServiceServer).GetListChannelMember(ctx, req.(*GetListChannelMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChannelMemberService_ServiceDesc is the grpc.ServiceDesc for ChannelMemberService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChannelMemberService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.channel.ChannelMemberService",
	HandlerType: (*ChannelMemberServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getListChannelMember",
			Handler:    _ChannelMemberService_GetListChannelMember_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "channel_member.proto",
}
