package ghfm

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

// GitHubPDF attempts to recreate this page (https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet) as a pdf.
func GitHubPDF(path string) {
	pdf := gofpdf.New(
		"P",      // orientationStr: "P" or "L" (Portrait or Landscape
		"mm",     // unitStr: "pt", "mm", "cm", or "in" (point, millimeter, centimeter, or inches)
		"Letter", // sizeStr: "A3", "A4", "A5", "Letter", "Legal", or "Tabloid".
		"",       // fontDirStr: Empty is fine when using core fonts.
	)
	pdf.AddPage()
	Setup(pdf)

	H1(pdf, "GitHub Markdown PDF demo")
	H2(pdf, "Headers")
	CodeBlock(pdf, cbContent())
	H1(pdf, "H1")
	H2(pdf, "H2")
	H3(pdf, "H3")
	H4(pdf, "H4")
	H5(pdf, "H5")
	H6(pdf, "H6")
	p(pdf, "Alternatively, for H1 and H2, an underline-ish style:")
	H1(pdf, "Alt-H1")
	H2(pdf, "Alt-H2")
	H2(pdf, "Emphasis")
	CodeBlock(pdf, emContent())
	span(pdf, "Emphasis, aka italics, with ")
	em(pdf, "This is a test string")
	H2(pdf, "Lists")
	p(pdf, "(In this example, leading and trailing spaces are shown with with dots: .)")
	CodeBlock(pdf, listContent())
	H2(pdf, "Links")
	CodeBlock(pdf, linksContent())
	link(pdf, "I'm an inline-style link", "https://www.google.com")
	H2(pdf, "Images")
	CodeBlock(pdf, imagesContent())
	H2(pdf, "Code and Syntax Highlighting")
	CodeBlock(pdf, inlineCodeContent())
	CodeBlock(pdf, syntaxHighContent())
	H2(pdf, "Tables")
	CodeBlock(pdf, tablesContent())
	H2(pdf, "Blockquotes")
	CodeBlock(pdf, blockquoteContent())
	H2(pdf, "Horizontal Rule")
	CodeBlock(pdf, hrContent())
	H2(pdf, "Line Breaks")

	err := pdf.OutputFileAndClose(path)
	if err != nil {
		fmt.Println(err)
	}
}

// ParagraphTest is just a test.
func ParagraphTest(path string) {
	pdf := gofpdf.New(
		"P",      // orientationStr: "P" or "L" (Portrait or Landscape
		"mm",     // unitStr: "pt", "mm", "cm", or "in" (point, millimeter, centimeter, or inches)
		"Letter", // sizeStr: "A3", "A4", "A5", "Letter", "Legal", or "Tabloid".
		"",       // fontDirStr: Empty is fine when using core fonts.
	)
	setup(pdf)

	H1(pdf, "Paragraph Test")

	span(pdf, "Some plain text")
	code2(pdf, "some code hehe lol")
	code2(pdf, "x > 42")
	span(pdf, "Some plain text")
	strong(pdf, "Some bold text")
	span(pdf, lorem())

	err := pdf.OutputFileAndClose(path)
	if err != nil {
		fmt.Println(err)
	}
}

// ListTest is just a test.
func ListTest(path string) {
	pdf := gofpdf.New("P", "mm", "Letter", "")
	Setup(pdf)
}
