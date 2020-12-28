package suvec

import (
	"log"
)

type Type int

const (
	Matrix64 Type = iota
	Box
	Point3d
	ErrorMatrix
)

type ErrorCode int

const (
	MatchingDimensions = iota
	NotSquareMatrix
	NotImplemented
	NotAScalarType
)

type Mat struct {
	typ        Type
	rows, cols int
	data       []float64
}

/** Iehandle Internal error handler
	"opration" failed with op1, and op2 (optional)
**/
func ehandle(e ErrorCode, op1, op2 *Mat, operation string, hint string) *Mat {
	switch e {
	case MatchingDimensions:
		log.Println("Error matching dimensions ", operation)
	case NotSquareMatrix:
		log.Println("Matrix not square ", operation)
	}

	return NewErrorMatrix()
}

/** New Create a new Vector of float32
	len Length, its a one-row, len-columns by default
	vals initial values

	By default its zero
	If vals are give they are filled from left to the right
	if vals is larger than len then the remeainign values are ignored
	Values are initialized to 0
**/
func New(rows, cols int, typ Type) *Mat {

	m := new(Mat)
	m.typ = typ
	m.rows = rows
	m.cols = cols

	switch m.typ {
	case Matrix64:
		m.data = make([]float64, rows*cols)
	case Box:
		m.data = make([]float64, 6)
	case Point3d:
		m.data = make([]float64, 3)
	}

	m.Zero()

	return m
}

/** NewMatrix is just short for New(n, m, Matrix64)
**/
func NewMatrix(r, c int) *Mat {
	return New(r, c, Matrix64)
}

/** NewVector creates a row-Vector with n elements
Short for : New(1, n, Matrix64).Init(vals)
**/
func NewVector(len int, vals ...float64) *Mat {
	m := New(1, len, Matrix64)
	for i, v := range vals {
		if i < m.rows*m.cols {
			m.data[i] = v
		}
	}
	return m
}

func NewScalar(val float64) *Mat {
	m := New(1, 1, Matrix64)
	m.Set(0, 0, val)
	return m
}

/** Error Matrix returns a valid Matrix, but can not be used for further computeation **/
func NewErrorMatrix() *Mat {
	m := New(0, 0, ErrorMatrix)
	return m
}

/** NewBox creates a 3x2 vector for a geometrical reprentation of a box */
func NewBox(x1, y1, z1, x2, y2, z2 float64) *Mat {
	m := New(3, 2, Box)
	m.data[0] = x1
	m.data[1] = y1
	m.data[2] = z1
	m.data[3] = x2
	m.data[4] = y2
	m.data[5] = z2
	return m
}

/** Init initializes the values of the matrix in row-order
**/
func (m *Mat) Init(vals ...float64) *Mat {
	m.Zero()

	for i, v := range vals {
		if i < m.rows*m.cols {
			m.data[i] = v
		}
	}
	return m
}

/** Duplicate creates a new Instance of the same type and copies the values
**/
func Duplicate(val *Mat) *Mat {
	m := new(Mat)
	m.typ = val.typ
	m.rows = val.rows
	m.cols = val.cols
	m.data = make([]float64, val.rows*val.cols)
	for i := 0; i < m.rows*m.cols; i++ {
		m.data[i] = val.data[i]
	}
	return m
}

/** Create a copy **/
func (m *Mat) Clone() *Mat {
	return Duplicate(m)
}

func (m *Mat) SetTo(val *Mat) *Mat {
	if m.rows == val.rows && m.cols == val.cols {
		for i := 0; i < m.rows*m.cols; i++ {
			m.data[i] = val.data[i]
		}
	} else {
		// errorhandling: not Implemented yet
	}
	return m
}

func (m *Mat) Zero() *Mat {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.data[i*m.cols+j] = 0
		}
	}
	return m
}

func Zeros(rows, cols int) *Mat {
	z := New(rows, cols, Matrix64)
	for i := 0; i < z.rows; i++ {
		for j := 0; j < z.cols; j++ {
			z.data[i*z.cols+j] = 0
		}
	}
	return z
}

func (m *Mat) Ones() *Mat {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.data[i*m.cols+j] = 1
		}
	}
	return m
}

func Ones(rows, cols int) *Mat {
	z := New(rows, cols, Matrix64)
	for i := 0; i < z.rows; i++ {
		for j := 0; j < z.cols; j++ {
			z.data[i*z.cols+j] = 1
		}
	}
	return z
}

func (m *Mat) Identity() *Mat {

	if m.rows != m.cols {
		return ehandle(NotSquareMatrix, m, nil, "Identity", "")
	}

	m.Zero()
	for i := 0; i < m.rows; i++ {
		m.data[i*m.cols+i] = 1
	}
	return m
}

func (m *Mat) Get(r, c int) float64 {
	return m.data[r*m.cols+c]
}

func (m *Mat) Cols() int {
	return m.cols
}

func (m *Mat) Rows() int {
	return m.cols
}

func (m *Mat) Set(r, c int, val float64) *Mat {
	m.data[r*m.cols+c] = val
	return m
}

func (m *Mat) IsScalar() bool {
	if m.rows == 1 && m.cols == 1 {
		return true
	}
	return false
}

func (m *Mat) IsVector() bool {
	if m.rows == 1 && m.cols > 1 {
		return true
	}
	if m.cols == 1 && m.rows > 1 {
		return true
	}
	return false
}

func (m *Mat) IsRowVector() bool {
	if m.rows == 1 && m.cols > 1 {
		return true
	}
	return false
}

func (m *Mat) IsColumnVector() bool {
	if m.cols == 1 && m.rows > 1 {
		return true
	}
	return false
}

func (m *Mat) IsMatrix() bool {
	if m.rows > 1 && m.cols > 1 {
		return true
	}
	return false
}

func (m *Mat) IsSameSize(val *Mat) bool {
	if m.rows != val.rows || m.cols != val.cols {
		return false
	}
	return true
}

/** T Transposes a Matrix **/
func (m *Mat) T() *Mat {
	res := New(m.cols, m.rows, m.typ)

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			res.data[j*m.rows+i] = m.data[i*m.cols+j]
		}
	}

	return res
}

/** Length of largest array dimension **/
func (m *Mat) Length() *Mat {

	log.Panic("Dont use this anymore")

	return nil
}

/** Float32 returns the value of a Scalar as float32 Datatype.
  - The input must be a scalar
 **/
func (m *Mat) Float32() float32 {
	if m.IsScalar() {
		return float32(m.data[0])
	}

	ehandle(NotAScalarType, m, nil, "Float32", "")

	return 0
}

/** Float64 returns the value of a Scalar as float64 Datatype.
  - The input must be a scalar
 **/
func (m *Mat) Float64() float64 {

	if m.IsScalar() {
		return float64(m.data[0])
	}

	ehandle(NotAScalarType, m, nil, "Float64", "")

	return 0
}
