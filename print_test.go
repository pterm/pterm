package pterm

import (
	"errors"
	"io"
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

func TestSprintfln(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		assert.Equal(t, randomString+"\n", Sprintfln(randomString))
	}
	assert.Equal(t, "Hello, World!\n", Sprintfln("Hello, %s!", "World"))
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
		Output = true
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				Print(randomString)
			})
			assert.Equal(t, randomString, out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		Output = false
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				Print(randomString)
			})
			assert.Equal(t, "", out)
		}
		Output = true
	})
}

func TestPrintln(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		Output = true
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				Println(randomString)
			})
			assert.Equal(t, randomString+"\n", out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		Output = false
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				Println(randomString)
			})
			assert.Equal(t, "", out)
		}
		Output = true
	})
}

func TestPrintf(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		Output = true
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
	})

	t.Run("disabled output", func(t *testing.T) {
		Output = false
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				Printf(randomString)
			})
			assert.Equal(t, "", out)
		}
		out := captureStdout(func(w io.Writer) {
			Printf("Hello, %s!", "World")
		})
		assert.Equal(t, "", out)
		Output = true
	})
}

func TestPrintfln(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		Output = true
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				Printfln(randomString)
			})
			assert.Equal(t, randomString+"\n", out)
		}
		out := captureStdout(func(w io.Writer) {
			Printfln("Hello, %s!", "World")
		})
		assert.Equal(t, "Hello, World!\n", out)
	})

	t.Run("disabled output", func(t *testing.T) {
		Output = false
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				Printfln(randomString)
			})
			assert.Equal(t, "", out)
		}
		out := captureStdout(func(w io.Writer) {
			Printfln("Hello, %s!", "World")
		})
		assert.Equal(t, "", out)
		Output = true
	})
}

func TestFprint(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		Output = true
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				Fprint(w, randomString)
			})
			assert.Equal(t, randomString, out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		Output = false
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				Fprint(w, randomString)
			})
			assert.Equal(t, "", out)
		}
		Output = true
	})
}

func TestFprintln(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		Output = true
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				Fprintln(w, randomString)
			})
			assert.Equal(t, randomString+"\n", out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		Output = false
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				Fprintln(w, randomString)
			})
			assert.Equal(t, "", out)
		}
		Output = true
	})
}

func TestPrinto(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		Output = true
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				Printo(randomString)
			})
			assert.Equal(t, "\r"+randomString, out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		Output = false
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				Printo(randomString)
			})
			assert.Equal(t, "", out)
		}
		Output = false
	})
}

func TestFprinto(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		Output = true
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				Fprinto(w, randomString)
			})
			assert.Equal(t, "\r"+randomString, out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		Output = false
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				Fprinto(w, randomString)
			})
			assert.Equal(t, "", out)
		}
		Output = true
	})
}

func TestSetDefaultOutput(t *testing.T) {
	SetDefaultOutput(os.Stdout)
}

func TestPrintOnError(t *testing.T) {
	t.Run("PrintOnError", func(t *testing.T) {
		result := captureStdout(func(w io.Writer) {
			PrintOnError(errors.New("hello world"))
		})
		assert.Contains(t, result, "hello world")
	})
}

func TestPrintIfError_WithoutError(t *testing.T) {
	t.Run("PrintIfError_WithoutError", func(t *testing.T) {
		result := captureStdout(func(w io.Writer) {
			PrintOnError(nil)
		})
		assert.Empty(t, result)
	})
}
