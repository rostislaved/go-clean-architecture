package postgres

import (
	"fmt"
	"time"
)

func postgresConnectionString(host, port, user, password, name, timeZone string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		host, port, user, password, name, timeZone)
}

type Config struct {
	Host     string `config:"envVar"`
	Port     string `config:"envVar"`
	Name     string
	User     string `config:"envVar"`
	Password string `config:"envVar"`
	TimeZone string

	MaxOpenConns    *int
	MaxIdleConns    *int
	ConnMaxLifetime *time.Duration
	ConnMaxIdleTime *time.Duration
}

func getLogString(cfg Config) (connectingString, connectedString string) {
	currentHostString := fmt.Sprintf("db host: [%s:%s].", cfg.Host, cfg.Port)

	connectingString = currentHostString + " Подключение..."
	connectedString = currentHostString + " Подключено!"

	return connectingString, connectedString
}
