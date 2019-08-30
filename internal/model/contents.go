package model

import "strings"

// Contents models the contents of a Paragraph, Blockquote, List Item, or Table Cell.
type Contents struct {
	Content []Text `json:"content"`
}

// AddContent to a Contents.
func (c *Contents) AddContent(text ...Text) {
	for _, txt := range text {
		if txt.Text != "" {
			c.Content = append(c.Content, txt)
		}
	}
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

// AllContent returns the values of all Text.Text contained in Contents.
func (c Contents) AllContent() []string {
	allContent := make([]string, len(c.Content))
	for i, text := range c.Content {
		allContent[i] = text.Text
	}
	return allContent
}

// JoinContent returns the values of all Text.Text contained in Contents.
func (c Contents) JoinContent() string {
	return strings.Join(c.AllContent(), "")
}
