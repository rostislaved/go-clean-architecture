package repohelpers

import (
	"fmt"
	"log"
	"strings"
)

const (
	mssql      = "mssql"
	postgres   = "postgres"
	clickhouse = "clickhouse"
)

// connectionString := repohelpers.GetConnectionString(cfg.Type, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
func GetConnectionString(Type, Host, Port, User, Password, Name string) (connectionString string) {
	switch Type {
	case mssql:
		if strings.Contains(Host, "\\") {
			connectionString = fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s",
				Host, User, Password, Name)
		} else {
			connectionString = fmt.Sprintf("server=%s;port=%s;user id=%s;password=%s;database=%s",
				Host, Port, User, Password, Name)
		}

		return
	case postgres:
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
			Host, Port, User, Password, Name, "Europe/Moscow")
	case clickhouse:
		return fmt.Sprintf(
			"http://%s:%s@%s:%s/%s",
			User,
			Password,
			Host,
			Port,
			Name,
		)
	default:
		log.Fatal("Неверный тип БД")
	}

	return
}
