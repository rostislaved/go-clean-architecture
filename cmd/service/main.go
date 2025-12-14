package main

import (
	_ "go.uber.org/automaxprocs"

	"github.com/rostislaved/go-clean-architecture/internal/app"
	"github.com/rostislaved/go-clean-architecture/internal/pkg/logger"
)

func main() {
	logger := logger.New()

	app, err := app.New(logger)
	if err != nil {
		logger.Error("Failed to create application", "err", err)

		return
	}

	err = app.Start()
	if err != nil {
		logger.Error("Failed to start application", "err", err)

		return
	}
}
