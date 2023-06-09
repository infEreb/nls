// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: agent.proto

package gen

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

const (
	AgentService_GetAgent_FullMethodName    = "/AgentService/GetAgent"
	AgentService_GetAgents_FullMethodName   = "/AgentService/GetAgents"
	AgentService_PostAgent_FullMethodName   = "/AgentService/PostAgent"
	AgentService_PutAgent_FullMethodName    = "/AgentService/PutAgent"
	AgentService_DeleteAgent_FullMethodName = "/AgentService/DeleteAgent"
)

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentServiceClient interface {
	GetAgent(ctx context.Context, in *GetAgentRequest, opts ...grpc.CallOption) (*GetAgentResponse, error)
	GetAgents(ctx context.Context, in *GetAgentsRequest, opts ...grpc.CallOption) (*GetAgentsResponse, error)
	PostAgent(ctx context.Context, in *PostAgentRequest, opts ...grpc.CallOption) (*PostAgentResponse, error)
	PutAgent(ctx context.Context, in *PutAgentRequest, opts ...grpc.CallOption) (*PutAgentResponse, error)
	DeleteAgent(ctx context.Context, in *DeleteAgentRequest, opts ...grpc.CallOption) (*DeleteAgentResponse, error)
}

type agentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentServiceClient(cc grpc.ClientConnInterface) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) GetAgent(ctx context.Context, in *GetAgentRequest, opts ...grpc.CallOption) (*GetAgentResponse, error) {
	out := new(GetAgentResponse)
	err := c.cc.Invoke(ctx, AgentService_GetAgent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) GetAgents(ctx context.Context, in *GetAgentsRequest, opts ...grpc.CallOption) (*GetAgentsResponse, error) {
	out := new(GetAgentsResponse)
	err := c.cc.Invoke(ctx, AgentService_GetAgents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) PostAgent(ctx context.Context, in *PostAgentRequest, opts ...grpc.CallOption) (*PostAgentResponse, error) {
	out := new(PostAgentResponse)
	err := c.cc.Invoke(ctx, AgentService_PostAgent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) PutAgent(ctx context.Context, in *PutAgentRequest, opts ...grpc.CallOption) (*PutAgentResponse, error) {
	out := new(PutAgentResponse)
	err := c.cc.Invoke(ctx, AgentService_PutAgent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) DeleteAgent(ctx context.Context, in *DeleteAgentRequest, opts ...grpc.CallOption) (*DeleteAgentResponse, error) {
	out := new(DeleteAgentResponse)
	err := c.cc.Invoke(ctx, AgentService_DeleteAgent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServiceServer is the server API for AgentService service.
// All implementations must embed UnimplementedAgentServiceServer
// for forward compatibility
type AgentServiceServer interface {
	GetAgent(context.Context, *GetAgentRequest) (*GetAgentResponse, error)
	GetAgents(context.Context, *GetAgentsRequest) (*GetAgentsResponse, error)
	PostAgent(context.Context, *PostAgentRequest) (*PostAgentResponse, error)
	PutAgent(context.Context, *PutAgentRequest) (*PutAgentResponse, error)
	DeleteAgent(context.Context, *DeleteAgentRequest) (*DeleteAgentResponse, error)
	mustEmbedUnimplementedAgentServiceServer()
}

// UnimplementedAgentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServiceServer struct {
}

func (UnimplementedAgentServiceServer) GetAgent(context.Context, *GetAgentRequest) (*GetAgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgent not implemented")
}
func (UnimplementedAgentServiceServer) GetAgents(context.Context, *GetAgentsRequest) (*GetAgentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgents not implemented")
}
func (UnimplementedAgentServiceServer) PostAgent(context.Context, *PostAgentRequest) (*PostAgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostAgent not implemented")
}
func (UnimplementedAgentServiceServer) PutAgent(context.Context, *PutAgentRequest) (*PutAgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutAgent not implemented")
}
func (UnimplementedAgentServiceServer) DeleteAgent(context.Context, *DeleteAgentRequest) (*DeleteAgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAgent not implemented")
}
func (UnimplementedAgentServiceServer) mustEmbedUnimplementedAgentServiceServer() {}

// UnsafeAgentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServiceServer will
// result in compilation errors.
type UnsafeAgentServiceServer interface {
	mustEmbedUnimplementedAgentServiceServer()
}

func RegisterAgentServiceServer(s grpc.ServiceRegistrar, srv AgentServiceServer) {
	s.RegisterService(&AgentService_ServiceDesc, srv)
}

func _AgentService_GetAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).GetAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_GetAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).GetAgent(ctx, req.(*GetAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_GetAgents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAgentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).GetAgents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_GetAgents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).GetAgents(ctx, req.(*GetAgentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_PostAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).PostAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_PostAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).PostAgent(ctx, req.(*PostAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_PutAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).PutAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_PutAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).PutAgent(ctx, req.(*PutAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_DeleteAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).DeleteAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_DeleteAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).DeleteAgent(ctx, req.(*DeleteAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentService_ServiceDesc is the grpc.ServiceDesc for AgentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAgent",
			Handler:    _AgentService_GetAgent_Handler,
		},
		{
			MethodName: "GetAgents",
			Handler:    _AgentService_GetAgents_Handler,
		},
		{
			MethodName: "PostAgent",
			Handler:    _AgentService_PostAgent_Handler,
		},
		{
			MethodName: "PutAgent",
			Handler:    _AgentService_PutAgent_Handler,
		},
		{
			MethodName: "DeleteAgent",
			Handler:    _AgentService_DeleteAgent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}
