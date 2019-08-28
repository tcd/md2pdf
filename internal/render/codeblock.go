package render

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/model"
)

// CodeBlock writes a `<pre>` style code block to a gofpdf.Fpdf.
func CodeBlock(pdf *gofpdf.Fpdf, contents model.Contents) {
	if len(contents.Content) != 1 {
		fmt.Println("CodeBlock: invalid argument.")
	}

	content := contents.Content[0].Text

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

// // CodeBlock writes a `<pre>` style code block to a gofpdf.Fpdf.
// func CodeBlock(pdf *gofpdf.Fpdf, contents Contents) {
// 	if len(contents.Content) != 1 {
// 		fmt.Println("CodeBlock: invalid argument.")
// 	}
//
// 	content := contents.Content[0].Content
//
// 	oldCellMargin := pdf.GetCellMargin()
//
// 	pdf.SetFont("courier", "", 11)
// 	pdf.SetFillColor(246, 248, 250)
// 	pdf.SetCellMargin(4)
//
// 	// usableWidth := model.ContentBoxWidth(pdf)
//
// 	// // x1, y1 := pdf.GetXY()
// 	// pdf.CellFormat(0, 0, "", "", 1, "TL", true, 0, "")
// 	// // x2, y2 := pdf.GetXY()
// 	// // pdf.MultiCell(0, 0, content, "", "", false)
// 	// pdf.CellFormat(0, 0, content, "", 1, "", true, 0, "")
// 	// pdf.CellFormat(0, 0, "", "", 1, "TL", true, 0, "")
// 	// pdf.Ln(6)
// 	// // x3, y3 := pdf.GetXY()
//
// 	// pdf.Rect(x1, y1, (x3 - x1), (y3 + y1), "F")
// 	// pdf.SetXY(x2, y2)
// 	pdf.CellFormat(0, 4, "", "", 1, "TL", true, 0, "")
// 	pdf.MultiCell(0, 5, content, "", "", true)
// 	pdf.CellFormat(0, 4, "", "", 0, "TL", true, 0, "")
// 	pdf.Ln(6)
//
// 	pdf.SetCellMargin(oldCellMargin)
// }

// // HighlightedCodeBlock TODO: comment.
// func HighlightedCodeBlock(pdf *gofpdf.Fpdf, contents Contents) {
// 	oldCellMargin := pdf.GetCellMargin()
//
// 	pdf.SetFont("courier", "", 11)
// 	pdf.SetFillColor(246, 248, 250)
// 	pdf.SetCellMargin(4)
//
// 	pdf.CellFormat(0, 4, "", "", 1, "TL", true, 0, "")
// 	for _, txt := range contents.Content {
// 		cbSpan(pdf, txt)
// 	}
// 	pdf.CellFormat(0, 4, "", "", 1, "TL", true, 0, "")
// 	pdf.Ln(6)
//
// 	pdf.SetCellMargin(oldCellMargin)
// }

// func cbSpan(pdf *gofpdf.Fpdf, txt Text) {
// 	r, g, b, err := parseStyle(txt)
// 	if err == nil {
// 		pdf.SetTextColor(r, g, b)
// 	} else {
// 		pdf.SetTextColor(36, 41, 46)
// 	}
//
// 	pdf.SetFont("courier", "", 11)
// 	pdf.SetFillColor(246, 248, 250)
// 	// width := pdf.GetStringWidth(txt.Content)
//
// 	var lineHt float64 = 6
// 	pdf.Write(lineHt, txt.Content)
// 	// pdf.Write()
// 	// pdf.MultiCell(0, lineHt, txt.Content, "", "", true)
// 	// pdf.Cell(width, lineHt, txt.Content)
// 	// pdf.CellFormat(width, lineHt, txt.Content, "", 1, "", true, 0, "")
// }

// func parseStyle(txt Text) (r, g, b int, err error) {
//
// 	if txt.Style == "" {
// 		err = fmt.Errorf("parseStyle: no Style attribute")
// 		return
// 	}
//
// 	re := regexp.MustCompile(`\bcolor:(#\w+)`)
// 	matches := re.FindAllStringSubmatch(txt.Style, -1)
// 	if len(matches) == 0 {
// 		err = fmt.Errorf("parseStyle: no hex code found")
// 		return
// 	}
// 	hexColor := matches[0][1]
//
// 	r, g, b, err = model.HexToRGB(hexColor)
// 	if err != nil {
// 		return
// 	}
//
// 	return
// }
