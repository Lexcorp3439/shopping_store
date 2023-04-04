package testutils

import (
	"crypto/rand"
	"math/big"
)

func GetRandomInt(max int64) int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(max))
	return n.Int64()
}
