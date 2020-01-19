package matrix_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tyz910/ray-tracer-challenge/internal/matrix"
	"github.com/tyz910/ray-tracer-challenge/internal/tuple"
)

// Matrix equality with identical matrices
func TestEqual(t *testing.T) {
	// Given
	a := matrix.New(4, 4, []float64{
		1.0, 2.0, 3.0, 4.0,
		5.0, 6.0, 7.0, 8.0,
		9.0, 8.0, 7.0, 6.0,
		5.0, 4.0, 3.0, 2.0,
	})

	b := matrix.New(4, 4, []float64{
		1.0, 2.0, 3.0, 4.0,
		5.0, 6.0, 7.0, 8.0,
		9.0, 8.0, 7.0, 6.0,
		5.0, 4.0, 3.0, 2.0,
	})

	// Then
	assert.True(t, a.Equal(b))
}

// Matrix equality with different matrices
func TestNotEqual(t *testing.T) {
	// Given
	a := matrix.New(4, 4, []float64{
		1.0, 2.0, 3.0, 4.0,
		5.0, 6.0, 7.0, 8.0,
		9.0, 8.0, 7.0, 6.0,
		5.0, 4.0, 3.0, 2.0,
	})

	b := matrix.New(4, 4, []float64{
		2.0, 3.0, 4.0, 5.0,
		6.0, 7.0, 8.0, 9.0,
		8.0, 7.0, 6.0, 5.0,
		4.0, 3.0, 2.0, 1.0,
	})

	// Then
	assert.False(t, a.Equal(b))
}

// Multiplying two matrices
func TestMatMul(t *testing.T) {
	// Given
	a := matrix.New(4, 4, []float64{
		1.0, 2.0, 3.0, 4.0,
		5.0, 6.0, 7.0, 8.0,
		9.0, 8.0, 7.0, 6.0,
		5.0, 4.0, 3.0, 2.0,
	})

	b := matrix.New(4, 4, []float64{
		-2.0, 1.0, 2.0, 3.0,
		3.0, 2.0, 1.0, -1.0,
		4.0, 3.0, 6.0, 5.0,
		1.0, 2.0, 7.0, 8.0,
	})

	// Then
	assert.True(t, a.MatMul(b).Equal(matrix.New(4, 4, []float64{
		20.0, 22.0, 50.0, 48.0,
		44.0, 54.0, 114.0, 108.0,
		40.0, 58.0, 110.0, 102.0,
		16.0, 26.0, 46.0, 42.0,
	})))
}

// A matrix multiplied by a tuple
func TestTupMul(t *testing.T) {
	// Given
	a := matrix.New(4, 4, []float64{
		1.0, 2.0, 3.0, 4.0,
		2.0, 4.0, 4.0, 2.0,
		8.0, 6.0, 4.0, 1.0,
		0.0, 0.0, 0.0, 1.0,
	})

	b := tuple.New(1.0, 2.0, 3.0, 1.0)

	// Then
	assert.True(t, a.TupMul(b).Equal(tuple.New(18.0, 24.0, 33.0, 1.0)))
}

// Multiplying a matrix by the identity matrix
func TestMatMulIdentity(t *testing.T) {
	// Given
	a := matrix.New(4, 4, []float64{
		0.0, 1.0, 2.0, 4.0,
		1.0, 2.0, 4.0, 8.0,
		2.0, 4.0, 8.0, 16.0,
		4.0, 8.0, 16.0, 32.0,
	})

	// Then
	assert.True(t, a.MatMul(matrix.Identity()).Equal(a))
}

// Multiplying the identity matrix by a tuple
func TestTupMulIdentity(t *testing.T) {
	// Given
	a := tuple.New(1, 2, 3, 4)

	// Then
	assert.True(t, matrix.Identity().TupMul(a).Equal(a))
}

