package pterm

import (
	"bytes"
	"errors"
	"io"
	"math"
	"strings"

	"github.com/pterm/pterm/internal"
)

// rgbLegendSteps is the number of cells the legend fades through in RGB mode.
const rgbLegendSteps = 10

// DefaultHeatmap contains standards, which can be used to print a HeatmapPrinter.
var DefaultHeatmap = HeatmapPrinter{
	AxisStyle:                  &ThemeDefault.HeatmapHeaderStyle,
	SeparatorStyle:             &ThemeDefault.HeatmapSeparatorStyle,
	VerticalSeparator:          "│",
	TopRightCornerSeparator:    "╰",
	TopLeftCornerSeparator:     "╯",
	BottomLeftCornerSeparator:  "╮",
	BottomRightCornerSeparator: "╭",
	HorizontalSeparator:        "─",
	TSeparator:                 "┬",
	TReverseSeparator:          "┴",
	LSeparator:                 "├",
	LReverseSeparator:          "┤",
	TCrossSeparator:            "┼",
	LegendLabel:                "Legend",
	Boxed:                      true,
	Grid:                       true,
	Legend:                     true,
	TextRGB:                    ThemeDefault.HeatmapTextRGB,
	RGBRange:                   ThemeDefault.HeatmapRGBRange,
	TextColor:                  ThemeDefault.HeatmapTextColor,
	Colors:                     ThemeDefault.HeatmapColors,

	EnableRGB: false,
}

// HeatmapData is the type that contains the data of a HeatmapPrinter.
type HeatmapData [][]float32

// HeatmapAxis contains the labels of the X and Y axis of a HeatmapPrinter.
type HeatmapAxis struct {
	XAxis []string
	YAxis []string
}

// HeatmapPrinter is able to render tables.
type HeatmapPrinter struct {
	HasHeader                  bool
	AxisStyle                  *Style
	VerticalSeparator          string
	TopRightCornerSeparator    string
	TopLeftCornerSeparator     string
	BottomLeftCornerSeparator  string
	BottomRightCornerSeparator string
	HorizontalSeparator        string
	TSeparator                 string
	TReverseSeparator          string
	LSeparator                 string
	LReverseSeparator          string
	TCrossSeparator            string
	LegendLabel                string
	SeparatorStyle             *Style
	Data                       HeatmapData
	Axis                       HeatmapAxis
	Boxed                      bool
	Grid                       bool
	OnlyColoredCells           bool
	LegendOnlyColoredCells     bool
	EnableComplementaryColor   bool
	Legend                     bool
	CellSize                   int
	Colors                     []Color
	TextColor                  Color
	EnableRGB                  bool
	RGBRange                   []RGB
	TextRGB                    RGB
	Writer                     io.Writer

	minValue float32
	maxValue float32
}

var complementaryColors = map[Color]Color{
	BgBlack:        FgLightWhite,
	BgRed:          FgCyan,
	BgGreen:        FgMagenta,
	BgYellow:       FgBlue,
	BgBlue:         FgYellow,
	BgMagenta:      FgGreen,
	BgCyan:         FgRed,
	BgWhite:        FgBlack,
	BgDefault:      FgBlack,
	BgDarkGray:     FgLightWhite,
	BgLightRed:     FgLightCyan,
	BgLightGreen:   FgLightMagenta,
	BgLightYellow:  FgLightBlue,
	BgLightBlue:    FgLightYellow,
	BgLightMagenta: FgLightGreen,
	BgLightCyan:    FgLightRed,
	BgLightWhite:   FgBlack,
}

// WithAxisData returns a new HeatmapPrinter, where the first line and row are headers.
func (p HeatmapPrinter) WithAxisData(hd HeatmapAxis) *HeatmapPrinter {
	p.HasHeader = true
	p.Axis = hd

	return &p
}

// WithAxisStyle returns a new HeatmapPrinter with a specific AxisStyle.
func (p HeatmapPrinter) WithAxisStyle(style *Style) *HeatmapPrinter {
	p.AxisStyle = style
	return &p
}

// WithSeparatorStyle returns a new HeatmapPrinter with a specific SeparatorStyle.
func (p HeatmapPrinter) WithSeparatorStyle(style *Style) *HeatmapPrinter {
	p.SeparatorStyle = style
	return &p
}

