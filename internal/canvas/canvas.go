package canvas

import (
	"fmt"
	"math"

	"github.com/tyz910/ray-tracer-challenge/internal/canvas/color"
)

// Canvas is a rectangular grid of pixels.
type Canvas struct {
	width, height int
	pixels        []color.Color
}

// New creates new canvas.
func New(width, height int) Canvas {
	if (width < 1) || (height < 1) {
		panic(fmt.Sprintf("invalid canvas size (%d, %d)", width, height))
	}

	return Canvas{
		width:  width,
		height: height,
		pixels: make([]color.Color, width*height),
	}
}

// Width returns the width of the canvas.
func (cnv Canvas) Width() int {
	return cnv.width
}

// Height returns the height of the canvas.
func (cnv Canvas) Height() int {
	return cnv.height
}

// Contains checks whether position (x, y) within canvas.
func (cnv Canvas) Contains(x, y int) bool {
	return (x >= 0) && (x < cnv.width) && (y >= 0) && (y < cnv.height)
}

func (cnv Canvas) index(x, y int) int {
	if !cnv.Contains(x, y) {
		panic(fmt.Sprintf("index (%d, %d) out of range for canvas size (%d, %d)", x, y, cnv.width, cnv.height))
	}

	return y*cnv.width + x
}

// Pixel returns the color of the pixel at position (x, y).
func (cnv Canvas) Pixel(x, y int) color.Color {
	return cnv.pixels[cnv.index(x, y)]
}

// SetPixel sets the color of the pixel at position (x, y) to c.
func (cnv Canvas) SetPixel(x, y int, c color.Color) {
	cnv.pixels[cnv.index(x, y)] = c
}

// WritePixel sets the color of the pixel at position (x, y) to c if  it's position within canvas.
func (cnv Canvas) WritePixel(x, y float64, c color.Color) {
	xPos := int(math.Ceil(x))
	yPos := int(math.Ceil(y))

	if cnv.Contains(xPos, yPos) {
		cnv.SetPixel(xPos, yPos, c)
	}
}
