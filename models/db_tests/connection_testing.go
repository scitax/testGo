package db_tests

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/scitax/testGo/models"
)

func PoolTest() error {
	// Create a connection pool.
	pool := models.CreateConnectionPool()

	// Get a connection from the pool.
	conn := pool.Get()
	defer conn.Close()

	// Use the connection to perform Redis operations.
	_, err := conn.Do("SET", "foo", "bar")
	if err != nil {
		return err
	}
	val, err := redis.String(conn.Do("GET", "foo"))
	if err != nil {
		return err
	}
	fmt.Println("foo:", val)

	_, err = conn.Do("DEL", "foo")
	if err != nil {
		return err
	}
	return nil
}
