package render

import (
	"strings"

	"github.com/jung-kurt/gofpdf"
)

// BasicP writes the contents of a paragraph without rendering
// any enclosed elements.
func BasicP(f *gofpdf.Fpdf, text string) {
	if len(text) == 0 {
		return
	}
	f.SetFont("helvetica", "", 12)
	f.SetFillColor(255, 255, 255)
	f.SetTextColor(36, 41, 46)
	f.Write(6, text)
	f.Ln(10)
}

// FullP will (when finished) write a paragraph with plain, bold, italic, bold/italic, and linked text.
func FullP(f *gofpdf.Fpdf, content Contents) {
	if len(content.Content) == 0 {
		return
	}

	f.SetFont("helvetica", "", 12)
	f.SetFillColor(255, 255, 255)
	f.SetTextColor(36, 41, 46)

	drawParagraphContent(f, content)

	f.Ln(10)
}

func drawParagraphContent(f *gofpdf.Fpdf, c Contents) {
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
			pStrike(f, txt)
		} else if txt.Code {
			pInlineCode(f, txt)
		} else {
			pSpan(f, txt, styleStr)
		}
	}
}

func pSpan(f *gofpdf.Fpdf, txt Text, styleStr string) {
	f.SetFont("helvetica", styleStr, 12)
	f.SetFillColor(255, 255, 255)
	if txt.HREF != "" {
		f.SetTextColor(3, 102, 214)
		f.WriteLinkString(6, txt.Content, txt.HREF)
	} else {
		f.SetTextColor(36, 41, 46)
		f.Write(6, txt.Content)
	}

}

func pInlineCode(f *gofpdf.Fpdf, txt Text) {
	oldCellMargin := f.GetCellMargin()
	f.SetFillColor(243, 243, 243)
	f.SetFont("courier", "", 12)
	width := (f.GetStringWidth(txt.Content) * 1.05)

	if txt.HREF != "" {
		f.SetTextColor(3, 102, 214)
		f.CellFormat(
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
		f.CellFormat(
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

	f.SetCellMargin(oldCellMargin)
}

// Strike writes a line of text with a line through it.
func pStrike(f *gofpdf.Fpdf, txt Text) {
	f.SetFont("helvetica", "", 12)
	f.SetFillColor(255, 255, 255)
	f.SetTextColor(36, 41, 46)
	width := f.GetStringWidth(txt.Content)
	var lineHt float64 = 6

	// Where the text starts, also where to start the strikethrough line.
	x := f.GetX()
	y := (f.GetY() + (lineHt / 2))

	f.SetLineWidth(0.25)
	f.SetDrawColor(36, 41, 46)

	if txt.HREF != "" {
		f.SetTextColor(3, 102, 214)
		f.WriteLinkString(lineHt, txt.Content, txt.HREF)
	} else {
		f.Write(lineHt, txt.Content)
	}

	f.Line(x, y, x+width, y)
}
