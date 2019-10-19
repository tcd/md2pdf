package md2pdf

import (
	"bytes"

	"github.com/tcd/md2pdf/internal/parse"
	"github.com/tcd/md2pdf/internal/renderer"
	rdr "github.com/tcd/md2pdf/internal/renderers/github"
)

func md2pdf(mdBytes []byte) ([]byte, error) {
	htmlBytes := mdBytes2htmlbytes(mdBytes)
	rbs := parse.Parse(htmlBytes)
	var buf bytes.Buffer
	var r rdr.Renderer
	err := renderer.RenderToWriter(r, rbs, &buf)

	if err != nil {
		return buf.Bytes(), err
	}
	return buf.Bytes(), nil
}
