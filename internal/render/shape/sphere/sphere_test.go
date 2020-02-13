package sphere_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tyz910/ray-tracer-challenge/internal/math/matrix"
	"github.com/tyz910/ray-tracer-challenge/internal/math/tuple"
	"github.com/tyz910/ray-tracer-challenge/internal/render/material"
	"github.com/tyz910/ray-tracer-challenge/internal/render/ray"
	"github.com/tyz910/ray-tracer-challenge/internal/render/shape/sphere"
)

// A ray intersects a sphere
func TestIntersect(t *testing.T) {
	tests := []struct {
		Name      string
		Ray       ray.Ray
		Sphere    *sphere.Sphere
		ExpectedT []float64
	}{
		{
			Name:      "A ray intersects a sphere at two points",
			Ray:       ray.New(tuple.Point(0.0, 0.0, -5.0), tuple.Vector(0.0, 0.0, 1.0)),
			Sphere:    sphere.New(),
			ExpectedT: []float64{4.0, 6.0},
		},

		{
			Name:      "A ray intersects a sphere at a tangent",
			Ray:       ray.New(tuple.Point(0.0, 1.0, -5.0), tuple.Vector(0.0, 0.0, 1.0)),
			Sphere:    sphere.New(),
			ExpectedT: []float64{5.0, 5.0},
		},

		{
			Name:      "A ray misses a sphere",
			Ray:       ray.New(tuple.Point(0.0, -2.0, -5.0), tuple.Vector(0.0, 0.0, 1.0)),
			Sphere:    sphere.New(),
			ExpectedT: []float64{},
		},

		{
			Name:      "A ray originates inside a sphere",
			Ray:       ray.New(tuple.Point(0.0, 0.0, 0.0), tuple.Vector(0.0, 0.0, 1.0)),
			Sphere:    sphere.New(),
			ExpectedT: []float64{-1.0, 1.0},
		},

		{
			Name:      "A sphere is behind a ray",
			Ray:       ray.New(tuple.Point(0.0, 0.0, 5.0), tuple.Vector(0.0, 0.0, 1.0)),
			Sphere:    sphere.New(),
			ExpectedT: []float64{-6.0, -4.0},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			// Given
			r := test.Ray
			s := test.Sphere

			// When
			xs := s.Intersect(r)

			// Then
			assert.Equal(t, len(test.ExpectedT), len(xs))
			for i, x := range test.ExpectedT {
				assert.Equal(t, x, xs[i].T())
			}
		})
	}
}

// A sphere's default transformation
func TestDefaultTransform(t *testing.T) {
	// Given
	s := sphere.New()

	// Then
	assert.True(t, s.Transform().Equal(matrix.Identity()))
}

// Changing a sphere's transformation
func TestSetTransform(t *testing.T) {
	// Given
	s := sphere.New()
	m := matrix.Translation(2.0, 3.0, 4.0)

	// When
	s.SetTransform(m)

	// Then
	assert.True(t, s.Transform().Equal(m))
}

// Intersect sets the object on the intersection
func TestIntersectObject(t *testing.T) {
	// Given
	r := ray.New(
		tuple.Point(0.0, 0.0, -5.0),
		tuple.Vector(0.0, 0.0, 1.0),
	)

	s := sphere.New()

	// When
	xs := s.Intersect(r)

	// Then
	assert.Equal(t, 2, len(xs))
	assert.Equal(t, s, xs[0].Object())
	assert.Equal(t, s, xs[1].Object())
}

// Intersecting a scaled sphere with a ray
func TestIntersectScaled(t *testing.T) {
	// Given
	r := ray.New(
		tuple.Point(0.0, 0.0, -5.0),
		tuple.Vector(0.0, 0.0, 1.0),
	)

	s := sphere.New()

	// When
	s.SetTransform(matrix.Scaling(2.0, 2.0, 2.0))
	xs := s.Intersect(r)

	// Then
	assert.Equal(t, 2, len(xs))
	assert.Equal(t, 3.0, xs[0].T())
	assert.Equal(t, 7.0, xs[1].T())
}

