// Package content returns strings containing markdown content.
package content

import (
	"strings"
)

// Lorem returns 10 sentences of Lorem Ipsum text.
func Lorem() string {
	return "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent vel dapibus orci. Donec nec dignissim lectus. Cras nisi diam, hendrerit quis ex ut, porttitor posuere augue. Nulla elit ligula, laoreet quis egestas quis, interdum eu sem. Etiam luctus, diam in lacinia facilisis, dui sem iaculis lacus, eget porttitor est quam non sapien. Fusce vestibulum accumsan interdum. Vivamus tristique congue tincidunt. Duis molestie turpis vel varius maximus. Fusce tristique dolor arcu, id feugiat tellus condimentum ac."
}

// Lorem2 returns a long paragraph of Lorem Ipsum text.
func Lorem2() string {
	return `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent vel dapibus orci. Donec nec dignissim lectus. Cras nisi diam, hendrerit quis ex ut, porttitor posuere augue. Nulla elit ligula, laoreet quis egestas quis, interdum eu sem. Etiam luctus, diam in lacinia facilisis, dui sem iaculis lacus, eget porttitor est quam non sapien. Fusce vestibulum accumsan interdum. Vivamus tristique congue tincidunt. Duis molestie turpis vel varius maximus. Fusce tristique dolor arcu, id feugiat tellus condimentum ac.
In hac habitasse platea dictumst. Proin libero lectus, rutrum in porttitor in, mattis vitae nulla. Nulla aliquet purus nec magna ultricies tincidunt. Etiam dignissim ultrices leo luctus feugiat. Nunc congue vitae urna sed gravida. Praesent viverra leo nulla, sit amet sodales magna ultricies a. Aenean ornare semper enim, sed ultrices est gravida eget. Nullam pharetra lorem id ullamcorper maximus. Ut nisi felis, vehicula rutrum vehicula non, volutpat eu ante. Donec aliquam lectus sit amet ultrices aliquet. Nullam tempus interdum ipsum, vitae tincidunt mi dictum eleifend. Aenean vitae ante id justo convallis volutpat id sed dolor. Nam quis lectus eget nunc porttitor commodo vitae et sapien.
Cras vitae sagittis purus. Mauris rutrum posuere mattis. Aenean enim tortor, suscipit sed aliquam et, ornare ac neque. Suspendisse feugiat augue laoreet aliquam viverra. Sed eleifend lacus sit amet dictum volutpat. Phasellus sit amet libero tincidunt, eleifend nulla vitae, consequat orci. Curabitur ante mauris, finibus at eleifend at, porta a dolor. Praesent varius eros id quam semper pretium. Nunc non magna ut urna aliquam cursus. Etiam ut bibendum tortor, eget iaculis sem. Fusce sodales porta arcu, ac convallis lectus porta non. Morbi non interdum ex. Sed pellentesque imperdiet sodales.
Nullam vestibulum sagittis ultrices. In dignissim non justo quis iaculis. Etiam a mi vulputate, sodales ex molestie, iaculis diam. Nullam faucibus nunc justo, sed tristique ipsum feugiat at. Praesent vel molestie nisi. In facilisis, quam et lacinia cursus, mi tortor facilisis enim, mollis luctus ante augue non orci. Vestibulum posuere, risus dignissim vehicula ultrices, libero nisi elementum lacus, sed tempus lectus quam quis massa. Curabitur sagittis tincidunt velit, in volutpat leo efficitur in. Duis maximus nisi in lacus tincidunt, eget finibus orci fermentum. Donec condimentum, dolor sed dictum molestie, ipsum ipsum pharetra sem, ac porttitor diam sapien id tortor. Integer aliquet metus felis, ut pretium lorem fermentum eget. Nam mi nunc, hendrerit et congue nec, vehicula et est. Pellentesque accumsan quam et lacus accumsan tempus. Donec congue consequat magna. Mauris eu arcu vitae lacus mattis fringilla. Mauris eu gravida purus.
Curabitur quis convallis risus, in tristique felis. In facilisis arcu sagittis egestas semper. Sed eu ex at dolor pellentesque eleifend. Nam metus leo, dictum quis ultricies sed, vulputate et erat. Integer vel fringilla nunc. Ut vulputate pellentesque ante, vel sodales nunc consequat ut. Aenean dictum tristique sem, a gravida erat pretium ut. Sed sagittis pulvinar ipsum nec fermentum. Donec gravida, metus ut ultrices consequat, est felis pharetra nulla, eget sodales arcu sem pellentesque dui.`
}

