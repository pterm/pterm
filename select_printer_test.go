package pterm_test

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"

	"atomicgo.dev/keyboard/keys"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal/vt"
)

var clusters = []pterm.Option{
	{Label: "kapp-k8s-at", Description: "Vienna", Detail: "3 nodes"},
	{Label: "kapp-k8s-de", Description: "Frankfurt", Detail: "5 nodes"},
	{Label: "kapp-k8s-us", Description: "Virginia", Detail: "12 nodes"},
	{Label: "staging", Value: "stg", Description: "shared playground", Detail: "2 nodes"},
	{Label: "legacy-eu", Description: "read only until March", Detail: "1 node", Disabled: true},
}

var languages = []string{"Go", "Rust", "TypeScript", "Python", "Zig", "Kotlin", "Swift", "Elixir", "Haskell", "OCaml"}

func TestSelectKeys(t *testing.T) {
	tests := []struct {
		name  string
		p     *pterm.SelectPrinter
		input string
		want  string
	}{
		{"enter takes the first", pterm.Select.WithItems(clusters), keyEnter, "kapp-k8s-at"},
		{"down", pterm.Select.WithItems(clusters), keyDown + keyEnter, "kapp-k8s-de"},
		{"up wraps and skips disabled", pterm.Select.WithItems(clusters), keyUp + keyEnter, "stg"},
		{"down wraps and skips disabled", pterm.Select.WithItems(clusters), strings.Repeat(keyDown, 4) + keyEnter, "kapp-k8s-at"},
		{"value", pterm.Select.WithItems(clusters), strings.Repeat(keyDown, 3) + keyEnter, "stg"},
		{"default option", pterm.Select.WithItems(clusters).WithDefaultOption("kapp-k8s-us"), keyEnter, "kapp-k8s-us"},
		{"default by value", pterm.Select.WithItems(clusters).WithDefaultOption("stg"), keyUp + keyEnter, "kapp-k8s-us"},
		{"filter", pterm.Select.WithItems(clusters), "k8sus" + keyEnter, "kapp-k8s-us"},
		{"filter backspace", pterm.Select.WithItems(clusters), "stagx" + keyBackspace + keyEnter, "stg"},
		{"no matches ignores enter", pterm.Select.WithOptions(languages), "zzz" + keyEnter + strings.Repeat(keyBackspace, 3) + keyEnter, "Go"},
		{"filter off ignores typing", pterm.Select.WithOptions(languages).WithFilter(false), "rust" + keyEnter, "Go"},
		{"end", pterm.Select.WithOptions(languages), "\x1b[4~" + keyEnter, "OCaml"},
		{"page down", pterm.Select.WithOptions(languages).WithMaxHeight(3), "\x1b[6~" + keyEnter, "Python"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(t)

			got, err := tt.p.WithInput(strings.NewReader(tt.input)).Show()
			if err != nil {
				t.Fatal(err)
			}

			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSelectSnapshots(t *testing.T) {
	view := func(p *pterm.SelectPrinter, input string) func() string {
		return func() string {
			return waiting(t, func() {
				_, _ = p.WithInput(strings.NewReader(input)).Show("Which cluster?")
			})
		}
	}

	snap(t, "Default", view(pterm.Select.WithOptions(languages[:4]), keyDown+keyEnter))
	snap(t, "Columns", view(pterm.Select.WithItems(clusters), keyDown+keyEnter))
	snap(t, "Filter", view(pterm.Select.WithItems(clusters), "k8s"+keyEnter))
	snap(t, "NoMatches", view(pterm.Select.WithItems(clusters), "zzz"))
	snap(t, "Scrolled", view(pterm.Select.WithOptions(languages).WithMaxHeight(4), strings.Repeat(keyDown, 5)+keyEnter))
	snap(t, "Groups", view(pterm.Select.WithItems([]pterm.Option{
		{Group: "Production", Label: "kapp-k8s-at"},
		{Group: "Production", Label: "kapp-k8s-de"},
		{Group: "Testing", Label: "staging"},
	}), keyEnter))
	snap(t, "Customized", view(pterm.Select.
		WithOptions(languages[:3]).
		WithPointer(">").
		WithKeyHints(false).
		WithFilterPlaceholder("type to search"), keyEnter))
}

func TestSelectNarrow(t *testing.T) {
	setup(t)

	term := vt.New(36, 24)

	view := waiting(t, func() {
		_, _ = pterm.Select.WithItems(clusters).WithWriter(term).WithInput(strings.NewReader(keyEnter)).Show("Which cluster?")
	})

	for line := range strings.SplitSeq(view, "\n") {
		if w := len([]rune(stripANSI(line))); w > 35 {
			t.Errorf("line %q is %d columns wide, want at most 35", stripANSI(line), w)
		}
	}

	if !strings.Contains(stripANSI(view), "…") {
		t.Errorf("expected cut descriptions in:\n%s", stripANSI(view))
	}
}

func TestSelectCollapses(t *testing.T) {
	term := setup(t)

	_, _ = pterm.Select.WithItems(clusters).WithInput(strings.NewReader(keyDown + keyEnter)).Show("Which cluster?")

	if got, want := term.Screen(), "✔ Which cluster? kapp-k8s-de"; got != want {
		t.Errorf("screen = %q, want %q", got, want)
	}
}

func TestSelectInFlow(t *testing.T) {
	term := setup(t)

	flow, _ := pterm.Flow.Start("deploy")

	var view string

	restore := pterm.SetViewHook(func(v string) { view = v })
	_, _ = pterm.Select.WithOptions(languages[:2]).WithInput(strings.NewReader(keyEnter)).Show("Language?")

	restore()

	_ = flow.Stop()

	// The question stands on the rail, the options are drawn next to it.
	lines := strings.Split(view, "\n")
	if !strings.HasPrefix(lines[0], "\x00") || strings.HasPrefix(lines[1], "\x00") {
		t.Errorf("only the question should be a node line: %q", view)
	}

	want := "┌ deploy\n│\n✔ Language? Go\n│\n└"
	if got := term.Screen(); got != want {
		t.Errorf("screen:\n%s\nwant:\n%s", got, want)
	}
}

func TestSelectNotInteractive(t *testing.T) {
	setup(t)

	restore := pterm.SetStdinIsTerminal(false)
	defer restore()

	got, err := pterm.Select.WithItems(clusters).WithDefaultOption("staging").Show()
	if err != nil || got != "stg" {
		t.Errorf("got %q, %v; want the default value", got, err)
	}

	preselected := []pterm.Option{{Label: "a"}, {Label: "b", Selected: true}}
	if got, err := pterm.Select.WithItems(preselected).Show(); err != nil || got != "b" {
		t.Errorf("got %q, %v; want the selected option", got, err)
	}

	if _, err := pterm.Select.WithItems(clusters).Show(); !errors.Is(err, pterm.ErrNotInteractive) {
		t.Errorf("err = %v, want ErrNotInteractive", err)
	}
}

func TestSelectErrors(t *testing.T) {
	term := setup(t)

	if _, err := pterm.Select.Show(); !errors.Is(err, pterm.ErrEmpty) {
		t.Errorf("err = %v, want ErrEmpty", err)
	}

	called := false

	_, err := pterm.Select.WithOptions(languages).
		WithInput(strings.NewReader(keyCtrlC)).
		WithOnInterrupt(func() { called = true }).
		Show("Language?")
	if !errors.Is(err, pterm.ErrInterrupted) || !called {
		t.Errorf("err = %v, handler called = %v", err, called)
	}

	if got, want := term.Screen(), "✖ Language? canceled"; got != want {
		t.Errorf("screen = %q, want %q", got, want)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := pterm.Select.WithOptions(languages).WithContext(ctx).WithInput(strings.NewReader(keyEnter)).Show(); !errors.Is(err, pterm.ErrCanceled) {
		t.Errorf("err = %v, want ErrCanceled", err)
	}

	if _, err := pterm.Select.WithOptions(languages).WithInput(strings.NewReader(keyDown)).Show(); !errors.Is(err, io.EOF) {
		t.Errorf("err = %v, want io.EOF when the input ends", err)
	}
}

func TestSelectTerminalKeys(t *testing.T) {
	setup(t)

	// The terminal delivers fast typing as one key with several runes.
	fromTerminal(t, typed("k8sde"), key(keys.Enter))

	got, err := pterm.Select.WithItems(clusters).Show()
	if err != nil || got != "kapp-k8s-de" {
		t.Errorf("got %q, %v; want kapp-k8s-de", got, err)
	}
}
