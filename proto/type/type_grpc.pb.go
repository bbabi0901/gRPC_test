// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: proto/type/type.proto

package _type

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

// TypeClient is the client API for Type service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TypeClient interface {
	ListTypes(ctx context.Context, in *ListTypesRequest, opts ...grpc.CallOption) (*ListTypesResponse, error)
	GetPokemonsOfType(ctx context.Context, in *GetPokemonsOfTypeRequest, opts ...grpc.CallOption) (*GetPokemonsOfTypeResponse, error)
	GetTypeInfo(ctx context.Context, in *GetTypeInfoRequest, opts ...grpc.CallOption) (*GetTypeInfoResponse, error)
}

type typeClient struct {
	cc grpc.ClientConnInterface
}

func NewTypeClient(cc grpc.ClientConnInterface) TypeClient {
	return &typeClient{cc}
}

func (c *typeClient) ListTypes(ctx context.Context, in *ListTypesRequest, opts ...grpc.CallOption) (*ListTypesResponse, error) {
	out := new(ListTypesResponse)
	err := c.cc.Invoke(ctx, "/type.Type/ListTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *typeClient) GetPokemonsOfType(ctx context.Context, in *GetPokemonsOfTypeRequest, opts ...grpc.CallOption) (*GetPokemonsOfTypeResponse, error) {
	out := new(GetPokemonsOfTypeResponse)
	err := c.cc.Invoke(ctx, "/type.Type/GetPokemonsOfType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *typeClient) GetTypeInfo(ctx context.Context, in *GetTypeInfoRequest, opts ...grpc.CallOption) (*GetTypeInfoResponse, error) {
	out := new(GetTypeInfoResponse)
	err := c.cc.Invoke(ctx, "/type.Type/GetTypeInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TypeServer is the server API for Type service.
// All implementations must embed UnimplementedTypeServer
// for forward compatibility
type TypeServer interface {
	ListTypes(context.Context, *ListTypesRequest) (*ListTypesResponse, error)
	GetPokemonsOfType(context.Context, *GetPokemonsOfTypeRequest) (*GetPokemonsOfTypeResponse, error)
	GetTypeInfo(context.Context, *GetTypeInfoRequest) (*GetTypeInfoResponse, error)
	mustEmbedUnimplementedTypeServer()
}

// UnimplementedTypeServer must be embedded to have forward compatible implementations.
type UnimplementedTypeServer struct {
}

func (UnimplementedTypeServer) ListTypes(context.Context, *ListTypesRequest) (*ListTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTypes not implemented")
}
func (UnimplementedTypeServer) GetPokemonsOfType(context.Context, *GetPokemonsOfTypeRequest) (*GetPokemonsOfTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPokemonsOfType not implemented")
}
func (UnimplementedTypeServer) GetTypeInfo(context.Context, *GetTypeInfoRequest) (*GetTypeInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTypeInfo not implemented")
}
func (UnimplementedTypeServer) mustEmbedUnimplementedTypeServer() {}

// UnsafeTypeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TypeServer will
// result in compilation errors.
type UnsafeTypeServer interface {
	mustEmbedUnimplementedTypeServer()
}

func RegisterTypeServer(s grpc.ServiceRegistrar, srv TypeServer) {
	s.RegisterService(&Type_ServiceDesc, srv)
}

func _Type_ListTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypeServer).ListTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/type.Type/ListTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypeServer).ListTypes(ctx, req.(*ListTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Type_GetPokemonsOfType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPokemonsOfTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypeServer).GetPokemonsOfType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/type.Type/GetPokemonsOfType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypeServer).GetPokemonsOfType(ctx, req.(*GetPokemonsOfTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Type_GetTypeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTypeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypeServer).GetTypeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/type.Type/GetTypeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypeServer).GetTypeInfo(ctx, req.(*GetTypeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Type_ServiceDesc is the grpc.ServiceDesc for Type service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Type_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "type.Type",
	HandlerType: (*TypeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTypes",
			Handler:    _Type_ListTypes_Handler,
		},
		{
			MethodName: "GetPokemonsOfType",
			Handler:    _Type_GetPokemonsOfType_Handler,
		},
		{
			MethodName: "GetTypeInfo",
			Handler:    _Type_GetTypeInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/type/type.proto",
}
