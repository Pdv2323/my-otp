package generate

import (
	"math/rand"
	"time"
)

func GenerateOtp() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(900000) + 100000
}
