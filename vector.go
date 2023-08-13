package wyvern

import "math"

type Vector []int

func (v Vector) DotProduct(w Vector) int {
	var (
		dp int
	)

	for i, c := range v {
		dp += (c * w[i])
	}

	return dp
}

func (v Vector) Magnitude() float64 {
	var sumOfSquaredComponents float64
	for _, c := range v {
		sumOfSquaredComponents += float64(c * c)
	}

	return math.Sqrt(sumOfSquaredComponents)
}

func (v Vector) sameSpace(w Vector) bool {
	return len(v) == len(w)
}

func (v Vector) Angle(w Vector) float64 {
	return math.Acos(float64(v.DotProduct(w)) / (v.Magnitude() * w.Magnitude()))
}
