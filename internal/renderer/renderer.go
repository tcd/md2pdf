package renderer

import (
	"io"

	"github.com/jung-kurt/gofpdf"
)

// Renderer implementers can draw content to a gofpdf.Fpdf.
type Renderer interface {
	Setup(*gofpdf.Fpdf)
	RenderHeader(*gofpdf.Fpdf, Header)
	RenderBlockquote(*gofpdf.Fpdf, Blockquote)
	RenderCodeblock(*gofpdf.Fpdf, Codeblock)
	RenderHR(*gofpdf.Fpdf)
	RenderImage(*gofpdf.Fpdf, Image)
	RenderList(*gofpdf.Fpdf, List)
	RenderParagraph(*gofpdf.Fpdf, Paragraph)
	RenderTable(*gofpdf.Fpdf, Table)
}

// RenderToFile outputs PDF content to a file.
func RenderToFile(r Renderer, rbs Renderables, path string) error {
	f := gofpdf.New("P", "mm", "Letter", "")
	r.Setup(f)
	for _, rb := range rbs.Renderables {
		switch rb.Type() {
		case BlockquoteType:
			r.RenderBlockquote(f, rb.(Blockquote))
		case CodeblockType:
			r.RenderCodeblock(f, rb.(Codeblock))
		case HeaderType:
			r.RenderHeader(f, rb.(Header))
		case HRType:
			r.RenderHR(f)
		case ImageType:
			r.RenderImage(f, rb.(Image))
		case ListType:
			r.RenderList(f, rb.(List))
		case ParagraphType:
			r.RenderParagraph(f, rb.(Paragraph))
		case TableType:
			r.RenderTable(f, rb.(Table))
		}
	}
	return f.OutputFileAndClose(path)
}

// RenderToWriter outputs PDF content to to an io.Writer.
func RenderToWriter(r Renderer, rbs Renderables, w io.Writer) error {
	f := gofpdf.New("P", "mm", "Letter", "")
	r.Setup(f)
	for _, rb := range rbs.Renderables {
		switch rb.Type() {
		case BlockquoteType:
			r.RenderBlockquote(f, rb.(Blockquote))
		case CodeblockType:
			r.RenderCodeblock(f, rb.(Codeblock))
		case HeaderType:
			r.RenderHeader(f, rb.(Header))
		case HRType:
			r.RenderHR(f)
		case ImageType:
			r.RenderImage(f, rb.(Image))
		case ListType:
			r.RenderList(f, rb.(List))
		case ParagraphType:
			r.RenderParagraph(f, rb.(Paragraph))
		case TableType:
			r.RenderTable(f, rb.(Table))
		}
	}
	return f.Output(w)
}
