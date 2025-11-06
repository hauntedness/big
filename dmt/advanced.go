package dmt

func Pow(c *Context, x, y *Decimal) (*Decimal, error) {
	dst := new(Decimal)
	return dst, SetPowTo(c, dst, x, y)
}

func SetPowTo(c *Context, dst, x, y *Decimal) error {
	_, err := c.Pow(dst, x, y)
	return err
}

// Exp return e**x.
func Exp(c *Context, x *Decimal) (*Decimal, error) {
	dst := new(Decimal)
	return dst, SetExpTo(c, dst, x)
}

// SetExpTo set dst to e**x.
func SetExpTo(c *Context, dst, x *Decimal) error {
	_, err := c.Exp(dst, x)
	return err
}

func Sqrt(c *Context, x *Decimal) (*Decimal, error) {
	d := new(Decimal)
	return d, SetSqrtTo(c, d, x)
}

func SetSqrtTo(c *Context, dst, x *Decimal) error {
	_, err := c.Sqrt(dst, x)
	return err
}

func Log10(c *Context, x *Decimal) (*Decimal, error) {
	dst := new(Decimal)
	return dst, SetLog10To(c, dst, x)
}

func SetLog10To(c *Context, dst, x *Decimal) error {
	_, err := c.Log10(dst, x)
	return err
}

func Ln(c *Context, x *Decimal) (*Decimal, error) {
	dst := new(Decimal)
	return dst, SetLnTo(c, dst, x)
}

func SetLnTo(c *Context, dst, x *Decimal) error {
	_, err := c.Ln(dst, x)
	return err
}
