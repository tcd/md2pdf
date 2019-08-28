package model

import "github.com/jung-kurt/gofpdf"

// PDF attaches useful methods to an Fpdf.
type PDF struct {
	*gofpdf.Fpdf // https://github.com/jung-kurt/gofpdf/blob/master/def.go#L499
}

// PageWidth of a pdf.
func (pdf PDF) PageWidth() float64 {
	width, _ := pdf.GetPageSize()
	return width
}

// LeftMargin of a pdf.
func (pdf PDF) LeftMargin() float64 {
	leftMargin, _, _, _ := pdf.GetMargins()
	return leftMargin
}

// TopMargin of a pdf.
func (pdf PDF) TopMargin() float64 {
	_, topMargin, _, _ := pdf.GetMargins()
	return topMargin
}

// RightMargin of a pdf.
func (pdf PDF) RightMargin() float64 {
	_, _, rightMargin, _ := pdf.GetMargins()
	return rightMargin
}

// BottomMargin of a pdf.
func (pdf PDF) BottomMargin() float64 {
	_, _, _, bottomMargin := pdf.GetMargins()
	return bottomMargin
}

// ContentBox returns the width and height of a page minus margins.
func (pdf PDF) ContentBox() (width, height float64) {
	tWidth, tHeight := pdf.GetPageSize()
	leftM, topM, rightM, bottomM := pdf.GetMargins()

	height = (tHeight - topM - bottomM)
	width = (tWidth - leftM - rightM)

	return
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
