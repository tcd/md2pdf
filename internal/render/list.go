package render

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/model"
)

var margins = map[int]float64{
	0: 22.5,
	1: 22.5,
	2: 27.5,
	3: 32.5,
}

// List writes a list with any level of indentation to a gofpdf.Fpdf.
func List(pdf *gofpdf.Fpdf, list model.ListContent) {
	level := 1
	pdf.SetLeftMargin(margins[level])
	pdf.SetRightMargin(21)
	pdf.SetFont("helvetica", "", 12)
	_, lineHt := pdf.GetFontSize() // need to call this after we call SetFont.
	lineHt *= 1.5

	for i, item := range list.Items {
		if list.Ordered {
			drawNumbering(pdf, lineHt, level, i+1)
		} else {
			drawBullet(pdf, lineHt, level)
		}
		drawContent(pdf, item.Contents)
		pdf.SetLeftMargin(margins[level])
		pdf.Ln(-1)
		pdf.Ln(2)
		if item.HasChildren() {
			drawList(pdf, item.Children, level+1)
		}
	}

	pdf.SetLeftMargin(20)
	pdf.SetRightMargin(20)
	pdf.Ln(4)
}

// Draw all items in a ListContent, anc call drawList for any sublists.
func drawList(pdf *gofpdf.Fpdf, list model.ListContent, level int) {
	pdf.SetLeftMargin(margins[level])
	pdf.SetFont("helvetica", "", 12)
	_, lineHt := pdf.GetFontSize() // need to call this after we call SetFont.
	lineHt *= 1.5

	lastIndex := len(list.Items) - 1
	for i, item := range list.Items {
		if list.Ordered {
			drawNumbering(pdf, lineHt, level, i+1)
		} else {
			drawBullet(pdf, lineHt, level)
		}
		drawContent(pdf, item.Contents)
		pdf.SetLeftMargin(margins[level])
		if i != lastIndex {
			pdf.Ln(-1)
			pdf.Ln(2)
		}
		if item.HasChildren() {
			drawList(pdf, item.Children, level+1)
		}
	}
	pdf.SetLeftMargin(margins[level-1])
	pdf.Ln(-1)
	pdf.Ln(2)
}
