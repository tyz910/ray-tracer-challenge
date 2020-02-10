package ray

import (
	"github.com/tyz910/ray-tracer-challenge/internal/matrix"
	"github.com/tyz910/ray-tracer-challenge/internal/tuple"
)

// Ray represents ray created by ray tracer. Ray have a starting point called the origin,
// and a vector called the direction which says where it points.
type Ray struct {
	origin    tuple.Tuple
	direction tuple.Tuple
}

// New creates new ray.
func New(origin tuple.Tuple, direction tuple.Tuple) Ray {
	if !origin.IsPoint() {
		panic("origin must be a point")
	}

	if !direction.IsVector() {
		panic("direction must be a vector")
	}

	return Ray{
		origin:    origin,
		direction: direction,
	}
}

// Origin returns starting point of the ray.
func (r Ray) Origin() tuple.Tuple {
	return r.origin
}

// Direction returns direction of the ray.
func (r Ray) Direction() tuple.Tuple {
	return r.direction
}

// Position computes the point at the given distance t along the ray.
func (r Ray) Position(t float64) tuple.Tuple {
	return r.origin.Add(r.direction.Mul(t))
}

// Transform applies the given transformation matrix to the ray,
// and returns a new ray with transformed origin and direction.
func (r Ray) Transform(m matrix.Matrix) Ray {
	return New(
		m.TupMul(r.origin),
		m.TupMul(r.direction),
	)
}
