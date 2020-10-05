package test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

func TestGenericPrinter(t *testing.T) {
	var genericPrinters []pterm.GenericPrinter

	prefixPrinter := []pterm.PrefixPrinter{pterm.Info, pterm.Success, pterm.Warning, pterm.Error, *pterm.Fatal.WithFatal(false)}
	for _, pp := range prefixPrinter {
		genericPrinters = append(genericPrinters, pp)
	}

	for _, p := range genericPrinters {
		for _, str := range randomStrings {
			t.Run("TestGenericPrinter_Sprint", func(t *testing.T) {
				out := p.Sprint(str)
				assert.NotEmpty(t, out, p)
				assert.NotEmpty(t, pterm.RemoveColors(out), p)
			})
			t.Run("TestGenericPrinter_Sprintln", func(t *testing.T) {
				out := p.Sprintln(str)
				assert.NotEmpty(t, out, p)
				assert.NotEmpty(t, pterm.RemoveColors(out), p)
			})
			t.Run("TestGenericPrinter_Sprintf", func(t *testing.T) {
				out := p.Sprintf(str+"%s World", "Hello")
				assert.NotEmpty(t, out, p)
				assert.NotEmpty(t, pterm.RemoveColors(out), p)
			})

			t.Run("TestGenericPrinter_Print", func(t *testing.T) {
				out := captureStdout(func(w io.Writer) {
					p.Print(str)
				})
				assert.NotEmpty(t, out, p)
				assert.NotEmpty(t, pterm.RemoveColors(out), p)
			})
			t.Run("TestGenericPrinter_Println", func(t *testing.T) {
				out := captureStdout(func(w io.Writer) {
					p.Println(str)
				})
				assert.NotEmpty(t, out, p)
				assert.NotEmpty(t, pterm.RemoveColors(out), p)
			})
			t.Run("TestGenericPrinter_Printf", func(t *testing.T) {
				out := captureStdout(func(w io.Writer) {
					p.Printf(str+"%s World", "Hello")
				})
				assert.NotEmpty(t, out, p)
				assert.NotEmpty(t, pterm.RemoveColors(out), p)
			})
		}
	}
}
