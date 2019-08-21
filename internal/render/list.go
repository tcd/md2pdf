package render

import (
	"strings"

	"github.com/jung-kurt/gofpdf"
)

// List is an array of ListItems.
type List struct {
	Items   []ListItem `json:"items"`
	Ordered bool       `json:"ordered"`
}

// AddItems appends ListItems to a List's Items.
func (ls *List) AddItems(listItem ...ListItem) {
	ls.Items = append(ls.Items, listItem...)
}

// ListItem is a single list bullet, and may contain a nested list.
type ListItem struct {
	Contents `json:"contents"`
	Children List `json:"children"`
}

// HasChildren returns true if a ListItem contains a nested list.
func (li ListItem) HasChildren() bool {
	if len(li.Children.Items) > 0 {
		return true
	}
	return false
}

const (
	levelOneMargin   = 22.5
	levelTwoMargin   = 27.5
	levelThreeMargin = 32.5
)

// AnyList writes an unordered list with up to three levels of indentation to a gofpdf.Fpdf.
func AnyList(f *gofpdf.Fpdf, list List) {
	f.SetLeftMargin(levelOneMargin)
	f.SetRightMargin(21)
	f.SetFont("helvetica", "", 12)
	_, lineHt := f.GetFontSize() // need to call this after we call SetFont.
	lineHt = lineHt * 1.5

	for _, item := range list.Items {
		levelOneBullett(f, lineHt)
		drawListItemContent(f, item.Contents, lineHt)
		f.SetLeftMargin(levelOneMargin)
		f.Ln(-1) // A negative value indicates the height of the last printed cell.
		f.Ln(2)

		if item.HasChildren() {
			f.SetLeftMargin(levelTwoMargin)
			lastIndex := len(item.Children.Items) - 1
			for i, item := range item.Children.Items {
				levelTwoBullet(f, lineHt)
				drawListItemContent(f, item.Contents, lineHt)
				f.SetLeftMargin(levelTwoMargin)
				if i != lastIndex {
					f.Ln(-1)
					f.Ln(2)
				}
				if item.HasChildren() {
					f.Ln(-1)
					f.Ln(2)
					f.SetLeftMargin(levelThreeMargin)
					lastIndex := len(item.Children.Items) - 1
					for i, item := range item.Children.Items {
						levelThreeBullet(f, lineHt)
						drawListItemContent(f, item.Contents, lineHt)
						f.SetLeftMargin(levelThreeMargin)
						if i != lastIndex {
							f.Ln(-1)
							f.Ln(2)
						}
					}
					f.SetLeftMargin(levelTwoMargin)
					f.Ln(-1)
					f.Ln(2)
				}
			}
			f.SetLeftMargin(levelOneMargin)
			f.Ln(-1)
			f.Ln(2)
		}
	}

	f.SetLeftMargin(20)
	f.SetRightMargin(20)
	f.Ln(4)
}

func levelOneBullett(f *gofpdf.Fpdf, lineHt float64) {
	f.SetFont("zapfdingbats", "", 5)
	f.CellFormat(
		5,      // width
		lineHt, // height
		"\x6c", // ●
		"",     // no border
		0,      // ln; 0 (to the right), 1 (to the beginning of the next line, like Ln()), and 2 (below).
		"RM",   // alignStr; Default is left middle. "L", "C" or "R" (left, center, right) + "T", "M", "B" or "A" (top, middle, bottom, baseline)
		false,  // no fill
		0,      // no link
		"",     // no linkStr
	)
	x := f.GetX()
	f.SetLeftMargin(x)
}

func levelTwoBullet(f *gofpdf.Fpdf, lineHt float64) {
	f.SetFont("zapfdingbats", "", 5)
	f.CellFormat(
		5,      // width
		lineHt, // height
		"\x6d", // ❍
		"",     // no border
		0,      // ln; 0 (to the right), 1 (to the beginning of the next line, like Ln()), and 2 (below).
		"RM",   // alignStr; Default is left middle. "L", "C" or "R" (left, center, right) + "T", "M", "B" or "A" (top, middle, bottom, baseline)
		false,  // no fill
		0,      // no link
		"",     // no linkStr
	)
	x := f.GetX()
	f.SetLeftMargin(x)
}

func levelThreeBullet(f *gofpdf.Fpdf, lineHt float64) {
	f.SetFont("zapfdingbats", "", 5)
	f.CellFormat(
		5,      // width
		lineHt, // height
		"\x6e", // ■
		"",     // no border
		0,      // ln; 0 (to the right), 1 (to the beginning of the next line, like Ln()), and 2 (below).
		"RM",   // alignStr; Default is left middle. "L", "C" or "R" (left, center, right) + "T", "M", "B" or "A" (top, middle, bottom, baseline)
		false,  // no fill
		0,      // no link
		"",     // no linkStr
	)
	x := f.GetX()
	f.SetLeftMargin(x)
}

func drawListItemContent(f *gofpdf.Fpdf, c Contents, lineHt float64) {
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
			liStrike(f, txt.Content, txt.HREF)
		} else if txt.Code {
			liInlineCode(f, txt.Content, txt.HREF)
		} else if txt.HREF != "" {
			f.SetFont("helvetica", styleStr, 12)
			f.SetTextColor(3, 102, 214)
			f.WriteLinkString(lineHt, txt.Content, txt.HREF)
		} else {
			liSpan(f, txt.Content, styleStr, lineHt)
		}
	}
}

func liSpan(f *gofpdf.Fpdf, text string, styleStr string, lineHt float64) {
	f.SetFont("helvetica", styleStr, 12)
	f.SetFillColor(255, 255, 255)
	f.SetTextColor(36, 41, 46)
	f.Write(lineHt, text)
}

func liInlineCode(f *gofpdf.Fpdf, text, href string) {
	oldCellMargin := f.GetCellMargin()
	f.SetFillColor(243, 243, 243)
	f.SetFont("courier", "", 12)
	width := (f.GetStringWidth(text) * 1.05)

	if len(href) != 0 {
		f.SetTextColor(3, 102, 214)
		f.CellFormat(
			width, // width; If 0, the cell extends up to the right margin.
			6,     // height;
			text,  // txtStr;
			"",    // borderStr; "" (no border), "1" (full border), "L", "T", "R" and "B" (left, top, right and bottom)
			0,     // ln; 0 (to the right), 1 (to the beginning of the next line, like Ln()), and 2 (below).
			"LM",  // alignStr; Default is left middle. "L", "C" or "R" (left, center, right) + "T", "M", "B" or "A" (top, middle, bottom, baseline)
			true,  // fill; true for color, false for transparent
			0,     // link;  Identifier returned by AddLink() or 0 for no internal link.
			href,  // linkStr; A target URL or empty for no external link. A non-zero value for link takes precedence over linkStr.
		)
	} else {
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

	f.SetCellMargin(oldCellMargin)
}

func liStrike(f *gofpdf.Fpdf, text string, href string) {
	f.SetFont("helvetica", "", 12)
	f.SetFillColor(255, 255, 255)
	f.SetTextColor(36, 41, 46)
	width := f.GetStringWidth(text)
	var lineHt float64 = 6

	// Where the text starts, also where to start the strikethrough line.
	x := f.GetX()
	y := (f.GetY() + (lineHt / 2))

	f.SetLineWidth(0.25)
	f.SetDrawColor(36, 41, 46)

	if len(href) != 0 {
		f.SetTextColor(3, 102, 214)
		f.WriteLinkString(lineHt, text, href)
	} else {
		f.Write(lineHt, text)
	}

	f.Line(x, y, x+width, y)
}
