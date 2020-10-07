package pterm

import (
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm/internal"
)

// Sprint functions

func TestSprint(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		assert.Equal(t, randomString, Sprint(randomString))
	}
}

func TestSprintf(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		assert.Equal(t, randomString, Sprintf(randomString))
	}
	assert.Equal(t, "Hello, World!", Sprintf("Hello, %s!", "World"))
}

func TestSprintln(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		assert.Equal(t, randomString+"\n", Sprintln(randomString))
	}
}

func TestSprinto(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		assert.Equal(t, "\r"+randomString, Sprinto(randomString))
	}
}

// Print functions

func TestPrint(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		out := captureStdout(func(w io.Writer) {
			Print(randomString)
		})
		assert.Equal(t, randomString, out)
	}
}

func TestPrintln(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		out := captureStdout(func(w io.Writer) {
			Println(randomString)
		})
		assert.Equal(t, randomString+"\n", out)
	}
}

func TestPrintf(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		out := captureStdout(func(w io.Writer) {
			Printf(randomString)
		})
		assert.Equal(t, randomString, out)
	}
	out := captureStdout(func(w io.Writer) {
		Printf("Hello, %s!", "World")
	})
	assert.Equal(t, "Hello, World!", out)
}

func TestFprint(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		out := captureStdout(func(w io.Writer) {
			Fprint(w, randomString)
		})
		assert.Equal(t, randomString, out)
	}
}

func TestFprintln(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		out := captureStdout(func(w io.Writer) {
			Fprintln(w, randomString)
		})
		assert.Equal(t, randomString+"\n", out)
	}
}

func TestPrinto(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		out := captureStdout(func(w io.Writer) {
			Printo(randomString)
		})
		assert.Equal(t, "\r"+randomString, out)
	}
}

func TestFprinto(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		out := captureStdout(func(w io.Writer) {
			Fprinto(w, randomString)
		})
		assert.Equal(t, "\r"+randomString, out)
	}
}

func TestRemoveColors(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		testString := Cyan(randomString)
		assert.Equal(t, randomString, RemoveColorFromString(testString))
	}
}

func TestGenericPrinter(t *testing.T) {
	var genericPrinters = []GenericPrinter{DefaultSection, DefaultHeader}

	prefixPrinter := []PrefixPrinter{Info, Success, Warning, Error, *Fatal.WithFatal(false)}
	for _, pp := range prefixPrinter {
		genericPrinters = append(genericPrinters, pp)
	}

	for _, p := range genericPrinters {
		for _, str := range internal.RandomStrings {
			t.Run("TestGenericPrinter_Sprint", func(t *testing.T) {
				out := p.Sprint(str)
				assert.NotEmpty(t, out, p)
				assert.NotEmpty(t, RemoveColorFromString(out), p)
			})
			t.Run("TestGenericPrinter_Sprintln", func(t *testing.T) {
				out := p.Sprintln(str)
				assert.NotEmpty(t, out, p)
				assert.NotEmpty(t, RemoveColorFromString(out), p)
			})
			t.Run("TestGenericPrinter_Sprintf", func(t *testing.T) {
				out := p.Sprintf(str+"%s World", "Hello")
				assert.NotEmpty(t, out, p)
				assert.NotEmpty(t, RemoveColorFromString(out), p)
			})

			t.Run("TestGenericPrinter_Print", func(t *testing.T) {
				out := captureStdout(func(w io.Writer) {
					p.Print(str)
				})
				assert.NotEmpty(t, out, p)
				assert.NotEmpty(t, RemoveColorFromString(out), p)
			})
			t.Run("TestGenericPrinter_Println", func(t *testing.T) {
				out := captureStdout(func(w io.Writer) {
					p.Println(str)
				})
				assert.NotEmpty(t, out, p)
				assert.NotEmpty(t, RemoveColorFromString(out), p)
			})
			t.Run("TestGenericPrinter_Printf", func(t *testing.T) {
				out := captureStdout(func(w io.Writer) {
					p.Printf(str+"%s World", "Hello")
				})
				assert.NotEmpty(t, out, p)
				assert.NotEmpty(t, RemoveColorFromString(out), p)
			})

			t.Run("TestGenericPrinterPrintPrintsSprint", func(t *testing.T) {
				out := captureStdout(func(w io.Writer) {
					p.Print(str)
				})
				assert.Equal(t, p.Sprint(str), out)
			})
			t.Run("TestGenericPrinterPrintlnPrintsSprintln", func(t *testing.T) {
				out := captureStdout(func(w io.Writer) {
					p.Println(str)
				})
				assert.Equal(t, p.Sprintln(str), out)
			})
		}
	}
}

func TestStyle_Add(t *testing.T) {
	assert.Equal(t, Style{FgRed, BgGreen}, Style{FgRed}.Add(Style{BgGreen}))
	assert.Equal(t, Style{FgRed, BgGreen, Bold}, Style{FgRed}.Add(Style{BgGreen}).Add(Style{Bold}))
	assert.Equal(t, Style{FgRed, BgGreen, Bold}, Style{FgRed}.Add(Style{BgGreen, Bold}))
	assert.Equal(t, Style{FgRed, BgGreen, Bold}, Style{FgRed}.Add(Style{BgGreen}, Style{Bold}))
}

// CaptureStdout captures everything written to the terminal and returns it as a string.
func captureStdout(f func(w io.Writer)) string {
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	SetDefaultOutput(w)

	f(w)

	_ = w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = originalStdout
	SetDefaultOutput(originalStdout)

	return string(out)
}
