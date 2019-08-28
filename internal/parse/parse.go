package parse

import (
	"strings"

	"github.com/tcd/md2pdf/internal/renderable"
	"golang.org/x/net/html"
)

// Parse gathers the data needed to render a PDF.
func Parse(inputHTML, outputPath string) renderable.Elements {
	elements := renderable.Elements{}

	doc := strings.NewReader(inputHTML)
	tokenizer := html.NewTokenizer(doc)

	for {
		tt := tokenizer.Next()

		if tt == html.ErrorToken {
			break // End of document
		}

		if tt == html.SelfClosingTagToken {
			T1 := tokenizer.Token()

			if T1.Data == "hr" {
				elements.Add(renderable.HR{})
			}
			if T1.Data == "img" {
				img := Image(T1)
				elements.Add(img)
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
				header := Header(tokenizer, T1)
				elements.Add(header)
			}
			if T1.Data == "pre" {
				cb := Codeblock(tokenizer)
				elements.Add(cb)
			}
			if T1.Data == "p" {
				// p := Paragraph(tokenizer)
				// elements.Add(p)
				AddParagraph(&elements, tokenizer)
			}
			if T1.Data == "blockquote" {
				bq := Blockquote(tokenizer)
				elements.Add(bq)
			}
			if T1.Data == "ol" || T1.Data == "ul" {
				ls := List(tokenizer)
				elements.Add(ls)
			}
			if T1.Data == "table" {
				table := Table(tokenizer)
				elements.Add(table)
			}
		}
	}

	return elements
}
