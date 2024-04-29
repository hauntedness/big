package decimal

import "github.com/cockroachdb/apd/v3"

func New(coeff int64, exponent int32) *apd.Decimal {
	return apd.New(coeff, exponent)
}

// NewFromString call apd.NewFromString but ignore returned Condition
func NewFromString(s string) (*apd.Decimal, error) {
	d, _, err := apd.NewFromString(s)
	return d, err
}

// NewFromFloat call apd.NewFromFloat with new(Decimal)
func NewFromFloat(f float64) (*apd.Decimal, error) {
	return new(Decimal).SetFloat64(f)
}
