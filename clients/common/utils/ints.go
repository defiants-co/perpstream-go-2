package utils

import (
	"math/big"
)

// maxBigint returns the maximum value for an int.
func MaxBigint() *big.Int {
	maxInt := big.NewInt(1)
	maxInt.Lsh(maxInt, 63).Sub(maxInt, big.NewInt(1)) // 2^63 - 1
	return maxInt
}
