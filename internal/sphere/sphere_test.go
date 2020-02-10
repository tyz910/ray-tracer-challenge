package sphere_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tyz910/ray-tracer-challenge/internal/matrix"
	"github.com/tyz910/ray-tracer-challenge/internal/ray"
	"github.com/tyz910/ray-tracer-challenge/internal/sphere"
	"github.com/tyz910/ray-tracer-challenge/internal/tuple"
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
