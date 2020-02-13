package render

import (
	"math"

	"github.com/tyz910/ray-tracer-challenge/internal/canvas/color"
	"github.com/tyz910/ray-tracer-challenge/internal/math/tuple"
	"github.com/tyz910/ray-tracer-challenge/internal/render/light"
	"github.com/tyz910/ray-tracer-challenge/internal/render/material"
)

// Lighting calculates color for the point on the surface. It expects five arguments:
// the material of the surface, the point being illuminated, the light source,
// and the eye and normal vectors from the Phong reflection model.
func Lighting(m material.Material, l light.Light, point, eyeVec, normalVec tuple.Tuple) color.Color {
	var ambient, diffuse, specular color.Color

	// combine the surface color with the light's color/intensity
	effectiveColor := m.Color().Hadamard(l.Intensity())

	// find the direction to the light source
	lightVec := l.Position().Sub(point).Normalize()

	// compute the ambient contribution
	ambient = effectiveColor.Mul(m.Ambient())

	// lightDotNormal represents the cosine of the angle between the light vector and the normal vector.
	// A negative number means the light is on the other side of the surface.
	lightDotNormal := lightVec.Dot(normalVec)

	if lightDotNormal < 0.0 {
		diffuse = color.Black()
		specular = color.Black()
	} else {
		// compute the diffuse contribution
		diffuse = effectiveColor.Mul(m.Diffuse() * lightDotNormal)

		// reflectDotEye represents the cosine of the angle between the reflection vector and the eye vector.
		// A negative number means the light reflects away from the eye.
		reflectDotEye := lightVec.Negate().Reflect(normalVec).Dot(eyeVec)

		if reflectDotEye <= 0.0 {
			specular = color.Black()
		} else {
			// compute the specular contribution
			factor := math.Pow(reflectDotEye, m.Shininess())
			specular = l.Intensity().Mul(m.Specular() * factor)
		}
	}

	// add the three contributions together to get the final shading
	return ambient.Add(diffuse).Add(specular)
}
