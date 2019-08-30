package render

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/model"
)

// H1 write a 1st level header to a gofpdf.Fpdf.
func H1(f *gofpdf.Fpdf, contents model.Contents) {
	f.SetFont("helvetica", "B", 24)
	f.SetTextColor(36, 41, 46)
	_, lineHt := f.GetFontSize()
	lineHt *= 1.7

	for _, txt := range contents.Content {
		f.Write(lineHt, txt.Text)
	}
	f.Ln(-1)

	x := f.GetX()
	y := f.GetY()
	f.SetLineWidth(0.09)
	f.SetDrawColor(191, 191, 191)
	f.Line(x, y, x+pageWidth, y)
	f.Ln(5)
}

// H2 write a 2nd level header to a gofpdf.Fpdf.
func H2(f *gofpdf.Fpdf, contents model.Contents) {
	f.SetFont("helvetica", "B", 18)
	f.SetTextColor(36, 41, 46)
	_, lineHt := f.GetFontSize()
	lineHt *= 1.6

	f.Ln(3)
	for _, txt := range contents.Content {
		f.Write(lineHt, txt.Text)
	}
	f.Ln(-1)

	x := f.GetX()
	y := f.GetY()
	f.SetLineWidth(0.09)
	f.SetDrawColor(191, 191, 191)
	f.Line(x, y, x+pageWidth, y)
	f.Ln(5)
}

// H3 write a 3rd level header to a gofpdf.Fpdf.
func H3(f *gofpdf.Fpdf, contents model.Contents) {
	f.SetFont("helvetica", "B", 15)
	f.SetTextColor(36, 41, 46)
	_, lineHt := f.GetFontSize()
	lineHt *= 1.5

	f.Ln(3)
	for _, txt := range contents.Content {
		f.Write(lineHt, txt.Text)
	}
	f.Ln(-1)
	f.Ln(4)
}

// H4 write a 4th level header to a gofpdf.Fpdf.
func H4(f *gofpdf.Fpdf, contents model.Contents) {
	f.SetFont("helvetica", "B", 12)
	f.SetTextColor(36, 41, 46)
	_, lineHt := f.GetFontSize()
	lineHt *= 1.5

	f.Ln(4)
	for _, txt := range contents.Content {
		f.Write(lineHt, txt.Text)
	}
	f.Ln(-1)
	f.Ln(3)
}

// H5 write a 5th level header to a gofpdf.Fpdf.
func H5(f *gofpdf.Fpdf, contents model.Contents) {
	f.SetFont("helvetica", "B", 10.5)
	f.SetTextColor(36, 41, 46)
	_, lineHt := f.GetFontSize()
	lineHt *= 1.5

	f.Ln(4)
	for _, txt := range contents.Content {
		f.Write(lineHt, txt.Text)
	}
	f.Ln(-1)
	f.Ln(3)
}

// H6 write a 6th level header to a gofpdf.Fpdf.
func H6(f *gofpdf.Fpdf, contents model.Contents) {
	f.SetFont("helvetica", "B", 10.5)
	f.SetTextColor(106, 115, 125)
	_, lineHt := f.GetFontSize()
	lineHt *= 1.5

	f.Ln(4)
	for _, txt := range contents.Content {
		f.Write(lineHt, txt.Text)
	}
	f.Ln(-1)
	f.Ln(3)
}
