package bigint

import "math/big"

func Multi[I int | int64](x, y I) *Int {
	xi := new(Int).SetInt64(int64(x))
	return xi.Mul(xi, new(Int).SetInt64(int64(y)))
}

func Sum(values ...*Int) *Int {
	dst := new(big.Int)
	for _, value := range values {
		dst.Add(dst, value)
	}
	return dst
}

func Product(values ...*Int) *Int {
	dst := new(big.Int)
	for _, value := range values {
		dst.Mul(dst, value)
	}
	return dst
}

func Mul(x, y *Int) *Int {
	return MulTo(nil, x, y)
}

func MulTo(dst *Int, x, y *Int) *Int {
	if dst == nil {
		dst = new(big.Int)
	}
	return dst.Mul(x, y)
}

func Div(x, y *Int) *Int {
	return DivTo(nil, x, y)
}

func DivTo(dst *Int, x, y *Int) *Int {
	if dst == nil {
		dst = new(big.Int)
	}
	return dst.Div(x, y)
}

func MulI64[I int | int64](x *Int, y I) *Int {
	return MulI64To(nil, x, y)
}

func MulI64To[I int | int64](dst *Int, x *Int, y I) *Int {
	if dst == nil {
		dst = new(big.Int)
	}
	return dst.Mul(x, new(big.Int).SetInt64(int64(y)))
}

func Add(x, y *Int) *Int {
	return AddTo(nil, x, y)
}

func AddTo(dst *Int, x, y *Int) *Int {
	if dst == nil {
		dst = new(big.Int)
	}
	return dst.Add(x, y)
}

func AddI64To[I int | int64](dst *Int, x *Int, y I) *Int {
	if dst == nil {
		dst = new(big.Int)
	}
	return dst.Add(x, new(big.Int).SetInt64(int64(y)))
}

func Sub(x, y *Int) *Int {
	return SubTo(nil, x, y)
}

func SubTo(dst *Int, x, y *Int) *Int {
	if dst == nil {
		dst = new(big.Int)
	}
	return dst.Sub(x, y)
}

func SubI64[I int | int64](x *Int, y I) *Int {
	return SubI64To(nil, x, y)
}

func SubI64To[I int | int64](dst *Int, x *Int, y I) *Int {
	if dst == nil {
		dst = new(big.Int)
	}
	return dst.Sub(x, new(big.Int).SetInt64(int64(y)))
}

func Neg(x *Int) *Int {
	return new(Int).Neg(x)
}
