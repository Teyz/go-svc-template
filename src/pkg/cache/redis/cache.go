package pkg_redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"

	"github.com/teyz/go-svc-template/pkg/errors"
)

func (c *cacheClient) Set(ctx context.Context, key string, value interface{}) error {
	err := c.rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to set key in the cache")
		return err
	}

	return nil
}

func (c *cacheClient) SetEx(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	err := c.rdb.SetEx(ctx, key, value, duration).Err()
	if err != nil {
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to set key in the cache")
		return err
	}

	return nil
}

func (c *cacheClient) TTL(ctx context.Context, key string) (time.Duration, error) {
	expiresAt, err := c.rdb.TTL(ctx, key).Result()
	if err != nil {
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to ttl key in the cache")
		return expiresAt, err
	}

	return expiresAt, nil
}

func (c *cacheClient) Get(ctx context.Context, key string) (string, error) {
	val, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		if err.Error() == redis.Nil.Error() {
			return "", errors.NewNotFoundError("key not found")
		}
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to get key from the cache")

		return "", err
	}

	return val, nil
}

func (c *cacheClient) Del(ctx context.Context, key string) error {
	_, err := c.rdb.Del(ctx, key).Result()
	if err != nil {
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to del key in the cache")
		return err
	}

	return nil
}

func (c *cacheClient) DelAll(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil // No keys to delete
	}

	_, err := c.rdb.Del(ctx, keys...).Result()
	if err != nil {
		log.Error().Err(err).
			Msg("unable to DelAll keys in the cache")
		return err
	}

	return nil
}

func (c *cacheClient) Incr(ctx context.Context, key string) error {
	err := c.rdb.Incr(ctx, key).Err()
	if err != nil {
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to incr key in the cache")
		return err
	}

	return nil
}

func (c *cacheClient) Decr(ctx context.Context, key string) error {
	err := c.rdb.Decr(ctx, key).Err()
	if err != nil {
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to decr key in the cache")
		return err
	}

	return nil
}

func (c *cacheClient) ExpiresAt(ctx context.Context, key string, tm time.Time) error {
	err := c.rdb.ExpireAt(ctx, key, tm).Err()
	if err != nil {
		log.
			Error().
			Err(err).
			Str("key", key).
			Msg("unable to expire key in the cache")
		return err
	}

	return nil
}

func (c *cacheClient) HSet(ctx context.Context, key, field string, value interface{}) error {
	err := c.rdb.HSet(ctx, key, field, value).Err()
	if err != nil {
		log.Error().Err(err).
			Str("key", key).
			Str("field", field).
			Msg("unable to set field in the hash")
		return err
	}

	return nil
}

func (c *cacheClient) HGet(ctx context.Context, key, field string) (string, error) {
	result, err := c.rdb.HGet(ctx, key, field).Result()
	if err != nil {
		log.Error().Err(err).
			Str("key", key).
			Str("field", field).
			Msg("unable to get field from the hash")
		return "", err
	}

	return result, nil
}

func (c *cacheClient) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	result, err := c.rdb.HGetAll(ctx, key).Result()
	if err != nil {
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to get all fields from the hash")
		return nil, err
	}

	return result, nil
}

func (c *cacheClient) HIncrBy(ctx context.Context, key, field string, incr int64) error {
	_, err := c.rdb.HIncrBy(ctx, key, field, incr).Result()
	if err != nil {
		log.Error().Err(err).
			Str("key", key).
			Str("field", field).
			Msg("unable to increment field in hash in the cache")
		return err
	}

	return nil
}

func (c *cacheClient) LPush(ctx context.Context, key string, value interface{}) error {
	err := c.rdb.LPush(ctx, key, value).Err()
	if err != nil {
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to lpush key in the cache")
		return err
	}

	return nil
}

func (c *cacheClient) LPushAll(ctx context.Context, key string, values ...interface{}) (int64, error) {
	val, err := c.rdb.LPush(ctx, key, values...).Result()
	if err != nil {
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to LPushAll key in the cache")
		return 0, err
	}

	return val, nil
}

func (c *cacheClient) LTrim(ctx context.Context, key string, start, stop int64) error {
	err := c.rdb.LTrim(ctx, key, start, stop).Err()
	if err != nil {
		if err.Error() == redis.Nil.Error() {
			return errors.NewNotFoundError("key not found")

		}
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to ltrim key in the cache")
		return err
	}

	return nil
}

func (c *cacheClient) LLen(ctx context.Context, key string) (int64, error) {
	val, err := c.rdb.LLen(ctx, key).Result()
	if err != nil {
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to llen")

		return 0, err
	}

	return val, nil
}

func (c *cacheClient) LPop(ctx context.Context, key string) (string, error) {
	val, err := c.rdb.LPop(ctx, key).Result()
	if err != nil {
		if err.Error() == redis.Nil.Error() {
			return "", errors.NewNotFoundError("key not found")
		}
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to lpop key in the cache")
		return "", err
	}

	return val, nil
}

