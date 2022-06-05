package domain

import "github.com/go-redis/redis/v7"

//RedisClient Redis Client
type RedisClient interface {
	GetClient() *redis.Client
}
