package redis

import (
	"context"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
	"github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-retry"
	"log"
	"os"
	"time"
)

//redisClient Redis client
type redisClient struct {
	SingletonRedis *redis.Client
}

//ProvideRedisClient Provide redis client
func ProvideRedisClient() domain.RedisClient {
	return &redisClient{SingletonRedis: InitRedis()}
}
func (r *redisClient) GetClient() *redis.Client {
	return r.SingletonRedis
}

//InitRedis Init redis
func InitRedis() *redis.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dsn := os.Getenv("REDIS_HOST")
	var client *redis.Client
	ctx := context.Background()
	if err := retry.Fibonacci(ctx, 1*time.Second, func(ctx context.Context) error {
		client = redis.NewClient(&redis.Options{
			Addr: dsn,
		})
		if _, err := client.Ping().Result(); err != nil {
			return retry.RetryableError(err)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return client
}
