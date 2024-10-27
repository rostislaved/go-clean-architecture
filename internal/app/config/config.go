package config

import "github.com/rostislaved/go-clean-architecture/internal/app/application/usecases"

type Config struct {
	Info        Info
	Application Application
	Adapters    Adapters
}
type Info struct {
	Name    string
	Version string
}

type Application struct {
	UseCases usecases.Config
}
