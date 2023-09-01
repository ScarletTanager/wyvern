package wyvern

import (
	"errors"
	"math"

	"golang.org/x/exp/constraints"
)

type Vector[N constraints.Float] []N

func (v Vector[N]) DotProduct(w Vector[N]) N {
	var (
		dp N
	)

	for i, c := range v {
		dp += (c * w[i])
	}

	return dp
}

func (v Vector[N]) Magnitude() float64 {
	var sumOfSquaredComponents float64
	for _, c := range v {
		sumOfSquaredComponents += float64(c * c)
	}

	return math.Sqrt(sumOfSquaredComponents)
}

func (v Vector[N]) sameDimension(w Vector[N]) bool {
	return len(v) == len(w)
}

func (v Vector[N]) Angle(w Vector[N]) float64 {
	return math.Acos(float64(v.DotProduct(w)) / (v.Magnitude() * w.Magnitude()))
}

// Multiply multiplies each component of the vector by the specified factor.  It
// returns the vector.
func (v Vector[N]) Multiply(f N) Vector[N] {
	for idx, _ := range v {
		v.multiplyComponent(idx, f)
	}

	return v
}

// Difference returns (v-other) as a new Vector - it does not modify the original Vector.
// If other has fewer dimensions than v, then the components of other are subtracted
// from the corresponding first n components of v, where n is the dimension of other.
// The remaining components are left unchanged (as if other had 0 values for those components).
func (v Vector[N]) Difference(other Vector[N]) Vector[N] {
	d := make(Vector[N], len(v))

	for ci, orig := range v {
		if ci < len(other) {
			d[ci] = orig - other[ci]
		} else {
			d[ci] = orig
		}
	}

	return d
}

func (v Vector[N]) MultiplyComponent(componentIndex int, f N) error {
	if componentIndex < 0 || componentIndex >= len(v) {
		return errors.New("Component index out of range for vector")
	}

	v.multiplyComponent(componentIndex, f)
	return nil
}

func (v Vector[N]) multiplyComponent(componentIndex int, f N) {
	prevVal := v[componentIndex]
	v[componentIndex] = prevVal * f
}
