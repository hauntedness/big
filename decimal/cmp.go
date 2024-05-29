package decimal

func Equal(x, y *Decimal) bool {
	return x.Cmp(y) == 0
}

func Less(x, y *Decimal) bool {
	return x.Cmp(y) < 0
}

func LessEq(x, y *Decimal) bool {
	return x.Cmp(y) <= 0
}

func Greater(x, y *Decimal) bool {
	return x.Cmp(y) > 0
}

func GreaterEq(x, y *Decimal) bool {
	return x.Cmp(y) >= 0
}
