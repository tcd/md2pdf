package render

import (
	"strings"

	"github.com/jung-kurt/gofpdf"
)

// BasicBlockquote writes text to a blockquote without rendering any enclosed elements.
func BasicBlockquote(f *gofpdf.Fpdf, text string) {
	oldCellMargin := f.GetCellMargin()

	f.SetFont("helvetica", "", 12)
	f.SetTextColor(106, 115, 125)
	f.SetDrawColor(223, 226, 229)
	f.SetLineWidth(1)
	f.SetCellMargin(4)

	_, lineHeight := f.GetFontSize()
	f.MultiCell(0, lineHeight*1.5, text, "LM", "", false)
	f.Ln(6)

	f.SetCellMargin(oldCellMargin)
}

// Blockquote does what BasicBlockquote don't.
func Blockquote(f *gofpdf.Fpdf, contents Contents) {
	allContent := contents.AllContent()
	text := strings.TrimSpace(strings.Join(allContent, ""))
	BasicBlockquote(f, text)
}
