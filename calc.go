package main

import (
	"fmt"
	"math"
	"mycodes/calc/cmath"
	"mycodes/calc/cnumeric"
)

func main() {
	m1 := cmath.Matrix{}
	m1.StartMatrix(2, 2, 3., 2., 5., -1.)
	m2 := cmath.Matrix{}
	m2.StartMatrix(2, 3, 6., 4, -2., 0., 7., 1.)

	fmt.Println("m1:")
	fmt.Println(m1)

	fmt.Println("m2:")
	fmt.Println(m2)

	fmt.Println("m1+m2:")
	z, err := m1.Add(m2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(z)
	}

	fmt.Println("\nm1*m2:")
	z, _ = m1.Product(m2)
	fmt.Println(z)

	v1 := cmath.Vector{X: 2., Y: 3., Z: 4.}
	v2 := cmath.Vector{X: 5., Y: 5., Z: 2.}

	fmt.Println("v1:", v1, " v2:", v2)

	x1, x2 := cnumeric.RootsPoly2(2, 3, -5)
	fmt.Printf("\n2x² + 3x - 5 = 0:\nx1 = %+1.2f x2 = %+1.2f\n", x1, x2)

	fmt.Println("\nLinear adjust of points {1., 1.2, 1.5, 1.3, 2., 2.3}")
	a1, a2, _ := cnumeric.LinearAdj(1., 1.2, 1.5, 1.3, 2., 2.3)
	fmt.Printf("b = %+1.2f m = %+1.2f\n", a1, a2)

	fmt.Println("\nLinear adjust of points {.23, -.54, -.3, -.54, .04, -.57}")
	a1, a2, _ = cnumeric.LinearAdj(.23, -.54, -.3, -.54, .04, -.57)
	fmt.Printf("b = %+1.2f m = %+1.2f\n", a1, a2)

	fmt.Println("\nLinear adjust of points {-.35, .2, .15, -.5, .23, .54, .35, .7}")
	a1, a2, _ = cnumeric.LinearAdj(-.35, .2, .15, -.5, .23, .54, .35, .7)
	fmt.Printf("b = %+1.2f m = %+1.2f\n", a1, a2)

	fmt.Println("\n", math.Sqrt(2))
}
