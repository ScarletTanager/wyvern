package wyvern_test

import (
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ScarletTanager/wyvern"
)

var _ = Describe("Vector", func() {
	var (
		v, w wyvern.Vector[float64]
	)

	BeforeEach(func() {
		v = wyvern.Vector[float64]{1, 2, 3}
		w = wyvern.Vector[float64]{4, 5, 6}
	})

	Describe("DotProduct", func() {
		It("computes the correct dot product", func() {
			Expect(v.DotProduct(w)).To(Equal(32.0))
		})
	})

	Describe("Magnitude", func() {
		BeforeEach(func() {
			v = wyvern.Vector[float64]{3, 4}
		})

		It("Returns the correct magnitude of the vector", func() {
			Expect(v.Magnitude()).To(Equal(5.0))
		})
	})

	Describe("Angle", func() {
		BeforeEach(func() {
			v = wyvern.Vector[float64]{1, 0}
			w = wyvern.Vector[float64]{0, 1}
		})

		It("Returns the angle in radians between the two vectors", func() {
			Expect(v.Angle(w)).To(Equal(math.Pi / 2))
		})
	})

	Describe("MultiplyComponent", func() {
		BeforeEach(func() {
			v = wyvern.Vector[float64]{3, 4}
			w = wyvern.Vector[float64]{3, 12}
			Expect(v).NotTo(Equal(w))
		})

		It("Multiples the specified component by the specified factor", func() {
			v.MultiplyComponent(1, 3.0)
			Expect(v).To(Equal(w))
		})
	})

	Describe("Multiply", func() {
		BeforeEach(func() {
			v = wyvern.Vector[float64]{3, 4}
			w = wyvern.Vector[float64]{9, 12}
			Expect(v).NotTo(Equal(w))
		})

		It("Multiples all components of the vector by the specified factor", func() {
			v.Multiply(3.0)
			Expect(v).To(Equal(w))
		})

		It("Returns the updated vector", func() {
			updated := v.Multiply(3.0)
			Expect(updated).To(Equal(w))
		})
	})
})
