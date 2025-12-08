package dmt

import (
	"math/big"

	"github.com/cockroachdb/apd/v3"
)

func New(coefficient int64, exponent int32) *apd.Decimal {
	return apd.New(coefficient, exponent)
}

// FromString call apd.NewFromString but ignore returned Condition
func FromString(s string) (*apd.Decimal, error) {
	d, _, err := apd.NewFromString(s)
	return d, err
}

// MustFromString panic instead of error.
func MustFromString(s string) *apd.Decimal {
	d, err := FromString(s)
	if err != nil {
		panic(err)
	}
	return d
}

// FromFloat call apd.NewFromFloat with new(Decimal)
func FromFloat(f float64) (*apd.Decimal, error) {
	return new(Decimal).SetFloat64(f)
}

// MustFromFloat panic instead of error.
func MustFromFloat(f float64) *apd.Decimal {
	ret, err := new(Decimal).SetFloat64(f)
	if err != nil {
		panic(err)
	}

	return ret
}

func FromInt[T int | int8 | int16 | int32 | int64](n T) *Decimal {
	return New(int64(n), 0)
}

func FromBigInt(coefficient *apd.BigInt, exponent int32) *apd.Decimal {
	return apd.NewWithBigInt(coefficient, exponent)
}

func FromMathBigInt(coeff *big.Int, exponent int32) *apd.Decimal {
	return apd.NewWithBigInt(new(apd.BigInt).SetMathBigInt(coeff), exponent)
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
