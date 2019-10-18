// Package github renders text in the style of GitHub using gofpdf.
package github

import (
	"log"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/renderer"
)

// Setup sets default GitHub-Flavored styles to a gofpdf.Fpdf.
func Setup(f *gofpdf.Fpdf) {
	f.SetMargins(20, 20, 20)     // left, top, right margins
	f.SetAutoPageBreak(true, 20) // bottom margin
	f.AddPage()
	f.SetFillColor(255, 255, 255)
	f.SetTextColor(36, 41, 46)
}

// SetMetaData for a pdf.
func SetMetaData(f *gofpdf.Fpdf, author, title string) {
	f.SetTitle(title, true)
	f.SetAuthor(author, true)
	f.SetCreator(author, true)
	f.SetCreationDate(time.Now())
}

type Renderer int

func (r Renderer) Setup(f *gofpdf.Fpdf) {
	f.SetMargins(20, 20, 20)     // left, top, right margins
	f.SetAutoPageBreak(true, 20) // bottom margin
	f.AddPage()
	f.SetFillColor(255, 255, 255)
	f.SetTextColor(36, 41, 46)
}

func (r Renderer) RenderBlockquote(f *gofpdf.Fpdf, b renderer.Blockquote) {
	Blockquote(f, b.Content)
}

func (r Renderer) RenderHeader(f *gofpdf.Fpdf, h renderer.Header) {
	switch h.Level {
	case "h1":
		H1(f, h.Content)
	case "h2":
		H2(f, h.Content)
	case "h3":
		H3(f, h.Content)
	case "h4":
		H4(f, h.Content)
	case "h5":
		H5(f, h.Content)
	case "h6":
		H6(f, h.Content)
	default:
		log.Printf("Error rendering header with content: %v\n", h.Content)
		return
	}
}

func (r Renderer) RenderCodeblock(f *gofpdf.Fpdf, c renderer.Codeblock) {
	if c.Class == "language-no-highlight" || c.Class == "" {
		CodeBlock(f, c.Content)
	} else {
		HighlightedCodeblock(f, c.Content, c.Class)
	}
}

func (r Renderer) RenderHR(f *gofpdf.Fpdf) {
	HR(f)
}

func (r Renderer) RenderImage(f *gofpdf.Fpdf, i renderer.Image) {
	Image(f, i.Src, i.Link)
}

func (r Renderer) RenderList(f *gofpdf.Fpdf, ls renderer.List) {
	List(f, ls.Content)
}

func (r Renderer) RenderParagraph(f *gofpdf.Fpdf, p renderer.Paragraph) {
	FullP(f, p.Content)
}

func (r Renderer) RenderTable(f *gofpdf.Fpdf, t renderer.Table) {
	Table(f, t.Content)
}
