// Code generated by protoc-gen-go.
// source: protobuf.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	protobuf.proto

It has these top-level messages:
	PingRequest
	PingResponse
*/
package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PingRequest struct {
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PingResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PingResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "protobuf.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "protobuf.PingResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Info service

type InfoClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type infoClient struct {
	cc *grpc.ClientConn
}

func NewInfoClient(cc *grpc.ClientConn) InfoClient {
	return &infoClient{cc}
}

func (c *infoClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/protobuf.Info/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Info service

type InfoServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

func RegisterInfoServer(s *grpc.Server, srv InfoServer) {
	s.RegisterService(&_Info_serviceDesc, srv)
}

func _Info_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Info/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Info_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Info",
	HandlerType: (*InfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Info_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf.proto",
}

func init() { proto.RegisterFile("protobuf.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x03, 0x33, 0x84, 0x38, 0x60, 0x7c, 0x25, 0x5e, 0x2e, 0xee, 0x80,
	0xcc, 0xbc, 0xf4, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x25, 0x0d, 0x2e, 0x1e, 0x08, 0xb7,
	0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x3d,
	0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc6, 0x35, 0xb2, 0xe7, 0x62, 0xf1, 0xcc, 0x4b,
	0xcb, 0x17, 0x32, 0xe7, 0x62, 0x01, 0xe9, 0x10, 0x12, 0xd5, 0x83, 0xdb, 0x81, 0x64, 0xa0, 0x94,
	0x18, 0xba, 0x30, 0xc4, 0x60, 0x25, 0x86, 0x24, 0x36, 0xb0, 0x84, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0xa2, 0xae, 0x58, 0x31, 0x9c, 0x00, 0x00, 0x00,
}