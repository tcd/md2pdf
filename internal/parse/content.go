package parse

import (
	"github.com/tcd/md2pdf/internal/model"
	"golang.org/x/net/html"
)

// recursively extract content from elements in list items, paragraphs, etc.
func parseContent(
	z *html.Tokenizer,
	startToken html.Token,
	parent model.Text,
	topLevel *model.Contents,
) {
	this := parent.Copy()
	currentKind := startToken.Data

	if currentKind == "p" {
		return
	}

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
			this.Text = string(z.Text())
		}
		if tt == html.SelfClosingTagToken {
			continue
		}
	}

	topLevel.AddContent(this)
}
