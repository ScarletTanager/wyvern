package wyvern

import "errors"

// A Matrix comprises one or more vectors.  These are passed in as column vectors.
type Matrix struct {
	Vectors []Vector
}

func NewMatrix(c []Vector) (Matrix, error) {
	// All vectors must have the same number of components
	var prevLen int
	for i, v := range c {
		if i > 0 {
			if len(v) != prevLen {
				return Matrix{}, errors.New("Vectors have different numbers of components")
			}
		}
		prevLen = len(v)
	}

	return Matrix{Vectors: c}, nil
}

func FromRows(c []Vector) (Matrix, error) {
	return Matrix{}, nil
}

func FromColumns(c []Vector) (Matrix, error) {
	return Matrix{}, nil
}

// Rows returns the set of row vectors, top to bottom, constituting the Matrix
func (a Matrix) Rows() []Vector {
	cols := a.Columns()
	rows := make([]Vector, len(cols[0]))
	for ri, _ := range rows {
		rows[ri] = make(Vector, len(cols))

		for ci, c := range cols {
			rows[ri][ci] = c[ri]
		}
	}

	return rows
}

// Columns returns the set of column vectors, left to right, consistituting the Matrix
func (a Matrix) Columns() []Vector {
	return a.Vectors
}

// Product multiplies two matrices.  a is the matrix on the left, b on the right.
// The matrix returned is a combination of the columns of a and of the rows of b.
func (a Matrix) Product(b Matrix) Matrix {
	return Matrix{}
}

func canBeMultiplied(a, b Matrix) bool {
	return true
}
