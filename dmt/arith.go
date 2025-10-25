package dmt

import (
	"errors"
	"fmt"
	"math"

	"github.com/cockroachdb/apd/v3"
)

// Add is for convenience to call AddTo
func Add(x *Decimal, y *Decimal) (*Decimal, error) {
	dst := new(Decimal)
	err := AddTo(dst, x, y)
	return dst, err
}

// AddTo sets dst to the sum x+y and return dst.
// if dst is nil, new(Decimal) will be used
func AddTo(dst *Decimal, x *Decimal, y *Decimal) error {
	_, err := apd.BaseContext.Add(dst, x, y)
	return err
}

// Sub is for convenience to call SubTo
func Sub(x *Decimal, y *Decimal) (*Decimal, error) {
	dst := new(Decimal)
	err := SubTo(dst, x, y)
	return dst, err
}

// SubTo sets dst to the difference x-y and return dst.
// if dst is nil, new(Decimal) will be used
func SubTo(dst *Decimal, x *Decimal, y *Decimal) error {
	_, err := apd.BaseContext.Sub(dst, x, y)
	return err
}

// Mul is for convenience to call MulTo
func Mul(x *Decimal, y *Decimal) (*Decimal, error) {
	dst := new(Decimal)
	err := MulTo(dst, x, y)
	return dst, err
}

// MulTo sets dst to the product x*y and return dst.
// if dst is nil, new(Decimal) will be used
func MulTo(dst *Decimal, x *Decimal, y *Decimal) error {
	_, err := apd.BaseContext.Mul(dst, x, y)
	return err
}

// Div is for convenience to call DivTo
func Div(c *Context, x *Decimal, y *Decimal) (*Decimal, error) {
	dst := new(Decimal)
	err := DivTo(dst, c, x, y)
	return dst, err
}

// DivTo sets dst to the quotient x/y for y != 0 and return dst.
// if dst is nil, new(Decimal) will be used
// c.Precision must be > 0.
// If an exact division is required, use a context with high precision
// and verify it was exact by checking the Inexact flag on the return Condition.
func DivTo(dst *Decimal, c *Context, x *Decimal, y *Decimal) error {
	_, err := c.Quo(dst, x, y)
	return err
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

func Rem(x *Decimal, y *Decimal, precision int) (*Decimal, error) {
	c := &Context{
		Precision:   uint32(precision), //nolint:gosec
		MaxExponent: apd.MaxExponent,
		MinExponent: apd.MinExponent,
		Traps:       apd.DefaultTraps,
	}

	var remainder = new(Decimal)

	_, err := c.Rem(remainder, x, y)
	if err != nil {
		return nil, err
	}

	return remainder, nil
}

// RelativeChange return value of (x - y) / x
func RelativeChange(x *Decimal, y *Decimal) (*Decimal, error) {
	sp := SafePresision(x) + SafePresision(y)
	c := apd.BaseContext.WithPrecision(uint32(sp))
	dif := new(Decimal)
	err := SubTo(dif, x, y)
	if err != nil {
		return nil, err
	}
	err = DivTo(dif, c, dif, x)
	if err != nil {
		return nil, err
	}
	return dif, nil
}

func Abs(x *Decimal) *Decimal {
	return new(Decimal).Abs(x)
}
