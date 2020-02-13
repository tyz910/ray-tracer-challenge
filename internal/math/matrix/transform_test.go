package matrix_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tyz910/ray-tracer-challenge/internal/math/matrix"
	"github.com/tyz910/ray-tracer-challenge/internal/math/tuple"
)

// Multiplying by a translation matrix
func TestTranslation(t *testing.T) {
	// Given
	transform := matrix.Translation(5.0, -3.0, 2.0)
	p := tuple.Point(-3.0, 4.0, 5.0)

	// Then
	assert.True(t, transform.TupMul(p).Equal(tuple.Point(2.0, 1.0, 7.0)))
}

// Multiplying by the inverse of a translation matrix
func TestTranslationInverse(t *testing.T) {
	// Given
	transform := matrix.Translation(5.0, -3.0, 2.0)
	inv := transform.Inverse()
	p := tuple.Point(-3.0, 4.0, 5.0)

	// Then
	assert.True(t, inv.TupMul(p).Equal(tuple.Point(-8.0, 7.0, 3.0)))
}

// Translation does not affect vectors
func TestTranslationVector(t *testing.T) {
	// Given
	transform := matrix.Translation(5.0, -3.0, 2.0)
	v := tuple.Vector(-3.0, 4.0, 5.0)

	// Then
	assert.True(t, transform.TupMul(v).Equal(v))
}

// A scaling matrix applied to a point
func TestScalingPoint(t *testing.T) {
	// Given
	transform := matrix.Scaling(2.0, 3.0, 4.0)
	p := tuple.Point(-4.0, 6.0, 8.0)

	// Then
	assert.True(t, transform.TupMul(p).Equal(tuple.Point(-8.0, 18.0, 32.0)))
}

// A scaling matrix applied to a vector
func TestScalingVector(t *testing.T) {
	// Given
	transform := matrix.Scaling(2.0, 3.0, 4.0)
	v := tuple.Vector(-4.0, 6.0, 8.0)

	// Then
	assert.True(t, transform.TupMul(v).Equal(tuple.Vector(-8.0, 18.0, 32.0)))
}

// Multiplying by the inverse of a scaling matrix
func TestScalingInverse(t *testing.T) {
	// Given
	transform := matrix.Scaling(2.0, 3.0, 4.0)
	inv := transform.Inverse()
	v := tuple.Vector(-4.0, 6.0, 8.0)

	// Then
	assert.True(t, inv.TupMul(v).Equal(tuple.Vector(-2.0, 2.0, 2.0)))
}

// Reflection is scaling by a negative value
func TestScalingNegative(t *testing.T) {
	// Given
	transform := matrix.Scaling(-1.0, 1.0, 1.0)
	p := tuple.Point(2.0, 3.0, 4.0)

	// Then
	assert.True(t, transform.TupMul(p).Equal(tuple.Point(-2.0, 3.0, 4.0)))
}

// Rotating a point around the x axis
func TestRotationX(t *testing.T) {
	// Given
	p := tuple.Point(0.0, 1.0, 0.0)
	halfQuarter := matrix.RotationX(math.Pi / 4.0)
	fullQuarter := matrix.RotationX(math.Pi / 2.0)

	// Then
	assert.True(t, halfQuarter.TupMul(p).Equal(tuple.Point(0.0, math.Sqrt(2.0)/2.0, math.Sqrt(2.0)/2.0)))
	assert.True(t, fullQuarter.TupMul(p).Equal(tuple.Point(0.0, 0.0, 1.0)))
}

// The inverse of an x-rotation rotates in the opposite direction
func TestInverseRotationX(t *testing.T) {
	// Given
	p := tuple.Point(0.0, 1.0, 0.0)
	halfQuarter := matrix.RotationX(math.Pi / 4.0)
	inv := halfQuarter.Inverse()

	// Then
	assert.True(t, inv.TupMul(p).Equal(tuple.Point(0.0, math.Sqrt(2.0)/2.0, -math.Sqrt(2.0)/2.0)))
}

// Rotating a point around the y axis
func TestRotationY(t *testing.T) {
	// Given
	p := tuple.Point(0.0, 0.0, 1.0)
	halfQuarter := matrix.RotationY(math.Pi / 4.0)
	fullQuarter := matrix.RotationY(math.Pi / 2.0)

	// Then
	assert.True(t, halfQuarter.TupMul(p).Equal(tuple.Point(math.Sqrt(2.0)/2.0, 0.0, math.Sqrt(2.0)/2.0)))
	assert.True(t, fullQuarter.TupMul(p).Equal(tuple.Point(1.0, 0.0, 0.0)))
}

