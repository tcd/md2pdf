package renderable

import (
	"log"

	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/model"
	"github.com/tcd/md2pdf/internal/render"
)

// Renderable implementers can return data that can be used to draw them to a gofpdf.Fpdf.
type Renderable interface {
	Render(*gofpdf.Fpdf)
}

// Blockquote implements the Renderable interface.
type Blockquote struct {
	Content model.Contents
}

// Render a blockquote.
func (b Blockquote) Render(pdf *gofpdf.Fpdf) {
	render.Blockquote(pdf, b.Content)
}

// Codeblock implements the Renderable interface.
type Codeblock struct {
	Class   string
	Content model.Contents
}

// Render a codeblock.
func (cb Codeblock) Render(pdf *gofpdf.Fpdf) {
	if cb.Class == "language-no-highlight" || cb.Class == "" {
		render.CodeBlock(pdf, cb.Content)
	} else {
		render.CodeBlock(pdf, cb.Content)
	}
}

// Header implements the Renderable interface.
type Header struct {
	Level   string
	Content model.Contents
}

// Render a header.
func (h Header) Render(pdf *gofpdf.Fpdf) {
	switch h.Level {
	case "h1":
		render.H1(pdf, h.Content)
	case "h2":
		render.H2(pdf, h.Content)
	case "h3":
		render.H3(pdf, h.Content)
	case "h4":
		render.H4(pdf, h.Content)
	case "h5":
		render.H5(pdf, h.Content)
	case "h6":
		render.H6(pdf, h.Content)
	default:
		log.Printf("Error rendering header with content: %v\n", h.Content)
		return
	}
}

// HR implements the Renderable interface.
type HR struct{}

// Render a horizontal rule.
func (hr HR) Render(pdf *gofpdf.Fpdf) {
	render.HR(pdf)
}

// Image implements the Renderable interface.
type Image struct {
	Src  string
	Link string
}

// Render an image.
func (img Image) Render(pdf *gofpdf.Fpdf) {
	render.Image(pdf, img.Src)
}

// List implements the Renderable interface.
type List struct {
	Content model.ListContent
}

// Render a list.
func (ls List) Render(pdf *gofpdf.Fpdf) {
	render.AnyList(pdf, ls.Content)
}

// Paragraph implements the Renderable interface.
type Paragraph struct {
	Content model.Contents
}

// Render a paragraph.
func (p Paragraph) Render(pdf *gofpdf.Fpdf) {
	render.FullP(pdf, p.Content)
}

// Table implements the Renderable interface.
type Table struct {
	Content model.TableContent
}

// Render a table.
func (t Table) Render(pdf *gofpdf.Fpdf) {
	render.Table(pdf, t.Content)
}
