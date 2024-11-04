package helpers

import (
	"math/rand"
	"time"
)

func RandomID() int64 {
	src := rand.NewSource(time.Now().UnixNano())
	random := rand.New(src)
	return random.Int63()
}
