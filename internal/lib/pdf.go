package lib

import "github.com/jung-kurt/gofpdf"

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
	pageWidth, _ := fpdf.GetPageSize()
	leftMargin, _, rightMargin, _ := fpdf.GetMargins()
	return (pageWidth - leftMargin - rightMargin)
}

// ContentBoxHeight returns the height of a page minus top and bottom margins.
func ContentBoxHeight(fpdf *gofpdf.Fpdf) float64 {
	_, pageHeight := fpdf.GetPageSize()
	_, topMargin, _, bottomMargin := fpdf.GetMargins()
	return (pageHeight - topMargin - bottomMargin)
}

// ContentBoxBottom returns the page height minus the bottom margin.
// Anything written below this needs to break to a new page.
func ContentBoxBottom(fpdf *gofpdf.Fpdf) float64 {
	_, pageHeight := fpdf.GetPageSize()
	_, _, _, bottomMargin := fpdf.GetMargins()
	return (pageHeight - bottomMargin)
}

// ContentBoxRight returns the page width minus the right margin.
// Anything written beyond this needs to wrap to a new line.
func ContentBoxRight(fpdf *gofpdf.Fpdf) float64 {
	pageWidth, _ := fpdf.GetPageSize()
	_, _, rightMargin, _ := fpdf.GetMargins()
	return (pageWidth - rightMargin)
}
