package pterm_test

import (
	"context"
	"errors"
	"slices"
	"strings"
	"testing"

	"atomicgo.dev/keyboard/keys"

	"github.com/pterm/pterm"
)

var caches = []pterm.Option{
	{Group: "Package managers", Label: "npm cache", Description: "~/.npm/_cacache", Detail: "1.2 GB", Selected: true},
	{Group: "Package managers", Label: "yarn cache", Description: "~/Library/Caches/Yarn", Detail: "640 MB"},
	{Group: "Package managers", Label: "go module cache", Value: "gomod", Description: "~/go/pkg/mod", Detail: "3.4 GB"},
	{Group: "Build caches", Label: "go build cache", Description: "~/Library/Caches/go-build", Detail: "5.8 GB", Selected: true},
	{Group: "Build caches", Label: "Bazel output base", Description: "in use by a running build", Detail: "7.2 GB", Disabled: true},
	{Group: "Containers", Label: "Docker images", Description: "dangling and unused images", Detail: "18.3 GB"},
	{Group: "Containers", Label: "Docker volumes", Description: "may hold databases", Detail: "4.7 GB"},
}

func TestMultiselectKeys(t *testing.T) {
	tests := []struct {
		name  string
		p     *pterm.MultiselectPrinter
		input string
		want  []string
	}{
		{"preselected", pterm.Multiselect.WithItems(caches), keyEnter, []string{"npm cache", "go build cache"}},
		{"toggle", pterm.Multiselect.WithItems(caches), " " + keyDown + " " + keyEnter, []string{"yarn cache", "go build cache"}},
		{"value", pterm.Multiselect.WithItems(caches), keyDown + keyDown + " " + keyEnter, []string{"npm cache", "gomod", "go build cache"}},
		{"all skips disabled", pterm.Multiselect.WithItems(caches), "a" + keyEnter, []string{"npm cache", "yarn cache", "gomod", "go build cache", "Docker images", "Docker volumes"}},
		{"none", pterm.Multiselect.WithItems(caches), "n" + keyEnter, []string{}},
		{"right and left", pterm.Multiselect.WithItems(caches), keyRight + keyLeft + keyEnter, []string{}},
		{"disabled cannot be toggled", pterm.Multiselect.WithItems(caches), strings.Repeat(keyDown, 4) + " " + keyEnter, []string{"npm cache", "go build cache"}},
		{"disabled keeps its state", pterm.Multiselect.WithItems([]pterm.Option{{Label: "core", Disabled: true, Selected: true}, {Label: "extra"}}), "n" + keyEnter, []string{"core"}},
		{"filter then all", pterm.Multiselect.WithItems(caches), "n/docker" + keyRight + keyEnter, []string{"Docker images", "Docker volumes"}},
		{"letters filter while filtering", pterm.Multiselect.WithItems(caches), "/an" + " " + keyEnter, []string{"npm cache", "yarn cache", "go build cache"}},
		{"backspace closes the filter", pterm.Multiselect.WithItems(caches), "/d" + keyBackspace + keyBackspace + "n" + keyEnter, []string{}},
		{"filter off", pterm.Multiselect.WithOptions(languages).WithFilter(false), "/" + " " + keyEnter, []string{"Go"}},
		{"default options", pterm.Multiselect.WithOptions(languages).WithDefaultOptions([]string{"Zig", "Go"}), keyEnter, []string{"Go", "Zig"}},
		{"custom keys", pterm.Multiselect.WithOptions(languages).WithKeySelect(keys.Enter).WithKeyConfirm(keys.Tab), keyDown + keyEnter + "\t", []string{"Rust"}},
		{"custom all key", pterm.Multiselect.WithOptions(languages[:2]).WithKeyAll('x'), "a" + "x" + keyEnter, []string{"Go", "Rust"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(t)

			got, err := tt.p.WithInput(strings.NewReader(tt.input)).Show()
			if err != nil {
				t.Fatal(err)
			}

			if !slices.Equal(got, tt.want) {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestMultiselectSnapshots(t *testing.T) {
	view := func(p *pterm.MultiselectPrinter, input string) func() string {
		return func() string {
			return waiting(t, func() {
				_, _ = p.WithInput(strings.NewReader(input)).Show("What should be cleaned?")
			})
		}
	}

	snap(t, "Default", view(pterm.Multiselect.WithOptions(languages[:4]), keyDown+" "+keyEnter))
	snap(t, "Groups", view(pterm.Multiselect.WithItems(caches), keyDown+keyEnter))
	snap(t, "Filter", view(pterm.Multiselect.WithItems(caches), "/gocache"+keyEnter))
	snap(t, "FilterEmpty", view(pterm.Multiselect.WithItems(caches), "/"+keyEnter))
	snap(t, "Scrolled", view(pterm.Multiselect.WithItems(caches).WithMaxHeight(3), strings.Repeat(keyDown, 4)+keyEnter))
	snap(t, "Customized", view(pterm.Multiselect.
		WithOptions(languages[:3]).
		WithCheckmark("[x]", "[ ]").
		WithPointer(">").
		WithFilter(false).
		WithDefaultOptions([]string{"Rust"}), keyEnter))
}

func TestMultiselectASCII(t *testing.T) {
	setup(t)
	pterm.SetGlyphs(pterm.GlyphsASCII)

	t.Cleanup(func() { pterm.SetGlyphs(pterm.GlyphsUnicode) })

	view := stripANSI(waiting(t, func() {
		_, _ = pterm.Multiselect.WithItems(caches).WithMaxHeight(3).WithInput(strings.NewReader(keyDown + keyEnter)).Show("Clean?")
	}))

	for _, r := range view {
		if r > 127 {
			t.Fatalf("view has the non-ASCII rune %q:\n%s", r, view)
		}
	}
}

func TestMultiselectCollapses(t *testing.T) {
	term := setup(t)

	_, _ = pterm.Multiselect.WithItems(caches).WithInput(strings.NewReader(keyEnter)).Show("What should be cleaned?")

	if got, want := term.Screen(), "✔ What should be cleaned? 2 of 7"; got != want {
		t.Errorf("screen = %q, want %q", got, want)
	}
}

func TestMultiselectShowSelectedOptions(t *testing.T) {
	term := setup(t)

	_, _ = pterm.Multiselect.WithItems(caches).
		WithShowSelectedOptions().
		WithInput(strings.NewReader("a" + keyEnter)).
		Show("What should be cleaned?")

	want := "✔ What should be cleaned? 6 of 7\n" +
		"  npm cache · yarn cache · go module cache · go build cache · Docker images\n" +
		"  Docker volumes"
	if got := term.Screen(); got != want {
		t.Errorf("screen:\n%s\nwant:\n%s", got, want)
	}
}

func TestMultiselectNotInteractive(t *testing.T) {
	setup(t)

	restore := pterm.SetStdinIsTerminal(false)
	defer restore()

	got, err := pterm.Multiselect.WithItems(caches).Show()
	if err != nil || !slices.Equal(got, []string{"npm cache", "go build cache"}) {
		t.Errorf("got %q, %v; want the preselected options", got, err)
	}

	if _, err := pterm.Multiselect.WithOptions(languages).Show(); !errors.Is(err, pterm.ErrNotInteractive) {
		t.Errorf("err = %v, want ErrNotInteractive", err)
	}
}

func TestMultiselectErrors(t *testing.T) {
	term := setup(t)

	if _, err := pterm.Multiselect.Show(); !errors.Is(err, pterm.ErrEmpty) {
		t.Errorf("err = %v, want ErrEmpty", err)
	}

	_, err := pterm.Multiselect.WithOptions(languages).
		WithInput(strings.NewReader(" " + keyCtrlC)).
		WithOnInterrupt(func() {}).
		Show("Languages?")
	if !errors.Is(err, pterm.ErrInterrupted) {
		t.Errorf("err = %v, want ErrInterrupted", err)
	}

	if got, want := term.Screen(), "✖ Languages? canceled"; got != want {
		t.Errorf("screen = %q, want %q", got, want)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := pterm.Multiselect.WithOptions(languages).WithContext(ctx).WithInput(strings.NewReader(keyEnter)).Show(); !errors.Is(err, pterm.ErrCanceled) {
		t.Errorf("err = %v, want ErrCanceled", err)
	}
}

func TestMultiselectTerminalKeys(t *testing.T) {
	setup(t)

	fromTerminal(t, typed("/"), typed("docker"), key(keys.Escape), key(keys.End), key(keys.Space), key(keys.Enter))

	got, err := pterm.Multiselect.WithItems(caches).Show()
	if err != nil || !slices.Equal(got, []string{"npm cache", "go build cache", "Docker volumes"}) {
		t.Errorf("got %q, %v", got, err)
	}
}
