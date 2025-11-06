package bigint

import (
	"math/big"

	"github.com/cockroachdb/apd/v3"
)

// Int
type Int = big.Int

func FromString(s string) (*Int, bool) {
	return new(big.Int).SetString(s, 10)
}

func FromBigInt(s *apd.BigInt) *Int {
	return s.MathBigInt()
}
