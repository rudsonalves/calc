package cmath

import (
	"math"
	"testing"
)

func TestX(t *testing.T) {
	v := Vector{1, 2, -3}
	if v.X() != 1. {
		t.Errorf("incorrect result: expected %1.3f, got %1.3f", v.X(), 1.)
	}
}

func TestY(t *testing.T) {
	v := Vector{1, 2, -3}
	if v.Y() != 2. {
		t.Errorf("incorrect result: expected %1.3f, got %1.3f", v.Y(), 2.)
	}
}

func TestZ(t *testing.T) {
	v := Vector{1, 2, -3}
	if v.Z() != -3. {
		t.Errorf("incorrect result: expected %1.3f, got %1.3f", v.Z(), -3.)
	}
}

func TestNorm(t *testing.T) {
	v := Vector{1, 2, -3}
	result := v.Norm()
	ans := math.Sqrt(14)
	if result != ans {
		t.Errorf("incorrect result: expected %1.3f, got %1.3f", ans, result)
	}
}

func TestIsEqual(t *testing.T) {
	vec := Vector{1, 2, 3}
	other := Vector{1, 2, 3}
	if !vec.IsEqual(other) {
		t.Errorf("incorrect result: expected %v, got %v", true, false)
	}
	other[2] += 1.
	if vec.IsEqual(other) {
		t.Errorf("incorrect result: expected %v, got %v", false, true)
	}
}

func TestAdd(t *testing.T) {
	v0 := Vector{1, -2, 3}
	v1 := Vector{2, 4, -6}
	v2 := Vector{-3, 2, -1}
	result := v0.Add(v1.Add(v2))
	ans := Vector{0, 4, -4}
	if !result.IsEqual(ans) {
		t.Errorf("incorrect result: expected %v, got %v", ans, result)
	}
}

func TestSub(t *testing.T) {
	v0 := Vector{-1, 2, -3}
	v1 := Vector{2, 4, -6}
	v2 := Vector{-3, 2, -1}
	result := v0.Sub(v1.Sub(v2))
	ans := Vector{-1 - (2 + 3), 2 - (4 - 2), -3 - (-6 + 1)}
	if !result.IsEqual(ans) {
		t.Errorf("incorrect result: expected %v, got %v", ans, result)
	}
}

func TestRealProd(t *testing.T) {
	v0 := Vector{-1, 2, -3}
	v1 := Vector{2, 4, -6}
	result := v0.RealProd(2).Add(v1.RealProd(-3))
	ans := Vector{-2 - 6, 4 - 12, -6 + 18}
	if !result.IsEqual(ans) {
		t.Errorf("incorrect result: expected %v, got %v", ans, result)
	}
}

func TestDotProd(t *testing.T) {
	v0 := Vector{-1, 2, -3}
	v1 := Vector{2, 4, -6}
	result := v0.DotProd(v1)
	ans := float64(-1*2 + 2*4 + 3*6)
	if result != ans {
		t.Errorf("incorrect result: expected %v, got %v", ans, result)
	}
}

func TestCrossProd(t *testing.T) {
	v0 := Vector{1, 2, 3}
	v1 := Vector{1, 2, -1}
	result := v0.CrossProd(v1)
	ans := Vector{-8, 4, 0}
	if !result.IsEqual(ans) {
		t.Errorf("incorrect result: expected %v, got %v", ans, result)
	}
}

func TestString(t *testing.T) {
	v0 := Vector{1, 2, 3}
	result := v0.String()
	ans := "(1.00i + 2.00j + 3.00k)"
	if ans != result {
		t.Errorf("incorrect result: expected %v, got %v", ans, result)
	}
}
