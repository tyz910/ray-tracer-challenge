package shape

import (
	"github.com/tyz910/ray-tracer-challenge/internal/math/tuple"
	"github.com/tyz910/ray-tracer-challenge/internal/render/material"
	"github.com/tyz910/ray-tracer-challenge/internal/render/ray"
)

// Shape is the interface implemented by objects that can be rendered.
type Shape interface {
	// Intersect returns the collection of intersections where the ray intersects the object.
	Intersect(r ray.Ray) Intersections

	// NormalAt returns the normal on the object at the given point.
	NormalAt(p tuple.Tuple) tuple.Tuple

	// Material returns the surface material of the object.
	Material() material.Material
}
