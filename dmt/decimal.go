package dmt

import (
	"errors"
	"fmt"
	"math"
	"math/big"

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

// Add is for convenience to call AddTo
func Add(x *Decimal, y *Decimal) (*Decimal, error) {
	return AddTo(new(Decimal), x, y)
}

// AddTo sets dst to the sum x+y and return dst.
// if dst is nil, new(Decimal) will be used
func AddTo(dst *Decimal, x *Decimal, y *Decimal) (*Decimal, error) {
	_, err := apd.BaseContext.Add(dst, x, y)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

// Sub is for convenience to call SubTo
func Sub(x *Decimal, y *Decimal) (*Decimal, error) {
	return SubTo(new(Decimal), x, y)
}

// SubTo sets dst to the difference x-y and return dst.
// if dst is nil, new(Decimal) will be used
func SubTo(dst *Decimal, x *Decimal, y *Decimal) (*Decimal, error) {
	_, err := apd.BaseContext.Sub(dst, x, y)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

// Mul is for convenience to call MulTo
func Mul(x *Decimal, y *Decimal) (*Decimal, error) {
	return MulTo(new(Decimal), x, y)
}

// MulTo sets dst to the product x*y and return dst.
// if dst is nil, new(Decimal) will be used
func MulTo(dst *Decimal, x *Decimal, y *Decimal) (*Decimal, error) {
	_, err := apd.BaseContext.Mul(dst, x, y)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

// Div is for convenience to call DivTo
func Div(c *Context, x *Decimal, y *Decimal) (*Decimal, error) {
	return DivTo(new(Decimal), c, x, y)
}

// DivTo sets dst to the quotient x/y for y != 0 and return dst.
// if dst is nil, new(Decimal) will be used
// c.Precision must be > 0.
// If an exact division is required, use a context with high precision
// and verify it was exact by checking the Inexact flag on the return Condition.
func DivTo(dst *Decimal, c *Context, x *Decimal, y *Decimal) (*Decimal, error) {
	_, err := c.Quo(dst, x, y)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

// SetScale is limit the precision and scale of dst
func SetScale(dst *Decimal, precision, scale int, rounder apd.Rounder) error {
	return setScaleTo(dst, precision, scale, rounder)
}

// SetScaleTo is similar to SetScale but avoid mutating the value of x
func SetScaleTo(dst *Decimal, x *Decimal, precision, scale int, rounder apd.Rounder) error {
	dst.Set(x)
	return setScaleTo(dst, precision, scale, rounder)
}

func setScaleTo(dst *Decimal, precision, scale int, rounder apd.Rounder) error {
	if dst.Form != apd.Finite || precision <= 0 {
		return nil
	}
	// Use +1 here because it is inverted later.
	if scale < math.MinInt32+1 || scale > math.MaxInt32 {
		return errors.New("scale out of range")
	}
	if scale > precision {
		return fmt.Errorf("scale (%d) must be between 0 and precision (%d)", scale, precision)
	}
	c := &apd.Context{
		Precision:   uint32(precision),
		MaxExponent: apd.MaxExponent,
		MinExponent: apd.MinExponent,
		Traps:       apd.InvalidOperation,
		Rounding:    rounder,
	}
	if _, err := c.Quantize(dst, dst, -int32(scale)); err != nil {
		var lt string
		switch v := precision - scale; v {
		case 0:
			lt = "1"
		default:
			lt = fmt.Sprintf("10^%d", v)
		}
		return fmt.Errorf("value with precision %d, scale %d must round to an absolute value less than %s", precision, scale, lt)
	}
	return nil
}

// Integral return integral part of x
func Integral(x *Decimal) *big.Int {
	return IntegralTo(new(apd.BigInt), x).MathBigInt()
}

// IntegralDst return set d to integral part of x.
func IntegralDst(dst, x *Decimal) (*Decimal, error) {
	_, err := RoundDownContext.RoundToIntegralValue(dst, x)
	return dst, err
}

// IntegralTo sets y by remove the fraction part of x.
// if y is nil, new(Decimal) will be used
func IntegralTo(dst *apd.BigInt, x *Decimal) *apd.BigInt {
	d := new(apd.Decimal)
	_, _ = IntegralDst(d, x)
	if d.Negative {
		dst.Set(dst.Neg(&d.Coeff))
	} else {
		dst.Set(&d.Coeff)
	}
	return dst
}

// Sum is for convenience to call SumTo
func Sum(values ...*Decimal) (*Decimal, error) {
	return SumTo(new(Decimal), values...)
}

// SumTo add all values to dst
func SumTo(dst *Decimal, values ...*Decimal) (*Decimal, error) {
	dst.SetInt64(0)
	if len(values) == 0 {
		return dst, nil
	}
	for _, value := range values {
		_, err := AddTo(dst, value, dst)
		if err != nil {
			return nil, err
		}
	}
	return dst, nil
}

// Product is for convenience to call ProductTo
func Product(values ...*Decimal) (*Decimal, error) {
	return ProductTo(new(Decimal), values...)
}

// ProductTo multiply all values to dst
func ProductTo(dst *Decimal, values ...*Decimal) (*Decimal, error) {
	if len(values) == 0 {
		return dst.SetInt64(0), nil
	}
	dst.SetInt64(1)
	for _, value := range values {
		_, err := MulTo(dst, value, dst)
		if err != nil {
			return nil, err
		}
	}
	return dst, nil
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

// RelativeChange return value of (x - y) / x
func RelativeChange(x *Decimal, y *Decimal) (*Decimal, error) {
	sp := SafePresision(x) + SafePresision(y)
	c := apd.BaseContext.WithPrecision(uint32(sp))
	tmp := new(Decimal)
	dif, err := SubTo(tmp, x, y)
	if err != nil {
		return nil, err
	}
	rc, err := DivTo(dif, c, dif, x)
	if err != nil {
		return nil, err
	}
	return rc, nil
}

func Abs(x *Decimal) *Decimal {
	return new(Decimal).Abs(x)
}
