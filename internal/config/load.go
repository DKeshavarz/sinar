package config

import "github.com/DKeshavarz/sinar/internal/interface/redis"

func New() *Config{

	r := redis.Config {
		Addr: GetEnv("REDIS_ADDR", "localhost:6379"),
		Password: GetEnv("REDIS_PASS", ""),
		DB: GetEnvAsInt("REDIS_DB", 0),
	}

	conf := Config{
		Redis: &r,
	}

	return &conf
}