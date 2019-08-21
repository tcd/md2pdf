package parse

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/render"
	"golang.org/x/net/html"
)

func parseP(pdf *gofpdf.Fpdf, z *html.Tokenizer) {
	var imageTokens []html.Token
	this := render.Contents{}

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}
		if tt == html.EndTagToken {
			T2 := z.Token()
			if T2.Data == "p" {
				break
			}
		}
		if tt == html.TextToken {
			this.AddContent(render.Text{Content: string(z.Text())})
		}
		if tt == html.StartTagToken {
			newText := render.Text{}
			parseContent(z, z.Token(), newText, &this)
		}
		if tt == html.SelfClosingTagToken {
			T2 := z.Token()
			if T2.Data == "img" {
				imageTokens = append(imageTokens, T2)
			} else {
				continue
			}
		}
	}

	render.FullP(pdf, this)
	if len(imageTokens) > 0 {
		for _, t := range imageTokens {
			parseImg(pdf, t)
		}
	}
}
