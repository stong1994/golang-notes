package main

import (
	"testing"
	"time"
)

func TestElapsed(t *testing.T) {
	defer Elapsed("TestElapsed")()
	time.Sleep(3 * time.Second)
}
