package pterm

import (
	"bytes"
	"errors"
	"io"
	"math"
	"strings"

	"github.com/pterm/pterm/internal"
)

// DefaultHeatmap contains standards, which can be used to print a HeatmapPrinter.
var DefaultHeatmap = HeatmapPrinter{
	AxisStyle:                  &ThemeDefault.HeatmapHeaderStyle,
	SeparatorStyle:             &ThemeDefault.HeatmapSeparatorStyle,
	VerticalSeparator:          "│",
	TopRightCornerSeparator:    "└",
	TopLeftCornerSeparator:     "┘",
	BottomLeftCornerSeparator:  "┐",
	BottomRightCornerSeparator: "┌",
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
	TextRGB:                    RGB{0, 0, 0, false},
	RGBRange:                   []RGB{{R: 255, G: 0, B: 0, Background: true}, {R: 255, G: 165, B: 0, Background: true}, {R: 0, G: 255, B: 0, Background: true}},
	TextColor:                  FgBlack,
	Colors:                     []Color{BgRed, BgLightRed, BgYellow, BgLightYellow, BgLightGreen, BgGreen},

	EnableRGB: false,
}

// HeatmapData is the type that contains the data of a HeatmapPrinter.
type HeatmapData [][]float32

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

	rgbLegendValue int
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

	if RawOutput {
		p.Legend = false
	}

	buffer := bytes.NewBufferString("")
	xAmount := len(p.Data[0]) - 1
	yAmount := len(p.Data) - 1
	p.minValue, p.maxValue = minMaxFloat32(p.Data)

	var data string
	for _, datum := range p.Data {
		for _, f := range datum {
			data += Sprintf("%v\n", f)
		}
	}

	if p.HasHeader {
		data, xAmount, yAmount = p.computeAxisData(data, xAmount, yAmount)
	}

	colWidth := internal.GetStringMaxWidth(data)
	legendColWidth := colWidth + 2

	if p.OnlyColoredCells && (p.CellSize > colWidth || !p.HasHeader) {
		colWidth = p.CellSize
	}

	if p.Boxed {
		p.renderSeparatorRow(buffer, colWidth, xAmount, true)
	}

	p.renderData(buffer, colWidth, xAmount, yAmount)

	if p.HasHeader {
		p.renderHeader(buffer, colWidth, xAmount)
	}

	if p.Boxed {
		p.renderSeparatorRow(buffer, colWidth, xAmount, false)
	}

	if p.Legend {
		p.renderLegend(buffer, legendColWidth)
	}

	buffer.WriteString("\n")

	return buffer.String(), nil
}

func (p HeatmapPrinter) computeAxisData(data string, xAmount, yAmount int) (string, int, int) {
	var header string
	for _, h := range p.Axis.XAxis {
		header += h + "\n"
	}
	for _, h := range p.Axis.YAxis {
		header += h + "\n"
	}

	if p.OnlyColoredCells {
		data = header
	} else {
		data += header
	}
	xAmount++
	yAmount++

	p.Axis.YAxis = append(p.Axis.YAxis, "")

	return data, xAmount, yAmount
}

func (p HeatmapPrinter) renderSeparatorRow(buffer *bytes.Buffer, colWidth, xAmount int, top bool) {
	tSep := p.TReverseSeparator
	rightSep := p.TopRightCornerSeparator
	leftSep := p.TopLeftCornerSeparator

	if top {
		tSep = p.TSeparator
		rightSep = p.BottomRightCornerSeparator
		leftSep = p.BottomLeftCornerSeparator
	} else {
		buffer.WriteString("\n")
	}
	buffer.WriteString(p.SeparatorStyle.Sprint(rightSep))
	for i := 0; i < xAmount+1; i++ {
		buffer.WriteString(strings.Repeat(p.SeparatorStyle.Sprint(p.HorizontalSeparator), colWidth))
		if i < xAmount {
			buffer.WriteString(p.SeparatorStyle.Sprint(tSep))
		}
	}
	buffer.WriteString(p.SeparatorStyle.Sprint(leftSep))

	if top {
		buffer.WriteString("\n")
	}
}

