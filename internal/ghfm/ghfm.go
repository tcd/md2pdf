// Package ghfm renders text in the style of GitHub using gofpdf.
package ghfm

import "github.com/jung-kurt/gofpdf"

const (
	pageWidth = 196 // mm
)

// Setup sets default GitHub-Flavored styles to a gofpdf.Fpdf.
func Setup(f *gofpdf.Fpdf) {
	f.AddPage()
	f.SetFont("helvetica", "", 16)
	f.SetMargins(10, 13, 10)
	f.SetFillColor(255, 255, 255)
	f.SetTextColor(36, 41, 46)
}

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
	f.Line(x, y, x+196, y)
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

// HR write a Horizontal Rule to a gofpdf.Fpdf.
func HR(f *gofpdf.Fpdf) {
	x := f.GetX()
	y := f.GetY()
	f.SetLineWidth(1)
	f.SetDrawColor(191, 191, 191)
	f.Line(x, y, x+pageWidth, y)
	f.Ln(2)
}
