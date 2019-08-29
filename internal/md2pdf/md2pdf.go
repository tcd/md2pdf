package md2pdf

import (
	"bytes"

	"github.com/tcd/md2pdf/internal/parse"
)

func md2pdf(mdBytes []byte) ([]byte, error) {
	htmlBytes := mdBytes2htmlbytes(mdBytes)
	elements := parse.Parse(htmlBytes)
	var buf bytes.Buffer
	err := elements.RenderToWriter(&buf)
	if err != nil {
		return buf.Bytes(), err
	}
	return buf.Bytes(), nil
}