func (p HeatmapPrinter) renderLegend(buffer *bytes.Buffer, legendColWidth int) {
	buffer.WriteString("\n")
	buffer.WriteString("\n")
	if p.Boxed {
		p.boxLegend(buffer, p.LegendLabel, legendColWidth)
	} else {
		p.generateLegend(buffer, p.LegendLabel, legendColWidth)
	}
}

func (p HeatmapPrinter) renderHeader(buffer *bytes.Buffer, colWidth int, xAmount int) {
	buffer.WriteString("\n")
	if p.Boxed {
		buffer.WriteString(p.SeparatorStyle.Sprint(p.LSeparator))
	}
	if p.Grid {
		for i := 0; i < xAmount+1; i++ {
			buffer.WriteString(strings.Repeat(p.SeparatorStyle.Sprint(p.HorizontalSeparator), colWidth))
			if i < xAmount {
				buffer.WriteString(p.SeparatorStyle.Sprint(p.TCrossSeparator))
			}
		}
	}
	if p.Boxed {
		buffer.WriteString(p.SeparatorStyle.Sprint(p.LReverseSeparator))
	}
	if p.Grid {
		buffer.WriteString("\n")
	}
	for j, f := range p.Axis.XAxis {
		if j == 0 {
			if p.Boxed {
				buffer.WriteString(p.SeparatorStyle.Sprint(p.VerticalSeparator))
			}
			ct := internal.CenterText(" ", colWidth)
			if len(ct) < colWidth {
				ct += strings.Repeat(" ", colWidth-len(ct))
			}
			buffer.WriteString(p.AxisStyle.Sprint(ct))
			if p.Grid {
				buffer.WriteString(p.SeparatorStyle.Sprint(p.VerticalSeparator))
			}
		}
		var ct string
		ct = internal.CenterText(Sprintf("%v", f), colWidth)
		if len(ct) < colWidth {
			ct += strings.Repeat(" ", colWidth-len(ct))
		}
		buffer.WriteString(p.AxisStyle.Sprint(ct))

		if j < xAmount {
			if !p.Boxed && j == xAmount-1 {
				continue
			}
			if p.Grid {
				buffer.WriteString(p.SeparatorStyle.Sprint(p.VerticalSeparator))
			}
		}
	}
}

func (p HeatmapPrinter) renderData(buffer *bytes.Buffer, colWidth int, xAmount int, yAmount int) {
	for i, datum := range p.Data {
		if p.Boxed {
			buffer.WriteString(p.SeparatorStyle.Sprint(p.VerticalSeparator))
		}
		for j, f := range datum {
			if j == 0 && p.HasHeader {
				ct := internal.CenterText(p.Axis.YAxis[i], colWidth)
				if len(ct) < colWidth {
					ct += strings.Repeat(" ", colWidth-len(ct))
				}
				buffer.WriteString(p.AxisStyle.Sprint(ct))
				if p.Grid {
					buffer.WriteString(p.SeparatorStyle.Sprint(p.VerticalSeparator))
				}
			}
			var ct string
			if p.OnlyColoredCells {
				ct = internal.CenterText(" ", colWidth)
			} else {
				ct = internal.CenterText(Sprintf("%v", f), colWidth)
			}
			if len(ct) < colWidth {
				if len(Sprintf("%v", f)) == 1 {
					ct += strings.Repeat(" ", colWidth-len(ct))
				} else {
					ct = strings.Repeat(" ", colWidth-len(ct)) + ct
				}
			}
			if p.EnableRGB {
				rgb := p.RGBRange[0].Fade(p.minValue, p.maxValue, f, p.RGBRange[1:]...)
				rgbStyle := NewRGBStyle(p.TextRGB, rgb)
				if p.EnableComplementaryColor {
					complimentary := NewRGB(internal.Complementary(rgb.R, rgb.G, rgb.B))
					rgbStyle = NewRGBStyle(complimentary, rgb)
				}
				buffer.WriteString(rgbStyle.Sprint(ct))
			} else {
				color := getColor(p.minValue, p.maxValue, f, p.Colors...)
				fgColor := p.TextColor
				if p.EnableComplementaryColor {
					fgColor = complementaryColors[color]
				}
				buffer.WriteString(fgColor.Sprint(color.Sprintf(ct)))
			}
			if j < xAmount {
				if !p.Boxed && p.HasHeader && j == xAmount-1 {
					continue
				}
				if p.Grid {
					buffer.WriteString(p.SeparatorStyle.Sprint(p.VerticalSeparator))
				}
			}
			if p.Boxed && !p.HasHeader && j == xAmount {
				buffer.WriteString(p.SeparatorStyle.Sprint(p.VerticalSeparator))
			}
		}

		if i < yAmount {
			if p.HasHeader && i == yAmount-1 {
				continue
			}
			buffer.WriteString("\n")
			if p.Boxed {
				buffer.WriteString(p.SeparatorStyle.Sprint(p.LSeparator))
			}
			if p.Grid {
				for i := 0; i < xAmount+1; i++ {
					buffer.WriteString(strings.Repeat(p.SeparatorStyle.Sprint(p.HorizontalSeparator), colWidth))
					if i < xAmount {
						buffer.WriteString(p.SeparatorStyle.Sprint(p.TCrossSeparator))
					}
				}
			}
			if p.Boxed {
				buffer.WriteString(p.SeparatorStyle.Sprint(p.LReverseSeparator))
			}
			if p.Grid {
				buffer.WriteString("\n")
			}
		}
	}
}

