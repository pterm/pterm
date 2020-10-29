package pterm

type TreeListItem struct {
	Children []TreeListItem
	ItemName string
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
	ListItem []TreeListItem
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
func (p TreeList) WithItems(items []TreeListItem) *TreeList {
	p.ListItem = items
	return &p
}

// Render prints the list to the terminal.
func (p TreeList) Render() {
	Println(p.Srender())
}

// Srender renders the list as a string.
func (p TreeList) Srender() string {
	return temp(p.ListItem, p, "")
}

func (items LvlTreeListItems) ConvertLeveledListToTreeListItems(lvl, startPos int) []TreeListItem {
	count := 0
	var treeListItems []TreeListItem
	for i := startPos; i < len(items); i++ {
		if hasNextLowerLevel(items, i) && lvl != 0 {
			treeListItems = append(treeListItems, TreeListItem{Children: nil, ItemName: items[i].Text})
			return treeListItems
		}
		if items[i].Level == lvl {
			treeListItems = append(treeListItems, TreeListItem{Children: nil, ItemName: items[i].Text})
			if hasChildren(items, i) {
				treeListItems[count].Children = items.ConvertLeveledListToTreeListItems(lvl+1, i+1)
			}
			count++
		}
	}
	return treeListItems
}

func temp(list []TreeListItem, p TreeList, prefix string) string {
	var ret string
	for i, item := range list {
		if len(list) > i+1 {
			if item.Children == nil {
				ret += prefix + p.Style.Sprint(p.CornerOngoing) + p.Style.Sprint(p.Horizontal) + item.ItemName + "\n"
			} else {
				ret += prefix + p.Style.Sprint(p.CornerOngoing) + p.Style.Sprint(p.T) + item.ItemName + "\n"
				ret += temp(item.Children, p, prefix+p.Style.Sprint(p.Vertical))
			}
		} else if len(list) == i+1 {
			if item.Children == nil {
				ret += prefix + p.Style.Sprint(p.Corner) + p.Style.Sprint(p.Horizontal) + item.ItemName + "\n"
			} else {
				ret += prefix + p.Style.Sprint(p.Corner) + p.Style.Sprint(p.T) + item.ItemName + "\n"
				ret += temp(item.Children, p, prefix+" ")
			}
		}
	}
	return ret
}

func hasChildren(list LvlTreeListItems, i int) bool {
	if len(list) < i+2 {
		return false
	}
	if list[i+1].Level > list[i].Level {
		return true
	}
	return false
}
func hasNextLowerLevel(list LvlTreeListItems, i int) bool {
	if len(list) < i+2 {
		return false
	}
	if list[i+1].Level < list[i].Level {
		return true
	}
	return false
}
