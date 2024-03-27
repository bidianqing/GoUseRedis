package main

import (
	"context"
	"time"

	"github.com/redis/rueidis"
)

var client rueidis.Client
var err error
var ctx = context.Background()

func main() {
	client, err = rueidis.NewClient(rueidis.ClientOption{
		InitAddress:  []string{"127.0.0.1:6379"},
		DisableCache: true,
	})
	if err != nil {
		panic(err)
	}
	defer client.Close()
}

func Set(key string, value string, expiration time.Duration) {
	if expiration > 0 {
		client.Do(ctx, client.B().Set().Key(key).Value(value).Ex(expiration).Build())
	} else {
		client.Do(ctx, client.B().Set().Key(key).Value(value).Build())
	}
}

func Get(key string) string {
	name, err := client.Do(ctx, client.B().Get().Key(key).Build()).ToString()
	if err != nil {
		panic(err)
	}

	return name
}

func Expire(key string, expiration time.Duration) {
	client.Do(ctx, client.B().Expire().Key(key).Seconds(int64(expiration.Seconds())).Build())
}

func Incr(key string) int64 {
	id, _ := client.Do(ctx, client.B().Incr().Key(key).Build()).ToInt64()

	return id
}

func Del(key ...string) {
	client.Do(ctx, client.B().Del().Key(key...).Build())
}
