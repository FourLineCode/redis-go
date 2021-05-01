package main

import (
	"github.com/FourLineCode/redis-go/pkg/redis"
)

func main() {
	r := redis.NewClient()

	r.Run()
}
