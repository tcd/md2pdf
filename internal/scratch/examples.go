package scratch

import (
	"fmt"
	"os"

	"github.com/jung-kurt/gofpdf"
)

// BulletList ...
func BulletList() {
	const fontHt = 18
	pdf := gofpdf.New("P", "mm", "A4", "")
	lineHt := pdf.PointToUnitConvert(fontHt)
	pdf.AddPage()
	for j := 1; j < 12; j++ {
		pdf.SetFont("zapfdingbats", "", fontHt/4)
		pdf.CellFormat(5, lineHt, "\x6c", "", 0, "R", false, 0, "")
		pdf.SetFont("times", "", 18)
		pdf.CellFormat(5, lineHt, fmt.Sprintf(" Item %d", j), "", 1, "L", false, 0, "")
	}
	err := pdf.OutputFileAndClose("bullet.pdf")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error writing to bullet.pdf: %s\n", err)
	}
}

// func example1() {
// 	const (
// 		fontPtSize = 18.0
// 		wd         = 100.0
// 	)
// 	pdf := gofpdf.New("P", "mm", "A4", "") // A4 210.0 x 297.0
// 	pdf.SetFont("Times", "", fontPtSize)
// 	_, lineHt := pdf.GetFontSize()
// 	pdf.AddPage()
// 	pdf.SetMargins(10, 10, 10)
// 	lines := pdf.SplitLines([]byte(lorem()), wd)
// 	ht := float64(len(lines)) * lineHt
// 	y := (297.0 - ht) / 2.0
// 	pdf.SetDrawColor(128, 128, 128)
// 	pdf.SetFillColor(255, 255, 210)
// 	x := (210.0 - (wd + 40.0)) / 2.0
// 	pdf.Rect(x, y-20.0, wd+40.0, ht+40.0, "FD")
// 	pdf.SetY(y)
// 	for _, line := range lines {
// 		pdf.CellFormat(190.0, lineHt, string(line), "", 1, "C", false, 0, "")
// 	}
// 	fileStr := example.Filename("Fpdf_Splitlines")
// 	err := pdf.OutputFileAndClose(fileStr)
// 	example.Summary(err, fileStr)
// }
