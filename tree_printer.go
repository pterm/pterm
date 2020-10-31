package pterm

type Tree struct {
	Children []Tree
	Text     string
}

type LeveledList []LeveledListItem

type LeveledListItem struct {
	Level int
	Text  string
}

// DefaultTreeList contains standards, which can be used to print a TreeList.
var DefaultTreeList = TreeList{
	Style:                &ThemeDefault.ListCornerStyle, // TODO: Add own `TreeStyle` to theme.
	TextStyle:            &ThemeDefault.ListTextStyle,   // TODO: Add own `TreeTextStyle` to theme.
	TopRightCornerString: "└",
	HorizontalString:     "─",
	TopRightDownString:   "├",
	VerticalString:       "│",
	RightDownLeftString:  "┬",
}

// TreeList is able to render a list.
type TreeList struct {
	Root                 Tree
	Style                *Style
	TextStyle            *Style
	TopRightCornerString string
	TopRightDownString   string
	HorizontalString     string
	VerticalString       string
	RightDownLeftString  string
}

// WithStyle returns a new list with a specific bullet style.
func (p TreeList) WithStyle(style *Style) *TreeList {
	p.Style = style
	return &p
}

// WithTextStyle returns a new list with a specific text style.
func (p TreeList) WithTextStyle(style *Style) *TreeList {
	p.TextStyle = style
	return &p
}

// WithTopRightCornerString returns a new list with a specific bullet.
func (p TreeList) WithTopRightCornerString(s string) *TreeList {
	p.TopRightCornerString = s
	return &p
}

// WithTopRightDownStringOngoing returns a new list with a specific bullet.
func (p TreeList) WithTopRightDownStringOngoing(s string) *TreeList {
	p.TopRightDownString = s
	return &p
}

// WithHorizontalString returns a new list with a specific bullet.
func (p TreeList) WithHorizontalString(s string) *TreeList {
	p.HorizontalString = s
	return &p
}

// WithVerticalString returns a new list with a specific bullet.
func (p TreeList) WithVerticalString(s string) *TreeList {
	p.VerticalString = s
	return &p
}

// WithRoot returns a new list with a specific bullet.
func (p TreeList) WithRoot(root Tree) *TreeList {
	p.Root = root
	return &p
}

// Render prints the list to the terminal.
func (p TreeList) Render() {
	Println(p.Srender())
}

// Srender renders the list as a string.
func (p TreeList) Srender() string {
	return walkOverTree(p.Root.Children, p, "")
}

func walkOverTree(list []Tree, p TreeList, prefix string) string {
	var ret string
	for i, item := range list {
		if len(list) > i+1 {
			if len(item.Children) == 0 {
				ret += prefix + p.Style.Sprint(p.TopRightDownString) + p.Style.Sprint(p.HorizontalString) + item.Text + "\n"
			} else {
				ret += prefix + p.Style.Sprint(p.TopRightDownString) + p.Style.Sprint(p.RightDownLeftString) + item.Text + "\n"
				ret += walkOverTree(item.Children, p, prefix+p.Style.Sprint(p.VerticalString))
			}
		} else if len(list) == i+1 {
			if len(item.Children) == 0 {
				ret += prefix + p.Style.Sprint(p.TopRightCornerString) + p.Style.Sprint(p.HorizontalString) + item.Text + "\n"
			} else {
				ret += prefix + p.Style.Sprint(p.TopRightCornerString) + p.Style.Sprint(p.RightDownLeftString) + item.Text + "\n"
				ret += walkOverTree(item.Children, p, prefix+" ")
			}
		}
	}
	return ret
}

// NewTreeFromLeveledList converts a TreeItems list to a Tree and returns it.
func NewTreeFromLeveledList(lvledList []LeveledListItem) Tree {
	root := &Tree{
		Children: []Tree{},
		Text:     lvledList[0].Text,
	}

	for _, record := range lvledList {
		last := root
		for i := 0; i < record.Level; i++ {
			var lastIndex int
			if len(last.Children) > 0 {
				lastIndex = len(last.Children) - 1
			} else {
				last.Children = append(last.Children, Tree{})
			}
			last = &last.Children[lastIndex]
		}
		last.Children = append(last.Children, Tree{
			Children: []Tree{},
			Text:     record.Text,
		})
	}

	return *root
}
