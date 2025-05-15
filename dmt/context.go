package dmt

import "github.com/cockroachdb/apd/v3"

var RoundDownContext = func() *apd.Context {
	c := apd.BaseContext
	c.Rounding = apd.RoundDown
	return &c
}()
