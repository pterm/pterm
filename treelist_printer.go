package pterm

type TreeListItem struct {
	Children []TreeListItem
	Text     string
}

type LvlTreeListItem struct {
	Level int
	Text  string
}

type LvlTreeListItems []LvlTreeListItem

// DefaultTreeList contains standards, which can be used to print a TreeList.
var DefaultTreeList = TreeList{
	Corner:        "└",
	Horizontal:    "─",
	CornerOngoing: "├",
	Vertical:      "│",
	T:             "┬",
	TextStyle:     &ThemeDefault.ListTextStyle,
	Style:         &ThemeDefault.ListCornerStyle,
}

// TreeList is able to render a list.
type TreeList struct {
	ListItem TreeListItem
	// LvlTreeListItems     LvlTreeListItems
	TextStyle     *Style
	Corner        string
	Style         *Style
	CornerOngoing string
	Horizontal    string
	Vertical      string
	T             string
}

// WithTextStyle returns a new list with a specific text style.
func (p TreeList) WithTextStyle(style *Style) *TreeList {
	p.TextStyle = style
	return &p
}

// WithCorner returns a new list with a specific bullet.
func (p TreeList) WithCorner(Corner string) *TreeList {
	p.Corner = Corner
	return &p
}

// WithStyle returns a new list with a specific bullet style.
func (p TreeList) WithStyle(style *Style) *TreeList {
	p.Style = style
	return &p
}

// WithCornerOngoing returns a new list with a specific bullet.
func (p TreeList) WithCornerOngoing(CornerOngoing string) *TreeList {
	p.CornerOngoing = CornerOngoing
	return &p
}

// WithBranch returns a new list with a specific bullet.
func (p TreeList) WithBranch(Branch string) *TreeList {
	p.Horizontal = Branch
	return &p
}

// WithRoot returns a new list with a specific bullet.
func (p TreeList) WithRoot(Root string) *TreeList {
	p.Vertical = Root
	return &p
}

// WithItems returns a new list with a specific bullet.
func (p TreeList) WithItems(items TreeListItem) *TreeList {
	p.ListItem = items
	return &p
}

// Render prints the list to the terminal.
func (p TreeList) Render() {
	Println(p.Srender())
}

// Srender renders the list as a string.
func (p TreeList) Srender() string {
	return temp(p.ListItem.Children, p, "")
}

// ConvertLeveledListToTreeListItems converts a LvlTreeListItems list to a TreeListItem and returns it.
func ConvertLeveledListToTreeListItems(lvledList LvlTreeListItems) TreeListItem {

	var root *TreeListItem
	root = &TreeListItem{
		Children: []TreeListItem{},
		Text:     lvledList[0].Text,
	}

	for _, record := range lvledList {
		last := root
		for i := 0; i < record.Level; i++ {
			var lastIndex int
			if len(last.Children) > 0 {
				lastIndex = len(last.Children) - 1
			} else {
				last.Children = append(last.Children, TreeListItem{})
			}
			last = &last.Children[lastIndex]
		}
		last.Children = append(last.Children, TreeListItem{
			Children: []TreeListItem{},
			Text:     record.Text,
		})
	}

	return *root
}

func temp(list []TreeListItem, p TreeList, prefix string) string {
	var ret string
	for i, item := range list {
		if len(list) > i+1 {
			if len(item.Children) == 0 {
				ret += prefix + p.Style.Sprint(p.CornerOngoing) + p.Style.Sprint(p.Horizontal) + item.Text + "\n"
			} else {
				ret += prefix + p.Style.Sprint(p.CornerOngoing) + p.Style.Sprint(p.T) + item.Text + "\n"
				ret += temp(item.Children, p, prefix+p.Style.Sprint(p.Vertical))
			}
		} else if len(list) == i+1 {
			if len(item.Children) == 0 {
				ret += prefix + p.Style.Sprint(p.Corner) + p.Style.Sprint(p.Horizontal) + item.Text + "\n"
			} else {
				ret += prefix + p.Style.Sprint(p.Corner) + p.Style.Sprint(p.T) + item.Text + "\n"
				ret += temp(item.Children, p, prefix+" ")
			}
		}
	}
	return ret
}
