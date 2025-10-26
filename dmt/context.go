package dmt

import "github.com/cockroachdb/apd/v3"

func WithRounding(precision int, rounding apd.Rounder) *Context {
	c := apd.BaseContext.WithPrecision(uint32(precision))
	c.Rounding = rounding
	return c
}

// WithPrecision returns a copy of c but with the specified precision.
func WithPrecision(precision int) *Context {
	return apd.BaseContext.WithPrecision(uint32(precision))
}
