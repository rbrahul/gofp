package gofp

import (
	"math/rand"
	"time"
)

// get
// set
// equal
// deepEual

func randomer() *rand.Rand {
	seed := rand.NewSource(time.Now().UnixNano())
	return rand.New(seed)
}
