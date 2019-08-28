package render

import (
	"strings"

	"github.com/jung-kurt/gofpdf"
)

const (
	levelOneMargin   = 22.5
	levelTwoMargin   = 27.5
	levelThreeMargin = 32.5
)

// AnyList writes an unordered list with up to three levels of indentation to a gofpdf.Fpdf.
func AnyList(pdf *gofpdf.Fpdf, list List) {
	pdf.SetLeftMargin(levelOneMargin)
	pdf.SetRightMargin(21)
	pdf.SetFont("helvetica", "", 12)
	_, lineHt := pdf.GetFontSize() // need to call this after we call SetFont.
	lineHt = lineHt * 1.5

	for _, item := range list.Items {
		levelOneBullett(pdf, lineHt)
		drawListItemContent(pdf, item.Contents, lineHt)
		pdf.SetLeftMargin(levelOneMargin)
		pdf.Ln(-1) // A negative value indicates the height of the last printed cell.
		pdf.Ln(2)

		if item.HasChildren() {
			pdf.SetLeftMargin(levelTwoMargin)
			lastIndex := len(item.Children.Items) - 1
			for i, item := range item.Children.Items {
				levelTwoBullet(pdf, lineHt)
				drawListItemContent(pdf, item.Contents, lineHt)
				pdf.SetLeftMargin(levelTwoMargin)
				if i != lastIndex {
					pdf.Ln(-1)
					pdf.Ln(2)
				}
				if item.HasChildren() {
					pdf.Ln(-1)
					pdf.Ln(2)
					pdf.SetLeftMargin(levelThreeMargin)
					lastIndex := len(item.Children.Items) - 1
					for i, item := range item.Children.Items {
						levelThreeBullet(pdf, lineHt)
						drawListItemContent(pdf, item.Contents, lineHt)
						pdf.SetLeftMargin(levelThreeMargin)
						if i != lastIndex {
							pdf.Ln(-1)
							pdf.Ln(2)
						}
					}
					pdf.SetLeftMargin(levelTwoMargin)
					pdf.Ln(-1)
					pdf.Ln(2)
				}
			}
			pdf.SetLeftMargin(levelOneMargin)
			pdf.Ln(-1)
			pdf.Ln(2)
		}
	}

	pdf.SetLeftMargin(20)
	pdf.SetRightMargin(20)
	pdf.Ln(4)
}

func levelOneBullett(pdf *gofpdf.Fpdf, lineHt float64) {
	pdf.SetTextColor(36, 41, 46)
	pdf.SetFont("zapfdingbats", "", 5)
	bulletChar := "\x6c" // ●
	pdf.CellFormat(5, lineHt, bulletChar, "", 0, "RM", false, 0, "")
	pdf.SetLeftMargin(pdf.GetX())
}

func levelTwoBullet(pdf *gofpdf.Fpdf, lineHt float64) {
	pdf.SetTextColor(36, 41, 46)
	pdf.SetFont("zapfdingbats", "", 5)
	bulletChar := "\x6d" // ❍
	pdf.CellFormat(5, lineHt, bulletChar, "", 0, "RM", false, 0, "")
	pdf.SetLeftMargin(pdf.GetX())
}

func levelThreeBullet(pdf *gofpdf.Fpdf, lineHt float64) {
	pdf.SetTextColor(36, 41, 46)
	pdf.SetFont("zapfdingbats", "", 5)
	bulletChar := "\x6e" // ■
	pdf.CellFormat(5, lineHt, bulletChar, "", 0, "RM", false, 0, "")
	pdf.SetLeftMargin(pdf.GetX())
}

func drawListItemContent(pdf *gofpdf.Fpdf, c Contents, lineHt float64) {
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
			liSpan(pdf, txt, styleStr, lineHt)
		}
	}
}

func liSpan(pdf *gofpdf.Fpdf, txt Text, styleStr string, lineHt float64) {
	pdf.SetFont("helvetica", styleStr, 12)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(36, 41, 46)
	if txt.HREF != "" {
		pdf.SetTextColor(3, 102, 214)
		pdf.WriteLinkString(lineHt, txt.Content, txt.HREF)
	} else {
		pdf.SetTextColor(36, 41, 46)
		pdf.Write(lineHt, txt.Content)
	}
}

func liInlineCode(pdf *gofpdf.Fpdf, txt Text, styleStr string) {
	pdf.SetFont("courier", styleStr, 12)
	pdf.SetFillColor(255, 255, 255)

	x := pdf.GetX()
	pageWidth, _ := pdf.GetPageSize()
	_, _, rightMargin, _ := pdf.GetMargins()
	strWdth := pdf.GetStringWidth(txt.Content)
	if x+strWdth > pageWidth-rightMargin {
		pdf.Ln(-1)
	}

	if txt.HREF != "" {
		pdf.SetTextColor(3, 102, 214)
		pdf.WriteLinkString(6, txt.Content, txt.HREF)
	} else {
		pdf.SetTextColor(36, 41, 46)
		pdf.Write(6, txt.Content)
	}
}

func liStrike(pdf *gofpdf.Fpdf, txt Text, styleStr string) {
	pdf.SetFont("helvetica", styleStr, 12)
	pdf.SetFillColor(255, 255, 255)
	width := pdf.GetStringWidth(txt.Content)
	var lineHt float64 = 6

	// Where the text starts, also where to start the strikethrough line.
	x := pdf.GetX()
	y := (pdf.GetY() + (lineHt / 2))

	pdf.SetLineWidth(0.25)
	pdf.SetDrawColor(36, 41, 46)

	if txt.HREF != "" {
		pdf.SetTextColor(3, 102, 214)
		pdf.WriteLinkString(lineHt, txt.Content, txt.HREF)
	} else {
		pdf.SetTextColor(36, 41, 46)
		pdf.Write(lineHt, txt.HREF)
	}

	pdf.Line(x, y, x+width, y)
}
