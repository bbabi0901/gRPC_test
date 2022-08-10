package main

import (
	"context"
	"log"
	"net"

	"github.com/bbabi0901/learngo/gRPC/gRPCPokemon/db"
	client "github.com/bbabi0901/learngo/gRPC/gRPCPokemon/grpc"
	abpb "github.com/bbabi0901/learngo/gRPC/gRPCPokemon/proto/ability"
	pkpb "github.com/bbabi0901/learngo/gRPC/gRPCPokemon/proto/pokemon"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	portNumber            = "8889"
	pokemonServerEndpoint = "localhost:8888"
)

type abilityServer struct {
	abpb.AbilityServer

	pokemonCli pkpb.PokemonClient
}

func (s *abilityServer) ListAbilities(ctx context.Context, req *abpb.ListAbilitiesRequeset) (*abpb.ListAbilitiesResponse, error) {
	var abilityInfos []*abpb.AbilityInfo

	for _, a := range db.AbilityInfos() {
		resp, err := s.pokemonCli.ListPokemons(ctx, &pkpb.ListPokemonsRequest{})
		if err != nil {
			log.Fatalf("Failed to listen Pokemon Server: %v", err)
		}
		for _, p := range resp.PokemonInfos {
			for _, pa := range p.Ability {
				if a.AbilityName == pa {
					a.PokemonName = append(a.PokemonName, p.PokemonName)
				}
			}
		}
		abilityInfos = append(abilityInfos, a)
	}
	return &abpb.ListAbilitiesResponse{AbilityInfos: abilityInfos}, nil
}

func (s *abilityServer) GetAbilityByID(ctx context.Context, req *abpb.GetAbilityByIDRequest) (*abpb.GetAbilityByIDResponse, error) {
	var abilityInfo *abpb.AbilityInfo

	for _, a := range db.AbilityInfos() {
		if a.AbilityId == req.AbilityId {
			abilityInfo = a
			resp, err := s.pokemonCli.ListPokemons(ctx, &pkpb.ListPokemonsRequest{})
			if err != nil {
				log.Fatalf("Failed to listen Pokemon Server: %v", err)
			}
			for _, p := range resp.PokemonInfos {
				for _, pa := range p.Ability {
					if a.AbilityName == pa {
						abilityInfo.PokemonName = append(abilityInfo.PokemonName, p.PokemonName)
					}
				}
			}
		}
	}
	if abilityInfo == nil {
		return nil, status.Error(codes.NotFound, "Ability is not found.")
	}
	return &abpb.GetAbilityByIDResponse{AbilityInfo: abilityInfo}, nil
}

func (s *abilityServer) GetAbilityByName(ctx context.Context, req *abpb.GetAbilityByNameRequest) (*abpb.GetAbilityByNameResponse, error) {
	var abilityInfo *abpb.AbilityInfo

	for _, a := range db.AbilityInfos() {
		if a.AbilityName == req.AbilityName {
			abilityInfo = a
			resp, err := s.pokemonCli.ListPokemons(ctx, &pkpb.ListPokemonsRequest{})
			if err != nil {
				log.Fatalf("Failed to listen Pokemon Server: %v", err)
			}
			for _, p := range resp.PokemonInfos {
				for _, pa := range p.Ability {
					if a.AbilityName == pa {
						abilityInfo.PokemonName = append(abilityInfo.PokemonName, p.PokemonName)
					}
				}
			}
		}
	}
	return &abpb.GetAbilityByNameResponse{AbilityInfo: abilityInfo}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	p, _ := client.GetPokemonRPC(pokemonServerEndpoint)
	abpb.RegisterAbilityServer(grpcServer, &abilityServer{pokemonCli: p.PokemonCli})

	log.Printf("Start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
