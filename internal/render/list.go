package render

import (
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/model"
)

const (
	levelOneMargin   = 22.5
	levelTwoMargin   = 27.5
	levelThreeMargin = 32.5
)

var margins = map[int]float64{
	0: 22.5,
	1: 22.5,
	2: 27.5,
	3: 32.5,
}

// List writes a list with any level of indentation to a gofpdf.Fpdf.
func List(pdf *gofpdf.Fpdf, list model.ListContent) {
	level := 1
	pdf.SetLeftMargin(margins[level])
	pdf.SetRightMargin(21)
	_, lineHt := pdf.GetFontSize() // need to call this after we call SetFont.
	// lineHt = lineHt * 1.5

	for i, item := range list.Items {
		if list.Ordered {
			drawNumbering(pdf, lineHt, level, i+1)
		} else {
			drawBullet(pdf, lineHt, level)
		}
		drawListItemContent(pdf, item.Contents)
		pdf.SetLeftMargin(margins[level])
		pdf.Ln(-1) // A negative value indicates the height of the last printed cell.
		pdf.Ln(2)
		if item.HasChildren() {
			drawList(pdf, item.Children, level+1)
		}
	}

	pdf.SetLeftMargin(20)
	pdf.SetRightMargin(20)
	pdf.Ln(4)
}

// Draw all items in a ListContent, anc call drawList for any sublists.
func drawList(pdf *gofpdf.Fpdf, list model.ListContent, level int) {
	pdf.SetLeftMargin(margins[level])
	pdf.SetFont("helvetica", "", 12)
	_, lineHt := pdf.GetFontSize() // need to call this after we call SetFont.
	// lineHt = lineHt * 1.5

	lastIndex := len(list.Items) - 1
	for i, item := range list.Items {
		if list.Ordered {
			drawNumbering(pdf, lineHt, level, i+1)
		} else {
			drawBullet(pdf, lineHt, level)
		}
		drawListItemContent(pdf, item.Contents)
		pdf.SetLeftMargin(margins[level])
		if i != lastIndex {
			pdf.Ln(-1)
			pdf.Ln(2)
		}
		if item.HasChildren() {
			drawList(pdf, item.Children, level+1)
		}
	}
	pdf.SetLeftMargin(margins[level-1])
	pdf.Ln(-1)
	pdf.Ln(2)
}

func drawListItemContent(pdf *gofpdf.Fpdf, c model.Contents) {
	for _, txt := range c.Content {
		var styles strings.Builder // "B" (bold), "I" (italic), "U" (underscore) or any combination.
		if txt.Bold {
			styles.WriteRune('B')
		}
		if txt.Italic {
			styles.WriteRune('I')
		}
		styleStr := styles.String()

		if txt.Strike {
			liStrike(pdf, txt, styleStr)
		} else if txt.Code {
			liInlineCode(pdf, txt, styleStr)
		} else {
			liSpan(pdf, txt, styleStr)
		}
	}
}

func liSpan(pdf *gofpdf.Fpdf, txt model.Text, styleStr string) {
	pdf.SetFont("helvetica", styleStr, 12)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(36, 41, 46)
	_, lineHt := pdf.GetFontSize()
	lineHt *= 1.5
	if txt.HREF != "" {
		pdf.SetTextColor(3, 102, 214)
		pdf.WriteLinkString(lineHt, txt.Text, txt.HREF)
	} else {
		pdf.SetTextColor(36, 41, 46)
		pdf.Write(lineHt, txt.Text)
	}
}

func liInlineCode(pdf *gofpdf.Fpdf, txt model.Text, styleStr string) {
	pdf.SetFont("courier", styleStr, 12)
	pdf.SetFillColor(255, 255, 255)
	var lineHt float64 = 6

	x := pdf.GetX()
	pageWidth, _ := pdf.GetPageSize()
	_, _, rightMargin, _ := pdf.GetMargins()
	strWdth := pdf.GetStringWidth(txt.Text)
	if x+strWdth > pageWidth-rightMargin {
		pdf.Ln(-1)
	}

	if txt.HREF != "" {
		pdf.SetTextColor(3, 102, 214)
		pdf.WriteLinkString(lineHt, txt.Text, txt.HREF)
	} else {
		// pdf.SetTextColor(36, 41, 46)
		pdf.SetTextColor(227, 98, 9)
		pdf.Write(lineHt, txt.Text)
	}
}

func liStrike(pdf *gofpdf.Fpdf, txt model.Text, styleStr string) {
	pdf.SetFont("helvetica", styleStr, 12)
	pdf.SetFillColor(255, 255, 255)
	width := pdf.GetStringWidth(txt.Text)
	var lineHt float64 = 6

	// Where the text starts, also where to start the strikethrough line.
	x := pdf.GetX()
	y := (pdf.GetY() + (lineHt / 2))

	pdf.SetLineWidth(0.25)
	pdf.SetDrawColor(36, 41, 46)

	if txt.HREF != "" {
		pdf.SetTextColor(3, 102, 214)
		pdf.WriteLinkString(lineHt, txt.Text, txt.HREF)
	} else {
		pdf.SetTextColor(36, 41, 46)
		pdf.Write(lineHt, txt.HREF)
	}

	pdf.Line(x, y, x+width, y)
}
