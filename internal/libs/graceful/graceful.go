package graceful

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/sync/errgroup"
)

type starter interface {
	Start(ctx context.Context) error
}

type graceful struct {
	processes []process
}

func New(processes ...process) *graceful {
	return &graceful{
		processes: processes,
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
				log.Println(err)
				log.Println("Start graceful shutdown")

				return err
			}

			return nil
		}

		g.Go(f)
	}

	err := g.Wait()
	if err != nil {
		fmt.Println("Application stopped gracefully")
		return err
	}

	return nil
}
