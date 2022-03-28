package pterm_test

import (
	"fmt"
	"os"
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
	content, err := printer.Srender()

	testza.AssertNoError(t, err)
	testza.AssertNotZero(t, content)
}

func TestBigTextPrinter_RenderRGB(t *testing.T) {
	printer := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromStringWithRGB("Hello", pterm.NewRGB(255, 0, 0)))
	content, err := printer.Srender()

	testza.AssertNoError(t, err)
	testza.AssertNotZero(t, content)
}

func TestBigTextPrinter_RenderRawOutput(t *testing.T) {
	printer := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("Hello"))

	pterm.DisableStyling()
	content, err := printer.Srender()
	pterm.EnableStyling()

	testza.AssertNoError(t, err)
	testza.AssertNotZero(t, content)
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

func TestLetter_WithRGB(t *testing.T) {
	p := pterm.Letter{}
	rgb := pterm.NewRGB(0, 0, 0)
	p2 := p.WithRGB(rgb)

	testza.AssertEqual(t, rgb, p2.RGB)
	testza.AssertZero(t, p.RGB)
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

func TestNewLettersFromTextWithRGB(t *testing.T) {
	e := pterm.Letters{
		pterm.Letter{
			String: "a",
			Style:  pterm.NewStyle(),
			RGB:    pterm.NewRGB(0, 0, 0),
		},
		pterm.Letter{
			String: "b",
			Style:  pterm.NewStyle(),
			RGB:    pterm.NewRGB(0, 0, 0),
		},
	}
	p := pterm.NewLettersFromStringWithRGB("ab", pterm.NewRGB(0, 0, 0))

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

func TestBigTextPrinter_WithWriter(t *testing.T) {
	p := pterm.BigTextPrinter{}
	s := os.Stderr
	p2 := p.WithWriter(s)

	testza.AssertEqual(t, s, p2.Writer)
	testza.AssertZero(t, p.Writer)
}
