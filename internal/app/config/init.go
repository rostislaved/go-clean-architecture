package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/rostislaved/cleanconfig"
)

func New() (config Config, err error) {
	config, err = cleanconfig.NewReader[Config]().
		Add(yaml.Parser(), file.Provider("config.yaml")).
		Read()
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
