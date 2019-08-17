package render

import "github.com/jung-kurt/gofpdf"

// H1 write a 1st level header to a gofpdf.Fpdf.
func H1(f *gofpdf.Fpdf, text string) {
	f.SetFont("helvetica", "B", 24)
	f.SetTextColor(36, 41, 46)

	f.Ln(5)
	f.Write(1.5, text)
	f.Ln(8)

	x := f.GetX()
	y := f.GetY()
	f.SetLineWidth(0.09)
	f.SetDrawColor(191, 191, 191)
	f.Line(x, y, x+pageWidth, y)
	f.Ln(6)
}

// H2 write a 2nd level header to a gofpdf.Fpdf.
func H2(f *gofpdf.Fpdf, text string) {
	f.SetFont("helvetica", "B", 18)
	f.SetTextColor(36, 41, 46)

	f.Ln(5)
	f.Write(1.5, text)
	f.Ln(6)

	x := f.GetX()
	y := f.GetY()
	f.SetLineWidth(0.09)
	f.SetDrawColor(191, 191, 191)
	f.Line(x, y, x+pageWidth, y)
	f.Ln(5)
}

// H3 write a 3rd level header to a gofpdf.Fpdf.
func H3(f *gofpdf.Fpdf, text string) {
	f.SetFont("helvetica", "B", 15)
	f.SetTextColor(36, 41, 46)

	f.Ln(5)
	f.Write(1.5, text)
	f.Ln(7)
}

// H4 write a 4th level header to a gofpdf.Fpdf.
func H4(f *gofpdf.Fpdf, text string) {
	f.SetFont("helvetica", "B", 12)
	f.SetTextColor(36, 41, 46)

	f.Ln(5)
	f.Write(1.5, text)
	f.Ln(7)
}

// H5 write a 5th level header to a gofpdf.Fpdf.
func H5(f *gofpdf.Fpdf, text string) {
	f.SetFont("helvetica", "B", 10.5)
	f.SetTextColor(36, 41, 46)

	f.Ln(5)
	f.Write(1.5, text)
	f.Ln(6)
}

// H6 write a 6th level header to a gofpdf.Fpdf.
func H6(f *gofpdf.Fpdf, text string) {
	f.SetFont("helvetica", "B", 10.5)
	f.SetTextColor(106, 115, 125)

	f.Ln(5)
	f.Write(1.5, text)
	f.Ln(6)
}
