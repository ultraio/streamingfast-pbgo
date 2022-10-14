// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: sf/firehose/v2/firehose.proto

package pbfirehose

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

// StreamClient is the client API for Stream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamClient interface {
	Blocks(ctx context.Context, in *Request, opts ...grpc.CallOption) (Stream_BlocksClient, error)
}

type streamClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamClient(cc grpc.ClientConnInterface) StreamClient {
	return &streamClient{cc}
}

func (c *streamClient) Blocks(ctx context.Context, in *Request, opts ...grpc.CallOption) (Stream_BlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &Stream_ServiceDesc.Streams[0], "/sf.firehose.v2.Stream/Blocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamBlocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Stream_BlocksClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type streamBlocksClient struct {
	grpc.ClientStream
}

func (x *streamBlocksClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServer is the server API for Stream service.
// All implementations should embed UnimplementedStreamServer
// for forward compatibility
type StreamServer interface {
	Blocks(*Request, Stream_BlocksServer) error
}

// UnimplementedStreamServer should be embedded to have forward compatible implementations.
type UnimplementedStreamServer struct {
}

func (UnimplementedStreamServer) Blocks(*Request, Stream_BlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method Blocks not implemented")
}

// UnsafeStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamServer will
// result in compilation errors.
type UnsafeStreamServer interface {
	mustEmbedUnimplementedStreamServer()
}

func RegisterStreamServer(s grpc.ServiceRegistrar, srv StreamServer) {
	s.RegisterService(&Stream_ServiceDesc, srv)
}

func _Stream_Blocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServer).Blocks(m, &streamBlocksServer{stream})
}

type Stream_BlocksServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type streamBlocksServer struct {
	grpc.ServerStream
}

func (x *streamBlocksServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

// Stream_ServiceDesc is the grpc.ServiceDesc for Stream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Stream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sf.firehose.v2.Stream",
	HandlerType: (*StreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Blocks",
			Handler:       _Stream_Blocks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sf/firehose/v2/firehose.proto",
}

// FetchClient is the client API for Fetch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FetchClient interface {
	Block(ctx context.Context, in *SingleBlockRequest, opts ...grpc.CallOption) (*SingleBlockResponse, error)
}

type fetchClient struct {
	cc grpc.ClientConnInterface
}

func NewFetchClient(cc grpc.ClientConnInterface) FetchClient {
	return &fetchClient{cc}
}

func (c *fetchClient) Block(ctx context.Context, in *SingleBlockRequest, opts ...grpc.CallOption) (*SingleBlockResponse, error) {
	out := new(SingleBlockResponse)
	err := c.cc.Invoke(ctx, "/sf.firehose.v2.Fetch/Block", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FetchServer is the server API for Fetch service.
// All implementations should embed UnimplementedFetchServer
// for forward compatibility
type FetchServer interface {
	Block(context.Context, *SingleBlockRequest) (*SingleBlockResponse, error)
}

// UnimplementedFetchServer should be embedded to have forward compatible implementations.
type UnimplementedFetchServer struct {
}

func (UnimplementedFetchServer) Block(context.Context, *SingleBlockRequest) (*SingleBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Block not implemented")
}

// UnsafeFetchServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FetchServer will
// result in compilation errors.
type UnsafeFetchServer interface {
	mustEmbedUnimplementedFetchServer()
}

func RegisterFetchServer(s grpc.ServiceRegistrar, srv FetchServer) {
	s.RegisterService(&Fetch_ServiceDesc, srv)
}

func _Fetch_Block_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetchServer).Block(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sf.firehose.v2.Fetch/Block",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetchServer).Block(ctx, req.(*SingleBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Fetch_ServiceDesc is the grpc.ServiceDesc for Fetch service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fetch_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sf.firehose.v2.Fetch",
	HandlerType: (*FetchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Block",
			Handler:    _Fetch_Block_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sf/firehose/v2/firehose.proto",
}
