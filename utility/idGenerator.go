package utility

import (
	"math/rand"
	"time"
)

func GenerateId() int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	min := int64(1)
	max := int64(999999999999)
	id := r.Int63n(max-min+1) + min
	return id
}
