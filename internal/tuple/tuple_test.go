package tuple_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tyz910/ray-tracer-challenge/internal/tuple"
)

// A tuple with w=1.0 is a point
func TestCreateTupleAsPoint(t *testing.T) {
	// Given
	a := tuple.New(4.3, -4.2, 3.1, 1.0)

	// Then
	assert.Equal(t, 4.3, a.X())
	assert.Equal(t, -4.2, a.Y())
	assert.Equal(t, 3.1, a.Z())
	assert.Equal(t, 1.0, a.W())
	assert.True(t, a.IsPoint())
	assert.False(t, a.IsVector())
}

// A tuple with w=0 is a vector
func TestCreateTupleAsVector(t *testing.T) {
	// Given
	a := tuple.New(4.3, -4.2, 3.1, 0.0)

	// Then
	assert.Equal(t, 4.3, a.X())
	assert.Equal(t, -4.2, a.Y())
	assert.Equal(t, 3.1, a.Z())
	assert.Equal(t, 0.0, a.W())
	assert.True(t, a.IsVector())
	assert.False(t, a.IsPoint())
}

// Point() creates tuples with w=1
func TestCreatePoint(t *testing.T) {
	// Given
	p := tuple.Point(4.0, -4.0, 3.0)

	// Then
	assert.True(t, p.Equal(tuple.New(4.0, -4.0, 3.0, 1.0)))
}

// Vector() creates tuples with w=0
func TestCreateVector(t *testing.T) {
	// Given
	p := tuple.Vector(4.0, -4.0, 3.0)

	// Then
	assert.True(t, p.Equal(tuple.New(4.0, -4.0, 3.0, 0.0)))
}
