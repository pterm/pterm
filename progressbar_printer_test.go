package pterm_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
)

// terminalEscapeRegexp matches any CSI escape sequence (colors, cursor
// movement, line clearing, ...). stripANSI (contract_test.go) only removes
// SGR color codes, but live printers also emit cursor-control sequences.
var terminalEscapeRegexp = regexp.MustCompile(`\x1b\[[0-9;?]*[A-Za-z]`)

// stripTerminalEscapes removes every terminal escape sequence from s, leaving
// only the text a user would actually see.
func stripTerminalEscapes(s string) string {
	return stripANSI(terminalEscapeRegexp.ReplaceAllString(s, ""))
}

// lastFrame returns the visible content of the last carriage-return frame in
// s. Live printers overwrite their line by emitting "\r" followed by the new
// frame, so the content after the final "\r" is what stays on screen.
func lastFrame(s string) string {
	frames := strings.Split(s, "\r")

	return stripTerminalEscapes(frames[len(frames)-1])
}

// waitFor polls cond until it returns true, failing the test with failMsg
// after a generous deadline. Polling keeps tests fast in the common case
// without relying on fixed sleeps that get flaky under -race.
func waitFor(t *testing.T, cond func() bool, failMsg func() string) {
	t.Helper()

	deadline := time.Now().Add(10 * time.Second)
	for time.Now().Before(deadline) {
		if cond() {
			return
		}

		time.Sleep(5 * time.Millisecond)
	}

	t.Fatal(failMsg())
}

// waitForOutput polls buf until its visible content (escape sequences
// stripped) contains want.
func waitForOutput(t *testing.T, buf fmt.Stringer, want string) {
	t.Helper()

	waitFor(t, func() bool {
		return strings.Contains(stripTerminalEscapes(buf.String()), want)
	}, func() string {
		return fmt.Sprintf("output did not contain %q within the timeout, got:\n%s", want, buf.String())
	})
}

// plainHalfBar returns a half-full progressbar with all decorators disabled
// and simple ASCII bar characters, so the bar geometry can be asserted
// exactly: MaxWidth 21 leaves 20 cells for the bar (one cell is reserved by
// the decorator separator), of which 10 must be filled at 50%.
func plainHalfBar() *pterm.ProgressbarPrinter {
	return pterm.DefaultProgressbar.
		WithTotal(10).
		WithCurrent(5).
		WithMaxWidth(21).
		WithShowTitle(false).
		WithShowCount(false).
		WithShowPercentage(false).
		WithShowElapsedTime(false).
		WithBarCharacter("#").
		WithLastCharacter("#").
		WithBarFiller("-").
		WithBarPartialCharacters(nil)
}

func TestProgressbarPrinter_BarIsHalfFilledAtFiftyPercent(t *testing.T) {
	buf := &syncBuffer{}

	p, err := plainHalfBar().WithWriter(buf).Start()
	require.NoError(t, err)

	defer p.Stop()

	frame := lastFrame(buf.String())
	assert.Equal(t, 10, strings.Count(frame, "#"), "10 of 20 bar cells must be filled at 5/10")
	assert.Equal(t, 10, strings.Count(frame, "-"), "10 of 20 bar cells must remain unfilled at 5/10")
	assert.Equal(t, "##########----------", strings.TrimRight(frame, " "), "filled cells must come first, unfilled cells last")
}

func TestProgressbarPrinter_SmoothBarRendersPartialEdge(t *testing.T) {
	buf := &syncBuffer{}

	// 1/3 of a 20-cell bar is 6.66 cells: 6 full blocks plus a partial glyph
	// covering 2/3 of the next cell (slot 5 of the 7 eighth-block glyphs).
	p, err := pterm.DefaultProgressbar.
		WithTotal(3).
		WithCurrent(1).
		WithMaxWidth(21).
		WithShowTitle(false).
		WithShowCount(false).
		WithShowPercentage(false).
		WithShowElapsedTime(false).
		WithWriter(buf).
		Start()
	require.NoError(t, err)

	defer p.Stop()

	frame := lastFrame(buf.String())
	assert.Contains(t, frame, strings.Repeat("█", 6)+"▋", "the bar edge must be drawn with a partial block glyph")
}

func TestProgressbarPrinter_ShowPercentage(t *testing.T) {
	buf := &syncBuffer{}

	p, err := plainHalfBar().WithShowPercentage(true).WithWriter(buf).Start()
	require.NoError(t, err)

	defer p.Stop()

	assert.Contains(t, lastFrame(buf.String()), "50%")
}

func TestProgressbarPrinter_ShowCount(t *testing.T) {
	buf := &syncBuffer{}

	p, err := plainHalfBar().WithShowCount(true).WithWriter(buf).Start()
	require.NoError(t, err)

	defer p.Stop()

	assert.Contains(t, lastFrame(buf.String()), "5/10")
}

