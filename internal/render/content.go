package render

// PContent1 returns content to write be used by FullP.
var PContent1 = []Text{
	Text{Content: "This is intended as a quick reference and showcase. For more complete info, see "},
	Text{
		Content: "John Gruber's original spec",
		HREF:    "http://daringfireball.net/projects/markdown/",
	},
	Text{Content: " and the "},
	Text{
		Content: "Github-flavored Markdown info page",
		HREF:    "http://github.github.com/github-flavored-markdown/",
	},
	Text{Content: "."},
}

// PContent2 returns content to write be used by FullP.
var PContent2 = []Text{
	Text{Content: "Note that there is also a "},
	Text{
		Content: "Cheatsheet specific to Markdown Here",
		HREF:    "./Markdown-Here-Cheatsheet",
	},
	Text{Content: " if that's what you're looking for. You can also check out "},
	Text{
		Content: "more Markdown tools",
		HREF:    "./Other-Markdown-Tools",
	},
	Text{Content: "."},
}

// PContentEmphasis returns content to write be used by FullP.
var PContentEmphasis = []Text{
	Text{Content: "Emphasis, aka italics, with "},
	Text{
		Content: "asterisks",
		Italic:  true,
	},
	Text{Content: " or "},
	Text{
		Content: "underscores",
		Italic:  true,
	},
}

// PContentStrong returns content to write be used by FullP.
var PContentStrong = []Text{
	Text{Content: "Strong emphasis, aka bold, with "},
	Text{
		Content: "asterisks",
		Bold:    true,
	},
	Text{Content: " or "},
	Text{
		Content: "underscores",
		Bold:    true,
	},
}

// PContentCombined returns content to write be used by FullP.
var PContentCombined = []Text{
	Text{Content: "Combined emphasis with "},
	Text{
		Content: "asterisks and ",
		Bold:    true,
	},
	Text{
		Content: "underscores",
		Italic:  true,
		Bold:    true,
	},
}

// PContentStrike returns content to write be used by FullP.
var PContentStrike = []Text{
	Text{Content: "Strikethrough uses two tildes. "},
	Text{
		Content: "Scratch this.",
		Strike:  true,
	},
	Text{Content: "This isn't that easy to do in a PDF."},
}

// PContentCodeAndSyntax returns content to write be used by FullP.
var PContentCodeAndSyntax = []Text{
	Text{Content: "Code blocks are part of the Markdown spec, but syntax highlighting isn't. However, many renderers -- like Github's and "},
	Text{
		Content: "Markdown here --",
		Italic:  true,
	},
	Text{Content: " support syntax highlighting. Which languages are supported and how those language names should be written will vary from renderer to renderer. "},
	Text{
		Content: "Markdown Here",
		Italic:  true,
	},
	Text{Content: " supports highlighting for dozens of languages (and not-really-languages, like diffs and HTTP headers); to see the complete list, and how to write the language names, see the "},
	Text{
		Content: "highlight.js demo page",
		HREF:    "https://highlightjs.org/static/demo/",
	},
	Text{Content: "."},
}

var inlineCodeContent1 = []Text{
	Text{Content: "Inline "},
	Text{
		Content: "code",
		Code:    true,
	},
	Text{Content: " has "},
	Text{
		Content: "back-ticks around",
		Code:    true,
	},
	Text{Content: " it."},
}

var inlineCodeContent2 = []Text{
	Text{Content: "Blocks of code are either fenced by lines with three back-ticks "},
	Text{
		Content: "```",
		Code:    true,
	},
	Text{Content: ", or are indented with four spaces. I recommend only using the fenced code blocks -- they're easier and only they support syntax highlighting."},
}

// TableContent1 returns TableContent data.
func TableContent1() TableContent {
	return TableContent{
		Rows: [][]string{
			{"Tables", "Are", "Cool"},
			{"col 3 is", "right-aligned", "$1600"},
			{"col 2 is", "centered", "$12"},
			{"zebra stripes", "are neat", "$1"},
		},
		Alignments: []string{"L", "C", "R"},
	}
}

// TableContent2 returns TableContent data.
func TableContent2() TableContent {
	return TableContent{
		Rows: [][]string{
			{"Markdown", "Less", "Pretty"},
			{"still", "renders", "nicely"},
			{"1", "2", "3"},
		},
		Alignments: []string{"L", "L", "L"},
	}
}
