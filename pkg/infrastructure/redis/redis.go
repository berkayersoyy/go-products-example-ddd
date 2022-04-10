package redis

import (
	"context"
	"github.com/berkayersoyy/go-products-example-ddd/pkg/domain"
	"github.com/go-redis/redis/v7"
	"github.com/sethvargo/go-retry"
	"log"
	"os"
	"time"
)

type redisClient struct {
	SingletonRedis *redis.Client
}

func ProvideRedisClient() domain.RedisClient {
	return &redisClient{SingletonRedis: InitRedis()}
}
func (r *redisClient) GetClient() *redis.Client {
	return r.SingletonRedis
}

func InitRedis() *redis.Client {
	//conf, err := config.LoadConfig("./")
	//if err != nil {
	//	panic(err)
	//}
	DSN := os.Getenv("REDIS_HOST")
	client := redis.NewClient(&redis.Options{
		Addr: DSN,
	})
	ctx := context.Background()
	if err := retry.Fibonacci(ctx, 1*time.Second, func(ctx context.Context) error {
		if _, err := client.Ping().Result(); err != nil {
			return retry.RetryableError(err)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return client
}
