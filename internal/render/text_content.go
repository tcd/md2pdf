package render

// Text models text strings that make up a Paragraph, Blockquote, List Item, or Table Cell.
type Text struct {
	Content string `json:"content"`
	Bold    bool   `json:"bold"`
	Italic  bool   `json:"italic"`
	Code    bool   `json:"code"`
	Strike  bool   `json:"strike"`
	HREF    string `json:"href"`
}

// Contents models the contents of a Paragraph, Blockquote, List Item, or Table Cell.
type Contents struct {
	Content []Text `json:"content"`
}

// AddContent to a Contents.
func (c *Contents) AddContent(text Text) {
	c.Content = append(c.Content, text)
}

// AddStr adds a string with no styles to Contents.
func (c *Contents) AddStr(str string) {
	text := Text{
		Content: str,
	}
	c.Content = append(c.Content, text)
}

// AllContent returns all enclosed Text.Content.
func (c Contents) AllContent() []string {
	allContent := make([]string, len(c.Content))
	for i, text := range c.Content {
		allContent[i] = text.Content
	}
	return allContent
}

// Copy returns a new Text struct with the same values as the Text it was called from.
// Except the Content field, that's empty.
func (txt Text) Copy() Text {
	var newText Text
	if txt.Bold {
		newText.Bold = true
	}
	if txt.Italic {
		newText.Italic = true
	}
	if txt.Code {
		newText.Code = true
	}
	if txt.Strike {
		newText.Strike = true
	}
	newText.HREF = txt.HREF[:]
	return newText
}