func TestProgressbarPrinter_StartWithCurrentAboveZero(t *testing.T) {
	buf := &syncBuffer{}

	p, err := plainHalfBar().
		WithCurrent(3).
		WithShowCount(true).
		WithShowPercentage(true).
		WithMaxWidth(60).
		WithWriter(buf).
		Start()
	require.NoError(t, err)

	defer p.Stop()

	frame := lastFrame(buf.String())
	assert.Contains(t, frame, "3/10", "the initial frame must reflect the preset current value")
	assert.Contains(t, frame, "30%")
}

func TestProgressbarPrinter_StartWithTitleRendersTitle(t *testing.T) {
	buf := &syncBuffer{}

	p, err := pterm.DefaultProgressbar.WithTotal(10).WithShowElapsedTime(false).WithWriter(buf).Start("My Title")
	require.NoError(t, err)

	defer p.Stop()

	assert.Equal(t, "My Title", p.Title)
	assert.Contains(t, lastFrame(buf.String()), "My Title")
}

func TestProgressbarPrinter_UpdateTitleRerenders(t *testing.T) {
	buf := &syncBuffer{}

	p, err := pterm.DefaultProgressbar.
		WithTotal(10).
		WithTitle("before title").
		WithShowElapsedTime(false).
		WithWriter(buf).
		Start()
	require.NoError(t, err)

	defer p.Stop()

	assert.Contains(t, lastFrame(buf.String()), "before title")

	p.UpdateTitle("after title")

	frame := lastFrame(buf.String())
	assert.Contains(t, frame, "after title", "the new title must be rendered immediately")
	assert.NotContains(t, frame, "before title", "the old title must be gone from the current frame")
	assert.Equal(t, "after title", p.Title)
}

func TestProgressbarPrinter_AddClampsAtTotalAndAutoStops(t *testing.T) {
	buf := &syncBuffer{}

	p, err := pterm.DefaultProgressbar.
		WithTotal(3).
		WithShowElapsedTime(false).
		WithWriter(buf).
		Start()
	require.NoError(t, err)

	defer p.Stop()

	p.Add(5)

	assert.Equal(t, 5, p.Current)
	assert.Equal(t, p.Current, p.Total, "overshooting Add must clamp the total to the current value")
	assert.False(t, p.IsActive, "the progressbar must auto-stop when reaching its total")
	assert.Contains(t, lastFrame(buf.String()), "100%")
}

func TestProgressbarPrinter_IncrementAutoStopsAtTotal(t *testing.T) {
	buf := &syncBuffer{}

	p, err := pterm.DefaultProgressbar.
		WithTotal(2).
		WithShowElapsedTime(false).
		WithWriter(buf).
		Start()
	require.NoError(t, err)

	defer p.Stop()

	p.Increment()
	assert.Equal(t, 1, p.Current)
	assert.True(t, p.IsActive, "the progressbar must keep running below its total")
	assert.Contains(t, lastFrame(buf.String()), "50%")

	p.Increment()
	assert.Equal(t, 2, p.Current)
	assert.False(t, p.IsActive, "the progressbar must auto-stop when reaching its total")
	assert.Contains(t, lastFrame(buf.String()), "100%")
}

func TestProgressbarPrinter_AddWithTotalZeroIsNoop(t *testing.T) {
	buf := &syncBuffer{}
	p := pterm.ProgressbarPrinter{}.WithTotal(0).WithWriter(buf)

	assert.Nil(t, p.Add(1337), "Add must bail out on a zero total")
	assert.Equal(t, 0, p.Current)
	assert.Empty(t, buf.String(), "nothing must be rendered for a zero total")
}

func TestProgressbarPrinter_RemoveWhenDoneClearsLine(t *testing.T) {
	buf := &syncBuffer{}

	p, err := plainHalfBar().WithRemoveWhenDone().WithWriter(buf).Start()
	require.NoError(t, err)

	_, err = p.Stop()
	require.NoError(t, err)

	out := buf.String()
	assert.True(t, strings.HasSuffix(out, "\r"+strings.Repeat(" ", 80)+"\r"),
		"stopping must blank the line and return the cursor to its start, got:\n%q", out)
	assert.Empty(t, lastFrame(out), "no bar content may remain visible after Stop")
}

func TestProgressbarPrinter_StopWithoutRemoveKeepsLine(t *testing.T) {
	buf := &syncBuffer{}

	p, err := plainHalfBar().WithWriter(buf).Start()
	require.NoError(t, err)

	_, err = p.Stop()
	require.NoError(t, err)

	out := buf.String()
	assert.True(t, strings.HasSuffix(out, "\n"), "stopping must finish the bar line with a newline")
	assert.Contains(t, lastFrame(out), "#####", "the final bar must stay visible")
}

