package redis

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/DKeshavarz/sinar/internal/usecase"
	"github.com/redis/go-redis/v9"
)

type RedisOTPStore struct {
	client *redis.Client
}

func New() usecase.OtpStore {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})

	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		fmt.Printf("failed to connect to Redis: %s", err)
	}

	return &RedisOTPStore{
		client: client,
	}
}

func (r *RedisOTPStore) Create(userID, otp string, ttl time.Duration) error {
	ctx := context.Background()
	
	if err := r.client.Set(ctx, userID, otp, ttl).Err(); err != nil {
		log.Println("error in create redis", err)
		return err
	}

	return  nil
}
func (r *RedisOTPStore) Get(userID string) (string, error) {
	ctx := context.Background()
	otp, err := r.client.Get(ctx, userID).Result()
	if err == redis.Nil {
		return "", errors.New("OTP not found")
	}
	if err != nil {
		return "", fmt.Errorf("failed to retrieve OTP: %w", err)
	}
	return otp, nil
}
func (r *RedisOTPStore) Delete(userID string) error {
	ctx := context.Background()
	if err := r.client.Del(ctx, userID).Err(); err != nil {
		log.Println("error in deleting redis", err)
		return err
	}
	return nil
}

func (r *RedisOTPStore) Close() error {
	return r.client.Close()
}