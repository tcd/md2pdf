package parse

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/render"
	"golang.org/x/net/html"
)

func parseImg(pdf *gofpdf.Fpdf, token html.Token) {
	var src string
	for _, a := range token.Attr {
		if a.Key == "src" {
			src = a.Val
		}
	}
	render.Image(pdf, src)
}
