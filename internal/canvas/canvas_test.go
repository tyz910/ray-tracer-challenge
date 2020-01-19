package canvas_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tyz910/ray-tracer-challenge/internal/canvas"
	"github.com/tyz910/ray-tracer-challenge/internal/color"
)

// Creating a canvas
func TestCreateCanvas(t *testing.T) {
	// Given
	c := canvas.New(10, 20)

	// Then
	assert.Equal(t, 10, c.Width())
	assert.Equal(t, 20, c.Height())

	for x := 0; x < c.Width(); x++ {
		for y := 0; y < c.Height(); y++ {
			// every pixel of c is color(0, 0, 0)
			assert.True(t, c.Pixel(x, y).Equal(color.New(0.0, 0.0, 0.0)))
		}
	}
}

// Writing pixels to a canvas
func TestWritePixel(t *testing.T) {
	// Given
	c := canvas.New(10, 20)
	red := color.New(1.0, 0.0, 0.0)

	// When
	c.SetPixel(2, 3, red)

	// Then
	assert.True(t, c.Pixel(2, 3).Equal(red))
}
