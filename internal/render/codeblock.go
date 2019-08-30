package render

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/alecthomas/chroma/quick"
	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/model"
)

// CodeBlock writes a `<pre>` style code block to a gofpdf.Fpdf.
func CodeBlock(pdf *gofpdf.Fpdf, contents model.Contents) {
	if len(contents.Content) != 1 {
		log.Println("CodeBlock: invalid argument.")
	}

	content := contents.Content[0].Text

	oldCellMargin := pdf.GetCellMargin()

	pdf.SetFont("courier", "", 11)
	pdf.SetFillColor(LightCodeBlockBG())
	pdf.SetTextColor(DefaultFG())
	pdf.SetCellMargin(4)

	pdf.CellFormat(0, 4, "", "", 1, "TL", true, 0, "")
	pdf.MultiCell(0, 5, content, "", "", true)
	pdf.CellFormat(0, 4, "", "", 0, "TL", true, 0, "")
	pdf.Ln(10)

	pdf.SetCellMargin(oldCellMargin)
}

// HighlightedCodeblock draws a codeblock with syntax highlighting provided by Chroma.
func HighlightedCodeblock(pdf *gofpdf.Fpdf, contents model.Contents, class string) {
	code := contents.Content[0].Text
	tokens, err := generateTokens(code, class)
	if err != nil {
		log.Println(err)
	}

	oldCellMargin := pdf.GetCellMargin()
	pdf.SetCellMargin(5)
	pdf.SetFont("courier", "", 11)
	pdf.SetFillColor(30, 30, 30)
	_, lineHt := pdf.GetFontSize()
	lineHt *= 1.25

	// Draw out the text on a temporary PDF to see how much space it takes up.
	testPDF := dummyPDF()
	testPDF.SetFont("courier", "", 11)
	y1 := testPDF.GetY()
	testPDF.Ln(lineHt)
	for _, token := range tokens {
		testPDF.Write(lineHt, token.Value)
	}
	testPDF.Ln(lineHt)
	y2 := testPDF.GetY()
	blockHt := y2 - y1
	testPDF.Close()

	// Get top left corner, avoid page breaking inside a block.
	x1, y1 := pdf.GetXY()
	if y1+blockHt > ContentBoxHeight(pdf) {
		pdf.AddPage()
		x1, y1 = pdf.GetXY()
	}
	// Draw the background.
	pdf.Rect(x1, y1, ContentBoxWidth(pdf), blockHt, "F")

	// Write out the content for real this time.
	pdf.SetXY(x1, y1)
	pdf.Ln(lineHt)
	for _, token := range tokens {
		r, g, b, _ := HexToRGB(styleMap[token.Type])
		pdf.SetTextColor(r, g, b)
		pdf.Write(lineHt, cleanString(token.Value))
	}
	pdf.Ln(lineHt)

	pdf.SetCellMargin(oldCellMargin)
	pdf.Ln(6)
}

func generateTokens(code, class string) ([]chromaToken, error) {
	var tokens []chromaToken
	var buf bytes.Buffer
	lang := strings.Replace(class, "language-", "", -1)
	err := quick.Highlight(&buf, code, lang, "json", "monokai")
	err = json.Unmarshal(buf.Bytes(), &tokens)
	if err != nil {
		return tokens, err
	}
	return tokens, nil
}

