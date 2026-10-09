package pterm_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

// foregroundCodes maps every exported foreground color to its SGR code.
var foregroundCodes = map[pterm.Color]int{
	pterm.FgBlack:        30,
	pterm.FgRed:          31,
	pterm.FgGreen:        32,
	pterm.FgYellow:       33,
	pterm.FgBlue:         34,
	pterm.FgMagenta:      35,
	pterm.FgCyan:         36,
	pterm.FgWhite:        37,
	pterm.FgDefault:      39,
	pterm.FgDarkGray:     90,
	pterm.FgLightRed:     91,
	pterm.FgLightGreen:   92,
	pterm.FgLightYellow:  93,
	pterm.FgLightBlue:    94,
	pterm.FgLightMagenta: 95,
	pterm.FgLightCyan:    96,
	pterm.FgLightWhite:   97,
}

// backgroundCodes maps every exported background color to its SGR code.
var backgroundCodes = map[pterm.Color]int{
	pterm.BgBlack:        40,
	pterm.BgRed:          41,
	pterm.BgGreen:        42,
	pterm.BgYellow:       43,
	pterm.BgBlue:         44,
	pterm.BgMagenta:      45,
	pterm.BgCyan:         46,
	pterm.BgWhite:        47,
	pterm.BgDefault:      49,
	pterm.BgDarkGray:     100,
	pterm.BgLightRed:     101,
	pterm.BgLightGreen:   102,
	pterm.BgLightYellow:  103,
	pterm.BgLightBlue:    104,
	pterm.BgLightMagenta: 105,
	pterm.BgLightCyan:    106,
	pterm.BgLightWhite:   107,
}

// optionCodes maps every exported style option to its SGR code.
var optionCodes = map[pterm.Color]int{
	pterm.Reset:         0,
	pterm.Bold:          1,
	pterm.Fuzzy:         2,
	pterm.Italic:        3,
	pterm.Underscore:    4,
	pterm.Blink:         5,
	pterm.FastBlink:     6,
	pterm.Reverse:       7,
	pterm.Concealed:     8,
	pterm.Strikethrough: 9,
}

// allColorCodes merges all exported color and option constants.
func allColorCodes() map[pterm.Color]int {
	all := make(map[pterm.Color]int, len(foregroundCodes)+len(backgroundCodes)+len(optionCodes))

	for _, m := range []map[pterm.Color]int{foregroundCodes, backgroundCodes, optionCodes} {
		for c, code := range m {
			all[c] = code
		}
	}

	return all
}

func TestColorSprintEmitsExactEscapeSequence(t *testing.T) {
	for color, code := range allColorCodes() {
		t.Run(fmt.Sprintf("code %d", code), func(t *testing.T) {
			expected := fmt.Sprintf("\x1b[%dmx\x1b[0m", code)
			assert.Equal(t, expected, color.Sprint("x"))
		})
	}
}

func TestColorString(t *testing.T) {
	for color, code := range allColorCodes() {
		assert.Equal(t, fmt.Sprintf("%d", code), color.String())
	}
}

func TestColorSprintFamilyExactOutput(t *testing.T) {
	assert.Equal(t, "\x1b[31mx\x1b[0m", pterm.FgRed.Sprint("x"))
	assert.Equal(t, "\x1b[31m1 2\x1b[0m", pterm.FgRed.Sprint(1, 2), "Sprint keeps fmt spacing rules inside the sequence")
	assert.Equal(t, "\x1b[31mn=1\x1b[0m", pterm.FgRed.Sprintf("n=%d", 1))
	assert.Equal(t, "\x1b[31mn=1\x1b[0m\n", pterm.FgRed.Sprintfln("n=%d", 1))
	assert.Equal(t, "\x1b[31mx\x1b[0m\n", pterm.FgRed.Sprintln("x"))
	assert.Equal(t, "", pterm.FgRed.Sprint(""), "empty input must not emit stray sequences")
}

func TestColorSprintWrapsEachLineIndividually(t *testing.T) {
	assert.Equal(t, "\x1b[31mone\x1b[0m\n\x1b[31mtwo\x1b[0m", pterm.FgRed.Sprint("one\ntwo"))
}

func TestColorSprintReopensColorAfterNestedReset(t *testing.T) {
	inner := pterm.FgGreen.Sprint("mid")
	expected := "\x1b[31ma\x1b[32mmid\x1b[0m\x1b[31mb\x1b[0m"

	// The inner reset is followed by the outer color again, so the text after
	// the nested segment keeps the outer color.
	assert.Equal(t, expected, pterm.FgRed.Sprint("a"+inner+"b"))
}

