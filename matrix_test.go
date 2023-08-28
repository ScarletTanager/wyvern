package wyvern_test

import (
	"github.com/ScarletTanager/wyvern"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Matrix", func() {
	var (
		mt wyvern.Matrix
		c  []wyvern.Vector
	)

	Describe("NewMatrix", func() {
		When("The vectors have the same numbers of components", func() {
			BeforeEach(func() {
				c = []wyvern.Vector{
					{1, 2, 3},
					{4, 5, 6},
					{10, -12, 37},
				}
			})

			It("Returns a matrix", func() {
				m, _ := wyvern.NewMatrix(c)
				Expect(m.Vectors).To(HaveLen(3))
			})

			It("Does not return an error", func() {
				_, e := wyvern.NewMatrix(c)
				Expect(e).NotTo(HaveOccurred())
			})
		})

		When("The vectors contain differing numbers of components", func() {
			BeforeEach(func() {
				c = []wyvern.Vector{
					{1, 2, 3},
					{4, 5},
					{10, -12, 37},
				}
			})

			It("Returns an empty Matrix and an error", func() {
				m, e := wyvern.NewMatrix(c)
				Expect(m).To(Equal(wyvern.Matrix{}))
				Expect(e).To(HaveOccurred())
			})
		})
	})

	Describe("Matrix methods", func() {
		JustBeforeEach(func() {
			mt, _ = wyvern.NewMatrix(c)
			Expect(mt).NotTo(Equal(wyvern.Matrix{}))
		})

		BeforeEach(func() {
			c = []wyvern.Vector{
				{1, 4, 9},
				{2, 5, 7},
				{13, 10, -4},
			}
		})

		Describe("Columns", func() {
			It("Returns the set of column vectors", func() {
				Expect(mt.Columns()).To(Equal(c))
			})
		})

		Describe("Rows", func() {
			It("Returns the set of row vectors", func() {
				Expect(mt.Rows()).To(Equal([]wyvern.Vector{
					{1, 2, 13},
					{4, 5, 10},
					{9, 7, -4},
				}))
			})
		})
	})
})
