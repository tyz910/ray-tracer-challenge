package material_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tyz910/ray-tracer-challenge/internal/canvas/color"
	"github.com/tyz910/ray-tracer-challenge/internal/render/material"
)

// The default material
func TestCreateMaterial(t *testing.T) {
	// Given
	m := material.New()

	// Then
	assert.Equal(t, color.New(1.0, 1.0, 1.0), m.Color())
	assert.Equal(t, 0.1, m.Ambient())
	assert.Equal(t, 0.9, m.Diffuse())
	assert.Equal(t, 0.9, m.Specular())
	assert.Equal(t, 200.0, m.Shininess())
}
