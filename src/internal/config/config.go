package config

import (
	pkg_redis "github.com/teyz/go-svc-template/pkg/cache/redis"
	pkg_config "github.com/teyz/go-svc-template/pkg/config"
	pkg_postgres "github.com/teyz/go-svc-template/pkg/database/postgres"
	pkg_http "github.com/teyz/go-svc-template/pkg/http"
)

type Config struct {
	ServiceConfig pkg_config.ServiceConfig

	HTTPServerConfig pkg_http.HTTPServerConfig
	PostgresConfig   pkg_postgres.PostgresConfig
	RedisConfig      pkg_redis.RedisConfig
}
