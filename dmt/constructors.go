package dmt

import (
	"math/big"

	"github.com/cockroachdb/apd/v3"
)

func WithContext(c *Context) ErrDecimal {
	return apd.MakeErrDecimal(c)
}

func New(coeff int64, exponent int32) *apd.Decimal {
	return apd.New(coeff, exponent)
}

func FromInt(coeff *big.Int, exponent int32) *apd.Decimal {
	return FromBigInt(new(apd.BigInt).SetMathBigInt(coeff), exponent)
}

func FromBigInt(coeff *apd.BigInt, exponent int32) *apd.Decimal {
	return apd.NewWithBigInt(coeff, exponent)
}

// NewFromString call apd.NewFromString but ignore returned Condition
func FromString(s string) (*apd.Decimal, error) {
	d, _, err := apd.NewFromString(s)
	return d, err
}

// NewFromFloat call apd.NewFromFloat with new(Decimal)
func FromFloat(f float64) (*apd.Decimal, error) {
	return new(Decimal).SetFloat64(f)
}

func FromIntN[T int | int8 | int16 | int32 | int64](n T) *Decimal {
	return New(int64(n), 0)
}

func FromUintN[T uint | uint8 | uint16 | uint32 | uint64](n T) *Decimal {
	return FromInt(new(big.Int).SetUint64(uint64(n)), 0)
}
