// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: user.proto

package user

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

// DataUserClient is the client API for DataUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataUserClient interface {
	GenerateJwtByToken(ctx context.Context, in *TokenGrand, opts ...grpc.CallOption) (*TokenJwt, error)
}

type dataUserClient struct {
	cc grpc.ClientConnInterface
}

func NewDataUserClient(cc grpc.ClientConnInterface) DataUserClient {
	return &dataUserClient{cc}
}

func (c *dataUserClient) GenerateJwtByToken(ctx context.Context, in *TokenGrand, opts ...grpc.CallOption) (*TokenJwt, error) {
	out := new(TokenJwt)
	err := c.cc.Invoke(ctx, "/user.DataUser/GenerateJwtByToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataUserServer is the server API for DataUser service.
// All implementations must embed UnimplementedDataUserServer
// for forward compatibility
type DataUserServer interface {
	GenerateJwtByToken(context.Context, *TokenGrand) (*TokenJwt, error)
	mustEmbedUnimplementedDataUserServer()
}

// UnimplementedDataUserServer must be embedded to have forward compatible implementations.
type UnimplementedDataUserServer struct {
}

func (UnimplementedDataUserServer) GenerateJwtByToken(context.Context, *TokenGrand) (*TokenJwt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateJwtByToken not implemented")
}
func (UnimplementedDataUserServer) mustEmbedUnimplementedDataUserServer() {}

// UnsafeDataUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataUserServer will
// result in compilation errors.
type UnsafeDataUserServer interface {
	mustEmbedUnimplementedDataUserServer()
}

func RegisterDataUserServer(s grpc.ServiceRegistrar, srv DataUserServer) {
	s.RegisterService(&DataUser_ServiceDesc, srv)
}

func _DataUser_GenerateJwtByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenGrand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataUserServer).GenerateJwtByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.DataUser/GenerateJwtByToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataUserServer).GenerateJwtByToken(ctx, req.(*TokenGrand))
	}
	return interceptor(ctx, in, info, handler)
}

// DataUser_ServiceDesc is the grpc.ServiceDesc for DataUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.DataUser",
	HandlerType: (*DataUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateJwtByToken",
			Handler:    _DataUser_GenerateJwtByToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
