package pterm

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm/internal"
)

func TestBoxPrinterNilPrint(t *testing.T) {
	proxyToDevNull()
	BoxPrinter{}.Render()
	BoxPrinter{}.WithText("abc").Render()
}

func TestBoxPrinter_Render(t *testing.T) {
	internal.TestPrintContains(t, func(w io.Writer, a interface{}) {
		DefaultBox.WithText(fmt.Sprint(a) + "\nabc\nHello, World!").Render()
	})
}

func TestBoxPrinter_WithBottomLeftCornerString(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithBottomLeftCornerString("-")

	assert.Equal(t, "-", p2.BottomLeftCornerString)
	assert.Empty(t, p.BottomLeftCornerString)
}

func TestBoxPrinter_WithBottomPadding(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithBottomPadding(5)

	assert.Equal(t, 5, p2.BottomPadding)
	assert.Empty(t, p.BottomPadding)
}

func TestBoxPrinter_WithBottomRightCornerString(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithBottomRightCornerString("-")

	assert.Equal(t, "-", p2.BottomRightCornerString)
	assert.Empty(t, p.BottomRightCornerString)
}

func TestBoxPrinter_WithBoxStyle(t *testing.T) {
	p := BoxPrinter{}
	s := NewStyle(FgRed, BgRed, Bold)
	p2 := p.WithBoxStyle(s)

	assert.Equal(t, s, p2.BoxStyle)
	assert.Empty(t, p.BoxStyle)
}

func TestBoxPrinter_WithLeftPadding(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithLeftPadding(5)

	assert.Equal(t, 5, p2.LeftPadding)
	assert.Empty(t, p.LeftPadding)
}

func TestBoxPrinter_WithRightPadding(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithRightPadding(5)

	assert.Equal(t, 5, p2.RightPadding)
	assert.Empty(t, p.RightPadding)
}

func TestBoxPrinter_WithText(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithText("-")

	assert.Equal(t, "-", p2.Text)
	assert.Empty(t, p.Text)
}

func TestBoxPrinter_WithTextStyle(t *testing.T) {
	p := BoxPrinter{}
	s := NewStyle(FgRed, BgRed, Bold)
	p2 := p.WithTextStyle(s)

	assert.Equal(t, s, p2.TextStyle)
	assert.Empty(t, p.TextStyle)
}

func TestBoxPrinter_WithTopLeftCornerString(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithTopLeftCornerString("-")

	assert.Equal(t, "-", p2.TopLeftCornerString)
	assert.Empty(t, p.TopLeftCornerString)
}

func TestBoxPrinter_WithTopPadding(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithTopPadding(5)

	assert.Equal(t, 5, p2.TopPadding)
	assert.Empty(t, p.TopPadding)
}

func TestBoxPrinter_WithInvalidTopPadding(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithTopPadding(-5)

	assert.Equal(t, 0, p2.TopPadding)
	assert.Empty(t, p.TopPadding)
}

func TestBoxPrinter_WithInvalidBottomPadding(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithBottomPadding(-5)

	assert.Equal(t, 0, p2.BottomPadding)
	assert.Empty(t, p.BottomPadding)
}

func TestBoxPrinter_WithInvalidLeftPadding(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithLeftPadding(-5)

	assert.Equal(t, 0, p2.LeftPadding)
	assert.Empty(t, p.LeftPadding)
}

func TestBoxPrinter_WithInvalidRightPadding(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithRightPadding(-5)

	assert.Equal(t, 0, p2.RightPadding)
	assert.Empty(t, p.RightPadding)
}

func TestBoxPrinter_WithTopRightCornerString(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithTopRightCornerString("-")

	assert.Equal(t, "-", p2.TopRightCornerString)
	assert.Empty(t, p.TopRightCornerString)
}

func TestBoxPrinter_WithVerticalString(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithVerticalString("-")

	assert.Equal(t, "-", p2.VerticalString)
	assert.Empty(t, p.VerticalString)
}
