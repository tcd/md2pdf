package render

import (
	"io"
	"log"

	gofpdf "github.com/tcd/gofpdf-1"
	"github.com/tcd/md2pdf/internal/model"
	"github.com/tcd/md2pdf/internal/renderer"
)

// Elements slice of Renderer that can be rendered to a PDF.
type Elements struct {
	Elements []renderer.Renderer `json:"elements"`
}

// Add one or more elements.
func (e *Elements) Add(elements ...renderer.Renderer) {
	e.Elements = append(e.Elements, elements...)
}

// RenderToFile writes elements to a file at the given path.
func (e Elements) RenderToFile(path string) error {
	pdf, err := gofpdf.New(
		gofpdf.PdfOptionPageSize(gofpdf.PageSizeLetter.W, gofpdf.PageSizeLetter.H), // W: 612, H: 792
		gofpdf.PdfOptionUnit(gofpdf.Unit_MM),
		gofpdf.PdfOptionMargin(20, 20, 20, 20),
	)
	if err != nil {
		log.Println(err)
	}

	err = pdf.AddTTFFont("helvetica", "/Users/clay/go/src/github.com/tcd/md2pdf/static/fonts/fonts/helvetica.ttf")
	if err != nil {
		log.Println(err)
	}

	err = pdf.AddTTFFont("courier", "/Users/clay/go/src/github.com/tcd/md2pdf/static/fonts/fonts/helvetica.ttf")
	if err != nil {
		log.Println(err)
	}

	err = pdf.SetFont("helvetica", "", 12)
	if err != nil {
		log.Println(err)
	}

	pdf.AddPage()
	for _, r := range e.Elements {
		r.Render(pdf)
	}
	return pdf.WritePdf(path)
}

// RenderToWriter writes PDF output to an io.Writer.
func (e Elements) RenderToWriter(w io.Writer) error {
	pdf, err := gofpdf.New(
		gofpdf.PdfOptionPageSize(gofpdf.PageSizeLetter.W, gofpdf.PageSizeLetter.H),
		gofpdf.PdfOptionUnit(gofpdf.Unit_MM),
		gofpdf.PdfOptionMargin(20, 20, 20, 20),
	)
	if err != nil {
		log.Println(err)
	}
	if err := pdf.AddTTFFont("helvetica", "/Users/clay/go/src/github.com/tcd/md2pdf/static/fonts/fonts/helvetica.ttf"); err != nil {
		log.Println(err)
	}
	if err := pdf.AddTTFFont("courier", "/Users/clay/go/src/github.com/tcd/md2pdf/static/fonts/fonts/helvetica.ttf"); err != nil {
		log.Println(err)
	}
	if err := pdf.SetFont("helvetica", "", 12); err != nil {
		log.Println(err)
	}

	pdf.AddPage()
	for _, r := range e.Elements {
		r.Render(pdf)
	}
	return pdf.Write(w)
}

// Blockquote implements the Renderer interface.
type Blockquote struct {
	Type    string
	Content model.Contents
}

// Render a blockquote.
func (b Blockquote) Render(pdf *gofpdf.Fpdf) {
	blockquote(pdf, b.Content)
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
		codeBlock(pdf, cb.Content)
	} else {
		highlightedCodeblock(pdf, cb.Content, cb.Class)
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
		h1(pdf, h.Content)
	case "h2":
		h2(pdf, h.Content)
	case "h3":
		h3(pdf, h.Content)
	case "h4":
		h4(pdf, h.Content)
	case "h5":
		h5(pdf, h.Content)
	case "h6":
		h6(pdf, h.Content)
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
func (h HR) Render(pdf *gofpdf.Fpdf) {
	hr(pdf)
}

// Image implements the Renderer interface.
type Image struct {
	Type string
	Src  string
	Link string
}

// Render an image.
func (img Image) Render(pdf *gofpdf.Fpdf) {
	// render.Image(pdf, img.Src, img.Link)
	err := pdf.WriteText(12, "image")
	if err != nil {
		log.Println(err)
	}
}

// List implements the Renderer interface.
type List struct {
	Type    string
	Content model.ListContent
}

// Render a list.
func (ls List) Render(pdf *gofpdf.Fpdf) {
	list(pdf, ls.Content)
}

// Paragraph implements the Renderer interface.
type Paragraph struct {
	Type    string
	Content model.Contents
}

// Render a paragraph.
func (p Paragraph) Render(pdf *gofpdf.Fpdf) {
	FullP(pdf, p.Content)
}

// Table implements the Renderer interface.
type Table struct {
	Type    string
	Content model.TableContent
}

// Render a table.
func (t Table) Render(pdf *gofpdf.Fpdf) {
	// render.Table(pdf, t.Content)
	err := pdf.WriteText(12, "table")
	if err != nil {
		log.Println(err)
	}
}
