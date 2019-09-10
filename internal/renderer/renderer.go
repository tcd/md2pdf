package renderer

import gofpdf "github.com/tcd/gofpdf-1"

// Renderer implementers can render content to a gofpdf.Fpdf.
type Renderer interface {
	Render(*gofpdf.Fpdf)
}

// RenderList slice of Renderer that can be rendered to a PDF.
type RenderList struct {
	Elements []Renderer `json:"elements"`
}

// Add one or more elements.
func (rl *RenderList) Add(elements ...Renderer) {
	rl.Elements = append(rl.Elements, elements...)
}
