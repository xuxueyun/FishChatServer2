// Code generated by protoc-gen-go.
// source: idgen.proto
// DO NOT EDIT!

package rpc

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

type Snowflake struct {
}

func (m *Snowflake) Reset()                    { *m = Snowflake{} }
func (m *Snowflake) String() string            { return proto.CompactTextString(m) }
func (*Snowflake) ProtoMessage()               {}
func (*Snowflake) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type Snowflake_Key struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Snowflake_Key) Reset()                    { *m = Snowflake_Key{} }
func (m *Snowflake_Key) String() string            { return proto.CompactTextString(m) }
func (*Snowflake_Key) ProtoMessage()               {}
func (*Snowflake_Key) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

func (m *Snowflake_Key) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Snowflake_Value struct {
	Value int64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Snowflake_Value) Reset()                    { *m = Snowflake_Value{} }
func (m *Snowflake_Value) String() string            { return proto.CompactTextString(m) }
func (*Snowflake_Value) ProtoMessage()               {}
func (*Snowflake_Value) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 1} }

func (m *Snowflake_Value) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Snowflake_NullRequest struct {
}

func (m *Snowflake_NullRequest) Reset()                    { *m = Snowflake_NullRequest{} }
func (m *Snowflake_NullRequest) String() string            { return proto.CompactTextString(m) }
func (*Snowflake_NullRequest) ProtoMessage()               {}
func (*Snowflake_NullRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 2} }

type Snowflake_UUID struct {
	Uuid uint64 `protobuf:"varint,1,opt,name=uuid" json:"uuid,omitempty"`
}

func (m *Snowflake_UUID) Reset()                    { *m = Snowflake_UUID{} }
func (m *Snowflake_UUID) String() string            { return proto.CompactTextString(m) }
func (*Snowflake_UUID) ProtoMessage()               {}
func (*Snowflake_UUID) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 3} }

func (m *Snowflake_UUID) GetUuid() uint64 {
	if m != nil {
		return m.Uuid
	}
	return 0
}

func init() {
	proto.RegisterType((*Snowflake)(nil), "rpc.Snowflake")
	proto.RegisterType((*Snowflake_Key)(nil), "rpc.Snowflake.Key")
	proto.RegisterType((*Snowflake_Value)(nil), "rpc.Snowflake.Value")
	proto.RegisterType((*Snowflake_NullRequest)(nil), "rpc.Snowflake.NullRequest")
	proto.RegisterType((*Snowflake_UUID)(nil), "rpc.Snowflake.UUID")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for IDGenServerRPC service

type IDGenServerRPCClient interface {
	Next(ctx context.Context, in *Snowflake_Key, opts ...grpc.CallOption) (*Snowflake_Value, error)
	GetUUID(ctx context.Context, in *Snowflake_NullRequest, opts ...grpc.CallOption) (*Snowflake_UUID, error)
}

type iDGenServerRPCClient struct {
	cc *grpc.ClientConn
}

func NewIDGenServerRPCClient(cc *grpc.ClientConn) IDGenServerRPCClient {
	return &iDGenServerRPCClient{cc}
}

func (c *iDGenServerRPCClient) Next(ctx context.Context, in *Snowflake_Key, opts ...grpc.CallOption) (*Snowflake_Value, error) {
	out := new(Snowflake_Value)
	err := grpc.Invoke(ctx, "/rpc.IDGenServerRPC/Next", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iDGenServerRPCClient) GetUUID(ctx context.Context, in *Snowflake_NullRequest, opts ...grpc.CallOption) (*Snowflake_UUID, error) {
	out := new(Snowflake_UUID)
	err := grpc.Invoke(ctx, "/rpc.IDGenServerRPC/GetUUID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IDGenServerRPC service

type IDGenServerRPCServer interface {
	Next(context.Context, *Snowflake_Key) (*Snowflake_Value, error)
	GetUUID(context.Context, *Snowflake_NullRequest) (*Snowflake_UUID, error)
}

func RegisterIDGenServerRPCServer(s *grpc.Server, srv IDGenServerRPCServer) {
	s.RegisterService(&_IDGenServerRPC_serviceDesc, srv)
}

func _IDGenServerRPC_Next_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Snowflake_Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IDGenServerRPCServer).Next(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.IDGenServerRPC/Next",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IDGenServerRPCServer).Next(ctx, req.(*Snowflake_Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _IDGenServerRPC_GetUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Snowflake_NullRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IDGenServerRPCServer).GetUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.IDGenServerRPC/GetUUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IDGenServerRPCServer).GetUUID(ctx, req.(*Snowflake_NullRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IDGenServerRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.IDGenServerRPC",
	HandlerType: (*IDGenServerRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Next",
			Handler:    _IDGenServerRPC_Next_Handler,
		},
		{
			MethodName: "GetUUID",
			Handler:    _IDGenServerRPC_GetUUID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idgen.proto",
}

func init() { proto.RegisterFile("idgen.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x4c, 0x49, 0x4f,
	0xcd, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x56, 0x2a, 0xe0, 0xe2,
	0x0c, 0xce, 0xcb, 0x2f, 0x4f, 0xcb, 0x49, 0xcc, 0x4e, 0x95, 0x92, 0xe4, 0x62, 0xf6, 0x4e, 0xad,
	0x14, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02,
	0xb3, 0xa5, 0x64, 0xb9, 0x58, 0xc3, 0x12, 0x73, 0x4a, 0x53, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x40,
	0x0c, 0xb0, 0x2c, 0x73, 0x10, 0x84, 0x23, 0xc5, 0xcb, 0xc5, 0xed, 0x57, 0x9a, 0x93, 0x13, 0x94,
	0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x25, 0xc5, 0xc5, 0x12, 0x1a, 0xea, 0xe9, 0x02, 0x32, 0xa9,
	0xb4, 0x34, 0x33, 0x05, 0xac, 0x96, 0x25, 0x08, 0xcc, 0x36, 0xaa, 0xe3, 0xe2, 0xf3, 0x74, 0x71,
	0x4f, 0xcd, 0x0b, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x0a, 0x0a, 0x70, 0x16, 0x32, 0xe0, 0x62, 0xf1,
	0x4b, 0xad, 0x28, 0x11, 0x12, 0xd2, 0x2b, 0x2a, 0x48, 0xd6, 0x83, 0x3b, 0x47, 0xcf, 0x3b, 0xb5,
	0x52, 0x4a, 0x04, 0x4d, 0x0c, 0xe2, 0x08, 0x2b, 0x2e, 0x76, 0xf7, 0xd4, 0x12, 0xb0, 0x15, 0x52,
	0x68, 0x0a, 0x90, 0x9d, 0x21, 0x8c, 0x26, 0x07, 0xd2, 0x90, 0xc4, 0x06, 0xf6, 0xbd, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0xdc, 0x56, 0xa6, 0xee, 0x0c, 0x01, 0x00, 0x00,
}