// LoremList returns a slice of placeholder strings.
func LoremList() []string {
	return []string{
		"Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod " +
			"tempor incididunt ut labore et dolore magna aliqua.",
		"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut " +
			"aliquip ex ea commodo consequat.",
		"Duis aute irure dolor in reprehenderit in voluptate velit esse cillum " +
			"dolore eu fugiat nulla pariatur.",
		"Excepteur sint occaecat cupidatat non proident, sunt in culpa qui " +
			"officia deserunt mollit anim id est laborum.",
	}
}

// HeaderContent returns a string with different level headers.
func HeaderContent() string {
	return `# H1
## H2
### H3
#### H4
##### H5
###### H6

Alternatively, for H1 and H2, an underline-ish style:

Alt-H1
======

Alt-H2
------`
}

// EmContent returns a string with bold, italic, and stikeout elements.
func EmContent() string {
	return `Emphasis, aka italics, with *asterisks* or _underscores_.

Strong emphasis, aka bold, with **asterisks** or __underscores__.

Combined emphasis with **asterisks and _underscores_**.

Strikethrough uses two tildes. ~~Scratch this.~~`
}

// DirtyListContent returns a markdown list with unicode characters that can't be rendered properly by gofpdf.
func DirtyListContent() string {
	return `1. First ordered list item
2. Another item
⋅⋅* Unordered sub-list.
1. Actual numbers don't matter, just that it's a number
⋅⋅1. Ordered sub-list
4. And another item.

⋅⋅⋅You can have properly indented paragraphs within list items. Notice the blank line above, and the leading spaces (at least one, but we'll use three here to also align the raw Markdown).

⋅⋅⋅To have a line break without a paragraph, you will need to use two trailing spaces.⋅⋅
⋅⋅⋅Note that this line is separate, but within the same paragraph.⋅⋅
⋅⋅⋅(This is contrary to the typical GFM line break behaviour, where trailing spaces are not required.)

* Unordered list can use asterisks
- Or minuses
+ Or pluses`
}

// ListContent returns a string containing several markdown lists.
func ListContent() string {
	return `1. First ordered list item
2. Another item
  * Unordered sub-list.
1. Actual numbers don't matter, just that it's a number
  1. Ordered sub-list
4. And another item.

   You can have properly indented paragraphs within list items. Notice the blank line above, and the leading spaces (at least one, but we'll use three here to also align the raw Markdown).

   To have a line break without a paragraph, you will need to use two trailing spaces...
   Note that this line is separate, but within the same paragraph...
   (This is contrary to the typical GFM line break behaviour, where trailing spaces are not required.)

* Unordered list can use asterisks
- Or minuses
+ Or pluses`
}

// LinksContent returns a string containing several markdown links
func LinksContent() string {
	return `[I'm an inline-style link](https://www.google.com)

[I'm an inline-style link with title](https://www.google.com "Google's Homepage")

[I'm a reference-style link][Arbitrary case-insensitive reference text]

[I'm a relative reference to a repository file](../blob/master/LICENSE)

[You can use numbers for reference-style link definitions][1]

Or leave it empty and use the [link text itself].

URLs and URLs in angle brackets will automatically get turned into links.
http://www.example.com or <http://www.example.com> and sometimes
example.com (but not on Github, for example).

Some text to show that the reference links can follow later.

[arbitrary case-insensitive reference text]: https://www.mozilla.org
[1]: http://slashdot.org
[link text itself]: http://www.reddit.com`
}

// ImagesContent returns a string containing several markdown images.
func ImagesContent() string {
	return `Here's our logo (hover to see the title text):

Inline-style:
![alt text](https://github.com/adam-p/markdown-here/raw/master/src/common/images/icon48.png "Logo Title Text 1")

Reference-style:
![alt text][logo]

[logo]: https://github.com/adam-p/markdown-here/raw/master/src/common/images/icon48.png "Logo Title Text 2"`
}

// InlineCodeContent returns a string with some words wrapped in backticks.
func InlineCodeContent() string {
	return "Inline `code` has `back-ticks around` it."
}

