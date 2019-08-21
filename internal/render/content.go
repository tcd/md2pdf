package render

// PContent1 returns content to write be used by FullP.
var PContent1 = Contents{
	Content: []Text{
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
	},
}

// PContent2 returns content to write be used by FullP.
var PContent2 = Contents{
	Content: []Text{
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
	},
}

// PContentEmphasis returns content to write be used by FullP.
var PContentEmphasis = Contents{
	Content: []Text{
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
	},
}

// PContentStrong returns content to write be used by FullP.
var PContentStrong = Contents{
	Content: []Text{
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
	},
}

// PContentCombined returns content to write be used by FullP.
var PContentCombined = Contents{
	Content: []Text{
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
	},
}

// PContentStrike returns content to write be used by FullP.
var PContentStrike = Contents{
	Content: []Text{
		Text{Content: "Strikethrough uses two tildes. "},
		Text{
			Content: "Scratch this.",
			Strike:  true,
		},
		Text{Content: "This isn't that easy to do in a PDF."},
	},
}

// PContentCodeAndSyntax returns content to write be used by FullP.
var PContentCodeAndSyntax = Contents{
	Content: []Text{
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
	},
}

var inlineCodeContent1 = Contents{
	Content: []Text{
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
	},
}

var inlineCodeContent2 = Contents{
	Content: []Text{
		Text{Content: "Blocks of code are either fenced by lines with three back-ticks "},
		Text{
			Content: "```",
			Code:    true,
		},
		Text{Content: ", or are indented with four spaces. I recommend only using the fenced code blocks -- they're easier and only they support syntax highlighting."},
	},
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

var bfParagraph1 = Contents{
	Content: []Text{
		Text{Content: "Blackfriday is a "},
		Text{
			Content: "Markdown",
			HREF:    "https://daringfireball.net/projects/markdown/",
		},
		Text{Content: " processor implemented in "},
		Text{
			Content: "Go",
			HREF:    "https://golang.org/",
		},
		Text{Content: ". It is paranoid about its input (so you can safely feed it user-supplied data), it is fast, it supports common extensions (tables, smart punctuation substitutions, etc.), and it is safe for all utf-8 (unicode) input."},
	},
}

func bfVersionList() List {
	var list List

	var item1 ListItem
	item1.AddStr("Cleaned up API")

	var item2 ListItem
	item2.AddStr("A separate call to ")
	item2.AddContent(Text{
		Content: "Parse",
		Code:    true,
		HREF:    "https://godoc.org/gopkg.in/russross/blackfriday.v2#Parse",
	})
	item2.AddStr(", which produces an abstract syntax tree for the document")

	var item3 ListItem
	item3.AddStr("Latest bug fixes")

	var item4 ListItem
	item4.AddStr("Flexibility to easily add your own rendering extensions")

	list.AddItems(item1, item2, item3, item4)

	return list
}

func bfFeaturesList() List {
	var list List

	var item1 ListItem
	item1.AddContent(Text{
		Content: "Compatibility",
		Bold:    true,
	})
	item1.AddStr(". The Markdown v1.0.3 test suite passes with the ")
	item1.AddContent(Text{
		Content: "--tidy",
		Code:    true,
	})
	item1.AddStr(" option. Without ")
	item1.AddContent(Text{
		Content: "--tidy",
		Code:    true,
	})
	item1.AddStr(", the differences are mostly in whitespace and entity escaping, where blackfriday is more consistent and cleaner.")

	var item2 ListItem
	item2.AddContent(Text{
		Content: "Common extensions",
		Bold:    true,
	})
	item2.AddStr(", including table support, fenced code blocks, autolinks, strikethroughs, non-strict emphasis, etc.")

	var item3 ListItem
	item3.AddContent(Text{
		Content: "Safety",
		Bold:    true,
	})
	item3.AddStr(". Blackfriday is paranoid when parsing, making it safe to feed untrusted user input without fear of bad things happening. The test suite stress tests this and there are no known inputs that make it crash.  If you find one, please let me know and send me the input that does it.")
	item3.AddStr("\nNOTE: \"safety\" in this context means ")
	item3.AddContent(Text{
		Content: "runtime safety only",
		Italic:  true,
	})
	item3.AddStr(". In order to protect yourself against JavaScript injection in untrusted content, see ")
	item3.AddContent(Text{
		Content: "this example",
		HREF:    "https://github.com/russross/blackfriday#sanitize-untrusted-content",
	})
	item3.AddStr(".")

	var item4 ListItem
	item4.AddContent(Text{
		Content: "Fast processing",
		Bold:    true,
	})
	item4.AddStr(". It is fast enough to render on-demand in most web applications without having to cache the output.")

	list.AddItems(item1, item2, item3, item4)

	return list
}

func exampleList() List {
	var list List

	var item1 ListItem
	item1.AddStr("Item 1. ")
	item1.AddStr("And some more text for good measure.")

	var item2 ListItem
	item2.AddStr("Item 2.")

	var item3 ListItem
	item3.AddStr("Item 3.")

	list.AddItems(item1, item2, item3)

	return list
}

func nestedList() List {
	var list List

	var item1 ListItem
	item1.AddStr("Item 1. ")
	item1.AddStr("And some more text for good measure.")

	var item2 ListItem
	item2.AddStr("Item 2.")

	var item3 ListItem
	item3.AddStr("Item 3.")
	var item3a ListItem
	item3a.AddStr("Item 3 has some sub points to it.")
	var item3b ListItem
	item3b.AddStr("Not too many though. ")
	item3b.AddStr("Just a few.")
	var item3c ListItem
	item3c.AddStr("Oh shit, it's kid has kids though")

	var item3c1 ListItem
	item3c1.AddStr("This is on the third level.")
	// var item3c2 ListItem
	// item3c2.AddStr("This is getting out of hand.")
	item3c.Children.AddItems(item3c1)
	// item3c.Children.AddItems(item3c1, item3c2)

	var item3d ListItem
	item3d.AddStr("Okay, back down to level 2.")

	item3.Children.AddItems(item3a, item3b, item3c, item3d)

	var item4 ListItem
	item4.AddStr("Item 4.")

	// var item5 ListItem
	// item5.AddStr("Item 5.")

	// list.AddItems(item1, item2, item3, item4, item5)
	list.AddItems(item1, item2, item3, item4)

	return list
}
