package parse

import (
	"log"

	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/render"
	"golang.org/x/net/html"
)

func parseHeader(pdf *gofpdf.Fpdf, z *html.Tokenizer, startToken html.Token) {
	headerLevel := startToken.Data

	tt := z.Next()
	if tt == html.TextToken {
		content := string(z.Text())
		switch headerLevel {
		case "h1":
			render.H1(pdf, content)
		case "h2":
			render.H2(pdf, content)
		case "h3":
			render.H3(pdf, content)
		case "h4":
			render.H4(pdf, content)
		case "h5":
			render.H5(pdf, content)
		case "h6":
			render.H6(pdf, content)
		default:
			log.Printf("Error parsing heaader with content: %q\n", content)
			return
		}
	}
}
