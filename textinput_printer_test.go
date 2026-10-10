package pterm_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"atomicgo.dev/keyboard/keys"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal/vt"
)

func TestTextInputKeys(t *testing.T) {
	tests := []struct {
		name  string
		p     *pterm.TextInputPrinter
		input string
		want  string
	}{
		{"typing", &pterm.TextInput, "hello world" + keyEnter, "hello world"},
		{"backspace", &pterm.TextInput, "helloo" + keyBackspace + keyEnter, "hello"},
		{"insert in the middle", &pterm.TextInput, "hllo" + keyLeft + keyLeft + keyLeft + "e" + keyEnter, "hello"},
		{"delete", &pterm.TextInput, "hello" + "\x1b[1~" + "\x1b[3~" + keyEnter, "ello"},
		{"home and end", &pterm.TextInput, "ello" + "\x01" + "h" + "\x05" + "!" + keyEnter, "hello!"},
		{"ctrl+w", &pterm.TextInput, "hello big world" + "\x17" + keyEnter, "hello big "},
		{"ctrl+u", &pterm.TextInput, "hello world" + keyLeft + "\x15" + keyEnter, "d"},
		{"ctrl+k", &pterm.TextInput, "hello world" + "\x01" + keyRight + "\x0b" + keyEnter, "h"},
		{"right stops at the end", &pterm.TextInput, "ab" + keyRight + keyRight + "c" + keyEnter, "abc"},
		{"unicode", &pterm.TextInput, "grüße 世界" + keyBackspace + keyEnter, "grüße 世"},
		{"default on empty", pterm.TextInput.WithDefaultValue("my-app"), keyEnter, "my-app"},
		{"typed beats default", pterm.TextInput.WithDefaultValue("my-app"), "web" + keyEnter, "web"},
		{"tab takes the default", pterm.TextInput.WithDefaultValue("my-app"), "\t-2" + keyEnter, "my-app-2"},
		{"ctrl+d submits", &pterm.TextInput, "done" + "\x04", "done"},
		{"multi-line", pterm.TextInput.WithMultiLine(), "one" + keyEnter + "two" + "\x04", "one\ntwo"},
		{"multi-line newline", pterm.TextInput.WithMultiLine(), "one\ntwo\x04", "one\ntwo"},
		{"multi-line join", pterm.TextInput.WithMultiLine(), "one" + keyEnter + keyBackspace + "two" + "\x04", "onetwo"},
		{"multi-line up", pterm.TextInput.WithMultiLine(), "one" + keyEnter + "two" + keyUp + "!" + "\x04", "one!\ntwo"},
		{"multi-line delete joins", pterm.TextInput.WithMultiLine(), "one" + keyEnter + "two" + keyUp + "\x1b[3~" + "\x04", "onetwo"},
		{"custom submit key", pterm.TextInput.WithMultiLine().WithSubmitKey(keys.Tab), "a" + keyEnter + "b\t", "a\nb"},
		{"mask", pterm.TextInput.WithMask("*"), "hunter2" + keyEnter, "hunter2"},
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

// The terminal reader delivers pasted text as one key per run of
// printable runes, with enter keys between the lines. A paste must land
// once, at the cursor, and the cursor must end up after it.
func TestTextInputPaste(t *testing.T) {
	tests := []struct {
		name string
		p    *pterm.TextInputPrinter
		keys []keys.Key
		want string
	}{
		{"single line stops at the line break", &pterm.TextInput, []keys.Key{typed("hello world"), key(keys.Enter), typed("ignored")}, "hello world"},
		{"multi-line keeps every line", pterm.TextInput.WithMultiLine(), []keys.Key{typed("hello world"), key(keys.Enter), typed("second line"), key(keys.CtrlD)}, "hello world\nsecond line"},
		{"crlf is one line break", pterm.TextInput.WithMultiLine(), []keys.Key{typed("a"), key(keys.Enter), key(keys.CtrlJ), typed("b"), key(keys.CtrlD)}, "a\nb"},
		{"paste at the cursor", &pterm.TextInput, []keys.Key{typed("hd"), key(keys.Left), typed("ello worl"), typed("!"), key(keys.Enter)}, "hello worl!d"},
		{"paste twice", &pterm.TextInput, []keys.Key{typed("ab"), typed("ab"), key(keys.Enter)}, "abab"},
		{"tabs become spaces", &pterm.TextInput, []keys.Key{typed("a\tb\x1b"), key(keys.Enter)}, "a b"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(t)
			fromTerminal(t, tt.keys...)

			got, err := tt.p.Show()
			if err != nil {
				t.Fatal(err)
			}

			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestTextInputSnapshots(t *testing.T) {
	view := func(p *pterm.TextInputPrinter, input string) func() string {
		return func() string {
			return waiting(t, func() {
				_, _ = p.WithInput(strings.NewReader(input)).Show("Project name")
			})
		}
	}

	snap(t, "Empty", view(&pterm.TextInput, keyEnter))
	snap(t, "Typed", view(&pterm.TextInput, "my-app"+keyLeft+keyLeft+keyEnter))
	snap(t, "Default", view(pterm.TextInput.WithDefaultValue("my-app"), keyEnter))
	snap(t, "Placeholder", view(pterm.TextInput.WithPlaceholder("lowercase, no spaces"), keyEnter))
	snap(t, "Mask", view(pterm.TextInput.WithMask("*"), "hunter2"+keyEnter))
	snap(t, "MultiLine", view(pterm.TextInput.WithMultiLine(), "Fix the renderer"+keyEnter+keyEnter+"It erased one row too many."+keyUp+"\x04"))
	snap(t, "Wrapped", view(pterm.TextInput.WithMultiLine(), strings.Repeat("The region erased one row too many. ", 4)+"\x04"))
}

func TestTextInputWraps(t *testing.T) {
	setup(t)

	term := vt.New(30, 24)
	input := strings.Repeat("abc def ", 10)

	view := waiting(t, func() {
		_, _ = pterm.TextInput.WithWriter(term).WithInput(strings.NewReader(input + keyEnter)).Show("Name")
	})

	plain := stripANSI(view)

	lines := strings.Split(plain, "\n")
	if len(lines) < 3 {
		t.Fatalf("expected the input to wrap:\n%s", plain)
	}

	for _, line := range lines {
		if w := len([]rune(line)); w > 29 {
			t.Errorf("line %q is %d columns wide, want at most 29", line, w)
		}
	}

	if got := strings.Join(strings.Fields(strings.TrimPrefix(plain, "? Name")), " "); got != strings.TrimSpace(input) {
		t.Errorf("wrapped text lost something: %q", got)
	}
}

func TestTextInputCollapses(t *testing.T) {
	tests := []struct {
		name  string
		p     *pterm.TextInputPrinter
		input string
		want  string
	}{
		{"text", &pterm.TextInput, "my-app" + keyEnter, "✔ Project name my-app"},
		{"default", pterm.TextInput.WithDefaultValue("my-app"), keyEnter, "✔ Project name my-app"},
		{"mask hides the length", pterm.TextInput.WithMask("*"), "abc" + keyEnter, "✔ Project name ********"},
		{"multi-line", pterm.TextInput.WithMultiLine(), "one\rtwo\rthree\x04", "✔ Project name one +2 lines"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			term := setup(t)

			_, _ = tt.p.WithInput(strings.NewReader(tt.input)).Show("Project name")

			if got := term.Screen(); got != tt.want {
				t.Errorf("screen = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestTextInputNotInteractive(t *testing.T) {
	setup(t)

	restore := pterm.SetStdinIsTerminal(false)
	defer restore()

	got, err := pterm.TextInput.WithDefaultValue("my-app").Show()
	if err != nil || got != "my-app" {
		t.Errorf("got %q, %v; want the default value", got, err)
	}

	if _, err := pterm.TextInput.Show(); !errors.Is(err, pterm.ErrNotInteractive) {
		t.Errorf("err = %v, want ErrNotInteractive", err)
	}
}

func TestTextInputErrors(t *testing.T) {
	term := setup(t)

	_, err := pterm.TextInput.
		WithInput(strings.NewReader("abc" + keyCtrlC)).
		WithOnInterrupt(func() {}).
		Show("Name")
	if !errors.Is(err, pterm.ErrInterrupted) {
		t.Errorf("err = %v, want ErrInterrupted", err)
	}

	if got, want := term.Screen(), "✖ Name canceled"; got != want {
		t.Errorf("screen = %q, want %q", got, want)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := pterm.TextInput.WithContext(ctx).WithInput(strings.NewReader("a")).Show(); !errors.Is(err, pterm.ErrCanceled) {
		t.Errorf("err = %v, want ErrCanceled", err)
	}
}
