package pterm_test

import (
	"bytes"
	"encoding/csv"
	"errors"
	"strings"
	"testing"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal/vt"
)

var cacheData = pterm.TableData{
	{"Cache", "State", "Size"},
	{"npm", "stale", "1.20 GB"},
	{"nub", "fresh", "393 MB"},
}

func TestTableSnapshots(t *testing.T) {
	snap(t, "Default", srender(t, pterm.Table.WithData(cacheData).Srender))
	snap(t, "NoHeader", srender(t, pterm.Table.WithHasHeader(false).WithData(cacheData).Srender))
	snap(t, "Boxed", srender(t, pterm.Table.WithBoxed().WithTitle("Caches").WithData(cacheData).Srender))
	snap(t, "BoxedRowSeparators", srender(t, pterm.Table.WithBoxed().WithRowSeparators().WithData(cacheData).Srender))
	snap(t, "RowSeparators", srender(t, pterm.Table.WithRowSeparators().WithData(cacheData).Srender))
	snap(t, "MultiLine", srender(t, pterm.Table.WithData(pterm.TableData{
		{"Command", "Description"},
		{"init", "Create a project\nin this folder"},
		{"build", "Compile"},
	}).Srender))
	snap(t, "Styles", srender(t, pterm.Table.WithData(cacheData).
		WithColumnStyle(1, pterm.ToneAccent.Style()).
		WithRowStyle(1, pterm.ToneError.Style()).Srender))
	snap(t, "AlternateRowStyle", srender(t, pterm.Table.WithData(append(cacheData, []string{"go", "fresh", "2.84 GB"})).
		WithAlternateRowStyle(pterm.ToneText.Style().On(pterm.ToneBorder)).Srender))
	snap(t, "CJK", srender(t, pterm.Table.WithData(pterm.TableData{
		{"Name", "Note"},
		{"张", "这是一个测试"},
		{"Paul", "ok"},
	}).Srender))
}

func TestTableDefaultLook(t *testing.T) {
	setup(t)

	out := capture(t, func() { _ = pterm.Table.WithData(cacheData).Render() })
	want := "CACHE  STATE     SIZE\n" +
		"─────────────────────\n" +
		"npm    stale  1.20 GB\n" +
		"nub    fresh   393 MB\n"

	if out != want {
		t.Errorf("output:\n%s\nwant:\n%s", out, want)
	}
}

func TestTableEmpty(t *testing.T) {
	setup(t)

	if _, err := pterm.Table.Srender(); !errors.Is(err, pterm.ErrEmpty) {
		t.Errorf("err = %v, want ErrEmpty", err)
	}

	if err := pterm.Table.WithData(pterm.TableData{{}}).Render(); !errors.Is(err, pterm.ErrEmpty) {
		t.Errorf("err = %v, want ErrEmpty", err)
	}
}

func TestTableAlignment(t *testing.T) {
	setup(t)

	data := pterm.TableData{{"a", "b"}, {"x", "1"}, {"long", "22"}}

	got, _ := pterm.Table.WithHeaderRule(false).WithData(data).WithColumnAlignment(0, pterm.AlignRight).
		WithColumnAlignment(1, pterm.AlignLeft).Srender()
	want := "   A  B\n   x  1\nlong  22"

	if plain(got) != want {
		t.Errorf("got:\n%s\nwant:\n%s", plain(got), want)
	}

	got, _ = pterm.Table.WithHeaderRule(false).WithData(data).WithRightAlignment().Srender()
	if want := "   A   B\n   x   1\nlong  22"; plain(got) != want {
		t.Errorf("right:\n%s\nwant:\n%s", plain(got), want)
	}

	got, _ = pterm.Table.WithHeaderRule(false).WithData(data).WithLeftAlignment().Srender()
	if want := "A     B\nx     1\nlong  22"; plain(got) != want {
		t.Errorf("left:\n%s\nwant:\n%s", plain(got), want)
	}
}

func TestTableNumberDetection(t *testing.T) {
	setup(t)

	got, _ := pterm.Table.WithHasHeader(false).WithData(pterm.TableData{
		{"a", "1m 34s", "v1.2.3", "12%"},
		{"bb", "5s", "v1.10.0", "-"},
	}).Srender()
	want := "a   1m 34s  v1.2.3   12%\nbb      5s  v1.10.0    -"

	if plain(got) != want {
		t.Errorf("got:\n%s\nwant:\n%s", plain(got), want)
	}
}

