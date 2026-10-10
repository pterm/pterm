package pterm_test

import (
	"errors"
	"math"
	"strings"
	"sync"
	"testing"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal/text"
	"github.com/pterm/pterm/internal/vt"
)

func TestMeterSnapshots(t *testing.T) {
	snap(t, "Single", srender(t, pterm.Meter.WithValue(0.34, "well under budget").Srender))
	snap(t, "Thresholds", srender(t, pterm.Meter.WithItems(
		pterm.MeterItem{Label: "cpu", Value: 0.42},
		pterm.MeterItem{Label: "memory", Value: 0.785},
		pterm.MeterItem{Label: "disk", Value: 0.93},
		pterm.MeterItem{Label: "quota", Value: 1.12},
		pterm.MeterItem{Label: "custom", Value: 0.5, Tone: pterm.ToneInfo},
	).Srender))
	snap(t, "ToneAndWidth", srender(t, pterm.Meter.WithValue(0.5).WithTone(pterm.TonePrimary).WithWidth(8).Srender))
}

func TestMeterEdgeCases(t *testing.T) {
	setup(t)

	if _, err := pterm.Meter.Srender(); !errors.Is(err, pterm.ErrEmpty) {
		t.Errorf("err = %v, want ErrEmpty", err)
	}

	out := capture(t, func() {
		_ = pterm.Meter.WithItems(
			pterm.MeterItem{Value: math.NaN()},
			pterm.MeterItem{Value: -1},
			pterm.MeterItem{Value: 3},
		).Render()
	})

	want := "░░░░░░░░░░░░░░░░░░░░     0%\n" +
		"░░░░░░░░░░░░░░░░░░░░  -100%\n" +
		"████████████████████   300%\n"
	if out != want {
		t.Errorf("output =\n%s\nwant\n%s", out, want)
	}

	s, _ := pterm.Meter.WithValue(0.5, "a label that does not fit").WithWriter(vt.New(20, 5)).Srender()
	if w := text.Width(s); w > 20 {
		t.Errorf("meter is %d wide on a 20 column terminal: %q", w, s)
	}
}

func TestSparklineSnapshots(t *testing.T) {
	snap(t, "Default", srender(t, pterm.Sparkline.WithValues(3, 5, 4, 8, 12, 9, 14, 18, 16, 21, 19, 24).WithLabel("requests per second").Srender))
	snap(t, "Thresholds", srender(t, pterm.Sparkline.
		WithValues(12, 34, 88, 56, 71, 23, 95, 41).
		WithRange(0, 100).
		WithThresholds(pterm.LoadThresholds()...).
		WithLabel("8 cores").Srender))
}

func TestSparklineEdgeCases(t *testing.T) {
	setup(t)

	if _, err := pterm.Sparkline.Srender(); !errors.Is(err, pterm.ErrEmpty) {
		t.Errorf("err = %v, want ErrEmpty", err)
	}

	out := capture(t, func() { _ = pterm.Sparkline.WithValues(1, 2, math.NaN(), 3).Render() })
	if out != "▁▅ █\n" {
		t.Errorf("output = %q", out)
	}

	out = capture(t, func() { _ = pterm.Sparkline.WithValues(5, 5, 5).Render() })
	if out != "▅▅▅\n" {
		t.Errorf("flat output = %q", out)
	}

	// A sparkline wider than the terminal keeps the most recent values.
	s, _ := pterm.Sparkline.WithValues(1, 2, 3, 4, 5, 6, 7, 8).WithRange(0, 8).WithWriter(vt.New(4, 5)).Srender()
	if pterm.RemoveColorFromString(s) != "▅▆▇█" {
		t.Errorf("cut output = %q", s)
	}

	s, _ = pterm.Sparkline.WithValues(1, 2, 3, 4, 5, 6, 7, 8).WithWidth(3).Srender()
	if text.Width(s) != 3 {
		t.Errorf("WithWidth(3) output = %q", s)
	}
}

var (
	commits = pterm.HeatmapData{
		{3, 6, 4, 7, 12, 24},
		{5, 3, 6, 4, 14, 26},
		{1, 4, 2, 5, 5, 6},
	}
	commitAxis = pterm.HeatmapAxis{
		XAxis: []string{"08", "10", "12", "14", "16", "18"},
		YAxis: []string{"Mon", "Tue", "Sat"},
	}
)

