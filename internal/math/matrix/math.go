package matrix

import (
	"fmt"

	"github.com/tyz910/ray-tracer-challenge/internal/math"
	"github.com/tyz910/ray-tracer-challenge/internal/math/tuple"
)

// Equal approximately compares two matrices.
func (m Matrix) Equal(m2 Matrix) bool {
	if (m.rows != m2.rows) || (m.columns != m2.columns) {
		return false
	}

	for i := range m.values {
		if !math.Equals(m.values[i], m2.values[i]) {
			return false
		}
	}

	return true
}

// MatMul multiplies the matrix by a matrix.
func (m Matrix) MatMul(m2 Matrix) Matrix {
	if m.columns != m2.rows {
		panic(fmt.Sprintf("invalid operand shapes for multiplication: %dx%d * %dx%d", m.rows, m.columns, m2.rows, m2.columns))
	}

	result := New(m.rows, m2.columns, nil)
	for row := 0; row < result.rows; row++ {
		for col := 0; col < result.columns; col++ {
			val := 0.0
			for i := 0; i < m.columns; i++ {
				val += m.Value(row, i) * m2.Value(i, col)
			}

			result.SetValue(row, col, val)
		}
	}

	return result
}

// TupMul multiplies the matrix by a tuple.
func (m Matrix) TupMul(t tuple.Tuple) tuple.Tuple {
	return m.MatMul(FromTuple(t)).ToTuple()
}

// Transpose transposes the matrix by turning it's rows into columns and it's columns into rows.
func (m Matrix) Transpose() Matrix {
	result := New(m.columns, m.rows, nil)
	for row := 0; row < result.rows; row++ {
		for col := 0; col < result.columns; col++ {
			result.SetValue(row, col, m.Value(col, row))
		}
	}

	return result
}

// Determinant returns the determinant of the matrix.
func (m Matrix) Determinant() float64 {
	if m.rows != m.columns {
		panic(fmt.Sprintf("invalid matrix shape %dx%d for determinant calculation", m.rows, m.columns))
	}

	if m.rows == 1 {
		return m.Value(0, 0)
	} else if m.rows == 2 {
		return m.Value(0, 0)*m.Value(1, 1) - m.Value(0, 1)*m.Value(1, 0)
	}

	det := 0.0
	for col := 0; col < m.columns; col++ {
		det += m.Value(0, col) * m.Cofactor(0, col)
	}

	return det
}

// Sub extracts submatrix by excluding given row and column from the matrix.
func (m Matrix) Sub(row, column int) Matrix {
	result := New(m.rows-1, m.columns-1, nil)
	resultRow := 0
	resultCol := 0

	for r := 0; r < m.rows; r++ {
		if r == row {
			continue
		} else if r > row {
			resultRow = r - 1
		} else {
			resultRow = r
		}

		for c := 0; c < m.columns; c++ {
			if c == column {
				continue
			} else if c > column {
				resultCol = c - 1
			} else {
				resultCol = c
			}

			result.SetValue(resultRow, resultCol, m.Value(r, c))
		}
	}

	return result
}

// Minor returns the minor of the matrix. It's the determinant of the submatrix.
func (m Matrix) Minor(row, column int) float64 {
	return m.Sub(row, column).Determinant()
}

// Cofactor returns the cofactor of the matrix. It's the positive or negative minor, depending on the given row and column.
func (m Matrix) Cofactor(row, column int) float64 {
	if (row+column)%2 != 0 {
		return -m.Minor(row, column)
	}

	return m.Minor(row, column)
}

// IsInvertible tests the matrix for invertibility.
func (m Matrix) IsInvertible() bool {
	return m.Determinant() != 0
}

// Inverse returns the inverse of the matrix.
func (m Matrix) Inverse() Matrix {
	if !m.IsInvertible() {
		panic("matrix is not invertible")
	}

	result := New(m.rows, m.columns, nil)
	det := m.Determinant()

	for row := 0; row < m.rows; row++ {
		for col := 0; col < m.columns; col++ {
			result.SetValue(col, row, m.Cofactor(row, col)/det)
		}
	}

	return result
}
