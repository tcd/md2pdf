package render

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/model"
)

// BasicP writes the contents of a paragraph without rendering
// any enclosed elements.
func BasicP(pdf *gofpdf.Fpdf, text string) {
	if len(text) == 0 {
		return
	}
	pdf.SetFont("helvetica", "", 12)
	pdf.SetFillColor(DefaultBG())
	pdf.SetTextColor(DefaultFG())
	pdf.Write(6, text)
	pdf.Ln(10)
}

// FullP will write a paragraph with plain, bold, italic, bold/italic, and linked text.
func FullP(pdf *gofpdf.Fpdf, contents model.Contents) {
	if len(contents.Content) == 0 {
		return
	}
	if len(contents.AllContent()) == 0 {
		return
	}
	drawContent(pdf, contents)
	pdf.Ln(10)
}
