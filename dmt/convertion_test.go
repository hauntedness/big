package dmt_test

import (
	"fmt"
	"log/slog"
	"math/big"
	"testing"

	"github.com/cockroachdb/apd/v3"
	"github.com/hauntedness/big/dmt"
)

func TestRoundToIntegral(t *testing.T) {
	tests := map[string]struct {
		ctx   *dmt.Context
		dst   *dmt.Decimal
		x     *dmt.Decimal
		check func(*big.Int, error) error
	}{
		"decimal only": {
			ctx: dmt.WithRounding(40, apd.Round05Up),
			dst: &dmt.Decimal{},
			x:   dmt.MustFromString("0.000001243"),
			check: func(res *big.Int, err error) error {
				if err != nil {
					return err
				}
				if res.Cmp(big.NewInt(0)) != 0 {
					return fmt.Errorf("%v: ", res)
				}
				return nil
			},
		},
		"negative decimal only": {
			ctx: dmt.WithRounding(40, apd.Round05Up),
			dst: &dmt.Decimal{},
			x:   dmt.MustFromString("-0.000001243"),
			check: func(res *big.Int, err error) error {
				if err != nil {
					return err
				}
				if res.Cmp(big.NewInt(0)) != 0 {
					return fmt.Errorf("%v: ", res)
				}
				return nil
			},
		},
		"integral and decimal": {
			ctx: dmt.WithRounding(40, apd.RoundHalfEven),
			dst: &dmt.Decimal{},
			x:   dmt.MustFromString("103.000001243"),
			check: func(res *big.Int, err error) error {
				if err != nil {
					return err
				}
				if res.Cmp(big.NewInt(103)) != 0 {
					return fmt.Errorf("%v: ", res)
				}
				return nil
			},
		},
		"negative integral and decimal": {
			ctx: dmt.WithRounding(40, apd.RoundHalfEven),
			dst: &dmt.Decimal{},
			x:   dmt.MustFromString("-103.000001243"),
			check: func(res *big.Int, err error) error {
				if err != nil {
					return err
				}
				if res.Cmp(big.NewInt(-103)) != 0 {
					return fmt.Errorf("%v: ", res)
				}
				return nil
			},
		},
		"integral and decimal with exponent": {
			ctx: dmt.WithRounding(40, apd.RoundHalfEven),
			dst: &dmt.Decimal{},
			x:   dmt.New(154266543, 10),
			check: func(res *big.Int, err error) error {
				if err != nil {
					return err
				}
				dst := big.NewInt(154266543)
				m := new(big.Int).Exp(big.NewInt(10), big.NewInt(10), nil)
				if res.Cmp(dst.Mul(dst, m)) != 0 {
					return fmt.Errorf("%v: ", res)
				}
				return nil
			},
		},
		"negative integral and decimal with exponent": {
			ctx: dmt.WithRounding(40, apd.RoundHalfEven),
			dst: &dmt.Decimal{},
			x:   dmt.New(-154266543, 10),
			check: func(res *big.Int, err error) error {
				if err != nil {
					return err
				}
				dst := big.NewInt(-154266543)
				m := new(big.Int).Exp(big.NewInt(10), big.NewInt(10), nil)
				if res.Cmp(dst.Mul(dst, m)) != 0 {
					return fmt.Errorf("%v: ", res)
				}
				return nil
			},
		},
		"negative integral and decimal with negative exponent": {
			ctx: dmt.WithRounding(40, apd.RoundHalfEven),
			dst: &dmt.Decimal{},
			x:   dmt.New(-154266543, -4),
			check: func(res *big.Int, err error) error {
				if err != nil {
					return err
				}
				dst := big.NewInt(-154266543)
				m := new(big.Int).Exp(big.NewInt(10), big.NewInt(4), nil)
				wa := dst.Div(dst, m)
				if res.Cmp(wa) != 0 {
					return fmt.Errorf("%v", res)
				}
				return nil
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := dmt.ToIntegral(tt.ctx, tt.dst, tt.x)
			if err := tt.check(result, err); err != nil {
				t.Errorf("RoundToIntegral() = %v", err)
			} else {
				slog.Info("RoundToIntegral", "value", result.String())
			}
		})
	}
}
