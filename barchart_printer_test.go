package pterm_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal/text"
	"github.com/pterm/pterm/internal/vt"
)

var (
	downloads = pterm.Bars{
		{Label: "npm", Value: 1284},
		{Label: "pnpm", Value: 932.5},
		{Label: "yarn", Value: 611},
		{Label: "bun", Value: 287},
	}
	quarters = pterm.Bars{
		{Label: "Q1", Value: 12.5},
		{Label: "Q2", Value: -4},
		{Label: "Q3", Value: -9.25},
		{Label: "Q4", Value: 18},
	}
	yearly = pterm.BarGroups{
		{Label: "Q1", Values: []float64{42, 51}},
		{Label: "Q2", Values: []float64{38, 57}},
		{Label: "Q3", Values: []float64{45, 63.5}},
	}
	timings = pterm.BarGroups{
		{Label: "build", Values: []float64{120, 45, 12}},
		{Label: "test", Values: []float64{210, 80, 5}},
		{Label: "deploy", Values: []float64{30, 20, 140}},
	}
)

func TestBarChartSnapshots(t *testing.T) {
	chart := pterm.BarChart

	snap(t, "Default", srender(t, chart.WithBars(downloads).Srender))
	snap(t, "Negative", srender(t, chart.WithBars(quarters).WithUnit("%").Srender))
	snap(t, "AllNegative", srender(t, chart.WithBars(pterm.Bars{{Label: "a", Value: -3}, {Label: "b", Value: -1.5}}).Srender))
	snap(t, "Grouped", srender(t, chart.WithSeries("2025", "2026").WithGroups(yearly).Srender))
	snap(t, "Stacked", srender(t, chart.WithSeries("cpu", "io", "network").WithStacked().WithUnit(" ms").WithGroups(timings).Srender))
	snap(t, "Plain", srender(t, chart.WithBars(downloads).WithShowTrack(false).WithShowValues(false).WithTone(pterm.ToneInfo).Srender))
	snap(t, "FixedWidth", srender(t, chart.WithBars(downloads).WithWidth(40).WithRange(0, 2000).Srender))
	snap(t, "Vertical", srender(t, chart.WithVertical().WithBars(downloads).Srender))
	snap(t, "VerticalNegative", srender(t, chart.WithVertical().WithHeight(6).WithBars(quarters).Srender))
	snap(t, "VerticalGrouped", srender(t, chart.WithVertical().WithSeries("2025", "2026").WithGroups(yearly).Srender))
	snap(t, "VerticalStacked", srender(t, chart.WithVertical().WithStacked().WithSeries("cpu", "io", "network").WithGroups(timings).Srender))
}

func TestBarChartOneLiner(t *testing.T) {
	setup(t)

	out := capture(t, func() {
		if err := pterm.BarChart.WithBars(pterm.Bars{{Label: "npm", Value: 12}}).Render(); err != nil {
			t.Fatal(err)
		}
	})

	if !strings.HasPrefix(out, "npm  ███") || !strings.HasSuffix(out, "  12\n") {
		t.Errorf("output = %q", out)
	}
}

func TestBarChartEmpty(t *testing.T) {
	setup(t)

	if _, err := pterm.BarChart.Srender(); !errors.Is(err, pterm.ErrEmpty) {
		t.Errorf("err = %v, want ErrEmpty", err)
	}

	if err := pterm.BarChart.WithGroups(pterm.BarGroups{}).Render(); !errors.Is(err, pterm.ErrEmpty) {
		t.Errorf("err = %v, want ErrEmpty", err)
	}
}

func TestBarChartUnknownWidth(t *testing.T) {
	setup(t)
	t.Setenv("COLUMNS", "")

	out := capture(t, func() { _ = pterm.BarChart.WithBars(downloads).Render() })

	for line := range strings.SplitSeq(strings.TrimSuffix(out, "\n"), "\n") {
		// label column (4) + gap + 40 columns of bars + gap + values (5)
		if w := text.Width(line); w != 4+2+40+2+5 {
			t.Errorf("line %q is %d wide", line, w)
		}
	}
}

