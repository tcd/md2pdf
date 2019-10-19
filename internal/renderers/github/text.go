package github

import (
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/model"
)

func drawContent(pdf *gofpdf.Fpdf, c model.Contents, fontSize float64) {
	r, g, b := pdf.GetFillColor()
	for _, txt := range c.Content {
		if txt.Text == "" {
			continue
		}
		var styles strings.Builder // "B" (bold), "I" (italic), "U" (underscore), "S" (strike) or any combination.
		if txt.Bold {
			styles.WriteRune('B')
		}
		if txt.Italic {
			styles.WriteRune('I')
		}
		if txt.Strike {
			styles.WriteRune('S')
		}
		styleStr := styles.String()

		if txt.Code {
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

	if txt.HREF != "" {
		pdf.SetTextColor(LinkFG())
		pdf.WriteLinkString(lineHt, txt.Text, txt.HREF)
	} else {
		pdf.SetTextColor(CodeOrangeFG())
		pdf.Write(lineHt, txt.Text)
	}
}
