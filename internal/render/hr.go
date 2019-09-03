package render

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/lib"
)

// HR draws a standard Horizontal Rule to a gofpdf.Fpdf.
func HR(pdf *gofpdf.Fpdf) {
	pdf.Ln(2)
	hr(pdf, 1.00)
	pdf.Ln(2)
}

// hr draws a horizontal rule with a given width.
func hr(pdf *gofpdf.Fpdf, lineWidth float64) {
	x := pdf.GetX()
	y := pdf.GetY()
	pdf.SetLineWidth(lineWidth)
	pdf.SetDrawColor(HrFG())
	pdf.Line(x, y, x+lib.ContentBoxWidth(pdf), y)
}
