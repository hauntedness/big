package dmt

import "github.com/cockroachdb/apd/v3"

var RoundDownContext = func() *Context {
	c := apd.BaseContext
	c.Rounding = apd.RoundDown
	return &c
}()

func WithRounding(precision int, rounding apd.Rounder) *Context {
	c := apd.BaseContext.WithPrecision(uint32(precision))
	c.Rounding = rounding
	return c
}
