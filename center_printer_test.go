package pterm_test

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

func TestCenterPrinter_WithCenterEachLineSeparately(t *testing.T) {
	p := pterm.CenterPrinter{}
	p2 := p.WithCenterEachLineSeparately()

	assert.True(t, p2.CenterEachLineSeparately)
	assert.False(t, p.CenterEachLineSeparately)
}

func TestCenterPrinterPrintMethods(t *testing.T) {
	p := pterm.DefaultCenter

	t.Run("Print", func(t *testing.T) {
		testPrintContains(t, func(_ io.Writer, a any) {
			p.Print(a)
		})
	})

	t.Run("Printf", func(t *testing.T) {
		testPrintfContains(t, func(_ io.Writer, format string, a any) {
			p.Printf(format, a)
		})
	})

	t.Run("Printfln", func(t *testing.T) {
		testPrintflnContains(t, func(_ io.Writer, format string, a any) {
			p.Printfln(format, a)
		})
	})

	t.Run("Println", func(t *testing.T) {
		testPrintlnContains(t, func(_ io.Writer, a any) {
			p.Println(a)
		})
	})

	t.Run("Sprint", func(t *testing.T) {
		testSprintContains(t, func(a any) string {
			return p.Sprint(a)
		})
	})

	t.Run("Sprintf", func(t *testing.T) {
		testSprintfContains(t, func(format string, a any) string {
			return p.Sprintf(format, a)
		})
	})

	t.Run("Sprintfln", func(t *testing.T) {
		testSprintflnContains(t, func(format string, a any) string {
			return p.Sprintfln(format, a)
		})
	})

	t.Run("Sprintln", func(t *testing.T) {
		testSprintlnContains(t, func(a any) string {
			return p.Sprintln(a)
		})
	})

	t.Run("PrintOnError", func(t *testing.T) {
		result := captureStdout(func(_ io.Writer) {
			p.PrintOnError(errors.New("hello world"))
		})
		assert.Contains(t, result, "hello world")
	})

	t.Run("PrintIfError_WithoutError", func(t *testing.T) {
		result := captureStdout(func(_ io.Writer) {
			p.PrintOnError(nil)
		})
		assert.Zero(t, result)
	})

	t.Run("PrintOnErrorf", func(t *testing.T) {
		result := captureStdout(func(_ io.Writer) {
			p.PrintOnErrorf("wrapping error : %w", errors.New("hello world"))
		})
		assert.Contains(t, result, "hello world")
	})

	t.Run("PrintIfError_WithoutErrorf", func(t *testing.T) {
		result := captureStdout(func(_ io.Writer) {
			p.PrintOnErrorf("", nil)
		})
		assert.Zero(t, result)
	})
}

func TestCenterPrinterPrintMethodsCenterSeparately(t *testing.T) {
	p := pterm.DefaultCenter.WithCenterEachLineSeparately()

	t.Run("Print", func(t *testing.T) {
		testPrintContains(t, func(_ io.Writer, a any) {
			p.Print(a)
		})
	})

	t.Run("Printf", func(t *testing.T) {
		testPrintfContains(t, func(_ io.Writer, format string, a any) {
			p.Printf(format, a)
		})
	})

	t.Run("Println", func(t *testing.T) {
		testPrintlnContains(t, func(_ io.Writer, a any) {
			p.Println(a)
		})
	})

	t.Run("Sprint", func(t *testing.T) {
		testSprintContains(t, func(a any) string {
			return p.Sprint(a)
		})
	})

	t.Run("Sprintf", func(t *testing.T) {
		testSprintfContains(t, func(format string, a any) string {
			return p.Sprintf(format, a)
		})
	})

	t.Run("Sprintln", func(t *testing.T) {
		testSprintlnContains(t, func(a any) string {
			return p.Sprintln(a)
		})
	})
}

func TestCenterPrinter_SprintLineLongerThanTerminal(t *testing.T) {
	p := pterm.DefaultCenter
	p2 := p.Sprint("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	assert.Contains(t, p2, "a")
}

func TestCenterPrinter_SprintLineLongerThanTerminalWithCenterEachLineSeparately(t *testing.T) {
	p := pterm.DefaultCenter
	p2 := p.WithCenterEachLineSeparately().Sprint("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	assert.Contains(t, p2, "a")
}

func TestCenterPrinter_WithWriter(t *testing.T) {
	p := pterm.CenterPrinter{}
	s := os.Stderr
	p2 := p.WithWriter(s)

	assert.Equal(t, s, p2.Writer)
	assert.Zero(t, p.Writer)
}
