package suvec

import (
	"math"
)

func (m *Mat) Add(val *Mat) *Mat {

	if !m.IsSameSize(val) {
		ehandle(MatchingDimensions, m, val, "Add", "")
		return nil
	}

	res := Duplicate(m)

	for i := 0; i < m.rows*m.cols; i++ {
		res.data[i] = res.data[i] + val.data[i]
	}

	return res
}

func (m *Mat) Sub(val *Mat) *Mat {

	if !m.IsSameSize(val) {
		ehandle(MatchingDimensions, m, val, "Sub", "")
		return nil
	}

	res := Duplicate(m)
	for i := 0; i < m.rows*m.cols; i++ {
		res.data[i] = res.data[i] - val.data[i]
	}

	return res
}

func (m *Mat) Mul(val *Mat) *Mat {

	// Scaling
	if m.IsScalar() {
		res := m.Clone()
		for i := 0; i < m.rows*m.cols; i++ {
			res.data[i] = res.data[i] * val.data[0]
		}
		return res
	}

	if m.cols != val.rows {
		return ehandle(MatchingDimensions, m, val, "Mul", "")
	}

	res := New(m.rows, val.cols, Matrix64)
	for i1 := 0; i1 < m.rows; i1++ {

		for j2 := 0; j2 < val.cols; j2++ {
			sum := 0.0
			for j1 := 0; j1 < m.cols; j1++ {
				sum = sum + m.Get(i1, j1)*val.Get(j1, j2)
			}
			res.Set(i1, j2, sum)
		}
	}

	return res
}

/** Norm2 (p=2)-Norm of Vector or the 2-Norm of a matrix
	- Input must be a Vector
**/
func (m *Mat) Norm2() *Mat {

	if m.IsScalar() {
		return m.Clone()
	}

	if m.IsVector() {

		sum := 0.0
		for j := 0; j < m.cols*m.rows; j++ {
			sum = sum + m.data[j]*m.data[j]
		}
		rowsqrt := math.Sqrt(sum)

		return NewScalar(rowsqrt)
	}

	return ehandle(NotImplemented, m, nil, "Norm", "")
}
