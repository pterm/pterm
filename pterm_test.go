package pterm

import (
	"github.com/pterm/pterm/internal"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
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
	t.Run("enabled output", func(t *testing.T) {
		for _, randomString := range internal.RandomStrings {
			out := internal.CaptureStdout(func(w io.Writer) {
				Print(randomString)
			})
			assert.Equal(t, randomString, out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		DisableOutput = true
		for _, randomString := range internal.RandomStrings {
			out := internal.CaptureStdout(func(w io.Writer) {
				Print(randomString)
			})
			assert.Equal(t, "", out)
		}
		DisableOutput = false
	})
}

func TestPrintln(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		for _, randomString := range internal.RandomStrings {
			out := internal.CaptureStdout(func(w io.Writer) {
				Println(randomString)
			})
			assert.Equal(t, randomString+"\n", out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		DisableOutput = true
		for _, randomString := range internal.RandomStrings {
			out := internal.CaptureStdout(func(w io.Writer) {
				Println(randomString)
			})
			assert.Equal(t, "", out)
		}
		DisableOutput = false
	})
}

func TestPrintf(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		for _, randomString := range internal.RandomStrings {
			out := internal.CaptureStdout(func(w io.Writer) {
				Printf(randomString)
			})
			assert.Equal(t, randomString, out)
		}
		out := internal.CaptureStdout(func(w io.Writer) {
			Printf("Hello, %s!", "World")
		})
		assert.Equal(t, "Hello, World!", out)
	})

	t.Run("disabled output", func(t *testing.T) {
		DisableOutput = true
		for _, randomString := range internal.RandomStrings {
			out := internal.CaptureStdout(func(w io.Writer) {
				Printf(randomString)
			})
			assert.Equal(t, "", out)
		}
		out := internal.CaptureStdout(func(w io.Writer) {
			Printf("Hello, %s!", "World")
		})
		assert.Equal(t, "", out)
		DisableOutput = false
	})
}

func TestFprint(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		for _, randomString := range internal.RandomStrings {
			out := internal.CaptureStdout(func(w io.Writer) {
				Fprint(w, randomString)
			})
			assert.Equal(t, randomString, out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		DisableOutput = true
		for _, randomString := range internal.RandomStrings {
			out := internal.CaptureStdout(func(w io.Writer) {
				Fprint(w, randomString)
			})
			assert.Equal(t, "", out)
		}
		DisableOutput = false
	})
}

func TestFprintln(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		for _, randomString := range internal.RandomStrings {
			out := internal.CaptureStdout(func(w io.Writer) {
				Fprintln(w, randomString)
			})
			assert.Equal(t, randomString+"\n", out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		DisableOutput = true
		for _, randomString := range internal.RandomStrings {
			out := internal.CaptureStdout(func(w io.Writer) {
				Fprintln(w, randomString)
			})
			assert.Equal(t, "", out)
		}
		DisableOutput = false
	})
}

func TestPrinto(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		for _, randomString := range internal.RandomStrings {
			out := internal.CaptureStdout(func(w io.Writer) {
				Printo(randomString)
			})
			assert.Equal(t, "\r"+randomString, out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		DisableOutput = true
		for _, randomString := range internal.RandomStrings {
			out := internal.CaptureStdout(func(w io.Writer) {
				Printo(randomString)
			})
			assert.Equal(t, "", out)
		}
		DisableOutput = false
	})
}

func TestFprinto(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		for _, randomString := range internal.RandomStrings {
			out := internal.CaptureStdout(func(w io.Writer) {
				Fprinto(w, randomString)
			})
			assert.Equal(t, "\r"+randomString, out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		DisableOutput = true
		for _, randomString := range internal.RandomStrings {
			out := internal.CaptureStdout(func(w io.Writer) {
				Fprinto(w, randomString)
			})
			assert.Equal(t, "", out)
		}
		DisableOutput = false
	})
}

func TestRemoveColors(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		testString := Cyan(randomString)
		assert.Equal(t, randomString, RemoveColorFromString(testString))
	}
}

func TestStyle_Add(t *testing.T) {
	assert.Equal(t, Style{FgRed, BgGreen}, Style{FgRed}.Add(Style{BgGreen}))
	assert.Equal(t, Style{FgRed, BgGreen, Bold}, Style{FgRed}.Add(Style{BgGreen}).Add(Style{Bold}))
	assert.Equal(t, Style{FgRed, BgGreen, Bold}, Style{FgRed}.Add(Style{BgGreen, Bold}))
	assert.Equal(t, Style{FgRed, BgGreen, Bold}, Style{FgRed}.Add(Style{BgGreen}, Style{Bold}))
}
