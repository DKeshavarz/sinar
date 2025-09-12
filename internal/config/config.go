package config

import "github.com/DKeshavarz/sinar/internal/interface/redis"

type Config struct{
	Redis *redis.Config
}

func Default() Config{
	return Config{}
}