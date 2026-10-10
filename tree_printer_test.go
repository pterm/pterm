package pterm_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal/vt"
)

var commandTree = []pterm.TreeNode{
	{Text: "doctor", Description: "diagnostics"},
	{Text: "mcp", Description: "manage MCP servers", Children: []pterm.TreeNode{
		{Text: "start", Description: "boots a server"},
	}},
	{Text: "disk", Description: "local disk utilities", Children: []pterm.TreeNode{
		{Text: "clean"},
	}},
}

func TestTreeSnapshots(t *testing.T) {
	snap(t, "Default", srender(t, pterm.Tree.WithNodes(commandTree...).Srender))
	snap(t, "Root", srender(t, pterm.Tree.WithRoot(pterm.TreeNode{Text: "mjw", Children: commandTree}).Srender))
	snap(t, "Indent", srender(t, pterm.Tree.WithIndent(2).WithNodes(commandTree...).Srender))
	snap(t, "Styled", srender(t, pterm.Tree.WithTextStyle(pterm.ToneAccent.Style()).WithLineTone(pterm.ToneMuted).
		WithNodes(commandTree...).Srender))
}

func TestTreeLook(t *testing.T) {
	setup(t)

	out := capture(t, func() { _ = pterm.Tree.WithNodes(commandTree...).Render() })
	want := "├─ doctor    diagnostics\n" +
		"├─ mcp       manage MCP servers\n" +
		"│  └─ start  boots a server\n" +
		"└─ disk      local disk utilities\n" +
		"   └─ clean\n"

	if out != want {
		t.Errorf("output:\n%s\nwant:\n%s", out, want)
	}
}

func TestTreeASCII(t *testing.T) {
	setup(t)
	pterm.SetGlyphs(pterm.GlyphsASCII)

	got, _ := pterm.Tree.WithNodes(commandTree...).Srender()
	if !strings.HasPrefix(plain(got), "|- doctor") || !strings.Contains(plain(got), "`- disk") {
		t.Errorf("got:\n%s", plain(got))
	}
}

func TestTreeEmpty(t *testing.T) {
	setup(t)

	if _, err := pterm.Tree.Srender(); !errors.Is(err, pterm.ErrEmpty) {
		t.Errorf("err = %v, want ErrEmpty", err)
	}
}

func TestTreeCutsDescriptions(t *testing.T) {
	setup(t)

	term := vt.New(24, 10)
	_ = pterm.Tree.WithWriter(term).WithNodes(commandTree...).Render()

	for _, line := range strings.Split(term.Screen(), "\n") {
		if w := len([]rune(line)); w > 24 {
			t.Errorf("line %q is %d columns wide", line, w)
		}
	}

	if !strings.Contains(term.Screen(), "manage MCP…") {
		t.Errorf("description not cut:\n%s", term.Screen())
	}
}

func TestLeveledListTree(t *testing.T) {
	root := pterm.LeveledList{
		{Level: 0, Text: "a"},
		{Level: 1, Text: "a1"},
		{Level: 3, Text: "a1x"}, // too deep, becomes a child of a1
		{Level: 0, Text: "b"},
		{Level: -1, Text: "c"},
	}.Tree()

	if len(root.Children) != 3 || root.Children[2].Text != "c" {
		t.Fatalf("root children = %+v", root.Children)
	}

	if got := root.Children[0].Children[0].Children[0].Text; got != "a1x" {
		t.Errorf("deep item = %q", got)
	}

	if (pterm.LeveledList{}).Tree().Children != nil {
		t.Error("empty list has children")
	}
}