func TestProgressbarPrinter_ElapsedTimeIsRoundedToFactor(t *testing.T) {
	buf := &syncBuffer{}

	p, err := plainHalfBar().
		WithShowElapsedTime(true).
		WithElapsedTimeRoundingFactor(time.Second).
		WithMaxWidth(60).
		WithWriter(buf).
		Start()
	require.NoError(t, err)

	defer p.Stop()

	p.SetStartedAt(time.Now().Add(-90 * time.Second))
	p.Add(1)

	visible := stripTerminalEscapes(buf.String())
	assert.Regexp(t, `1m3[01]s`, visible, "the elapsed time must be rendered rounded to full seconds")
	assert.NotContains(t, visible, ".", "a rounded elapsed time must not contain fractional seconds")
}

// TestProgressbarPrinter_ZeroElapsedTimeRoundingFactor pins the regression
// where a rounding factor of zero must not blow up the elapsed time rendering;
// instead the raw, unrounded duration is shown.
func TestProgressbarPrinter_ZeroElapsedTimeRoundingFactor(t *testing.T) {
	buf := &syncBuffer{}

	p, err := plainHalfBar().
		WithShowElapsedTime(true).
		WithElapsedTimeRoundingFactor(0).
		WithMaxWidth(60).
		WithWriter(buf).
		Start()
	require.NoError(t, err)

	defer p.Stop()

	p.SetStartedAt(time.Now().Add(-90 * time.Second))

	assert.NotPanics(t, func() { p.Add(1) })
	assert.Contains(t, stripTerminalEscapes(buf.String()), "1m30.", "the unrounded elapsed time must be rendered")
}

func TestProgressbarPrinter_GetElapsedTime(t *testing.T) {
	buf := &syncBuffer{}

	p, err := plainHalfBar().WithWriter(buf).Start()
	require.NoError(t, err)

	defer p.Stop()

	p.SetStartedAt(time.Now().Add(-time.Minute))

	assert.GreaterOrEqual(t, p.GetElapsedTime(), time.Minute)
}

func TestProgressbarPrinter_PrintWhileActiveReprintsBar(t *testing.T) {
	buf := &syncBuffer{}

	p, err := pterm.DefaultProgressbar.
		WithTotal(10).
		WithCurrent(5).
		WithTitle("interplay title").
		WithShowElapsedTime(false).
		WithWriter(buf).
		Start()
	require.NoError(t, err)

	defer p.Stop()

	pterm.Fprintln(buf, "regular text")

	out := buf.String()
	frames := strings.Split(out, "\r")
	assert.Contains(t, frames, "regular text\n", "the printed text must end up intact on its own line")

	visible := stripTerminalEscapes(out)
	assert.Greater(t, strings.LastIndex(visible, "interplay title"), strings.Index(visible, "regular text"),
		"the bar must be re-rendered after the printed text")

	frame := lastFrame(out)
	assert.Contains(t, frame, "interplay title", "the re-rendered bar must be the last frame")
	assert.Contains(t, frame, "50%")
}

func TestProgressbarPrinter_MaxWidthIsClampedToTerminal(t *testing.T) {
	buf := &syncBuffer{}

	// A MaxWidth wider than the (forced 80 column) terminal must be clamped,
	// so a full bar line occupies exactly the terminal width.
	p, err := plainHalfBar().WithCurrent(10).WithMaxWidth(200).WithWriter(buf).Start()
	require.NoError(t, err)

	defer p.Stop()

	assert.Len(t, lastFrame(buf.String()), 80, "the rendered line must be clamped to the terminal width")
}

func TestProgressbarPrinter_TinyTerminalStillShowsDecorators(t *testing.T) {
	pterm.SetForcedTerminalSize(1, 1)
	t.Cleanup(func() { pterm.SetForcedTerminalSize(terminalWidth, terminalHeight) })

	buf := &syncBuffer{}

	p, err := pterm.DefaultProgressbar.WithTotal(2).WithShowElapsedTime(false).WithWriter(buf).Start()
	require.NoError(t, err)

	defer p.Stop()

	p.Add(1)

	assert.Equal(t, 1, p.Current)
	assert.Contains(t, lastFrame(buf.String()), "1/2", "decorators must render even when no space is left for the bar")
}

func TestProgressbarPrinter_RawOutputPrintsPlainTitle(t *testing.T) {
	restoreGlobalStyling(t)
	pterm.DisableStyling()

	buf := &syncBuffer{}

	p, err := pterm.DefaultProgressbar.
		WithTotal(10).
		WithTitle("raw title").
		WithShowElapsedTime(false).
		WithWriter(buf).
		Start()
	require.NoError(t, err)

	defer p.Stop()

	out := buf.String()
	assert.Contains(t, out, "raw title\n", "raw output must print the title as a plain line")
	assert.NotContains(t, out, "\x1b[", "raw output must not contain escape codes")
}
