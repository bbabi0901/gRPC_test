package main

import (
	"context"
	"log"
	"net/http"

	abpb "github.com/bbabi0901/learngo/gRPC/gRPCPokemon/proto/ability"
	pkpb "github.com/bbabi0901/learngo/gRPC/gRPCPokemon/proto/pokemon"
	typepb "github.com/bbabi0901/learngo/gRPC/gRPCPokemon/proto/type"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

const (
	portNumber              = "8887"
	abTypeServerPortNumber  = "8889"
	pokemonServerPortNumber = "8888"
	abilityServerEndpoint   = "localhost:" + abTypeServerPortNumber
	pokemonServerEndpoint   = "localhost:" + pokemonServerPortNumber
)

func main() {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	options := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	pkConnection(ctx, mux, pokemonServerEndpoint, options)
	abConnection(ctx, mux, abilityServerEndpoint, options)
	typeConnection(ctx, mux, abilityServerEndpoint, options)

	log.Printf("Start HTTP server on %s port", portNumber)
	if err := http.ListenAndServe(":"+portNumber, mux); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}

func pkConnection(ctx context.Context, mux *runtime.ServeMux, endpoint string, options []grpc.DialOption) {
	err := pkpb.RegisterPokemonHandlerFromEndpoint(ctx, mux, endpoint, options)
	if err != nil {
		log.Fatalf("Failed to register gRPC gateway: %v", err)
	}
}

func abConnection(ctx context.Context, mux *runtime.ServeMux, endpoint string, options []grpc.DialOption) {
	err := abpb.RegisterAbilityHandlerFromEndpoint(ctx, mux, endpoint, options)
	if err != nil {
		log.Fatalf("Failed to register gRPC gateway: %v", err)
	}
}

func typeConnection(ctx context.Context, mux *runtime.ServeMux, endpoint string, options []grpc.DialOption) {
	err := typepb.RegisterTypeHandlerFromEndpoint(ctx, mux, endpoint, options)
	if err != nil {
		log.Fatalf("Failed to register gRPC gateway: %v", err)
	}
}
