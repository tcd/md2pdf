package md2pdf

import (
	"fmt"
	"io/ioutil"

	bf "gopkg.in/russross/blackfriday.v2"
)

// Md2PDF converts a markdown file to a PDF file.
func Md2PDF(inPath, outPath string) error {
	htmlString, err := Md2HTML(inPath)
	if err != nil {
		return err
	}

	err = Parse(htmlString, outPath)
	if err != nil {
		return err
	}

	return nil
}

// Md2HTML converts a markdown file to an html string.
func Md2HTML(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	output := bf.Run(
		bytes,
		// bf.WithExtensions(bf.CommonExtensions|bf.Tables|bf.Strikethrough),
		bf.WithExtensions(bf.CommonExtensions),
		bf.WithRenderer(bf.NewHTMLRenderer(
			bf.HTMLRendererParameters{
				Flags: bf.CommonHTMLFlags | bf.SkipHTML,
			})),
	)
	return string(output), nil
}

// Md2HTMLFile parses a markdown file with blackfriday and writes the output to a file.
func Md2HTMLFile(inPath, outPath string) {
	bytes, err := ioutil.ReadFile(inPath)
	if err != nil {
		fmt.Println(err)
	}

	output := bf.Run(
		bytes,
		// bf.WithExtensions(bf.CommonExtensions|bf.Tables|bf.Strikethrough),
		bf.WithExtensions(bf.CommonExtensions),
		bf.WithRenderer(bf.NewHTMLRenderer(
			bf.HTMLRendererParameters{
				Flags: bf.CommonHTMLFlags | bf.SkipHTML,
			})),
	)

	err = ioutil.WriteFile(outPath, output, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
