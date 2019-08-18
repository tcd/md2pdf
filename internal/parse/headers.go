package parse

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/render"
	"golang.org/x/net/html"
)

func parseH1(pdf *gofpdf.Fpdf, z *html.Tokenizer) {
	tt := z.Next()
	if tt == html.TextToken {
		content := z.Text()
		render.H1(pdf, string(content))
	}
}
func parseH2(pdf *gofpdf.Fpdf, z *html.Tokenizer) {
	tt := z.Next()
	if tt == html.TextToken {
		content := z.Text()
		render.H2(pdf, string(content))
	}
}
func parseH3(pdf *gofpdf.Fpdf, z *html.Tokenizer) {
	tt := z.Next()
	if tt == html.TextToken {
		content := z.Text()
		render.H3(pdf, string(content))
	}
}
func parseH4(pdf *gofpdf.Fpdf, z *html.Tokenizer) {
	tt := z.Next()
	if tt == html.TextToken {
		content := z.Text()
		render.H4(pdf, string(content))
	}
}
func parseH5(pdf *gofpdf.Fpdf, z *html.Tokenizer) {
	tt := z.Next()
	if tt == html.TextToken {
		content := z.Text()
		render.H5(pdf, string(content))
	}
}
func parseH6(pdf *gofpdf.Fpdf, z *html.Tokenizer) {
	tt := z.Next()
	if tt == html.TextToken {
		content := z.Text()
		render.H6(pdf, string(content))
	}
}
