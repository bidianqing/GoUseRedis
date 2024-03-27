package main

import (
	"context"
	"fmt"
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

	val := HGetAll("tom")
	fmt.Println(val)
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

func HSet(key string, fieldValues map[string]string) {
	client.HSet(ctx, key, fieldValues)
}

func HGetAll(key string) map[string]string {
	val, _ := client.HGetAll(ctx, key).Result()
	return val
}
