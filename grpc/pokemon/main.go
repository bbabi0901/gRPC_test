package main

import (
	"context"
	"log"
	"net"

	db "github.com/bbabi0901/learngo/gRPC/gRPCPokemon/db"
	pkpb "github.com/bbabi0901/learngo/gRPC/gRPCPokemon/proto/pokemon"
	"github.com/bbabi0901/learngo/gRPC/gRPCPokemon/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const portNumber = "8888"

type pokemonServer struct {
	pkpb.PokemonServer
}

func (s *pokemonServer) ListPokemons(ctx context.Context, req *pkpb.ListPokemonsRequest) (*pkpb.ListPokemonsResponse, error) {
	return &pkpb.ListPokemonsResponse{PokemonInfos: db.PokemonInfos()}, nil
}

func (s *pokemonServer) GetPokemonByID(ctx context.Context, req *pkpb.GetPokemonByIDRequest) (*pkpb.GetPokemonByIDResponse, error) {
	var pokemonInfo *pkpb.PokemonInfo

	for _, p := range db.PokemonInfos() {
		if p.PokemonId == req.PokemonId {
			pokemonInfo = p
			break
		}
	}
	if pokemonInfo == nil {
		return nil, status.Error(codes.NotFound, "Pokemon is not found.")
	}
	return &pkpb.GetPokemonByIDResponse{PokemonInfo: pokemonInfo}, nil
}

func (s *pokemonServer) GetPokemonByName(ctx context.Context, req *pkpb.GetPokemonByNameRequest) (*pkpb.GetPokemonByNameResponse, error) {
	var pokemonInfo *pkpb.PokemonInfo

	for _, p := range db.PokemonInfos() {
		if p.PokemonName == req.PokemonName {
			pokemonInfo = p
			break
		}
	}
	if pokemonInfo == nil {
		return nil, status.Error(codes.NotFound, "Pokemon is not found.")
	}
	return &pkpb.GetPokemonByNameResponse{PokemonInfo: pokemonInfo}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	utils.HandleErr(err)

	grpcServer := grpc.NewServer()
	pkpb.RegisterPokemonServer(grpcServer, &pokemonServer{})

	log.Printf("Start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
