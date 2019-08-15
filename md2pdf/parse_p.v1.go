package md2pdf

import (
	"bytes"
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/ghfm"
	"golang.org/x/net/html"
)

var opts = []string{
	"a",
	"br",
	"blockquote",
	"code",
	"del",
	"em",
	"h1",
	"h2",
	"h3",
	"h4",
	"h5",
	"h6",
	"hr",
	"img",
	"p",
	"pre",
	"strong",

	"ol",
	"ul",
	"li",

	"dl",
	"dt",
	"dd",

	"table",
	"thead",
	"tr",
	"th",
	"tbody",
	"td",
}

// Tag is an html tag.
type Tag struct {
	Kind       string
	Contents   [][]byte
	Attributes map[string]string
	Children   []Tag
}

// AddChild to a Tag.
func (t *Tag) AddChild(tag Tag) {
	t.Children = append(t.Children, tag)
}

// AddContents appends bytes to a Tags Contents.
func (t *Tag) AddContents(b []byte) {
	t.Contents = append(t.Contents, b)
}

// ConsolidateContents of a Tag and it's children.
func (t Tag) ConsolidateContents() string {
	joinChar := []byte{' '}
	joined := bytes.Join(t.Contents, joinChar)
	cleaned := bytes.Replace(joined, []byte("  "), []byte(" "), -1)
	cleaned = bytes.Replace(cleaned, []byte("  "), []byte(" "), -1)
	cleaned = bytes.Replace(cleaned, []byte(" ."), []byte("."), -1)
	return string(cleaned)
}

func parseP(pdf *gofpdf.Fpdf, z *html.Tokenizer) {
	topLevel := Tag{
		Kind: "p",
	}

	for {
		ct := z.Token()
		if ct.Type == html.TextToken {
			topLevel.AddContents(z.Text())
			tt2 := z.Next()
			if tt2 == html.StartTagToken {
				if parseChild(z, &topLevel) {
					break
				}
			}
		}
		if ct.Type == html.EndTagToken {
			if ct.Data == "p" {
				break
			}
		}

		tt := z.Next()
		if tt == html.TextToken {
			topLevel.AddContents(z.Text())
			tt2 := z.Next()
			if tt2 == html.StartTagToken {
				if parseChild(z, &topLevel) {
					break
				}
			}
		}
		if tt == html.EndTagToken {
			T2 := z.Token()
			if T2.Data == "p" {
				break
			}
		}
		if tt == html.StartTagToken {
			if parseChild(z, &topLevel) {
				break
			}
		}
	}

	fmt.Println(topLevel.ConsolidateContents())
	ghfm.BasicP(pdf, topLevel.ConsolidateContents())
}

func parseChild(z *html.Tokenizer, above *Tag) bool {
	this := Tag{}
	currentToken := z.Token()
	for _, v := range opts {
		if currentToken.Data == v {
			this.Kind = v
		}
	}

	tt := z.Next()
	if tt == html.TextToken {
		above.AddContents(z.Text())

	}
	if tt == html.SelfClosingTagToken {
		// Not doing anything here yet.
	}
	if tt == html.StartTagToken {
		parseChild(z, above)
	}
	if tt == html.EndTagToken {
		T1 := z.Token()
		if T1.Data == "p" {
			return true
		}
		if T1.Data == this.Kind {
			return false
		}
	}
	return false
}

func parseImg(token html.Token) Tag {
	var src, alt, title string
	for _, a := range token.Attr {
		if a.Key == "src" {
			src = a.Val
		}
		if a.Key == "alt" {
			alt = a.Val
		}
		if a.Key == "title" {
			title = a.Val
		}
	}
	return Tag{
		Kind: "img",
		Attributes: map[string]string{
			"src":   src,
			"alt":   alt,
			"title": title,
		},
	}
}
