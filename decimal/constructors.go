package decimal

import (
	"math/big"

	"github.com/cockroachdb/apd/v3"
)

func New(coeff int64, exponent int32) *apd.Decimal {
	return apd.New(coeff, exponent)
}

func NewFromInt(coeff *big.Int, exponent int32) *apd.Decimal {
	return apd.NewWithBigInt(new(apd.BigInt).SetMathBigInt(coeff), exponent)
}

func NewFromBigInt(coeff *apd.BigInt, exponent int32) *apd.Decimal {
	return apd.NewWithBigInt(new(apd.BigInt).Set(coeff), exponent)
}

// NewFromString call apd.NewFromString but ignore returned Condition
func NewFromString(s string) (*apd.Decimal, error) {
	d, _, err := apd.NewFromString(s)
	return d, err
}

// NewFromFloat call apd.NewFromFloat with new(Decimal)
func NewFromFloat(f float64) (*apd.Decimal, error) {
	return new(Decimal).SetFloat64(f)
}

func ToDecimal[T int | int8 | int16 | int32 | int64](i T) *Decimal {
	return New(int64(i), 0)
}
