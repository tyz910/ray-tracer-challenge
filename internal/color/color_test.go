package color_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tyz910/ray-tracer-challenge/internal/color"
)

// Colors are (red, green, blue) tuples
func TestCreateColor(t *testing.T) {
	// Given
	c := color.New(-0.5, 0.4, 1.7)

	// Then
	assert.Equal(t, -0.5, c.Red())
	assert.Equal(t, 0.4, c.Green())
	assert.Equal(t, 1.7, c.Blue())
}
