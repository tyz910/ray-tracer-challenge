package tuple_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	mathUtil "github.com/tyz910/ray-tracer-challenge/internal/math"
	"github.com/tyz910/ray-tracer-challenge/internal/math/tuple"
)

// Adding two tuples
func TestAddTuples(t *testing.T) {
	// Given
	a1 := tuple.New(3.0, -2.0, 5.0, 1.0)
	a2 := tuple.New(-2.0, 3.0, 1.0, 0.0)

	// Then
	assert.True(t, a1.Add(a2).Equal(tuple.New(1.0, 1.0, 6.0, 1.0)))
}

// Subtracting two points
func TestSubtractPoints(t *testing.T) {
	// Given
	p1 := tuple.Point(3.0, 2.0, 1.0)
	p2 := tuple.Point(5.0, 6.0, 7.0)

	// Then
	assert.True(t, p1.Sub(p2).Equal(tuple.Vector(-2.0, -4.0, -6.0)))
}

// Subtracting a vector from a point
func TestSubtractVectorFromPoint(t *testing.T) {
	// Given
	p := tuple.Point(3.0, 2.0, 1.0)
	v := tuple.Vector(5.0, 6.0, 7.0)

	// Then
	assert.True(t, p.Sub(v).Equal(tuple.Point(-2.0, -4.0, -6.0)))
}

// Subtracting two vectors
func TestSubtractVectors(t *testing.T) {
	// Given
	v1 := tuple.Vector(3.0, 2.0, 1.0)
	v2 := tuple.Vector(5.0, 6.0, 7.0)

	// Then
	assert.True(t, v1.Sub(v2).Equal(tuple.Vector(-2.0, -4.0, -6.0)))
}

// Subtracting a vector from the zero vector
func TestSubtractVectorFromZero(t *testing.T) {
	// Given
	zero := tuple.Vector(0.0, 0.0, 0.0)
	v := tuple.Vector(1.0, -2.0, 3.0)

	// Then
	assert.True(t, zero.Sub(v).Equal(tuple.Vector(-1.0, 2.0, -3.0)))
}

// Negating a tuple
func TestNegate(t *testing.T) {
	// Given
	a := tuple.New(1.0, -2.0, 3.0, -4.0)

	// Then
	assert.True(t, a.Negate().Equal(tuple.New(-1.0, 2.0, -3.0, 4.0)))
}

// Multiplying a tuple by a scalar
func TestMulScalar(t *testing.T) {
	// Given
	a := tuple.New(1.0, -2.0, 3.0, -4.0)

	// Then
	assert.True(t, a.Mul(3.5).Equal(tuple.New(3.5, -7.0, 10.5, -14.0)))
}

// Multiplying a tuple by a fraction
func TestMulFraction(t *testing.T) {
	// Given
	a := tuple.New(1.0, -2.0, 3.0, -4.0)

	// Then
	assert.True(t, a.Mul(0.5).Equal(tuple.New(0.5, -1.0, 1.5, -2.0)))
}

// Dividing a tuple by a scalar
func TestDiv(t *testing.T) {
	// Given
	a := tuple.New(1.0, -2.0, 3.0, -4.0)

	// Then
	assert.True(t, a.Div(2.0).Equal(tuple.New(0.5, -1.0, 1.5, -2.0)))
}

// Computing the magnitude of vector
func TestMagnitude(t *testing.T) {
	tests := []struct {
		V tuple.Tuple
		M float64
	}{
		{V: tuple.Vector(1.0, 0.0, 0.0), M: 1.0},
		{V: tuple.Vector(0.0, 1.0, 0.0), M: 1.0},
		{V: tuple.Vector(0.0, 0.0, 1.0), M: 1.0},
		{V: tuple.Vector(1.0, 2.0, 3.0), M: math.Sqrt(14)},
		{V: tuple.Vector(-1.0, -2.0, -3.0), M: math.Sqrt(14)},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Computing the magnitude of %v", test.V), func(t *testing.T) {
			// Given
			v := test.V

			// Then
			assert.True(t, mathUtil.Equals(v.Magnitude(), test.M))
		})
	}
}

// Normalizing vector
func TestNormalize(t *testing.T) {
	tests := []struct {
		V tuple.Tuple
		N tuple.Tuple
	}{
		{V: tuple.Vector(4.0, 0.0, 0.0), N: tuple.Vector(1.0, 0.0, 0.0)},
		{V: tuple.Vector(1.0, 2.0, 3.0), N: tuple.Vector(0.26726, 0.53452, 0.80178)},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Normalizing %v", test.V), func(t *testing.T) {
			// Given
			v := test.V

			// Then
			assert.True(t, v.Normalize().Equal(test.N))
		})
	}
}

// The magnitude of a normalized vector
func TestMagnitudeNormalized(t *testing.T) {
	// Given
	v := tuple.Vector(1.0, 2.0, 3.0)

	// When
	norm := v.Normalize()

	// Then
	assert.True(t, mathUtil.Equals(norm.Magnitude(), 1.0))
}

// The dot product of two tuples
func TestDotProduct(t *testing.T) {
	// Given
	a := tuple.Vector(1.0, 2.0, 3.0)
	b := tuple.Vector(2.0, 3.0, 4.0)

	// Then
	assert.True(t, mathUtil.Equals(a.Dot(b), 20.0))
}

// The cross product of two vectors
func TestCrossProduct(t *testing.T) {
	// Given
	a := tuple.Vector(1.0, 2.0, 3.0)
	b := tuple.Vector(2.0, 3.0, 4.0)

	// Then
	assert.True(t, a.Cross(b).Equal(tuple.Vector(-1.0, 2.0, -1.0)))
	assert.True(t, b.Cross(a).Equal(tuple.Vector(1.0, -2.0, 1.0)))
}

// Reflecting a vector
func TestReflect(t *testing.T) {
	tests := []struct {
		Name string
		V    tuple.Tuple
		N    tuple.Tuple
		R    tuple.Tuple
	}{
		{
			Name: "Reflecting a vector approaching at 45°",
			V:    tuple.Vector(1.0, -1.0, 0.0),
			N:    tuple.Vector(0.0, 1.0, 0.0),
			R:    tuple.Vector(1.0, 1.0, 0.0),
		},

		{
			Name: "Reflecting a vector off a slanted surface",
			V:    tuple.Vector(0.0, -1.0, 0.0),
			N:    tuple.Vector(math.Sqrt(2.0)/2.0, math.Sqrt(2.0)/2.0, 0.0),
			R:    tuple.Vector(1.0, 0.0, 0.0),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			// Given
			v := test.V
			n := test.N

			// When
			r := v.Reflect(n)

			// Then
			assert.True(t, r.Equal(test.R))
		})
	}
}
