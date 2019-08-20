// Package render renders text in the style of GitHub using gofpdf.
package render

import (
	"time"

	"github.com/jung-kurt/gofpdf"
)

const (
	pageWidth = 176 // (216 - 40 for margins)
)

// Setup sets default GitHub-Flavored styles to a gofpdf.Fpdf.
func Setup(f *gofpdf.Fpdf) {
	f.AddPage()
	f.SetFont("helvetica", "", 16)
	f.SetMargins(20, 13, 20)
	f.SetFillColor(255, 255, 255)
	f.SetTextColor(36, 41, 46)
}

func setMetaData(f *gofpdf.Fpdf, author, title string) {
	f.SetTitle(title, true)
	f.SetAuthor(author, true)
	f.SetCreator(author, true)
	f.SetCreationDate(time.Now())
}

// Text models text strings that make up a Paragraph, Blockquote, List Item, or Table Cell.
type Text struct {
	Content string
	Bold    bool
	Italic  bool
	Code    bool
	Strike  bool
	HREF    string
}

// Contents models the contents of a Paragraph, Blockquote, List Item, or Table Cell.
type Contents struct {
	Content []Text
}

// AddContent to a to Contents.
func (p *Contents) AddContent(text Text) {
	p.Content = append(p.Content, text)
}

// Copy returns a pointer to a new Text struct with the same
// values as the Text it was called from.
// Except the Content field, that's empty.
func (txt Text) Copy() *Text {
	var newText Text
	if txt.Bold {
		newText.Bold = true
	}
	if txt.Italic {
		newText.Italic = true
	}
	if txt.Code {
		newText.Code = true
	}
	if txt.Strike {
		newText.Strike = true
	}
	newText.HREF = txt.HREF
	return &newText
}
