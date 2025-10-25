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
func SetScale(ctx *Context, dst *Decimal, scale int) error {
	return setScaleTo(ctx, dst, scale)
}

// SetScaleTo is similar to SetScale but avoid mutating the value of x
func SetScaleTo(ctx *Context, dst *Decimal, x *Decimal, precision, scale int, rounder apd.Rounder) error {
	dst.Set(x)
	return setScaleTo(ctx, dst, scale)
}

func setScaleTo(ctx *Context, dst *Decimal, scale int) error {
	if dst.Form != apd.Finite {
		return nil
	}
	// Use +1 here because it is inverted later.
	if scale < math.MinInt32+1 || scale > math.MaxInt32 {
		return errors.New("scale out of range")
	}
	if scale > int(ctx.Precision) {
		return fmt.Errorf("scale (%d) must be between 0 and precision (%d)", scale, ctx.Precision)
	}
	_, err := ctx.Quantize(dst, dst, -int32(scale))
	return err
}

func Rem(ctx *Context, x *Decimal, y *Decimal) (*Decimal, error) {
	var remainder = new(Decimal)

	_, err := ctx.Rem(remainder, x, y)
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
