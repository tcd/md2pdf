package render

import (
	"fmt"
	"time"

	"github.com/jung-kurt/gofpdf"
)

func setMetaData(f *gofpdf.Fpdf, author, title string) {
	f.SetTitle(title, true)
	f.SetAuthor(author, true)
	f.SetCreator(author, true)
	f.SetCreationDate(time.Now())
}

// Write a linebreak; used after writing paragraphs.
func br(f *gofpdf.Fpdf) {
	f.Ln(1)
}

// Write a segment of plain text.
func span1(f *gofpdf.Fpdf, text string) {
	f.SetFont("helvetica", "", 12)
	f.SetFillColor(255, 255, 255)
	f.SetTextColor(36, 41, 46)
	f.SetCellMargin(0)

	_, lineHeight := f.GetFontSize()
	f.CellFormat(
		0,          // width; If 0, the cell extends up to the right margin.
		lineHeight, // height;
		text,       // txtStr;
		"",         // borderStr; "" (no border), "1" (full border), "L", "T", "R" and "B" (left, top, right and bottom)
		0,          // ln; 0 (to the right), 1 (to the beginning of the next line, like Ln()), and 2 (below).
		"",         // alignStr; Default is left middle. "L", "C" or "R" (left, center, right) + "T", "M", "B" or "A" (top, middle, bottom, baseline)
		false,      // fill; true for color, false for transparent
		0,          // link;  Identifier returned by AddLink() or 0 for no internal link.
		"",         // linkStr; A target URL or empty for no external link. A non-zero value for link takes precedence over linkStr.
	)
}

// Write a segment of plain text.
func span2(f *gofpdf.Fpdf, text string) {
	// _, lineHeight := f.GetFontSize()
	f.SetFont("helvetica", "", 12)
	f.SetFillColor(255, 255, 255)
	f.SetTextColor(36, 41, 46)
	f.Write(1, text)
	f.Write(1, " ")
}

func span(f *gofpdf.Fpdf, text string) {
	f.SetFont("helvetica", "", 12)
	f.SetFillColor(255, 255, 255)
	f.SetTextColor(36, 41, 46)

	_, lineHeight := f.GetFontSize()
	// width := (f.GetStringWidth(text) * 1.05)
	// f.CellFormat(
	// 	width,      // width; If 0, the cell extends up to the right margin.
	// 	lineHeight, // height;
	// 	text,       // txtStr;
	// 	"",         // borderStr; "" (no border), "1" (full border), "L", "T", "R" and "B" (left, top, right and bottom)
	// 	0,          // ln; 0 (to the right), 1 (to the beginning of the next line, like Ln()), and 2 (below).
	// 	"LT",       // alignStr; Default is left middle. "L", "C" or "R" (left, center, right) + "T", "M", "B" or "A" (top, middle, bottom, baseline)
	// 	false,      // fill; true for color, false for transparent
	// 	0,          // link;  Identifier returned by AddLink() or 0 for no internal link.
	// 	"",         // linkStr; A target URL or empty for no external link. A non-zero value for link takes precedence over linkStr.
	// )
	f.MultiCell(0, lineHeight, text, "", "LT", false)
}

// write italic text
func em(f *gofpdf.Fpdf, text string) {
	fontSize, lineHeight := f.GetFontSize()
	f.SetFont("", "I", fontSize)

	f.CellFormat(
		0,          // width; If 0, the cell extends up to the right margin.
		lineHeight, // height;
		text,       // txtStr;
		"",         // borderStr; "" (no border), "1" (full border), "L", "T", "R" and "B" (left, top, right and bottom)
		0,          // ln; 0 (to the right), 1 (to the beginning of the next line, like Ln()), and 2 (below).
		"",         // alignStr; Default is left middle. "L", "C" or "R" (left, center, right) + "T", "M", "B" or "A" (top, middle, bottom, baseline)
		false,      // fill; true for color, false for transparent
		0,          // link;  Identifier returned by AddLink() or 0 for no internal link.
		"",         // linkStr; A target URL or empty for no external link. A non-zero value for link takes precedence over linkStr.
	)
}

