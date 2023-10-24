package grpcAdapter

import (
	"log"
	"net"

	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/grpc-adapter/controller"
	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/grpc-adapter/generated"

	"google.golang.org/grpc"
)

type GrpcAdapter struct {
	start func() error
}

func New() *GrpcAdapter {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	generated.RegisterApiServer(server, controller.Server{})

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
