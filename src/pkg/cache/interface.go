package pkg_cache

import (
	"context"
	"time"
)

type Cache interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (string, error)
	TTL(ctx context.Context, key string) (time.Duration, error)
	SetEx(ctx context.Context, key string, value interface{}, duration time.Duration) error
	Del(ctx context.Context, key string) error
	DelAll(ctx context.Context, keys ...string) error
	Incr(ctx context.Context, key string) error
	Decr(ctx context.Context, key string) error
	SAdd(ctx context.Context, key string, value interface{}) (int64, error)
	SAddAll(ctx context.Context, key string, values ...interface{}) (int64, error)
	SRem(ctx context.Context, key string, value interface{}) (int64, error)
	SCard(ctx context.Context, key string) (int64, error)
	SMembers(ctx context.Context, key string) ([]string, error)
	SIsMember(ctx context.Context, key string, value interface{}) (bool, error)
	ExpiresAt(ctx context.Context, key string, tm time.Time) error
	LPush(ctx context.Context, key string, value interface{}) error
	LPushAll(ctx context.Context, key string, values ...interface{}) (int64, error)
	LTrim(ctx context.Context, key string, start, stop int64) error
	LRange(ctx context.Context, key string) ([]string, error)
	LLen(ctx context.Context, key string) (int64, error)
	LPop(ctx context.Context, key string) (string, error)
	RPop(ctx context.Context, key string) (string, error)
	Expire(ctx context.Context, key string, duration time.Duration) error
	ZAdd(ctx context.Context, key string, value interface{}) error
	ZAddWithScore(ctx context.Context, key string, score float64, value interface{}) error
	ZRem(ctx context.Context, key string, value interface{}) (int64, error)
	ZPopMin(ctx context.Context, key string, nb int64) ([]string, error)
	ZCount(ctx context.Context, key string) (int64, error)
	ZRange(ctx context.Context, key string) ([]string, error)
	HSet(ctx context.Context, key, field string, value interface{}) error
	HGet(ctx context.Context, key, field string) (string, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	HIncrBy(ctx context.Context, key, field string, incr int64) error
}
