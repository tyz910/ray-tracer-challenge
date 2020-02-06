package matrix

import "math"

// Translation creates a new translation matrix.
// This transformation matrix used to move an object along given axes.
func Translation(x, y, z float64) Matrix {
	return New(4, 4, []float64{
		1.0, 0.0, 0.0, x,
		0.0, 1.0, 0.0, y,
		0.0, 0.0, 1.0, z,
		0.0, 0.0, 0.0, 1.0,
	})
}

// Translation creates a new scaling matrix.
// This transformation matrix used to alter size of an object along given axes.
func Scaling(x, y, z float64) Matrix {
	return New(4, 4, []float64{
		x, 0.0, 0.0, 0.0,
		0.0, y, 0.0, 0.0,
		0.0, 0.0, z, 0.0,
		0.0, 0.0, 0.0, 1.0,
	})
}

// RotationX creates a new rotation matrix for x axis.
// This transformation matrix used to rotate an object clockwise around x axis by r radians.
func RotationX(r float64) Matrix {
	cos := math.Cos(r)
	sin := math.Sin(r)

	return New(4, 4, []float64{
		1.0, 0.0, 0.0, 0.0,
		0.0, cos, -sin, 0.0,
		0.0, sin, cos, 0.0,
		0.0, 0.0, 0.0, 1.0,
	})
}

// RotationY creates a new rotation matrix for y axis.
// This transformation matrix used to rotate an object clockwise around y axis by r radians.
func RotationY(r float64) Matrix {
	cos := math.Cos(r)
	sin := math.Sin(r)

	return New(4, 4, []float64{
		cos, 0.0, sin, 0.0,
		0.0, 1.0, 0.0, 0.0,
		-sin, 0.0, cos, 0.0,
		0.0, 0.0, 0.0, 1.0,
	})
}

// RotationZ creates a new rotation matrix for z axis.
// This transformation matrix used to rotate an object clockwise around z axis by r radians.
func RotationZ(r float64) Matrix {
	cos := math.Cos(r)
	sin := math.Sin(r)

	return New(4, 4, []float64{
		cos, -sin, 0.0, 0.0,
		sin, cos, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	})
}

// Shearing creates a new shearing matrix.
// This transformation matrix used to slant the shape of an object. It changes each component of an object in proportion to the other two components.
// So the x component changes in proportion to y and z, y changes in proportion to x and z, and z changes in proportion to x and y.
func Shearing(xy, xz, yx, yz, zx, zy float64) Matrix {
	return New(4, 4, []float64{
		1.0, xy, xz, 0.0,
		yx, 1.0, yz, 0.0,
		zx, zy, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	})
}

// Transform applies transformations in sequence.
func Transform(transformations ...Matrix) Matrix {
	transform := Identity()
	for i := len(transformations) - 1; i >= 0; i-- {
		transform = transform.MatMul(transformations[i])
	}

	return transform
}
