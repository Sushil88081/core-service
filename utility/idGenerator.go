package utility

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateId() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	min := int64(1)
	max := int64(999999999999)
	id := r.Int63n(max-min+1) + min
	idStr := strconv.FormatInt(id, 10)
	return idStr
}
