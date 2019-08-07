package md2pdf

import "github.com/jung-kurt/gofpdf"

const (
	pageWidth = 196 // mm
)

// Write a first level header
func h1(f *gofpdf.Fpdf, text string) {
	f.SetFont("helvetica", "B", 24)
	f.SetTextColor(36, 41, 46)

	f.Ln(5)
	f.Write(1.5, text)
	f.Ln(8)

	x := f.GetX()
	y := f.GetY()
	f.SetLineWidth(0.09)
	f.SetDrawColor(191, 191, 191)
	f.Line(x, y, x+196, y)
	f.Ln(6)
}

// Write a second level header
func h2(f *gofpdf.Fpdf, text string) {
	f.SetFont("helvetica", "B", 18)
	f.SetTextColor(36, 41, 46)

	f.Ln(5)
	f.Write(1.5, text)
	f.Ln(6)

	x := f.GetX()
	y := f.GetY()
	f.SetLineWidth(0.09)
	f.SetDrawColor(191, 191, 191)
	f.Line(x, y, x+196, y)
	f.Ln(5)
}

// Write a third level header
func h3(f *gofpdf.Fpdf, text string) {
	f.SetFont("helvetica", "B", 15)
	f.SetTextColor(36, 41, 46)

	f.Ln(5)
	f.Write(1.5, text)
	f.Ln(7)
}

// Write a fourth level header
func h4(f *gofpdf.Fpdf, text string) {
	f.SetFont("helvetica", "B", 12)
	f.SetTextColor(36, 41, 46)

	f.Ln(5)
	f.Write(1.5, text)
	f.Ln(7)
}

// Write a fifth level header
func h5(f *gofpdf.Fpdf, text string) {
	f.SetFont("helvetica", "B", 10.5)
	f.SetTextColor(36, 41, 46)

	f.Ln(5)
	f.Write(1.5, text)
	f.Ln(6)
}

// Write a sixth level header
func h6(f *gofpdf.Fpdf, text string) {
	f.SetFont("helvetica", "B", 10.5)
	f.SetTextColor(106, 115, 125)

	f.Ln(5)
	f.Write(1.5, text)
	f.Ln(6)
}

// Write a Horizontal Rule
func hr(f *gofpdf.Fpdf) {
	x := f.GetX()
	y := f.GetY()
	f.SetLineWidth(1)
	f.SetDrawColor(191, 191, 191)
	f.Line(x, y, x+pageWidth, y)
	f.Ln(2)
}
