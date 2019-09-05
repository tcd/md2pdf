package render

import (
	"strings"

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
	drawBlockquoteContent(pdf, contents, 12)

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

func drawBlockquoteContent(pdf *gofpdf.Fpdf, c model.Contents, fontSize float64) {
	for _, txt := range c.Content {
		if txt.Text == "" {
			continue
		}
		var styles strings.Builder // "B" (bold), "I" (italic), "U" (underscore) or any combination.
		if txt.Bold {
			styles.WriteRune('B')
		}
		if txt.Italic {
			styles.WriteRune('I')
		}
		styleStr := styles.String()

		if txt.Strike {
			strikeBQ(pdf, txt, styleStr, fontSize)
		} else {
			spanBQ(pdf, txt, styleStr, fontSize)
		}
	}
}

func spanBQ(pdf *gofpdf.Fpdf, txt model.Text, styleStr string, fontSize float64) {
	pdf.SetFillColor(DefaultBG())
	pdf.SetFont("helvetica", styleStr, fontSize)
	_, lineHt := pdf.GetFontSize()
	lineHt *= 1.5
	if txt.HREF != "" {
		pdf.SetTextColor(LinkFG())
		pdf.WriteLinkString(lineHt, txt.Text, txt.HREF)
	} else {
		pdf.SetTextColor(BlockquoteFG())
		pdf.Write(lineHt, txt.Text)
	}
}

func strikeBQ(pdf *gofpdf.Fpdf, txt model.Text, styleStr string, fontSize float64) {
	pdf.SetFillColor(DefaultBG())
	pdf.SetDrawColor(DefaultFG())
	pdf.SetFont("helvetica", styleStr, fontSize)
	width := pdf.GetStringWidth(txt.Text)
	_, lineHt := pdf.GetFontSize()
	lineHt *= 1.5

	// Where the text starts, also where to start the strikethrough line.
	x := pdf.GetX()
	y := (pdf.GetY() + (lineHt / 2))
	pdf.SetLineWidth(0.25)

	if txt.HREF != "" {
		pdf.SetTextColor(LinkFG())
		pdf.WriteLinkString(lineHt, txt.Text, txt.HREF)
	} else {
		pdf.SetTextColor(BlockquoteFG())
		pdf.Write(lineHt, txt.Text)
	}

	pdf.Line(x, y, x+width, y)
}
