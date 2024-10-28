package graceful

import (
	"context"
	"log/slog"

	"golang.org/x/sync/errgroup"
)

type starter interface {
	Start(ctx context.Context) error
}

type graceful struct {
	processes []process
	logger    *slog.Logger
}

func New(processes ...process) *graceful {
	return &graceful{
		processes: processes,
		logger:    slog.Default(),
	}
}

func (gr *graceful) Start(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	for _, process := range gr.processes {
		process := process // TODO remove if go > 1.22

		if process.disabled {
			continue
		}

		f := func() error {
			err := process.starter.Start(ctx)
			if err != nil {
				gr.logger.Error(err.Error())
				gr.logger.Info("Start graceful shutdown")

				return err
			}

			return nil
		}

		g.Go(f)
	}

	err := g.Wait()
	if err != nil {
		gr.logger.Info("Application stopped gracefully")

		return err
	}

	gr.logger.Info("Every process stopped by itself with no error")

	return nil
}

func (gr *graceful) SetLogger(l *slog.Logger) {
	gr.logger = l
}
