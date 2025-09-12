package config

import (
	"github.com/DKeshavarz/sinar/internal/interface/redis"
	"github.com/DKeshavarz/sinar/internal/interface/sms"
)

func New() *Config{

	r := redis.Config {
		Addr: GetEnv("REDIS_ADDR", "localhost:6379"),
		Password: GetEnv("REDIS_PASS", ""),
		DB: GetEnvAsInt("REDIS_DB", 0),
	}

	s := sms.Config {
		ApiKey: GetEnv("SMS_APIKEY", ""),
		Sender: GetEnv("SMS_SENDER", ""),
	}
	conf := Config{
		Redis: &r,
		SMS: &s,
	}

	return &conf
}