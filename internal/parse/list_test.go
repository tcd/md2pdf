package parse

import (
	"encoding/json"
	"log"
	"strings"
	"testing"

	"github.com/tcd/md2pdf/internal/model"
	"golang.org/x/net/html"
)

func TestParseList(t *testing.T) {
	inputHTML := `
	<ul>
		<li><a href="https://github.com/jgm/pandoc">jgm/pandoc</a> - Universal markup converter</li>
		<li><a href="https://github.com/mandolyte/mdtopdf">mandolyte/mdtopdf</a> - Markdown to PDF</li>
		<li><a href="https://github.com/ajstarks/deck">ajstarks/deck</a> - Slide Decks</li>
		<li><a href="https://github.com/johnfercher/maroto">johnfercher/maroto</a> - A maroto way to create PDFs.</li>
	</ul>`
	outputElement := model.ListContent{
		Ordered: false,
		Items: []model.ListItem{
			model.ListItem{
				Contents: model.Contents{
					Content: []model.Text{
						model.Text{
							Text: "jgm/pandoc",
							HREF: "https://github.com/jgm/pandoc",
						},
						model.Text{Text: " - Universal markup converter"},
					},
				},
			},
			model.ListItem{
				Contents: model.Contents{
					Content: []model.Text{
						model.Text{
							Text: "mandolyte/mdtopdf",
							HREF: "https://github.com/mandolyte/mdtopdf",
						},
						model.Text{Text: " - Markdown to PDF"},
					},
				},
			},
			model.ListItem{
				Contents: model.Contents{
					Content: []model.Text{
						model.Text{
							Text: "ajstarks/deck",
							HREF: "https://github.com/ajstarks/deck",
						},
						model.Text{Text: " - Slide Decks"},
					},
				},
			},
			model.ListItem{
				Contents: model.Contents{
					Content: []model.Text{
						model.Text{
							Text: "johnfercher/maroto",
							HREF: "https://github.com/johnfercher/maroto",
						},
						model.Text{Text: " - A maroto way to create PDFs."},
					},
				},
			},
		},
	}

	doc := strings.NewReader(inputHTML)
	z := html.NewTokenizer(doc)
	z.Next()
	output := parseEntries(z, z.Token())

	haveBytes, err := json.Marshal(output)
	if err != nil {
		log.Fatal(err)
	}
	wantBytes, err := json.Marshal(outputElement)
	if err != nil {
		log.Fatal(err)
	}
	areEqual, err := AreEqualJSON(haveBytes, wantBytes)
	if err != nil {
		log.Fatal(err)
	}

	if !areEqual {
		t.Errorf("parseList: results do not match")
	}
}