func (p HeatmapPrinter) generateLegend(buffer *bytes.Buffer, legend string, legendColWidth int) {
	buffer.WriteString(p.AxisStyle.Sprint(legend))
	if p.Grid {
		buffer.WriteString(p.SeparatorStyle.Sprintf("%s", p.VerticalSeparator))
	} else {
		buffer.WriteString(" ")
	}
	if p.EnableRGB {
		p.generateRGBLegend(buffer, legendColWidth)
	} else {
		p.generateColorLegend(buffer, legendColWidth)
	}
}

func (p HeatmapPrinter) generateColorLegend(buffer *bytes.Buffer, legendColWidth int) {
	for i, color := range p.Colors {
		// the first color is the min value and the last color is the max value
		var f float32
		if i == 0 {
			f = p.minValue
		} else if i == len(p.Colors)-1 {
			f = p.maxValue
		} else {
			f = p.minValue + (p.maxValue-p.minValue)*float32(i)/float32(len(p.Colors)-1)
		}
		fgColor := p.TextColor
		if p.EnableComplementaryColor {
			fgColor = complementaryColors[color]
		}
		buffer.WriteString(fgColor.Sprint(color.Sprint(centerAndShorten(f, legendColWidth, p.LegendOnlyColoredCells))))
		if p.Grid && i < len(p.Colors)-1 && !p.LegendOnlyColoredCells {
			buffer.WriteString(p.SeparatorStyle.Sprintf("%s", p.VerticalSeparator))
		}
	}
}

func (p HeatmapPrinter) generateRGBLegend(buffer *bytes.Buffer, legendColWidth int) {
	p.rgbLegendValue = 10
	steps := len(p.RGBRange)
	if steps < p.rgbLegendValue {
		steps = p.rgbLegendValue
	}
	if p.LegendOnlyColoredCells {
		steps *= 3
	}
	for i := 0; i < steps; i++ {
		// the first color is the min value and the last color is the max value
		var f float32
		if i == 0 {
			f = p.minValue
		} else if i == steps-1 {
			f = p.maxValue
		} else {
			f = p.minValue + (p.maxValue-p.minValue)*float32(i)/float32(steps-1)
		}
		rgb := p.RGBRange[0].Fade(p.minValue, p.maxValue, f, p.RGBRange[1:]...)
		rgbStyle := NewRGBStyle(p.TextRGB, rgb)
		if p.EnableComplementaryColor {
			complimentary := NewRGB(internal.Complementary(rgb.R, rgb.G, rgb.B))
			rgbStyle = NewRGBStyle(complimentary, rgb)
		}
		if p.LegendOnlyColoredCells {
			buffer.WriteString(rgbStyle.Sprint(centerAndShorten(f, 1, p.LegendOnlyColoredCells)))
		} else {
			buffer.WriteString(rgbStyle.Sprint(centerAndShorten(f, legendColWidth, p.LegendOnlyColoredCells)))
		}
		if p.Grid && i < steps-1 && !p.LegendOnlyColoredCells {
			buffer.WriteString(p.SeparatorStyle.Sprintf("%s", p.VerticalSeparator))
		}
	}
}

