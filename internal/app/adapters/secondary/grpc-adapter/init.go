package grpcAdapter

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/grpc-adapter/generated"
)

type GrpcAdapter struct {
	client generated.ApiClient
}

func New() *GrpcAdapter {
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := generated.NewApiClient(conn)

	return &GrpcAdapter{
		client: client,
	}
}
