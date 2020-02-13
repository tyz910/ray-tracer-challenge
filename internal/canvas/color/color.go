package color

// Color is a (red, green, blue) tuple.
type Color struct {
	r, g, b float64
}

// New creates new color.
func New(r, g, b float64) Color {
	return Color{r, g, b}
}

// Red returns red component of color.
func (c Color) Red() float64 {
	return c.r
}

// Green returns green component of color.
func (c Color) Green() float64 {
	return c.g
}

// Blue returns blue component of color.
func (c Color) Blue() float64 {
	return c.b
}
