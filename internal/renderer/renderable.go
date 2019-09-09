package renderer

import (
	"log"

	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/model"
	"github.com/tcd/md2pdf/internal/render"
)

// Renderer implementers can draw content to a gofpdf.Fpdf.
type Renderer interface {
	Render(*gofpdf.Fpdf)
}

// Blockquote implements the Renderer interface.
type Blockquote struct {
	Type    string
	Content model.Contents
}

// Render a blockquote.
func (b Blockquote) Render(pdf *gofpdf.Fpdf) {
	render.Blockquote(pdf, b.Content)
}

// Codeblock implements the Renderer interface.
type Codeblock struct {
	Type    string
	Class   string
	Content model.Contents
}

// Render a codeblock.
func (cb Codeblock) Render(pdf *gofpdf.Fpdf) {
	if cb.Class == "language-no-highlight" || cb.Class == "" {
		render.CodeBlock(pdf, cb.Content)
	} else {
		render.HighlightedCodeblock(pdf, cb.Content, cb.Class)
	}
}

// Header implements the Renderer interface.
type Header struct {
	Type    string
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

// HR implements the Renderer interface.
type HR struct {
	Type string
}

// Render a horizontal rule.
func (hr HR) Render(pdf *gofpdf.Fpdf) {
	render.HR(pdf)
}

// Image implements the Renderer interface.
type Image struct {
	Type string
	Src  string
	Link string
}

// Render an image.
func (img Image) Render(pdf *gofpdf.Fpdf) {
	render.Image(pdf, img.Src, img.Link)
}

// List implements the Renderer interface.
type List struct {
	Type    string
	Content model.ListContent
}

// Render a list.
func (ls List) Render(pdf *gofpdf.Fpdf) {
	render.List(pdf, ls.Content)
}

// Paragraph implements the Renderer interface.
type Paragraph struct {
	Type    string
	Content model.Contents
}

// Render a paragraph.
func (p Paragraph) Render(pdf *gofpdf.Fpdf) {
	render.FullP(pdf, p.Content)
}

// Table implements the Renderer interface.
type Table struct {
	Type    string
	Content model.TableContent
}

// Render a table.
func (t Table) Render(pdf *gofpdf.Fpdf) {
	render.Table(pdf, t.Content)
}
