package pterm

// TreeNode is used as items in a Tree.
type TreeNode struct {
	Children []TreeNode
	Text     string
}

// LeveledList is a list, which contains multiple LeveledListItem.
type LeveledList []LeveledListItem

// LeveledListItem combines a text with a specific level.
// The level is the indent, which would normally be seen in a BulletList.
type LeveledListItem struct {
	Level int
	Text  string
}

// DefaultTree contains standards, which can be used to render a Tree.
var DefaultTree = Tree{
	TreeStyle:            &ThemeDefault.TreeStyle,
	TextStyle:            &ThemeDefault.TreeTextStyle,
	TopRightCornerString: "└",
	HorizontalString:     "─",
	TopRightDownString:   "├",
	VerticalString:       "│",
	RightDownLeftString:  "┬",
}

// Tree is able to render a list.
type Tree struct {
	Root                 TreeNode
	TreeStyle            *Style
	TextStyle            *Style
	TopRightCornerString string
	TopRightDownString   string
	HorizontalString     string
	VerticalString       string
	RightDownLeftString  string
}

// WithStyle returns a new list with a specific bullet style.
func (p Tree) WithStyle(style *Style) *Tree {
	p.TreeStyle = style
	return &p
}

// WithTextStyle returns a new list with a specific text style.
func (p Tree) WithTextStyle(style *Style) *Tree {
	p.TextStyle = style
	return &p
}

// WithTopRightCornerString returns a new list with a specific bullet.
func (p Tree) WithTopRightCornerString(s string) *Tree {
	p.TopRightCornerString = s
	return &p
}

// WithTopRightDownStringOngoing returns a new list with a specific bullet.
func (p Tree) WithTopRightDownStringOngoing(s string) *Tree {
	p.TopRightDownString = s
	return &p
}

// WithHorizontalString returns a new list with a specific bullet.
func (p Tree) WithHorizontalString(s string) *Tree {
	p.HorizontalString = s
	return &p
}

// WithVerticalString returns a new list with a specific bullet.
func (p Tree) WithVerticalString(s string) *Tree {
	p.VerticalString = s
	return &p
}

// WithRoot returns a new list with a specific bullet.
func (p Tree) WithRoot(root TreeNode) *Tree {
	p.Root = root
	return &p
}

// Render prints the list to the terminal.
func (p Tree) Render() {
	Println(p.Srender())
}

// Srender renders the list as a string.
func (p Tree) Srender() string {
	if p.TreeStyle == nil {
		p.TreeStyle = NewStyle()
	}
	if p.TextStyle == nil {
		p.TextStyle = NewStyle()
	}

	return walkOverTree(p.Root.Children, p, "")
}

// walkOverTree is a recursive function,
// which analyzes a Tree and connects the items with specific characters.
// Returns Tree as string.
func walkOverTree(list []TreeNode, p Tree, prefix string) string {
	var ret string
	for i, item := range list {
		if len(list) > i+1 { // if not last in list
			if len(item.Children) == 0 { // if there are no children
				ret += prefix + p.TreeStyle.Sprint(p.TopRightDownString) + p.TreeStyle.Sprint(p.HorizontalString) + item.Text + "\n"
			} else { // if there are children
				ret += prefix + p.TreeStyle.Sprint(p.TopRightDownString) + p.TreeStyle.Sprint(p.RightDownLeftString) + item.Text + "\n"
				ret += walkOverTree(item.Children, p, prefix+p.TreeStyle.Sprint(p.VerticalString))
			}
		} else if len(list) == i+1 { // if last in list
			if len(item.Children) == 0 { // if there are no children
				ret += prefix + p.TreeStyle.Sprint(p.TopRightCornerString) + p.TreeStyle.Sprint(p.HorizontalString) + item.Text + "\n"
			} else { // if there are children
				ret += prefix + p.TreeStyle.Sprint(p.TopRightCornerString) + p.TreeStyle.Sprint(p.RightDownLeftString) + item.Text + "\n"
				ret += walkOverTree(item.Children, p, prefix+" ")
			}
		}
	}
	return ret
}

// NewTreeFromLeveledList converts a TreeItems list to a TreeNode and returns it.
func NewTreeFromLeveledList(leveledListItems []LeveledListItem) TreeNode {
	root := &TreeNode{
		Children: []TreeNode{},
		Text:     leveledListItems[0].Text,
	}

	for _, record := range leveledListItems {
		last := root
		for i := 0; i < record.Level; i++ {
			var lastIndex int
			lastIndex = len(last.Children) - 1
			last = &last.Children[lastIndex]
		}
		last.Children = append(last.Children, TreeNode{
			Children: []TreeNode{},
			Text:     record.Text,
		})
	}

	return *root
}
