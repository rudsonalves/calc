package cmath

import (
	"math"
	"testing"
)

func TestGetElement(t *testing.T) {
	m, _ := StartMatrix(3, 2, 1, 2, 3, 4, 5, 6)
	v10, err := m.GetElement(1, 0)
	if err != nil {
		t.Error("incorrect result: expected error nil.")
	}
	if v10 != 3. {
		t.Errorf("incorrect result: expected v10 = 3., found %v", v10)
	}

	v12, err := m.GetElement(1, 2)
	if err == nil {
		t.Error("incorrect result: expected error.")
	}
	if v12 != 0. {
		t.Errorf("incorrect result: expected v12 = 0., found %v", v12)
	}

	v21, err := m.GetElement(2, 1)
	if err != nil {
		t.Error("incorrect result: expected error nil.")
	}
	if v21 != 6. {
		t.Errorf("incorrect result: expected v21 = 6., found %v", v21)
	}
}

func TestSetElement(t *testing.T) {
	m, _ := StartMatrix(3, 2, 1, 2, 3, 4, 5, 6)

	err := m.SetElement(1, 0, 0)
	if err != nil {
		t.Error("incorrect result: expected error nil.")
	}

	err = m.SetElement(1, 2, 0)
	if err == nil {
		t.Error("incorrect result: expected error.")
	}

	err = m.SetElement(2, 1, 0)
	if err != nil {
		t.Error("incorrect result: expected error nil.")
	}

	if m.elems[1][0] != 0. || m.elems[2][1] != 0. {
		t.Error("incorrect result: expected v10, v21 = 0., 0.")
	}
}

func TestStartMatrix(t *testing.T) {
	_, err := StartMatrix(4, 2, 1, 2, 3, 4, 5, 6)
	if err == nil {
		t.Errorf("incorrect result: expected error c*r, got nil")
	}

	m, err := StartMatrix(2, 3, 1, 2, 3, 4, 5, 6)
	if err != nil && m.cols == 3 && m.rows == 2 {
		t.Error("incorrect result: expected error nil, got error or matrix 3x2.")
	}
	elems := [2][3]float64{{1, 2, 3}, {4, 5, 6}}
	for r := 0; r < 2; r++ {
		for c := 0; c < 3; c++ {
			e, _ := m.GetElement(r, c)
			if elems[r][c] != e {
				t.Errorf("incorrect result: expected m[%d][%d] = %f, find %f.", r, c, elems[r][c], e)
			}
			t.Logf("m[%d][%d] = %f, find %f.", r, c, elems[r][c], e)
		}
	}
}

func TestStartZerosMatrix(t *testing.T) {
	m := StartZerosMatrix(2, 2)

	sum := 0.
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			v, err := m.GetElement(r, c)
			if err != nil {
				t.Error("incorrect result: expected error nil.")
			}
			sum += math.Abs(v)
		}
	}
	if sum != 0. {
		t.Error("incorrect result: expected sum == 0.")
	}
	t.Log("All elements are zeros.")
}

func TestSetMatrix(t *testing.T) {
	m, _ := StartMatrix(3, 2, 1, 2, 3, 4, 5, 6)

	err := m.SetMatrix(1, 2, 3, 4)
	if err == nil {
		t.Errorf("incorrect result: expected error c*r, got nil")
	}

	err = m.SetMatrix(6, 5, 4, 3, 2, 1)
	if err != nil && m.cols == 3 && m.rows == 2 {
		t.Error("incorrect result: expected error nil, got error or matrix 3x2.")
	}
	t.Log("Number of elements test passed.")
	mt := [3][2]float64{{6, 5}, {4, 3}, {2, 1}}
	for r := 0; r < 3; r++ {
		for c := 0; c < 2; c++ {
			vmt := mt[r][c]
			vm := m.elems[r][c]
			if vmt != vm {
				t.Errorf("incorrect result: expected m[%d][%d] = %f, find %f.", r, c, vmt, vm)
			}
		}
	}
	t.Log("Elements check passed.")
}

func TestSetZeros(t *testing.T) {
	m, _ := StartMatrix(2, 3, 1, 2, 3, 4, 5, 6)
	m.SetZeros()

	sum := 0.
	for r := 0; r < 2; r++ {
		for c := 0; c < 3; c++ {
			sum += math.Abs(m.elems[r][c])
		}
	}
	if sum != 0. {
		t.Error("the elements of matrix isn't zeros.")
	}
}

