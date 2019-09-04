package render

import (
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/lib"
	"github.com/tcd/md2pdf/internal/model"
)

func drawContent(pdf *gofpdf.Fpdf, c model.Contents, fontSize float64) {
	r, g, b := pdf.GetFillColor()
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
			strike(pdf, txt, styleStr, fontSize)
		} else if txt.Code {
			inlineCode(pdf, txt, styleStr, fontSize)
		} else {
			span(pdf, txt, styleStr, fontSize)
		}
	}
	pdf.SetFillColor(r, g, b)
}

func span(pdf *gofpdf.Fpdf, txt model.Text, styleStr string, fontSize float64) {
	pdf.SetFillColor(DefaultBG())
	pdf.SetFont("helvetica", styleStr, fontSize)
	_, lineHt := pdf.GetFontSize()
	lineHt *= 1.5
	if txt.HREF != "" {
		pdf.SetTextColor(LinkFG())
		pdf.WriteLinkString(lineHt, txt.Text, txt.HREF)
	} else {
		pdf.SetTextColor(DefaultFG())
		pdf.Write(lineHt, txt.Text)
	}
}

func inlineCode(pdf *gofpdf.Fpdf, txt model.Text, styleStr string, fontSize float64) {
	pdf.SetFillColor(DefaultBG())
	pdf.SetFont("courier", styleStr, fontSize)
	_, lineHt := pdf.GetFontSize()
	lineHt *= 1.5

	// Go to the next line before writing if the code span is too long.
	x := pdf.GetX()
	strWdth := pdf.GetStringWidth(txt.Text)
	if x+strWdth > lib.ContentBoxRight(pdf) {
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

func strike(pdf *gofpdf.Fpdf, txt model.Text, styleStr string, fontSize float64) {
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
		pdf.SetTextColor(DefaultFG())
		pdf.Write(lineHt, txt.Text)
	}

	pdf.Line(x, y, x+width, y)
}