// WithData returns a new HeatmapPrinter with specific Data.
func (p HeatmapPrinter) WithData(data [][]float32) *HeatmapPrinter {
	p.Data = data
	return &p
}

// WithTextColor returns a new HeatmapPrinter with a specific TextColor.
// This sets EnableComplementaryColor to false.
func (p HeatmapPrinter) WithTextColor(color Color) *HeatmapPrinter {
	p.TextColor = color
	p.EnableComplementaryColor = false

	return &p
}

// WithTextRGB returns a new HeatmapPrinter with a specific TextRGB.
// This sets EnableComplementaryColor to false.
func (p HeatmapPrinter) WithTextRGB(rgb RGB) *HeatmapPrinter {
	p.TextRGB = rgb
	p.EnableComplementaryColor = false

	return &p
}

// WithBoxed returns a new HeatmapPrinter with a box around the table.
// If set to true, Grid will be set to true too.
func (p HeatmapPrinter) WithBoxed(b ...bool) *HeatmapPrinter {
	p.Boxed = internal.WithBoolean(b)
	if p.Boxed && !p.Grid {
		p.Grid = true
	}

	return &p
}

// WithGrid returns a new HeatmapPrinter with a grid.
// If set to false, Boxed will be set to false too.
func (p HeatmapPrinter) WithGrid(b ...bool) *HeatmapPrinter {
	b2 := internal.WithBoolean(b)

	p.Grid = b2
	if !b2 && p.Boxed {
		p.Boxed = false
	}

	return &p
}

// WithEnableRGB returns a new HeatmapPrinter with RGB colors.
func (p HeatmapPrinter) WithEnableRGB(b ...bool) *HeatmapPrinter {
	p.EnableRGB = internal.WithBoolean(b)
	return &p
}

// WithOnlyColoredCells returns a new HeatmapPrinter with only colored cells.
func (p HeatmapPrinter) WithOnlyColoredCells(b ...bool) *HeatmapPrinter {
	b2 := internal.WithBoolean(b)
	p.OnlyColoredCells = b2

	return &p
}

// WithLegendOnlyColoredCells returns a new HeatmapPrinter with legend with only colored cells.
// This sets the Legend to true.
func (p HeatmapPrinter) WithLegendOnlyColoredCells(b ...bool) *HeatmapPrinter {
	b2 := internal.WithBoolean(b)

	p.LegendOnlyColoredCells = b2
	if b2 {
		p.Legend = true
	}

	return &p
}

// WithEnableComplementaryColor returns a new HeatmapPrinter with complement color.
func (p HeatmapPrinter) WithEnableComplementaryColor(b ...bool) *HeatmapPrinter {
	p.EnableComplementaryColor = internal.WithBoolean(b)
	return &p
}

// WithLegend returns a new HeatmapPrinter with a legend.
func (p HeatmapPrinter) WithLegend(b ...bool) *HeatmapPrinter {
	p.Legend = internal.WithBoolean(b)
	return &p
}

// WithCellSize returns a new HeatmapPrinter with a specific cell size.
// This only works if there is no header and OnlyColoredCells == true!
func (p HeatmapPrinter) WithCellSize(i int) *HeatmapPrinter {
	p.CellSize = i
	return &p
}

// WithLegendLabel returns a new HeatmapPrinter with a specific legend tag.
// This sets the Legend to true.
func (p HeatmapPrinter) WithLegendLabel(s string) *HeatmapPrinter {
	p.LegendLabel = s
	p.Legend = true

	return &p
}

// WithRGBRange returns a new HeatmapPrinter with a specific RGBRange.
func (p HeatmapPrinter) WithRGBRange(rgb ...RGB) *HeatmapPrinter {
	p.RGBRange = rgb
	return &p
}

// WithColors returns a new HeatmapPrinter with a specific Colors.
func (p HeatmapPrinter) WithColors(colors ...Color) *HeatmapPrinter {
	p.Colors = colors
	return &p
}

// WithWriter sets the Writer.
func (p HeatmapPrinter) WithWriter(writer io.Writer) *HeatmapPrinter {
	p.Writer = writer
	return &p
}

