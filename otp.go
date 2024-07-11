package otp

import (
	"time"

	"math/rand"
)

func VerifyOtp(otp, newOtp int) bool {
	if otp != newOtp {
		return false
	}

	return true
}

func GenerateOtp(n int) int {
	if n < 1 {
		return 0
	}

	min := 1
	for i := 1; i < n; i++ {
		min *= 10
	}
	max := min*10 - 1

	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
