// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dfuse/funnel/v1/funnel.proto

package funnel

import (
	context "context"
	fmt "fmt"
	deos "github.com/eoscanada/bstream/pb/dfuse/codecs/deos"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type StreamBlockRequest struct {
	FromBlockNum         int64    `protobuf:"varint,2,opt,name=fromBlockNum,proto3" json:"fromBlockNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamBlockRequest) Reset()         { *m = StreamBlockRequest{} }
func (m *StreamBlockRequest) String() string { return proto.CompactTextString(m) }
func (*StreamBlockRequest) ProtoMessage()    {}
func (*StreamBlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3f872ed66cead6b, []int{0}
}

func (m *StreamBlockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamBlockRequest.Unmarshal(m, b)
}
func (m *StreamBlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamBlockRequest.Marshal(b, m, deterministic)
}
func (m *StreamBlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamBlockRequest.Merge(m, src)
}
func (m *StreamBlockRequest) XXX_Size() int {
	return xxx_messageInfo_StreamBlockRequest.Size(m)
}
func (m *StreamBlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamBlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamBlockRequest proto.InternalMessageInfo

func (m *StreamBlockRequest) GetFromBlockNum() int64 {
	if m != nil {
		return m.FromBlockNum
	}
	return 0
}

func init() {
	proto.RegisterType((*StreamBlockRequest)(nil), "dfuse.search.v1.StreamBlockRequest")
}

func init() { proto.RegisterFile("dfuse/funnel/v1/funnel.proto", fileDescriptor_d3f872ed66cead6b) }

var fileDescriptor_d3f872ed66cead6b = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0x49, 0x2b, 0x2d,
	0x4e, 0xd5, 0x4f, 0x2b, 0xcd, 0xcb, 0x4b, 0xcd, 0xd1, 0x2f, 0x33, 0x84, 0xb2, 0xf4, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0xf8, 0xc1, 0xb2, 0x7a, 0xc5, 0xa9, 0x89, 0x45, 0xc9, 0x19, 0x7a, 0x65,
	0x86, 0x52, 0x50, 0xe5, 0xc9, 0xf9, 0x29, 0xa9, 0xc9, 0xc5, 0xfa, 0x29, 0xa9, 0xf9, 0x10, 0x02,
	0xa2, 0x5c, 0xc9, 0x82, 0x4b, 0x28, 0xb8, 0xa4, 0x28, 0x35, 0x31, 0xd7, 0x29, 0x27, 0x3f, 0x39,
	0x3b, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x89, 0x8b, 0x27, 0xad, 0x28, 0x1f, 0x22,
	0xe6, 0x57, 0x9a, 0x2b, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x1c, 0x84, 0x22, 0x66, 0x14, 0xc9, 0xc5,
	0xe6, 0x06, 0xb6, 0x58, 0xc8, 0x9f, 0x8b, 0x07, 0xc9, 0x8c, 0x62, 0x21, 0x65, 0x3d, 0x34, 0x37,
	0xe8, 0x61, 0x5a, 0x21, 0x25, 0x01, 0x55, 0x04, 0x71, 0x97, 0x1e, 0xd8, 0x49, 0x60, 0x05, 0x06,
	0x8c, 0x4e, 0x1c, 0x51, 0x6c, 0x10, 0x3f, 0x25, 0xb1, 0x81, 0x5d, 0x69, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0xdf, 0x14, 0x20, 0x96, 0xf4, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FunnelClient is the client API for Funnel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FunnelClient interface {
	StreamBlocks(ctx context.Context, in *StreamBlockRequest, opts ...grpc.CallOption) (Funnel_StreamBlocksClient, error)
}

type funnelClient struct {
	cc *grpc.ClientConn
}

func NewFunnelClient(cc *grpc.ClientConn) FunnelClient {
	return &funnelClient{cc}
}

func (c *funnelClient) StreamBlocks(ctx context.Context, in *StreamBlockRequest, opts ...grpc.CallOption) (Funnel_StreamBlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Funnel_serviceDesc.Streams[0], "/dfuse.search.v1.Funnel/StreamBlocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &funnelStreamBlocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Funnel_StreamBlocksClient interface {
	Recv() (*deos.Block, error)
	grpc.ClientStream
}

type funnelStreamBlocksClient struct {
	grpc.ClientStream
}

func (x *funnelStreamBlocksClient) Recv() (*deos.Block, error) {
	m := new(deos.Block)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FunnelServer is the server API for Funnel service.
type FunnelServer interface {
	StreamBlocks(*StreamBlockRequest, Funnel_StreamBlocksServer) error
}

// UnimplementedFunnelServer can be embedded to have forward compatible implementations.
type UnimplementedFunnelServer struct {
}

func (*UnimplementedFunnelServer) StreamBlocks(req *StreamBlockRequest, srv Funnel_StreamBlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamBlocks not implemented")
}

func RegisterFunnelServer(s *grpc.Server, srv FunnelServer) {
	s.RegisterService(&_Funnel_serviceDesc, srv)
}

func _Funnel_StreamBlocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamBlockRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FunnelServer).StreamBlocks(m, &funnelStreamBlocksServer{stream})
}

type Funnel_StreamBlocksServer interface {
	Send(*deos.Block) error
	grpc.ServerStream
}

type funnelStreamBlocksServer struct {
	grpc.ServerStream
}

func (x *funnelStreamBlocksServer) Send(m *deos.Block) error {
	return x.ServerStream.SendMsg(m)
}

var _Funnel_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dfuse.search.v1.Funnel",
	HandlerType: (*FunnelServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamBlocks",
			Handler:       _Funnel_StreamBlocks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dfuse/funnel/v1/funnel.proto",
}
