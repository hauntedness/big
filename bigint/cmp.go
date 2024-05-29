package bigint

import "math/big"

func Equal(x, y *big.Int) bool {
	return x.Cmp(y) == 0
}

func Less(x, y *big.Int) bool {
	return x.Cmp(y) < 0
}

func LessEq(x, y *big.Int) bool {
	return x.Cmp(y) <= 0
}

func Greater(x, y *big.Int) bool {
	return x.Cmp(y) > 0
}

func GreaterEq(x, y *big.Int) bool {
	return x.Cmp(y) >= 0
}
