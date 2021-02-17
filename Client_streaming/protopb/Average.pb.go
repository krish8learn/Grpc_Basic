// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Average.proto

package protopb

import (
	context "context"
	fmt "fmt"
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

type Request struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bd4c87b2d4a7036, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type Response struct {
	Avg                  float64  `protobuf:"fixed64,1,opt,name=Avg,proto3" json:"Avg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bd4c87b2d4a7036, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetAvg() float64 {
	if m != nil {
		return m.Avg
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "Client_streaming.Request")
	proto.RegisterType((*Response)(nil), "Client_streaming.Response")
}

func init() { proto.RegisterFile("Average.proto", fileDescriptor_0bd4c87b2d4a7036) }

var fileDescriptor_0bd4c87b2d4a7036 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x75, 0x2c, 0x4b, 0x2d,
	0x4a, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x70, 0xce, 0xc9, 0x4c, 0xcd,
	0x2b, 0x89, 0x2f, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xcd, 0xcc, 0x4b, 0x57, 0x52, 0xe4, 0x62, 0x0f,
	0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe3, 0x62, 0xcb, 0x2b, 0xcd, 0x4d, 0x4a, 0x2d,
	0x92, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x82, 0xf2, 0x94, 0x64, 0xb8, 0x38, 0x82, 0x52, 0x8b,
	0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x04, 0xb8, 0x98, 0x1d, 0xcb, 0xd2, 0xc1, 0x0a, 0x18, 0x83,
	0x40, 0x4c, 0x23, 0x7f, 0x2e, 0x76, 0xa8, 0x1d, 0x42, 0x2e, 0x08, 0xa6, 0xa4, 0x1e, 0xba, 0x4d,
	0x7a, 0x50, 0x6b, 0xa4, 0xa4, 0xb0, 0x49, 0x41, 0x8c, 0x57, 0x62, 0xd0, 0x60, 0x74, 0xe2, 0x8c,
	0x62, 0x07, 0x3b, 0xb6, 0x20, 0x29, 0x89, 0x0d, 0xcc, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x1a, 0xe4, 0x4d, 0x87, 0xc6, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AverageClient is the client API for Average service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AverageClient interface {
	Average(ctx context.Context, opts ...grpc.CallOption) (Average_AverageClient, error)
}

type averageClient struct {
	cc *grpc.ClientConn
}

func NewAverageClient(cc *grpc.ClientConn) AverageClient {
	return &averageClient{cc}
}

func (c *averageClient) Average(ctx context.Context, opts ...grpc.CallOption) (Average_AverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Average_serviceDesc.Streams[0], "/Client_streaming.Average/Average", opts...)
	if err != nil {
		return nil, err
	}
	x := &averageAverageClient{stream}
	return x, nil
}

type Average_AverageClient interface {
	Send(*Request) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type averageAverageClient struct {
	grpc.ClientStream
}

func (x *averageAverageClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *averageAverageClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AverageServer is the server API for Average service.
type AverageServer interface {
	Average(Average_AverageServer) error
}

// UnimplementedAverageServer can be embedded to have forward compatible implementations.
type UnimplementedAverageServer struct {
}

func (*UnimplementedAverageServer) Average(srv Average_AverageServer) error {
	return status.Errorf(codes.Unimplemented, "method Average not implemented")
}

func RegisterAverageServer(s *grpc.Server, srv AverageServer) {
	s.RegisterService(&_Average_serviceDesc, srv)
}

func _Average_Average_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AverageServer).Average(&averageAverageServer{stream})
}

type Average_AverageServer interface {
	SendAndClose(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type averageAverageServer struct {
	grpc.ServerStream
}

func (x *averageAverageServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *averageAverageServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Average_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Client_streaming.Average",
	HandlerType: (*AverageServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Average",
			Handler:       _Average_Average_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "Average.proto",
}