package suvec

func (m *Mat) Max() *Mat {

	r := New(1, m.cols, Matrix64)

	for j := 0; j < m.cols; j++ {
		maxValue := m.data[0*m.cols+j]
		for i := 1; i < m.rows; i++ {
			if m.data[i*m.cols+j] > maxValue {
				maxValue = m.data[i*m.cols+j]
			}
		}
		r.Set(0, j, maxValue)
	}

	return r
}
