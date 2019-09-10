package render

import (
	gofpdf "github.com/tcd/gofpdf-1"
	"github.com/tcd/md2pdf/internal/lib"
)

// hr draws a standard Horizontal Rule to a gofpdf.Fpdf.
func hr(pdf *gofpdf.Fpdf) {
	// pdf.SetFont("helvetica", "", 12)
	pdf.Ln(2)
	drawHR(pdf, 1.00)
	pdf.Ln(5)
}

// hr draws a horizontal rule with a given width.
func drawHR(pdf *gofpdf.Fpdf, lineWidth float64) {
	x, y := pdf.XY()
	pdf.SetLineWidth(lineWidth)
	pdf.SetRGBStrokeColor(HrFG())
	pdf.Line(x, y, x+lib.ContentBoxWidth(pdf), y)
}
