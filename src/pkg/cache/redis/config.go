package pkg_redis

type RedisConfig struct {
	CacheHost string `env:"CACHE_HOST"`
	CachePort uint16 `env:"CACHE_PORT"`
}
