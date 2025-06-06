package pterm_test

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/MarvinJWendt/testza"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal"
)

// pterm.Sprint functions

func TestSprint(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		testza.AssertEqual(t, randomString, pterm.Sprint(randomString))
	}
}

func TestSprintf(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		testza.AssertEqual(t, randomString, pterm.Sprintf("%s", randomString))
	}
	testza.AssertEqual(t, "Hello, World!", pterm.Sprintf("Hello, %s!", "World"))
}

func TestSprintfln(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		testza.AssertEqual(t, randomString+"\n", pterm.Sprintfln("%s", randomString))
	}
	testza.AssertEqual(t, "Hello, World!\n", pterm.Sprintfln("Hello, %s!", "World"))
}

func TestSprintln(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		testza.AssertEqual(t, randomString+"\n", pterm.Sprintln(randomString))
	}
}

func TestSprinto(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		testza.AssertEqual(t, "\r"+randomString, pterm.Sprinto(randomString))
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
			testza.AssertEqual(t, randomString, out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		pterm.Output = false
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Print(randomString)
			})
			testza.AssertEqual(t, "", out)
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
			testza.AssertEqual(t, randomString+"\n", out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		pterm.Output = false
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Println(randomString)
			})
			testza.AssertEqual(t, "", out)
		}
		pterm.Output = true
	})
}

func TestPrintf(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		pterm.Output = true
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Printf("%s", randomString)
			})
			testza.AssertEqual(t, randomString, out)
		}
		out := captureStdout(func(w io.Writer) {
			pterm.Printf("Hello, %s!", "World")
		})
		testza.AssertEqual(t, "Hello, World!", out)
	})

	t.Run("disabled output", func(t *testing.T) {
		pterm.Output = false
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Printf("%s", randomString)
			})
			testza.AssertEqual(t, "", out)
		}
		out := captureStdout(func(w io.Writer) {
			pterm.Printf("Hello, %s!", "World")
		})
		testza.AssertEqual(t, "", out)
		pterm.Output = true
	})
}

func TestPrintfln(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		pterm.Output = true
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Printfln("%s", randomString)
			})
			testza.AssertEqual(t, randomString+"\n", out)
		}
		out := captureStdout(func(w io.Writer) {
			pterm.Printfln("Hello, %s!", "World")
		})
		testza.AssertEqual(t, "Hello, World!\n", out)
	})

	t.Run("disabled output", func(t *testing.T) {
		pterm.Output = false
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Printfln("%s", randomString)
			})
			testza.AssertEqual(t, "", out)
		}
		out := captureStdout(func(w io.Writer) {
			pterm.Printfln("Hello, %s!", "World")
		})
		testza.AssertEqual(t, "", out)
		pterm.Output = true
	})
}

func TestFprint(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		pterm.Output = true
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				// set default to null to confirm that its correctly using the provided writer
				pterm.SetDefaultOutput(nil)
				pterm.Fprint(w, randomString)
			})
			testza.AssertEqual(t, randomString, out)
		}
	})

	t.Run("confirm defaults to default output when no writer provided", func(t *testing.T) {
		pterm.Output = true
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Fprint(nil, randomString)
			})
			testza.AssertEqual(t, randomString, out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		pterm.Output = false
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Fprint(w, randomString)
			})
			testza.AssertEqual(t, "", out)
		}
		pterm.Output = true
	})
}

func TestFprintln(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		pterm.Output = true
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				// set default to null to confirm that its correctly using the provided writer
				pterm.SetDefaultOutput(nil)
				pterm.Fprintln(w, randomString)
			})
			testza.AssertEqual(t, randomString+"\n", out)
		}
	})

	t.Run("confirm defaults to default output when no writer provided", func(t *testing.T) {
		pterm.Output = true
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Fprintln(nil, randomString)
			})
			testza.AssertEqual(t, randomString+"\n", out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		pterm.Output = false
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Fprintln(w, randomString)
			})
			testza.AssertEqual(t, "", out)
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
			testza.AssertEqual(t, "\r"+randomString, out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		pterm.Output = false
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Printo(randomString)
			})
			testza.AssertEqual(t, "", out)
		}
		pterm.Output = false
	})
}

func TestFprinto(t *testing.T) {
	t.Run("enabled output", func(t *testing.T) {
		pterm.Output = true
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				// set default to null to confirm that its correctly using the provided writer
				pterm.SetDefaultOutput(nil)
				pterm.Fprinto(w, randomString)
			})
			testza.AssertEqual(t, "\r"+randomString, out)
		}
	})

	t.Run("confirm defaults to default output when no writer provided", func(t *testing.T) {
		pterm.Output = true
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Fprinto(nil, randomString)
			})
			testza.AssertEqual(t, "\r"+randomString, out)
		}
	})

	t.Run("disabled output", func(t *testing.T) {
		pterm.Output = false
		for _, randomString := range internal.RandomStrings {
			out := captureStdout(func(w io.Writer) {
				pterm.Fprinto(w, randomString)
			})
			testza.AssertEqual(t, "", out)
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
		testza.AssertContains(t, result, "hello world")
	})
}

func TestPrintIfError_WithoutError(t *testing.T) {
	t.Run("PrintIfError_WithoutError", func(t *testing.T) {
		result := captureStdout(func(w io.Writer) {
			pterm.PrintOnError(nil)
		})
		testza.AssertZero(t, result)
	})
}

func TestPrintOnErrorf(t *testing.T) {
	t.Run("PrintOnErrorf", func(t *testing.T) {
		result := captureStdout(func(w io.Writer) {
			pterm.PrintOnErrorf("wrapping error : %w", errors.New("hello world"))
		})
		testza.AssertContains(t, result, "hello world")
	})
}

func TestPrintIfErrorf_WithoutError(t *testing.T) {
	t.Run("PrintIfErrorf_WithoutError", func(t *testing.T) {
		result := captureStdout(func(w io.Writer) {
			pterm.PrintOnErrorf("", nil)
		})
		testza.AssertZero(t, result)
	})
}
