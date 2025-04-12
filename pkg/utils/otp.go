package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateOTPWithLocalRand возвращает 4-значный OTP, используя локальный генератор
func GenerateOTPWithLocalRand() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := r.Intn(10000)
	return fmt.Sprintf("%04d", otp)
}
