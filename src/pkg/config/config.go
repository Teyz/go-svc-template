package pkg_config

import "github.com/caarlos0/env/v10"

type ServiceConfig struct {
	ServiceName string `env:"SERVICE_NAME"`
	Environment string `env:"ENVRIONMENT" envDefault:"local"`
}

func ParseConfig[T any](cfg *T) error {
	return env.Parse(cfg)
}
