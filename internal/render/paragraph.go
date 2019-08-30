package render

import (
	"strings"

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

// FullP will (when finished) write a paragraph with plain, bold, italic, bold/italic, and linked text.
func FullP(pdf *gofpdf.Fpdf, contents model.Contents) {
	if len(contents.Content) == 0 {
		return
	}
	if len(contents.AllContent()) == 0 {
		return
	}

	drawParagraphContent(pdf, contents)
	pdf.Ln(10)
}

func drawParagraphContent(pdf *gofpdf.Fpdf, c model.Contents) {
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
			pStrike(pdf, txt)
		} else if txt.Code {
			pInlineCode(pdf, txt, styleStr)
		} else {
			pSpan(pdf, txt, styleStr)
		}
	}
}

func pSpan(pdf *gofpdf.Fpdf, txt model.Text, styleStr string) {
	pdf.SetFont("helvetica", styleStr, 12)
	_, lineHt := pdf.GetFontSize()
	lineHt *= 1.5
	pdf.SetFillColor(DefaultBG())
	if txt.HREF != "" {
		pdf.SetTextColor(LinkFG())
		pdf.WriteLinkString(lineHt, txt.Text, txt.HREF)
	} else {
		pdf.SetTextColor(DefaultFG())
		pdf.Write(lineHt, txt.Text)
	}
}

func pInlineCode(pdf *gofpdf.Fpdf, txt model.Text, styleStr string) {
	pdf.SetFont("courier", styleStr, 12)
	_, lineHt := pdf.GetFontSize()
	lineHt *= 1.5
	pdf.SetFillColor(DefaultBG())

	x := pdf.GetX()
	pageWidth, _ := pdf.GetPageSize()
	_, _, rightMargin, _ := pdf.GetMargins()
	strWdth := pdf.GetStringWidth(txt.Text)
	if x+strWdth > pageWidth-rightMargin {
		pdf.Ln(-1)
	}

	if txt.HREF != "" {
		pdf.SetTextColor(LinkFG())
		pdf.WriteLinkString(lineHt, txt.Text, txt.HREF)
	} else {
		pdf.SetTextColor(CodeOrangeFG())
		pdf.Write(lineHt, txt.Text)
	}
}

// Strike writes a line of text with a line through it.
func pStrike(pdf *gofpdf.Fpdf, txt model.Text) {
	pdf.SetFont("helvetica", "", 12)
	pdf.SetFillColor(DefaultBG())
	pdf.SetTextColor(DefaultFG())
	width := pdf.GetStringWidth(txt.Text)
	_, lineHt := pdf.GetFontSize()
	lineHt *= 1.5

	// Where the text starts, also where to start the strikethrough line.
	x := pdf.GetX()
	y := (pdf.GetY() + (lineHt / 2))

	pdf.SetLineWidth(0.25)
	pdf.SetDrawColor(DefaultFG())

	if txt.HREF != "" {
		pdf.SetTextColor(LinkFG())
		pdf.WriteLinkString(lineHt, txt.Text, txt.HREF)
	} else {
		pdf.Write(lineHt, txt.Text)
	}

	pdf.Line(x, y, x+width, y)
}
