package renderer

import "github.com/tcd/md2pdf/internal/model"

type RenderableType int

const (
	UnknownType RenderableType = iota
	BlockquoteType
	CodeblockType
	HeaderType
	HRType
	ImageType
	ListType
	ParagraphType
	TableType
)

func (rt RenderableType) String() string {
	var s string
	switch rt {
	case BlockquoteType:
		s = "blockquote"
	case CodeblockType:
		s = "codeblock"
	case HeaderType:
		s = "header"
	case HRType:
		s = "hr"
	case ImageType:
		s = "image"
	case ListType:
		s = "list"
	case ParagraphType:
		s = "paragraph"
	case TableType:
		s = "table"
	default:
		s = ""
	}
	return s
}

type Renderable interface {
	Type() RenderableType
}

type Renderables struct {
	Renderables []Renderable `json:"renderables"`
}

// Add one or more elements.
func (rbs *Renderables) Add(r ...Renderable) {
	rbs.Renderables = append(rbs.Renderables, r...)
}

type Blockquote struct{ Content model.Contents }

func (b Blockquote) Type() RenderableType { return BlockquoteType }

type Codeblock struct {
	Class   string
	Content model.Contents
}

func (c Codeblock) Type() RenderableType { return CodeblockType }

type Header struct {
	Level   string
	Content model.Contents
}

func (h Header) Type() RenderableType { return HeaderType }

type HR struct{}

func (h HR) Type() RenderableType { return HRType }

type Image struct {
	Src  string
	Link string
}

func (i Image) Type() RenderableType { return ImageType }

type List struct{ Content model.ListContent }

func (l List) Type() RenderableType { return ListType }

type Paragraph struct{ Content model.Contents }

func (p Paragraph) Type() RenderableType { return ParagraphType }

type Table struct{ Content model.TableContent }

func (t Table) Type() RenderableType { return TableType }
