package color

type Colorizer struct {
}

func New() *Colorizer {
	return &Colorizer{}
}

func (c *Colorizer) GetColor(val string) (r, g, b uint32) {
	switch val {
	case "red":
		return c.red()
	case "yellow":
		return c.yellow()
	case "purple":
		return c.purple()
	default:
		return c.red()
	}
}

func (c *Colorizer) red() (r, g, b uint32) {
	return uint32(255), uint32(0), uint32(0)
}

func (c *Colorizer) purple() (r, g, b uint32) {
	return uint32(128), uint32(0), uint32(128)
}

func (c *Colorizer) yellow() (r, g, b uint32) {
	return uint32(255), uint32(255), uint32(0)
}
