package httpAdapter

import (
	"context"
	"log"

	"github.com/rostislaved/go-clean-architecture/internal/pkg/helpers"
)

func (a *HttpAdapter) Start() {
	startMsg := "Сервер запущен."

	log.Println(startMsg)
	a.logger.Info(startMsg, helpers.GetFunctionName())

	a.notify <- a.server.ListenAndServe()
	close(a.notify)
}

func (a *HttpAdapter) Notify() <-chan error {
	return a.notify
}

func (a *HttpAdapter) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), a.shutdownTimeout)
	defer cancel()

	err := a.server.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}
