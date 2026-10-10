package pterm

import (
	"os"
	"strings"
	"sync"
	"sync/atomic"
)

// Glyphs are the symbols printers draw with. GlyphsUnicode is the default;
// GlyphsASCII is used on terminals that cannot show Unicode.
type Glyphs struct {
	Success  string // ✔ a finished step, a success message
	Error    string // ✖ a failed step, an error message
	Warning  string // ▲ a warning
	Info     string // ● a note
	Debug    string // ◦ debug output
	Question string // ? a prompt waiting for an answer
	Pending  string // ○ something not started yet
	Skipped  string // – something that was skipped
	Canceled string // ✖ something interrupted

	Pointer   string // ❯ the cursor in a list
	Bullet    string // • a list item
	Checked   string // ◉ a selected option
	Unchecked string // ○ an option that is not selected
	Dot       string // · the separator between inline items

	Bar       string // ┃ the left bar of a callout
	RailStart string // ┌ the top of a flow
	Rail      string // │ the line of a flow
	RailEnd   string // └ the bottom of a flow

	Rule     string // ─ horizontal rules and table header lines
	Ellipsis string // … the end of cut text

	MeterFill  string   // █ the filled part of a meter or bar
	MeterTrack string   // ░ the empty part of a meter or bar
	Eighths    []string // ▏▎▍▌▋▊▉ partial horizontal blocks, 1/8 to 7/8
	Columns    []string // ▁▂▃▄▅▆▇█ vertical blocks, 1/8 to 8/8
	Spinner    []string // ⠋⠙⠹... spinner frames

	Box  BoxGlyphs
	Tree TreeGlyphs
}

// BoxGlyphs draw frames.
type BoxGlyphs struct {
	TopLeft, TopRight, BottomLeft, BottomRight string
	Horizontal, Vertical                       string
	TeeLeft, TeeRight, TeeTop, TeeBottom, Cross string
}

// TreeGlyphs draw trees.
type TreeGlyphs struct {
	Branch string // ├─
	Last   string // └─
	Line   string // │
	Space  string // the indent under a last branch
}

// GlyphsUnicode is the default glyph set.
var GlyphsUnicode = Glyphs{
	Success:  "✔",
	Error:    "✖",
	Warning:  "▲",
	Info:     "●",
	Debug:    "◦",
	Question: "?",
	Pending:  "○",
	Skipped:  "–",
	Canceled: "✖",

	Pointer:   "❯",
	Bullet:    "•",
	Checked:   "◉",
	Unchecked: "○",
	Dot:       "·",

	Bar:       "┃",
	RailStart: "┌",
	Rail:      "│",
	RailEnd:   "└",

	Rule:     "─",
	Ellipsis: "…",

	MeterFill:  "█",
	MeterTrack: "░",
	Eighths:    []string{"▏", "▎", "▍", "▌", "▋", "▊", "▉"},
	Columns:    []string{"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█"},
	Spinner:    []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},

	Box: BoxGlyphs{
		TopLeft: "╭", TopRight: "╮", BottomLeft: "╰", BottomRight: "╯",
		Horizontal: "─", Vertical: "│",
		TeeLeft: "├", TeeRight: "┤", TeeTop: "┬", TeeBottom: "┴", Cross: "┼",
	},
	Tree: TreeGlyphs{Branch: "├─", Last: "└─", Line: "│ ", Space: "  "},
}

// GlyphsASCII only uses printable ASCII, for terminals without Unicode
// fonts such as the Linux console.
var GlyphsASCII = Glyphs{
	Success:  "v",
	Error:    "x",
	Warning:  "!",
	Info:     "i",
	Debug:    ".",
	Question: "?",
	Pending:  "o",
	Skipped:  "-",
	Canceled: "x",

	Pointer:   ">",
	Bullet:    "*",
	Checked:   "*",
	Unchecked: "o",
	Dot:       "-",

	Bar:       "|",
	RailStart: "+",
	Rail:      "|",
	RailEnd:   "`",

	Rule:     "-",
	Ellipsis: "~",

	MeterFill:  "#",
	MeterTrack: ".",
	Eighths:    []string{"", "", "", "", "", "", ""},
	Columns:    []string{".", ".", ":", ":", "|", "|", "#", "#"},
	Spinner:    []string{"|", "/", "-", "\\"},

	Box: BoxGlyphs{
		TopLeft: "+", TopRight: "+", BottomLeft: "+", BottomRight: "+",
		Horizontal: "-", Vertical: "|",
		TeeLeft: "+", TeeRight: "+", TeeTop: "+", TeeBottom: "+", Cross: "+",
	},
	Tree: TreeGlyphs{Branch: "|-", Last: "`-", Line: "| ", Space: "  "},
}

var (
	activeGlyphs atomic.Pointer[Glyphs]
	glyphsOnce   sync.Once
)

// SetGlyphs makes g the glyph set of every printer.
func SetGlyphs(g Glyphs) {
	glyphsOnce.Do(func() {})
	activeGlyphs.Store(&g)
}

// CurrentGlyphs returns the active glyph set. On first use it is
// GlyphsASCII when PTERM_ASCII is set to anything but "0", when TERM is
// "linux" (the Linux console), or on a Windows console that cannot show
// Unicode; otherwise it is GlyphsUnicode.
func CurrentGlyphs() Glyphs {
	glyphsOnce.Do(func() {
		g := GlyphsUnicode
		if asciiTerminal() {
			g = GlyphsASCII
		}

		activeGlyphs.Store(&g)
	})

	return *activeGlyphs.Load()
}

func asciiTerminal() bool {
	if v := os.Getenv("PTERM_ASCII"); v != "" {
		return v != "0" && !strings.EqualFold(v, "false")
	}

	return os.Getenv("TERM") == "linux" || legacyConsole()
}