// Transposing a matrix
func TestTranspose(t *testing.T) {
	// Given
	a := matrix.New(4, 4, []float64{
		0.0, 9.0, 3.0, 0.0,
		9.0, 8.0, 0.0, 8.0,
		1.0, 8.0, 5.0, 3.0,
		0.0, 0.0, 5.0, 8.0,
	})

	// Then
	assert.True(t, a.Transpose().Equal(matrix.New(4, 4, []float64{
		0.0, 9.0, 1.0, 0.0,
		9.0, 8.0, 8.0, 0.0,
		3.0, 0.0, 5.0, 5.0,
		0.0, 8.0, 3.0, 8.0,
	})))
}

// Transposing the identity matrix
func TestTransposeIdentity(t *testing.T) {
	// Given
	a := matrix.Identity().Transpose()

	// Then
	assert.True(t, a.Equal(matrix.Identity()))
}

// Calculating the determinant of a 2x2 matrix
func TestDeterminant2x2(t *testing.T) {
	// Given
	a := matrix.New(2, 2, []float64{
		1.0, 5.0,
		-3.0, 2.0,
	})

	// Then
	assert.Equal(t, 17.0, a.Determinant())
}

// A submatrix of a 3x3 matrix is a 2x2 matrix
func TestSub3x3(t *testing.T) {
	// Given
	a := matrix.New(3, 3, []float64{
		1.0, 5.0, 0.0,
		-3.0, 2.0, 7.0,
		0.0, 6.0, -3.0,
	})

	// Then
	assert.True(t, a.Sub(0, 2).Equal(matrix.New(2, 2, []float64{
		-3.0, 2.0,
		0.0, 6.0,
	})))
}

// A submatrix of a 4x4 matrix is a 3x3 matrix
func TestSub4x4(t *testing.T) {
	// Given
	a := matrix.New(4, 4, []float64{
		-6.0, 1.0, 1.0, 6.0,
		-8.0, 5.0, 8.0, 6.0,
		-1.0, 0.0, 8.0, 2.0,
		-7.0, 1.0, -1.0, 1.0,
	})

	// Then
	assert.True(t, a.Sub(2, 1).Equal(matrix.New(3, 3, []float64{
		-6.0, 1.0, 6.0,
		-8.0, 8.0, 6.0,
		-7.0, -1.0, 1.0,
	})))
}

// Calculating a minor of a 3x3 matrix
func TestMinor3x3(t *testing.T) {
	// Given
	a := matrix.New(3, 3, []float64{
		3.0, 5.0, 0.0,
		2.0, -1.0, -7.0,
		6.0, -1.0, 5.0,
	})

	b := a.Sub(1, 0)

	// Then
	assert.Equal(t, 25.0, b.Determinant())
	assert.Equal(t, 25.0, a.Minor(1, 0))
}

// Calculating a cofactor of a 3x3 matrix
func TestCofactor3x3(t *testing.T) {
	// Given
	a := matrix.New(3, 3, []float64{
		3.0, 5.0, 0.0,
		2.0, -1.0, -7.0,
		6.0, -1.0, 5.0,
	})

	// Then
	assert.Equal(t, -12.0, a.Minor(0, 0))
	assert.Equal(t, -12.0, a.Cofactor(0, 0))
	assert.Equal(t, 25.0, a.Minor(1, 0))
	assert.Equal(t, -25.0, a.Cofactor(1, 0))
}

// Calculating the determinant of a 3x3 matrix
func TestDeterminant3x3(t *testing.T) {
	// Given
	a := matrix.New(3, 3, []float64{
		1.0, 2.0, 6.0,
		-5.0, 8.0, -4.0,
		2.0, 6.0, 4.0,
	})

	// Then
	assert.Equal(t, 56.0, a.Cofactor(0, 0))
	assert.Equal(t, 12.0, a.Cofactor(0, 1))
	assert.Equal(t, -46.0, a.Cofactor(0, 2))
	assert.Equal(t, -196.0, a.Determinant())
}

// Calculating the determinant of a 4x4 matrix
func TestDeterminant4x4(t *testing.T) {
	// Given
	a := matrix.New(4, 4, []float64{
		-2.0, -8.0, 3.0, 5.0,
		-3.0, 1.0, 7.0, 3.0,
		1.0, 2.0, -9.0, 6.0,
		-6.0, 7.0, 7.0, -9.0,
	})

	// Then
	assert.Equal(t, 690.0, a.Cofactor(0, 0))
	assert.Equal(t, 447.0, a.Cofactor(0, 1))
	assert.Equal(t, 210.0, a.Cofactor(0, 2))
	assert.Equal(t, 51.0, a.Cofactor(0, 3))
	assert.Equal(t, -4071.0, a.Determinant())
}

