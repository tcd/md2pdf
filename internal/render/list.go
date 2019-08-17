package render

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

// http://www.fpdf.org/en/script/script56.php
// https://github.com/ajstarks/deck/blob/master/cmd/pdfdeck/pdfdeck.go#L270

// OrderedList renders a list to a gofpdf.Fpdf.
func OrderedList(f *gofpdf.Fpdf, items []string) {
	oldLeftMargin, _, _, _ := f.GetMargins() // need to call this before we call SetLeftMargin.
	f.SetLeftMargin(22.5)
	f.SetFont("helvetica", "", 12)
	_, lineHt := f.GetFontSize() // need to call this after we call SetFont.
	lineHt = lineHt * 1.8

	for i, item := range items {
		f.SetFont("helvetica", "", 12)
		f.CellFormat(
			5,                       // width
			lineHt,                  // height
			fmt.Sprintf("%d.", i+1), // order :)
			"",                      // no border
			0,                       // ln; 0 (to the right), 1 (to the beginning of the next line, like Ln()), and 2 (below).
			"RM",                    // alignStr; Default is left middle. "L", "C" or "R" (left, center, right) + "T", "M", "B" or "A" (top, middle, bottom, baseline)
			false,                   // no fill
			0,                       // no link
			"",                      // no linkStr
		)
		f.SetFont("helvetica", "", 12)
		// f.CellFormat(5, lineHt, item, "", 1, "LM", false, 0, "")
		f.MultiCell(0, lineHt, item, "", "LM", false)
	}
	f.Ln(4)
	f.SetLeftMargin(oldLeftMargin)
}

// UnorderedList renders a list to a gofpdf.Fpdf.
func UnorderedList(f *gofpdf.Fpdf, items []string) {
	oldLeftMargin, _, _, _ := f.GetMargins()
	f.SetLeftMargin(22.5)
	f.SetFont("helvetica", "", 12)
	_, lineHt := f.GetFontSize()
	lineHt = lineHt * 1.7

	for _, item := range items {
		f.SetFont("zapfdingbats", "", 5)
		f.CellFormat(
			5,      // width
			lineHt, // height
			"\x6c", // ●
			"",     // no border
			0,      // ln; 0 (to the right), 1 (to the beginning of the next line, like Ln()), and 2 (below).
			"RM",   // alignStr; Default is left middle. "L", "C" or "R" (left, center, right) + "T", "M", "B" or "A" (top, middle, bottom, baseline)
			false,  // no fill
			0,      // no link
			"",     // no linkStr
		)
		f.SetFont("helvetica", "", 12)
		// f.CellFormat(5, lineHt, item, "", 1, "LM", false, 0, "")
		f.MultiCell(0, lineHt, item, "", "LM", false)
	}
	f.Ln(4)
	f.SetLeftMargin(oldLeftMargin)
}
