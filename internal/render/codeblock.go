package render

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

// CodeBlock writes a `<pre>` style code block to a gofpdf.Fpdf.
func CodeBlock(pdf *gofpdf.Fpdf, contents Contents) {
	if len(contents.Content) != 1 {
		fmt.Println("CodeBlock: invalid argument.")
	}

	content := contents.Content[0].Content

	oldCellMargin := pdf.GetCellMargin()

	pdf.SetFont("courier", "", 11)
	pdf.SetFillColor(246, 248, 250)
	pdf.SetCellMargin(4)

	pdf.CellFormat(0, 4, "", "", 1, "TL", true, 0, "")
	pdf.MultiCell(0, 5, content, "", "", true)
	pdf.CellFormat(0, 4, "", "", 0, "TL", true, 0, "")
	pdf.Ln(6)

	pdf.SetCellMargin(oldCellMargin)
}
