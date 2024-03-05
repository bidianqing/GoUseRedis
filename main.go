package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "name", "tom", time.Second*60).Err()
	if err != nil {
		fmt.Print(err)
	}

	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("name", val)

	age, err := rdb.Incr(ctx, "age").Result()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("name", age)
}
