package parse

import (
	"bytes"
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/render"
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
				done := parseChild(z, &topLevel)
				if done {
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
				done := parseChild(z, &topLevel)
				if done {
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
			done := parseChild(z, &topLevel)
			if done {
				break
			}
		}
	}

	fmt.Println(topLevel.ConsolidateContents())
	render.BasicP(pdf, topLevel.ConsolidateContents())
}

func parseChild(z *html.Tokenizer, parent *Tag) bool {
	this := Tag{}
	currentToken := z.Token()
	for _, v := range opts {
		if currentToken.Data == v {
			this.Kind = v
		}
	}

	tt := z.Next()
	if tt == html.TextToken {
		this.AddContents(z.Text())
		tt2 := z.Next()
		if tt2 == html.EndTagToken {
			T2 := z.Token()
			if T2.Data == this.Kind {
				parent.AddChild(this)
				return false
			}
		}
	}

	if tt == html.StartTagToken {
		parseChild(z, &this)
	}
	if tt == html.EndTagToken {
		T1 := z.Token()
		if T1.Data == "p" {
			return true
		}
		if T1.Data == this.Kind {
			parent.AddChild(this)
			return false
		}
	}
	if tt == html.SelfClosingTagToken {
		// Not dealing with this yet.
	}
	return false
}
