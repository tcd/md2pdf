package render

import "github.com/jung-kurt/gofpdf"

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