// Testing an invertible matrix for invertibility
func TestInvertible(t *testing.T) {
	// Given
	a := matrix.New(4, 4, []float64{
		6.0, 4.0, 4.0, 4.0,
		5.0, 5.0, 7.0, 6.0,
		4.0, -9.0, 3.0, -7.0,
		9.0, 1.0, 7.0, -6.0,
	})

	// Then
	assert.Equal(t, -2120.0, a.Determinant())
	assert.True(t, a.IsInvertible())
}

// Testing a noninvertible matrix for invertibility
func TestNonInvertible(t *testing.T) {
	// Given
	a := matrix.New(4, 4, []float64{
		-4.0, 2.0, -2.0, -3.0,
		9.0, 6.0, 2.0, 6.0,
		0.0, -5.0, 1.0, -5.0,
		0.0, 0.0, 0.0, 0.0,
	})

	// Then
	assert.Equal(t, 0.0, a.Determinant())
	assert.False(t, a.IsInvertible())
}

// Calculating the inverse of a matrix
func TestInverse(t *testing.T) {
	tests := []struct {
		A matrix.Matrix
		B matrix.Matrix
	}{
		{
			A: matrix.New(4, 4, []float64{
				-5.0, 2.0, 6.0, -8.0,
				1.0, -5.0, 1.0, 8.0,
				7.0, 7.0, -6.0, -7.0,
				1.0, -3.0, 7.0, 4.0,
			}),
			B: matrix.New(4, 4, []float64{
				0.21805, 0.45113, 0.24060, -0.04511,
				-0.80827, -1.45677, -0.44361, 0.52068,
				-0.07895, -0.22368, -0.05263, 0.19737,
				-0.52256, -0.81391, -0.30075, 0.30639,
			}),
		},

		{
			A: matrix.New(4, 4, []float64{
				8.0, -5.0, 9.0, 2.0,
				7.0, 5.0, 6.0, 1.0,
				-6.0, 0.0, 9.0, 6.0,
				-3.0, 0.0, -9.0, -4.0,
			}),
			B: matrix.New(4, 4, []float64{
				-0.15385, -0.15385, -0.28205, -0.53846,
				-0.07692, 0.12308, 0.02564, 0.03077,
				0.35897, 0.35897, 0.43590, 0.92308,
				-0.69231, -0.69231, -0.76923, -1.92308,
			}),
		},

		{
			A: matrix.New(4, 4, []float64{
				9.0, 3.0, 0.0, 9.0,
				-5.0, -2.0, -6.0, -3.0,
				-4.0, 9.0, 6.0, 4.0,
				-7.0, 6.0, 6.0, 2.0,
			}),
			B: matrix.New(4, 4, []float64{
				-0.04074, -0.07778, 0.14444, -0.22222,
				-0.07778, 0.03333, 0.36667, -0.33333,
				-0.02901, -0.14630, -0.10926, 0.12963,
				0.17778, 0.06667, -0.26667, 0.33333,
			}),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Calculating the inverse of a matrix #%d", i), func(t *testing.T) {
			// Given
			a := test.A

			// Then
			assert.True(t, a.Inverse().Equal(test.B))
		})
	}
}

// Multiplying a product by its inverse
func TestMulInverse(t *testing.T) {
	// Given
	a := matrix.New(4, 4, []float64{
		3.0, -9.0, 7.0, 3.0,
		3.0, -8.0, 2.0, -9.0,
		-4.0, 4.0, 4.0, 1.0,
		-6.0, 5.0, -1.0, 1.0,
	})

	b := matrix.New(4, 4, []float64{
		8, 2, 2, 2,
		3, -1, 7, 0,
		7, 0, 5, 4,
		6, -2, 0, 5,
	})

	c := a.MatMul(b)

	// Then
	assert.True(t, c.MatMul(b.Inverse()).Equal(a))
}
