package config

import (
	"github.com/DKeshavarz/sinar/internal/interface/redis"
	"github.com/DKeshavarz/sinar/internal/interface/sms"
)

type Config struct {
	Redis *redis.Config
	SMS   *sms.Config
}

func Default() Config {
	return Config{}
}
