package parse

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/render"
	"golang.org/x/net/html"
)

func parseCodeblock(pdf *gofpdf.Fpdf, z *html.Tokenizer) {
	tt := z.Next()
	if tt == html.StartTagToken {
		T1 := z.Token()
		if T1.Data == "code" {
			tt2 := z.Next()
			if tt2 == html.TextToken {
				content := z.Text()
				render.CodeBlock(pdf, string(content))
			}
		}
	}
	if tt == html.TextToken {
		content := z.Text()
		render.CodeBlock(pdf, string(content))
	}
}