func (c *cacheClient) RPop(ctx context.Context, key string) (string, error) {
	val, err := c.rdb.RPop(ctx, key).Result()
	if err != nil {
		if err.Error() == redis.Nil.Error() {
			return "", errors.NewNotFoundError("key not found")
		}
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to rpop key in the cache")
		return "", err
	}

	return val, nil
}

func (c *cacheClient) Expire(ctx context.Context, key string, duration time.Duration) error {
	err := c.rdb.Expire(ctx, key, duration).Err()
	if err != nil {
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to expire key in the cache")
		return err
	}

	return nil
}

func (c *cacheClient) ZAdd(ctx context.Context, key string, value interface{}) error {
	err := c.rdb.ZAdd(ctx, key, redis.Z{
		Score:  float64(time.Now().UnixMilli()),
		Member: value,
	}).Err()
	if err != nil {
		if err.Error() == redis.Nil.Error() {
			return errors.NewNotFoundError("key not found")
		}
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to zadd key in the cache")
		return err
	}

	return nil
}

func (c *cacheClient) ZAddWithScore(ctx context.Context, key string, score float64, value interface{}) error {
	err := c.rdb.ZAdd(ctx, key, redis.Z{
		Score:  score,
		Member: value,
	}).Err()
	if err != nil {
		if err.Error() == redis.Nil.Error() {
			return errors.NewNotFoundError("key not found")
		}
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to ZAddWithScore key in the cache")
		return err
	}

	return nil
}

func (c *cacheClient) ZRem(ctx context.Context, key string, value interface{}) (int64, error) {
	val, err := c.rdb.ZRem(ctx, key, value).Result()
	if err != nil {
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to ZRem key in the cache")
		return 0, err
	}

	return val, nil
}

func (c *cacheClient) ZPopMin(ctx context.Context, key string, nb int64) ([]string, error) {
	val, err := c.rdb.ZPopMin(ctx, key, nb).Result()
	if err != nil {
		if err.Error() == redis.Nil.Error() {
			return nil, errors.NewNotFoundError("key not found")
		}
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to zpopmin key in the cache")
		return nil, err
	}

	var members []string
	for _, member := range val {
		members = append(members, member.Member.(string))
	}

	return members, nil
}

func (c *cacheClient) ZCount(ctx context.Context, key string) (int64, error) {
	val, err := c.rdb.ZCount(ctx, key, "-inf", "+inf").Result()
	if err != nil {
		if err.Error() == redis.Nil.Error() {
			return 0, errors.NewNotFoundError("key not found")
		}
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to zcount key in the cache")
		return 0, err
	}

	return val, nil
}

func (c *cacheClient) ZRange(ctx context.Context, key string) ([]string, error) {
	val, err := c.rdb.ZRange(ctx, key, 0, -1).Result()
	if err != nil {
		if err.Error() == redis.Nil.Error() {
			return nil, errors.NewNotFoundError("key not found")
		}
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to zrange key in the cache")
		return nil, err
	}

	return val, nil
}

func (c *cacheClient) LRange(ctx context.Context, key string) ([]string, error) {
	val, err := c.rdb.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		if err.Error() == redis.Nil.Error() {
			return nil, errors.NewNotFoundError("key not found")
		}
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to lrange key in the cache")
		return nil, err
	}

	return val, nil
}

func (c *cacheClient) SAdd(ctx context.Context, key string, value interface{}) (int64, error) {
	val, err := c.rdb.SAdd(ctx, key, value).Result()
	if err != nil {
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to sadd key in the cache")
		return 0, err
	}

	return val, nil
}

func (c *cacheClient) SAddAll(ctx context.Context, key string, values ...interface{}) (int64, error) {
	val, err := c.rdb.SAdd(ctx, key, values...).Result()
	if err != nil {
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to saddAll key in the cache")
		return 0, err
	}

	return val, nil
}

func (c *cacheClient) SRem(ctx context.Context, key string, value interface{}) (int64, error) {

	val, err := c.rdb.SRem(ctx, key, value).Result()
	if err != nil {
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to SRem key in the cache")
		return 0, err
	}

	return val, nil
}

func (c *cacheClient) SCard(ctx context.Context, key string) (int64, error) {
	val, err := c.rdb.SCard(ctx, key).Result()
	if err != nil {
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to scard key in the cache")
		return 0, err
	}

	return val, nil
}

func (c *cacheClient) SIsMember(ctx context.Context, key string, value interface{}) (bool, error) {
	val, err := c.rdb.SIsMember(ctx, key, value).Result()
	if err != nil {
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to SIsMember key in the cache")
		return false, err
	}

	return val, nil
}

func (c *cacheClient) SMembers(ctx context.Context, key string) ([]string, error) {
	values, err := c.rdb.SMembers(ctx, key).Result()
	if err != nil {
		log.Error().Err(err).
			Str("key", key).
			Msg("unable to SMembers key in the cache")
		return nil, err
	}

	return values, nil
}
