package main

import (
	"context"
	"log/slog"
	"os"

	_ "go.uber.org/automaxprocs"

	"github.com/rostislaved/go-clean-architecture/internal/app"
	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/os-signal-adapter"
	"github.com/rostislaved/go-clean-architecture/internal/app/config"
	"github.com/rostislaved/go-clean-architecture/internal/libs/graceful"
	"github.com/rostislaved/go-clean-architecture/internal/libs/helpers"
)

func main() {
	cfg := config.New()

	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	l := slog.New(h)

	app := app.New(l, cfg)

	gr := graceful.New(
		graceful.NewProcess(app.HttpAdapter),
		graceful.NewProcess(app.PprofAdapter),
		graceful.NewProcess(app.NatsAdapterSubscriber),
		graceful.NewProcess(app.KafkaAdapterSubscriber),
		graceful.NewProcess(os_signal_adapter.New()),
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := gr.Start(ctx)
	if err != nil {
		l.Error(err.Error(), "source", helpers.GetFunctionName())

		panic(err)
	}
}
