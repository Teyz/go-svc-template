package pkg_redis

import (
	"context"

	"github.com/redis/go-redis/v9"

	pkg_cache "github.com/teyz/go-svc-template/pkg/cache"
)

type cacheClient struct {
	rdb redis.UniversalClient
}

func NewRedisCache(ctx context.Context, rdb redis.UniversalClient) pkg_cache.Cache {
	return &cacheClient{rdb: rdb}
}
