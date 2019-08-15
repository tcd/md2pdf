package md2pdf

import (
	"github.com/jung-kurt/gofpdf"
	. "github.com/tcd/md2pdf/internal/ghfm"
	"golang.org/x/net/html"
)

// func parse(pdf *gofpdf.Fpdf, z *html.Tokenizer) {}

// func parseA(pdf *gofpdf.Fpdf, z *html.Tokenizer, token html.Token) {
// 	var content, href string
//
// 	for _, a := range token.Attr {
// 		if a.Key == "href" {
// 			href = a.Val
// 		}
// 	}
//
// 	tt := z.Next()
// 	if tt == html.StartTagToken {
// 		fmt.Println("Aint nobody got time for that")
// 	}
// 	if tt == html.TextToken {
// 		T1 := z.Text()
// 		tt2 := z.Next()
// 		if tt2 == html.EndTagToken {
// 			content = string(T1)
// 		} else {
// 			fmt.Println("Aint nobody got time for that")
// 		}
// 	}
//
// 	// link(pdf, content, href)
// }

func parseEm(pdf *gofpdf.Fpdf, z *html.Tokenizer) {}

func parseStrong(pdf *gofpdf.Fpdf, z *html.Tokenizer) {}

func parseDel(pdf *gofpdf.Fpdf, z *html.Tokenizer) {}

func parsePre(pdf *gofpdf.Fpdf, z *html.Tokenizer) {
	tt := z.Next()
	if tt == html.StartTagToken {
		T1 := z.Token()
		if T1.Data == "code" {
			// for _, a := range T1.Attr {
			// 	if a.Key == "class" {
			// 		class := a.Val // Use this for syntax highlighting.
			// 	}
			// }
			tt2 := z.Next()
			if tt2 == html.TextToken {
				content := z.Text()
				CodeBlock(pdf, string(content))
			}
			// tt3 := z.Next()
			// if tt3 == html.EndTagToken {
			// 	if T1.Data == "code" {
			// 		tt4 := z.Next()
			// 		if tt4 == html.EndTagToken {
			// 			if T1.Data == "pre" {
			// 				CodeBlock(pdf, string(content))
			// 			}
			// 		}
			// 	}
			// }
		}
	}

	// if tt == html.TextToken {
	// 	content := z.Text()
	// 	CodeBlock(pdf, string(content))
	// }
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
