package parse

import (
	"log"

	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/render"
	"golang.org/x/net/html"
)

func parseHeader(pdf *gofpdf.Fpdf, z *html.Tokenizer, startToken html.Token) {
	var imageTokens []html.Token
	this := render.Contents{}

	headerLevel := startToken.Data

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}
		if tt == html.EndTagToken {
			T1 := z.Token()
			if T1.Data == headerLevel {
				break
			}
		}
		if tt == html.TextToken {
			this.AddStr(string(z.Text()))
		}
		if tt == html.StartTagToken {
			newText := render.Text{}
			parseContent(z, z.Token(), newText, &this)
		}
		if tt == html.SelfClosingTagToken {
			T1 := z.Token()
			if T1.Data == "img" {
				imageTokens = append(imageTokens, T1)
			} else {
				continue
			}
		}
	}

	switch headerLevel {
	case "h1":
		render.H1(pdf, this)
	case "h2":
		render.H2(pdf, this)
	case "h3":
		render.H3(pdf, this)
	case "h4":
		render.H4(pdf, this)
	case "h5":
		render.H5(pdf, this)
	case "h6":
		render.H6(pdf, this)
	default:
		log.Printf("Error parsing heaader with content: %v\n", this.Content)
		return
	}
}