func TestColorPrintWritesExactlySprintToDefaultWriter(t *testing.T) {
	// Print/Sprint delegation across all TextPrinters is covered by the
	// contract tests; this locks the exact bytes for a representative color.
	out := captureStdout(func(_ io.Writer) { pterm.FgBlue.Print("x") })
	assert.Equal(t, "\x1b[34mx\x1b[0m", out)
}

func TestColorAliasesMatchTheirColor(t *testing.T) {
	aliases := map[string]struct {
		alias func(...any) string
		color pterm.Color
	}{
		"Red":     {pterm.Red, pterm.FgRed},
		"Green":   {pterm.Green, pterm.FgGreen},
		"Yellow":  {pterm.Yellow, pterm.FgYellow},
		"Blue":    {pterm.Blue, pterm.FgBlue},
		"Magenta": {pterm.Magenta, pterm.FgMagenta},
		"Cyan":    {pterm.Cyan, pterm.FgCyan},
		"Gray":    {pterm.Gray, pterm.FgGray},
		"Normal":  {pterm.Normal, pterm.FgDefault},
	}

	for name, tc := range aliases {
		assert.Equal(t, tc.color.Sprint("x"), tc.alias("x"), "alias %s", name)
	}
}

func TestFgGrayAndBgGrayAreAliases(t *testing.T) {
	assert.Equal(t, pterm.FgDarkGray, pterm.FgGray)
	assert.Equal(t, pterm.BgDarkGray, pterm.BgGray)
}

func TestColorToStyle(t *testing.T) {
	assert.Equal(t, &pterm.Style{pterm.FgCyan}, pterm.FgCyan.ToStyle())
}

// Style tests.

func TestNewStyle(t *testing.T) {
	assert.Equal(t, &pterm.Style{pterm.FgRed, pterm.BgBlue, pterm.Bold}, pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold))
}

func TestStyleCodeCombinesAllColors(t *testing.T) {
	style := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)

	assert.Equal(t, "31;44;1", style.Code())
	assert.Equal(t, "31;44;1", style.String())
	assert.Equal(t, "", pterm.Style{}.Code(), "empty style has no code")
}

func TestStyleSprintEmitsCombinedSGRSequence(t *testing.T) {
	style := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)

	assert.Equal(t, "\x1b[31;44;1mx\x1b[0m", style.Sprint("x"))
	assert.Equal(t, "\x1b[31;44;1mn=1\x1b[0m", style.Sprintf("n=%d", 1))
	assert.Equal(t, "\x1b[31;44;1mx\x1b[0m\n", style.Sprintln("x"))
	assert.Equal(t, "\x1b[31;44;1mn=1\x1b[0m\n", style.Sprintfln("n=%d", 1))
}

func TestStyleSprintWrapsEachLineIndividually(t *testing.T) {
	assert.Equal(t, "\x1b[31;1mone\x1b[0m\n\x1b[31;1mtwo\x1b[0m", pterm.NewStyle(pterm.FgRed, pterm.Bold).Sprint("one\ntwo"))
}

func TestStyleSprintReopensStyleAfterNestedReset(t *testing.T) {
	inner := pterm.FgGreen.Sprint("mid")
	expected := "\x1b[31;1ma\x1b[32mmid\x1b[0m\x1b[31;1mb\x1b[0m"

	assert.Equal(t, expected, pterm.NewStyle(pterm.FgRed, pterm.Bold).Sprint("a"+inner+"b"))
}

func TestStylePrintWritesExactlySprint(t *testing.T) {
	style := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)

	assert.Equal(t, style.Sprint("x"), captureStdout(func(_ io.Writer) { style.Print("x") }))
	assert.Equal(t, style.Sprintln("x"), captureStdout(func(_ io.Writer) { style.Println("x") }))
	assert.Equal(t, style.Sprintf("n=%d", 1), captureStdout(func(_ io.Writer) { style.Printf("n=%d", 1) }))
	assert.Equal(t, style.Sprintfln("n=%d", 1), captureStdout(func(_ io.Writer) { style.Printfln("n=%d", 1) }))
}

