package pterm_test

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"

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
		assert.Equal(t, randomString, pterm.Sprintf(randomString))
	}
	assert.Equal(t, "Hello, World!", pterm.Sprintf("Hello, %s!", "World"))
}

func TestSprintfln(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		assert.Equal(t, randomString+"\n", pterm.Sprintfln(randomString))
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
		pterm.Output = true
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Print(randomString)
			})
			assert.Equal(t, randomString, out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		pterm.Output = false
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Print(randomString)
			})
			assert.Equal(t, "", out)
		}
		pterm.Output = true
	})
}

func TestPrintln(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		pterm.Output = true
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Println(randomString)
			})
			assert.Equal(t, randomString+"\n", out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		pterm.Output = false
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Println(randomString)
			})
			assert.Equal(t, "", out)
		}
		pterm.Output = true
	})
}

func TestPrintf(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		pterm.Output = true
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Printf(randomString)
			})
			assert.Equal(t, randomString, out)
		}
		out := captureStdout(func(w io.Writer) {
			pterm.Printf("Hello, %s!", "World")
		})
		assert.Equal(t, "Hello, World!", out)
	})

	t.Run("disabled output", func(t *testing.T) {
		pterm.Output = false
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Printf(randomString)
			})
			assert.Equal(t, "", out)
		}
		out := captureStdout(func(w io.Writer) {
			pterm.Printf("Hello, %s!", "World")
		})
		assert.Equal(t, "", out)
		pterm.Output = true
	})
}

func TestPrintfln(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		pterm.Output = true
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Printfln(randomString)
			})
			assert.Equal(t, randomString+"\n", out)
		}
		out := captureStdout(func(w io.Writer) {
			pterm.Printfln("Hello, %s!", "World")
		})
		assert.Equal(t, "Hello, World!\n", out)
	})

	t.Run("disabled output", func(t *testing.T) {
		pterm.Output = false
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Printfln(randomString)
			})
			assert.Equal(t, "", out)
		}
		out := captureStdout(func(w io.Writer) {
			pterm.Printfln("Hello, %s!", "World")
		})
		assert.Equal(t, "", out)
		pterm.Output = true
	})
}

func TestFprint(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		pterm.Output = true
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Fprint(w, randomString)
			})
			assert.Equal(t, randomString, out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		pterm.Output = false
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Fprint(w, randomString)
			})
			assert.Equal(t, "", out)
		}
		pterm.Output = true
	})
}

func TestFprintln(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		pterm.Output = true
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Fprintln(w, randomString)
			})
			assert.Equal(t, randomString+"\n", out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		pterm.Output = false
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Fprintln(w, randomString)
			})
			assert.Equal(t, "", out)
		}
		pterm.Output = true
	})
}

func TestPrinto(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		pterm.Output = true
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Printo(randomString)
			})
			assert.Equal(t, "\r"+randomString, out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		pterm.Output = false
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Printo(randomString)
			})
			assert.Equal(t, "", out)
		}
		pterm.Output = false
	})
}

func TestFprinto(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		pterm.Output = true
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Fprinto(w, randomString)
			})
			assert.Equal(t, "\r"+randomString, out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		pterm.Output = false
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Fprinto(w, randomString)
			})
			assert.Equal(t, "", out)
		}
		pterm.Output = true
	})
}

func TestSetDefaultOutput(t *testing.T) {
	pterm.SetDefaultOutput(os.Stdout)
}

func TestPrintOnError(t *testing.T) {
	t.Run("PrintOnError", func(t *testing.T) {
		result := captureStdout(func(w io.Writer) {
			pterm.PrintOnError(errors.New("hello world"))
		})
		assert.Contains(t, result, "hello world")
	})
}

func TestPrintIfError_WithoutError(t *testing.T) {
	t.Run("PrintIfError_WithoutError", func(t *testing.T) {
		result := captureStdout(func(w io.Writer) {
			pterm.PrintOnError(nil)
		})
		assert.Empty(t, result)
	})
}