// Srender renders the HeatmapPrinter as a string.
func (p HeatmapPrinter) Srender() (string, error) {
	if err := p.errCheck(); err != nil {
		return "", err
	}

	if p.SeparatorStyle == nil {
		p.SeparatorStyle = DefaultHeatmap.SeparatorStyle
	}

	if p.AxisStyle == nil {
		p.AxisStyle = DefaultHeatmap.AxisStyle
	}

	if rawOutput() {
		p.Legend = false
	}

	p.minValue, p.maxValue = minMaxFloat32(p.Data)

	colWidth := p.columnWidth()
	legendColWidth := colWidth + 2

	if p.OnlyColoredCells && (p.CellSize > colWidth || !p.HasHeader) {
		colWidth = p.CellSize
	}

	buffer := bytes.NewBufferString("")

	if p.Boxed {
		p.writeGridLine(buffer, colWidth, p.BottomRightCornerSeparator, p.TSeparator, p.BottomLeftCornerSeparator)
		buffer.WriteString("\n")
	}

	p.writeRows(buffer, colWidth)

	if p.HasHeader {
		p.writeXAxisRow(buffer, colWidth)
	}

	if p.Boxed {
		buffer.WriteString("\n")
		p.writeGridLine(buffer, colWidth, p.TopRightCornerSeparator, p.TReverseSeparator, p.TopLeftCornerSeparator)
	}

	if p.Legend {
		p.writeLegend(buffer, legendColWidth)
	}

	buffer.WriteString("\n")

	return buffer.String(), nil
}

// columnWidth returns the width of the widest cell content: every data value,
// plus every axis label when a header is rendered. With OnlyColoredCells the
// values themselves are not printed, so only the axis labels count.
func (p HeatmapPrinter) columnWidth() int {
	var width int

	if !p.HasHeader || !p.OnlyColoredCells {
		for _, row := range p.Data {
			for _, value := range row {
				width = max(width, internal.GetStringMaxWidth(Sprintf("%v", value)))
			}
		}
	}

	if p.HasHeader {
		for _, label := range p.Axis.XAxis {
			width = max(width, internal.GetStringMaxWidth(label))
		}

		for _, label := range p.Axis.YAxis {
			width = max(width, internal.GetStringMaxWidth(label))
		}
	}

	return width
}

// gridColumns returns the number of rendered columns: one per data column,
// plus the Y-axis label column when a header is rendered.
func (p HeatmapPrinter) gridColumns() int {
	if p.HasHeader {
		return len(p.Data[0]) + 1
	}

	return len(p.Data[0])
}

// writeGridLine writes one horizontal boundary line, e.g. "┌──┬──┐":
// left separator, one segment per column joined by cross separators, right
// separator. Used for the top and bottom edge of the box.
func (p HeatmapPrinter) writeGridLine(buffer *bytes.Buffer, colWidth int, left, cross, right string) {
	buffer.WriteString(p.SeparatorStyle.Sprint(left))

	for i := 0; i < p.gridColumns(); i++ {
		if i > 0 {
			buffer.WriteString(p.SeparatorStyle.Sprint(cross))
		}

		buffer.WriteString(strings.Repeat(p.SeparatorStyle.Sprint(p.HorizontalSeparator), colWidth))
	}

	buffer.WriteString(p.SeparatorStyle.Sprint(right))
}

// writeRowSeparator writes the horizontal line between two grid rows,
// e.g. "\n├──┼──┤\n". The edge separators appear only when boxed, the
// segments and crosses only when the grid is enabled.
func (p HeatmapPrinter) writeRowSeparator(buffer *bytes.Buffer, colWidth int) {
	buffer.WriteString("\n")

	if p.Boxed {
		buffer.WriteString(p.SeparatorStyle.Sprint(p.LSeparator))
	}

	if p.Grid {
		for i := 0; i < p.gridColumns(); i++ {
			if i > 0 {
				buffer.WriteString(p.SeparatorStyle.Sprint(p.TCrossSeparator))
			}

			buffer.WriteString(strings.Repeat(p.SeparatorStyle.Sprint(p.HorizontalSeparator), colWidth))
		}
	}

	if p.Boxed {
		buffer.WriteString(p.SeparatorStyle.Sprint(p.LReverseSeparator))
	}

	if p.Grid {
		buffer.WriteString("\n")
	}
}

