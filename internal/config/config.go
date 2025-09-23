package config

import (
	"github.com/DKeshavarz/sinar/internal/interface/redis"
	"github.com/DKeshavarz/sinar/pkg/sms"
)

type Config struct {
	Redis *redis.Config
	SMS   *sms.Config
}

func Default() Config {
	return Config{}
}
