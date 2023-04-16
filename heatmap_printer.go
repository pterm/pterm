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
	Boxed:                      true,
	Grid:                       true,
	TextRGB:                    RGB{0, 0, 0, false},
	RGBRange:                   []RGB{{R: 255, G: 0, B: 0, Background: true}, {R: 255, G: 165, B: 0, Background: true}, {R: 0, G: 255, B: 0, Background: true}},
	TextColor:                  FgBlack,
	Colors:                     []Color{BgRed, BgLightRed, BgYellow, BgLightYellow, BgLightGreen, BgGreen},

	IsRGB: false,
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
	SeparatorStyle             *Style
	Data                       HeatmapData
	Axis                       HeatmapAxis
	Boxed                      bool
	Grid                       bool
	OnlyColoredCells           bool
	ComplementColor            bool
	CellSize                   int
	Colors                     []Color
	TextColor                  Color
	IsRGB                      bool
	RGBRange                   []RGB
	TextRGB                    RGB
	Writer                     io.Writer

	minValue float32
	maxValue float32
}

var colorComplement = map[Color]Color{
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
func (p HeatmapPrinter) WithTextColor(color Color) *HeatmapPrinter {
	p.TextColor = color
	p.IsRGB = false
	p.ComplementColor = false
	return &p
}

// WithTextRGB returns a new HeatmapPrinter with a specific TextRGB.
func (p HeatmapPrinter) WithTextRGB(rgb RGB) *HeatmapPrinter {
	p.TextRGB = rgb
	p.IsRGB = true
	p.ComplementColor = false
	return &p
}

// WithBoxed returns a new HeatmapPrinter with a box around the table.
func (p HeatmapPrinter) WithBoxed(b ...bool) *HeatmapPrinter {
	p.Boxed = internal.WithBoolean(b)
	return &p
}

// WithGrid returns a new HeatmapPrinter with a grid.
func (p HeatmapPrinter) WithGrid(b ...bool) *HeatmapPrinter {
	b2 := internal.WithBoolean(b)
	p.Grid = b2
	if !b2 && p.Boxed {
		p.Boxed = false
	}
	return &p
}

// WithRGB returns a new HeatmapPrinter with RGB colors.
func (p HeatmapPrinter) WithRGB(b ...bool) *HeatmapPrinter {
	p.IsRGB = internal.WithBoolean(b)
	return &p
}

// WithOnlyColoredCells returns a new HeatmapPrinter with only colored cells.
func (p HeatmapPrinter) WithOnlyColoredCells(b ...bool) *HeatmapPrinter {
	b2 := internal.WithBoolean(b)
	p.OnlyColoredCells = b2
	return &p
}

// WithComplementColor returns a new HeatmapPrinter with complement color.
func (p HeatmapPrinter) WithComplementColor(b ...bool) *HeatmapPrinter {
	p.ComplementColor = internal.WithBoolean(b)
	return &p
}

// WithCellSize returns a new HeatmapPrinter with a specific cell size.
// This only works if there is no header and OnlyColoredCells == true!
func (p HeatmapPrinter) WithCellSize(i int) *HeatmapPrinter {
	p.CellSize = i
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
		p.SeparatorStyle = NewStyle()
	}
	if p.AxisStyle == nil {
		p.AxisStyle = NewStyle()
	}

	var buffer bytes.Buffer
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
	}

	lineLength := internal.GetStringMaxWidth(data)

	if p.OnlyColoredCells && (p.CellSize > lineLength || !p.HasHeader) {
		lineLength = p.CellSize
	}

	if p.Boxed {
		buffer.WriteString(p.SeparatorStyle.Sprint(p.BottomRightCornerSeparator))
		for i := 0; i < xAmount+1; i++ {
			buffer.WriteString(strings.Repeat(p.SeparatorStyle.Sprint(p.HorizontalSeparator), lineLength))
			if i < xAmount {
				buffer.WriteString(p.SeparatorStyle.Sprint(p.TSeparator))
			}
		}
		buffer.WriteString(p.SeparatorStyle.Sprint(p.BottomLeftCornerSeparator))
		buffer.WriteString("\n")
	}

	for i, datum := range p.Data {
		if p.Boxed {
			buffer.WriteString(p.SeparatorStyle.Sprint(p.VerticalSeparator))
		}
		for j, f := range datum {
			if j == 0 && p.HasHeader {
				ct := internal.CenterText(p.Axis.YAxis[i], lineLength)
				if len(ct) < lineLength {
					ct += strings.Repeat(" ", lineLength-len(ct))
				}
				buffer.WriteString(p.AxisStyle.Sprint(ct))
				if p.Grid {
					buffer.WriteString(p.SeparatorStyle.Sprint(p.VerticalSeparator))
				}
			}
			var ct string
			if p.OnlyColoredCells {
				ct = internal.CenterText(" ", lineLength)
			} else {
				ct = internal.CenterText(Sprintf("%v", f), lineLength)
			}
			if len(ct) < lineLength {
				ct += strings.Repeat(" ", lineLength-len(ct))
			}
			if p.IsRGB {
				rgb := p.RGBRange[0].Fade(p.minValue, p.maxValue, f, p.RGBRange[1:]...)
				rgbStyle := NewRGBStyle(p.TextRGB, rgb)
				if p.ComplementColor {
					complimentary := NewRGB(internal.Complementary(rgb.R, rgb.G, rgb.B))
					rgbStyle = NewRGBStyle(complimentary, rgb)
				}
				buffer.WriteString(rgbStyle.Sprint(ct))
			} else {
				color := getColor(p.minValue, p.maxValue, f, p.Colors...)
				fgColor := p.TextColor
				if p.ComplementColor {
					fgColor = colorComplement[color]
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
					buffer.WriteString(strings.Repeat(p.SeparatorStyle.Sprint(p.HorizontalSeparator), lineLength))
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

	if p.HasHeader {
		buffer.WriteString("\n")
		if p.Boxed {
			buffer.WriteString(p.SeparatorStyle.Sprint(p.LSeparator))
		}
		if p.Grid {
			for i := 0; i < xAmount+1; i++ {
				buffer.WriteString(strings.Repeat(p.SeparatorStyle.Sprint(p.HorizontalSeparator), lineLength))
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
				ct := internal.CenterText(" ", lineLength)
				if len(ct) < lineLength {
					ct += strings.Repeat(" ", lineLength-len(ct))
				}
				buffer.WriteString(p.AxisStyle.Sprint(ct))
				if p.Grid {
					buffer.WriteString(p.SeparatorStyle.Sprint(p.VerticalSeparator))
				}
			}
			var ct string
			ct = internal.CenterText(Sprintf("%v", f), lineLength)
			if len(ct) < lineLength {
				ct += strings.Repeat(" ", lineLength-len(ct))
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

	if p.Boxed {
		buffer.WriteString("\n")
		buffer.WriteString(p.SeparatorStyle.Sprint(p.TopRightCornerSeparator))
		for i := 0; i < xAmount+1; i++ {
			buffer.WriteString(strings.Repeat(p.SeparatorStyle.Sprint(p.HorizontalSeparator), lineLength))
			if i < xAmount {
				buffer.WriteString(p.SeparatorStyle.Sprint(p.TReverseSeparator))
			}
		}
		buffer.WriteString(p.SeparatorStyle.Sprint(p.TopLeftCornerSeparator))
	}
	buffer.WriteString("\n")

	return buffer.String(), nil
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
