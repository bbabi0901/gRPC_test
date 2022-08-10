package main

import (
	"context"
	"log"
	"net"

	db "github.com/bbabi0901/learngo/gRPC/gRPCPokemon/db"
	client "github.com/bbabi0901/learngo/gRPC/gRPCPokemon/grpc"
	pkpb "github.com/bbabi0901/learngo/gRPC/gRPCPokemon/proto/pokemon"
	typepb "github.com/bbabi0901/learngo/gRPC/gRPCPokemon/proto/type"
	"github.com/bbabi0901/learngo/gRPC/gRPCPokemon/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	portNumber            = "6666"
	pokemonServerEndpoint = "localhost:8888"
)

type typeServer struct {
	typepb.TypeServer

	pokemonCli pkpb.PokemonClient
}

func (s *typeServer) ListTypes(ctx context.Context, req *typepb.ListTypesRequest) (*typepb.ListTypesResponse, error) {
	return &typepb.ListTypesResponse{TypeInfos: db.TypeInfos()}, nil
}

func (s *typeServer) GetPokemonsOfType(ctx context.Context, req *typepb.GetPokemonsOfTypeRequest) (*typepb.GetPokemonsOfTypeResponse, error) {
	var typePokemonInfo *typepb.PokemonsOfType

	for _, t := range db.TypePokemonInfos() {
		if t.TypeName == req.TypeName {
			typePokemonInfo = t
			typePokemonInfo.PokemonName = nil

			resp, err := s.pokemonCli.ListPokemons(ctx, &pkpb.ListPokemonsRequest{})
			if err != nil {
				log.Fatalf("Failed to listen Pokemon Server: %v", err)
			}

			for _, p := range resp.PokemonInfos {
				for _, pt := range p.Type {
					if t.TypeName == pt {
						typePokemonInfo.PokemonName = append(typePokemonInfo.PokemonName, p.PokemonName)
					}
				}
			}

		}
	}
	if typePokemonInfo == nil {
		return nil, status.Error(codes.NotFound, "Type is not found.")
	}
	return &typepb.GetPokemonsOfTypeResponse{PokemonsOfType: typePokemonInfo}, nil
}

func (s *typeServer) GetTypeInfo(ctx context.Context, req *typepb.GetTypeInfoRequest) (*typepb.GetTypeInfoResponse, error) {
	var typeInfo *typepb.TypeInfo

	for _, t := range db.TypeInfos() {
		if t.TypeName == req.TypeName {
			typeInfo = t
			break
		}
	}
	if typeInfo == nil {
		return nil, status.Error(codes.NotFound, "Type is not found.")
	}
	return &typepb.GetTypeInfoResponse{TypeInfo: typeInfo}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	utils.HandleErr(err)

	grpcServer := grpc.NewServer()
	p, _ := client.GetPokemonRPC(pokemonServerEndpoint)
	typepb.RegisterTypeServer(grpcServer, &typeServer{pokemonCli: p.PokemonCli})

	log.Printf("Start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
