package parse

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/model"
	"github.com/tcd/md2pdf/internal/render"
	"golang.org/x/net/html"
)

// List ...
func List(pdf *gofpdf.Fpdf, z *html.Tokenizer) {
	content := parseEntries(z)
	render.AnyList(pdf, content)
}

func parseEntries(z *html.Tokenizer) model.ListContent {
	this := model.ListContent{}
	if z.Token().Data == "ol" {
		this.Ordered = true
	}

	for {
		tt := z.Next()
		if tt == html.EndTagToken {
			T1 := z.Token()
			if T1.Data == "ol" || T1.Data == "ul" {
				break
			}
		}
		if tt == html.StartTagToken {
			T1 := z.Token()
			if T1.Data == "li" {
				newEntry := parseEntry(z)
				this.AddItems(newEntry)
			}
		}
	}

	return this
}

func parseEntry(z *html.Tokenizer) model.ListItem {
	this := model.ListItem{}

	for {
		tt := z.Next()
		if tt == html.TextToken {
			content := string(z.Text())
			this.Contents.AddStr(content)
			// if content != "\n" {
			// 	if len(content) >= 3 {
			// 		if content[len(content)-2:] == "\n\n" {
			// 			this.Contents.AddStr(content[:len(content)-2])
			// 		}
			// 	}
			// } else {
			// 	this.Contents.AddStr(content)
			// }
		}
		if tt == html.EndTagToken {
			T1 := z.Token()
			if T1.Data == "li" {
				break
			}
		}
		if tt == html.StartTagToken {
			T1 := z.Token()
			if T1.Data != "ul" && T1.Data != "ol" {
				blankContent := model.Text{}
				parseContent(z, T1, blankContent, &this.Contents)
			}
			if T1.Data == "ul" || T1.Data == "ol" {
				nestedList := parseEntries(z)
				this.Children = nestedList
			}
		}
	}

	return this
}
