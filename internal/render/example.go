package render

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"github.com/tcd/md2pdf/internal/content"
)

// GitHubPDF attempts to recreate this page (https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet) as a pdf.
func GitHubPDF(outPath string) {
	pdf := gofpdf.New("P", "mm", "Letter", "")
	Setup(pdf)
	SetMetaData(pdf, "Clay Dunston", "md2pdf github test")

	H1(pdf, "GitHub Markdown PDF demo")
	FullP(pdf, PContent1)
	FullP(pdf, PContent2)
	H5(pdf, "Table of Contents")
	// AnyList(pdf, content.TocContent)
	H2(pdf, "Headers")
	CodeBlock(pdf, content.HeaderContent())
	H1(pdf, "H1")
	H2(pdf, "H2")
	H3(pdf, "H3")
	H4(pdf, "H4")
	H5(pdf, "H5")
	H6(pdf, "H6")
	BasicP(pdf, "Alternatively, for H1 and H2, an underline-ish style:")
	H1(pdf, "Alt-H1")
	H2(pdf, "Alt-H2")
	H2(pdf, "Emphasis")
	CodeBlock(pdf, content.EmContent())
	FullP(pdf, PContentEmphasis)
	FullP(pdf, PContentStrong)
	FullP(pdf, PContentCombined)
	FullP(pdf, PContentStrike)
	H2(pdf, "Lists")
	CodeBlock(pdf, content.ListContent())
	H2(pdf, "Links")
	CodeBlock(pdf, content.LinksContent())
	FullP(pdf, Contents{
		Content: []Text{Text{
			Content: "I'm an inline-style link",
			HREF:    "https://www.google.com/",
		}},
	})
	FullP(pdf, Contents{
		Content: []Text{Text{
			Content: "Unfortunately, we can't put titles on pdf links :(",
			HREF:    "https://www.google.com/",
		}},
	})
	FullP(pdf, Contents{
		Content: []Text{Text{
			Content: "I'm a reference-style link",
			HREF:    "https://www.mozilla.org",
		}},
	})
	FullP(pdf, Contents{
		Content: []Text{Text{
			Content: "We also can't link to relative files :'(",
			HREF:    "",
		}},
	})
	FullP(pdf, Contents{
		Content: []Text{Text{
			Content: "You can use numbers for reference-style link definitions",
			HREF:    "http://slashdot.org",
		}},
	})
	FullP(pdf, Contents{
		Content: []Text{
			Text{Content: "Or leave it empty and use the "},
			Text{Content: "link text itself", HREF: "http://www.reddit.com"},
			Text{Content: "."},
		},
	})
	FullP(pdf, Contents{
		Content: []Text{
			Text{Content: "URLs and URLs in angle brackets will automatically get turned into links. "},
			Text{Content: "http://www.example.com", HREF: "http://www.example.com"},
			Text{Content: " or "},
			Text{Content: "http://www.example.com", HREF: "http://www.example.com"},
			Text{Content: " and sometimes example.com (but not on Github, for example)."},
		},
	})
	FullP(pdf, Contents{
		Content: []Text{Text{
			Content: "Some text to show that the reference links can follow later.",
		}},
	})
	H2(pdf, "Images")
	CodeBlock(pdf, content.ImagesContent())
	BasicP(pdf, "Here's our logo (hover to see the title text):")
	BasicP(pdf, "Inline-style: ")
	Image(pdf, "https://github.com/adam-p/markdown-here/raw/master/src/common/images/icon48.png")
	BasicP(pdf, "Reference-style:")
	Image(pdf, "/Users/clay/go/src/github.com/tcd/md2pdf/static/images/icon48.png")
	H2(pdf, "Code and Syntax Highlighting")
	FullP(pdf, PContentCodeAndSyntax)
	CodeBlock(pdf, "Inline `code` has `back-ticks around` it.")
	FullP(pdf, inlineCodeContent1)
	FullP(pdf, inlineCodeContent2)
	CodeBlock(pdf, content.InlineCodeContent())
	CodeBlock(pdf, content.SyntaxHighContent())
	H2(pdf, "Tables")
	BasicP(pdf, "Tables aren't part of the core Markdown spec, but they are part of GFM and Markdown Here supports them. They are an easy way of adding tables to your email -- a task that would otherwise require copy-pasting from another application.")
	CodeBlock(pdf, content.TablesContent())
	BasicP(pdf, "Colons can be used to align columns.")
	Table(pdf, TableContent1())
	BasicP(pdf, "There must be at least 3 dashes separating each header cell. The outer pipes (|) are optional, and you don't need to make the raw Markdown line up prettily. You can also use inline Markdown.")
	Table(pdf, TableContent2())
	H2(pdf, "Blockquotes")
	CodeBlock(pdf, content.BlockquoteContent())
	BasicBlockquote(pdf, "Blockquotes are very handy in email to emulate reply text. This line is part of the same quote.")
	BasicP(pdf, "Quote break.")
	BasicBlockquote(pdf, "This is a very long line that will still be quoted properly when it wraps. Oh boy let's keep writing to make sure this is long enough to actually wrap for everyone. Oh, you can put Markdown into a blockquote.")
	H2(pdf, "Inline HTML")
	BasicP(pdf, "You can also use raw HTML in your Markdown, and it'll mostly work pretty well.")
	CodeBlock(pdf, content.InlineHTMLContent())
	H2(pdf, "Horizontal Rule")
	BasicP(pdf, "Three or more...")
	HR(pdf)
	BasicP(pdf, "Hyphens")
	HR(pdf)
	BasicP(pdf, "Asterisks")
	HR(pdf)
	BasicP(pdf, "Underscores")
	CodeBlock(pdf, content.HRContent())
	H2(pdf, "Line Breaks")
	BasicP(pdf, `My basic recommendation for learning how line breaks work is to experiment and discover -- hit <Enter> once (i.e., insert one newline), then hit it twice (i.e., insert two newlines), see what happens. You'll soon learn to get what you want. "Markdown Toggle" is your friend.`)
	BasicP(pdf, "Here are some things to try out:")
	CodeBlock(pdf, content.LineBreakContents())
	H2(pdf, "YouTube Videos")
	BasicP(pdf, "They can't be added directly but you can add an image with a link to the video like this:")
	CodeBlock(pdf, content.YouTubeContent1())
	BasicP(pdf, "Or, in pure Markdown, but losing the image sizing and border:")
	CodeBlock(pdf, content.YouTubeContent2())
	BasicP(pdf, "Referencing a bug by #bugID in your git commit links it to the slip. For example #1.")
	FullP(pdf, Contents{
		Content: []Text{
			Text{
				Content: "Original Markdown Cheatsheet",
				HREF:    "https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet",
			},
		},
	})

	err := pdf.OutputFileAndClose(outPath)
	if err != nil {
		fmt.Println(err)
	}
}
