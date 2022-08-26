package cmath

import (
	"fmt"
	"math"
)

// Vector represent a 3D vector
type Vector struct {
	X float64
	Y float64
	Z float64
}

// Norm return the Euclidean norm of a 3D vector
func (v Vector) Norm() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2))
}

// Add return the addition of current vector with the vector w
func (v Vector) Add(w Vector) Vector {
	return Vector{v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}

// Sub return the subtraction of a current vector to a vector w
func (v Vector) Sub(w Vector) Vector {
	return Vector{v.X - w.X, v.Y - w.Y, v.Z - w.Z}
}

// RealProd return the product of current vector with the real number r
func (v Vector) RealProd(r float64) Vector {
	return Vector{v.X * r, v.Y * r, v.Z * r}
}

// DotProd return the dot product of current vector with the vector w
func (v Vector) DotProd(w Vector) float64 {
	return v.X*w.X + v.Y*w.Y + v.Z*w.Z
}

// CrossProd return the cross product of current vector with the vector w
func (v Vector) CrossProd(w Vector) Vector {
	return Vector{
		v.Y*w.Z - v.Z*w.Y,
		v.Z*w.X - v.X*w.Z,
		v.X*w.Y - v.Y*w.X,
	}
}

// String format Vector printing
func (v Vector) String() string {
	return fmt.Sprintf("(%3.2fi + %3.2fj + %3.2fk)", v.X, v.Y, v.Z)
}
