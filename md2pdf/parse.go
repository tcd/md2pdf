package md2pdf

import (
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/ghfm"
	"golang.org/x/net/html"
)

// Parse renders a PDF from an html string generated by blackfriday
func parse(inputHTML, outputPath string) error {

	doc := strings.NewReader(inputHTML)
	tokenizer := html.NewTokenizer(doc)
	pdf := gofpdf.New("P", "mm", "Letter", "")
	ghfm.Setup(pdf)

	for {
		tt := tokenizer.Next()

		if tt == html.ErrorToken {
			break
		}

		if tt == html.SelfClosingTagToken {
			T1 := tokenizer.Token()

			if T1.Data == "hr" {
				ghfm.HR(pdf)
			}
		}

		if tt == html.StartTagToken {
			T1 := tokenizer.Token()

			if T1.Data == "h1" {
				parseH1(pdf, tokenizer)
			}
			if T1.Data == "h2" {
				parseH2(pdf, tokenizer)
			}
			if T1.Data == "h3" {
				parseH3(pdf, tokenizer)
			}
			if T1.Data == "h4" {
				parseH4(pdf, tokenizer)
			}
			if T1.Data == "h5" {
				parseH5(pdf, tokenizer)
			}
			if T1.Data == "h6" {
				parseH6(pdf, tokenizer)
			}
			if T1.Data == "pre" {
				parsePre(pdf, tokenizer)
			}
			if T1.Data == "p" {
				parseP(pdf, tokenizer)
			}
		}
	}

	err := pdf.OutputFileAndClose(outputPath)
	if err != nil {
		return err
	}
	return nil
}
