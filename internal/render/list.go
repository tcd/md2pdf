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
			liInlineCode(pdf, txt)
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

func liInlineCode(pdf *gofpdf.Fpdf, txt Text) {
	oldCellMargin := pdf.GetCellMargin()
	pdf.SetFillColor(243, 243, 243)
	pdf.SetFont("courier", "", 12)
	width := (pdf.GetStringWidth(txt.Content) * 1.05)
	pdf.SetCellMargin(2)

	x := pdf.GetX()
	pageWidth, _ := pdf.GetPageSize()
	_, _, rightMargin, _ := pdf.GetMargins()
	strWdth := pdf.GetStringWidth(txt.Content)
	if x+strWdth > pageWidth-rightMargin {
		pdf.Ln(-1)
	}

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
