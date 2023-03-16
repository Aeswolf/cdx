package cd

type CDX struct {
	src,
	dest string
}

func NewCdx(s string) *CDX {
	return &CDX{
		src: s,
	}
}

func (c *CDX) SetSrc(src string) {
	c.src = src
}

func (c *CDX) SetDest(dest string) {
	c.dest = dest
}

func (c CDX) GetSrc() string {
	return c.src
}

func (c CDX) GetDest() string {
	return c.dest
}