// Intersecting a translated sphere with a ray
func TestIntersectTranslated(t *testing.T) {
	// Given
	r := ray.New(
		tuple.Point(0.0, 0.0, -5.0),
		tuple.Vector(0.0, 0.0, 1.0),
	)

	s := sphere.New()

	// When
	s.SetTransform(matrix.Translation(5.0, 0.0, 0.0))
	xs := s.Intersect(r)

	// Then
	assert.Equal(t, 0, len(xs))
}

// The normal on a sphere
func TestNormal(t *testing.T) {
	tests := []struct {
		Name   string
		Point  tuple.Tuple
		Normal tuple.Tuple
	}{
		{
			Name:   "point on the x axis",
			Point:  tuple.Point(1.0, 0.0, 0.0),
			Normal: tuple.Vector(1.0, 0.0, 0.0),
		},

		{
			Name:   "point on the y axis",
			Point:  tuple.Point(0.0, 1.0, 0.0),
			Normal: tuple.Vector(0.0, 1.0, 0.0),
		},

		{
			Name:   "point on the z axis",
			Point:  tuple.Point(0.0, 0.0, 1.0),
			Normal: tuple.Vector(0.0, 0.0, 1.0),
		},

		{
			Name:   "nonaxial point",
			Point:  tuple.Point(math.Sqrt(3.0)/3.0, math.Sqrt(3.0)/3.0, math.Sqrt(3.0)/3.0),
			Normal: tuple.Vector(math.Sqrt(3.0)/3.0, math.Sqrt(3.0)/3.0, math.Sqrt(3.0)/3.0),
		},
	}

	for _, test := range tests {
		t.Run("The normal on a sphere at a "+test.Name, func(t *testing.T) {
			// Given
			s := sphere.New()

			// When
			n := s.NormalAt(test.Point)

			// Then
			assert.True(t, n.Equal(test.Normal))
		})
	}
}

// The normal is a normalized vector
func TestNormalized(t *testing.T) {
	// Given
	s := sphere.New()

	// When
	n := s.NormalAt(tuple.Point(math.Sqrt(3.0)/3.0, math.Sqrt(3.0)/3.0, math.Sqrt(3.0)/3.0))

	// Then
	assert.True(t, n.Equal(n.Normalize()))
}

// Computing the normal on a translated sphere
func TestNormalTranslated(t *testing.T) {
	// Given
	s := sphere.New()
	s.SetTransform(matrix.Translation(0.0, 1.0, 0.0))

	// When
	n := s.NormalAt(tuple.Point(0.0, 1.70711, -0.70711))

	// Then
	assert.True(t, n.Equal(tuple.Vector(0.0, 0.70711, -0.70711)))
}

// Computing the normal on a transformed sphere
func TestNormalTransformed(t *testing.T) {
	// Given
	s := sphere.New()
	m := matrix.Scaling(1.0, 0.5, 1.0).MatMul(matrix.RotationZ(math.Pi / 5.0))
	s.SetTransform(m)

	// When
	n := s.NormalAt(tuple.Point(0.0, math.Sqrt(2.0)/2.0, -math.Sqrt(2.0)/2.0))

	// Then
	assert.True(t, n.Equal(tuple.Vector(0.0, 0.97014, -0.24254)))
}

// A sphere has a default material
func TestDefaultMaterial(t *testing.T) {
	// Given
	s := sphere.New()

	// When
	m := s.Material()
	// Then
	assert.Equal(t, material.New(), m)
}

// A sphere may be assigned a material
func TestAssignMaterial(t *testing.T) {
	// Given
	s := sphere.New()
	m := material.New()
	m.SetAmbient(1.0)

	// When
	s.SetMaterial(m)

	// Then
	assert.Equal(t, m, s.Material())
}