func TestStyleAdd(t *testing.T) {
	assert.Equal(t, pterm.Style{pterm.FgRed, pterm.BgGreen}, pterm.Style{pterm.FgRed}.Add(pterm.Style{pterm.BgGreen}))
	assert.Equal(t, pterm.Style{pterm.FgRed, pterm.BgGreen, pterm.Bold}, pterm.Style{pterm.FgRed}.Add(pterm.Style{pterm.BgGreen}).Add(pterm.Style{pterm.Bold}))
	assert.Equal(t, pterm.Style{pterm.FgRed, pterm.BgGreen, pterm.Bold}, pterm.Style{pterm.FgRed}.Add(pterm.Style{pterm.BgGreen, pterm.Bold}))
	assert.Equal(t, pterm.Style{pterm.FgRed, pterm.BgGreen, pterm.Bold}, pterm.Style{pterm.FgRed}.Add(pterm.Style{pterm.BgGreen}, pterm.Style{pterm.Bold}))
}

func TestStyleRemoveColor(t *testing.T) {
	style := pterm.Style{pterm.FgRed, pterm.BgBlue, pterm.Bold}

	assert.Equal(t, pterm.Style{pterm.FgRed, pterm.Bold}, style.RemoveColor(pterm.BgBlue))
	assert.Equal(t, pterm.Style{pterm.FgRed}, style.RemoveColor(pterm.BgBlue, pterm.Bold))
	assert.Equal(t, pterm.Style{pterm.Bold, pterm.Bold}.RemoveColor(pterm.Bold), pterm.Style{}, "duplicates are removed too")
}

// RemoveColorFromString must strip exactly what Sprint added (roundtrip).

func TestRemoveColorFromStringRoundtrip(t *testing.T) {
	inputs := []string{
		"Hello, PTerm!",
		"one\ntwo\nthree",
		"ünïcödé ✓ 汉字",
		"100% done",
	}

	// Style itself is not a TextPrinter (its Print methods return nothing),
	// so the common denominator here is Sprint.
	printers := map[string]interface{ Sprint(...any) string }{
		"Color fg":            pterm.FgRed,
		"Color bg":            pterm.BgCyan,
		"Color light":         pterm.FgLightMagenta,
		"Style fg+bg+bold":    pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		"Style single option": pterm.NewStyle(pterm.Underscore),
		"RGB fg":              pterm.NewRGB(12, 34, 56),
		"RGBStyle":            pterm.NewRGBStyle(pterm.RGB{R: 1, G: 2, B: 3}, pterm.RGB{R: 4, G: 5, B: 6}).AddOptions(pterm.Bold),
	}

	for name, printer := range printers {
		t.Run(name, func(t *testing.T) {
			for _, input := range inputs {
				assert.Equal(t, input, pterm.RemoveColorFromString(printer.Sprint(input)))
			}
		})
	}

	t.Run("nested styling", func(t *testing.T) {
		nested := pterm.FgRed.Sprint("a" + pterm.FgGreen.Sprint("mid") + "b")
		assert.Equal(t, "amidb", pterm.RemoveColorFromString(nested))
	})

	t.Run("RGB background leaves the clear-to-EOL sequence behind", func(t *testing.T) {
		// Known limitation: RGB background Sprint appends "\x1b[K" (clear to
		// end of line) so the color fills the row, but RemoveColorFromString
		// only strips SGR sequences and OSC 8 hyperlinks, not this control
		// sequence. If this assertion starts failing because "\x1b[K" is now
		// stripped, fold this case into the roundtrip table above.
		out := pterm.RemoveColorFromString(pterm.NewRGB(12, 34, 56, true).Sprint("filled"))
		assert.Equal(t, "filled\x1b[K", out)
	})
}

// Disabling colors must not only stop new sequences from being emitted, it
// must also strip escape codes already embedded in the input (renderCode).

func TestDisableColorStripsEmbeddedEscapeCodes(t *testing.T) {
	restoreGlobalStyling(t)
	pterm.DisableColor()

	assert.Equal(t, "plain", pterm.FgRed.Sprint("plain"))
	assert.Equal(t, "green text", pterm.FgRed.Sprint("\x1b[32mgreen\x1b[0m text"))
	assert.Equal(t, "styled", pterm.NewStyle(pterm.FgRed, pterm.Bold).Sprint("\x1b[1mstyled\x1b[0m"))
}

func TestEnableColorRestoresSequences(t *testing.T) {
	restoreGlobalStyling(t)

	pterm.DisableColor()
	assert.Equal(t, "x", pterm.FgRed.Sprint("x"))

	pterm.EnableColor()
	assert.Equal(t, "\x1b[31mx\x1b[0m", pterm.FgRed.Sprint("x"))
}
