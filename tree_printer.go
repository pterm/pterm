package pterm

import (
	"io"
	"strings"
)

// TreeNode is used as items in a TreePrinter.
type TreeNode struct {
	Children []TreeNode
	Text     string
}

// LeveledList is a list, which contains multiple LeveledListItem.
type LeveledList []LeveledListItem

// LeveledListItem combines a text with a specific level.
// The level is the indent, which would normally be seen in a BulletListPrinter.
type LeveledListItem struct {
	Level int
	Text  string
}

// DefaultTree contains standards, which can be used to render a TreePrinter.
var DefaultTree = TreePrinter{
	TreeStyle:            &ThemeDefault.TreeStyle,
	TextStyle:            &ThemeDefault.TreeTextStyle,
	TopRightCornerString: "└",
	HorizontalString:     "─",
	TopRightDownString:   "├",
	VerticalString:       "│",
	RightDownLeftString:  "┬",
	Indent:               2,
}

// TreePrinter is able to render a list.
type TreePrinter struct {
	Root                 TreeNode
	TreeStyle            *Style
	TextStyle            *Style
	TopRightCornerString string
	TopRightDownString   string
	HorizontalString     string
	VerticalString       string
	RightDownLeftString  string
	Indent               int
	Writer               io.Writer
}

// WithTreeStyle returns a new list with a specific tree style.
func (p TreePrinter) WithTreeStyle(style *Style) *TreePrinter {
	p.TreeStyle = style
	return &p
}

// WithTextStyle returns a new list with a specific text style.
func (p TreePrinter) WithTextStyle(style *Style) *TreePrinter {
	p.TextStyle = style
	return &p
}

// WithTopRightCornerString returns a new list with a specific TopRightCornerString.
func (p TreePrinter) WithTopRightCornerString(s string) *TreePrinter {
	p.TopRightCornerString = s
	return &p
}

// WithTopRightDownStringOngoing returns a new list with a specific TopRightDownString.
func (p TreePrinter) WithTopRightDownStringOngoing(s string) *TreePrinter {
	p.TopRightDownString = s
	return &p
}

// WithHorizontalString returns a new list with a specific HorizontalString.
func (p TreePrinter) WithHorizontalString(s string) *TreePrinter {
	p.HorizontalString = s
	return &p
}

// WithVerticalString returns a new list with a specific VerticalString.
func (p TreePrinter) WithVerticalString(s string) *TreePrinter {
	p.VerticalString = s
	return &p
}

// WithRoot returns a new list with a specific Root.
func (p TreePrinter) WithRoot(root TreeNode) *TreePrinter {
	p.Root = root
	return &p
}

// WithIndent returns a new list with a specific amount of spacing between the levels.
// Indent must be at least 1.
func (p TreePrinter) WithIndent(indent int) *TreePrinter {
	if indent < 1 {
		indent = 1
	}
	p.Indent = indent
	return &p
}

// WithWriter sets the Writer.
func (p TreePrinter) WithWriter(writer io.Writer) *TreePrinter {
	p.Writer = writer
	return &p
}

// Render prints the list to the terminal.
func (p TreePrinter) Render() error {
	s, _ := p.Srender()
	Fprintln(p.Writer, s)

	return nil
}

// Srender renders the list as a string.
func (p TreePrinter) Srender() (string, error) {
	if p.TreeStyle == nil {
		p.TreeStyle = NewStyle()
	}
	if p.TextStyle == nil {
		p.TextStyle = NewStyle()
	}

	var result string
	if p.Root.Text != "" {
		result += p.Root.Text + "\n"
	}
	result += walkOverTree(p.Root.Children, p, "")
	return result, nil
}

// walkOverTree is a recursive function,
// which analyzes a TreePrinter and connects the items with specific characters.
// Returns TreePrinter as string.
func walkOverTree(list []TreeNode, p TreePrinter, prefix string) string {
	var ret string
	for i, item := range list {
		if len(list) > i+1 { // if not last in list
			if len(item.Children) == 0 { // if there are no children
				ret += prefix + p.TreeStyle.Sprint(p.TopRightDownString) + strings.Repeat(p.TreeStyle.Sprint(p.HorizontalString), p.Indent) +
					p.TextStyle.Sprint(item.Text) + "\n"
			} else { // if there are children
				ret += prefix + p.TreeStyle.Sprint(p.TopRightDownString) + strings.Repeat(p.TreeStyle.Sprint(p.HorizontalString), p.Indent-1) +
					p.TreeStyle.Sprint(p.RightDownLeftString) + p.TextStyle.Sprint(item.Text) + "\n"
				ret += walkOverTree(item.Children, p, prefix+p.TreeStyle.Sprint(p.VerticalString)+strings.Repeat(" ", p.Indent-1))
			}
		} else if len(list) == i+1 { // if last in list
			if len(item.Children) == 0 { // if there are no children
				ret += prefix + p.TreeStyle.Sprint(p.TopRightCornerString) + strings.Repeat(p.TreeStyle.Sprint(p.HorizontalString), p.Indent) +
					p.TextStyle.Sprint(item.Text) + "\n"
			} else { // if there are children
				ret += prefix + p.TreeStyle.Sprint(p.TopRightCornerString) + strings.Repeat(p.TreeStyle.Sprint(p.HorizontalString), p.Indent-1) +
					p.TreeStyle.Sprint(p.RightDownLeftString) + p.TextStyle.Sprint(item.Text) + "\n"
				ret += walkOverTree(item.Children, p, prefix+strings.Repeat(" ", p.Indent))
			}
		}
	}
	return ret
}
