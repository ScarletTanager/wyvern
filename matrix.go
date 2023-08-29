package wyvern

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// A Matrix comprises one or more vectors.  These are passed in as column vectors.
type Matrix[N constraints.Float] struct {
	columns []Vector[N]
}

func sameDimensionCount[N constraints.Float](c []Vector[N]) bool {
	var prevLen int
	for i, v := range c {
		if i > 0 {
			if len(v) != prevLen {
				return false
			}
		}
		prevLen = len(v)
	}

	return true
}

// func NewMatrix(c []Vector) (Matrix, error) {
// 	// All vectors must have the same number of components
// 	if sameDimensionCount(c) {
// 		return Matrix{columns: c}, nil
// 	}
// 	return Matrix{}, errors.New("Vectors have different numbers of components")
// }

func FromRows[N constraints.Float](rows []Vector[N]) (Matrix[N], error) {
	// All vectors must have the same number of components
	if sameDimensionCount(rows) {
		cols := make([]Vector[N], len(rows[0]))
		for ci, _ := range cols {
			cols[ci] = make(Vector[N], len(rows))
			for ri, _ := range rows {
				cols[ci][ri] = rows[ri][ci]
			}
		}

		return Matrix[N]{columns: cols}, nil
	}
	return Matrix[N]{}, errors.New("Vectors have different numbers of components")
}

func FromColumns[N constraints.Float](c []Vector[N]) (Matrix[N], error) {
	// All vectors must have the same number of components
	if sameDimensionCount(c) {
		return Matrix[N]{columns: c}, nil
	}
	return Matrix[N]{}, errors.New("Vectors have different numbers of components")
}

// Rows returns the set of row vectors, top to bottom, constituting the Matrix
func (a Matrix[N]) Rows() []Vector[N] {
	cols := a.Columns()

	if cols == nil {
		return nil
	}

	rows := make([]Vector[N], len(cols[0]))
	for ri, _ := range rows {
		rows[ri] = make(Vector[N], len(cols))

		for ci, c := range cols {
			rows[ri][ci] = c[ri]
		}
	}

	return rows
}

// Columns returns the set of column vectors, left to right, consistituting the Matrix
func (a Matrix[N]) Columns() []Vector[N] {
	return a.columns
}

// Product multiplies two matrices.  a is the matrix on the left, b on the right.
// The matrix returned is a combination of the columns of a and of the rows of b.
func (a Matrix[N]) Product(b Matrix[N]) Matrix[N] {
	return Matrix[N]{}
}

func canBeMultiplied[N constraints.Float](a, b Matrix[N]) bool {
	return true
}
