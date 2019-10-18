package parse

import (
	"bytes"

	"github.com/tcd/md2pdf/internal/renderer"
	"golang.org/x/net/html"
)

// Parse gathers the data needed to render a PDF.
func Parse(inputHTML []byte) renderer.Renderables {
	rbs := renderer.Renderables{}
	doc := bytes.NewReader(inputHTML)
	tokenizer := html.NewTokenizer(doc)

	for {
		tt := tokenizer.Next()

		if tt == html.ErrorToken {
			break // End of document
		}

		if tt == html.SelfClosingTagToken {
			T1 := tokenizer.Token()

			if T1.Data == "hr" {
				rbs.Add(renderer.HR{})
			}
			if T1.Data == "img" {
				img := Image(T1)
				rbs.Add(img)
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
				rbs.Add(header)
			}
			if T1.Data == "pre" {
				cb := Codeblock(tokenizer)
				rbs.Add(cb)
			}
			if T1.Data == "p" {
				AddParagraph(&rbs, tokenizer)
			}
			if T1.Data == "blockquote" {
				bq := Blockquote(tokenizer)
				rbs.Add(bq)
			}
			if T1.Data == "ol" || T1.Data == "ul" {
				ls := List(tokenizer, T1)
				rbs.Add(ls)
			}
			if T1.Data == "table" {
				table := Table(tokenizer)
				rbs.Add(table)
			}
		}
	}

	return rbs
}
