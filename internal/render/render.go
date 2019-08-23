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

// ColorCodes has methods for returning color values.
type ColorCodes uint8

// DefaultFG rgb values.
func (cc ColorCodes) DefaultFG() (r, g, b int) {
	r = 36
	g = 41
	b = 46
	return
}

// DefaultBG rgb values.
func (cc ColorCodes) DefaultBG() (r, g, b int) {
	r = 255
	g = 255
	b = 255
	return
}

// LinkFG rgb values.
func (cc ColorCodes) LinkFG() (r, g, b int) {
	r = 3
	g = 102
	b = 214
	return
}
