package render_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tyz910/ray-tracer-challenge/internal/canvas/color"
	"github.com/tyz910/ray-tracer-challenge/internal/math/tuple"
	"github.com/tyz910/ray-tracer-challenge/internal/render"
	"github.com/tyz910/ray-tracer-challenge/internal/render/light"
	"github.com/tyz910/ray-tracer-challenge/internal/render/material"
)

func TestLighting(t *testing.T) {
	tests := []struct {
		Name      string
		EyeVec    tuple.Tuple
		NormalVec tuple.Tuple
		Light     light.Light
		Color     color.Color
	}{
		{
			Name:      "Lighting with the eye between the light and the surface",
			EyeVec:    tuple.Vector(0.0, 0.0, -1.0),
			NormalVec: tuple.Vector(0.0, 0.0, -1.0),
			Light:     light.New(tuple.Point(0.0, 0.0, -10.0), color.White()),
			Color:     color.New(1.9, 1.9, 1.9),
		},

		{
			Name:      "Lighting with the eye between light and surface, eye offset 45°",
			EyeVec:    tuple.Vector(0.0, math.Sqrt(2.0)/2.0, -math.Sqrt(2.0)/2.0),
			NormalVec: tuple.Vector(0.0, 0.0, -1.0),
			Light:     light.New(tuple.Point(0.0, 0.0, -10.0), color.White()),
			Color:     color.New(1.0, 1.0, 1.0),
		},

		{
			Name:      "Lighting with eye opposite surface, light offset 45°",
			EyeVec:    tuple.Vector(0.0, 0.0, -1.0),
			NormalVec: tuple.Vector(0.0, 0.0, -1.0),
			Light:     light.New(tuple.Point(0.0, 10.0, -10.0), color.White()),
			Color:     color.New(0.7364, 0.7364, 0.7364),
		},

		{
			Name:      "Lighting with eye in the path of the reflection vector",
			EyeVec:    tuple.Vector(0.0, -math.Sqrt(2.0)/2.0, -math.Sqrt(2.0)/2.0),
			NormalVec: tuple.Vector(0.0, 0.0, -1.0),
			Light:     light.New(tuple.Point(0.0, 10.0, -10.0), color.White()),
			Color:     color.New(1.6364, 1.6364, 1.6364),
		},

		{
			Name:      "Lighting with the light behind the surface",
			EyeVec:    tuple.Vector(0.0, 0.0, -1.0),
			NormalVec: tuple.Vector(0.0, 0.0, -1.0),
			Light:     light.New(tuple.Point(0.0, 0.0, 10.0), color.White()),
			Color:     color.New(0.1, 0.1, 0.1),
		},
	}

	// Background
	m := material.New()
	p := tuple.Point(0.0, 0.0, 0.0)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			// Given
			l := test.Light
			eyeVec := test.EyeVec
			normalVec := test.NormalVec

			// When
			result := render.Lighting(m, l, p, eyeVec, normalVec)

			// Then
			assert.True(t, test.Color.Equal(result))
		})
	}
}
