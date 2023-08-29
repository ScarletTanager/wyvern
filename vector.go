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

func (v Vector[N]) Multiply(f N) {
	for idx, _ := range v {
		v.multiplyComponent(idx, f)
	}
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
