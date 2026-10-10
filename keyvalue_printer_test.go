package pterm_test

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal/vt"
)

var pairs = []pterm.KeyValuePair{
	pterm.KV("Version", "0.1.0"),
	pterm.KV("Path", "~/.local/fleet/dotfiles/src/cli/main.ts"),
	pterm.KV("Uptime", "1m 34s").WithSuffix("since last boot"),
}

func TestKeyValueSnapshots(t *testing.T) {
	snap(t, "Default", srender(t, pterm.KeyValue.WithPairs(pairs...).Srender))
	snap(t, "Accent", srender(t, pterm.KeyValue.WithAccentValues().WithPairs(pairs...).Srender))
	snap(t, "Wrapped", srender(t, pterm.KeyValue.WithPairs(
		pterm.KV("Description", strings.Repeat("a long value that wraps ", 6)),
		pterm.KV("Lines", "first\nsecond"),
	).Srender))
}

func TestKeyValueLook(t *testing.T) {
	setup(t)

	out := capture(t, func() { _ = pterm.KeyValue.WithPairs(pairs...).Render() })
	want := "  Version  0.1.0\n" +
		"  Path     ~/.local/fleet/dotfiles/src/cli/main.ts\n" +
		"  Uptime   1m 34s since last boot\n"

	if out != want {
		t.Errorf("output:\n%s\nwant:\n%s", out, want)
	}
}

func TestKeyValueWrapsWithHangingIndent(t *testing.T) {
	setup(t)

	term := vt.New(30, 10)
	_ = pterm.KeyValue.WithWriter(term).WithIndent(0).WithPairs(
		pterm.KV("Key", "one two three four five six seven"),
		pterm.KV("Unit", "12").WithSuffix("MB"),
		pterm.KV("Empty", ""),
	).Render()

	want := "Key    one two three four five\n       six seven\nUnit   12 MB\nEmpty"
	if term.Screen() != want {
		t.Errorf("screen:\n%s\nwant:\n%s", term.Screen(), want)
	}
}

func TestKeyValueUnknownWidth(t *testing.T) {
	setup(t)
	t.Setenv("COLUMNS", "")

	var buf bytes.Buffer

	long := strings.Repeat("word ", 50)
	_ = pterm.KeyValue.WithWriter(&buf).WithPairs(pterm.KV("k", long)).Render()

	if strings.Count(buf.String(), "\n") != 1 {
		t.Errorf("value wrapped with an unknown width: %q", buf.String())
	}
}

func TestKeyValueBreaksLongWords(t *testing.T) {
	setup(t)

	term := vt.New(20, 10)
	_ = pterm.KeyValue.WithWriter(term).WithIndent(0).WithPairs(pterm.KV("k", strings.Repeat("x", 30))).Render()

	for _, line := range strings.Split(term.Screen(), "\n") {
		if len(line) > 20 {
			t.Errorf("line %q is wider than 20", line)
		}
	}
}

func TestKVFormatsValues(t *testing.T) {
	if kv := pterm.KV("n", 42); kv.Value != "42" {
		t.Errorf("value = %q", kv.Value)
	}
}

func TestKeyValueEmpty(t *testing.T) {
	setup(t)

	if _, err := pterm.KeyValue.Srender(); !errors.Is(err, pterm.ErrEmpty) {
		t.Errorf("err = %v, want ErrEmpty", err)
	}
}
