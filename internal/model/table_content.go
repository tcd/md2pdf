package model

import (
	"github.com/jung-kurt/gofpdf"
)

// TableContent represents the contents of a table element.
type TableContent struct {
	Rows       [][]Contents `json:"rows"`
	Alignments []string     `json:"alignments"` // "L", "C", or "R"
}

// AddRow to TableContent.
func (tc *TableContent) AddRow(cols []Contents) {
	tc.Rows = append(tc.Rows, cols)
}

// AddRows to TableContent.
func (tc *TableContent) AddRows(cols ...[]Contents) {
	tc.Rows = append(tc.Rows, cols...)
}

// Headers returns the first slice of string in Rows.
func (tc TableContent) Headers() []Contents {
	if len(tc.Rows) != 0 {
		return tc.Rows[0]
	}
	return []Contents{}
}

// Body returns all rows after the first.
func (tc TableContent) Body() [][]Contents {
	if len(tc.Rows) > 0 {
		return tc.Rows[1:]
	}
	return [][]Contents{}
}

// ColCount returns the number of header columns.
func (tc TableContent) ColCount() int {
	return len(tc.Headers())
}

// GetColumn returns one string for every cell in a column at the given index.
// First column is 0, not 1.
func (tc TableContent) GetColumn(index int) []Contents {
	columns := make([]Contents, len(tc.Rows))
	for i, row := range tc.Rows {
		if len(row) > index {
			columns[i] = row[index]
		} else {
			columns[i] = Contents{}
		}
	}
	return columns
}

// longestWidths returns the width of the longest cell in each column.
func (tc TableContent) longestWidths(pdf *gofpdf.Fpdf) []float64 {
	widths := make([]float64, tc.ColCount())
	for i := 0; i < tc.ColCount(); i++ {
		col := tc.GetColumn(i)
		longest := ""
		for _, cell := range col {
			if len(cell.JoinContent()) > len(longest) {
				longest = cell.JoinContent()
			}
		}
		widths[i] = pdf.GetStringWidth(longest) * 1.5
	}
	return widths
}

// headerWidths returns the width needed to hold the header cell in each column.
func (tc TableContent) headerWidths(pdf *gofpdf.Fpdf) []float64 {
	widths := make([]float64, tc.ColCount())
	for i, header := range tc.Headers() {
		widths[i] = pdf.GetStringWidth(header.JoinContent()) * 1.5
	}
	return widths
}

// Widths returns the width needed to hold the longest cell in each column.
func (tc TableContent) Widths(pdf *gofpdf.Fpdf, cellMargin float64) []float64 {
	colCount := tc.ColCount()
	cellPadding := (cellMargin * 2)
	tableWidth := (ContentBoxWidth(pdf) - (cellPadding * float64(colCount)))
	widths := tc.longestWidths(pdf)
	if sum(widths...) <= tableWidth {
		return widths
	}

	finalWs := make([]float64, colCount)

	if len(widths) == 2 {
		portions := percentages(widths...)
		// larger := portions[0]
		if portions[0] > portions[1] {
			finalWs[0] = tableWidth * portions[0]
			finalWs[1] = tableWidth - finalWs[0]
		} else {
			finalWs[1] = tableWidth * portions[1]
			finalWs[0] = tableWidth - finalWs[1]
		}
		return finalWs
	}

	hWidths := tc.headerWidths(pdf)

	for i := range finalWs {
		var finalWidth float64
		portions := percentages(widths...)
		if portions[i] == 1 {
			finalWs[i] = tableWidth + cellPadding
			continue
		}
		if widths[i] < tableWidth && tableWidth*portions[i] < hWidths[i] {
			finalWidth = hWidths[i] + cellPadding
			tableWidth = tableWidth - (hWidths[i] + cellPadding)
			widths[i] = 0
		} else {
			finalWidth = (tableWidth * (portions[i] + cellPadding))
		}
		finalWs[i] = finalWidth
	}

	return finalWs
}

// Return the sum of any number of `float64`s.
func sum(items ...float64) (sum float64) {
	for i := range items {
		sum += items[i]
	}
	return
}

// Returns the percentage of all items out of their sum.
func percentages(items ...float64) []float64 {
	total := sum(items...)
	percentages := make([]float64, len(items))
	for i := range items {
		percentages[i] = items[i] / total
	}
	return percentages
}

// Return a new slice with one item removed.
func remove(slice []float64, s int) []float64 {
	return append(slice[:s], slice[s+1:]...)
}

// ContentBoxWidth returns the width of a width minus left and right margins.
func ContentBoxWidth(fpdf *gofpdf.Fpdf) float64 {
	tWidth, _ := fpdf.GetPageSize()
	leftM, _, rightM, _ := fpdf.GetMargins()
	return (tWidth - leftM - rightM)
}
