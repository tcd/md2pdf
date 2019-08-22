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
