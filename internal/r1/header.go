package render

import (
	"strings"

	gofpdf "github.com/tcd/gofpdf-1"
	"github.com/tcd/md2pdf/internal/model"
)

// h1 write a 1st level header to a gofpdf.Fpdf.
func h1(pdf *gofpdf.Fpdf, contents model.Contents) {
	// pdf.SetFont("helvetica", "B", 24)
	pdf.SetTextColor(DefaultFG())
	// _, lineHt := pdf.GetFontSize()
	// lineHt *= 1.7

	drawHeaderContent(pdf, contents, 24)

	drawHR(pdf, 0.09)
	// pdf.Ln(-1)
	// pdf.Ln(6)
}

// h2 write a 2nd level header to a gofpdf.Fpdf.
func h2(pdf *gofpdf.Fpdf, contents model.Contents) {
	// pdf.SetFont("helvetica", "B", 18)
	pdf.SetTextColor(DefaultFG())
	// _, lineHt := pdf.GetFontSize()
	// lineHt *= 1.6

	// pdf.Ln(3)
	drawHeaderContent(pdf, contents, 18)
	// pdf.Ln(-1)

	drawHR(pdf, 0.09)
	// pdf.Ln(7)
}

// h3 write a 3rd level header to a gofpdf.Fpdf.
func h3(pdf *gofpdf.Fpdf, contents model.Contents) {
	// pdf.SetFont("helvetica", "B", 15)
	pdf.SetTextColor(DefaultFG())
	// _, lineHt := pdf.GetFontSize()
	// lineHt *= 1.5

	// pdf.Ln(3)
	drawHeaderContent(pdf, contents, 15)
	// pdf.Ln(-1)
	// pdf.Ln(4)
}

// h4 write a 4th level header to a gofpdf.Fpdf.
func h4(pdf *gofpdf.Fpdf, contents model.Contents) {
	// pdf.SetFont("helvetica", "B", 12)
	pdf.SetTextColor(DefaultFG())
	// _, lineHt := pdf.GetFontSize()
	// lineHt *= 1.5

	// pdf.Ln(2)
	drawHeaderContent(pdf, contents, 12)
	// pdf.Ln(-1)
	// pdf.Ln(3)
}

// h5 write a 5th level header to a gofpdf.Fpdf.
func h5(pdf *gofpdf.Fpdf, contents model.Contents) {
	// pdf.SetFont("helvetica", "B", 10.5)
	pdf.SetTextColor(DefaultFG())
	// _, lineHt := pdf.GetFontSize()
	// lineHt *= 1.5

	// pdf.Ln(2)
	drawHeaderContent(pdf, contents, 10.5)
	// pdf.Ln(-1)
	// pdf.Ln(3)
}

// h6 write a 6th level header to a gofpdf.Fpdf.
func h6(pdf *gofpdf.Fpdf, contents model.Contents) {
	// pdf.SetFont("helvetica", "B", 10.5)
	pdf.SetTextColor(106, 115, 125)
	// _, lineHt := pdf.GetFontSize()
	// lineHt *= 1.5

	// pdf.Ln(2)
	drawHeaderContent(pdf, contents, 10.5)
	// pdf.Ln(-1)
	// pdf.Ln(3)
}

// Like drawContent, but everything is bold by default.
func drawHeaderContent(pdf *gofpdf.Fpdf, c model.Contents, fontSize float64) {
	for _, txt := range c.Content {
		if txt.Text == "" {
			continue
		}

		var styles strings.Builder
		styles.WriteRune('B')
		if txt.Italic {
			styles.WriteRune('I')
		}
		styleStr := styles.String()

		if txt.Code {
			inlineCode(pdf, txt, styleStr, fontSize)
		} else if txt.Strike {
			strike(pdf, txt, styleStr, fontSize)
		} else {
			span(pdf, txt, styleStr, fontSize)
		}
	}
}
