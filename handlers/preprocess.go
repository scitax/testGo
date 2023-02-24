package handlers

import (
	"time"
)

func Preproces() error {
	// Simulate a long preprocessing task
	time.Sleep(10 * time.Second)
	PreprocesStatus = true

	return nil
}
