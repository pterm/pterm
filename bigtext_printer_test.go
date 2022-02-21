package pterm_test

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/MarvinJWendt/testza"

	"github.com/pterm/pterm"
)

func TestBigTextPrinterNilPrint(t *testing.T) {
	p := pterm.BigTextPrinter{}
	p.Render()
}

func TestBigTextPrinter_Render(t *testing.T) {
	printer := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("Hello"))
	content := captureStdout(func(w io.Writer) {
		printer.Render()
	})
	testza.AssertNotZero(t, content)
	testza.SnapshotCreateOrValidate(t, t.Name(), content)
	// DisableStyling
	pterm.DisableStyling()
	content = captureStdout(func(w io.Writer) {
		printer.Render()
	})
	testza.SnapshotCreateOrValidate(t, t.Name()+"NoStyling", content)
	pterm.EnableStyling()
}

func TestBigTextPrinter_RenderRawOutput(t *testing.T) {
	printer := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("Hello"))
	content := captureStdout(func(w io.Writer) {
		printer.Render()
	})
	testza.AssertNotZero(t, content)
	testza.SnapshotCreateOrValidate(t, t.Name(), content)
	// DisableStyling
	pterm.DisableStyling()
	content = captureStdout(func(w io.Writer) {
		printer.Render()
	})
	testza.SnapshotCreateOrValidate(t, t.Name()+"NoStyling", content)
	pterm.EnableStyling()
}

func TestBigTextPrinter_WithBigCharacters(t *testing.T) {
	e := map[string]string{"a": "b", "c": "d"}
	p := pterm.BigTextPrinter{}
	p2 := p.WithBigCharacters(e)

	testza.AssertEqual(t, e, p2.BigCharacters)
	testza.AssertZero(t, p.BigCharacters)
}

func TestBigTextPrinter_WithLetters(t *testing.T) {
	e := pterm.Letters{
		pterm.Letter{
			String: "test",
			Style:  pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Letter{
			String: "test2",
			Style:  pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
	}
	p := pterm.BigTextPrinter{}
	p2 := p.WithLetters(e)

	testza.AssertEqual(t, e, p2.Letters)
	testza.AssertZero(t, p.Letters)
}

func TestLetter_WithString(t *testing.T) {
	e := "Hello, World!"
	p := pterm.Letter{}
	p2 := p.WithString(e)

	testza.AssertEqual(t, e, p2.String)
	testza.AssertZero(t, p.String)
}

func TestLetter_WithStyle(t *testing.T) {
	p := pterm.Letter{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgRed, pterm.Bold)
	p2 := p.WithStyle(s)

	testza.AssertEqual(t, s, p2.Style)
	testza.AssertZero(t, p.Style)
}

func TestNewLettersFromText(t *testing.T) {
	e := pterm.Letters{
		pterm.Letter{
			String: "a",
			Style:  &pterm.ThemeDefault.LetterStyle,
		},
		pterm.Letter{
			String: "b",
			Style:  &pterm.ThemeDefault.LetterStyle,
		},
	}
	p := pterm.NewLettersFromString("ab")

	testza.AssertEqual(t, e, p)
}

func TestNewLettersFromTextWithStyle(t *testing.T) {
	e := pterm.Letters{
		pterm.Letter{
			String: "a",
			Style:  pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Letter{
			String: "b",
			Style:  pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
	}
	p := pterm.NewLettersFromStringWithStyle("ab", pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold))

	testza.AssertEqual(t, e, p)
}

func TestDefaultLettersMaxHeight(t *testing.T) {
	maxHeight := 5
	chars := pterm.DefaultBigText.BigCharacters
	for s, l := range chars {
		h := strings.Count(l, "\n")
		testza.AssertTrue(t, h <= maxHeight, fmt.Sprintf("'%s' is too high", s))
	}
}
