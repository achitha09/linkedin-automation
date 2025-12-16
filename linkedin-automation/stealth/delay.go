package stealth

import (
	"math/rand"
	"time"
)

func RandomDelay(minMs int, maxMs int) {
	rand.Seed(time.Now().UnixNano())
	delay := rand.Intn(maxMs-minMs) + minMs
	time.Sleep(time.Duration(delay) * time.Millisecond)
}
