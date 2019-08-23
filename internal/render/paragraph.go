package render

import (
	"strings"

	"github.com/jung-kurt/gofpdf"
)

// BasicP writes the contents of a paragraph without rendering
// any enclosed elements.
func BasicP(pdf *gofpdf.Fpdf, text string) {
	if len(text) == 0 {
		return
	}
	pdf.SetFont("helvetica", "", 12)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(36, 41, 46)
	pdf.Write(6, text)
	pdf.Ln(10)
}

// FullP will (when finished) write a paragraph with plain, bold, italic, bold/italic, and linked text.
func FullP(pdf *gofpdf.Fpdf, content Contents) {
	if len(content.Content) == 0 {
		return
	}

	pdf.SetFont("helvetica", "", 12)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(36, 41, 46)

	drawParagraphContent(pdf, content)

	pdf.Ln(10)
}

func drawParagraphContent(pdf *gofpdf.Fpdf, c Contents) {
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
			pStrike(pdf, txt)
		} else if txt.Code {
			pInlineCode(pdf, txt)
		} else {
			pSpan(pdf, txt, styleStr)
		}
	}
}

func pSpan(pdf *gofpdf.Fpdf, txt Text, styleStr string) {
	pdf.SetFont("helvetica", styleStr, 12)
	pdf.SetFillColor(255, 255, 255)
	if txt.HREF != "" {
		pdf.SetTextColor(3, 102, 214)
		pdf.WriteLinkString(6, txt.Content, txt.HREF)
	} else {
		pdf.SetTextColor(36, 41, 46)
		pdf.Write(6, txt.Content)
	}

}

func pInlineCode(pdf *gofpdf.Fpdf, txt Text) {
	oldCellMargin := pdf.GetCellMargin()
	pdf.SetFillColor(243, 243, 243)
	pdf.SetFont("courier", "", 12)
	width := (pdf.GetStringWidth(txt.Content) * 1.05)

	if txt.HREF != "" {
		pdf.SetTextColor(3, 102, 214)
		pdf.CellFormat(
			width,       // width; If 0, the cell extends up to the right margin.
			6,           // height;
			txt.Content, // txtStr;
			"",          // borderStr; "" (no border), "1" (full border), "L", "T", "R" and "B" (left, top, right and bottom)
			0,           // ln; 0 (to the right), 1 (to the beginning of the next line, like Ln()), and 2 (below).
			"LM",        // alignStr; Default is left middle. "L", "C" or "R" (left, center, right) + "T", "M", "B" or "A" (top, middle, bottom, baseline)
			true,        // fill; true for color, false for transparent
			0,           // link;  Identifier returned by AddLink() or 0 for no internal link.
			txt.HREF,    // linkStr; A target URL or empty for no external link. A non-zero value for link takes precedence over linkStr.
		)
	} else {
		pdf.CellFormat(
			width,       // width; If 0, the cell extends up to the right margin.
			6,           // height;
			txt.Content, // txtStr;
			"",          // borderStr; "" (no border), "1" (full border), "L", "T", "R" and "B" (left, top, right and bottom)
			0,           // ln; 0 (to the right), 1 (to the beginning of the next line, like Ln()), and 2 (below).
			"LM",        // alignStr; Default is left middle. "L", "C" or "R" (left, center, right) + "T", "M", "B" or "A" (top, middle, bottom, baseline)
			true,        // fill; true for color, false for transparent
			0,           // link;  Identifier returned by AddLink() or 0 for no internal link.
			"",          // linkStr; A target URL or empty for no external link. A non-zero value for link takes precedence over linkStr.
		)
	}

	pdf.SetCellMargin(oldCellMargin)
}

// Strike writes a line of text with a line through it.
func pStrike(pdf *gofpdf.Fpdf, txt Text) {
	pdf.SetFont("helvetica", "", 12)
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(36, 41, 46)
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
		pdf.Write(lineHt, txt.Content)
	}

	pdf.Line(x, y, x+width, y)
}
