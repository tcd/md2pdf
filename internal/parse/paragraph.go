package parse

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/model"
	"github.com/tcd/md2pdf/internal/render"
	"golang.org/x/net/html"
)

// Paragraph ...
func Paragraph(pdf *gofpdf.Fpdf, z *html.Tokenizer, blockquote bool) {
	contents, imgTokens := parseP(z)

	if blockquote {
		render.Blockquote(pdf, contents)
	} else {
		render.FullP(pdf, contents)
	}

	if len(imgTokens) > 0 {
		for _, t := range imgTokens {
			Image(pdf, t)
		}
	}
}

func parseP(z *html.Tokenizer) (model.Contents, []html.Token) {
	var imgTokens []html.Token
	this := model.Contents{}

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}
		if tt == html.EndTagToken {
			T1 := z.Token()
			if T1.Data == "p" || T1.Data == "blockquote" {
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

	return this, imgTokens
}