func (p HeatmapPrinter) boxLegend(buffer *bytes.Buffer, legend string, legendColWidth int) {
	buffer.WriteString(p.SeparatorStyle.Sprint(p.BottomRightCornerSeparator))

	p.generateSeparatorRow(buffer, legend, legendColWidth, true)

	buffer.WriteString(p.SeparatorStyle.Sprint(p.BottomLeftCornerSeparator))
	buffer.WriteString("\n")
	buffer.WriteString(p.SeparatorStyle.Sprintf("%s", p.VerticalSeparator))

	p.generateLegend(buffer, legend, legendColWidth)

	buffer.WriteString(p.SeparatorStyle.Sprintf("%s", p.VerticalSeparator))
	buffer.WriteString("\n")

	buffer.WriteString(p.SeparatorStyle.Sprint(p.TopRightCornerSeparator))

	p.generateSeparatorRow(buffer, legend, legendColWidth, false)

	buffer.WriteString(p.SeparatorStyle.Sprint(p.TopLeftCornerSeparator))
}

func (p HeatmapPrinter) generateSeparatorRow(buffer *bytes.Buffer, legend string, legendColWidth int, top bool) {
	p.rgbLegendValue = 10
	steps := len(p.RGBRange)
	if steps < p.rgbLegendValue {
		steps = p.rgbLegendValue
	}
	if p.LegendOnlyColoredCells {
		steps *= 3
	}

	var xValue int
	if p.EnableRGB {
		xValue = len(p.RGBRange)
		if xValue < p.rgbLegendValue {
			xValue = p.rgbLegendValue
		}
	} else {
		xValue = len(p.Colors)
	}

	for i := 0; i < xValue+1; i++ {
		if i == 0 {
			firstLength := len(legend)
			buffer.WriteString(strings.Repeat(p.SeparatorStyle.Sprint(p.HorizontalSeparator), firstLength))
		} else {
			if p.LegendOnlyColoredCells {
				if p.EnableRGB {
					buffer.WriteString(strings.Repeat(p.SeparatorStyle.Sprint(p.HorizontalSeparator), steps/(xValue)))
				} else {
					buffer.WriteString(strings.Repeat(p.SeparatorStyle.Sprint(p.HorizontalSeparator), legendColWidth))
				}
			} else {
				buffer.WriteString(strings.Repeat(p.SeparatorStyle.Sprint(p.HorizontalSeparator), legendColWidth))
			}
		}
		if i < xValue && !p.LegendOnlyColoredCells || i == 0 {
			if top {
				buffer.WriteString(p.SeparatorStyle.Sprint(p.TSeparator))
			} else {
				buffer.WriteString(p.SeparatorStyle.Sprint(p.TReverseSeparator))
			}
		}
	}
}

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
	ct := internal.CenterText(value, lineLength)
	if len(ct) < lineLength {
		if len(Sprintf("%v", f)) == 1 {
			ct += strings.Repeat(" ", lineLength-len(ct))
		} else {
			ct = strings.Repeat(" ", lineLength-len(ct)) + ct
		}
	}

	return ct
}

func getColor(min float32, max float32, current float32, colors ...Color) Color {
	// split the range into equal parts
	// and assign a color to each part
	// the last color is assigned to the max value
	// and the first color to the min value
	// the rest of the colors are assigned to the
	// middle values
	step := (max - min) / float32(len(colors))
	for i := range colors {
		if current >= min+float32(i)*step && current < min+float32(i+1)*step {
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

		for i := 1; i < len(p.Data); i++ {
			if len(p.Data[i]) != len(p.Axis.XAxis) {
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

// return min and max value of a slice
func minMaxFloat32(s [][]float32) (float32, float32) {
	var min, max float32
	min = math.MaxFloat32
	max = -math.MaxFloat32

	for _, r := range s {
		for _, c := range r {
			if c < min {
				min = c
			}
			if c > max {
				max = c
			}
		}
	}
	return min, max
}
