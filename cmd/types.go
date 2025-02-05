package main

import (
	"github.com/go-redis/redis"
	"time"
)

type RedisClientInterface interface {
	Ping() *redis.StatusCmd
	Get(string) *redis.StringCmd
	Set(string, interface{}, time.Duration) *redis.StatusCmd
}