// writeRows writes one grid row per data row, prefixed with the Y-axis label
// when a header is rendered, separated by grid lines.
func (p HeatmapPrinter) writeRows(buffer *bytes.Buffer, colWidth int) {
	lastCol := len(p.Data[0]) - 1
	lastRow := len(p.Data) - 1

	for rowIdx, row := range p.Data {
		if p.Boxed {
			buffer.WriteString(p.SeparatorStyle.Sprint(p.VerticalSeparator))
		}

		if p.HasHeader {
			buffer.WriteString(p.AxisStyle.Sprint(padCell(p.Axis.YAxis[rowIdx], colWidth, false)))

			if p.Grid {
				buffer.WriteString(p.SeparatorStyle.Sprint(p.VerticalSeparator))
			}
		}

		for colIdx, value := range row {
			buffer.WriteString(p.sprintCell(value, colWidth))
			p.writeCellSeparator(buffer, colIdx == lastCol)
		}

		if rowIdx < lastRow {
			p.writeRowSeparator(buffer, colWidth)
		}
	}
}

// writeCellSeparator writes the vertical separator following a data cell.
// After the last cell of a row it becomes the closing border: with a header
// it requires both box and grid, without a header the box alone suffices
// (matching the historical layout).
func (p HeatmapPrinter) writeCellSeparator(buffer *bytes.Buffer, lastCol bool) {
	switch {
	case !lastCol:
		if p.Grid {
			buffer.WriteString(p.SeparatorStyle.Sprint(p.VerticalSeparator))
		}

	case p.HasHeader:
		if p.Boxed && p.Grid {
			buffer.WriteString(p.SeparatorStyle.Sprint(p.VerticalSeparator))
		}

	default:
		if p.Boxed {
			buffer.WriteString(p.SeparatorStyle.Sprint(p.VerticalSeparator))
		}
	}
}

// writeXAxisRow writes the last grid row: a blank cell above the Y-axis
// labels followed by the X-axis labels.
func (p HeatmapPrinter) writeXAxisRow(buffer *bytes.Buffer, colWidth int) {
	p.writeRowSeparator(buffer, colWidth)

	if p.Boxed {
		buffer.WriteString(p.SeparatorStyle.Sprint(p.VerticalSeparator))
	}

	buffer.WriteString(p.AxisStyle.Sprint(padCell(" ", colWidth, false)))

	if p.Grid {
		buffer.WriteString(p.SeparatorStyle.Sprint(p.VerticalSeparator))
	}

	lastCol := len(p.Axis.XAxis) - 1

	for colIdx, label := range p.Axis.XAxis {
		buffer.WriteString(p.AxisStyle.Sprint(padCell(label, colWidth, false)))

		if colIdx < lastCol {
			if p.Grid {
				buffer.WriteString(p.SeparatorStyle.Sprint(p.VerticalSeparator))
			}
		} else if p.Boxed && p.Grid {
			buffer.WriteString(p.SeparatorStyle.Sprint(p.VerticalSeparator))
		}
	}
}

// sprintCell formats and colors a single data cell. The cell text is the
// value itself, or a blank cell with OnlyColoredCells. Values wider than one
// character are right-aligned within the cell.
func (p HeatmapPrinter) sprintCell(value float32, colWidth int) string {
	text := " "
	if !p.OnlyColoredCells {
		text = Sprintf("%v", value)
	}

	cell := padCell(text, colWidth, len(Sprintf("%v", value)) > 1)

	if p.EnableRGB {
		return p.rgbStyleFor(value).Sprint(cell)
	}

	background := getColor(p.minValue, p.maxValue, value, p.Colors...)

	return p.textColorFor(background).Sprint(background.Sprint(cell))
}

// rgbStyleFor fades the value into the configured RGB range and pairs the
// resulting background with the text color (or its computed complement).
func (p HeatmapPrinter) rgbStyleFor(value float32) RGBStyle {
	background := p.RGBRange[0].Fade(p.minValue, p.maxValue, value, p.RGBRange[1:]...)

	if p.EnableComplementaryColor {
		return NewRGBStyle(NewRGB(internal.Complementary(background.R, background.G, background.B)), background)
	}

	return NewRGBStyle(p.TextRGB, background)
}

// textColorFor returns the text color to use on the given background color.
func (p HeatmapPrinter) textColorFor(background Color) Color {
	if p.EnableComplementaryColor {
		return complementaryColors[background]
	}

	return p.TextColor
}

// writeLegend writes the legend below the heatmap, boxed if the heatmap
// itself is boxed.
func (p HeatmapPrinter) writeLegend(buffer *bytes.Buffer, legendColWidth int) {
	buffer.WriteString("\n\n")

	if p.Boxed {
		p.writeBoxedLegend(buffer, legendColWidth)
	} else {
		p.writeLegendRow(buffer, legendColWidth)
	}
}

