package clickhouse

import (
	"crypto/tls"
	"fmt"
	"log/slog"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"

	"github.com/rostislaved/go-clean-architecture/internal/pkg/helpers"
)

func New(l *slog.Logger, cfg Config) driver.Conn {
	currentHostString := fmt.Sprintf("db host: [%s:%s].", cfg.Host, cfg.Port)

	l.Info(currentHostString+" Подключение...", "source", helpers.GetFunctionName())

	options := clickhouse.Options{
		Addr: []string{"<CLICKHOUSE_SECURE_NATIVE_HOSTNAME>:9440"},
		Auth: clickhouse.Auth{
			Database: cfg.Name,
			Username: cfg.User,
			Password: cfg.Password,
		},
		ClientInfo: clickhouse.ClientInfo{
			Products: []struct {
				Name    string
				Version string
			}{
				{Name: "an-example-go-client", Version: "0.1"},
			},
		},

		Debugf: func(format string, v ...interface{}) {
			fmt.Printf(format, v)
		},
		TLS: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	conn, err := clickhouse.Open(&options)
	if err != nil {
		l.Error(err.Error(), "source", helpers.GetFunctionName())

		panic(err)
	}

	l.Info(currentHostString+" Подключено!", "source", helpers.GetFunctionName())

	return conn
}

type Config struct {
	Type       string
	Host       string `config:"envVar"`
	Port       string `config:"envVar"`
	User       string `config:"envVar"`
	Password   string `config:"envVar"`
	Name       string
	Procedures map[string]string
}
