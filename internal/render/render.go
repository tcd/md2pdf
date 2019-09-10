// Package render renders text in the style of GitHub using gofpdf.
package render

import (
	"time"

	"github.com/jung-kurt/gofpdf"
)

// Setup sets default GitHub-Flavored styles to a gofpdf.Fpdf.
func Setup(f *gofpdf.Fpdf) {
	f.SetMargins(20, 20, 20)     // left, top, right margins
	f.SetAutoPageBreak(true, 20) // bottom margin
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

// DefaultFG rgb values (36, 41, 46)
func DefaultFG() (r, g, b uint8) {
	r = 36
	g = 41
	b = 46
	return
}

// DefaultBG rgb values (255, 255, 255)
func DefaultBG() (r, g, b uint8) {
	r = 255
	g = 255
	b = 255
	return
}

// LinkFG rgb values (3, 102, 214)
func LinkFG() (r, g, b uint8) {
	r = 3
	g = 102
	b = 214
	return
}

// CodeSpanBG rgb values: (243, 243, 243)
func CodeSpanBG() (r, g, b uint8) {
	r = 243
	g = 243
	b = 243
	return
}

// LightCodeBlockBG rgb values: (246, 248, 250)
func LightCodeBlockBG() (r, g, b uint8) {
	r = 246
	g = 248
	b = 250
	return
}

// TableCellBG rgb values: (246, 248, 250)
func TableCellBG() (r, g, b uint8) {
	r = 246
	g = 248
	b = 250
	return
}

// BlockquoteFG rgb values: (106, 115, 125)
func BlockquoteFG() (r, g, b uint8) {
	r = 106
	g = 115
	b = 125
	return
}

// BlockquoteBorder rgb values: (223, 226, 229)
func BlockquoteBorder() (r, g, b uint8) {
	r = 223
	g = 226
	b = 229
	return
}

// BlockquoteBG rgb values: (223, 226, 229)
func BlockquoteBG() (r, g, b uint8) {
	r = 223
	g = 226
	b = 229
	return
}

// HeaderUnderlineFG rgb values: (234, 236, 239)
func HeaderUnderlineFG() (r, g, b uint8) {
	r = 234
	g = 236
	b = 239
	return
}

// HrFG rgb values: (191, 191, 191)
func HrFG() (r, g, b uint8) {
	r = 191
	g = 191
	b = 191
	return
}

// CodeOrangeFG rgb values: (227, 98, 9)
func CodeOrangeFG() (r, g, b uint8) {
	r = 227
	g = 98
	b = 9
	return
}
