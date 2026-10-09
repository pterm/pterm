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
	RightDownLeftString:  "",
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
	// RightDownLeftString replaces the last HorizontalString in front of nodes
	// that have children (e.g. "┬" renders "├─┬ parent" instead of "├── parent").
	// If empty, the connector consists of HorizontalString only.
	RightDownLeftString string
	Indent              int
	Writer              io.Writer
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

// WithRightDownLeftString returns a new list with a specific RightDownLeftString.
func (p TreePrinter) WithRightDownLeftString(s string) *TreePrinter {
	p.RightDownLeftString = s
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
	s, err := p.Srender()
	if err != nil {
		return err
	}

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

	var result strings.Builder
	if p.Root.Text != "" {
		result.WriteString(p.TextStyle.Sprint(p.Root.Text))
		result.WriteByte('\n')
	}

	result.WriteString(walkOverTree(p.Root.Children, p, ""))

	return result.String(), nil
}

// walkOverTree is a recursive function,
// which analyzes a TreePrinter and connects the items with specific characters.
// Returns TreePrinter as string.
func walkOverTree(list []TreeNode, p TreePrinter, prefix string) string {
	var ret strings.Builder

	for i, item := range list {
		last := i == len(list)-1

		connector := p.TopRightDownString
		childPrefix := prefix + p.TreeStyle.Sprint(p.VerticalString) + strings.Repeat(" ", p.Indent+1)

		if last {
			connector = p.TopRightCornerString
			childPrefix = prefix + strings.Repeat(" ", p.Indent+2)
		}

		branch := strings.Repeat(p.HorizontalString, p.Indent)
		if len(item.Children) > 0 && p.RightDownLeftString != "" {
			branch = strings.Repeat(p.HorizontalString, max(p.Indent-1, 0)) + p.RightDownLeftString
		}

		ret.WriteString(prefix)
		ret.WriteString(p.TreeStyle.Sprint(connector + branch))
		ret.WriteByte(' ')
		ret.WriteString(p.TextStyle.Sprint(item.Text))
		ret.WriteByte('\n')

		if len(item.Children) > 0 {
			ret.WriteString(walkOverTree(item.Children, p, childPrefix))
		}
	}

	return ret.String()
}
