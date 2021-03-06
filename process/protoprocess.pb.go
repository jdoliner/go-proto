// Code generated by protoc-gen-go.
// source: process/protoprocess.proto
// DO NOT EDIT!

package protoprocess

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"

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
const _ = proto.ProtoPackageIsVersion1

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for API service

type APIClient interface {
	Do(ctx context.Context, opts ...grpc.CallOption) (API_DoClient, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Do(ctx context.Context, opts ...grpc.CallOption) (API_DoClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_API_serviceDesc.Streams[0], c.cc, "/protoprocess.API/Do", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIDoClient{stream}
	return x, nil
}

type API_DoClient interface {
	Send(*google_protobuf.BytesValue) error
	Recv() (*google_protobuf.BytesValue, error)
	grpc.ClientStream
}

type aPIDoClient struct {
	grpc.ClientStream
}

func (x *aPIDoClient) Send(m *google_protobuf.BytesValue) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aPIDoClient) Recv() (*google_protobuf.BytesValue, error) {
	m := new(google_protobuf.BytesValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for API service

type APIServer interface {
	Do(API_DoServer) error
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Do_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(APIServer).Do(&aPIDoServer{stream})
}

type API_DoServer interface {
	Send(*google_protobuf.BytesValue) error
	Recv() (*google_protobuf.BytesValue, error)
	grpc.ServerStream
}

type aPIDoServer struct {
	grpc.ServerStream
}

func (x *aPIDoServer) Send(m *google_protobuf.BytesValue) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aPIDoServer) Recv() (*google_protobuf.BytesValue, error) {
	m := new(google_protobuf.BytesValue)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protoprocess.API",
	HandlerType: (*APIServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Do",
			Handler:       _API_Do_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x28, 0xca, 0x4f,
	0x4e, 0x2d, 0x2e, 0xd6, 0x07, 0xd2, 0x25, 0xf9, 0x50, 0x8e, 0x1e, 0x98, 0x23, 0xc4, 0x83, 0x2c,
	0x26, 0x25, 0x97, 0x9e, 0x9f, 0x9f, 0x9e, 0x93, 0x0a, 0x51, 0x98, 0x54, 0x9a, 0xa6, 0x5f, 0x5e,
	0x94, 0x58, 0x50, 0x90, 0x5a, 0x04, 0x55, 0x6d, 0xe4, 0xcd, 0xc5, 0xec, 0x18, 0xe0, 0x29, 0xe4,
	0xc2, 0xc5, 0xe4, 0x92, 0x2f, 0x24, 0xad, 0x07, 0x51, 0xad, 0x07, 0x53, 0xad, 0xe7, 0x54, 0x59,
	0x92, 0x5a, 0x1c, 0x96, 0x98, 0x53, 0x9a, 0x2a, 0x85, 0x4f, 0x52, 0x89, 0x41, 0x83, 0xd1, 0x80,
	0x31, 0x89, 0x0d, 0x2c, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x44, 0x68, 0x7e, 0x7b, 0x9f,
	0x00, 0x00, 0x00,
}
