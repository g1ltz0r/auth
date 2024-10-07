package helpers

import (
	"crypto/rand"
	"math"
	"math/big"
)

// GetRandID returns a uniform random value in [0, MaxInt64)
func GetRandID() int64 {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	return nBig.Int64()
}
