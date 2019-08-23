package parse

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/render"
	"golang.org/x/net/html"
)

func parseCodeblock(pdf *gofpdf.Fpdf, z *html.Tokenizer) {
	tt := z.Next()
	this := render.Contents{}
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

	if class == "language-no-highlight" || class == "" {
		render.CodeBlock(pdf, this)
	} else {
		render.CodeBlock(pdf, this)
	}
}