// Write bold text
func strong(f *gofpdf.Fpdf, text string) {
	fontSize, _ := f.GetFontSize()
	f.SetFont("", "B", fontSize)

	_, lineHeight := f.GetFontSize()
	width := (f.GetStringWidth(text) * 1.05)
	f.CellFormat(
		width,      // width; If 0, the cell extends up to the right margin.
		lineHeight, // height;
		text,       // txtStr;
		"",         // borderStr; "" (no border), "1" (full border), "L", "T", "R" and "B" (left, top, right and bottom)
		0,          // ln; 0 (to the right), 1 (to the beginning of the next line, like Ln()), and 2 (below).
		"LT",       // alignStr; Default is left middle. "L", "C" or "R" (left, center, right) + "T", "M", "B" or "A" (top, middle, bottom, baseline)
		false,      // fill; true for color, false for transparent
		0,          // link;  Identifier returned by AddLink() or 0 for no internal link.
		"",         // linkStr; A target URL or empty for no external link. A non-zero value for link takes precedence over linkStr.
	)
}

// Write BoldItalic text
func boldItalic(f *gofpdf.Fpdf, text string) {
	fontSize, lineHeight := f.GetFontSize()
	f.SetFont("", "BI", fontSize)
	f.Write(lineHeight, text)
}

// Write text with a line through it
func strike(f *gofpdf.Fpdf, text string) {
	width := f.GetStringWidth(text)
	fontSize, lineHeight := f.GetFontSize()
	x := f.GetX()
	y := (f.GetY() + (lineHeight / 2))

	f.SetLineWidth(1 / fontSize)
	f.SetDrawColor(36, 41, 46)

	f.Write(lineHeight, text)
	f.Line(x, y, x+width, y)
}

func code(f *gofpdf.Fpdf, text string) {
	// _, lineHeight := f.GetFontSize()
	f.SetFont("courier", "", 12)
	f.Write(1, text)
	f.Write(1, " ")
}

func code2(f *gofpdf.Fpdf, text string) {
	f.SetFont("courier", "", 12)
	_, lineHeight := f.GetFontSize()
	width := (f.GetStringWidth(text) * 1.05)
	f.SetFillColor(230, 234, 237)

	f.CellFormat(
		width,      // width; If 0, the cell extends up to the right margin.
		lineHeight, // height;
		text,       // txtStr;
		"",         // borderStr; "" (no border), "1" (full border), "L", "T", "R" and "B" (left, top, right and bottom)
		0,          // ln; 0 (to the right), 1 (to the beginning of the next line, like Ln()), and 2 (below).
		"LT",       // alignStr; Default is left middle. "L", "C" or "R" (left, center, right) + "T", "M", "B" or "A" (top, middle, bottom, baseline)
		true,       // fill; true for color, false for transparent
		0,          // link;  Identifier returned by AddLink() or 0 for no internal link.
		"",         // linkStr; A target URL or empty for no external link. A non-zero value for link takes precedence over linkStr.
	)
}

// WriteLinkString is the gofpdf function for writing links.
func Link1(f *gofpdf.Fpdf, text string, href string) {
	fontSize, lineHeight := f.GetFontSize()
	f.SetTextColor(3, 102, 214)
	f.SetFont("helvetica", "", fontSize)
	htmlLink := fmt.Sprintf(`<a href="%s">%s</a>`, href, text)
	html := f.HTMLBasicNew()
	html.Write(lineHeight, htmlLink)
}

// WriteLinkString is the gofpdf function for writing links.
func link2(f *gofpdf.Fpdf, text string, href string) {
	fontSize, _ := f.GetFontSize()
	f.SetTextColor(3, 102, 214)
	f.SetFont("helvetica", "", fontSize)

	// f.CellFormat(
	// 	0,          // width; If 0, the cell extends up to the right margin.
	// 	lineHeight, // height;
	// 	text,       // txtStr;
	// 	"",         // borderStr; "" (no border), "1" (full border), "L", "T", "R" and "B" (left, top, right and bottom)
	// 	1,          // ln; 0 (to the right), 1 (to the beginning of the next line, like Ln()), and 2 (below).
	// 	"",         // alignStr; Default is left middle. "L", "C" or "R" (left, center, right) + "T", "M", "B" or "A" (top, middle, bottom, baseline)
	// 	false,      // fill; true for color, false for transparent
	// 	0,          // link;  Identifier returned by AddLink() or 0 for no internal link.
	// 	"",         // linkStr; A target URL or empty for no external link. A non-zero value for link takes precedence over linkStr.
	// )
}
