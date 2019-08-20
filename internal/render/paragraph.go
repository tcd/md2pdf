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
// Still working on strike and inline code.
func FullP(f *gofpdf.Fpdf, text []Text) {
	if len(text) == 0 {
		return
	}

	f.SetFont("helvetica", "", 12)
	var lineHt float64 = 6
	f.SetFillColor(255, 255, 255)
	f.SetTextColor(36, 41, 46)

	for _, t := range text {
		var styles strings.Builder // "B" (bold), "I" (italic), "U" (underscore) or any combination.
		if t.Bold {
			styles.WriteRune('B')
		}
		if t.Italic {
			styles.WriteRune('I')
		}
		styleStr := styles.String()

		if t.Strike {
			Strike(f, t.Content, t.HREF)
		} else if t.Code {
			InlineCode(f, t.Content)
		} else if t.HREF != "" {
			f.SetFont("helvetica", styleStr, 12)
			f.SetTextColor(3, 102, 214)
			f.WriteLinkString(lineHt, t.Content, t.HREF)
		} else {
			Span(f, t.Content, styleStr)
		}
	}

	f.Ln(10)
}

// Span writes a line of text plain text.
// Intended for use in FullP, Blockquote, & CodeBlock.
func Span(f *gofpdf.Fpdf, text string, styleStr string) {
	f.SetFont("helvetica", styleStr, 12)
	f.SetFillColor(255, 255, 255)
	f.SetTextColor(36, 41, 46)

	// _, lineHt := f.GetFontSize()
	// fmt.Println(lineHt)
	var lineHt float64 = 6
	f.Write(lineHt, text)
}

// InlineCode writes code with a grey background in courier font.
func InlineCode(f *gofpdf.Fpdf, text string) {
	f.SetFont("courier", "", 12)
	// _, lineHeight := f.GetFontSize()
	width := (f.GetStringWidth(text) * 1.1)
	f.SetFillColor(243, 243, 243)
	// cellMargin := f.GetCellMargin()
	// fmt.Println(cellMargin)

	f.CellFormat(
		width, // width; If 0, the cell extends up to the right margin.
		6,     // height;
		text,  // txtStr;
		"",    // borderStr; "" (no border), "1" (full border), "L", "T", "R" and "B" (left, top, right and bottom)
		0,     // ln; 0 (to the right), 1 (to the beginning of the next line, like Ln()), and 2 (below).
		"LM",  // alignStr; Default is left middle. "L", "C" or "R" (left, center, right) + "T", "M", "B" or "A" (top, middle, bottom, baseline)
		true,  // fill; true for color, false for transparent
		0,     // link;  Identifier returned by AddLink() or 0 for no internal link.
		"",    // linkStr; A target URL or empty for no external link. A non-zero value for link takes precedence over linkStr.
	)
}

// Strike writes a line of text with a line through it.
func Strike(f *gofpdf.Fpdf, text string, link string) {
	f.SetFont("helvetica", "", 12)
	f.SetFillColor(255, 255, 255)
	f.SetTextColor(36, 41, 46)
	width := f.GetStringWidth(text)
	// fontSize, lineHt := f.GetFontSize()
	var lineHt float64 = 6
	// lineHt = 6

	// Where the text starts, also where to start the strikethrough line.
	x := f.GetX()
	y := (f.GetY() + (lineHt / 2))

	// f.SetLineWidth(1 / fontSize)
	f.SetLineWidth(0.25)
	f.SetDrawColor(36, 41, 46)

	if len(link) != 0 {
		f.SetTextColor(3, 102, 214)
		f.WriteLinkString(lineHt, text, link)
	} else {
		f.Write(lineHt, text)
	}

	f.Line(x, y, x+width, y)
}