// SyntaxHighContent returns a string containing markdown codeblocks.
func SyntaxHighContent() string {
	content := []string{
		"```javascript",
		`var s = "JavaScript syntax highlighting";`,
		"alert(s);",
		"```",
		"",
		"```python",
		`s = "Python syntax highlighting"`,
		"print s",
		"```",
		"",
		"```",
		"No language indicated, so no syntax highlighting.",
		"But let's throw in a <b>tag</b>.",
		"```",
	}

	return strings.Join(content, "\n")
}

// TablesContent returns a string containing markdown tables.
func TablesContent() string {
	return `Colons can be used to align columns.

| Tables        | Are           | Cool  |
| ------------- |:-------------:| -----:|
| col 3 is      | right-aligned | $1600 |
| col 2 is      | centered      |   $12 |
| zebra stripes | are neat      |    $1 |

There must be at least 3 dashes separating each header cell.
The outer pipes (|) are optional, and you don't need to make the
raw Markdown line up prettily. You can also use inline Markdown.

Markdown | Less | Pretty
--- | --- | ---
*Still* | _renders_ | **nicely**
1 | 2 | 3
`
}

// BlockquoteContent returns markdown blockquotes.
func BlockquoteContent() string {
	return `> Blockquotes are very handy in email to emulate reply text.
> This line is part of the same quote.

Quote break.

> This is a very long line that will still be quoted properly when it wraps. Oh boy let's keep writing to make sure this is long enough to actually wrap for everyone. Oh, you can *put* **Markdown** into a blockquote. `
}

// InlineHTMLContent returns an html definition list.
func InlineHTMLContent() string {
	return `<dl>
  <dt>Definition list</dt>
    <dd>Is something people use sometimes.</dd>

    <dt>Markdown in HTML</dt>
  <dd>Does *not* work **very** well. Use HTML <em>tags</em>.</dd>
</dl>
`
}

// HRContent returns a string with several horizontal rules.
func HRContent() string {
	return `Three or more...

---

Hyphens

***

Asterisks

___

Underscores`
}

// LineBreakContents returns strings interspersed with empty lines.
func LineBreakContents() string {
	return `Here's a line for us to start with.

This line is separated from the one above by two newlines, so it will be a *separate paragraph*.

This line is also a separate paragraph, but...
This line is only separated by a single newline, so it's a separate line in the *same paragraph*.`
}

// YouTubeContent1 returns an HTML link to a YouTube video.
func YouTubeContent1() string {
	return `<a href="http://www.youtube.com/watch?feature=player_embedded&v=YOUTUBE_VIDEO_ID_HERE" target="_blank">
  <img src="http://img.youtube.com/vi/YOUTUBE_VIDEO_ID_HERE/0.jpg" alt="IMAGE ALT TEXT HERE" width="240" height="180" border="10" />
</a>`
}

// YouTubeContent2 returns a markdown link to a YouTube video.
func YouTubeContent2() string {
	return "[![IMAGE ALT TEXT HERE](http://img.youtube.com/vi/YOUTUBE_VIDEO_ID_HERE/0.jpg)](http://www.youtube.com/watch?v=YOUTUBE_VIDEO_ID_HERE)"
}

// TocContent returns a slice of short strings.
var TocContent = []string{
	"Headers",
	"Emphasis",
	"Lists",
	"Links",
	"Images",
	"Code and Syntax Highlighting",
	"Tables",
	"Blockquotes",
	"Inline HTML",
	"Horizontal Rule",
	"Line Breaks",
	"YouTube Videos",
}

// ListContent1 returns a slice of string.
var ListContent1 = []string{
	"First ordered list item",
	"Another item",
	"Unordered sub-list.",
}

// ListContent2 returns a slice of string.
var ListContent2 = []string{
	"Actual numbers don't matter, just that it's a number",
	"Ordered sub-list",
	`And another item.

You can have properly indented paragraphs within list items. Notice the blank line above, and the leading spaces (at least one, but we'll use three here to also align the raw Markdown).

To have a line break without a paragraph, you will need to use two trailing spaces.
Note that this line is separate, but within the same paragraph.
(This is contrary to the typical GFM line break behaviour, where trailing spaces are not required.)
`,
}

// ListContent3 returns a slice of string.
var ListContent3 = []string{
	"Unordered list can use asterisks",
	"Or minuses",
	"Or pluses",
}
