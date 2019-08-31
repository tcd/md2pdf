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
