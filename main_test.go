package main

import "testing"

func SimpleTest(t *testing.T) {
	if 2*2 != 4 {
		t.Error("not passed")
	}
}
