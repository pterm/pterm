package pterm_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/pterm/pterm/internal/testhelper"

	"github.com/pterm/pterm"
)

func TestBigTextPrinterNilPrint(t *testing.T) {
	p := pterm.BigTextPrinter{}
	p.Render()
}

func TestBigTextPrinter_Render(t *testing.T) {
	printer := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("Hello"))
	content, err := printer.Srender()

	testhelper.AssertNoError(t, err)
	testhelper.AssertNotZero(t, content)
}

func TestBigTextPrinter_RenderRGB(t *testing.T) {
	printer := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromStringWithRGB("Hello", pterm.NewRGB(255, 0, 0)))
	content, err := printer.Srender()

	testhelper.AssertNoError(t, err)
	testhelper.AssertNotZero(t, content)
}

func TestBigTextPrinter_RenderRawOutput(t *testing.T) {
	printer := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("Hello"))

	pterm.DisableStyling()

	content, err := printer.Srender()

	pterm.EnableStyling()

	testhelper.AssertNoError(t, err)
	testhelper.AssertNotZero(t, content)
}

func TestBigTextPrinter_WithBigCharacters(t *testing.T) {
	e := map[string]string{"a": "b", "c": "d"}
	p := pterm.BigTextPrinter{}
	p2 := p.WithBigCharacters(e)

	testhelper.AssertEqual(t, e, p2.BigCharacters)
	testhelper.AssertZero(t, p.BigCharacters)
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

	testhelper.AssertEqual(t, e, p2.Letters)
	testhelper.AssertZero(t, p.Letters)
}

func TestLetter_WithString(t *testing.T) {
	e := "Hello, World!"
	p := pterm.Letter{}
	p2 := p.WithString(e)

	testhelper.AssertEqual(t, e, p2.String)
	testhelper.AssertZero(t, p.String)
}

func TestLetter_WithStyle(t *testing.T) {
	p := pterm.Letter{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgRed, pterm.Bold)
	p2 := p.WithStyle(s)

	testhelper.AssertEqual(t, s, p2.Style)
	testhelper.AssertZero(t, p.Style)
}

func TestLetter_WithRGB(t *testing.T) {
	p := pterm.Letter{}
	rgb := pterm.NewRGB(0, 0, 0)
	p2 := p.WithRGB(rgb)

	testhelper.AssertEqual(t, rgb, p2.RGB)
	testhelper.AssertZero(t, p.RGB)
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

	testhelper.AssertEqual(t, e, p)
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

	testhelper.AssertEqual(t, e, p)
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

	testhelper.AssertEqual(t, e, p)
}

func TestDefaultLettersMaxHeight(t *testing.T) {
	maxHeight := 5

	chars := pterm.DefaultBigText.BigCharacters
	for s, l := range chars {
		h := strings.Count(l, "\n")
		testhelper.AssertTrue(t, h <= maxHeight, fmt.Sprintf("'%s' is too high", s))
	}
}

func TestBigTextPrinter_WithWriter(t *testing.T) {
	p := pterm.BigTextPrinter{}
	s := os.Stderr
	p2 := p.WithWriter(s)

	testhelper.AssertEqual(t, s, p2.Writer)
	testhelper.AssertZero(t, p.Writer)
}
