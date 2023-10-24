package grpcAdapter

import (
	"context"
	"time"

	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/secondary/grpc-adapter/generated"
)

func (a GrpcAdapter) SendMessage(msg string) (resp string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	message := &generated.Message{Body: msg}

	r, err := a.client.SendMessage(ctx, message)
	if err != nil {
		return "", err
	}

	return r.Body, nil
}
