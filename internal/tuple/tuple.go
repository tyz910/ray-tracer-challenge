package tuple

import (
	"fmt"
)

const (
	pointW  = 1.0
	vectorW = 0.0
)

// Tuple contains four components: three for spatial coordinates and one to distinguish a point from a vector.
type Tuple struct {
	x, y, z, w float64
}

// New creates new tuple.
func New(x, y, z, w float64) Tuple {
	return Tuple{x, y, z, w}
}

// Point creates new tuple with w=1.
func Point(x, y, z float64) Tuple {
	return New(x, y, z, pointW)
}

// Vector creates new tuple with w=0.
func Vector(x, y, z float64) Tuple {
	return New(x, y, z, vectorW)
}

// X returns the x-axis coordinate.
func (t Tuple) X() float64 {
	return t.x
}

// Y returns the y-axis coordinate.
func (t Tuple) Y() float64 {
	return t.y
}

// Z returns the z-axis coordinate.
func (t Tuple) Z() float64 {
	return t.z
}

// W returns the vector/point flag.
func (t Tuple) W() float64 {
	return t.w
}

// IsPoint checks whether tuple is a point.
func (t Tuple) IsPoint() bool {
	return t.w == pointW
}

// IsVector checks whether tuple is a vector.
func (t Tuple) IsVector() bool {
	return t.w == vectorW
}

func (t Tuple) String() string {
	if t.IsVector() {
		return fmt.Sprintf("vector(%.1f, %.1f, %.1f)", t.x, t.y, t.z)
	}

	if t.IsPoint() {
		return fmt.Sprintf("point(%.1f, %.1f, %.1f)", t.x, t.y, t.z)
	}

	return fmt.Sprintf("tuple(%.1f, %.1f, %.1f, %.1f)", t.x, t.y, t.z, t.w)
}
