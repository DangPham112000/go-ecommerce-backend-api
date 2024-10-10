package random

import (
	"crypto/rand"
	"math/big"
)

func GenerateSixDigitOTP() int {
	// 6-digit OTP means values between 100000 and 999999
	min := 100000
	max := 999999
	rangeVal := big.NewInt(int64(max - min + 1))

	// Generate a secure random number
	num, _ := rand.Int(rand.Reader, rangeVal)

	// Add the offset to get a 6-digit number
	otp := int(num.Int64()) + min
	return otp
}
