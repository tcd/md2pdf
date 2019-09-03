package render

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/model"
)

// Table writes a table with nicely split lines to a gofpdf.Fpdf.
// https://github.com/jung-kurt/gofpdf/blob/master/fpdf_test.go#L525
func Table(f *gofpdf.Fpdf, table model.TableContent) {
	oldCellMargin := f.GetCellMargin()
	oldLineWidth := f.GetLineWidth()

	var tbMargin float64 = 2 // top & bottom cell margins
	var lrMargin float64 = 3 // left & right cell margins

	f.SetFont("helvetica", "B", 12) // Header Font
	f.SetCellMargin(lrMargin)       // left & right cell margins
	f.SetTextColor(DefaultFG())     // text color
	f.SetDrawColor(BlockquoteBG())  // border color
	f.SetLineWidth(0.3)             // outline width
	f.SetFillColor(TableCellBG())   // background color for every other cell

	widths := table.Widths(f, lrMargin)

	for i, header := range table.Headers() {
		f.CellFormat(
			widths[i]+lrMargin,   // width
			10,                   // height
			header.JoinContent(), // textStr
			"1",                  // borderStr; 1 for full border.
			0,                    // ln; position after the call.
			"CM",                 // alignStr; "L", "C" or "R" + "T", "M", "B" or "A" (left, center, right + top, middle, bottom, baseline)
			false,                // fill
			0,                    // link
			"",                   // linkStr
		)
	}
	f.Ln(-1) // Move to next row

	f.SetFont("helvetica", "", 12) // Body Font

	fill := false
	_, lineHt := f.GetFontSize()
	alignments := table.Alignments
	_, pageHeight := f.GetPageSize()
	_, _, _, bottomMargin := f.GetMargins()

	for _, row := range table.Body() {

		startX, y := f.GetXY()
		x := startX
		height := 0.00

		// Determine the height needed for a cell.
		for i, cell := range row {
			var lines [][]byte
			if f.GetStringWidth(cell.JoinContent()) < widths[i] {
				lines = make([][]byte, 1)
				lines[0] = []byte(cell.JoinContent())
			} else {
				lines = f.SplitLines([]byte(cell.JoinContent()), widths[i]-lrMargin*2)
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

		for i, cell := range row {
			width := widths[i] + lrMargin
			// Draw the cell, fill every other row.
			if fill {
				f.Rect(x, y, width, height, "FD")
			} else {
				f.Rect(x, y, width, height, "D")
			}
			f.SetX(x)
			f.MultiCell(width, lineHt+(tbMargin*2), cell.JoinContent(), "", alignments[i], false)
			x += width
			f.SetXY(x, y)
		}
		f.SetXY(startX, y+height-8.1)
		f.Ln(-1) // Move to next row
		fill = !fill
	}
	f.Ln(5)
	f.SetCellMargin(oldCellMargin)
	f.SetLineWidth(oldLineWidth)
}
