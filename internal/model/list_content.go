package model

// ListContent is an array of ListItems.
type ListContent struct {
	Items   []ListItem `json:"items"`
	Ordered bool       `json:"ordered"`
}

// AddItems appends ListItems to a List's Items.
func (ls *ListContent) AddItems(listItem ...ListItem) {
	ls.Items = append(ls.Items, listItem...)
}

// NewItem creates a new ListItem whose contents are comprised of
// the given Text arguments and adds that ListItem to a List's Items.
func (ls *ListContent) NewItem(texts ...Text) {
	newItem := ListItem{}
	newItem.AddContent(texts...)
	ls.AddItems(newItem)
}

// ListItem is a single list bullet, and may contain a nested list.
type ListItem struct {
	Contents `json:"contents"`
	Children ListContent `json:"children"`
}

// HasChildren returns true if a ListItem contains a nested list.
func (li ListItem) HasChildren() bool {
	if len(li.Children.Items) > 0 {
		return true
	}
	return false
}
