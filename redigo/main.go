package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}

	value, err := c.Do("set", "name", "lily")
	if err != nil {
		panic(err)
	}
	fmt.Println(value)
}
