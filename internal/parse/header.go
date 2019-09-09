package parse

import (
	"github.com/tcd/md2pdf/internal/model"
	"github.com/tcd/md2pdf/internal/renderer"
	"golang.org/x/net/html"
)

// Header gathers the data needed to render a header.
// TODO: Deal with images in headers.
func Header(z *html.Tokenizer, startToken html.Token) renderer.Header {
	level, content, _ := parseHeader(z, startToken)

	return renderer.Header{
		Type:    "header",
		Level:   level,
		Content: content,
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