type chromaToken struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// DarkPlus Style
var styleMap = map[string]string{
	// Other
	"Background":     "#1e1e1e",
	"Text":           "#d4d4d4",
	"TextWhitespace": "#d4d4d4",
	"Other":          "#d4d4d4",
	"Punctuation":    "#d4d4d4",
	"Error":          "#d16969",
	// Comments
	"Comment":          "#608b4e",
	"CommentHashbang":  "#608b4e",
	"CommentMultiline": "#608b4e",
	"CommentPreproc":   "#608b4e",
	"CommentSingle":    "#608b4e",
	"CommentSpecial":   "#608b4e",
	// Generic
	"Generic":           "#d4d4d4",
	"GenericDeleted":    "#d16969",
	"GenericEmph":       "#d4d4d4",
	"GenericError":      "#d16969",
	"GenericHeading":    "#d4d4d4",
	"GenericInserted":   "#608b4e",
	"GenericOutput":     "#d4d4d4",
	"GenericPrompt":     "#d4d4d4",
	"GenericStrong":     "#d4d4d4",
	"GenericSubheading": "#d4d4d4",
	"GenericTraceback":  "#d16969",
	"GenericUnderline":  "#d4d4d4",
	// Keywords
	"Keyword":            "#569cd6",
	"KeywordConstant":    "#569cd6",
	"KeywordDeclaration": "#569cd6",
	"KeywordNamespace":   "#569cd6",
	"KeywordPseudo":      "#569cd6",
	"KeywordReserved":    "#569cd6",
	"KeywordType":        "#569cd6",
	// Literal?
	"Literal":     "#d7ba7d",
	"LiteralDate": "#d7ba7d",
	// Numbers
	"LiteralNumber":            "#b5cea8",
	"LiteralNumberBin":         "#b5cea8",
	"LiteralNumberFloat":       "#b5cea8",
	"LiteralNumberHex":         "#b5cea8",
	"LiteralNumberInteger":     "#b5cea8",
	"LiteralNumberIntegerLong": "#b5cea8",
	"LiteralNumberOct":         "#b5cea8",
	// Operators
	"Operator":     "#c586c0",
	"OperatorWord": "#c586c0",
	// Strings
	"LiteralString":         "#ce9178",
	"LiteralStringBacktick": "#ce9178",
	"LiteralStringChar":     "#d7ba7d",
	"LiteralStringDoc":      "#ce9178",
	"LiteralStringDouble":   "#ce9178",
	"LiteralStringEscape":   "#d7ba7d",
	"LiteralStringHeredoc":  "#ce9178",
	"LiteralStringInterpol": "#569cd6",
	"LiteralStringOther":    "#ce9178",
	"LiteralStringRegex":    "#d16969",
	"LiteralStringSingle":   "#ce9178",
	"LiteralStringSymbol":   "#ce9178",
	// Names
	"Name":                 "#9cdcfe",
	"NameAttribute":        "#9cdcfe",
	"NameBuiltin":          "#569cd6",
	"NameBuiltinPseudo":    "#569cd6",
	"NameClass":            "#4ec9b0",
	"NameConstant":         "#9cdcfe",
	"NameDecorator":        "#dcdcaa",
	"NameEntity":           "#9cdcfe",
	"NameException":        "#9cdcfe",
	"NameFunction":         "#dcdcaa",
	"NameLabel":            "#9cdcfe",
	"NameNamespace":        "#4ec9b0",
	"NameOther":            "#9cdcfe",
	"NameTag":              "#9cdcfe",
	"NameVariable":         "#9cdcfe",
	"NameVariableClass":    "#9cdcfe",
	"NameVariableGlobal":   "#9cdcfe",
	"NameVariableInstance": "#9cdcfe",
}

// HexToRGB takes a hex color value and returns three ints with its red, green, and blue values.
// Source: https://stackoverflow.com/a/54200713/7687024
func HexToRGB(s string) (r, g, b int, err error) {
	var errInvalidFormat = errors.New("invalid format")

	if len(s) == 0 {
		return r, g, b, errInvalidFormat
	}
	if s[0] != '#' {
		return r, g, b, errInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}

	var rVal, gVal, bVal uint8
	switch len(s) {
	case 7:
		rVal = hexToByte(s[1])<<4 + hexToByte(s[2])
		gVal = hexToByte(s[3])<<4 + hexToByte(s[4])
		bVal = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		rVal = hexToByte(s[1]) * 17
		gVal = hexToByte(s[2]) * 17
		bVal = hexToByte(s[3]) * 17
	default:
		err = errInvalidFormat
	}

	r = int(rVal)
	g = int(gVal)
	b = int(bVal)
	return
}
