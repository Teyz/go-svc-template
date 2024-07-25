package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/teyz/go-svc-template/internal/config"
	database_postgres "github.com/teyz/go-svc-template/internal/database/postgres"
	service_v1 "github.com/teyz/go-svc-template/internal/service/v1"
	pkg_redis "github.com/teyz/go-svc-template/pkg/cache/redis"
	pkg_config "github.com/teyz/go-svc-template/pkg/config"
	pkg_postgres "github.com/teyz/go-svc-template/pkg/database/postgres"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)

	cfg := &config.Config{}
	err := pkg_config.ParseConfig(cfg)
	if err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to parse config")
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	cacheConnection := pkg_redis.GetConnection(ctx, &cfg.RedisConfig)
	cacheRedis := pkg_redis.NewRedisCache(ctx, cacheConnection)

	databaseConnection, err := pkg_postgres.NewDatabaseConnection(ctx, &cfg.PostgresConfig)
	if err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to create database connection")
	}
	databaseClient := database_postgres.NewClient(ctx, databaseConnection)

	exampleStoreService, err := service_v1.NewExampleStoreService(ctx, databaseClient, cacheRedis)
	if err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to create example store service")
	}

	// create http server
	httpServer, err := handlers_http.NewServer(ctx, cfg.HTTPServerConfig, service)
	if err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to create http server")
	}

	// setup http server
	if err := httpServer.Setup(ctx); err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to setup http server")
	}

	// start http server
	if err := httpServer.Start(ctx); err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to start http server")
	}

	<-sigs
	cancel()

	// stop http server
	if err := httpServer.Stop(ctx); err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to stop http server")
	}

	os.Exit(0)
}
