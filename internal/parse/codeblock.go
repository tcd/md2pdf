package parse

import (
	"github.com/tcd/md2pdf/internal/model"
	"github.com/tcd/md2pdf/internal/renderer"
	"golang.org/x/net/html"
)

// Codeblock gathers the data needed to render a codeblock.
func Codeblock(z *html.Tokenizer) renderer.Codeblock {
	content, class := parseCodeblock(z)

	return renderer.Codeblock{
		Class:   class,
		Content: content,
	}
}

func parseCodeblock(z *html.Tokenizer) (model.Contents, string) {
	tt := z.Next()
	this := model.Contents{}
	var class string

	if tt == html.TextToken {
		content := z.Text()
		this.AddStr(string(content))
	}

	if tt == html.StartTagToken {
		T1 := z.Token()
		if T1.Data == "code" {
			for _, a := range T1.Attr {
				if a.Key == "class" {
					class = a.Val
				}
			}
			tt2 := z.Next()
			if tt2 == html.TextToken {
				content := z.Text()
				this.AddStr(string(content))
			}
		}
	}

	return this, class
}
