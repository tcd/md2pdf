package model

// Contents models the contents of a Paragraph, Blockquote, List Item, or Table Cell.
type Contents struct {
	Content []Text `json:"content"`
}

// AddContent to a Contents.
func (c *Contents) AddContent(text Text) {
	if text.Text == "" {
		return
	}
	c.Content = append(c.Content, text)
}

// AddStr adds a string with no styles to Contents.
func (c *Contents) AddStr(str string) {
	if str == "" {
		return
	}
	text := Text{
		Text: str,
	}
	c.Content = append(c.Content, text)
}

// AllContent returns all enclosed Text.Content.
func (c Contents) AllContent() []string {
	allContent := make([]string, len(c.Content))
	for i, text := range c.Content {
		allContent[i] = text.Text
	}
	return allContent
}
