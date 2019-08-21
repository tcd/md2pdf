package parse

import (
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/render"
	"golang.org/x/net/html"
)

// HTML renders a PDF from an html string generated by blackfriday
func HTML(inputHTML, outputPath string) error {

	doc := strings.NewReader(inputHTML)
	tokenizer := html.NewTokenizer(doc)
	pdf := gofpdf.New("P", "mm", "Letter", "")
	render.Setup(pdf)

	for {
		tt := tokenizer.Next()

		if tt == html.ErrorToken {
			break // End of document
		}

		if tt == html.SelfClosingTagToken {
			T1 := tokenizer.Token()

			if T1.Data == "hr" {
				render.HR(pdf)
			}
			if T1.Data == "br" {
				// decide how to deal with these?
			}
			if T1.Data == "img" {
				parseImg(pdf, T1)
			}
		}

		if tt == html.StartTagToken {
			T1 := tokenizer.Token()

			if T1.Data == "h1" ||
				T1.Data == "h2" ||
				T1.Data == "h3" ||
				T1.Data == "h4" ||
				T1.Data == "h5" ||
				T1.Data == "h6" {
				parseHeader(pdf, tokenizer, T1)
			}
			if T1.Data == "pre" {
				parsePre(pdf, tokenizer)
			}
			if T1.Data == "p" {
				parseP(pdf, tokenizer)
			}
			if T1.Data == "blockquote" {
				// parseBlockquote(pdf, tokenizer)
			}
			if T1.Data == "ol" || T1.Data == "ul" {
				// parseList(pdf, tokenizer)
			}
			if T1.Data == "dl" {
				// Deal with definition lists?
			}
			if T1.Data == "table" {
				parseTable(pdf, tokenizer)
			}
		}
	}

	return pdf.OutputFileAndClose(outputPath)
}

func parseContent(
	z *html.Tokenizer,
	startToken html.Token,
	parent render.Text,
	topLevel *render.Contents,
) {
	this := parent.Copy()
	currentKind := startToken.Data

	if currentKind == "a" {
		for _, a := range startToken.Attr {
			if a.Key == "href" {
				this.HREF = a.Val
			}
		}
	}
	if currentKind == "strong" {
		this.Bold = true
	}
	if currentKind == "em" {
		this.Italic = true
	}
	if currentKind == "code" {
		this.Code = true
	}
	if currentKind == "del" {
		this.Strike = true
	}

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}
		if tt == html.EndTagToken {
			T2 := z.Token()
			if T2.Data == currentKind {
				break
			}
		}
		if tt == html.StartTagToken {
			T2 := z.Token()
			parseContent(z, T2, this, topLevel)
		}
		if tt == html.TextToken {
			this.Content = string(z.Text())
		}
		if tt == html.SelfClosingTagToken {
			continue
		}
	}

	topLevel.AddContent(this)
}
