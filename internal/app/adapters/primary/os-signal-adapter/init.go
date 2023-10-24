package osSignalAdapter

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type OsSignalAdapter struct {
	chError chan error
}

func New() *OsSignalAdapter {
	return &OsSignalAdapter{
		chError: make(chan error, 1),
	}
}

func (a *OsSignalAdapter) Start() {
	osSignCh := make(chan os.Signal, 1)

	signal.Notify(osSignCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	sig := <-osSignCh
	err := fmt.Errorf("\nполучен сигнал %s, завершаю работу\n", sig.String()) //nolint:stylecheck

	a.chError <- err
}

func (a *OsSignalAdapter) Notify() <-chan error {
	return a.chError
}
