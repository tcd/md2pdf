package model

// Text models text strings that make up a Paragraph, Blockquote, List Item, or Table Cell.
type Text struct {
	Text   string `json:"text"`
	Bold   bool   `json:"bold"`
	Italic bool   `json:"italic"`
	Code   bool   `json:"code"`
	Strike bool   `json:"strike"`
	HREF   string `json:"href"`
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
