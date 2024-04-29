package bigint

import (
	"math/big"
)

// Int
type Int = big.Int

func NewFromString(s string) (*Int, bool) {
	return new(big.Int).SetString(s, 10)
}
