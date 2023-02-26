package main

import (
	"testing"

	"github.com/scitax/testGo/models/db_tests"
)

func ConnectioPoolTesting(t *testing.T) {
	if err := db_tests.PoolTest(); err != nil {
		t.Error("Connection pool test didn't passed")
	}
}

func RedisSearch(t *testing.T) {
	if err := db_tests.SearchIndexTest(); err != nil {
		t.Error("Redis search test didn't passed")
	}
}
