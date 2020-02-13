package light

import (
	"github.com/tyz910/ray-tracer-challenge/internal/canvas/color"
	"github.com/tyz910/ray-tracer-challenge/internal/math/tuple"
)

// Light represents a light source with no size, existing at a single point in space.
type Light struct {
	position  tuple.Tuple
	intensity color.Color
}

// New creates new point light.
func New(position tuple.Tuple, intensity color.Color) Light {
	return Light{
		position:  position,
		intensity: intensity,
	}
}

// Position returns the position of the light source.
func (l Light) Position() tuple.Tuple {
	return l.position
}

// Intensity returns the color of the light source.
func (l Light) Intensity() color.Color {
	return l.intensity
}
