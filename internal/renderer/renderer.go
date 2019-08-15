package renderer

import (
	"bytes"

	bf "github.com/russross/blackfriday"
)

// PDF is a type that implements the Renderer interface for PDF output.
//
// Do not create this directly, instead use the PDFRenderer function.
type PDF struct {
}

// PDFRenderer creates and configures a PDF object, which
// satisfies the Renderer interface.
//
// flags is a set of PDF_* options ORed together (currently no such options
// are defined).
func PDFRenderer(flags int) bf.Renderer {
	return &PDF{}
}

// GetFlags parses options for a PDF renderer
func (options *PDF) GetFlags() int {
	return 0
}

// ============================================================================
// block-level callbacks
// ============================================================================

func (options *PDF) BlockCode(out *bytes.Buffer, text []byte, infoString string) {}

func (options *PDF) BlockQuote(out *bytes.Buffer, text []byte) {}

// BlockHtml doesn't do anything, we're not worrying about raw HTML
func (options *PDF) BlockHtml(out *bytes.Buffer, text []byte) {}

func (options *PDF) Header(out *bytes.Buffer, text func() bool, level int, id string)      {}
func (options *PDF) HRule(out *bytes.Buffer)                                               {}
func (options *PDF) List(out *bytes.Buffer, text func() bool, flags int)                   {}
func (options *PDF) ListItem(out *bytes.Buffer, text []byte, flags int)                    {}
func (options *PDF) Paragraph(out *bytes.Buffer, text func() bool)                         {}
func (options *PDF) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {}
func (options *PDF) TableRow(out *bytes.Buffer, text []byte)                               {}
func (options *PDF) TableHeaderCell(out *bytes.Buffer, text []byte, flags int)             {}
func (options *PDF) TableCell(out *bytes.Buffer, text []byte, flags int)                   {}
func (options *PDF) Footnotes(out *bytes.Buffer, text func() bool)                         {}
func (options *PDF) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int)          {}
func (options *PDF) TitleBlock(out *bytes.Buffer, text []byte)                             {}

// ============================================================================
// Span-level callbacks
// ============================================================================

func (options *PDF) AutoLink(out *bytes.Buffer, link []byte, kind int)                 {}
func (options *PDF) CodeSpan(out *bytes.Buffer, text []byte)                           {}
func (options *PDF) DoubleEmphasis(out *bytes.Buffer, text []byte)                     {}
func (options *PDF) Emphasis(out *bytes.Buffer, text []byte)                           {}
func (options *PDF) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte)    {}
func (options *PDF) LineBreak(out *bytes.Buffer)                                       {}
func (options *PDF) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {}
func (options *PDF) RawHtmlTag(out *bytes.Buffer, tag []byte)                          {}
func (options *PDF) TripleEmphasis(out *bytes.Buffer, text []byte)                     {}
func (options *PDF) StrikeThrough(out *bytes.Buffer, text []byte)                      {}
func (options *PDF) FootnoteRef(out *bytes.Buffer, ref []byte, id int)                 {}

// ============================================================================
// Low-level callbacks
// ============================================================================

func (options *PDF) Entity(out *bytes.Buffer, entity []byte)   {}
func (options *PDF) NormalText(out *bytes.Buffer, text []byte) {}

// ============================================================================
// Header and footer
// ============================================================================

func (options *PDF) DocumentHeader(out *bytes.Buffer) {}
func (options *PDF) DocumentFooter(out *bytes.Buffer) {}
