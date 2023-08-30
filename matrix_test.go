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
			It("Returns a copy of the column vectors", func() {
				colsBefore := mt.Columns()
				Expect(colsBefore).To(Equal(c))
				colsBefore[0].Multiply(2.0)
				colsAfter := mt.Columns()
				Expect(colsAfter).To(Equal(c))
				Expect(colsBefore).NotTo(Equal(colsAfter))
			})
		})

		Describe("Rows", func() {
			It("Returns a copy of the set of row vectors", func() {
				rowsBefore := mt.Rows()
				Expect(rowsBefore).To(Equal([]wyvern.Vector[float64]{
					{1, 2, 13},
					{4, 5, 10},
					{9, 7, -4},
				}))
				rowsBefore[0].Multiply(5.0)
				rowsAfter := mt.Rows()
				Expect(rowsAfter).To(Equal([]wyvern.Vector[float64]{
					{1, 2, 13},
					{4, 5, 10},
					{9, 7, -4},
				}))
				Expect(rowsBefore).NotTo(Equal(rowsAfter))
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
				Expect(mt.Rows()).To(Equal(originalRows))
				mt.MultiplyRow(rowIndex, 3.0)
				Expect(mt.Rows()).NotTo(Equal(originalRows))
				Expect(mt.Rows()[rowIndex]).To(Equal(originalRows[rowIndex].Multiply(3.0)))
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

		Describe("MultiplyColumn", func() {
			var (
				columnIndex     int
				originalColumns []wyvern.Vector[float64]
				factor          float64
			)

			BeforeEach(func() {
				columnIndex = 2
				factor = 4.0
			})

			JustBeforeEach(func() {
				originalColumns = mt.Columns()
			})

			It("Multiplies the specified column by the specified factor", func() {
				Expect(mt.Columns()[columnIndex]).To(Equal(originalColumns[columnIndex]))
				mt.MultiplyColumn(columnIndex, factor)
				Expect(mt.Columns()[columnIndex]).NotTo(Equal(originalColumns[columnIndex]))
				Expect(mt.Columns()[columnIndex]).To(Equal(originalColumns[columnIndex].Multiply(factor)))
			})

			When("The specified column index is not valid for the matrix", func() {
				BeforeEach(func() {
					columnIndex = 7
				})

				It("Returns an error and does not modify the matrix", func() {
					e := mt.MultiplyColumn(columnIndex, factor)
					Expect(e).To(HaveOccurred())
					Expect(mt.Columns()).To(Equal(originalColumns))
				})
			})
		})
	})
})
