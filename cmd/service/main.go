package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/rostislaved/graceful"

	_ "go.uber.org/automaxprocs"

	"github.com/rostislaved/go-clean-architecture/internal/app"
	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/os-signal-adapter"
	"github.com/rostislaved/go-clean-architecture/internal/app/config"
	"github.com/rostislaved/go-clean-architecture/internal/pkg/helpers"
)

func main() {
	cfg := config.New()

	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})
	l := slog.New(h)

	app := app.New(l, cfg)

	gr := graceful.New(
		graceful.NewProcess(os_signal_adapter.New()),
		graceful.NewProcess(app.HttpAdapter),
		graceful.NewProcess(app.GrpcAdapter),
		graceful.NewProcess(app.PprofAdapter),
		graceful.NewProcess(app.NatsAdapterSubscriber),
		graceful.NewProcess(app.KafkaAdapterSubscriber),
	)

	err := gr.Start(context.Background())
	if err != nil {
		l.Error(err.Error(), "source", helpers.GetFunctionName())

		panic(err)
	}
}
