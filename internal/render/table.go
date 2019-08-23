package render

import (
	"github.com/jung-kurt/gofpdf"
)

// Table writes a table with nicely split lines to a gofpdf.Fpdf.
// https://github.com/jung-kurt/gofpdf/blob/master/fpdf_test.go#L525
func Table(f *gofpdf.Fpdf, table TableContent) {
	oldCellMargin := f.GetCellMargin()
	oldLineWidth := f.GetLineWidth()

	var widths []float64
	var tbMargin float64 = 2 // top & bottom cell margins
	var lrMargin float64 = 3 // left & right cell margins

	f.SetFont("helvetica", "B", 12) // Header Font
	f.SetCellMargin(lrMargin)       // left & right cell margins
	f.SetTextColor(36, 41, 46)      // text color
	f.SetDrawColor(223, 226, 229)   // border color
	f.SetLineWidth(0.3)             // outline width
	f.SetFillColor(246, 248, 250)   // background color for every other cell

	widths = table.Widths(f, lrMargin)

	for i, header := range table.Headers() {
		f.CellFormat(
			widths[i]+lrMargin, // width
			10,                 // height
			header,             // textStr
			"1",                // borderStr; 1 for full border.
			0,                  // ln; position after the call.
			"CM",               // alignStr; "L", "C" or "R" + "T", "M", "B" or "A" (left, center, right + top, middle, bottom, baseline)
			false,              // fill
			0,                  // link
			"",                 // linkStr
		)
	}
	f.Ln(-1)

	f.SetFont("helvetica", "", 12) // Body Font

	fill := false
	alignments := table.Alignments
	_, pageHeight := f.GetPageSize()
	_, _, _, bottomMargin := f.GetMargins()

	for _, row := range table.Body() {

		startX, y := f.GetXY()
		x := startX
		height := 0.00
		_, lineHt := f.GetFontSize()

		for i, txt := range row {
			var lines [][]byte
			if f.GetStringWidth(txt) < widths[i] {
				lines = make([][]byte, 1)
				lines[0] = []byte(txt)
			} else {
				lines = f.SplitLines([]byte(txt), widths[i]-lrMargin*2)
			}
			h := float64(len(lines))*lineHt + tbMargin*2*float64(len(lines))
			if h > height {
				height = h
			}
		}

		// add a new page if the height of the row doesn't fit on the page
		if f.GetY()+height > pageHeight-bottomMargin {
			f.AddPage()
			y = f.GetY()
		}

		for i, txt := range row {
			width := widths[i] + lrMargin
			if fill {
				f.Rect(x, y, width, height, "FD")
			} else {
				f.Rect(x, y, width, height, "D")
			}
			f.SetX(x)
			f.MultiCell(width, lineHt+(tbMargin*2), txt, "", alignments[i], false)
			x += width
			f.SetXY(x, y)
		}
		f.SetXY(startX, y+height-8.1)
		f.Ln(-1)
		fill = !fill
	}
	f.Ln(5)
	f.SetCellMargin(oldCellMargin)
	f.SetLineWidth(oldLineWidth)
}

// TableContent represents the contents of a table element.
type TableContent struct {
	Rows       [][]string `json:"rows"`
	Alignments []string   `json:"alignments"` // "L", "C", or "R"
}

// AddRow to TableContent.
func (tc *TableContent) AddRow(cols []string) {
	tc.Rows = append(tc.Rows, cols)
}

// AddRows to TableContent.
func (tc *TableContent) AddRows(cols ...[]string) {
	tc.Rows = append(tc.Rows, cols...)
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

// longestWidths returns the width of the longest cell in each column.
func (tc TableContent) longestWidths(f *gofpdf.Fpdf) []float64 {
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

// headerWidths returns the width needed to hold the header cell in each column.
func (tc TableContent) headerWidths(f *gofpdf.Fpdf) []float64 {
	widths := make([]float64, tc.ColCount())
	for i, header := range tc.Headers() {
		widths[i] = f.GetStringWidth(header) * 1.5
	}
	return widths
}

// Widths returns the width needed to hold the longest cell in each column.
func (tc TableContent) Widths(f *gofpdf.Fpdf, cellMargin float64) []float64 {
	pageWidth, _ := f.GetPageSize()
	leftM, _, rightM, _ := f.GetMargins()
	colCount := tc.ColCount()
	cellPadding := (cellMargin * 2)
	tableWidth := (pageWidth - leftM - rightM - (cellPadding * float64(colCount)))

	widths := tc.longestWidths(f)
	if sum(widths...) <= tableWidth {
		return widths
	}

	hWidths := tc.headerWidths(f)
	finalWs := make([]float64, colCount)

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
