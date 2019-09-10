package render

import (
	"fmt"
	"log"
	"strings"

	gofpdf "github.com/tcd/gofpdf-1"
	"github.com/tcd/md2pdf/internal/model"
)

func drawContent(pdf *gofpdf.Fpdf, c model.Contents, fontSize float64) {
	// r, g, b := pdf.GetFillColor()
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
	pdf.Ln(0)
	// pdf.SetRGBFillColor(r, g, b)
}

func span(pdf *gofpdf.Fpdf, txt model.Text, styleStr string, fontSize float64) {
	// pdf.SetRGBFillColor(DefaultBG())
	// pdf.SetFont("helvetica", styleStr, fontSize)
	err := pdf.SetFont("helvetica", "", fontSize)
	if err != nil {
		log.Println(err)
	}

	lineHt := (fontSize * (25.4 / 72)) * 1.5
	// _, lineHt := pdf.GetFontSize()
	// lineHt *= 1.5
	// width, err := pdf.MeasureTextWidth(txt.Text, gofpdf.TextOption{})
	// if err != nil {
	// 	log.Println(err)
	// }
	// if txt.HREF != "" {
	// 	pdf.SetTextColor(LinkFG())
	// pdf.AddExternalLink(txt.HREF, pdf.X(), pdf.Y(), width, lineHt)
	// } else {
	// pdf.SetTextColor(DefaultFG())
	fmt.Println(pdf.X(), pdf.Y())
	if pdf.X() > pdf.MarginRight() {
	}
	err = pdf.WriteText(lineHt, txt.Text)
	// err = pdf.Cell(width, lineHt, txt.Text)
	pdf.Ln(10)
	// err = pdf.WriteTextOpts(lineHt, txt.Text, gofpdf.CellOption{Float: gofpdf.Right}, gofpdf.TextOption{Stroke: true})
	// err = pdf.WriteText((fontSize * (25.4 / 72)), txt.Text)
	if err != nil {
		log.Println(err)
	}
}

func inlineCode(pdf *gofpdf.Fpdf, txt model.Text, styleStr string, fontSize float64) {
	// pdf.SetRGBFillColor(DefaultBG())
	pdf.SetFont("courier", "", fontSize)

	// width, err := pdf.MeasureTextWidth(txt.Text, gofpdf.TextOption{})
	// if err != nil {
	// 	log.Println(err)
	// }
	// _, lineHt := pdf.GetFontSize()
	// lineHt *= 1.5

	// if txt.HREF != "" {
	// 	pdf.SetTextColor(LinkFG())
	// 	pdf.AddExternalLink(txt.HREF, pdf.X(), pdf.Y(), width, lineHt)
	// } else {
	pdf.SetTextColor(CodeOrangeFG())
	pdf.WriteText(fontSize, txt.Text)
	// }
}

func strike(pdf *gofpdf.Fpdf, txt model.Text, styleStr string, fontSize float64) {
	pdf.SetRGBFillColor(DefaultBG())
	pdf.SetRGBStrokeColor(DefaultFG())
	pdf.SetFont("helvetica", styleStr, fontSize)
	// width, err := pdf.MeasureTextWidth(txt.Text, gofpdf.TextOption{})
	// if err != nil {
	// 	log.Println(err)
	// }
	// _, lineHt := pdf.GetFontSize()
	// lineHt *= 1.5

	// Where the text starts, also where to start the strikethrough line.
	// x := pdf.X()
	// y := (pdf.Y() + (lineHt / 2))
	pdf.SetLineWidth(0.25)

	// if txt.HREF != "" {
	// 	pdf.SetTextColor(LinkFG())
	// 	pdf.AddExternalLink(txt.HREF, pdf.X(), pdf.Y(), width, lineHt)
	// } else {
	pdf.SetTextColor(DefaultFG())
	pdf.WriteText(fontSize, txt.Text)
	// }

	// pdf.Line(x, y, x+width, y)
}
