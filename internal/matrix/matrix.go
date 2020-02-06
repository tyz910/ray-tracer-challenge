package matrix

import (
	"fmt"

	"github.com/tyz910/ray-tracer-challenge/internal/tuple"
)

// Matrix is a grid of numbers.
type Matrix struct {
	rows, columns int
	values        []float64
}

// New creates new matrix.
func New(rows, columns int, values []float64) Matrix {
	if rows < 1 || columns < 1 {
		panic(fmt.Sprintf("invalid matrix shape %dx%d", rows, columns))
	}

	expectedLen := rows * columns
	if values == nil {
		values = make([]float64, expectedLen)
	} else if len(values) != expectedLen {
		panic(fmt.Sprintf("invalid number of elements for matrix %dx%d. expected: %d", rows, columns, expectedLen))
	}

	return Matrix{
		rows:    rows,
		columns: columns,
		values:  values,
	}
}

// FromTuple creates new 4x1 matrix from tuple.
func FromTuple(t tuple.Tuple) Matrix {
	return New(4, 1, []float64{
		t.X(),
		t.Y(),
		t.Z(),
		t.W(),
	})
}

// Identity creates new identity matrix. Multiplying a matrix by the identity matrix returns the original matrix.
func Identity() Matrix {
	return New(4, 4, []float64{
		1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	})
}

// Rows returns the number of rows in the matrix.
func (m Matrix) Rows() int {
	return m.rows
}

// Columns returns the number of columns in the matrix.
func (m Matrix) Columns() int {
	return m.columns
}

// index returns index for given row and column in values slice.
func (m Matrix) index(row, column int) int {
	if (row < 0) || (row > m.rows-1) || (column < 0) || (column > m.columns-1) {
		panic(fmt.Sprintf("index (%d, %d) out of range for matrix %dx%d", row, column, m.rows, m.columns))
	}

	return row*m.columns + column
}

// Value returns the value at position (row, column).
func (m Matrix) Value(row, column int) float64 {
	return m.values[m.index(row, column)]
}

// SetValue sets the value at position (row, column).
func (m Matrix) SetValue(row, column int, value float64) {
	m.values[m.index(row, column)] = value
}

// ToTuple converts 4x1 matrix to tuple.
func (m Matrix) ToTuple() tuple.Tuple {
	return tuple.New(
		m.Value(0, 0),
		m.Value(1, 0),
		m.Value(2, 0),
		m.Value(3, 0),
	)
}
