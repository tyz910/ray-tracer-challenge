package matrix_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tyz910/ray-tracer-challenge/internal/math/matrix"
)

// Constructing and inspecting a 4x4 matrix
func TestCreateMatrix4x4(t *testing.T) {
	// Given
	m := matrix.New(4, 4, []float64{
		1.0, 2.0, 3.0, 4.0,
		5.5, 6.5, 7.5, 8.5,
		9.0, 10.0, 11.0, 12.0,
		13.5, 14.5, 15.5, 16.5,
	})

	// Then
	assert.Equal(t, 1.0, m.Value(0, 0))
	assert.Equal(t, 4.0, m.Value(0, 3))
	assert.Equal(t, 5.5, m.Value(1, 0))
	assert.Equal(t, 7.5, m.Value(1, 2))
	assert.Equal(t, 11.0, m.Value(2, 2))
	assert.Equal(t, 13.5, m.Value(3, 0))
	assert.Equal(t, 15.5, m.Value(3, 2))
}

// A 2x2 matrix ought to be representable
func TestCreateMatrix2x2(t *testing.T) {
	// Given
	m := matrix.New(2, 2, []float64{
		-3.0, 5.0,
		1.0, -2.0,
	})

	// Then
	assert.Equal(t, -3.0, m.Value(0, 0))
	assert.Equal(t, 5.0, m.Value(0, 1))
	assert.Equal(t, 1.0, m.Value(1, 0))
	assert.Equal(t, -2.0, m.Value(1, 1))
}

// A 3x3 matrix ought to be representable
func TestCreateMatrix3x3(t *testing.T) {
	// Given
	m := matrix.New(3, 3, []float64{
		-3.0, 5.0, 0.0,
		1.0, -2.0, -7.0,
		0.0, 1.0, 1.0,
	})

	// Then
	assert.Equal(t, -3.0, m.Value(0, 0))
	assert.Equal(t, -2.0, m.Value(1, 1))
	assert.Equal(t, 1.0, m.Value(2, 2))
}
