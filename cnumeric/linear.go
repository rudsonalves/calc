/*
Copyright 2022 The Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.
*/

// Pacote cnumeric contém funções para a calculo
// de raízes, funções estatísticas e ajustes de curvas.
package cnumeric

import (
	"fmt"
	"math"
)

// LinearAdj returns the coefficients a1 and a2 of the linear
// fit y = a1 + a2*x for the data sets (x1, y1), (x2, y2),
// (x3, y3), ... passed in the input. An error is generated if
// the passed dataset number is not even.
func LinearAdj(v ...float64) (float64, float64, error) {
	if len(v)%2 != 0 {
		return 0, 0, fmt.Errorf("v must be x1, y1, x2, y2, ... with an even number of elements")
	}
	x := []float64{}
	y := []float64{}

	for i := 0; i < len(v); i++ {
		if i%2 == 0 {
			x = append(x, v[i])
		} else {
			y = append(y, v[i])
		}
	}

	N := float64(len(x))
	sX := sumX(x)
	sX2 := sumX2(x)
	sY := sumX(y)
	sXY := sumXY(x, y)

	d := N*sX2 - math.Pow(sX, 2)
	a1 := (sX2*sY - sX*sXY) / d
	a2 := (N*sXY - sX*sY) / d

	return a1, a2, nil
}

// sumX returns the sum of all elements passed.
func sumX(x []float64) float64 {
	s := 0.
	for i := 0; i < len(x); i++ {
		s += x[i]
	}

	return s
}

// sumX2 returns the sum of xi² from all elements passed.
func sumX2(x []float64) float64 {
	s := 0.
	for i := 0; i < len(x); i++ {
		s += math.Pow(x[i], 2)
	}

	return s
}

// sumXY returns the sum of xi*yi from all passed elements.
func sumXY(x, y []float64) float64 {
	s := 0.
	for i := 0; i < len(x); i++ {
		s += x[i] * y[i]
	}

	return s
}

// ArithmeticMean returns the arithmetic mean of the x...
func ArithmeticMean(x ...float64) float64 {
	s := 0.
	for _, xi := range x {
		s += xi
	}

	return s / float64(len(x))
}

// QuadraticMean returns the mean square of the x...
func QuadraticMean(x ...float64) float64 {
	s := 0.
	for _, xi := range x {
		s += math.Pow(xi, 2)
	}

	return math.Sqrt(s / float64(len(x)))
}

// MidRange returns the average value of the values ​​passed.
func MidRange(x ...float64) float64 {
	return (Max(x...) + Min(x...)) / 2.
}

// Max returns the maximum of the values ​​passed.
func Max(x ...float64) float64 {
	xmax := x[0]
	for i := 1; i < len(x); i++ {
		if xmax < x[i] {
			xmax = x[i]
		}
	}

	return xmax
}

// Min returns the minimum of the values ​​passed.
func Min(x ...float64) float64 {
	xmin := x[0]
	for i := 0; i < len(x); i++ {
		if xmin > x[i] {
			xmin = x[i]
		}
	}

	return xmin
}
