package dmt

import (
	"github.com/cockroachdb/apd/v3"
)

// Decimal
type Decimal = apd.Decimal

// Context
type Context = apd.Context

// Condition
type Condition = apd.Condition

// ErrDecimal
type ErrDecimal = apd.ErrDecimal

// Sum is for convenience to call SumTo
func Sum(values ...*Decimal) (*Decimal, error) {
	dst := new(Decimal)
	err := SumTo(dst, values...)
	return dst, err
}

// SumTo add all values to dst
func SumTo(dst *Decimal, values ...*Decimal) error {
	dst.SetInt64(0)
	if len(values) == 0 {
		return nil
	}
	for _, value := range values {
		err := AddTo(dst, value, dst)
		if err != nil {
			return err
		}
	}
	return nil
}

// Product is for convenience to call ProductTo
func Product(values ...*Decimal) (*Decimal, error) {
	dst := new(Decimal)
	err := ProductTo(dst, values...)
	return dst, err
}

// ProductTo multiply all values to dst
func ProductTo(dst *Decimal, values ...*Decimal) error {
	if len(values) == 0 {
		dst.SetInt64(0)
		return nil
	}
	dst.SetInt64(1)
	for _, value := range values {
		err := MulTo(dst, value, dst)
		if err != nil {
			return err
		}
	}
	return nil
}

// SafePresision get a safe presision which have enough space for x
//   - SafePresision take x.Exponent into account even if x == 0
func SafePresision(x *Decimal) int64 {
	s := x.NumDigits()
	if x.Exponent > 0 {
		// 1234000
		s += int64(x.Exponent)
	} else if e := -int64(x.Exponent); e > s {
		// 0.000000123
		s = e
	}
	// 123.4567 no risk

	// result
	return s
}
