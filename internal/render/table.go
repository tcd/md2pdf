package render

import (
	"log"

	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/lib"
	"github.com/tcd/md2pdf/internal/model"
)

// Table writes a table to a gofpdf.Fpdf.
func Table(f *gofpdf.Fpdf, table model.TableContent) {
	tableData := mockTable(f, table)
	if len(tableData.Rows) == 0 {
		log.Println("Error rendering table")
		return
	}
	bodyData := tableData.Rows[1:]

	oldCellMargin := f.GetCellMargin()
	oldLineWidth := f.GetLineWidth()

	var tbMargin float64 = 2 // top & bottom cell margins
	var lrMargin float64 = 3 // left & right cell margins

	f.SetFont("helvetica", "B", 12) // Header Font
	f.SetLineWidth(0.3)             // outline width
	f.SetCellMargin(lrMargin)       // left & right cell margins
	f.SetTextColor(DefaultFG())     // text color
	f.SetDrawColor(BlockquoteBG())  // border color
	f.SetFillColor(TableCellBG())   // background color for every other cell

	widths := table.Widths(f, lrMargin)

	for i, header := range table.Headers() {
		f.CellFormat(widths[i]+lrMargin, 10, header.JoinContent(), "1", 0, "CM", false, 0, "")
	}
	f.Ln(0)                        // Move to next row
	f.SetFont("helvetica", "", 12) // Set font for table body.

	fill := false
	pageWidth, _ := f.GetPageSize()

	for r, row := range table.Body() {
		if bodyData[r].Cells[0].Y < f.GetY() {
			f.AddPage()
		}
		for c, cell := range row {
			x, y, width, height := bodyData[r].Cells[c].Data()
			if fill {
				f.Rect(x, y, width, height, "FD")
			} else {
				f.Rect(x, y, width, height, "D")
			}
			f.SetXY(x, y+(tbMargin/2))
			f.SetLeftMargin(x)
			f.SetRightMargin(pageWidth - (x + width))
			drawContent(f, cell, 12)
		}
		f.SetMargins(20, 20, 20)
		f.Ln(-1)
		fill = !fill
	}

	f.Ln(-1)
	f.Ln(5)
	f.SetCellMargin(oldCellMargin)
	f.SetLineWidth(oldLineWidth)
}

// Draw out a table on a temporary fpdf to figure out cell sizes to draw
// on our primary pdf.
func mockTable(f *gofpdf.Fpdf, table model.TableContent) tableData {
	var data tableData

	pdf := gofpdf.New("P", "mm", "Letter", "")
	Setup(pdf)
	pdf.SetXY(f.GetXY())

	var tbMargin float64 = 2          // top & bottom cell margins
	var lrMargin float64 = 3          // left & right cell margins
	pdf.SetCellMargin(lrMargin)       // left & right cell margins
	pdf.SetLineWidth(0.3)             // outline width
	pdf.SetFont("helvetica", "B", 12) // Header Font
	widths := table.Widths(pdf, lrMargin)

	data.NewRow()
	for i, cell := range table.Headers() {
		x, y := pdf.GetXY()
		data.AddCell(x, y, widths[i]+lrMargin, 10)
		pdf.CellFormat(widths[i]+lrMargin, 10, cell.JoinContent(), "1", 0, "CM", false, 0, "")
	}
	pdf.Ln(-1)

	// Need to set font so GetFontSize returns appropreate values.
	pdf.SetFont("helvetica", "", 12)
	_, lineHt := pdf.GetFontSize()
	alignments := table.Alignments

	for _, row := range table.Body() {
		data.NewRow()

		startX, y := pdf.GetXY()
		x := startX
		height := 0.00

		// Determine the height needed for a cell.
		for i, cell := range row {
			var lines []string
			if pdf.GetStringWidth(cell.JoinContent()) < widths[i] {
				lines = []string{cell.JoinContent()}
			} else {
				lines = pdf.SplitText(cell.JoinContent(), widths[i]-lrMargin*2)
			}
			h := float64(len(lines))*lineHt + tbMargin*2*float64(len(lines))
			if h > height {
				height = h
			}
		}

		// add a new page if the height of the row doesn't fit on the page
		if pdf.GetY()+height > lib.ContentBoxBottom(pdf) {
			pdf.AddPage()
			y = pdf.GetY()
		}

		for i, cell := range row {
			width := widths[i] + lrMargin
			data.AddCell(x, y, width, height) // Store the values needed for drawing each cell.
			pdf.Rect(x, y, width, height, "D")
			pdf.SetX(x)
			pdf.MultiCell(width, lineHt+(tbMargin*2), cell.JoinContent(), "", alignments[i], false)
			x += width
			pdf.SetXY(x, y)
		}
		pdf.SetXY(startX, y+height-8.1) // FIXME: 8.1 is a magic number. Just obtained from trial and error. Needs to be replaced with a proper calculation.
		pdf.Ln(-1)
	}

	pdf.Close() // Close our mock pdf.
	return data
}

// Information needed to draw a cell.
type cell struct {
	X, Y, Width, Height float64
}

// Wrapper to succinctly extract all field values.
func (c cell) Data() (x, y, width, height float64) {
	return c.X, c.Y, c.Width, c.Height
}

// Wrapper around a slice of cell.
type row struct {
	Cells []cell
}

// Create a cell struct from data and add it to a row.
func (r *row) AddCell(x, y, width, height float64) {
	r.Cells = append(r.Cells, cell{x, y, width, height})
}

// Wrapper around a slice of row.
type tableData struct {
	Rows []row
}

// Add a new row to populate with data.
func (td *tableData) NewRow() {
	td.Rows = append(td.Rows, row{})
}

// Add a cell to the last row.
func (td *tableData) AddCell(x, y, width, height float64) {
	if len(td.Rows) == 0 {
		td.NewRow()
	}
	td.Rows[len(td.Rows)-1].AddCell(x, y, width, height)
}

// BasicTable writes a table with nicely split lines to a gofpdf.Fpdf.
// https://github.com/jung-kurt/gofpdf/blob/master/fpdf_test.go#L525
func BasicTable(f *gofpdf.Fpdf, table model.TableContent) {
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
