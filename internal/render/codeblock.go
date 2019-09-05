package render

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"

	"github.com/alecthomas/chroma/quick"
	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/lib"
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
	_, _, _, oldBottomMargin := pdf.GetMargins()

	pdf.SetCellMargin(5)
	pdf.SetAutoPageBreak(true, 0)
	pdf.SetFont("courier", "", 11)
	pdf.SetFillColor(30, 30, 30)
	_, lineHt := pdf.GetFontSize()
	lineHt *= 1.25

	// Draw out the text on a temporary PDF to see how much space it takes up.
	testPDF := gofpdf.New("P", "mm", "Letter", "")
	Setup(testPDF)
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
	if y1+blockHt > lib.ContentBoxHeight(pdf) {
		pdf.AddPage()
		x1, y1 = pdf.GetXY()
	}
	// Draw the background.
	pdf.Rect(x1, y1, lib.ContentBoxWidth(pdf), blockHt, "F")

	// Write out the content for real this time.
	pdf.SetXY(x1, y1)
	pdf.Ln(lineHt)
	for _, token := range tokens {
		r, g, b, _ := lib.HexToRGB(styleMap[token.Type])
		pdf.SetTextColor(r, g, b)
		pdf.Write(lineHt, lib.CleanString(token.Value))
	}
	pdf.Ln(lineHt)

	pdf.SetCellMargin(oldCellMargin)
	pdf.SetAutoPageBreak(true, oldBottomMargin)
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

var (
	black       = "#1e1e1e"
	white       = "#d4d4d4"
	red         = "#d16969"
	orange      = "#ce9178"
	lightYellow = "#dcdcaa"
	darkYellow  = "#d7ba7d"
	lightGreen  = "#b5cea8"
	green       = "#608b4e"
	cyan        = "#4ec9b0"
	lightBlue   = "#9cdcfe"
	darkBlue    = "#569cd6"
	magenta     = "#c586c0"
)

// DarkPlus Style
var styleMap = map[string]string{
	// Other
	"Background":     black,
	"Text":           white,
	"TextWhitespace": white,
	"Other":          white,
	"Punctuation":    white,
	"Error":          red,
	// Comments
	"Comment":          green,
	"CommentHashbang":  green,
	"CommentMultiline": green,
	"CommentPreproc":   green,
	"CommentSingle":    green,
	"CommentSpecial":   green,
	// Generic
	"Generic":           white,
	"GenericDeleted":    red,
	"GenericEmph":       white,
	"GenericError":      red,
	"GenericHeading":    white,
	"GenericInserted":   green,
	"GenericOutput":     white,
	"GenericPrompt":     white,
	"GenericStrong":     white,
	"GenericSubheading": white,
	"GenericTraceback":  red,
	"GenericUnderline":  white,
	// Keywords
	"Keyword":            darkBlue,
	"KeywordConstant":    darkBlue,
	"KeywordDeclaration": darkBlue,
	"KeywordNamespace":   darkBlue,
	"KeywordPseudo":      darkBlue,
	"KeywordReserved":    darkBlue,
	"KeywordType":        darkBlue,
	// Literal?
	"Literal":     darkYellow,
	"LiteralDate": darkYellow,
	// Numbers
	"LiteralNumber":            lightGreen,
	"LiteralNumberBin":         lightGreen,
	"LiteralNumberFloat":       lightGreen,
	"LiteralNumberHex":         lightGreen,
	"LiteralNumberInteger":     lightGreen,
	"LiteralNumberIntegerLong": lightGreen,
	"LiteralNumberOct":         lightGreen,
	// Operators
	"Operator":     magenta,
	"OperatorWord": magenta,
	// Strings
	"LiteralString":         orange,
	"LiteralStringBacktick": orange,
	"LiteralStringChar":     darkYellow,
	"LiteralStringDoc":      orange,
	"LiteralStringDouble":   orange,
	"LiteralStringEscape":   darkYellow,
	"LiteralStringHeredoc":  orange,
	"LiteralStringInterpol": darkBlue,
	"LiteralStringOther":    orange,
	"LiteralStringRegex":    red,
	"LiteralStringSingle":   orange,
	"LiteralStringSymbol":   orange,
	// Names
	"Name":                 lightBlue,
	"NameAttribute":        lightBlue,
	"NameBuiltin":          darkBlue,
	"NameBuiltinPseudo":    darkBlue,
	"NameClass":            cyan,
	"NameConstant":         lightBlue,
	"NameDecorator":        lightYellow,
	"NameEntity":           lightBlue,
	"NameException":        lightBlue,
	"NameFunction":         lightYellow,
	"NameLabel":            lightBlue,
	"NameNamespace":        cyan,
	"NameOther":            lightBlue,
	"NameTag":              darkBlue,
	"NameVariable":         lightBlue,
	"NameVariableClass":    lightBlue,
	"NameVariableGlobal":   lightBlue,
	"NameVariableInstance": lightBlue,
}
