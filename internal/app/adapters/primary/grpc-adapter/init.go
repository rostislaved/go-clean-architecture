package grpc_adapter

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/grpc-adapter/generated"
	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/grpc-adapter/handlers"
)

type GrpcAdapter struct {
	start func() error
}

func New() *GrpcAdapter {
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	generated.RegisterApiServer(server, handlers.Server{})

	startFunc := func() error {
		err = server.Serve(listener)

		return err
	}

	return &GrpcAdapter{
		start: startFunc,
	}
}

func (a GrpcAdapter) Start() {
	err := a.start()
	if err != nil {
		log.Fatal(err)
	}
}
