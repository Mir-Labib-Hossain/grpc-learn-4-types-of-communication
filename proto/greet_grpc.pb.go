// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.26.1
// source: proto/greet.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	GreetService_SayHello_FullMethodName                       = "/greet_service.GreetService/SayHello"
	GreetService_SayHelloServerStreaming_FullMethodName        = "/greet_service.GreetService/SayHelloServerStreaming"
	GreetService_SayHelloClientStreaming_FullMethodName        = "/greet_service.GreetService/SayHelloClientStreaming"
	GreetService_SayHelloBiDirectionalStreaming_FullMethodName = "/greet_service.GreetService/SayHelloBiDirectionalStreaming"
)

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetServiceClient interface {
	SayHello(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*HelloRes, error)
	SayHelloServerStreaming(ctx context.Context, in *NameList, opts ...grpc.CallOption) (GreetService_SayHelloServerStreamingClient, error)
	SayHelloClientStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_SayHelloClientStreamingClient, error)
	SayHelloBiDirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_SayHelloBiDirectionalStreamingClient, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) SayHello(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*HelloRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HelloRes)
	err := c.cc.Invoke(ctx, GreetService_SayHello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) SayHelloServerStreaming(ctx context.Context, in *NameList, opts ...grpc.CallOption) (GreetService_SayHelloServerStreamingClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[0], GreetService_SayHelloServerStreaming_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceSayHelloServerStreamingClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_SayHelloServerStreamingClient interface {
	Recv() (*HelloRes, error)
	grpc.ClientStream
}

type greetServiceSayHelloServerStreamingClient struct {
	grpc.ClientStream
}

func (x *greetServiceSayHelloServerStreamingClient) Recv() (*HelloRes, error) {
	m := new(HelloRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) SayHelloClientStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_SayHelloClientStreamingClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[1], GreetService_SayHelloClientStreaming_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceSayHelloClientStreamingClient{ClientStream: stream}
	return x, nil
}

type GreetService_SayHelloClientStreamingClient interface {
	Send(*HelloReq) error
	CloseAndRecv() (*MessageList, error)
	grpc.ClientStream
}

type greetServiceSayHelloClientStreamingClient struct {
	grpc.ClientStream
}

func (x *greetServiceSayHelloClientStreamingClient) Send(m *HelloReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceSayHelloClientStreamingClient) CloseAndRecv() (*MessageList, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MessageList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) SayHelloBiDirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_SayHelloBiDirectionalStreamingClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[2], GreetService_SayHelloBiDirectionalStreaming_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceSayHelloBiDirectionalStreamingClient{ClientStream: stream}
	return x, nil
}

type GreetService_SayHelloBiDirectionalStreamingClient interface {
	Send(*HelloReq) error
	Recv() (*HelloRes, error)
	grpc.ClientStream
}

type greetServiceSayHelloBiDirectionalStreamingClient struct {
	grpc.ClientStream
}

func (x *greetServiceSayHelloBiDirectionalStreamingClient) Send(m *HelloReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceSayHelloBiDirectionalStreamingClient) Recv() (*HelloRes, error) {
	m := new(HelloRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
// All implementations must embed UnimplementedGreetServiceServer
// for forward compatibility
type GreetServiceServer interface {
	SayHello(context.Context, *NoParam) (*HelloRes, error)
	SayHelloServerStreaming(*NameList, GreetService_SayHelloServerStreamingServer) error
	SayHelloClientStreaming(GreetService_SayHelloClientStreamingServer) error
	SayHelloBiDirectionalStreaming(GreetService_SayHelloBiDirectionalStreamingServer) error
	mustEmbedUnimplementedGreetServiceServer()
}

// UnimplementedGreetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (UnimplementedGreetServiceServer) SayHello(context.Context, *NoParam) (*HelloRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreetServiceServer) SayHelloServerStreaming(*NameList, GreetService_SayHelloServerStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloServerStreaming not implemented")
}
func (UnimplementedGreetServiceServer) SayHelloClientStreaming(GreetService_SayHelloClientStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloClientStreaming not implemented")
}
func (UnimplementedGreetServiceServer) SayHelloBiDirectionalStreaming(GreetService_SayHelloBiDirectionalStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloBiDirectionalStreaming not implemented")
}
func (UnimplementedGreetServiceServer) mustEmbedUnimplementedGreetServiceServer() {}

// UnsafeGreetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServiceServer will
// result in compilation errors.
type UnsafeGreetServiceServer interface {
	mustEmbedUnimplementedGreetServiceServer()
}

func RegisterGreetServiceServer(s grpc.ServiceRegistrar, srv GreetServiceServer) {
	s.RegisterService(&GreetService_ServiceDesc, srv)
}

func _GreetService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GreetService_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).SayHello(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_SayHelloServerStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NameList)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).SayHelloServerStreaming(m, &greetServiceSayHelloServerStreamingServer{ServerStream: stream})
}

type GreetService_SayHelloServerStreamingServer interface {
	Send(*HelloRes) error
	grpc.ServerStream
}

type greetServiceSayHelloServerStreamingServer struct {
	grpc.ServerStream
}

func (x *greetServiceSayHelloServerStreamingServer) Send(m *HelloRes) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetService_SayHelloClientStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).SayHelloClientStreaming(&greetServiceSayHelloClientStreamingServer{ServerStream: stream})
}

type GreetService_SayHelloClientStreamingServer interface {
	SendAndClose(*MessageList) error
	Recv() (*HelloReq, error)
	grpc.ServerStream
}

type greetServiceSayHelloClientStreamingServer struct {
	grpc.ServerStream
}

func (x *greetServiceSayHelloClientStreamingServer) SendAndClose(m *MessageList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceSayHelloClientStreamingServer) Recv() (*HelloReq, error) {
	m := new(HelloReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetService_SayHelloBiDirectionalStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).SayHelloBiDirectionalStreaming(&greetServiceSayHelloBiDirectionalStreamingServer{ServerStream: stream})
}

type GreetService_SayHelloBiDirectionalStreamingServer interface {
	Send(*HelloRes) error
	Recv() (*HelloReq, error)
	grpc.ServerStream
}

type greetServiceSayHelloBiDirectionalStreamingServer struct {
	grpc.ServerStream
}

func (x *greetServiceSayHelloBiDirectionalStreamingServer) Send(m *HelloRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceSayHelloBiDirectionalStreamingServer) Recv() (*HelloReq, error) {
	m := new(HelloReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetService_ServiceDesc is the grpc.ServiceDesc for GreetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet_service.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _GreetService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloServerStreaming",
			Handler:       _GreetService_SayHelloServerStreaming_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayHelloClientStreaming",
			Handler:       _GreetService_SayHelloClientStreaming_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SayHelloBiDirectionalStreaming",
			Handler:       _GreetService_SayHelloBiDirectionalStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/greet.proto",
}