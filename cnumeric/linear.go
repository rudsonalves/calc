package cnumeric

import (
	"fmt"
	"math"
)

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

func sumX(x []float64) float64 {
	s := 0.
	for i := 0; i < len(x); i++ {
		s += x[i]
	}

	return s
}

func sumX2(x []float64) float64 {
	s := 0.
	for i := 0; i < len(x); i++ {
		s += math.Pow(x[i], 2)
	}

	return s
}

func sumXY(x, y []float64) float64 {
	s := 0.
	for i := 0; i < len(x); i++ {
		s += x[i] * y[i]
	}

	return s
}

func ArithmeticMean(x ...float64) float64 {
	s := 0.
	for _, xi := range x {
		s += xi
	}

	return s / float64(len(x))
}

func QuadraticMean(x ...float64) float64 {
	s := 0.
	for _, xi := range x {
		s += math.Pow(xi, 2)
	}

	return math.Sqrt(s / float64(len(x)))
}

func MidRange(x ...float64) float64 {
	return (Max(x...) + Min(x...)) / 2.
}

func Max(x ...float64) float64 {
	xmax := x[0]
	for _, xi := range x {
		if xmax < xi {
			xmax = xi
		}
	}

	return xmax
}

func Min(x ...float64) float64 {
	xmin := x[0]
	for _, xi := range x {
		if xmin > xi {
			xmin = xi
		}
	}

	return xmin
}
