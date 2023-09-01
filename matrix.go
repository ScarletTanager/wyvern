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
	for i, _ := range c {
		if i > 0 {
			if !c[i].sameDimension(c[i-1]) {
				return false
			}
		}
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

// Row returns the specified row as a Vector.  Returns nil if the index is out of bounds.
// The returned Vector is a copy and can be modified without affecting the
// source Matrix.
func (a Matrix[N]) Row(rowIndex int) (Vector[N], error) {
	if !a.isValidRowIndex(rowIndex) {
		return nil, errors.New("Row index out of bounds")
	}

	row := make(Vector[N], len(a.columns))
	for ci, c := range a.columns {
		row[ci] = c[rowIndex]
	}

	return row, nil
}

// Rows returns the set of row vectors, top to bottom, constituting the Matrix.
// Both Rows() and Columns() return **copies** of the component vectors - modifying
// the returned slice does not affect the original Matrix.
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

// Column returns the specified column as a Vector.  Returns nil if the index is out of bounds.
// The returned Vector is a copy and can be modified without affecting the
// source Matrix.
func (a Matrix[N]) Column(columnIndex int) (Vector[N], error) {
	if !a.isValidColumnIndex(columnIndex) {
		return nil, errors.New("Column index out of bounds")
	}

	col := make(Vector[N], len(a.columns[columnIndex]))
	for compIdx, val := range a.columns[columnIndex] {
		col[compIdx] = val
	}

	return col, nil
}

// Columns returns the set of column vectors, left to right, consistituting the Matrix.
// Both Rows() and Columns() return **copies** of the component vectors - modifying
// the returned slice does not affect the original Matrix.
func (a Matrix[N]) Columns() []Vector[N] {
	cols := make([]Vector[N], len(a.columns))
	for ci, c := range a.columns {
		cols[ci] = make(Vector[N], len(c))
		for compIdx, val := range c {
			cols[ci][compIdx] = val
		}
	}
	return cols
}

// ReplaceRow replaces the row at the specified index with the supplied Vector.
// An error is returned if either (a) the index is out of bounds or (b) the
// Vector has the wrong dimension.
func (a Matrix[N]) ReplaceRow(rowIndex int, src Vector[N]) error {
	if !a.isValidRowIndex(rowIndex) {
		return errors.New("Row index out of bounds")
	}

	if len(src) != len(a.columns) {
		return errors.New("Vector has wrong dimension")
	}

	for ci, c := range a.columns {
		c[rowIndex] = src[ci]
	}

	return nil
}

// ReplaceColumn replaces the column at the specified index with the supplied Vector.
// An error is returned if either (a) the index is out of bounds or (b) the
// Vector as the wrong demension.
func (a Matrix[N]) ReplaceColumn(columnIndex int, src Vector[N]) error {
	if !a.isValidColumnIndex(columnIndex) {
		return errors.New("Column index out of bounds")
	}

	if !src.sameDimension(a.columns[0]) {
		return errors.New("Vector has wrong dimension")
	}

	a.columns[columnIndex] = src
	return nil
}

// MultiplyRow multiplies the specified row by the given factor.
// Returns an error if the row index is out of range.
func (a Matrix[N]) MultiplyRow(rowIndex int, factor N) error {
	if !a.isValidRowIndex(rowIndex) {
		return errors.New("Row index out of range")
	}

	for _, col := range a.columns {
		orig := col[rowIndex]
		col[rowIndex] = orig * factor
	}

	return nil
}

// MultiplyColumn multiplies the specified column by the given factor.
// Returns an error if the colun index is out of range.
func (a Matrix[N]) MultiplyColumn(columnIndex int, factor N) error {
	if !a.isValidColumnIndex(columnIndex) {
		return errors.New("Row index out of range")
	}
	a.columns[columnIndex].Multiply(factor)
	return nil
}

func (a Matrix[N]) isValidRowIndex(index int) bool {
	return a.isValidIndex(index, true)
}

func (a Matrix[N]) isValidColumnIndex(index int) bool {
	return a.isValidIndex(index, false)
}

func (a Matrix[N]) isValidIndex(index int, isRowIndex bool) bool {
	if index < 0 {
		return false
	}

	if isRowIndex {
		if index >= len(a.columns[0]) {
			return false
		}
	} else if index >= len(a.columns) {
		return false
	}

	return true
}

func (a Matrix[N]) eliminate(srcIdx, destIdx int) {

}

func (a Matrix[N]) combineRows(srcIdx, destIdx int, factor N) Matrix[N] {
	result := Matrix[N]{
		columns: make([]Vector[N], len(a.columns)),
	}

	for ci, srcColumn := range a.columns {
		result.columns[ci] = make(Vector[N], len(srcColumn))
		for compIdx, val := range srcColumn {
			if compIdx != destIdx {
				result.columns[ci][compIdx] = val
			} else {
				result.columns[ci][compIdx] = val + (factor * srcColumn[srcIdx])
			}
		}
	}

	return result
}

// Product multiplies two matrices.  a is the matrix on the left, b on the right.
// The matrix returned is a combination of the columns of a and of the rows of b.
// Not yet implemented
func (a Matrix[N]) Product(b Matrix[N]) (Matrix[N], error) {
	return Matrix[N]{}, nil
}

func canBeMultiplied[N constraints.Float](a, b Matrix[N]) bool {
	return true
}
