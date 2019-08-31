package render

import "github.com/jung-kurt/gofpdf"

// HR writes a Horizontal Rule to a gofpdf.Fpdf.
func HR(pdf *gofpdf.Fpdf) {
	x := pdf.GetX()
	y := pdf.GetY()
	pdf.SetLineWidth(1)
	pdf.SetDrawColor(HrFG())
	pdf.Line(x, y, x+pageWidth, y)
	pdf.Ln(2)
}