func TestProduct(t *testing.T) {
	m0, _ := StartMatrix(2, 2, 3, 2, 5, -1)
	m1, _ := StartMatrix(3, 2, 6, 4, -2, 0, 7, 1)
	if _, err := m0.Product(m1); err == nil {
		t.Error("incorrect result: expected err != nil, because m0.cols != m1.rows.")
	}
	t.Log("Return nil != nil, because m0.cols != m1.rows. It is ok")

	m1, _ = StartMatrix(2, 3, 6, 4, -2, 0, 7, 1)
	m2, _ := StartMatrix(2, 3, 18, 26, -4, 30, 13, -11)
	m3, _ := m0.Product(m1)
	if !m2.IsEqual(m3) {
		t.Errorf("incorrect result: expected \n%v, got\n%v.", m2, m3)
	}
	t.Logf("Return expected product \n%v.", m3)
}

func TestIsEqualMatrix(t *testing.T) {
	m0, _ := StartMatrix(3, 3, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	m1, _ := StartMatrix(3, 4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	if m0.IsEqual(m1) {
		t.Error("incorrect result: expected false. Diferents by sizes.")
	}
	m1, _ = StartMatrix(3, 3, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	if !m0.IsEqual(m1) {
		t.Error("incorrect result: expected true.")
	}
	m1.SetElement(2, 2, 10)
	if m0.IsEqual(m1) {
		t.Error("incorrect result: expected false.")
	}
}

func TestRealProductMatrix(t *testing.T) {
	m0, _ := StartMatrix(3, 3, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	m1, _ := StartMatrix(3, 3, 2, 4, 6, 8, 10, 12, 14, 16, 18)
	m3 := m0.RealProduct(2)
	if !m1.IsEqual(m3) {
		t.Error("incorrect result: expected true.")
	}
}

func TestAddMatrix(t *testing.T) {
	m0, _ := StartMatrix(2, 2, 1, 2, 3, 4)
	m1, _ := StartMatrix(2, 1, 5, 6)

	_, err := m0.Add(m1)
	if err == nil {
		t.Error("incorrect result: expected error.")
	}

	m1, _ = StartMatrix(2, 2, 5, 6, 7, 8)
	m3, err := m0.Add(m1)
	m4, _ := StartMatrix(2, 2, 6, 8, 10, 12)
	if err != nil {
		t.Error("incorrect result: expected err is nil.")
	}
	if !m4.IsEqual(m3) {
		t.Error("incorrect result: expected is equal (true).")
	}
}

func TestSubMatrix(t *testing.T) {
	m0, _ := StartMatrix(2, 2, 1, 2, 3, 4)
	m1, _ := StartMatrix(2, 1, 5, 6)

	_, err := m0.Sub(m1)
	if err == nil {
		t.Error("incorrect result: expected err is nil.")
	}

	m1, _ = StartMatrix(2, 2, 5, 6, 7, 8)
	m3, err := m0.Sub(m1)
	m4, _ := StartMatrix(2, 2, -4, -4, -4, -4)
	if err != nil {
		t.Error("incorrect result: expected err is nil.")
	}
	if !m4.IsEqual(m3) {
		t.Error("incorrect result: expected is equal (true).")
	}
}

func TestDet2(t *testing.T) {
	m0, _ := StartMatrix(2, 1, 1, 2)
	_, err := m0.Det2()
	if err == nil {
		t.Error("incorrect result: expected err != nil.")
	}

	m0, _ = StartMatrix(2, 2, 1, 2, 3, 4)
	d, err := m0.Det2()
	if err != nil {
		t.Error("incorrect result: expected err is nil.")
	}
	if d != -2. {
		t.Errorf("incorrect result: expected d = %1.2f, got %1.2f.", -2., d)
	}
}

func TestDet3(t *testing.T) {
	m0, _ := StartMatrix(2, 2, 1, 2, 3, 4)
	_, err := m0.Det3()
	if err == nil {
		t.Error("incorrect result: expected err != nil.")
	}

	m0, _ = StartMatrix(3, 3, 1, 9, 5, 3, 7, 8, 10, 4, 2)
	d, err := m0.Det3()
	if err != nil {
		t.Error("incorrect result: expected err is nil.")
	}
	if d != 358. {
		t.Errorf("incorrect result: expected d = %1.2f, got %1.2f.", 358., d)
	}
}

func TestStringMatrix(t *testing.T) {
	m, _ := StartMatrix(2, 2, 1, 2, 3, 4)
	str := m.String()
	if str != "[+1.00 +2.00]\n[+3.00 +4.00]\n" {
		t.Error("incorrect result: expected err is nil.")
	}

	m = Matrix{}
	str = m.String()
	if str != "[]\n" {
		t.Error("incorrect result: expected err is nil.")
	}
}
