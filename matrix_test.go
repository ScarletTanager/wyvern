package wyvern_test

import (
	"github.com/ScarletTanager/wyvern"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Matrix", func() {
	var (
		mt wyvern.Matrix[float64]
		c  []wyvern.Vector[float64]
	)

	Describe("Constructors", func() {
		When("The vectors have the same numbers of components", func() {
			BeforeEach(func() {
				c = []wyvern.Vector[float64]{
					{1, 2, 3},
					{4, 5, 6},
					{10, -12, 37},
				}
			})

			Describe("FromRows", func() {
				It("Returns a non-nil matrix", func() {
					m, _ := wyvern.FromRows(c)
					Expect(m.Columns()).NotTo(BeNil())
				})

				It("Does not return an error", func() {
					_, e := wyvern.FromRows(c)
					Expect(e).NotTo(HaveOccurred())
				})

				It("Populates the matrix using the input vectors as rows", func() {
					m, _ := wyvern.FromRows(c)
					rows := m.Rows()
					Expect(rows).NotTo(BeNil())
					Expect(rows).To(Equal(c))

					cols := m.Columns()
					Expect(cols).NotTo(BeNil())
					Expect(cols).To(Equal([]wyvern.Vector[float64]{
						{1, 4, 10},
						{2, 5, -12},
						{3, 6, 37},
					}))
				})
			})

			Describe("FromColumns", func() {
				It("Returns a matrix", func() {
					m, _ := wyvern.FromColumns(c)
					Expect(m.Columns()).NotTo(BeNil())
				})

				It("Does not return an error", func() {
					_, e := wyvern.FromColumns(c)
					Expect(e).NotTo(HaveOccurred())
				})

				It("Populates the matrix using the input vectors as rows", func() {
					m, _ := wyvern.FromColumns(c)

					cols := m.Columns()
					Expect(cols).NotTo(BeNil())
					Expect(cols).To(Equal(c))

					rows := m.Rows()
					Expect(rows).NotTo(BeNil())
					Expect(rows).To(Equal([]wyvern.Vector[float64]{
						{1, 4, 10},
						{2, 5, -12},
						{3, 6, 37},
					}))
				})
			})
		})

		When("The vectors contain differing numbers of components", func() {
			BeforeEach(func() {
				c = []wyvern.Vector[float64]{
					{1, 2, 3},
					{4, 5},
					{10, -12, 37},
				}
			})

			Describe("FromRows", func() {
				It("Returns an empty Matrix and an error", func() {
					m, e := wyvern.FromRows(c)
					Expect(m).To(Equal(wyvern.Matrix[float64]{}))
					Expect(e).To(HaveOccurred())
				})
			})

			Describe("FromColumns", func() {
				It("Returns an empty Matrix and an error", func() {
					m, e := wyvern.FromColumns(c)
					Expect(m).To(Equal(wyvern.Matrix[float64]{}))
					Expect(e).To(HaveOccurred())
				})
			})
		})
	})

	Describe("Matrix methods", func() {
		JustBeforeEach(func() {
			mt, _ = wyvern.FromColumns(c)
			Expect(mt).NotTo(Equal(wyvern.Matrix[float64]{}))
		})

		BeforeEach(func() {
			c = []wyvern.Vector[float64]{
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
				Expect(mt.Rows()).To(Equal([]wyvern.Vector[float64]{
					{1, 2, 13},
					{4, 5, 10},
					{9, 7, -4},
				}))
			})
		})

		Describe("Product", func() {

		})

		Describe("MultiplyRow", func() {
			var (
				rowIndex     int
				originalRows []wyvern.Vector[float64]
			)

			BeforeEach(func() {
				rowIndex = 1
			})

			JustBeforeEach(func() {
				originalRows = mt.Rows()
			})

			It("Multiplies the specified row by the specified factor", func() {
				Expect(originalRows[rowIndex]).To(Equal(wyvern.Vector[float64]{4, 5, 10}))
				mt.MultiplyRow(rowIndex, 3.0)
				originalRows[rowIndex].Multiply(3.0)
				Expect(mt.Rows()[rowIndex]).To(Equal(originalRows[rowIndex]))
			})

			When("The specified row index is not valid for the matrix", func() {
				BeforeEach(func() {
					rowIndex = 5
				})

				It("Returns an error and does not modify the matrix", func() {
					e := mt.MultiplyRow(rowIndex, 3.0)
					Expect(e).To(HaveOccurred())
					Expect(mt.Rows()).To(Equal(originalRows))
				})
			})
		})
	})
})
