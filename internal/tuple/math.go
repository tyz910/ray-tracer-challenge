package tuple

import (
	"math"

	"github.com/tyz910/ray-tracer-challenge/internal/util"
)

// Equal approximately compares two tuples.
func (t Tuple) Equal(t2 Tuple) bool {
	return util.Equals(t.x, t2.x) && util.Equals(t.y, t2.y) && util.Equals(t.z, t2.z) && util.Equals(t.w, t2.w)
}

// Add adds two tuples by adding the corresponding components.
func (t Tuple) Add(o Tuple) Tuple {
	return New(t.x+o.x, t.y+o.y, t.z+o.z, t.w+o.w)
}

// Sub subtracts two tuples by subtracting corresponding components.
func (t Tuple) Sub(o Tuple) Tuple {
	return New(t.x-o.x, t.y-o.y, t.z-o.z, t.w-o.w)
}

// Negate negates each component of the tuple.
func (t Tuple) Negate() Tuple {
	return New(-t.x, -t.y, -t.z, -t.w)
}

// Mul multiplies each component of the tuple by the scalar.
func (t Tuple) Mul(n float64) Tuple {
	return New(n*t.x, n*t.y, n*t.z, n*t.w)
}

// Div divides each component of the tuple by the scalar.
func (t Tuple) Div(n float64) Tuple {
	return New(t.x/n, t.y/n, t.z/n, t.w/n)
}

// Magnitude returns the length of the vector.
func (t Tuple) Magnitude() float64 {
	if !(t.IsVector()) {
		panic("invalid operand type - must be a vector")
	}

	return math.Sqrt(t.x*t.x + t.y*t.y + t.z*t.z + t.w*t.w)
}

// Normalize normalizes the vector to a unit vector with magnitude=1.
func (t Tuple) Normalize() Tuple {
	if !(t.IsVector()) {
		panic("invalid operand type - must be a vector")
	}

	return t.Div(t.Magnitude())
}

// Dot returns the dot product of the two vectors.
func (t Tuple) Dot(t2 Tuple) float64 {
	if !(t.IsVector() && t2.IsVector()) {
		panic("invalid operand types - must be a vectors")
	}

	return t.x*t2.x + t.y*t2.y + t.z*t2.z + t.w*t2.w
}

// Cross returns the cross product of the two vectors.
func (t Tuple) Cross(t2 Tuple) Tuple {
	if !(t.IsVector() && t2.IsVector()) {
		panic("invalid operand types - must be a vectors")
	}

	return Vector(t.y*t2.z-t.z*t2.y, t.z*t2.x-t.x*t2.z, t.x*t2.y-t.y*t2.x)
}
