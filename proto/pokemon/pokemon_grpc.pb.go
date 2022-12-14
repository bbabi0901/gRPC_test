// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: proto/pokemon/pokemon.proto

package pokemon

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

// PokemonClient is the client API for Pokemon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PokemonClient interface {
	ListPokemons(ctx context.Context, in *ListPokemonsRequest, opts ...grpc.CallOption) (*ListPokemonsResponse, error)
	GetPokemonByID(ctx context.Context, in *GetPokemonByIDRequest, opts ...grpc.CallOption) (*GetPokemonByIDResponse, error)
	GetPokemonByName(ctx context.Context, in *GetPokemonByNameRequest, opts ...grpc.CallOption) (*GetPokemonByNameResponse, error)
}

type pokemonClient struct {
	cc grpc.ClientConnInterface
}

func NewPokemonClient(cc grpc.ClientConnInterface) PokemonClient {
	return &pokemonClient{cc}
}

func (c *pokemonClient) ListPokemons(ctx context.Context, in *ListPokemonsRequest, opts ...grpc.CallOption) (*ListPokemonsResponse, error) {
	out := new(ListPokemonsResponse)
	err := c.cc.Invoke(ctx, "/pokemon.Pokemon/ListPokemons", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonClient) GetPokemonByID(ctx context.Context, in *GetPokemonByIDRequest, opts ...grpc.CallOption) (*GetPokemonByIDResponse, error) {
	out := new(GetPokemonByIDResponse)
	err := c.cc.Invoke(ctx, "/pokemon.Pokemon/GetPokemonByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonClient) GetPokemonByName(ctx context.Context, in *GetPokemonByNameRequest, opts ...grpc.CallOption) (*GetPokemonByNameResponse, error) {
	out := new(GetPokemonByNameResponse)
	err := c.cc.Invoke(ctx, "/pokemon.Pokemon/GetPokemonByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PokemonServer is the server API for Pokemon service.
// All implementations must embed UnimplementedPokemonServer
// for forward compatibility
type PokemonServer interface {
	ListPokemons(context.Context, *ListPokemonsRequest) (*ListPokemonsResponse, error)
	GetPokemonByID(context.Context, *GetPokemonByIDRequest) (*GetPokemonByIDResponse, error)
	GetPokemonByName(context.Context, *GetPokemonByNameRequest) (*GetPokemonByNameResponse, error)
	mustEmbedUnimplementedPokemonServer()
}

// UnimplementedPokemonServer must be embedded to have forward compatible implementations.
type UnimplementedPokemonServer struct {
}

func (UnimplementedPokemonServer) ListPokemons(context.Context, *ListPokemonsRequest) (*ListPokemonsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPokemons not implemented")
}
func (UnimplementedPokemonServer) GetPokemonByID(context.Context, *GetPokemonByIDRequest) (*GetPokemonByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPokemonByID not implemented")
}
func (UnimplementedPokemonServer) GetPokemonByName(context.Context, *GetPokemonByNameRequest) (*GetPokemonByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPokemonByName not implemented")
}
func (UnimplementedPokemonServer) mustEmbedUnimplementedPokemonServer() {}

// UnsafePokemonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PokemonServer will
// result in compilation errors.
type UnsafePokemonServer interface {
	mustEmbedUnimplementedPokemonServer()
}

func RegisterPokemonServer(s grpc.ServiceRegistrar, srv PokemonServer) {
	s.RegisterService(&Pokemon_ServiceDesc, srv)
}

func _Pokemon_ListPokemons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPokemonsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServer).ListPokemons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.Pokemon/ListPokemons",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServer).ListPokemons(ctx, req.(*ListPokemonsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pokemon_GetPokemonByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPokemonByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServer).GetPokemonByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.Pokemon/GetPokemonByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServer).GetPokemonByID(ctx, req.(*GetPokemonByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pokemon_GetPokemonByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPokemonByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServer).GetPokemonByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokemon.Pokemon/GetPokemonByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServer).GetPokemonByName(ctx, req.(*GetPokemonByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Pokemon_ServiceDesc is the grpc.ServiceDesc for Pokemon service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pokemon_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pokemon.Pokemon",
	HandlerType: (*PokemonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPokemons",
			Handler:    _Pokemon_ListPokemons_Handler,
		},
		{
			MethodName: "GetPokemonByID",
			Handler:    _Pokemon_GetPokemonByID_Handler,
		},
		{
			MethodName: "GetPokemonByName",
			Handler:    _Pokemon_GetPokemonByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/pokemon/pokemon.proto",
}
