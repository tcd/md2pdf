package parse

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/render"
	"golang.org/x/net/html"
)

func parseP(pdf *gofpdf.Fpdf, z *html.Tokenizer, blockquote bool) {
	var imageTokens []html.Token
	this := render.Contents{}

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
			this.AddContent(render.Text{Content: string(z.Text())})
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

	if blockquote {
		// render.FullP(pdf, this)
		render.Blockquote(pdf, this)
	} else {
		render.FullP(pdf, this)
	}

	if len(imageTokens) > 0 {
		for _, t := range imageTokens {
			parseImg(pdf, t)
		}
	}
}
