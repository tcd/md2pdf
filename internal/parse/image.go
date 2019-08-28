package parse

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/render"
	"golang.org/x/net/html"
)

// Image ...
func Image(pdf *gofpdf.Fpdf, token html.Token) {
	src := parseImg(token)
	render.Image(pdf, src)
}

func parseImg(token html.Token) string {
	var src string
	for _, a := range token.Attr {
		if a.Key == "src" {
			src = a.Val
		}
	}
	return src
}
