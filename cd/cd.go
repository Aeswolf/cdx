package cd

type cdx struct {
	src,
	dest string
}

func NewCdx(s string) *cdx {
	return &cdx{
		src: s,
	}
}

func (c *cdx) SetSrc(src string) {
	c.src = src
}

func (c *cdx) SetDest(dest string) {
	c.dest = dest
}

func (c cdx) GetSrc() string {
	return c.src
}

func (c cdx) GetDest() string {
	return c.dest
}
