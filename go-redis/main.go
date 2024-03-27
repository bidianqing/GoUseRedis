package main

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var client *redis.Client

func main() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()
}

func Set(key string, value string, expiration time.Duration) {
	client.Set(ctx, key, value, expiration)
}

func Get(key string) string {
	value, _ := client.Get(ctx, key).Result()
	return value
}

func Expire(key string, expiration time.Duration) {
	client.Expire(ctx, key, expiration)
}

func Incr(key string) int64 {
	value, _ := client.Incr(ctx, key).Result()
	return value
}

func Del(key ...string) {
	client.Del(ctx, key...)
}
