package dmt

import "math/big"

// StringFixed format d as FixedPoint:
// 'f' -ddddd.dddd, no exponent
func StringFixed(d *Decimal) string {
	return d.Text('f')
}

func ToBigFloat(x *Decimal) (*big.Float, bool) {
	return new(big.Float).SetString(x.Text('G'))
}

func MustToIntegral(c *Context, x *Decimal) *big.Int {
	dst := new(Decimal)
	res, err := ToIntegral(c, dst, x)
	if err != nil {
		panic(err)
	}
	return res
}

func ToIntegral(c *Context, dst, x *Decimal) (*big.Int, error) {
	_, err := c.RoundToIntegralValue(dst, x)
	if err != nil {
		return nil, err
	}
	if x.Negative {
		bigint := dst.Coeff.MathBigInt()
		return bigint.Neg(bigint), nil
	} else {
		return dst.Coeff.MathBigInt(), nil
	}
}
