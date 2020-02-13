package sphere

import (
	"math"

	"github.com/tyz910/ray-tracer-challenge/internal/math/matrix"
	"github.com/tyz910/ray-tracer-challenge/internal/math/tuple"
	"github.com/tyz910/ray-tracer-challenge/internal/render/material"
	"github.com/tyz910/ray-tracer-challenge/internal/render/ray"
	"github.com/tyz910/ray-tracer-challenge/internal/render/shape"
)

// Sphere represents a sphere object.
type Sphere struct {
	transform matrix.Matrix
	material  material.Material
}

// New creates new sphere.
func New() *Sphere {
	return &Sphere{
		transform: matrix.Identity(),
		material:  material.New(),
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

// Material returns the surface material of the sphere.
func (s *Sphere) Material() material.Material {
	return s.material
}

// SetMaterial changes the surface material of the sphere.
func (s *Sphere) SetMaterial(m material.Material) {
	s.material = m
}

// Intersect returns the collection of intersections where the ray intersects the sphere.
func (s *Sphere) Intersect(r ray.Ray) shape.Intersections {
	rLocal := r.Transform(s.transform.Inverse())

	// the vector from the sphere's center, to the ray origin
	sphereToRay := rLocal.Origin().Sub(tuple.Point(0.0, 0.0, 0.0))

	a := rLocal.Direction().Dot(rLocal.Direction())
	b := 2.0 * rLocal.Direction().Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1.0

	discriminant := b*b - 4.0*a*c
	if discriminant < 0.0 {
		return shape.Intersections{}
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2.0 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2.0 * a)

	i1 := shape.NewIntersection(t1, s)
	i2 := shape.NewIntersection(t2, s)

	if t1 > t2 {
		return shape.Intersections{i2, i1}
	}

	return shape.Intersections{i1, i2}
}

// NormalAt returns the normal on the sphere at the given point.
func (s *Sphere) NormalAt(p tuple.Tuple) tuple.Tuple {
	ti := s.transform.Inverse()
	pLocal := ti.TupMul(p)
	nLocal := pLocal.Sub(tuple.Point(0.0, 0.0, 0.0))
	nWorld := ti.Transpose().TupMul(nLocal).AsVector()

	return nWorld.Normalize()
}
