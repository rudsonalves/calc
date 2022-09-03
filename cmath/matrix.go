/*
Copyright 2022 The Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.
*/

/*
Package cmath package has types, attributes and methods for calculating
matrices and vectors. This package contains two files with
matrix and vector definitions.
*/
package cmath

import (
	"fmt"
)

/*
Matrix declares a matrix type that allows you to do addition,
subtraction, and product operations according to matrix
manipulation rules.

Array elements can be initialized by the StratMatrix and
ZerosMatrix methods.
*/
type Matrix struct {
	elems [][]float64
	rows  int
	cols  int
}

// GetElement returns the matrix elements of the row r
// and column c. An error is generated if the requested
// element exceeds the matrix size.
func (m Matrix) GetElement(r, c int) (float64, error) {
	if r >= 0 && r < m.rows && c >= 0 && c < m.cols {
		return m.elems[r][c], nil
	}
	return 0., fmt.Errorf("there is not element [%d][%d] in this matrix", r, c)
}

// SetElement changes the matrix elements of row r and column
// c with value v. An error is generated if the requested
// element exceeds the matrix size.
func (m *Matrix) SetElement(r, c int, v float64) error {
	if r >= 0 && r < m.rows && c >= 0 && c < m.cols {
		m.elems[r][c] = v
		return nil
	}
	return fmt.Errorf("there is not element [%d][%d] in this matrix", r, c)
}

// StartMatrix starts a matrix with r rows and c columns with the
// elements passed by list e. An error is generated if the number
// of elements does not match the product r*c.
func StartMatrix(r int, c int, e ...float64) (Matrix, error) {
	if r*c != len(e) {
		return Matrix{}, fmt.Errorf("rows (%d) x columns (%d) is different from the number of matrix elements (%d)", r, c, len(e))
	}

	m := newMatrix(r, c)

	for i := 0; i < r; i++ {
		row := []float64{}
		for j := 0; j < c; j++ {
			row = append(row, e[i*c+j])
		}
		m.elems = append(m.elems, row)
	}

	return m, nil
}

// StartZerosMatrix starts a matrix with zero elements with r rows
// and c columns.
func StartZerosMatrix(r, c int) Matrix {
	m := newMatrix(r, c)

	for i := 0; i < r; i++ {
		row := make([]float64, c)
		m.elems = append(m.elems, row)
	}

	return m
}

// newMatrix return a new zeros elements Matrix with r rows
// and c cols.
func newMatrix(r, c int) Matrix {
	var m Matrix
	m.elems = [][]float64{}
	m.rows = r
	m.cols = c

	return m
}

// SetMatrix reset current matrix elements with list e. An
// error is generated if the number of elements passed does
// not match the product m.rows*m.cols.
func (m *Matrix) SetMatrix(e ...float64) error {
	if m.rows*m.cols != len(e) {
		return fmt.Errorf("rows (%d) x columns (%d) is different from the number of matrix elements (%d)", m.rows, m.cols, len(e))
	}

	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			m.elems[r][c] = e[r*m.cols+c]
		}
	}

	return nil
}

// SetZeros reset the current matrix elements.
func (m *Matrix) SetZeros() {
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			m.elems[r][c] = 0.
		}
	}
}

// Product returns the product between the current matrix, m, and
// the matrix other. An error is generated if the number of columns of
// m is different from the number of rows of w.
func (m Matrix) Product(other Matrix) (Matrix, error) {
	if m.cols != other.rows {
		return Matrix{}, fmt.Errorf("in A*B the A.columns must be equal to B.rows")
	}

	result := StartZerosMatrix(m.rows, other.cols)

	for r := 0; r < result.rows; r++ {
		for c := 0; c < result.cols; c++ {
			for k := 0; k < result.rows; k++ {
				result.elems[r][c] += m.elems[r][k] * other.elems[k][c]
			}
		}
	}

	return result, nil
}

// IsEqual compares the current matrix, m, with the matrix other
// and returns true if they are equal
func (m Matrix) IsEqual(other Matrix) bool {
	if m.cols != other.cols || m.rows != other.rows {
		return false
	}

	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			if m.elems[r][c] != other.elems[r][c] {
				return false
			}
		}
	}

	return true
}

// RealProduct retorna uma matriz produto da Matrizes
// corrente, m, por uma constante real c.
func (m Matrix) RealProduct(real float64) Matrix {
	result := StartZerosMatrix(m.rows, m.cols)

	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			result.elems[r][c] = m.elems[r][c] * real
		}
	}

	return result
}

// Add returns the addition of the current matrix, m, by the
// other matrix. An error is generated if the arrays have
// different sizes.
func (m Matrix) Add(other Matrix) (Matrix, error) {
	if m.rows != other.rows || m.cols != other.cols {
		return Matrix{}, fmt.Errorf("it is not possible to add matrices of different sizes")
	}
	result := StartZerosMatrix(m.rows, m.cols)

	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			result.elems[r][c] = m.elems[r][c] + other.elems[r][c]
		}
	}

	return result, nil
}

// Sub returns the subtraction of the current matrix, m, by the
// other matrix. An error is generated if the arrays have different
// sizes.
func (m Matrix) Sub(other Matrix) (Matrix, error) {
	if m.rows != other.rows || m.cols != other.cols {
		return Matrix{}, fmt.Errorf("it is not possible to subtract matrices of different sizes")
	}
	result := StartZerosMatrix(m.rows, m.cols)

	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			result.elems[r][c] = m.elems[r][c] - other.elems[r][c]
		}
	}

	return result, nil
}

// Det2 returns the determinant of a 2x2 matrix. Returns an error
// if the array is not 2x2.
func (m Matrix) Det2() (float64, error) {
	if m.rows != 2 || m.cols != 2 {
		return 0., fmt.Errorf("is not a 2x2 matrix")
	}

	e := m.elems
	det := e[0][0]*e[1][1] - e[0][1]*e[1][0]

	return det, nil
}

// Det3 returns the determinant of a 3x3 matrix. Returns an
// error if the array is not 3x3.
func (m Matrix) Det3() (float64, error) {
	if m.rows != 3 || m.cols != 3 {
		return 0., fmt.Errorf("is not a 3x3 matrix")
	}
	e := m.elems
	det := e[0][0]*(e[1][1]*e[2][2]-e[1][2]*e[2][1]) -
		e[0][1]*(e[1][0]*e[2][2]-e[1][2]*e[2][0]) +
		e[0][2]*(e[1][0]*e[2][1]-e[1][1]*e[2][0])

	return det, nil
}

// String creates formatted output for the array and makes the
// Matrix part of the types that satisfy the fmt.Stringer interface.
func (m Matrix) String() string {
	if len(m.elems) > 0 {
		str := "["
		for r := 0; r < m.rows; r++ {
			for c := 0; c < m.cols; c++ {
				str += fmt.Sprintf("%+3.2f ", m.elems[r][c])
			}
			str = str[:len(str)-1] + "]\n["
		}
		str = str[:len(str)-1]

		return str
	}
	return "[]\n"
}
