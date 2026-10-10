package pterm

import (
	"io"
	"strings"

	"github.com/pterm/pterm/internal/text"
)

// Definition is a term and what it means, such as a flag and its help.
type Definition struct {
	Term        string
	Description string
}

// Def returns a definition of term.
func Def(term, description string) Definition {
	return Definition{Term: term, Description: description}
}

// DefinitionGroup is a list of definitions under an optional heading.
type DefinitionGroup struct {
	// Title is printed above the group, uppercase and muted. Empty prints
	// no heading.
	Title string
	// Definitions are the entries of the group.
	Definitions []Definition
}

// DefinitionsPrinter renders terms and their descriptions in two aligned
// columns, like the help of a command line tool:
//
//	FLAGS
//	  -v, --verbose  print what is happening
//	  -o, --output   where to write the result, defaults to standard
//	                 output
//
// Descriptions wrap under themselves. A term much wider than the others
// gets its description on the next line, so one long term does not push
// every description to the right.
type DefinitionsPrinter struct {
	// Groups are the groups in order.
	Groups []DefinitionGroup
	// TermStyle styles the terms.
	TermStyle Style
	// DescriptionStyle styles the descriptions.
	DescriptionStyle Style
	// HeadingStyle styles the group headings.
	HeadingStyle Style
	// HeadingUppercase prints the headings in uppercase.
	HeadingUppercase bool
	// Indent is the number of spaces before the terms of a group with a
	// heading. Groups without a heading are not indented.
	Indent int
	// Gap is the number of spaces between the term column and the
	// descriptions.
	Gap int
	// MaxTermWidth is the widest the term column gets. Longer terms put
	// their description on the next line. Zero picks two fifths of the
	// terminal width, and no limit when the width is unknown.
	MaxTermWidth int
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
}

// Definitions is the default definitions printer.
var Definitions = DefinitionsPrinter{
	TermStyle:        ToneAccent.Style(),
	DescriptionStyle: ToneText.Style(),
	HeadingStyle:     ToneMuted.Style(),
	HeadingUppercase: true,
	Indent:           2,
	Gap:              2,
}

// WithDefinitions returns a copy of p with defs as its only group, without
// a heading.
func (p DefinitionsPrinter) WithDefinitions(defs ...Definition) *DefinitionsPrinter {
	p.Groups = []DefinitionGroup{{Definitions: defs}}
	return &p
}

// WithGroup returns a copy of p with one more group, under the heading
// title.
func (p DefinitionsPrinter) WithGroup(title string, defs ...Definition) *DefinitionsPrinter {
	p.Groups = append(append([]DefinitionGroup(nil), p.Groups...), DefinitionGroup{Title: title, Definitions: defs})
	return &p
}

// WithGroups returns a copy of p that renders groups.
func (p DefinitionsPrinter) WithGroups(groups ...DefinitionGroup) *DefinitionsPrinter {
	p.Groups = groups
	return &p
}

// WithTermStyle returns a copy of p whose terms use s.
func (p DefinitionsPrinter) WithTermStyle(s Style) *DefinitionsPrinter {
	p.TermStyle = s
	return &p
}

// WithDescriptionStyle returns a copy of p whose descriptions use s.
func (p DefinitionsPrinter) WithDescriptionStyle(s Style) *DefinitionsPrinter {
	p.DescriptionStyle = s
	return &p
}

// WithHeadingStyle returns a copy of p whose headings use s.
func (p DefinitionsPrinter) WithHeadingStyle(s Style) *DefinitionsPrinter {
	p.HeadingStyle = s
	return &p
}

// WithHeadingUppercase returns a copy of p that prints headings in
// uppercase, or as written with false.
func (p DefinitionsPrinter) WithHeadingUppercase(b ...bool) *DefinitionsPrinter {
	p.HeadingUppercase = flag(b)
	return &p
}

// WithIndent returns a copy of p that indents grouped terms by n spaces.
func (p DefinitionsPrinter) WithIndent(n int) *DefinitionsPrinter {
	p.Indent = max(n, 0)
	return &p
}

// WithGap returns a copy of p with n spaces between terms and
// descriptions.
func (p DefinitionsPrinter) WithGap(n int) *DefinitionsPrinter {
	p.Gap = max(n, 1)
	return &p
}

// WithMaxTermWidth returns a copy of p whose term column is at most n
// columns wide.
func (p DefinitionsPrinter) WithMaxTermWidth(n int) *DefinitionsPrinter {
	p.MaxTermWidth = max(n, 0)
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p DefinitionsPrinter) WithWriter(w io.Writer) *DefinitionsPrinter {
	p.Writer = w
	return &p
}

func (p DefinitionsPrinter) out() io.Writer { return orDefault(p.Writer) }

// Render prints the definitions.
func (p DefinitionsPrinter) Render() error {
	s, err := p.Srender()
	return renderTo(p.out(), s, err)
}

// Srender returns the definitions. It returns ErrEmpty when there are none.
func (p DefinitionsPrinter) Srender() (string, error) {
	c := canvasFor(p.out())
	column := p.termColumn(c)

	if column < 0 {
		return "", ErrEmpty
	}

	var lines []string

	for _, g := range p.Groups {
		if len(g.Definitions) == 0 && g.Title == "" {
			continue
		}

		if len(lines) > 0 {
			lines = append(lines, "")
		}

		indent := 0

		if g.Title != "" {
			title := g.Title
			if p.HeadingUppercase {
				title = upperVisible(title)
			}

			lines = append(lines, c.style(p.HeadingStyle, c.cut(title, c.width)))
			indent = p.Indent
		}

		for _, d := range g.Definitions {
			lead := strings.Repeat(" ", indent) + c.style(p.TermStyle, d.Term)
			body := []segment{{d.Description, p.DescriptionStyle}}
			lines = append(lines, hanging(c, lead, column, p.Gap, body)...)
		}
	}

	if len(lines) == 0 {
		return "", ErrEmpty
	}

	return strings.Join(lines, "\n"), nil
}

// termColumn returns where the term column ends, counted from the left
// edge and including the indent, or -1 without definitions.
func (p DefinitionsPrinter) termColumn(c canvas) int {
	limit := p.MaxTermWidth
	if limit == 0 && c.width > 0 {
		limit = c.width * 2 / 5
	}

	column, longest := -1, 0

	for _, g := range p.Groups {
		indent := 0
		if g.Title != "" {
			indent = p.Indent
		}

		for _, d := range g.Definitions {
			w := indent + text.Width(d.Term)
			longest = max(longest, w)

			if limit <= 0 || w <= limit {
				column = max(column, w)
			}
		}
	}

	if column < 0 && longest > 0 {
		// Every term is too long: put all descriptions on their own line,
		// one indent in.
		return p.Indent
	}

	return column
}
