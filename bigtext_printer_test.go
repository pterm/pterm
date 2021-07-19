package pterm_test

import (
	"io"
	"strings"
	"testing"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

func TestBigTextPrinterNilPrint(t *testing.T) {
	p := pterm.BigTextPrinter{}
	p.Render()
}

func TestBigTextPrinter_Render(t *testing.T) {
	testDoesOutput(t, func(w io.Writer) {
		pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("Hello")).Render()
	})
}

func TestBigTextPrinter_RenderRawOutput(t *testing.T) {
	pterm.DisableStyling()
	testDoesOutput(t, func(w io.Writer) {
		pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("Hello")).Render()
	})
	pterm.EnableStyling()
}

func TestBigTextPrinter_WithBigCharacters(t *testing.T) {
	e := map[string]string{"a": "b", "c": "d"}
	p := pterm.BigTextPrinter{}
	p2 := p.WithBigCharacters(e)

	assert.Equal(t, e, p2.BigCharacters)
	assert.Empty(t, p.BigCharacters)
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

	assert.Equal(t, e, p2.Letters)
	assert.Empty(t, p.Letters)
}

func TestLetter_WithString(t *testing.T) {
	e := "Hello, World!"
	p := pterm.Letter{}
	p2 := p.WithString(e)

	assert.Equal(t, e, p2.String)
	assert.Empty(t, p.String)
}

func TestLetter_WithStyle(t *testing.T) {
	p := pterm.Letter{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgRed, pterm.Bold)
	p2 := p.WithStyle(s)

	assert.Equal(t, s, p2.Style)
	assert.Empty(t, p.Style)
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

	assert.Equal(t, e, p)
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

	assert.Equal(t, e, p)
}

func TestDefaultLettersMaxHeight(t *testing.T) {
	maxHeight := 5
	chars := pterm.DefaultBigText.BigCharacters
	for s, l := range chars {
		h := strings.Count(l, "\n")
		assert.LessOrEqualf(t, h, maxHeight, "'%s' is too high", s)
	}
}