func TestTableCutsToWidth(t *testing.T) {
	setup(t)

	term := vt.New(30, 10)
	data := pterm.TableData{
		{"Name", "Path"},
		{"npm", "~/.npm/_cacache/index-v5/aa/bb"},
		{"go", "~/go/pkg/mod/cache/download"},
	}

	if err := pterm.Table.WithWriter(term).WithData(data).Render(); err != nil {
		t.Fatal(err)
	}

	for _, line := range strings.Split(term.Screen(), "\n") {
		if w := len([]rune(line)); w > 30 {
			t.Errorf("line %q is %d columns wide", line, w)
		}
	}

	if !strings.Contains(term.Screen(), "…") {
		t.Errorf("no ellipsis in:\n%s", term.Screen())
	}

	// The widest column shrinks; the narrow one keeps its content.
	if !strings.Contains(term.Screen(), "npm   ~/.npm/") {
		t.Errorf("narrow column was cut:\n%s", term.Screen())
	}
}

func TestTableBoxedFitsWidth(t *testing.T) {
	setup(t)

	term := vt.New(20, 10)
	_ = pterm.Table.WithWriter(term).WithBoxed().WithData(pterm.TableData{
		{"Name", "Description"},
		{"x", "a description that is far too long"},
	}).Render()

	for _, line := range strings.Split(term.Screen(), "\n") {
		if w := len([]rune(line)); w != 20 {
			t.Errorf("line %q is %d columns wide, want 20", line, w)
		}
	}
}

func TestTableUnknownWidthNeverCuts(t *testing.T) {
	setup(t)
	t.Setenv("COLUMNS", "")

	var buf bytes.Buffer

	long := strings.Repeat("x", 300)
	_ = pterm.Table.WithWriter(&buf).WithData(pterm.TableData{{"A"}, {long}}).Render()

	if !strings.Contains(buf.String(), long) {
		t.Error("cell was cut with an unknown width")
	}
}

func TestTableANSIInCells(t *testing.T) {
	setup(t)

	red := pterm.ToneError.Sprint("12")
	got, _ := pterm.Table.WithHeaderUppercase(false).WithData(pterm.TableData{
		{pterm.ToneAccent.Sprint("head"), "n"},
		{red, "1"},
		{"345", "2"},
	}).Srender()

	if !strings.Contains(got, red) {
		t.Errorf("styled cell lost its style: %q", got)
	}

	if want := "head  N"; strings.Split(plain(got), "\n")[0] == want {
		t.Error("header was uppercased without HeaderUppercase")
	}

	// Escape sequences do not count as width: "12" right-aligns under "345".
	if lines := strings.Split(plain(got), "\n"); lines[2] != "  12  1" {
		t.Errorf("row = %q", lines[2])
	}
}

func TestTableUppercaseKeepsEscapes(t *testing.T) {
	setup(t)

	got, _ := pterm.Table.WithData(pterm.TableData{{pterm.ToneAccent.Sprint("name")}, {"x"}}).Srender()
	if !strings.Contains(got, "NAME") || !strings.Contains(got, "\x1b[38;2;") {
		t.Errorf("header = %q", got)
	}
}

func TestTableMinWidthsAndRaggedRows(t *testing.T) {
	setup(t)

	got, _ := pterm.Table.WithHeaderRule(false).WithColumnMinWidths(6).WithData(pterm.TableData{
		{"a", "b", "c"},
		{"x"},
	}).Srender()

	if want := "A       B  C\nx"; plain(got) != want {
		t.Errorf("got %q, want %q", plain(got), want)
	}
}

func TestTableCSV(t *testing.T) {
	setup(t)

	p := pterm.Table.WithCSVReader(csv.NewReader(strings.NewReader("a,b\n1,2")))
	if len(p.Data) != 2 || p.Data[1][1] != "2" {
		t.Errorf("data = %v", p.Data)
	}

	broken := pterm.Table.WithData(cacheData).WithCSVReader(csv.NewReader(strings.NewReader("a,\"b")))
	if len(broken.Data) != len(cacheData) {
		t.Error("a failing reader replaced the data")
	}
}

func TestTableBuildersDoNotShare(t *testing.T) {
	setup(t)

	base := pterm.Table.WithColumnStyle(0, pterm.ToneAccent.Style())
	a := base.WithColumnStyle(1, pterm.ToneError.Style())
	b := base.WithColumnStyle(1, pterm.ToneSuccess.Style())

	if a.ColumnStyles[1] == b.ColumnStyles[1] {
		t.Error("builders share their column styles")
	}

	if len(pterm.Table.ColumnStyles) != 0 {
		t.Error("the default table was changed")
	}
}

// plain strips colors from s.
func plain(s string) string {
	return pterm.RemoveColorFromString(s)
}
