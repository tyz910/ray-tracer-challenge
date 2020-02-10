package ray_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tyz910/ray-tracer-challenge/internal/matrix"
	"github.com/tyz910/ray-tracer-challenge/internal/ray"
	"github.com/tyz910/ray-tracer-challenge/internal/tuple"
)

// Creating and querying a ray
func TestCreate(t *testing.T) {
	// Given
	origin := tuple.Point(1.0, 2.0, 3.0)
	direction := tuple.Vector(4.0, 5.0, 6.0)

	// When
	r := ray.New(origin, direction)

	// Then
	assert.True(t, r.Origin().Equal(origin))
	assert.True(t, r.Direction().Equal(direction))
}

// Computing a point from a distance
func TestPosition(t *testing.T) {
	// Given
	r := ray.New(
		tuple.Point(2.0, 3.0, 4.0),
		tuple.Vector(1.0, 0.0, 0.0),
	)

	// Then
	assert.True(t, r.Position(0.0).Equal(tuple.Point(2.0, 3.0, 4.0)))
	assert.True(t, r.Position(1.0).Equal(tuple.Point(3.0, 3.0, 4.0)))
	assert.True(t, r.Position(-1.0).Equal(tuple.Point(1.0, 3.0, 4.0)))
	assert.True(t, r.Position(2.5).Equal(tuple.Point(4.5, 3.0, 4.0)))
}

// Translating a ray
func TestTranslation(t *testing.T) {
	// Given
	r := ray.New(
		tuple.Point(1.0, 2.0, 3.0),
		tuple.Vector(0.0, 1.0, 0.0),
	)

	m := matrix.Translation(3.0, 4.0, 5.0)

	// When
	r2 := r.Transform(m)

	// Then
	assert.True(t, r2.Origin().Equal(tuple.Point(4.0, 6.0, 8.0)))
	assert.True(t, r2.Direction().Equal(tuple.Vector(0.0, 1.0, 0.0)))
}

// Scaling a ray
func TestScaling(t *testing.T) {
	// Given
	r := ray.New(
		tuple.Point(1.0, 2.0, 3.0),
		tuple.Vector(0.0, 1.0, 0.0),
	)

	m := matrix.Scaling(2.0, 3.0, 4.0)

	// When
	r2 := r.Transform(m)

	// Then
	assert.True(t, r2.Origin().Equal(tuple.Point(2.0, 6.0, 12.0)))
	assert.True(t, r2.Direction().Equal(tuple.Vector(0.0, 3.0, 0.0)))
}
