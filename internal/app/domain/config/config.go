package config

import (
	"time"
)

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
	UpdateService UpdateService
}

type UpdateService struct {
	UpdateInterval time.Duration
}
