package render

import "github.com/jung-kurt/gofpdf"

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
	f.Ln(11)
}
