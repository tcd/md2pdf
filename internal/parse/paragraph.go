package parse

import (
	"github.com/tcd/md2pdf/internal/model"
	"github.com/tcd/md2pdf/internal/renderer"
	"golang.org/x/net/html"
)

// AddParagraph adds a Paragraph Element and any contained Image Elements.
func AddParagraph(r *renderer.Renderables, z *html.Tokenizer) {
	contents, imgTokens := parseP(z)

	r.Add(renderer.Paragraph{
		Content: contents,
	})

	if len(imgTokens) > 0 {
		for _, t := range imgTokens {
			img := Image(t)
			r.Add(img)
		}
	}
}

// Blockquote gathers the data needed to render a blockquote.
func Blockquote(z *html.Tokenizer) renderer.Blockquote {
	contents, _ := parseP(z)
	return renderer.Blockquote{
		Content: contents,
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
			content := string(z.Text())
			if content != "\n" {
				this.AddStr(content)
			}
		}
		if tt == html.StartTagToken {
			newText := model.Text{}
			parseContent(z, z.Token(), newText, &this)
		}
		if tt == html.SelfClosingTagToken {
			T1 := z.Token()
			if T1.Data == "img" {
				imgTokens = append(imgTokens, T1)
			}
			if T1.Data == "br" {
				this.AddStr("\n")
			}
		}
	}

	return this, imgTokens
}
