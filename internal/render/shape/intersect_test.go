package shape_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tyz910/ray-tracer-challenge/internal/render/shape"
	"github.com/tyz910/ray-tracer-challenge/internal/render/shape/sphere"
)

// An intersection encapsulates t and object
func TestCreate(t *testing.T) {
	// Given
	s := sphere.New()

	// When
	i := shape.NewIntersection(3.5, s)

	// Then
	assert.Equal(t, 3.5, i.T())
	assert.Equal(t, s, i.Object())
}

// The hit, when all intersections have positive t
func TestHitPositive(t *testing.T) {
	// Given
	s := sphere.New()
	i1 := shape.NewIntersection(1.0, s)
	i2 := shape.NewIntersection(2.0, s)
	xs := shape.Intersections{i1, i2}

	// When
	i := xs.Hit()

	// Then
	assert.Equal(t, i1, i)
}

// The hit, when some intersections have negative t
func TestHitNegative(t *testing.T) {
	// Given
	s := sphere.New()
	i1 := shape.NewIntersection(-1.0, s)
	i2 := shape.NewIntersection(1.0, s)
	xs := shape.Intersections{i2, i1}

	// When
	i := xs.Hit()

	// Then
	assert.Equal(t, i2, i)
}

// The hit, when all intersections have negative t
func TestHitAllNegative(t *testing.T) {
	// Given
	s := sphere.New()
	i1 := shape.NewIntersection(-2.0, s)
	i2 := shape.NewIntersection(-1.0, s)
	xs := shape.Intersections{i2, i1}

	// When
	i := xs.Hit()

	// Then
	assert.Nil(t, i)
}

// The hit is always the lowest nonnegative intersection
func TestHitLowest(t *testing.T) {
	// Given
	s := sphere.New()
	i1 := shape.NewIntersection(5.0, s)
	i2 := shape.NewIntersection(7.0, s)
	i3 := shape.NewIntersection(-3.0, s)
	i4 := shape.NewIntersection(2.0, s)
	xs := shape.Intersections{i1, i2, i3, i4}

	// When
	i := xs.Hit()

	// Then
	assert.Equal(t, i4, i)
}
