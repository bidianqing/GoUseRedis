package main

import (
	"context"
	"fmt"
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

	Set("name", "bidianqing", time.Minute*2)

	name := Get("name")
	fmt.Println(name)
}

func Set(key string, value string, duration time.Duration) {
	err = client.Do(ctx, client.B().Set().Key(key).Value(value).Ex(duration).Build()).Error()
	if err != nil {
		panic(err)
	}
}

func Get(key string) string {
	name, err := client.Do(ctx, client.B().Get().Key(key).Build()).ToString()
	if err != nil {
		panic(err)
	}

	return name
}