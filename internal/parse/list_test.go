package parse

import (
	"log"
	"strings"
	"testing"

	"github.com/tcd/md2pdf/internal/model"
	"golang.org/x/net/html"
)

func TestParseList(t *testing.T) {
	input := `<ul>
		<li><a href="https://github.com/jgm/pandoc">jgm/pandoc</a> - Universal markup converter</li>
		<li><a href="https://github.com/mandolyte/mdtopdf">mandolyte/mdtopdf</a> - Markdown to PDF</li>
		<li><a href="https://github.com/ajstarks/deck">ajstarks/deck</a> - Slide Decks</li>
		<li><a href="https://github.com/johnfercher/maroto">johnfercher/maroto</a> - A maroto way to create PDFs.</li>
	</ul>`
	var want model.ListContent
	want.NewItem(
		model.Text{
			Text: "jgm/pandoc",
			HREF: "https://github.com/jgm/pandoc",
		},
		model.Text{Text: " - Universal markup converter"},
	)
	want.NewItem(
		model.Text{
			Text: "mandolyte/mdtopdf",
			HREF: "https://github.com/mandolyte/mdtopdf",
		},
		model.Text{Text: " - Markdown to PDF"},
	)
	want.NewItem(
		model.Text{
			Text: "ajstarks/deck",
			HREF: "https://github.com/ajstarks/deck",
		},
		model.Text{Text: " - Slide Decks"},
	)
	want.NewItem(
		model.Text{
			Text: "johnfercher/maroto",
			HREF: "https://github.com/johnfercher/maroto",
		},
		model.Text{Text: " - A maroto way to create PDFs."},
	)

	z := html.NewTokenizer(strings.NewReader(input))
	z.Next()
	have := parseEntries(z, z.Token())

	areEqual, err := AreStructsEqual(have, want)
	if err != nil {
		log.Fatal(err)
	}

	if !areEqual {
		t.Errorf("parseList: results do not match")
	}
}
