package pterm_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/MarvinJWendt/testza"

	"github.com/pterm/pterm"
)

func TestBigTextPrinterNilPrint(t *testing.T) {
	p := pterm.BigTextPrinter{}
	p.Render()
}

func TestBigTextPrinter_WithMethods(t *testing.T) {
	testWithMethods(t, pterm.BigTextPrinter{})
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
