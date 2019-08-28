package renderable

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/render"
)

// Elements slice of Renderables that can be rendered to a PDF.
type Elements struct {
	Elements []Renderable `json:"elements"`
}

// Add an Element.
func (e *Elements) Add(element Renderable) {
	e.Elements = append(e.Elements, element)
}

// // Add one or more elements.
// func (e *Elements) Add(elements ...Renderable) {
// 	e.Elements = append(e.Elements, elements...)
// }

// Render a PDF from Elements.
func (e Elements) Render(outPath string) (err error) {
	pdf := gofpdf.New("P", "mm", "Letter", "")
	render.Setup(pdf)
	for _, r := range e.Elements {
		r.Render(pdf)
	}
	err = pdf.OutputFileAndClose(outPath)
	return
}
