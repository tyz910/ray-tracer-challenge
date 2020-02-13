package color

// Black returns the black color.
func Black() Color {
	return New(0.0, 0.0, 0.0)
}

// White returns the white color.
func White() Color {
	return New(1.0, 1.0, 1.0)
}

// Red returns the red color.
func Red() Color {
	return New(1.0, 0.0, 0.0)
}

// Magenta returns the magenta color.
func Magenta() Color {
	return New(1.0, 0.0, 1.0)
}
