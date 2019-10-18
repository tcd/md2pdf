package github

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/lib"
)

func drawBullet(pdf *gofpdf.Fpdf, lineHt float64, level int) {
	// Make sure there's enough room to draw the bullet and the list item contents.
	if pdf.GetY()+lineHt > lib.ContentBoxBottom(pdf) {
		pdf.AddPage()
	}
	var bulletChar string
	switch level {
	case 1:
		bulletChar = "\x6c" // ●
	case 2:
		bulletChar = "\x6d" // ❍
	case 3:
		bulletChar = "\x6e" // ■
	default:
		bulletChar = "\x6e" // ■
	}
	pdf.SetTextColor(DefaultFG())
	pdf.SetFont("zapfdingbats", "", 6)
	pdf.CellFormat(6, lineHt, bulletChar, "", 0, "RM", false, 0, "")
	pdf.SetLeftMargin(pdf.GetX())
}

func drawNumbering(pdf *gofpdf.Fpdf, lineHt float64, level, number int) {
	var numberList map[int]string
	switch level {
	case 1:
		numberList = levelOneNumbers
	case 2:
		numberList = levelTwoNumbers
	case 3:
		numberList = levelThreeNumbers
	default:
		numberList = levelOneNumbers
	}
	pdf.SetTextColor(DefaultFG())
	pdf.SetFont("helvetica", "", 12)
	pdf.CellFormat(5, lineHt, numberList[number], "", 0, "RM", false, 0, "")
	pdf.SetLeftMargin(pdf.GetX())
}

func loweralpha() string {
	p := make([]byte, 26)
	for i := range p {
		p[i] = 'a' + byte(i)
	}
	return string(p)
}

var levelOneNumbers = map[int]string{
	1:  "1.",
	2:  "2.",
	3:  "3.",
	4:  "4.",
	5:  "5.",
	6:  "6.",
	7:  "7.",
	8:  "8.",
	9:  "9.",
	10: "10.",
	11: "11.",
	12: "12.",
	13: "13.",
	14: "14.",
	15: "15.",
	16: "16.",
	17: "17.",
	18: "18.",
	19: "19.",
	20: "20.",
	21: "21.",
	22: "22.",
	23: "23.",
	24: "24.",
	25: "25.",
	26: "26.",
}

var levelTwoNumbers = map[int]string{
	1:  "i.",
	2:  "ii.",
	3:  "iii.",
	4:  "iv.",
	5:  "v.",
	6:  "vi.",
	7:  "vii.",
	8:  "viii.",
	9:  "ix.",
	10: "x.",
	11: "xi.",
	12: "xii.",
	13: "xiii.",
	14: "xiv.",
	15: "xv.",
	16: "xvi.",
	17: "xvii.",
	18: "xviii.",
	19: "xix.",
	20: "xx.",
	21: "xxi.",
	22: "xxii.",
	23: "xxiii.",
	24: "xxiv.",
	25: "xxv.",
	26: "xxvi.",
}

var levelThreeNumbers = map[int]string{
	1:  "a.",
	2:  "b.",
	3:  "c.",
	4:  "d.",
	5:  "e.",
	6:  "f.",
	7:  "g.",
	8:  "h.",
	9:  "i.",
	10: "j.",
	11: "k.",
	12: "l.",
	13: "m.",
	14: "n.",
	15: "o.",
	16: "p.",
	17: "q.",
	18: "r.",
	19: "s.",
	20: "t.",
	21: "u.",
	22: "v.",
	23: "w.",
	24: "x.",
	25: "y.",
	26: "z.",
}
