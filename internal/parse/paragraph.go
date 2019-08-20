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
			parseChild(z, z.Token(), &newText, &this)
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

	render.FullP(pdf, this.Content)
	if len(imageTokens) > 0 {
		for _, t := range imageTokens {
			parseImg(pdf, t)
		}
	}
}

// Add a nested tags content to a Paragraph, return true if it finds a closing `</p>` tag.
func parseChild(z *html.Tokenizer, startToken html.Token, parent *render.Text, topLevel *render.Contents) {
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
			parseChild(z, T2, this, topLevel)
		}
		if tt == html.TextToken {
			this.Content = string(z.Text())
		}
		if tt == html.SelfClosingTagToken {
			continue
		}
	}

	topLevel.AddContent(*this)
}
