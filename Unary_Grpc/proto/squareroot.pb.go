// Code generated by protoc-gen-go. DO NOT EDIT.
// source: squareroot.proto

package proto

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
	Input                int32    `protobuf:"varint,1,opt,name=Input,proto3" json:"Input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_c51098fddf27cb29, []int{0}
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

func (m *Request) GetInput() int32 {
	if m != nil {
		return m.Input
	}
	return 0
}

type Response struct {
	Output               int32    `protobuf:"varint,1,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c51098fddf27cb29, []int{1}
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

func (m *Response) GetOutput() int32 {
	if m != nil {
		return m.Output
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "GRPC_Practice.Request")
	proto.RegisterType((*Response)(nil), "GRPC_Practice.Response")
}

func init() { proto.RegisterFile("squareroot.proto", fileDescriptor_c51098fddf27cb29) }

var fileDescriptor_c51098fddf27cb29 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x2e, 0x2c, 0x4d,
	0x2c, 0x4a, 0x2d, 0xca, 0xcf, 0x2f, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x75, 0x0f,
	0x0a, 0x70, 0x8e, 0x0f, 0x28, 0x4a, 0x4c, 0x2e, 0xc9, 0x4c, 0x4e, 0x55, 0x92, 0xe7, 0x62, 0x0f,
	0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe1, 0x62, 0xf5, 0xcc, 0x2b, 0x28, 0x2d, 0x91,
	0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x82, 0x70, 0x94, 0x94, 0xb8, 0x38, 0x82, 0x52, 0x8b, 0x0b,
	0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xc4, 0xb8, 0xd8, 0xf2, 0x4b, 0x4b, 0x10, 0x4a, 0xa0, 0x3c, 0x23,
	0x5f, 0x2e, 0xae, 0x60, 0xb8, 0x3d, 0x42, 0xf6, 0x28, 0x3c, 0x31, 0x3d, 0x14, 0x0b, 0xf5, 0xa0,
	0xb6, 0x49, 0x89, 0x63, 0x88, 0x43, 0x2c, 0x51, 0x62, 0x70, 0x62, 0x8f, 0x62, 0x05, 0xbb, 0x35,
	0x89, 0x0d, 0x4c, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x73, 0x7e, 0x4b, 0xa9, 0xc6, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SquarerootClient is the client API for Squareroot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SquarerootClient interface {
	//unary streaming
	//The error being sent is of type INVALID_ARGUMENT
	Squareroot(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type squarerootClient struct {
	cc *grpc.ClientConn
}

func NewSquarerootClient(cc *grpc.ClientConn) SquarerootClient {
	return &squarerootClient{cc}
}

func (c *squarerootClient) Squareroot(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/GRPC_Practice.Squareroot/Squareroot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SquarerootServer is the server API for Squareroot service.
type SquarerootServer interface {
	//unary streaming
	//The error being sent is of type INVALID_ARGUMENT
	Squareroot(context.Context, *Request) (*Response, error)
}

// UnimplementedSquarerootServer can be embedded to have forward compatible implementations.
type UnimplementedSquarerootServer struct {
}

func (*UnimplementedSquarerootServer) Squareroot(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Squareroot not implemented")
}

func RegisterSquarerootServer(s *grpc.Server, srv SquarerootServer) {
	s.RegisterService(&_Squareroot_serviceDesc, srv)
}

func _Squareroot_Squareroot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SquarerootServer).Squareroot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GRPC_Practice.Squareroot/Squareroot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SquarerootServer).Squareroot(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Squareroot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "GRPC_Practice.Squareroot",
	HandlerType: (*SquarerootServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Squareroot",
			Handler:    _Squareroot_Squareroot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "squareroot.proto",
}
