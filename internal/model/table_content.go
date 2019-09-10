package model

import (
	gofpdf "github.com/tcd/gofpdf-1"
	"github.com/tcd/md2pdf/internal/lib"
)

// TableContent represents the contents of a table element.
type TableContent struct {
	Rows       [][]Contents `json:"rows"`
	Alignments []string     `json:"alignments"` // "L", "C", or "R"
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
		// widths[i] = pdf.GetStringWidth(longest) * 1.5
		widths[i], _ = pdf.MeasureTextWidth(longest, gofpdf.TextOption{})
	}
	return widths
}

// headerWidths returns the width needed to hold the header cell in each column.
func (tc TableContent) headerWidths(pdf *gofpdf.Fpdf) []float64 {
	widths := make([]float64, tc.ColCount())
	for i, header := range tc.Headers() {
		// widths[i] = pdf.GetStringWidth(header.JoinContent()) * 1.5
		widths[i], _ = pdf.MeasureTextWidth(header.JoinContent(), gofpdf.TextOption{})
	}
	return widths
}

// Widths returns the width needed to hold the longest cell in each column.
func (tc TableContent) Widths(pdf *gofpdf.Fpdf, cellMargin float64) []float64 {
	colCount := tc.ColCount()
	cellPadding := (cellMargin * 2)
	tableWidth := (lib.ContentBoxWidth(pdf) - (cellPadding * float64(colCount)))
	headerWidths := tc.headerWidths(pdf)
	widths := tc.longestWidths(pdf)
	if sum(widths...) <= tableWidth {
		return widths
	}

	finalWidths := make([]float64, colCount)

	for i := range finalWidths {
		remainingWidth := tableWidth - sum(headerWidths...)
		portions := percentages(widths...)
		finalWidths[i] = (remainingWidth * portions[i]) + headerWidths[i] + cellMargin
	}
	return finalWidths
}

// WidthsAlt is an older version of Widths.
// Still trying to figure out the best way to implement this.
func (tc TableContent) WidthsAlt(pdf *gofpdf.Fpdf, cellMargin float64) []float64 {
	colCount := tc.ColCount()
	cellPadding := (cellMargin * 2)
	tableWidth := (lib.ContentBoxWidth(pdf) - (cellPadding * float64(colCount)))
	headerWidths := tc.headerWidths(pdf)
	widths := tc.longestWidths(pdf)
	if sum(widths...) <= tableWidth {
		return widths
	}

	finalWidths := make([]float64, colCount)

	if len(widths) == 2 {
		remainingWidth := tableWidth - sum(headerWidths...)
		portions := percentages(widths...)
		finalWidths[0] = (remainingWidth * portions[0]) + headerWidths[0] + cellMargin
		finalWidths[1] = (remainingWidth * portions[1]) + headerWidths[1] + cellMargin
		return finalWidths
	}

	for i := range finalWidths {
		var finalWidth float64
		portions := percentages(widths...)
		if portions[i] == 1 {
			finalWidths[i] = tableWidth + cellPadding
			continue
		}
		if widths[i] < tableWidth && tableWidth*portions[i] < headerWidths[i] {
			finalWidth = headerWidths[i] + cellPadding
			tableWidth = tableWidth - (headerWidths[i] + cellPadding)
			widths[i] = 0
		} else {
			finalWidth = (tableWidth * (portions[i] + cellPadding))
		}
		finalWidths[i] = finalWidth
	}

	return finalWidths
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
