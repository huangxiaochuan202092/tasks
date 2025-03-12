package utils

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("failed to connect redis: %v", err)
	}
	log.Println("redis connected success")
}

// RedisClient 初始化 Redis 客户端

// SetVerificationCode 设置验证码到 Redis
func SetVerificationCode(email, code string) error {
	ctx := context.Background()
	return RedisClient.Set(ctx, email, code, 5*time.Minute).Err()
}

// GetVerificationCode 从 Redis 获取验证码
func GetVerificationCode(email string) (string, error) {
	ctx := context.Background()
	return RedisClient.Get(ctx, email).Result()
}

// DelVerificationCode 删除验证码 redis
func DelVerificationCode(email string) error {
	ctx := context.Background()
	return RedisClient.Del(ctx, email).Err()
}
