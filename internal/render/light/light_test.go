package light_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tyz910/ray-tracer-challenge/internal/canvas/color"
	"github.com/tyz910/ray-tracer-challenge/internal/math/tuple"
	"github.com/tyz910/ray-tracer-challenge/internal/render/light"
)

// A point light has a position and intensity
func TestCreateLight(t *testing.T) {
	// Given
	p := tuple.Point(0.0, 0.0, 0.0)
	i := color.New(1.0, 1.0, 1.0)

	// When
	l := light.New(p, i)

	// Then
	assert.True(t, l.Position().Equal(p))
	assert.True(t, l.Intensity().Equal(i))
}