func TestHeatmapSnapshots(t *testing.T) {
	heatmap := pterm.Heatmap.WithData(commits)

	snap(t, "Default", srender(t, heatmap.WithAxisData(commitAxis).WithLegendLabel("commits").Srender))
	snap(t, "NoAxis", srender(t, heatmap.WithLegend(false).Srender))
	snap(t, "Values", srender(t, heatmap.WithAxisData(commitAxis).WithShowValues().WithGradient(pterm.TrafficTones()...).Srender))
	snap(t, "Grid", srender(t, heatmap.WithAxisData(commitAxis).WithShowValues().WithGrid().Srender))
	snap(t, "Boxed", srender(t, heatmap.WithShowValues().WithBoxed().WithRange(0, 30).Srender))
}

func TestHeatmapEdgeCases(t *testing.T) {
	setup(t)

	if _, err := pterm.Heatmap.Srender(); !errors.Is(err, pterm.ErrEmpty) {
		t.Errorf("err = %v, want ErrEmpty", err)
	}

	// Ragged rows and missing values leave blank cells, and missing axis
	// labels are blank too.
	s, err := pterm.Heatmap.WithData(pterm.HeatmapData{{1, 2, 3}, {4}, {math.NaN(), 6}}).
		WithAxisData(pterm.HeatmapAxis{YAxis: []string{"a"}}).
		WithLegend(false).
		Srender()
	if err != nil {
		t.Fatal(err)
	}

	if got := len(strings.Split(s, "\n")); got != 3 {
		t.Errorf("got %d lines:\n%s", got, s)
	}

	pterm.SetGlyphs(pterm.GlyphsASCII)
	defer pterm.SetGlyphs(pterm.GlyphsUnicode)

	out := capture(t, func() { _ = pterm.Heatmap.WithData(pterm.HeatmapData{{0, 5, 10}}).WithCellWidth(1).Render() })
	if out != ". | #\n\n0 .....:::::|||||##### 10\n" {
		t.Errorf("plain ASCII output = %q", out)
	}
}

func TestBigTextSnapshots(t *testing.T) {
	snap(t, "Default", func() string { return pterm.BigText.Sprint("PTerm") })
	snap(t, "Tone", func() string { return pterm.BigText.WithTone(pterm.ToneSuccess).Sprint("ok!") })
	snap(t, "Offset", srender(t, pterm.BigText.WithText("PT").WithGradientOffset(0.5).Srender))
	snap(t, "Letters", srender(t, pterm.BigText.WithLetters(
		pterm.LettersFromString("P", pterm.ToneError.Bold()),
		pterm.LettersFromString("Term"),
	).Srender))
	snap(t, "Lines", func() string { return pterm.BigText.Sprint("a\nb") })
}

func TestBigText(t *testing.T) {
	setup(t)

	if _, err := pterm.BigText.Srender(); !errors.Is(err, pterm.ErrEmpty) {
		t.Errorf("err = %v, want ErrEmpty", err)
	}

	// Unknown characters are skipped, lowercase letters use uppercase.
	if pterm.BigText.Sprint("a€") != pterm.BigText.Sprint("A") {
		t.Error("unknown characters or lowercase letters render differently")
	}

	if a, b := pterm.BigText.Sprint("PTerm"), pterm.BigText.WithGradientOffset(2).Sprint("PTerm"); a != b {
		t.Error("an offset of 2 should bring the gradient back to the start")
	}

	pterm.DisableStyling()

	if got := pterm.BigText.Sprint("PTerm"); got != "PTerm" {
		t.Errorf("unstyled = %q", got)
	}

	pterm.EnableStyling()
	pterm.SetGlyphs(pterm.GlyphsASCII)

	defer pterm.SetGlyphs(pterm.GlyphsUnicode)

	if got := pterm.RemoveColorFromString(pterm.BigText.Sprint("Q")); strings.ContainsAny(got, "█▀▄") {
		t.Errorf("ASCII big text has block characters:\n%s", got)
	}

	out := capture(t, func() { pterm.BigText.Println("I") })
	if out != "██\n██\n██\n██\n██\n" && out != "##\n##\n##\n##\n##\n" {
		t.Errorf("Println = %q", out)
	}
}

func TestChartsRenderConcurrently(t *testing.T) {
	setup(t)

	chart := pterm.BarChart.WithSeries("a", "b").WithGroups(yearly)
	want, _ := chart.Srender()

	var wg sync.WaitGroup

	for range 8 {
		wg.Go(func() {
			for range 20 {
				if got, _ := chart.Srender(); got != want {
					t.Error("concurrent renders differ")
				}

				_ = pterm.BigText.Sprint("go")
				_, _ = pterm.Heatmap.WithData(commits).Srender()
				_, _ = pterm.Meter.WithValue(0.5).Srender()
				_, _ = pterm.Sparkline.WithValues(1, 2, 3).Srender()
			}
		})
	}

	wg.Wait()
}
