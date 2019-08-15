package ghfm

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

// GitHubPDF attempts to recreate this page (https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet) as a pdf.
func GitHubPDF(outPath string) {
	pdf := gofpdf.New(
		"P",      // orientationStr: "P" or "L" (Portrait or Landscape)
		"mm",     // unitStr: "pt", "mm", "cm", or "in" (point, millimeter, centimeter, or inches)
		"Letter", // sizeStr: "A3", "A4", "A5", "Letter", "Legal", or "Tabloid".
		"",       // fontDirStr: Empty is fine when using core fonts.
	)
	Setup(pdf)

	H1(pdf, "GitHub Markdown PDF demo")
	H2(pdf, "Headers")
	CodeBlock(pdf, cbContent())
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
	CodeBlock(pdf, emContent())
	BasicP(pdf, "Emphasis, aka italics, with asterisks or underscores.")
	BasicP(pdf, "Strong emphasis, aka bold, with asterisks or underscores.")
	BasicP(pdf, "Combined emphasis with asterisks and underscores.")
	BasicP(pdf, "Strikethrough uses two tildes. Scratch this.")
	H2(pdf, "Lists")
	CodeBlock(pdf, listContent())
	H2(pdf, "Links")
	CodeBlock(pdf, linksContent())
	link(pdf, "I'm an inline-style link", "https://www.google.com")
	H2(pdf, "Images")
	CodeBlock(pdf, imagesContent())
	BasicP(pdf, "Here's our logo (hover to see the title text):")
	image1(pdf, "https://github.com/adam-p/markdown-here/raw/master/src/common/images/icon48.png", "Logo Title Text 1", "")
	H2(pdf, "Code and Syntax Highlighting")
	BasicP(pdf, "Code blocks are part of the Markdown spec, but syntax highlighting isn't. However, many renderers -- like Github's and Markdown Here -- support syntax highlighting. Which languages are supported and how those language names should be written will vary from renderer to renderer. Markdown Here supports highlighting for dozens of languages (and not-really-languages, like diffs and HTTP headers); to see the complete list, and how to write the language names, see the highlight.js demo page.")
	CodeBlock(pdf, inlineCodeContent())
	CodeBlock(pdf, syntaxHighContent())
	BasicP(pdf, "Blocks of code are either fenced by lines with three back-ticks ```, or are indented with four spaces. I recommend only using the fenced code blocks -- they're easier and only they support syntax highlighting.")
	H2(pdf, "Tables")
	BasicP(pdf, "Tables aren't part of the core Markdown spec, but they are part of GFM and Markdown Here supports them. They are an easy way of adding tables to your email -- a task that would otherwise require copy-pasting from another application.")
	CodeBlock(pdf, tablesContent())
	BasicP(pdf, "Colons can be used to align columns.")
	BasicP(pdf, "There must be at least 3 dashes separating each header cell. The outer pipes (|) are optional, and you don't need to make the raw Markdown line up prettily. You can also use inline Markdown.")
	H2(pdf, "Blockquotes")
	CodeBlock(pdf, blockquoteContent())
	BasicP(pdf, "Blockquotes are very handy in email to emulate reply text. This line is part of the same quote.")
	BasicP(pdf, "Quote break.")
	BasicP(pdf, "This is a very long line that will still be quoted properly when it wraps. Oh boy let's keep writing to make sure this is long enough to actually wrap for everyone. Oh, you can put Markdown into a blockquote.")
	H2(pdf, "Inline HTML")
	BasicP(pdf, "You can also use raw HTML in your Markdown, and it'll mostly work pretty well.")
	CodeBlock(pdf, inlineHTMLContent())
	H2(pdf, "Horizontal Rule")
	BasicP(pdf, "Three or more...")
	HR(pdf)
	BasicP(pdf, "Hyphens")
	HR(pdf)
	BasicP(pdf, "Asterisks")
	HR(pdf)
	BasicP(pdf, "Underscores")
	CodeBlock(pdf, hrContent())
	H2(pdf, "Line Breaks")
	BasicP(pdf, `My basic recommendation for learning how line breaks work is to experiment and discover -- hit <Enter> once (i.e., insert one newline), then hit it twice (i.e., insert two newlines), see what happens. You'll soon learn to get what you want. "Markdown Toggle" is your friend.`)
	BasicP(pdf, "Here are some things to try out:")
	CodeBlock(pdf, lineBreakContents())
	H2(pdf, "YouTube Videos")
	BasicP(pdf, "They can't be added directly but you can add an image with a link to the video like this:")
	CodeBlock(pdf, youTubeContent1())
	BasicP(pdf, "Or, in pure Markdown, but losing the image sizing and border:")
	CodeBlock(pdf, youTubeContent2())
	BasicP(pdf, "Referencing a bug by #bugID in your git commit links it to the slip. For example #1.")
	link(pdf, "Original Markdown Cheatsheet", "https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet")

	err := pdf.OutputFileAndClose(outPath)
	if err != nil {
		fmt.Println(err)
	}
}
