package config

import (
	"github.com/rostislaved/go-clean-architecture/internal/app/application/usecases"
	"github.com/rostislaved/go-clean-architecture/internal/app/infrastructure/clickhouse"
	"github.com/rostislaved/go-clean-architecture/internal/app/infrastructure/mongo"
	"github.com/rostislaved/go-clean-architecture/internal/app/infrastructure/postgres"
)

type Config struct {
	Info           Info
	Application    Application
	Adapters       Adapters
	Infrastructure Infrastructure
}
type Info struct {
	Name    string
	Version string
}

type Application struct {
	UseCases usecases.Config
}

type Infrastructure struct {
	Databases Databases
}

type Databases struct {
	Postgres   postgres.Config
	Clickhouse clickhouse.Config
	Mongo      mongo.Config
}
