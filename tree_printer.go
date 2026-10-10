package pterm

import (
	"io"
	"strings"

	"github.com/pterm/pterm/internal/text"
)

// TreeNode is one node of a tree: its text, an optional description that
// is shown in an aligned column, and its children.
type TreeNode struct {
	Text        string
	Description string
	Children    []TreeNode
}

// LeveledList is a flat way to write a tree: every item states its depth.
// Tree turns it into a TreeNode.
type LeveledList []LeveledListItem

// LeveledListItem is one item of a LeveledList.
type LeveledListItem struct {
	Level       int
	Text        string
	Description string
}

// Tree returns the tree that l describes, under a root without text. An
// item that is more than one level deeper than the item before it is moved
// up to be that item's child.
func (l LeveledList) Tree() TreeNode {
	var root TreeNode

	prev := -1

	for _, item := range l {
		level := min(max(item.Level, 0), prev+1)
		prev = level

		parent := &root
		for range level {
			parent = &parent.Children[len(parent.Children)-1]
		}

		parent.Children = append(parent.Children, TreeNode{Text: item.Text, Description: item.Description})
	}

	return root
}

// TreePrinter renders a tree with its lines in the Border tone and an
// optional, aligned description column:
//
//	├─ doctor    diagnostics
//	├─ mcp       manage MCP servers
//	│  └─ start  boots a server
//	└─ disk      local disk utilities
//	   └─ clean
//
// The text of the root is printed above its children when it is set.
type TreePrinter struct {
	// Root is the tree. Its text is the first line; its children follow.
	Root TreeNode
	// LineTone colors the branch lines.
	LineTone Tone
	// RootStyle styles the text of the root.
	RootStyle Style
	// TextStyle styles the text of the other nodes.
	TextStyle Style
	// DescriptionStyle styles the descriptions.
	DescriptionStyle Style
	// Lines are the branch glyphs. Empty fields use the glyph set.
	Lines TreeGlyphs
	// Indent is the number of extra columns every level is indented by.
	Indent int
	// Gap is the space between the widest node and the descriptions.
	Gap int
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
}

// Tree is the default tree.
var Tree = TreePrinter{
	LineTone:         ToneBorder,
	RootStyle:        ToneText.Bold(),
	TextStyle:        ToneText.Style(),
	DescriptionStyle: ToneMuted.Style(),
	Gap:              2,
}

// WithRoot returns a copy of p that renders root.
func (p TreePrinter) WithRoot(root TreeNode) *TreePrinter {
	p.Root = root
	return &p
}

// WithNodes returns a copy of p that renders nodes under a root without
// text.
func (p TreePrinter) WithNodes(nodes ...TreeNode) *TreePrinter {
	p.Root = TreeNode{Children: nodes}
	return &p
}

// WithLineTone returns a copy of p that draws its lines in tone t.
func (p TreePrinter) WithLineTone(t Tone) *TreePrinter {
	p.LineTone = t
	return &p
}

// WithRootStyle returns a copy of p whose root text uses s.
func (p TreePrinter) WithRootStyle(s Style) *TreePrinter {
	p.RootStyle = s
	return &p
}

// WithTextStyle returns a copy of p whose node text uses s.
func (p TreePrinter) WithTextStyle(s Style) *TreePrinter {
	p.TextStyle = s
	return &p
}

// WithDescriptionStyle returns a copy of p whose descriptions use s.
func (p TreePrinter) WithDescriptionStyle(s Style) *TreePrinter {
	p.DescriptionStyle = s
	return &p
}

// WithLines returns a copy of p that draws its branches with g.
func (p TreePrinter) WithLines(g TreeGlyphs) *TreePrinter {
	p.Lines = g
	return &p
}

// WithIndent returns a copy of p that indents every level by n more
// columns.
func (p TreePrinter) WithIndent(n int) *TreePrinter {
	p.Indent = max(n, 0)
	return &p
}

// WithGap returns a copy of p with n spaces before the descriptions.
func (p TreePrinter) WithGap(n int) *TreePrinter {
	p.Gap = max(n, 1)
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p TreePrinter) WithWriter(w io.Writer) *TreePrinter {
	p.Writer = w
	return &p
}

func (p TreePrinter) out() io.Writer { return orDefault(p.Writer) }

// Render prints the tree.
func (p TreePrinter) Render() error {
	s, err := p.Srender()
	return renderTo(p.out(), s, err)
}

// treeLine is one rendered node before the description column is known.
type treeLine struct {
	prefix string // branch lines, unstyled
	text   string
	desc   string
	root   bool
}

// Srender returns the tree. It returns ErrEmpty for a root without text
// and children.
func (p TreePrinter) Srender() (string, error) {
	if p.Root.Text == "" && len(p.Root.Children) == 0 {
		return "", ErrEmpty
	}

	c := canvasFor(p.out())
	g := p.glyphs(c)

	var lines []treeLine
	if p.Root.Text != "" {
		lines = append(lines, treeLine{text: p.Root.Text, desc: p.Root.Description, root: true})
	}

	lines = p.walk(g, c.g.Rule, p.Root.Children, "", lines)

	column := 0

	for _, l := range lines {
		if l.desc != "" {
			column = max(column, text.Width(l.prefix)+text.Width(l.text))
		}
	}

	out := make([]string, len(lines))

	for i, l := range lines {
		st := p.TextStyle
		if l.root {
			st = p.RootStyle
		}

		node := l.text
		if c.width > 0 {
			node = c.cut(node, max(c.width-text.Width(l.prefix), 1))
		}

		line := c.paint(p.LineTone, l.prefix) + c.style(st, node)
		descColumn := column + p.Gap

		// A description needs room for at least a few characters; on a very
		// narrow terminal it is left out rather than cut to nothing.
		if l.desc != "" && (c.width <= 0 || descColumn+3 <= c.width) {
			desc := l.desc
			if c.width > 0 {
				desc = c.cut(desc, c.width-descColumn)
			}

			used := text.Width(l.prefix) + text.Width(node)
			line += strings.Repeat(" ", max(descColumn-used, 1)) + c.style(p.DescriptionStyle, desc)
		}

		out[i] = line
	}

	return strings.Join(out, "\n"), nil
}

// walk appends the lines of nodes and their children.
func (p TreePrinter) walk(g TreeGlyphs, rule string, nodes []TreeNode, prefix string, lines []treeLine) []treeLine {
	extra := strings.Repeat(" ", p.Indent)
	dash := strings.Repeat(rule, p.Indent)

	for i, n := range nodes {
		branch, next := g.Branch, g.Line
		if i == len(nodes)-1 {
			branch, next = g.Last, g.Space
		}

		lines = append(lines, treeLine{prefix: prefix + branch + dash + " ", text: n.Text, desc: n.Description})
		lines = p.walk(g, rule, n.Children, prefix+next+extra+" ", lines)
	}

	return lines
}

// glyphs returns the branch glyphs: p.Lines, filled up from the canvas.
func (p TreePrinter) glyphs(c canvas) TreeGlyphs {
	g := p.Lines
	if g.Branch == "" {
		g.Branch = c.g.Tree.Branch
	}

	if g.Last == "" {
		g.Last = c.g.Tree.Last
	}

	if g.Line == "" {
		g.Line = c.g.Tree.Line
	}

	if g.Space == "" {
		g.Space = c.g.Tree.Space
	}

	return g
}
