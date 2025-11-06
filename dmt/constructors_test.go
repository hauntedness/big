package dmt_test

import (
	"log/slog"
	"math"
	"testing"

	"github.com/cockroachdb/apd/v3"
	"github.com/hauntedness/big/dmt"
)

func TestMustFromFloat(t *testing.T) {
	tests := map[string]struct {
		input float64
		check func(*apd.Decimal, error) error
	}{
		"test 0": {
			input: 0,
			check: func(d *apd.Decimal, err error) error {
				return err
			},
		},
		"test -0.1": {
			input: float64(-0.1),
			check: func(d *apd.Decimal, err error) error {
				return err
			},
		},
		"test 0.3": {
			input: float64(0.1) + float64(0.2),
			check: func(d *apd.Decimal, err error) error {
				return err
			},
		},
		"test Nan": {
			input: math.NaN(),
			check: func(d *apd.Decimal, err error) error {
				return err
			},
		},
		"test Inf": {
			input: math.Inf(-1),
			check: func(d *apd.Decimal, err error) error {
				return err
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := dmt.MustFromFloat(tt.input)
			if err := tt.check(res, err); err != nil {
				t.Fatalf("MustFromFloat() failed: %v", err)
			} else {
				slog.Info("convert float to decimal", "result", res)
			}
		})
	}
}
