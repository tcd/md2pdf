package render

// List is an array of ListItems.
type List struct {
	Items   []ListItem `json:"items"`
	Ordered bool       `json:"ordered"`
}

// AddItems appends ListItems to a List's Items.
func (ls *List) AddItems(listItem ...ListItem) {
	ls.Items = append(ls.Items, listItem...)
}

// ListItem is a single list bullet, and may contain a nested list.
type ListItem struct {
	Contents `json:"contents"`
	Children List `json:"children"`
}

// HasChildren returns true if a ListItem contains a nested list.
func (li ListItem) HasChildren() bool {
	if len(li.Children.Items) > 0 {
		return true
	}
	return false
}
