package parse

import (
	"log"

	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/model"
	"github.com/tcd/md2pdf/internal/render"
	"golang.org/x/net/html"
)

// Header ...
// TODO: Deal with styles in headers.
// TODO: Deal with images in headers.
func Header(pdf *gofpdf.Fpdf, z *html.Tokenizer, startToken html.Token) {
	level, content, _ := parseHeader(z, startToken)

	switch level {
	case "h1":
		render.H1(pdf, content)
	case "h2":
		render.H2(pdf, content)
	case "h3":
		render.H3(pdf, content)
	case "h4":
		render.H4(pdf, content)
	case "h5":
		render.H5(pdf, content)
	case "h6":
		render.H6(pdf, content)
	default:
		log.Printf("Error parsing header with content: %v\n", content)
		return
	}
}

func parseHeader(z *html.Tokenizer, startToken html.Token) (headerLevel string, contents model.Contents, imgTokens []html.Token) {
	headerLevel = startToken.Data
	this := model.Contents{}

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
			newText := model.Text{}
			parseContent(z, z.Token(), newText, &this)
		}
		if tt == html.SelfClosingTagToken {
			T1 := z.Token()
			if T1.Data == "img" {
				imgTokens = append(imgTokens, T1)
			} else {
				continue
			}
		}
	}

	contents = this
	return
}
