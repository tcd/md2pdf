package render

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/model"
)

// 20 is the default margin;
// Start at 22.5 and increment by 7.5.
var margins = map[int]float64{
	0:  20.0,
	1:  22.5,
	2:  30.0,
	3:  37.5,
	4:  45.0,
	5:  52.5,
	6:  60.0,
	7:  67.5,
	8:  72.5,
	9:  80.0,
	10: 87.5,
}

// List writes a list with any level of indentation to a gofpdf.Fpdf.
func List(pdf *gofpdf.Fpdf, list model.ListContent) {
	level := 1
	pdf.SetLeftMargin(margins[level])
	pdf.SetRightMargin(21)

	drawList(pdf, list, level)

	pdf.SetLeftMargin(20)
	pdf.SetRightMargin(20)
	pdf.SetFont("helvetica", "", 12)
}

// Draw all items in a ListContent, anc call drawList for any sublists.
func drawList(pdf *gofpdf.Fpdf, list model.ListContent, level int) {
	pdf.SetLeftMargin(margins[level])
	pdf.SetFont("helvetica", "", 12)
	_, oldLineHt := pdf.GetFontSize()
	lineHt := oldLineHt * 1.5

	for i, item := range list.Items {
		if list.Ordered {
			drawNumbering(pdf, lineHt, level, i+1)
		} else {
			drawBullet(pdf, lineHt, level)
		}
		drawContent(pdf, item.Contents, 12)
		pdf.SetLeftMargin(margins[level])
		pdf.Ln(-1)
		if item.HasChildren() {
			drawList(pdf, item.Children, level+1)
		}
	}
	pdf.SetLeftMargin(margins[level-1])
	if level == 1 {
		pdf.Ln(-1)
		pdf.Ln(1)
	} else {
		// pdf.Ln(1 / 1.5)
		pdf.Ln(0.33)
	}
}
