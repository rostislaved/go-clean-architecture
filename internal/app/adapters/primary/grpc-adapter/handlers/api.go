package handlers

import (
	"context"

	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/grpc-adapter/generated"
)

type Server struct {
	generated.UnimplementedApiServer
}

func (s Server) SendMessage(ctx context.Context, message *generated.Message) (*generated.Message, error) {
	m := generated.Message{Body: "Ответ"}
	// Тут вызов сервиса

	return &m, nil
}
