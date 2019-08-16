package md2pdf

import (
	"github.com/jung-kurt/gofpdf"
	. "github.com/tcd/md2pdf/internal/ghfm"
	"golang.org/x/net/html"
)

// func parse(pdf *gofpdf.Fpdf, z *html.Tokenizer) {}

func parseEm(pdf *gofpdf.Fpdf, z *html.Tokenizer) {}

func parseStrong(pdf *gofpdf.Fpdf, z *html.Tokenizer) {}

func parseDel(pdf *gofpdf.Fpdf, z *html.Tokenizer) {}

func parseImg(pdf *gofpdf.Fpdf, token html.Token) {
	var src string
	for _, a := range token.Attr {
		if a.Key == "src" {
			src = a.Val
		}
	}
	Image(pdf, src)
}

func parsePre(pdf *gofpdf.Fpdf, z *html.Tokenizer) {
	tt := z.Next()
	if tt == html.StartTagToken {
		T1 := z.Token()
		if T1.Data == "code" {
			tt2 := z.Next()
			if tt2 == html.TextToken {
				content := z.Text()
				CodeBlock(pdf, string(content))
			}
		}
	}
	if tt == html.TextToken {
		content := z.Text()
		CodeBlock(pdf, string(content))
	}
}

func parseH1(pdf *gofpdf.Fpdf, z *html.Tokenizer) {
	tt := z.Next()
	if tt == html.TextToken {
		content := z.Text()
		H1(pdf, string(content))
	}
}
func parseH2(pdf *gofpdf.Fpdf, z *html.Tokenizer) {
	tt := z.Next()
	if tt == html.TextToken {
		content := z.Text()
		H2(pdf, string(content))
	}
}
func parseH3(pdf *gofpdf.Fpdf, z *html.Tokenizer) {
	tt := z.Next()
	if tt == html.TextToken {
		content := z.Text()
		H3(pdf, string(content))
	}
}
func parseH4(pdf *gofpdf.Fpdf, z *html.Tokenizer) {
	tt := z.Next()
	if tt == html.TextToken {
		content := z.Text()
		H4(pdf, string(content))
	}
}
func parseH5(pdf *gofpdf.Fpdf, z *html.Tokenizer) {
	tt := z.Next()
	if tt == html.TextToken {
		content := z.Text()
		H5(pdf, string(content))
	}
}
func parseH6(pdf *gofpdf.Fpdf, z *html.Tokenizer) {
	tt := z.Next()
	if tt == html.TextToken {
		content := z.Text()
		H6(pdf, string(content))
	}
}
