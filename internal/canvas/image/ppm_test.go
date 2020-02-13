package image_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tyz910/ray-tracer-challenge/internal/canvas"
	"github.com/tyz910/ray-tracer-challenge/internal/canvas/color"
	"github.com/tyz910/ray-tracer-challenge/internal/canvas/image"
)

// Constructing the PPM header
func TestHeader(t *testing.T) {
	// Given
	c := canvas.New(5, 3)

	// When
	ppm := image.NewPPM(c).String()

	// Then
	lines := strings.Split(ppm, "\n")
	assert.Equal(t, []string{
		"P3",
		"5 3",
		"255",
	}, lines[:3])
}

// Constructing the PPM pixel data
func TestPixelData(t *testing.T) {
	// Given
	c := canvas.New(5, 3)
	c1 := color.New(1.5, 0.0, 0.0)
	c2 := color.New(0.0, 0.5, 0.0)
	c3 := color.New(-0.5, 0.0, 1.0)

	// When
	c.SetPixel(0, 0, c1)
	c.SetPixel(2, 1, c2)
	c.SetPixel(4, 2, c3)
	ppm := image.NewPPM(c).String()

	// Then
	lines := strings.Split(ppm, "\n")
	assert.Equal(t, []string{
		"255 0 0 0 0 0 0 0 0 0 0 0 0 0 0",
		"0 0 0 0 0 0 0 128 0 0 0 0 0 0 0",
		"0 0 0 0 0 0 0 0 0 0 0 0 0 0 255",
	}, lines[3:6])
}

// Splitting long lines in PPM files
func TestLongLines(t *testing.T) {
	// Given
	c := canvas.New(10, 2)

	// When
	for x := 0; x < c.Width(); x++ {
		for y := 0; y < c.Height(); y++ {
			c.SetPixel(x, y, color.New(1.0, 0.8, 0.6))
		}
	}
	ppm := image.NewPPM(c).String()

	// Then
	lines := strings.Split(ppm, "\n")
	assert.Equal(t, []string{
		"255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204",
		"153 255 204 153 255 204 153 255 204 153 255 204 153",
		"255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204",
		"153 255 204 153 255 204 153 255 204 153 255 204 153",
	}, lines[3:7])
}

// PPM files are terminated by a newline character
func TestFileTerminate(t *testing.T) {
	// Given
	c := canvas.New(5, 3)

	// When
	ppm := image.NewPPM(c).String()

	// Then
	assert.Equal(t, "\n", ppm[len(ppm)-1:])
}
