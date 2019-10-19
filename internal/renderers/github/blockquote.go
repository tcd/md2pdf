package github

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/model"
)

// BasicBlockquote writes text to a blockquote without rendering any enclosed elements.
func BasicBlockquote(pdf *gofpdf.Fpdf, text string) {
	oldCellMargin := pdf.GetCellMargin()

	pdf.SetFont("helvetica", "", 12)
	pdf.SetTextColor(BlockquoteFG())
	pdf.SetDrawColor(BlockquoteBorder())
	pdf.SetLineWidth(1)
	pdf.SetCellMargin(4)

	_, lineHeight := pdf.GetFontSize()
	pdf.MultiCell(0, lineHeight*1.5, text, "LM", "", false)
	pdf.Ln(4)

	pdf.SetCellMargin(oldCellMargin)
}

// Blockquote does what BasicBlockquote don't.
func Blockquote(pdf *gofpdf.Fpdf, contents model.Contents) {
	oldLeftMargin, _, _, _ := pdf.GetMargins()

	pdf.SetFont("helvetica", "", 12)
	pdf.SetTextColor(BlockquoteFG())
	pdf.SetDrawColor(BlockquoteBorder())
	pdf.SetLineWidth(1)

	x, y := pdf.GetXY()
	height := mockBlockquote(pdf, contents.JoinContent())
	pdf.Line(x, y, x, y+height)

	pdf.SetLeftMargin(oldLeftMargin + 4)
	drawContent(pdf, contents, 12)

	pdf.SetLeftMargin(oldLeftMargin)
	pdf.Ln(-1)
	pdf.Ln(4)
}

// Returns the height needed to hold blockquote contents.
func mockBlockquote(mainPdf *gofpdf.Fpdf, text string) float64 {
	pdf := gofpdf.New("P", "mm", "Letter", "")
	Setup(pdf)
	pdf.SetXY(mainPdf.GetXY())

	y1 := pdf.GetY()
	pdf.SetFont("helvetica", "", 12)
	pdf.SetLineWidth(1)
	pdf.SetCellMargin(4)
	_, lineHeight := pdf.GetFontSize()
	pdf.MultiCell(0, lineHeight*1.5, text, "LM", "", false)
	y2 := pdf.GetY()

	return y2 - y1
}
