// Package render renders text in the style of GitHub using gofpdf.
package render

import "github.com/jung-kurt/gofpdf"

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