// writeBoxedLegend wraps the legend row in a box.
func (p HeatmapPrinter) writeBoxedLegend(buffer *bytes.Buffer, legendColWidth int) {
	buffer.WriteString(p.SeparatorStyle.Sprint(p.BottomRightCornerSeparator))
	p.writeLegendSeparatorRow(buffer, legendColWidth, true)
	buffer.WriteString(p.SeparatorStyle.Sprint(p.BottomLeftCornerSeparator))
	buffer.WriteString("\n")

	buffer.WriteString(p.SeparatorStyle.Sprint(p.VerticalSeparator))
	p.writeLegendRow(buffer, legendColWidth)
	buffer.WriteString(p.SeparatorStyle.Sprint(p.VerticalSeparator))
	buffer.WriteString("\n")

	buffer.WriteString(p.SeparatorStyle.Sprint(p.TopRightCornerSeparator))
	p.writeLegendSeparatorRow(buffer, legendColWidth, false)
	buffer.WriteString(p.SeparatorStyle.Sprint(p.TopLeftCornerSeparator))
}

// writeLegendRow writes the legend label followed by one colored cell per
// legend step.
func (p HeatmapPrinter) writeLegendRow(buffer *bytes.Buffer, legendColWidth int) {
	buffer.WriteString(p.AxisStyle.Sprint(p.LegendLabel))

	if p.Grid {
		buffer.WriteString(p.SeparatorStyle.Sprint(p.VerticalSeparator))
	} else {
		buffer.WriteString(" ")
	}

	if p.EnableRGB {
		p.writeRGBLegendCells(buffer, legendColWidth)
	} else {
		p.writeColorLegendCells(buffer, legendColWidth)
	}
}

// writeColorLegendCells writes one legend cell per configured color, fading
// linearly from the minimum to the maximum data value.
func (p HeatmapPrinter) writeColorLegendCells(buffer *bytes.Buffer, legendColWidth int) {
	for i, color := range p.Colors {
		value := p.legendValue(i, len(p.Colors))
		cell := centerAndShorten(value, legendColWidth, p.LegendOnlyColoredCells)
		buffer.WriteString(p.textColorFor(color).Sprint(color.Sprint(cell)))

		if p.Grid && i < len(p.Colors)-1 && !p.LegendOnlyColoredCells {
			buffer.WriteString(p.SeparatorStyle.Sprint(p.VerticalSeparator))
		}
	}
}

// writeRGBLegendCells writes the RGB legend cells. With LegendOnlyColoredCells
// the legend becomes a smooth gradient: three times as many cells, each one
// character wide and without values.
func (p HeatmapPrinter) writeRGBLegendCells(buffer *bytes.Buffer, legendColWidth int) {
	steps := max(len(p.RGBRange), rgbLegendSteps)
	cellWidth := legendColWidth

	if p.LegendOnlyColoredCells {
		steps *= 3
		cellWidth = 1
	}

	for i := 0; i < steps; i++ {
		value := p.legendValue(i, steps)
		buffer.WriteString(p.rgbStyleFor(value).Sprint(centerAndShorten(value, cellWidth, p.LegendOnlyColoredCells)))

		if p.Grid && i < steps-1 && !p.LegendOnlyColoredCells {
			buffer.WriteString(p.SeparatorStyle.Sprint(p.VerticalSeparator))
		}
	}
}

// writeLegendSeparatorRow writes a horizontal edge of the boxed legend. The
// first segment spans the legend label and is always followed by a cross
// separator; every further segment spans one legend cell, with crosses in
// between unless the legend is a gradient (LegendOnlyColoredCells).
func (p HeatmapPrinter) writeLegendSeparatorRow(buffer *bytes.Buffer, legendColWidth int, top bool) {
	cross := p.TReverseSeparator
	if top {
		cross = p.TSeparator
	}

	segments := len(p.Colors)
	segmentWidth := legendColWidth

	if p.EnableRGB {
		segments = max(len(p.RGBRange), rgbLegendSteps)
		if p.LegendOnlyColoredCells {
			// the gradient legend renders three one-character cells per step
			segmentWidth = 3
		}
	}

	buffer.WriteString(strings.Repeat(p.SeparatorStyle.Sprint(p.HorizontalSeparator), len(p.LegendLabel)))
	buffer.WriteString(p.SeparatorStyle.Sprint(cross))

	for i := 0; i < segments; i++ {
		buffer.WriteString(strings.Repeat(p.SeparatorStyle.Sprint(p.HorizontalSeparator), segmentWidth))

		if !p.LegendOnlyColoredCells && i < segments-1 {
			buffer.WriteString(p.SeparatorStyle.Sprint(cross))
		}
	}
}

