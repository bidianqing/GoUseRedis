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

	Set("name", "bidianqing", time.Second*30)

	name := Get("name")
	fmt.Println(name)

	Hset("user:1", map[string]string{
		"name": "bidianqing",
		"age":  "22",
	})

	Expire("user:1", time.Second*30)
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

func Hset(key string, fieldValues map[string]string) {
	fieldValue := client.B().Hset().Key(key).FieldValue()
	for field, value := range fieldValues {
		fieldValue.FieldValue(field, value)
	}

	client.Do(ctx, fieldValue.Build())
}

func Expire(key string, duration time.Duration) {
	client.Do(ctx, client.B().Expire().Key(key).Seconds(int64(duration.Seconds())).Build())
}
