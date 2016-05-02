package main

import (
	"math/rand"
	"time"
)

// BuildAddress merges the host and
// port of a remote host, and returns
// the combined string.
func BuildAddress(host, port string) string {
	return host + ":" + port
}

// RandomInRange returns a random number within
// min and max provided
func RandomInRange(min, max int64) int64 {
	rand.Seed(time.Now().Unix())
	return rand.Int63n(max-min) + min
}
