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
	f.SetMargins(20, 15, 20)
	f.AddPage()
	f.SetFillColor(255, 255, 255)
	f.SetTextColor(36, 41, 46)
}

// SetMetaData for a pdf.
func SetMetaData(f *gofpdf.Fpdf, author, title string) {
	f.SetTitle(title, true)
	f.SetAuthor(author, true)
	f.SetCreator(author, true)
	f.SetCreationDate(time.Now())
}

// ContentBox returns the width and height of a page minus margins.
func ContentBox(fpdf *gofpdf.Fpdf) (width, height float64) {
	tWidth, tHeight := fpdf.GetPageSize()
	leftM, topM, rightM, bottomM := fpdf.GetMargins()

	height = (tHeight - topM - bottomM)
	width = (tWidth - leftM - rightM)
	return
}

// ContentBoxWidth returns the width of a width minus left and right margins.
func ContentBoxWidth(fpdf *gofpdf.Fpdf) float64 {
	tWidth, _ := fpdf.GetPageSize()
	leftM, _, rightM, _ := fpdf.GetMargins()
	return (tWidth - leftM - rightM)
}

// ContentBoxHeight returns the height of a page minus top and bottom margins.
func ContentBoxHeight(fpdf *gofpdf.Fpdf) float64 {
	_, tHeight := fpdf.GetPageSize()
	_, topM, _, bottomM := fpdf.GetMargins()
	return (tHeight - topM - bottomM)
}

// DefaultFG rgb values.
func DefaultFG() (r, g, b int) {
	r = 36
	g = 41
	b = 46
	return
}

// DefaultBG rgb values.
func DefaultBG() (r, g, b int) {
	r = 255
	g = 255
	b = 255
	return
}

// LinkFG rgb values. (<a>)
func LinkFG() (r, g, b int) {
	r = 3
	g = 102
	b = 214
	return
}

// CodeSpanBG rgb values. (<code>)
func CodeSpanBG() (r, g, b int) {
	r = 243
	g = 243
	b = 243
	return
}

// CodeBlockBG rgb values. (<code>)
func CodeBlockBG() (r, g, b int) {
	r = 246
	g = 248
	b = 250
	return
}

// TableCellBG rgb values.
func TableCellBG() (r, g, b int) {
	r = 246
	g = 248
	b = 250
	return
}

// BlockquoteFG rgb values.
func BlockquoteFG() (r, g, b int) {
	r = 106
	g = 115
	b = 125
	return
}

// BlockquoteBorder rgb values.
func BlockquoteBorder() (r, g, b int) {
	r = 223
	g = 226
	b = 229
	return
}

// BlockquoteBG rgb values.
func BlockquoteBG() (r, g, b int) {
	r = 223
	g = 226
	b = 229
	return
}

// HeaderUnderlineFG rgb values.
func HeaderUnderlineFG() (r, g, b int) {
	r = 234
	g = 236
	b = 239
	return
}

// CodePurpleFG rgb values.
func CodePurpleFG() (r, g, b int) {
	r = 111
	g = 66
	b = 193
	return
}

// CodeOrangeFG rgb values.
func CodeOrangeFG() (r, g, b int) {
	r = 227
	g = 98
	b = 9
	return
}
