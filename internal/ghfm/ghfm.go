// Package ghfm renders text in the style of GitHub using gofpdf.
package ghfm

import "github.com/jung-kurt/gofpdf"

const (
	pageWidth = 176 // (216 - 40 for margins)
)

// Setup sets default GitHub-Flavored styles to a gofpdf.Fpdf.
func Setup(f *gofpdf.Fpdf) {
	f.AddPage()
	f.SetFont("helvetica", "", 16)
	f.SetMargins(20, 13, 20)
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

// HR writes a Horizontal Rule to a gofpdf.Fpdf.
func HR(f *gofpdf.Fpdf) {
	x := f.GetX()
	y := f.GetY()
	f.SetLineWidth(1)
	f.SetDrawColor(191, 191, 191)
	f.Line(x, y, x+pageWidth, y)
	f.Ln(2)
}

// CodeBlock writes a `<pre>` style code block to a gofpdf.Fpdf.
// TODO: syntax highlighting..............
func CodeBlock(f *gofpdf.Fpdf, text string) {
	oldCellMargin := f.GetCellMargin()

	f.SetFont("courier", "", 11)
	f.SetFillColor(246, 248, 250)
	f.SetCellMargin(4)

	f.CellFormat(0, 4, "", "", 1, "TL", true, 0, "")
	f.MultiCell(0, 5, text, "", "", true)
	f.CellFormat(0, 4, "", "", 1, "TL", true, 0, "")
	f.Ln(6)

	f.SetCellMargin(oldCellMargin)
}

// BasicP writes the contents of a paragraph without rendering
// any enclosed elements.
func BasicP(f *gofpdf.Fpdf, text string) {
	if len(text) == 0 {
		return
	}
	f.SetFont("helvetica", "", 12)
	f.SetFillColor(255, 255, 255)
	f.SetTextColor(36, 41, 46)
	f.Write(6, text)
	f.Ln(11)
}

// BasicBlockquote writes text to a blockquote without rendering
// any enclosed elements.
func BasicBlockquote(f *gofpdf.Fpdf, text string) {
	oldCellMargin := f.GetCellMargin()

	f.SetFont("helvetica", "", 12)
	f.SetTextColor(106, 115, 125)
	f.SetDrawColor(223, 226, 229)
	f.SetLineWidth(1)
	f.SetCellMargin(4)

	_, lineHeight := f.GetFontSize()
	f.MultiCell(0, lineHeight*1.5, text, "L", "", false)
	f.Ln(6)

	f.SetCellMargin(oldCellMargin)
}
