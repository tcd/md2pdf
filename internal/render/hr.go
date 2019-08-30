package render

import "github.com/jung-kurt/gofpdf"

// HR writes a Horizontal Rule to a gofpdf.Fpdf.
func HR(f *gofpdf.Fpdf) {
	x := f.GetX()
	y := f.GetY()
	f.SetLineWidth(1)
	f.SetDrawColor(HrFG())
	f.Line(x, y, x+pageWidth, y)
	f.Ln(2)
}
