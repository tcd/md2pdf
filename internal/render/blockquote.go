package render

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
	pdf.Ln(5)

	pdf.SetCellMargin(oldCellMargin)
}

// Blockquote does what BasicBlockquote don't.
func Blockquote(pdf *gofpdf.Fpdf, contents model.Contents) {
	BasicBlockquote(pdf, contents.JoinContent())
}

func mockBlockquote(pdf *gofpdf.Fpdf, text string) {

}
