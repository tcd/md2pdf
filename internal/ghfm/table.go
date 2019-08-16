package ghfm

import (
	"github.com/jung-kurt/gofpdf"
)

// Table writes a table to a gofpdf.Fpdf.
func Table(f *gofpdf.Fpdf, table TableContent) {
	f.SetFont("helvetica", "", 12)
	oldCellMargin := f.GetCellMargin()
	oldLineWidth := f.GetLineWidth()
	headers := table.Headers()
	widths := table.Widths(f)
	body := table.Body()

	f.SetFont("helvetica", "B", 12)
	f.SetTextColor(36, 41, 46)    // text color
	f.SetDrawColor(223, 226, 229) // border color
	f.SetLineWidth(0.3)           // outline width
	f.SetFillColor(255, 255, 255) // background color
	f.SetCellMargin(3)
	for i, header := range headers {
		f.CellFormat(
			widths[i], // width
			10,        // height
			header,    // textStr
			"1",       // borderStr; 1 for full border.
			0,         // ln; position after the call.
			"CM",      // alignStr; "L", "C" or "R" + "T", "M", "B" or "A" (left, center, right + top, middle, bottom, baseline)
			false,     // fill
			0,         // link
			"",        // linkStr
		)
	}
	f.Ln(-1)

	f.SetFont("helvetica", "", 12)
	f.SetFillColor(246, 248, 250)
	fill := false
	for _, row := range body {
		for i, col := range row {
			f.CellFormat(
				widths[i],
				10,
				col,
				"1",
				0,
				table.Alignments[i],
				fill,
				0,
				"",
			)
		}
		fill = !fill
		f.Ln(-1)
	}

	f.Ln(5)
	f.SetCellMargin(oldCellMargin)
	f.SetLineWidth(oldLineWidth)
}

// TableContent represents the contents of a table element.
type TableContent struct {
	Rows       [][]string
	Alignments []string
}

// Headers returns the first slice of string in Rows.
func (tc TableContent) Headers() []string {
	if len(tc.Rows) != 0 {
		return tc.Rows[0]
	}
	return []string{}
}

// Body returns all rows after the first.
func (tc TableContent) Body() [][]string {
	if len(tc.Rows) > 0 {
		return tc.Rows[1:]
	}
	return [][]string{}
}

// ColCount returns the number of header columns.
func (tc TableContent) ColCount() int {
	return len(tc.Headers())
}

// GetColumn returns one string for every cell in a column at the given index.
// First column is 0, not 1.
func (tc TableContent) GetColumn(index int) []string {
	columns := make([]string, len(tc.Rows))
	for i, row := range tc.Rows {
		if len(row) > index {
			columns[i] = row[index]
		} else {
			columns[i] = ""
		}
	}
	return columns
}

// Widths returns the width needed to hold the longest cell in each column.
func (tc TableContent) Widths(f *gofpdf.Fpdf) []float64 {
	widths := make([]float64, tc.ColCount())
	for i := 0; i < tc.ColCount(); i++ {
		col := tc.GetColumn(i)
		longest := ""
		for _, cell := range col {
			if len(cell) > len(longest) {
				longest = cell
			}
		}
		widths[i] = f.GetStringWidth(longest) * 1.5
	}
	return widths
}
