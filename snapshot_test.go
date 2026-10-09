package pterm_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal/snapshot"
)

// testSnapshot renders the same output twice — once with styling enabled and
// once in plain mode — and compares both against committed snapshots. This
// locks the exact output of every printer so any rendering change shows up as
// a snapshot diff in the pull request.
func testSnapshot(t *testing.T, name string, render func(t *testing.T) string) {
	t.Helper()
	t.Run(name, func(t *testing.T) {
		pterm.EnableStyling()
		t.Run("styled", func(t *testing.T) {
			snapshot.Assert(t, render(t))
		})
		pterm.DisableStyling()
		t.Run("plain", func(t *testing.T) {
			snapshot.Assert(t, render(t))
		})
		pterm.EnableStyling()
	})
}

// srender adapts a Srender method for testSnapshot, failing the test on error.
func srender(t *testing.T, render func() (string, error)) string {
	t.Helper()

	s, err := render()
	assert.NoError(t, err)

	return s
}

func TestSnapshots(t *testing.T) {
	testSnapshot(t, "BasicText", func(_ *testing.T) string {
		return pterm.DefaultBasicText.Sprint("Hello, PTerm!")
	})

	testSnapshot(t, "Info", func(_ *testing.T) string {
		return pterm.Info.Sprint("This is an info message")
	})
	testSnapshot(t, "Success", func(_ *testing.T) string {
		return pterm.Success.Sprint("This is a success message")
	})
	testSnapshot(t, "Warning", func(_ *testing.T) string {
		return pterm.Warning.Sprint("This is a warning message")
	})
	testSnapshot(t, "Error", func(_ *testing.T) string {
		return pterm.Error.Sprint("This is an error message")
	})
	testSnapshot(t, "Description", func(_ *testing.T) string {
		return pterm.Description.Sprint("This is a description message")
	})
	testSnapshot(t, "Debug", func(_ *testing.T) string {
		pterm.EnableDebugMessages()

		defer pterm.DisableDebugMessages()

		return pterm.Debug.Sprint("This is a debug message")
	})
	testSnapshot(t, "PrefixMultiline", func(_ *testing.T) string {
		return pterm.Info.Sprint("First line\nSecond line\nThird line")
	})

	testSnapshot(t, "Header", func(_ *testing.T) string {
		return pterm.DefaultHeader.Sprint("PTerm Header")
	})
	testSnapshot(t, "HeaderFullWidth", func(_ *testing.T) string {
		return pterm.DefaultHeader.WithFullWidth().Sprint("PTerm Header")
	})

	testSnapshot(t, "Section", func(_ *testing.T) string {
		return pterm.DefaultSection.Sprint("Section Title")
	})

	testSnapshot(t, "Paragraph", func(_ *testing.T) string {
		return pterm.DefaultParagraph.Sprint("This is a longer paragraph text that will be wrapped by the paragraph printer to fit into the width of the terminal, demonstrating word wrapping behavior.")
	})

	testSnapshot(t, "Center", func(_ *testing.T) string {
		return pterm.DefaultCenter.Sprint("centered text\nwith multiple lines")
	})

	testSnapshot(t, "Box", func(_ *testing.T) string {
		return pterm.DefaultBox.Sprint("Boxed content")
	})
	testSnapshot(t, "BoxWithTitle", func(_ *testing.T) string {
		return pterm.DefaultBox.WithTitle("Title").Sprint("Boxed content\nwith two lines")
	})

	testSnapshot(t, "BulletList", func(t *testing.T) string {
		return srender(t, pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
			{Level: 0, Text: "Level 0"},
			{Level: 1, Text: "Level 1", TextStyle: pterm.NewStyle(pterm.FgGreen)},
			{Level: 2, Text: "Level 2", Bullet: ">", BulletStyle: pterm.NewStyle(pterm.FgCyan)},
		}).Srender)
	})

	testSnapshot(t, "Tree", func(t *testing.T) string {
		return srender(t, pterm.DefaultTree.WithRoot(pterm.TreeNode{
			Text: "Root",
			Children: []pterm.TreeNode{
				{Text: "First child", Children: []pterm.TreeNode{
					{Text: "First grandchild"},
					{Text: "Second grandchild"},
				}},
				{Text: "Second child"},
			},
		}).Srender)
	})

	testSnapshot(t, "Table", func(t *testing.T) string {
		return srender(t, pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{
			{"Firstname", "Lastname", "Email"},
			{"Paul", "Dean", "nisi.dictum.augue@velitAliquam.co.uk"},
			{"Callie", "Mckay", "egestas.nunc.sed@est.com"},
			{"Libby", "Camacho", "aliquet.lobortis@semper.com"},
		}).Srender)
	})
	testSnapshot(t, "TableBoxed", func(t *testing.T) string {
		return srender(t, pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(pterm.TableData{
			{"Firstname", "Lastname"},
			{"Paul", "Dean"},
		}).Srender)
	})

	testSnapshot(t, "Panel", func(t *testing.T) string {
		return srender(t, pterm.DefaultPanel.WithPanels(pterm.Panels{
			{{Data: "First panel"}, {Data: "Second panel\nwith two lines"}},
			{{Data: "Third panel"}},
		}).Srender)
	})

	testSnapshot(t, "BarChartHorizontal", func(t *testing.T) string {
		return srender(t, pterm.DefaultBarChart.WithHorizontal().WithShowValue().WithBars(pterm.Bars{
			{Label: "A", Value: 10},
			{Label: "B", Value: 20},
			{Label: "C", Value: 5},
		}).Srender)
	})
	testSnapshot(t, "BarChartVertical", func(t *testing.T) string {
		return srender(t, pterm.DefaultBarChart.WithShowValue().WithBars(pterm.Bars{
			{Label: "A", Value: 10},
			{Label: "B", Value: 20},
			{Label: "C", Value: 5},
		}).Srender)
	})

	testSnapshot(t, "BigText", func(t *testing.T) string {
		return srender(t, pterm.DefaultBigText.WithLetters(pterm.NewLettersFromStringWithStyle("PT", pterm.NewStyle(pterm.FgCyan))).Srender)
	})

	testSnapshot(t, "Heatmap", func(t *testing.T) string {
		return srender(t, pterm.DefaultHeatmap.WithAxisData(pterm.HeatmapAxis{
			XAxis: []string{"a", "b", "c"},
			YAxis: []string{"1", "2"},
		}).WithData([][]float32{
			{1, 2, 3},
			{4, 5, 6},
		}).WithLegend().Srender)
	})

	testSnapshot(t, "Color", func(_ *testing.T) string {
		return pterm.FgRed.Sprint("red") + " " + pterm.FgGreen.Sprint("green") + " " + pterm.BgBlue.Sprint("blue bg")
	})
	testSnapshot(t, "Style", func(_ *testing.T) string {
		return pterm.NewStyle(pterm.FgRed, pterm.BgBlack, pterm.Bold).Sprint("bold red on black")
	})
	testSnapshot(t, "RGB", func(_ *testing.T) string {
		return pterm.NewRGB(255, 0, 255).Sprint("magenta rgb")
	})
	testSnapshot(t, "RGBBackground", func(_ *testing.T) string {
		return pterm.NewRGB(0, 128, 255, true).Sprint("rgb background")
	})
	testSnapshot(t, "RGBStyle", func(_ *testing.T) string {
		return pterm.NewRGBStyle(pterm.RGB{R: 255, G: 0, B: 0}, pterm.RGB{R: 0, G: 0, B: 128}).AddOptions(pterm.Bold).Sprint("styled rgb")
	})

	testSnapshot(t, "Logger", func(_ *testing.T) string {
		var buf bytes.Buffer
		logger := pterm.DefaultLogger.WithWriter(&buf).WithTime(false)
		logger.Info("An info log", logger.Args("key", "value"))
		logger.Warn("A warning log")
		logger.Error("An error log", logger.Args("code", 42, "spaced", "a quoted value"))

		return buf.String()
	})

	testSnapshot(t, "LoggerBlock", func(t *testing.T) string {
		pterm.SetForcedTerminalSize(80, 24)
		t.Cleanup(func() { pterm.SetForcedTerminalSize(0, 0) })

		var buf bytes.Buffer
		logger := pterm.DefaultLogger.WithWriter(&buf).WithTime(false).WithMaxWidth(60)
		logger.Info("A message that is too long for a single line wraps under its own first line", logger.Args("key", "value", "another", "value with spaces"))
		logger.Warn("Multiline and long values align under the value start", logger.Args("stack", "first line\nsecond line", "token", "averylongsingletokenthatmustbebrokenhardbecauseitdoesnotfit"))

		return buf.String()
	})
}
