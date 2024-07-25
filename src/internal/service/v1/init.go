package service_v1

import (
	"context"
	"fmt"
	"time"

	"github.com/teyz/go-svc-template/internal/database"
	pkg_cache "github.com/teyz/go-svc-template/pkg/cache"
)

const (
	exampleCacheDuration = time.Hour * 24
)

func generateExampleCacheKeyWithID(id string) string {
	return fmt.Sprintf("go-svc-template:channel:id:%v", id)
}

func generateExamplesCacheKey() string {
	return "go-svc-template:examples"
}

type service struct {
	store database.Database
	cache pkg_cache.Cache
}

func NewExampleStoreService(ctx context.Context, store database.Database, cache pkg_cache.Cache) (*service, error) {
	return &service{
		store: store,
		cache: cache,
	}, nil
}
