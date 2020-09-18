package pterm

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
)

var (
	Print   = color.Print
	Printf  = color.Printf
	Println = color.Println

	Sprint  = color.Sprint
	Sprintf = color.Sprintf
)

// Foreground colors. basic foreground colors 30 - 37
const (
	FgBlack Color = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta // 品红
	FgCyan    // 青色
	FgWhite
	// FgDefault revert default FG
	FgDefault Color = 39
)

// Extra foreground color 90 - 97(非标准)
const (
	FgDarkGray Color = iota + 90 // 亮黑（灰）
	FgLightRed
	FgLightGreen
	FgLightYellow
	FgLightBlue
	FgLightMagenta
	FgLightCyan
	FgLightWhite
	// FgGray is alias of FgDarkGray
	FgGray Color = 90 // 亮黑（灰）
)

// Background colors. basic background colors 40 - 47
const (
	BgBlack Color = iota + 40
	BgRed
	BgGreen
	BgYellow // BgBrown like yellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
	// BgDefault revert default BG
	BgDefault Color = 49
)

// Extra background color 100 - 107(非标准)
const (
	BgDarkGray Color = iota + 100
	BgLightRed
	BgLightGreen
	BgLightYellow
	BgLightBlue
	BgLightMagenta
	BgLightCyan
	BgLightWhite
	// BgGray is alias of BgDarkGray
	BgGray Color = 100
)

// Option settings
const (
	Reset         Color = iota // 0 重置所有设置
	Bold                       // 1 加粗
	Fuzzy                      // 2 模糊(不是所有的终端仿真器都支持)
	Italic                     // 3 斜体(不是所有的终端仿真器都支持)
	Underscore                 // 4 下划线
	Blink                      // 5 闪烁
	FastBlink                  // 5 快速闪烁(未广泛支持)
	Reverse                    // 7 颠倒的 交换背景色与前景色
	Concealed                  // 8 隐匿的
	Strikethrough              // 9 删除的，删除线(未广泛支持)
)

var (
	Red     = FgRed.Sprint
	Cyan    = FgCyan.Sprint
	Gray    = FgGray.Sprint
	Blue    = FgBlue.Sprint
	Black   = FgBlack.Sprint
	Green   = FgGreen.Sprint
	White   = FgWhite.Sprint
	Yellow  = FgYellow.Sprint
	Magenta = FgMagenta.Sprint

	// special

	Normal = color.Normal.Sprint

	// extra light

	LightRed     = FgLightRed.Sprint
	LightCyan    = FgLightCyan.Sprint
	LightBlue    = FgLightBlue.Sprint
	LightGreen   = FgLightGreen.Sprint
	LightWhite   = FgLightWhite.Sprint
	LightYellow  = FgLightYellow.Sprint
	LightMagenta = FgLightMagenta.Sprint
)

type Color uint8

func (c Color) Sprintln(a ...interface{}) string {
	return color.RenderCode(color.Color(c).String(), a...) + "\n"
}

func (c Color) Sprint(a ...interface{}) string {
	return color.RenderCode(color.Color(c).String(), a...)
}

func (c Color) Sprintf(format string, a ...interface{}) string {
	return color.RenderString(color.Color(c).String(), fmt.Sprintf(format, a...))
}

func (c Color) Println(a ...interface{}) {
	Print(c.Sprintln(a...))
}

func (c Color) Print(a ...interface{}) {
	Print(c.Sprint(a...))
}

func (c Color) Printf(format string, a ...interface{}) {
	Print(c.Sprintf(format, a))
}

// String converts the color a string. eg "35"
func (c Color) String() string {
	return fmt.Sprintf("%d", c)
}

type Style []Color

func New(colors ...Color) Style {
	return colors
}

func (s Style) Sprint(a ...interface{}) string {
	return color.RenderCode(s.String(), a...)
}

func (s Style) Sprintln(a ...interface{}) string {
	return color.RenderCode(s.String(), a...) + "\n"
}

func (s Style) Sprintf(format string, a ...interface{}) string {
	return color.RenderString(s.String(), fmt.Sprintf(format, a...))
}

func (s Style) Print(a ...interface{}) {
	Print(s.Sprint(a...))
}

func (s Style) Println(a ...interface{}) {
	Println(s.Sprint(a...))
}

func (s Style) Printf(format string, a ...interface{}) {
	Print(s.Sprintf(format, a...))
}

// Code convert to code string. returns like "32;45;3"
func (s Style) Code() string {
	return s.String()
}

// String convert to code string. returns like "32;45;3"
func (s Style) String() string {
	return colors2code(s...)
}

// convert colors to code. return like "32;45;3"
func colors2code(colors ...Color) string {
	if len(colors) == 0 {
		return ""
	}

	var codes []string
	for _, c := range colors {
		codes = append(codes, c.String())
	}

	return strings.Join(codes, ";")
}
