package dmt

import (
	"github.com/cockroachdb/apd/v3"
)

func With(c *Context) ErrDecimal {
	return apd.MakeErrDecimal(c)
}

func New(coefficient int64, exponent int32) *apd.Decimal {
	return apd.New(coefficient, exponent)
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

func FromInt[T int | int8 | int16 | int32 | int64](n T) *Decimal {
	return New(int64(n), 0)
}

func FromBigInt(coefficient *apd.BigInt, exponent int32) *apd.Decimal {
	return apd.NewWithBigInt(coefficient, exponent)
}

func FromUint[T uint | uint8 | uint16 | uint32 | uint64](n T) *Decimal {
	return FromBigInt(new(apd.BigInt).SetUint64(uint64(n)), 0)
}

func Zero() *Decimal {
	return New(0, 0)
}

func One() *Decimal {
	return New(1, 0)
}