func TestBarChartFitsNarrowTerminals(t *testing.T) {
	setup(t)

	for _, width := range []int{20, 30, 50} {
		term := vt.New(width, 24)

		for _, chart := range []*pterm.BarChartPrinter{
			pterm.BarChart.WithBars(downloads),
			pterm.BarChart.WithBars(quarters),
			pterm.BarChart.WithSeries("a", "b").WithGroups(yearly).WithWidth(200),
			pterm.BarChart.WithStacked().WithGroups(timings),
		} {
			s, err := chart.WithWriter(term).Srender()
			if err != nil {
				t.Fatal(err)
			}

			for line := range strings.SplitSeq(s, "\n") {
				if w := text.Width(line); w > width {
					t.Errorf("width %d: line %q is %d wide", width, line, w)
				}
			}
		}
	}
}

func TestBarChartLongAndWideLabels(t *testing.T) {
	setup(t)

	s, err := pterm.BarChart.WithBars(pterm.Bars{
		{Label: "日本語のラベル", Value: 3},
		{Label: "a label that is far too long for the label column", Value: 5},
		{Label: "\x1b[1mbold\x1b[0m", Value: 1},
	}).Srender()
	if err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(s, "\n")
	for _, line := range lines {
		if w := text.Width(line); w > 80 {
			t.Errorf("line %q is %d wide", line, w)
		}
	}

	// All bars start in the same column.
	start := -1

	for _, line := range lines {
		plain := pterm.RemoveColorFromString(line)

		i := strings.Index(plain, "█")
		if i < 0 {
			t.Fatalf("no bar in %q", plain)
		}

		col := text.Width(plain[:i])
		if start >= 0 && col != start {
			t.Errorf("bars start at %d and %d", start, col)
		}

		start = col
	}
}

func TestBarChartFormat(t *testing.T) {
	setup(t)

	s, _ := pterm.BarChart.WithBars(pterm.Bars{{Label: "disk", Value: 0.5}}).
		WithFormat(func(v float64) string { return fmt.Sprintf("%.0f", v*100) }).
		WithUnit("%").
		Srender()

	if !strings.HasSuffix(pterm.RemoveColorFromString(s), "  50%") {
		t.Errorf("output = %q", s)
	}
}

func TestBarChartRaw(t *testing.T) {
	setup(t)
	pterm.DisableStyling()

	s, _ := pterm.BarChart.WithBars(downloads[:2]).Srender()
	if s != "npm: 1284\npnpm: 932.5" {
		t.Errorf("raw = %q", s)
	}

	s, _ = pterm.BarChart.WithSeries("2025", "2026").WithGroups(yearly[:1]).Srender()
	if s != "Q1: 2025=42, 2026=51" {
		t.Errorf("raw = %q", s)
	}
}

func TestBarChartASCII(t *testing.T) {
	setup(t)
	pterm.SetGlyphs(pterm.GlyphsASCII)

	snap(t, "ASCII", func() string {
		pterm.SetGlyphs(pterm.GlyphsASCII)
		defer pterm.SetGlyphs(pterm.GlyphsUnicode)

		s, _ := pterm.BarChart.WithBars(quarters).WithWidth(40).Srender()

		v, _ := pterm.BarChart.WithVertical().WithHeight(4).WithBars(quarters).Srender()

		return s + "\n\n" + v
	})
}

func TestBarChartDoesNotChangeDefault(t *testing.T) {
	setup(t)

	before := fmt.Sprintf("%+v", pterm.BarChart)
	_ = pterm.BarChart.WithVertical().WithBars(downloads).WithTones(pterm.ToneError).WithWidth(3)

	if after := fmt.Sprintf("%+v", pterm.BarChart); after != before {
		t.Error("a With method changed the default bar chart")
	}
}
