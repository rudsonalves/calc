/*
Copyright 2022 The Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.
*/

package cnumeric

import (
	"math"
)

// RootsPoly2 returns the real or complex roots of the second
// -order polynomial ax²+bx+c=0
func RootsPoly2(a, b, c float64) (x1, x2 interface{}) {
	a2 := 2. * a
	delta := math.Pow(b, 2) - 4*a*c
	if delta >= 0 {
		// Real roots
		sq := math.Sqrt(delta)
		x1 = (-b + sq) / a2
		x2 = (-b - sq) / a2
		return
	}
	// Complex roots
	x1 = complex(-b/a2, math.Sqrt(-delta)/a2)
	x2 = complex(-b/a2, -math.Sqrt(-delta)/a2)
	return
}
