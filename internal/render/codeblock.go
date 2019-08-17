package render

import "github.com/jung-kurt/gofpdf"

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
