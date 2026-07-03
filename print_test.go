package pterm_test

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal"
)

// pterm.Sprint functions

func TestSprint(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		assert.Equal(t, randomString, pterm.Sprint(randomString))
	}
}

func TestSprintf(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		assert.Equal(t, randomString, pterm.Sprintf("%s", randomString))
	}

	assert.Equal(t, "Hello, World!", pterm.Sprintf("Hello, %s!", "World"))
}

func TestSprintfln(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		assert.Equal(t, randomString+"\n", pterm.Sprintfln("%s", randomString))
	}

	assert.Equal(t, "Hello, World!\n", pterm.Sprintfln("Hello, %s!", "World"))
}

func TestSprintln(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		assert.Equal(t, randomString+"\n", pterm.Sprintln(randomString))
	}
}

func TestSprinto(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		assert.Equal(t, "\r"+randomString, pterm.Sprinto(randomString))
	}
}

// Print functions

func TestPrint(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		pterm.EnableOutput()

		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(_ io.Writer) {
				pterm.Print(randomString)
			})
			assert.Equal(t, randomString, out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		pterm.DisableOutput()

		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(_ io.Writer) {
				pterm.Print(randomString)
			})
			assert.Equal(t, "", out)
		}

		pterm.EnableOutput()
	})
}

func TestPrintln(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		pterm.EnableOutput()

		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(_ io.Writer) {
				pterm.Println(randomString)
			})
			assert.Equal(t, randomString+"\n", out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		pterm.DisableOutput()

		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(_ io.Writer) {
				pterm.Println(randomString)
			})
			assert.Equal(t, "", out)
		}

		pterm.EnableOutput()
	})
}

func TestPrintf(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		pterm.EnableOutput()

		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(_ io.Writer) {
				pterm.Printf("%s", randomString)
			})
			assert.Equal(t, randomString, out)
		}

		out := captureStdout(func(_ io.Writer) {
			pterm.Printf("Hello, %s!", "World")
		})
		assert.Equal(t, "Hello, World!", out)
	})

	t.Run("disabled output", func(t *testing.T) {
		pterm.DisableOutput()

		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(_ io.Writer) {
				pterm.Printf("%s", randomString)
			})
			assert.Equal(t, "", out)
		}

		out := captureStdout(func(_ io.Writer) {
			pterm.Printf("Hello, %s!", "World")
		})
		assert.Equal(t, "", out)

		pterm.EnableOutput()
	})
}

func TestPrintfln(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		pterm.EnableOutput()

		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(_ io.Writer) {
				pterm.Printfln("%s", randomString)
			})
			assert.Equal(t, randomString+"\n", out)
		}

		out := captureStdout(func(_ io.Writer) {
			pterm.Printfln("Hello, %s!", "World")
		})
		assert.Equal(t, "Hello, World!\n", out)
	})

	t.Run("disabled output", func(t *testing.T) {
		pterm.DisableOutput()

		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(_ io.Writer) {
				pterm.Printfln("%s", randomString)
			})
			assert.Equal(t, "", out)
		}

		out := captureStdout(func(_ io.Writer) {
			pterm.Printfln("Hello, %s!", "World")
		})
		assert.Equal(t, "", out)

		pterm.EnableOutput()
	})
}

func TestFprint(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		pterm.EnableOutput()

		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				// set default to null to confirm that its correctly using the provided writer
				pterm.SetDefaultOutput(nil)
				pterm.Fprint(w, randomString)
			})
			assert.Equal(t, randomString, out)
		}
	})

	t.Run("confirm defaults to default output when no writer provided", func(t *testing.T) {
		pterm.EnableOutput()

		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(_ io.Writer) {
				pterm.Fprint(nil, randomString)
			})
			assert.Equal(t, randomString, out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		pterm.DisableOutput()

		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Fprint(w, randomString)
			})
			assert.Equal(t, "", out)
		}

		pterm.EnableOutput()
	})
}

func TestFprintln(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		pterm.EnableOutput()

		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				// set default to null to confirm that its correctly using the provided writer
				pterm.SetDefaultOutput(nil)
				pterm.Fprintln(w, randomString)
			})
			assert.Equal(t, randomString+"\n", out)
		}
	})

	t.Run("confirm defaults to default output when no writer provided", func(t *testing.T) {
		pterm.EnableOutput()

		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(_ io.Writer) {
				pterm.Fprintln(nil, randomString)
			})
			assert.Equal(t, randomString+"\n", out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		pterm.DisableOutput()

		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Fprintln(w, randomString)
			})
			assert.Equal(t, "", out)
		}

		pterm.EnableOutput()
	})
}

func TestPrinto(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		pterm.EnableOutput()

		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(_ io.Writer) {
				pterm.Printo(randomString)
			})
			assert.Equal(t, "\r"+randomString, out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		pterm.DisableOutput()

		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(_ io.Writer) {
				pterm.Printo(randomString)
			})
			assert.Equal(t, "", out)
		}

		pterm.DisableOutput()
	})
}

func TestFprinto(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		pterm.EnableOutput()

		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				// set default to null to confirm that its correctly using the provided writer
				pterm.SetDefaultOutput(nil)
				pterm.Fprinto(w, randomString)
			})
			assert.Equal(t, "\r"+randomString, out)
		}
	})

	t.Run("confirm defaults to default output when no writer provided", func(t *testing.T) {
		pterm.EnableOutput()

		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(_ io.Writer) {
				pterm.Fprinto(nil, randomString)
			})
			assert.Equal(t, "\r"+randomString, out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		pterm.DisableOutput()

		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Fprinto(w, randomString)
			})
			assert.Equal(t, "", out)
		}

		pterm.EnableOutput()
	})
}

func TestSetDefaultOutput(_ *testing.T) {
	pterm.SetDefaultOutput(os.Stdout)
}

func TestPrintOnError(t *testing.T) {
	t.Run("PrintOnError", func(t *testing.T) {
		result := captureStdout(func(_ io.Writer) {
			pterm.PrintOnError(errors.New("hello world"))
		})
		assert.Contains(t, result, "hello world")
	})
}

func TestPrintIfError_WithoutError(t *testing.T) {
	t.Run("PrintIfError_WithoutError", func(t *testing.T) {
		result := captureStdout(func(_ io.Writer) {
			pterm.PrintOnError(nil)
		})
		assert.Zero(t, result)
	})
}

func TestPrintOnErrorf(t *testing.T) {
	t.Run("PrintOnErrorf", func(t *testing.T) {
		result := captureStdout(func(_ io.Writer) {
			pterm.PrintOnErrorf("wrapping error : %w", errors.New("hello world"))
		})
		assert.Contains(t, result, "hello world")
	})
}

func TestPrintIfErrorf_WithoutError(t *testing.T) {
	t.Run("PrintIfErrorf_WithoutError", func(t *testing.T) {
		result := captureStdout(func(_ io.Writer) {
			pterm.PrintOnErrorf("", nil)
		})
		assert.Zero(t, result)
	})
}
