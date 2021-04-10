package pterm

import (
	"github.com/pterm/pterm/internal"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
	"strings"
)

func TestBigTextPrinterNilPrint(t *testing.T) {
	p := BigTextPrinter{}
	p.Render()
}

func TestBigTextPrinter_Render(t *testing.T) {
	internal.TestDoesOutput(t, func(w io.Writer) {
		DefaultBigText.WithLetters(NewLettersFromString("Hello")).Render()
	})
}

func TestBigTextPrinter_WithBigCharacters(t *testing.T) {
	e := map[string]string{"a": "b", "c": "d"}
	p := BigTextPrinter{}
	p2 := p.WithBigCharacters(e)

	assert.Equal(t, e, p2.BigCharacters)
	assert.Empty(t, p.BigCharacters)
}

func TestBigTextPrinter_WithLetters(t *testing.T) {
	e := Letters{
		Letter{
			String: "test",
			Style:  NewStyle(FgRed, BgBlue, Bold),
		},
		Letter{
			String: "test2",
			Style:  NewStyle(FgRed, BgBlue, Bold),
		},
	}
	p := BigTextPrinter{}
	p2 := p.WithLetters(e)

	assert.Equal(t, e, p2.Letters)
	assert.Empty(t, p.Letters)
}

func TestLetter_WithString(t *testing.T) {
	e := "Hello, World!"
	p := Letter{}
	p2 := p.WithString(e)

	assert.Equal(t, e, p2.String)
	assert.Empty(t, p.String)
}

func TestLetter_WithStyle(t *testing.T) {
	p := Letter{}
	s := NewStyle(FgRed, BgRed, Bold)
	p2 := p.WithStyle(s)

	assert.Equal(t, s, p2.Style)
	assert.Empty(t, p.Style)
}

func TestNewLettersFromText(t *testing.T) {
	e := Letters{
		Letter{
			String: "a",
			Style:  &ThemeDefault.LetterStyle,
		},
		Letter{
			String: "b",
			Style:  &ThemeDefault.LetterStyle,
		},
	}
	p := NewLettersFromString("ab")

	assert.Equal(t, e, p)
}

func TestNewLettersFromTextWithStyle(t *testing.T) {
	e := Letters{
		Letter{
			String: "a",
			Style:  NewStyle(FgRed, BgBlue, Bold),
		},
		Letter{
			String: "b",
			Style:  NewStyle(FgRed, BgBlue, Bold),
		},
	}
	p := NewLettersFromStringWithStyle("ab", NewStyle(FgRed, BgBlue, Bold))

	assert.Equal(t, e, p)
}

func TestDefaultLettersMaxHeight(t *testing.T) {
	maxHeight := 5
	chars := DefaultBigText.BigCharacters
	for s, l := range(chars) {
		h := strings.Count(l, "\n")
		assert.LessOrEqualf(t, h, maxHeight, "'%s' is too high", s)
	}
}
