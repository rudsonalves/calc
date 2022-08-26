package cmath

import (
	"fmt"
)

// Matrix represent a rowsxcolumns matrix
type Matrix struct {
	elements [][]float64
	rows     int
	columns  int
}

// StartMatrix start a matrix with r rows and c columns with e elements
func (m *Matrix) StartMatrix(r int, c int, e ...float64) error {
	if r*c != len(e) {
		return fmt.Errorf("rows (%d) x columns (%d) is different from the number of matrix elements (%d)", r, c, len(e))
	}

	for i := 0; i < r; i++ {
		row := []float64{}
		for j := 0; j < c; j++ {
			row = append(row, e[i*c+j])
		}
		m.elements = append(m.elements, row)
	}
	m.rows = r
	m.columns = c

	return nil
}

// ZerosMatrix start a zeros matrix of r rows and c columns
func (m *Matrix) ZerosMatrix(r, c int) {
	m.rows = r
	m.columns = c

	for i := 0; i < r; i++ {
		row := make([]float64, c)
		m.elements = append(m.elements, row)
	}
}

// String format Matrix printing
func (m Matrix) String() string {
	str := "["
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.columns; c++ {
			str += fmt.Sprintf("%+3.2f ", m.elements[r][c])
		}
		str = str[:len(str)-1] + "]\n["
	}
	str = str[:len(str)-1]

	return str
}

// GetElement return the [i,j] Matrix element
func (m Matrix) GetElement(i, j int) float64 {
	return m.elements[i][j]
}

// SetElement set the [i,j] element with v value
func (m *Matrix) SetElement(i, j int, v float64) {
	m.elements[i][j] = v
}

// Product return the product the current Matrix and w Matrix
func (m Matrix) Product(w Matrix) (Matrix, error) {
	if m.columns != w.rows {
		return Matrix{}, fmt.Errorf("in A*B the A.columns must be equal to B.rows")
	}

	z := Matrix{}
	z.ZerosMatrix(m.rows, w.columns)

	for i := 0; i < z.rows; i++ {
		for j := 0; j < z.columns; j++ {
			for k := 0; k < z.rows; k++ {
				z.elements[i][j] += m.elements[i][k] * w.elements[k][j]
			}
		}
	}

	return z, nil
}

// Compare return true if current Matrix is equal to w Matrix
func (m Matrix) Compare(w Matrix) bool {
	if m.columns != w.columns || m.rows != w.rows {
		return false
	}

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			if m.elements[i][j] != w.elements[i][j] {
				return false
			}
		}
	}

	return true
}

// RealProduct returns the product between the current Matrix
// and a real number c
func (m Matrix) RealProduct(c float64) Matrix {
	w := Matrix{}
	w.ZerosMatrix(m.rows, m.columns)

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			w.elements[i][j] = m.elements[i][j] * c
		}
	}

	return w
}

// Add return the addition between the current Matrix
// and a w Matrix
func (m Matrix) Add(w Matrix) (Matrix, error) {
	if m.rows != w.rows || m.columns != w.columns {
		return Matrix{}, fmt.Errorf("it is not possible to add matrices of different sizes")
	}
	z := Matrix{}
	z.ZerosMatrix(m.rows, m.columns)

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			z.elements[i][j] = m.elements[i][j] + w.elements[i][j]
		}
	}

	return z, nil
}

// Sub return the subtraction between the current Matrix
// and a w Matrix
func (m Matrix) Sub(w Matrix) (Matrix, error) {
	if m.rows != w.rows || m.columns != w.columns {
		return Matrix{}, fmt.Errorf("it is not possible to subtract matrices of different sizes")
	}
	z := Matrix{}
	z.ZerosMatrix(m.rows, m.columns)

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			z.elements[i][j] = m.elements[i][j] - w.elements[i][j]
		}
	}

	return z, nil
}

// Det2 return the determinant of the current Matrix 2x2
func (m Matrix) Det2() (float64, error) {
	if m.rows != 2 || m.columns != 2 {
		return 0., fmt.Errorf("is not a 2x2 matrix")
	}

	e := m.elements
	d := e[0][0]*e[1][1] - e[0][1]*e[1][0]

	return d, nil
}

// Det3 return the determinant of the current Matrix 3x3
func (m Matrix) Det3() (float64, error) {
	if m.rows != 3 || m.columns != 3 {
		return 0., fmt.Errorf("is not a 3x3 matrix")
	}
	e := m.elements
	d := e[0][0]*(e[1][1]*e[2][2]-e[1][2]*e[2][1]) -
		e[0][1]*(e[1][0]*e[2][2]-e[1][2]*e[2][0]) +
		e[0][2]*(e[1][0]*e[2][1]-e[1][1]*e[2][0])

	return d, nil
}
