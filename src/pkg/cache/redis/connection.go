package pkg_redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func GetConnection(ctx context.Context, cfg *RedisConfig) redis.UniversalClient {
	return redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{fmt.Sprintf("%s:%d", cfg.CacheHost, cfg.CachePort)},
		DB:    0,
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			log.Info().Msg("connected to the cache")
			return nil
		},
		ContextTimeoutEnabled: true,
		MinIdleConns:          10,
	})
}
