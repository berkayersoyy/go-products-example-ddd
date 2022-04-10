package domain

import "github.com/go-redis/redis/v7"

type RedisClient interface {
	GetClient() *redis.Client
}
