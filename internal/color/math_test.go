package color_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tyz910/ray-tracer-challenge/internal/color"
)

// Adding colors
func TestAddColors(t *testing.T) {
	// Given
	c1 := color.New(0.9, 0.6, 0.75)
	c2 := color.New(0.7, 0.1, 0.25)

	// Then
	assert.True(t, c1.Add(c2).Equal(color.New(1.6, 0.7, 1.0)))
}

// Subtracting colors
func TestSubColors(t *testing.T) {
	// Given
	c1 := color.New(0.9, 0.6, 0.75)
	c2 := color.New(0.7, 0.1, 0.25)

	// Then
	assert.True(t, c1.Sub(c2).Equal(color.New(0.2, 0.5, 0.5)))
}

// Multiplying a color by a scalar
func TestMulScalar(t *testing.T) {
	// Given
	c := color.New(0.2, 0.3, 0.4)

	// Then
	assert.True(t, c.Mul(2.0).Equal(color.New(0.4, 0.6, 0.8)))
}

// Multiplying colors
func TestHadamard(t *testing.T) {
	// Given
	c1 := color.New(1.0, 0.2, 0.4)
	c2 := color.New(0.9, 1.0, 0.1)

	// Then
	assert.True(t, c1.Hadamard(c2).Equal(color.New(0.9, 0.2, 0.04)))
}
