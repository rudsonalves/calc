/*
Copyright 2022 The Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.
*/

package cmath

import (
	"fmt"
	"math"
)

// Vector returns a three-dimensional vector supporting
// conventional vector operations.
type Vector [3]float64

// X return the x component of Vector
func (v Vector) X() float64 {
	return v[0]
}

// Y return the y component of Vector
func (v Vector) Y() float64 {
	return v[1]
}

// Z return the z component of Vector
func (v Vector) Z() float64 {
	return v[2]
}

// Norm returns the norm (module) of the current vets.
func (v Vector) Norm() float64 {
	return math.Sqrt(math.Pow(v[0], 2) + math.Pow(v[1], 2) + math.Pow(v[2], 2))
}

// Add returns the vector resulting from the sum of the current
// vector, v, by the other vecto.
func (v Vector) Add(other Vector) Vector {
	return Vector{v[0] + other[0], v[1] + other[1], v[2] + other[2]}
}

// Sub returns the vector resulting from the subtraction of the
// current vector, v, by the other vector.
func (v Vector) Sub(other Vector) Vector {
	return Vector{v[0] - other[0], v[1] - other[1], v[2] - other[2]}
}

// IsEqual returns true if the current vector, v, is equal to
// the other vector.
func (v Vector) IsEqual(other Vector) bool {
	if v[0] == other[0] && v[1] == other[1] && v[2] == other[2] {
		return true
	}
	return false
}

// RealProd returns the resultant vector of the product of the
// current vector, v, by the real constant.
func (v Vector) RealProd(real float64) Vector {
	return Vector{v[0] * real, v[1] * real, v[2] * real}
}

// DotProd returns the dot product of the current vector, v, by
// the other vector.
func (v Vector) DotProd(other Vector) float64 {
	return v[0]*other[0] + v[1]*other[1] + v[2]*other[2]
}

// CrossProd returns the cross product between the current vector,
// v, and the other vector.
func (v Vector) CrossProd(other Vector) Vector {
	return Vector{
		v[1]*other[2] - v[2]*other[1],
		v[2]*other[0] - v[0]*other[2],
		v[0]*other[1] - v[1]*other[0],
	}
}

// String creates formatted output for the Matrix and makes the
// Matrix part of the types that satisfy the fmt.Stringer interface.
func (v Vector) String() string {
	return fmt.Sprintf("(%3.2fi + %3.2fj + %3.2fk)", v[0], v[1], v[2])
}