// legendValue returns the value represented by legend cell i of steps, spread
// linearly from the minimum to the maximum data value.
func (p HeatmapPrinter) legendValue(i, steps int) float32 {
	switch i {
	case 0:
		return p.minValue
	case steps - 1:
		return p.maxValue
	default:
		return p.minValue + (p.maxValue-p.minValue)*float32(i)/float32(steps-1)
	}
}

// padCell centers text in a cell of the given width. internal.CenterText
// drops the remainder space on odd padding, so the missing column is re-added
// on the right, or on the left when padLeft is set (right-aligning the cell
// content). Widths are measured in bytes, consistent with CenterText.
func padCell(text string, width int, padLeft bool) string {
	cell := internal.CenterText(text, width)

	missing := width - len(cell)
	if missing <= 0 {
		return cell
	}

	if padLeft {
		return strings.Repeat(" ", missing) + cell
	}

	return cell + strings.Repeat(" ", missing)
}

// centerAndShorten formats a legend value into a cell of lineLength
// characters, reducing the value's precision if it does not fit. With
// onlyColor the cell stays blank and only carries the background color.
func centerAndShorten(f float32, lineLength int, onlyColor bool) string {
	value := ""
	if !onlyColor {
		value = Sprintf("%.2v", f)
	}

	if len(value) > lineLength {
		value = value[:lineLength]
		if strings.HasSuffix(value, ".") {
			value = Sprintf("%.1v", f)
			lineLength = len(value)
		}
	}

	return padCell(value, lineLength, len(Sprintf("%v", f)) > 1)
}

// getColor buckets current into len(colors) equal parts of the [minStep,
// maxStep] range: the first color covers the minimum, the last color the
// maximum (and, as the fallback, everything not matched by a bucket, e.g.
// when all values are equal).
func getColor(minStep float32, maxStep float32, current float32, colors ...Color) Color {
	step := (maxStep - minStep) / float32(len(colors))

	for i := range colors {
		if current >= minStep+float32(i)*step && current < minStep+float32(i+1)*step {
			return colors[i]
		}
	}

	return colors[len(colors)-1]
}

// Render prints the HeatmapPrinter to the terminal.
func (p HeatmapPrinter) Render() error {
	s, err := p.Srender()
	if err != nil {
		return err
	}

	Fprintln(p.Writer, s)

	return nil
}

func (p HeatmapPrinter) errCheck() error {
	if p.HasHeader {
		if p.Axis.XAxis == nil {
			return errors.New("x axis is nil")
		}

		if p.Axis.YAxis == nil {
			return errors.New("y axis is nil")
		}

		if len(p.Axis.XAxis) == 0 {
			return errors.New("x axis is empty")
		}

		if len(p.Axis.YAxis) == 0 {
			return errors.New("y axis is empty")
		}

		for _, row := range p.Data {
			if len(row) != len(p.Axis.XAxis) {
				return errors.New("x axis length does not match data")
			}
		}

		if len(p.Axis.YAxis) != len(p.Data) {
			return errors.New("y axis length does not match data")
		}
	}

	if p.Data == nil {
		return errors.New("data is nil")
	}

	if len(p.Data) == 0 {
		return errors.New("data is empty")
	}

	// check if p.Data[n] has the same length
	for i := 1; i < len(p.Data); i++ {
		if len(p.Data[i]) != len(p.Data[0]) {
			return errors.New("data is not rectangular")
		}
	}

	return nil
}

// minMaxFloat32 returns the smallest and largest value in the data.
func minMaxFloat32(s [][]float32) (float32, float32) {
	minValue := float32(math.MaxFloat32)
	maxValue := float32(-math.MaxFloat32)

	for _, row := range s {
		for _, value := range row {
			minValue = min(minValue, value)
			maxValue = max(maxValue, value)
		}
	}

	return minValue, maxValue
}
