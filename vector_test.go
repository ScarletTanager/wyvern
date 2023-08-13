package wyvern_test

import (
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ScarletTanager/wyvern"
)

var _ = Describe("Vector", func() {
	var (
		v, w wyvern.Vector
	)

	BeforeEach(func() {
		v = wyvern.Vector{1, 2, 3}
		w = wyvern.Vector{4, 5, 6}
	})

	Describe("Dot", func() {
		It("computes the correct dot product", func() {
			Expect(v.DotProduct(w)).To(Equal(32))
		})
	})

	Describe("Magnitude", func() {
		BeforeEach(func() {
			v = wyvern.Vector{3, 4}
		})

		It("Returns the correct magnitude of the vector", func() {
			Expect(v.Magnitude()).To(Equal(5.0))
		})
	})

	Describe("Angle", func() {
		BeforeEach(func() {
			v = wyvern.Vector{1, 0}
			w = wyvern.Vector{0, 1}
		})

		It("Returns the angle in radians between the two vectors", func() {
			Expect(v.Angle(w)).To(Equal(math.Pi / 2))
		})
	})
})
