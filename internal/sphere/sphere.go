package sphere

import (
	"math"

	"github.com/tyz910/ray-tracer-challenge/internal/intersect"
	"github.com/tyz910/ray-tracer-challenge/internal/matrix"
	"github.com/tyz910/ray-tracer-challenge/internal/ray"
	"github.com/tyz910/ray-tracer-challenge/internal/tuple"
)

// Sphere represents a sphere object.
type Sphere struct {
	transform matrix.Matrix
}

// New creates new sphere.
func New() *Sphere {
	return &Sphere{
		transform: matrix.Identity(),
	}
}

// Transform returns the transformation matrix assigned to the sphere.
func (s *Sphere) Transform() matrix.Matrix {
	return s.transform
}

// SetTransform assigns transformation matrix to the sphere.
func (s *Sphere) SetTransform(m matrix.Matrix) {
	s.transform = m
}

// Intersect returns the collection of intersections where the ray intersects the sphere.
func (s *Sphere) Intersect(r ray.Ray) intersect.Intersections {
	r2 := r.Transform(s.transform.Inverse())

	// the vector from the sphere's center, to the ray origin
	sphereToRay := r2.Origin().Sub(tuple.Point(0.0, 0.0, 0.0))

	a := r2.Direction().Dot(r2.Direction())
	b := 2.0 * r2.Direction().Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1.0

	discriminant := b*b - 4.0*a*c
	if discriminant < 0.0 {
		return intersect.Intersections{}
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2.0 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2.0 * a)

	i1 := intersect.New(t1, s)
	i2 := intersect.New(t2, s)

	if t1 > t2 {
		return intersect.Intersections{i2, i1}
	}

	return intersect.Intersections{i1, i2}
}
