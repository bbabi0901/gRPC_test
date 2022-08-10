package grpc_client

import (
	"sync"

	"google.golang.org/grpc"

	abpb "github.com/bbabi0901/learngo/gRPC/gRPCPokemon/proto/ability"
	pkpb "github.com/bbabi0901/learngo/gRPC/gRPCPokemon/proto/pokemon"
	typepb "github.com/bbabi0901/learngo/gRPC/gRPCPokemon/proto/type"
	"github.com/bbabi0901/learngo/gRPC/gRPCPokemon/utils"
)

var (
	once sync.Once
	cli  pkpb.PokemonClient
)

type PokemonRPC struct {
	Conn       *grpc.ClientConn
	PokemonCli pkpb.PokemonClient
	ChunkSize  int
}

type AbilityRPC struct {
	Conn       *grpc.ClientConn
	AbilityCli abpb.AbilityClient
	chunkSize  int
}

type TypeRPC struct {
	Conn      *grpc.ClientConn
	TypeCli   typepb.TypeClient
	ChunkSize int
}

func GetPokemonRPC(serviceHost string) (c PokemonRPC, err error) {
	once.Do(func() {
		c.Conn, err = grpc.Dial(
			serviceHost,
			grpc.WithBlock(),
			grpc.WithInsecure(),
		)
		utils.HandleErr(err)
		c.PokemonCli = pkpb.NewPokemonClient(c.Conn)
	})
	return
}

func GetAbilityRPC(serviceHost string) (c AbilityRPC, err error) {
	once.Do(func() {
		c.Conn, err = grpc.Dial(
			serviceHost,
			grpc.WithBlock(),
			grpc.WithInsecure(),
		)
		utils.HandleErr(err)
		c.AbilityCli = abpb.NewAbilityClient(c.Conn)
	})
	return
}

func GetTypeRPC(serviceHost string) (c TypeRPC, err error) {
	once.Do(func() {
		c.Conn, err = grpc.Dial(
			serviceHost,
			grpc.WithBlock(),
			grpc.WithInsecure(),
		)
		utils.HandleErr(err)
		c.TypeCli = typepb.NewTypeClient(c.Conn)
	})
	return
}
