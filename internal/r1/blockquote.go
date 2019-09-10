package render

import (
	"log"
	"strings"

	gofpdf "github.com/tcd/gofpdf-1"
	"github.com/tcd/md2pdf/internal/model"
)

// basicBlockquote writes text to a blockquote without rendering any enclosed elements.
func basicBlockquote(pdf *gofpdf.Fpdf, text string) {
	// oldCellMargin := pdf.GetCellMargin()

	err := pdf.SetFont("helvetica", "", 12)
	if err != nil {
		log.Println(err)
	}
	pdf.SetTextColor(BlockquoteFG())
	// pdf.SetDrawColor(BlockquoteBorder())
	// pdf.
	pdf.SetLineWidth(1)
	// pdf.SetCellMargin(4)

	_, lineHeight := pdf.GetFontSize()
	// pdf.MultiCell(0, lineHeight*1.5, text, "LM", "", false)
	pdf.MultiCell(0, lineHeight*1.5, text)
	pdf.Ln(4)

	// pdf.SetCellMargin(oldCellMargin)
}

// blockquote does what BasicBlockquote don't.
func blockquote(pdf *gofpdf.Fpdf, contents model.Contents) {
	// oldLeftMargin, _, _, _ := pdf.GetMargins()

	err := pdf.SetFont("helvetica", "", 12)
	if err != nil {
		log.Println(err)
	}
	pdf.SetTextColor(BlockquoteFG())
	// pdf.SetDrawColor(BlockquoteBorder())
	pdf.SetLineWidth(1)

	x, y := pdf.XY()
	height := mockBlockquote(pdf, contents.JoinContent())
	pdf.Line(x, y, x, y+height)

	// pdf.SetMarginLeft(oldLeftMargin + 4)
	drawBlockquoteContent(pdf, contents, 12)

	// pdf.SetMarginLeft(oldLeftMargin)
	pdf.Ln(-1)
	pdf.Ln(4)
}

// Returns the height needed to hold blockquote contents.
func mockBlockquote(mainPdf *gofpdf.Fpdf, text string) float64 {
	pdf := newPdf()
	pdf.SetXY(mainPdf.XY())

	y1 := pdf.Y()
	err := pdf.SetFont("helvetica", "", 12)
	if err != nil {
		log.Println(err)
	}
	pdf.SetLineWidth(1)
	// pdf.SetCellMargin(4)
	_, lineHeight := pdf.GetFontSize()
	// pdf.MultiCell(0, lineHeight*1.5, text, "LM", "", false)
	pdf.MultiCell(0, lineHeight*1.5, text)
	y2 := pdf.Y()

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
	pdf.SetRGBFillColor(DefaultBG())
	// err := pdf.SetFont(DefaultFont, "", 12)
	// if err != nil {
	// 	log.Println(err)
	// }
	width, err := pdf.MeasureTextWidth(txt.Text, gofpdf.TextOption{})
	if err != nil {
		log.Println(err)
	}
	_, lineHt := pdf.GetFontSize()
	lineHt *= 1.5

	if txt.HREF != "" {
		pdf.SetTextColor(LinkFG())
		pdf.AddExternalLink(txt.HREF, pdf.X(), pdf.Y(), width, lineHt)
	} else {
		pdf.SetTextColor(BlockquoteFG())
		pdf.WriteText(lineHt, txt.Text)
	}
}

func strikeBQ(pdf *gofpdf.Fpdf, txt model.Text, styleStr string, fontSize float64) {
	pdf.SetRGBFillColor(DefaultBG())
	// err := pdf.SetFont(DefaultFont, "", 12)
	// if err != nil {
	// 	log.Println(err)
	// }
	width, err := pdf.MeasureTextWidth(txt.Text, gofpdf.TextOption{})
	if err != nil {
		log.Println(err)
	}
	_, lineHt := pdf.GetFontSize()
	lineHt *= 1.5

	// Where the text starts, also where to start the strikethrough line.
	x := pdf.X()
	y := (pdf.Y() + (lineHt / 2))
	pdf.SetLineWidth(0.25)

	if txt.HREF != "" {
		pdf.SetTextColor(LinkFG())
		pdf.AddExternalLink(txt.HREF, pdf.X(), pdf.Y(), width, lineHt)
	} else {
		pdf.SetTextColor(BlockquoteFG())
		pdf.WriteText(lineHt, txt.Text)
	}

	pdf.Line(x, y, x+width, y)
}
