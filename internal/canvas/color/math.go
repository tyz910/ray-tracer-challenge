package color

import (
	"github.com/tyz910/ray-tracer-challenge/internal/math"
)

// Equal approximately compares two colors.
func (c Color) Equal(c2 Color) bool {
	return math.Equals(c.Red(), c2.Red()) && math.Equals(c.Green(), c2.Green()) && math.Equals(c.Blue(), c2.Blue())
}

// Add adds two colors by adding the corresponding components.
func (c Color) Add(c2 Color) Color {
	return New(c.r+c2.r, c.g+c2.g, c.b+c2.b)
}

// Sub subtracts two colors by subtracting corresponding components.
func (c Color) Sub(c2 Color) Color {
	return New(c.r-c2.r, c.g-c2.g, c.b-c2.b)
}

// Mul multiplies each component of the color by the scalar.
func (c Color) Mul(n float64) Color {
	return New(n*c.r, n*c.g, n*c.b)
}

// Hadamard blends two colors by multiplying corresponding components of each color to form a new color.
func (c Color) Hadamard(c2 Color) Color {
	return New(c.r*c2.r, c.g*c2.g, c.b*c2.b)
}
