package models

import (
	"fmt"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
)

func CreateConnectionPool() *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			// Connect to Redis.
			conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")))
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
		MaxIdle:     10,
		MaxActive:   100,
		IdleTimeout: time.Minute,
		Wait:        true,
	}
}

func IndexVm() error {
	return nil
}
