package md2pdf

import (
	"io/ioutil"

	bf "gopkg.in/russross/blackfriday.v2"
)

// Given a byte slice of markdown content, returns byte slice of html content.
func mdBytes2htmlbytes(input []byte) []byte {
	var extensions = bf.NoIntraEmphasis | bf.Tables | bf.FencedCode | bf.Autolink | bf.Strikethrough | bf.SpaceHeadings | bf.BackslashLineBreak | bf.HeadingIDs
	var flags = bf.UseXHTML | bf.SkipHTML
	return bf.Run(
		input,
		bf.WithExtensions(extensions),
		bf.WithRenderer(bf.NewHTMLRenderer(bf.HTMLRendererParameters{Flags: flags})),
	)
}

func mdFile2htmlBytes(path string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}, err
	}
	output := mdBytes2htmlbytes(bytes)
	return output, nil
}
