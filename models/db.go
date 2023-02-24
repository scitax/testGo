package models

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

var client *redis.Client

func Init() {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "",
		DB:       0,
	})

}
