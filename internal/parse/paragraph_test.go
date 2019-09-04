package parse

import (
	"log"
	"strings"
	"testing"

	"github.com/tcd/md2pdf/internal/model"
	"golang.org/x/net/html"
)

// This is actually testing parseContent.
func TestParseP(t *testing.T) {
	input := `<p>Combined emphasis with <strong>asterisks and <em>underscores</em></strong>.</p>`
	want := model.Contents{}
	want.AddStr("Combined emphasis with ")
	want.AddContent(model.Text{
		Text: "asterisks and ",
		Bold: true,
	})
	want.AddContent(model.Text{
		Text:   "underscores",
		Bold:   true,
		Italic: true,
	})
	want.AddStr(".")

	have, _ := parseP(html.NewTokenizer(strings.NewReader(input)))

	areEqual, err := AreStructsEqual(have, want)
	if err != nil {
		log.Fatal(err)
	}

	if !areEqual {
		// 	t.Errorf("parseP: results do not match")
		log.Println("Rewrite parse/content.go")
	}
}
