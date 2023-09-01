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

		Describe("Row", func() {
			var (
				rowIndex int
			)

			BeforeEach(func() {
				c = []wyvern.Vector[float64]{
					{4, 7},
					{5, 13},
					{12, -6},
					{0, 9},
				}

				rowIndex = 1
			})

			It("Returns the specified row Vector (as a copy)", func() {
				row, e := mt.Row(rowIndex)
				Expect(row).NotTo(BeNil())
				Expect(e).NotTo(HaveOccurred())
				Expect(row).To(Equal(wyvern.Vector[float64]{7, 13, -6, 9}))

				row.Multiply(5.0)
				r2, _ := mt.Row(rowIndex)
				Expect(r2).To(Equal(wyvern.Vector[float64]{7, 13, -6, 9}))
				Expect(row).NotTo(Equal(r2))
			})

			When("The row index is out of bounds", func() {
				BeforeEach(func() {
					rowIndex = 2
				})

				It("Returns nil and an error", func() {
					row, e := mt.Row(rowIndex)
					Expect(row).To(BeNil())
					Expect(e).To(HaveOccurred())
				})
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

		Describe("Column", func() {
			var (
				columnIndex int
			)

			BeforeEach(func() {
				c = []wyvern.Vector[float64]{
					{12, 3, 9},
					{24, 17, -47},
					{-8, 99, 104},
					{5, 5, 11},
				}

				columnIndex = 2
			})

			It("Returns the specified column Vector (as a copy)", func() {
				col, e := mt.Column(columnIndex)
				Expect(col).NotTo(BeNil())
				Expect(e).NotTo(HaveOccurred())
				Expect(col).To(Equal(c[columnIndex]))

				col.Multiply(4.5)
				c2, _ := mt.Column(columnIndex)
				Expect(c2).To(Equal(c[columnIndex]))
				Expect(col).NotTo(Equal(c2))
			})

			When("The column index is out of bounds", func() {
				BeforeEach(func() {
					columnIndex = 4
				})

				It("Returns nil and an error", func() {
					col, e := mt.Column(columnIndex)
					Expect(col).To(BeNil())
					Expect(e).To(HaveOccurred())
				})
			})
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

		Describe("ReplaceRow", func() {
			var (
				rowIndex int
				newRow   wyvern.Vector[float64]
			)

			BeforeEach(func() {
				c = []wyvern.Vector[float64]{
					{9, 7, 5.5},
					{3, 6, 11},
					{2, 2, -13},
					{0, 0, 74},
				}

				rowIndex = 2
				newRow = wyvern.Vector[float64]{3, 3, 12, -16}
			})

			It("Replaces the specified row with the supplied vector", func() {
				before, _ := mt.Row(rowIndex)
				e := mt.ReplaceRow(rowIndex, newRow)
				Expect(e).NotTo(HaveOccurred())
				after, _ := mt.Row(rowIndex)
				Expect(after).To(Equal(newRow))
				Expect(after).NotTo(Equal(before))
			})

			When("The row index is out of bounds", func() {
				BeforeEach(func() {
					rowIndex = 3
				})

				It("Returns an error and leaves the matrix unchanged", func() {
					before := mt.Rows()
					e := mt.ReplaceRow(rowIndex, newRow)
					Expect(e).To(HaveOccurred())
					after := mt.Rows()
					Expect(after).To(Equal(before))
				})
			})

			When("The supplied vector has the wrong dimension", func() {
				BeforeEach(func() {
					newRow = wyvern.Vector[float64]{3, 3, 12}
				})

				It("Returns an error and leaves the matrix unchanged", func() {
					before := mt.Rows()
					e := mt.ReplaceRow(rowIndex, newRow)
					Expect(e).To(HaveOccurred())
					after := mt.Rows()
					Expect(after).To(Equal(before))
				})
			})
		})

		Describe("ReplaceColumn", func() {
			var (
				columnIndex int
				newColumn   wyvern.Vector[float64]
			)

			BeforeEach(func() {
				c = []wyvern.Vector[float64]{
					{9, 7, 5.5},
					{3, 6, 11},
					{2, 2, -13},
					{0, 0, 74},
				}

				columnIndex = 3
				newColumn = wyvern.Vector[float64]{3, 3, 12}
			})

			It("Replaces the specified row with the supplied vector", func() {
				before, _ := mt.Column(columnIndex)
				e := mt.ReplaceColumn(columnIndex, newColumn)
				Expect(e).NotTo(HaveOccurred())
				after, _ := mt.Column(columnIndex)
				Expect(after).To(Equal(newColumn))
				Expect(after).NotTo(Equal(before))
			})

			When("The row index is out of bounds", func() {
				BeforeEach(func() {
					columnIndex = 4
				})

				It("Returns an error and leaves the matrix unchanged", func() {
					before := mt.Columns()
					e := mt.ReplaceColumn(columnIndex, newColumn)
					Expect(e).To(HaveOccurred())
					after := mt.Columns()
					Expect(after).To(Equal(before))
				})
			})

			When("The supplied vector has the wrong dimension", func() {
				BeforeEach(func() {
					newColumn = wyvern.Vector[float64]{3, 3, 12, 17, 5}
				})

				It("Returns an error and leaves the matrix unchanged", func() {
					before := mt.Columns()
					e := mt.ReplaceColumn(columnIndex, newColumn)
					Expect(e).To(HaveOccurred())
					after := mt.Columns()
					Expect(after).To(Equal(before))
				})
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

		Describe("Product", func() {
			var (
				mtA, mtB            wyvern.Matrix[float64]
				colsA, colsB, colsU []wyvern.Vector[float64]
			)

			BeforeEach(func() {
				colsA = []wyvern.Vector[float64]{
					{1, 0, -3},
					{0, 1, 0},
					{0, 0, 1},
				}

				colsB = []wyvern.Vector[float64]{
					{2, 0, 6},
					{1, 4, 3},
					{0, 2, 5},
				}

				colsU = []wyvern.Vector[float64]{
					{2, 0, 0},
					{1, 4, 0},
					{0, 2, 5},
				}
			})

			JustBeforeEach(func() {
				mtA, _ = wyvern.FromColumns(colsA)
				mtB, _ = wyvern.FromColumns(colsB)
			})

			It("Returns the matrix product of the two matrices", func() {
				product, _ := mtA.Product(mtB)
				Expect(product.Columns()).To(Equal(colsU))
			})
		})
	})
})
