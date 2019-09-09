package renderer

import (
	"io"

	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/render"
)

// Elements slice of Renderer that can be rendered to a PDF.
type Elements struct {
	Elements []Renderer `json:"elements"`
}

// Add one or more elements.
func (e *Elements) Add(elements ...Renderer) {
	e.Elements = append(e.Elements, elements...)
}

// RenderToFile writes elements to a file at the given path.
func (e Elements) RenderToFile(path string) error {
	pdf := gofpdf.New("P", "mm", "Letter", "")
	render.Setup(pdf)
	for _, r := range e.Elements {
		r.Render(pdf)
	}
	return pdf.OutputFileAndClose(path)
}

// RenderToWriter writes PDF output to an io.Writer.
func (e Elements) RenderToWriter(w io.Writer) error {
	pdf := gofpdf.New("P", "mm", "Letter", "")
	render.Setup(pdf)
	for _, r := range e.Elements {
		r.Render(pdf)
	}
	return pdf.Output(w)
}
