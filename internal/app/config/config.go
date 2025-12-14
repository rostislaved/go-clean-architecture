package config

import (
	"github.com/rostislaved/go-clean-architecture/internal/app/application/usecases"
	"github.com/rostislaved/go-clean-architecture/internal/pkg/clickhouse"
	"github.com/rostislaved/go-clean-architecture/internal/pkg/mongo"
	"github.com/rostislaved/go-clean-architecture/internal/pkg/postgres"
)

type Config struct {
	Application    Application
	Adapters       Adapters
	Infrastructure Infrastructure
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