// Rotating a point around the z axis
func TestRotationZ(t *testing.T) {
	// Given
	p := tuple.Point(0.0, 1.0, 0.0)
	halfQuarter := matrix.RotationZ(math.Pi / 4.0)
	fullQuarter := matrix.RotationZ(math.Pi / 2.0)

	// Then
	assert.True(t, halfQuarter.TupMul(p).Equal(tuple.Point(-math.Sqrt(2.0)/2.0, math.Sqrt(2.0)/2.0, 0.0)))
	assert.True(t, fullQuarter.TupMul(p).Equal(tuple.Point(-1.0, 0.0, 0.0)))
}

// A shearing transformation
func TestShearing(t *testing.T) {
	tests := []struct {
		Name      string
		Transform []float64
		Point     tuple.Tuple
		Expected  tuple.Tuple
	}{
		{
			Name:      "x in proportion to y",
			Transform: []float64{1.0, 0.0, 0.0, 0.0, 0.0, 0.0},
			Point:     tuple.Point(2.0, 3.0, 4.0),
			Expected:  tuple.Point(5.0, 3.0, 4.0),
		},

		{
			Name:      "x in proportion to z",
			Transform: []float64{0.0, 1.0, 0.0, 0.0, 0.0, 0.0},
			Point:     tuple.Point(2.0, 3.0, 4.0),
			Expected:  tuple.Point(6.0, 3.0, 4.0),
		},

		{
			Name:      "y in proportion to x",
			Transform: []float64{0.0, 0.0, 1.0, 0.0, 0.0, 0.0},
			Point:     tuple.Point(2.0, 3.0, 4.0),
			Expected:  tuple.Point(2.0, 5.0, 4.0),
		},

		{
			Name:      "y in proportion to z",
			Transform: []float64{0.0, 0.0, 0.0, 1.0, 0.0, 0.0},
			Point:     tuple.Point(2.0, 3.0, 4.0),
			Expected:  tuple.Point(2.0, 7.0, 4.0),
		},

		{
			Name:      "z in proportion to x",
			Transform: []float64{0.0, 0.0, 0.0, 0.0, 1.0, 0.0},
			Point:     tuple.Point(2.0, 3.0, 4.0),
			Expected:  tuple.Point(2.0, 3.0, 6.0),
		},

		{
			Name:      "z in proportion to y",
			Transform: []float64{0.0, 0.0, 0.0, 0.0, 0.0, 1.0},
			Point:     tuple.Point(2.0, 3.0, 4.0),
			Expected:  tuple.Point(2.0, 3.0, 7.0),
		},
	}

	for _, test := range tests {
		t.Run("A shearing transformation moves "+test.Name, func(t *testing.T) {
			// Given
			transform := matrix.Shearing(
				test.Transform[0],
				test.Transform[1],
				test.Transform[2],
				test.Transform[3],
				test.Transform[4],
				test.Transform[5],
			)

			p := test.Point

			// Then
			assert.True(t, transform.TupMul(p).Equal(test.Expected))
		})
	}
}

// Individual transformations are applied in sequence
func TestTransformSequence(t *testing.T) {
	// Given
	p := tuple.Point(1.0, 0.0, 1.0)
	a := matrix.RotationX(math.Pi / 2.0)
	b := matrix.Scaling(5.0, 5.0, 5.0)
	c := matrix.Translation(10.0, 5.0, 7.0)

	// When
	p2 := a.TupMul(p) // apply rotation first

	// Then
	assert.True(t, p2.Equal(tuple.Point(1.0, -1.0, 0.0)))

	// When
	p3 := b.TupMul(p2) // then apply scaling

	// Then
	assert.True(t, p3.Equal(tuple.Point(5.0, -5.0, 0.0)))

	// When
	p4 := c.TupMul(p3) // then apply translation

	// Then
	assert.True(t, p4.Equal(tuple.Point(15.0, 0.0, 7.0)))
}

// Chained transformations must be applied in reverse order
func TestTransformChained(t *testing.T) {
	// Given
	p := tuple.Point(1.0, 0.0, 1.0)
	a := matrix.RotationX(math.Pi / 2.0)
	b := matrix.Scaling(5.0, 5.0, 5.0)
	c := matrix.Translation(10.0, 5.0, 7.0)

	// When
	transform := c.MatMul(b).MatMul(a)

	// Then
	assert.True(t, transform.TupMul(p).Equal(tuple.Point(15.0, 0.0, 7.0)))
}
