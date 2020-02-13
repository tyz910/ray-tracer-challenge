package math

import (
	"math"
)

const epsilon = 0.00001

// Equals approximately compares two floats.
func Equals(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}
