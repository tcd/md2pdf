package lib

import gofpdf "github.com/tcd/gofpdf-1"

var width float64 = 612
var height float64 = 792

// ContentBox returns the width and height of a page minus margins.
func ContentBox(fpdf *gofpdf.Fpdf) (width, height float64) {
	// pb := fpdf.GetPageBoundary(gofpdf.PageBoundaryCrop)
	// pageWidth, pageHeight := pb.Size.W, pb.Size.H
	pageWidth, pageHeight := width, height
	leftMargin, topMargin, rightMargin, bottomMargin := fpdf.Margins()

	height = (pageHeight - topMargin - bottomMargin)
	width = (pageWidth - leftMargin - rightMargin)
	return
}

// ContentBoxWidth returns the width of a width minus left and right margins.
func ContentBoxWidth(fpdf *gofpdf.Fpdf) float64 {
	pageWidth := width
	leftMargin, _, rightMargin, _ := fpdf.Margins()
	return (pageWidth - leftMargin - rightMargin)
}

// ContentBoxHeight returns the height of a page minus top and bottom margins.
func ContentBoxHeight(fpdf *gofpdf.Fpdf) float64 {
	pageHeight := height
	_, topMargin, _, bottomMargin := fpdf.Margins()
	return (pageHeight - topMargin - bottomMargin)
}

// ContentBoxBottom returns the page height minus the bottom margin.
// Anything written below this needs to break to a new page.
func ContentBoxBottom(fpdf *gofpdf.Fpdf) float64 {
	pageHeight := height
	_, _, _, bottomMargin := fpdf.Margins()
	return (pageHeight - bottomMargin)
}

// ContentBoxRight returns the page width minus the right margin.
// Anything written beyond this needs to wrap to a new line.
func ContentBoxRight(fpdf *gofpdf.Fpdf) float64 {
	pageWidth := width
	_, _, rightMargin, _ := fpdf.Margins()
	return (pageWidth - rightMargin)
}
