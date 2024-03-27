// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: dice.proto

package pb

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

// RollDicerClient is the client API for RollDicer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RollDicerClient interface {
	RollDice(ctx context.Context, in *RollDiceRequest, opts ...grpc.CallOption) (*RollDiceReply, error)
}

type rollDicerClient struct {
	cc grpc.ClientConnInterface
}

func NewRollDicerClient(cc grpc.ClientConnInterface) RollDicerClient {
	return &rollDicerClient{cc}
}

func (c *rollDicerClient) RollDice(ctx context.Context, in *RollDiceRequest, opts ...grpc.CallOption) (*RollDiceReply, error) {
	out := new(RollDiceReply)
	err := c.cc.Invoke(ctx, "/proto.RollDicer/RollDice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RollDicerServer is the server API for RollDicer service.
// All implementations must embed UnimplementedRollDicerServer
// for forward compatibility
type RollDicerServer interface {
	RollDice(context.Context, *RollDiceRequest) (*RollDiceReply, error)
	mustEmbedUnimplementedRollDicerServer()
}

// UnimplementedRollDicerServer must be embedded to have forward compatible implementations.
type UnimplementedRollDicerServer struct {
}

func (UnimplementedRollDicerServer) RollDice(context.Context, *RollDiceRequest) (*RollDiceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollDice not implemented")
}
func (UnimplementedRollDicerServer) mustEmbedUnimplementedRollDicerServer() {}

// UnsafeRollDicerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RollDicerServer will
// result in compilation errors.
type UnsafeRollDicerServer interface {
	mustEmbedUnimplementedRollDicerServer()
}

func RegisterRollDicerServer(s grpc.ServiceRegistrar, srv RollDicerServer) {
	s.RegisterService(&RollDicer_ServiceDesc, srv)
}

func _RollDicer_RollDice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollDiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RollDicerServer).RollDice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RollDicer/RollDice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RollDicerServer).RollDice(ctx, req.(*RollDiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RollDicer_ServiceDesc is the grpc.ServiceDesc for RollDicer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RollDicer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RollDicer",
	HandlerType: (*RollDicerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RollDice",
			Handler:    _RollDicer_RollDice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dice.proto",
}