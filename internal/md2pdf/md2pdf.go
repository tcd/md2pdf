package md2pdf

import (
	"io/ioutil"
	"os"

	"github.com/tcd/md2pdf/internal/parse"
	bf "gopkg.in/russross/blackfriday.v2"
)

// Md2PDF converts a markdown file to a PDF file.
func Md2PDF(inPath, outPath string) error {
	htmlString, err := Md2HTML(inPath)
	if err != nil {
		return err
	}

	err = parse.HTML(htmlString, outPath)
	if err != nil {
		return err
	}

	return nil
}

// Defines the extensions that are used
var extensions = bf.NoIntraEmphasis | bf.Tables | bf.FencedCode | bf.Autolink | bf.Strikethrough | bf.SpaceHeadings | bf.BackslashLineBreak | bf.HeadingIDs

// Defines the HTML rendering flags that are used
var flags = bf.UseXHTML | bf.SkipHTML

// Md2HTML converts a markdown file to an html string.
func Md2HTML(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	output := bf.Run(
		bytes,
		bf.WithExtensions(extensions),
		bf.WithRenderer(
			bf.NewHTMLRenderer(bf.HTMLRendererParameters{Flags: flags}),
		),
	)

	return string(output), nil
}

// Md2HTMLFile parses a markdown file with blackfriday and writes the output to a file.
func Md2HTMLFile(inPath, outPath string) error {
	bytes, err := ioutil.ReadFile(inPath)
	if err != nil {
		return err
	}

	output := bf.Run(
		bytes,
		bf.WithExtensions(extensions),
		bf.WithRenderer(
			bf.NewHTMLRenderer(bf.HTMLRendererParameters{Flags: flags}),
		),
	)

	err = ioutil.WriteFile(outPath, []byte(output), os.FileMode(0644))
	if err != nil {
		return err
	}
	return nil
}
