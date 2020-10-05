package test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

// Sprint functions

func TestSprint(t *testing.T) {
	for _, randomString := range randomStrings {
		assert.Equal(t, randomString, pterm.Sprint(randomString))
	}
}

func TestSprintf(t *testing.T) {
	for _, randomString := range randomStrings {
		assert.Equal(t, randomString, pterm.Sprintf(randomString))
	}
	assert.Equal(t, "Hello, World!", pterm.Sprintf("Hello, %s!", "World"))
}

func TestSprintln(t *testing.T) {
	for _, randomString := range randomStrings {
		assert.Equal(t, randomString+"\n", pterm.Sprintln(randomString))
	}
}

func TestSprinto(t *testing.T) {
	for _, randomString := range randomStrings {
		assert.Equal(t, "\r"+randomString, pterm.Sprinto(randomString))
	}
}

// Print functions

func TestPrint(t *testing.T) {
	for _, randomString := range randomStrings {
		out := captureStdout(func(w io.Writer) {
			pterm.Print(randomString)
		})
		assert.Equal(t, randomString, out)
	}
}

func TestPrintln(t *testing.T) {
	for _, randomString := range randomStrings {
		out := captureStdout(func(w io.Writer) {
			pterm.Println(randomString)
		})
		assert.Equal(t, randomString+"\n", out)
	}
}

func TestPrintf(t *testing.T) {
	for _, randomString := range randomStrings {
		out := captureStdout(func(w io.Writer) {
			pterm.Printf(randomString)
		})
		assert.Equal(t, randomString, out)
	}
	out := captureStdout(func(w io.Writer) {
		pterm.Printf("Hello, %s!", "World")
	})
	assert.Equal(t, "Hello, World!", out)
}

func TestFprint(t *testing.T) {
	for _, randomString := range randomStrings {
		out := captureStdout(func(w io.Writer) {
			pterm.Fprint(w, randomString)
		})
		assert.Equal(t, randomString, out)
	}
}

func TestFprintln(t *testing.T) {
	for _, randomString := range randomStrings {
		out := captureStdout(func(w io.Writer) {
			pterm.Fprintln(w, randomString)
		})
		assert.Equal(t, randomString+"\n", out)
	}
}

func TestPrinto(t *testing.T) {
	for _, randomString := range randomStrings {
		out := captureStdout(func(w io.Writer) {
			pterm.Printo(randomString)
		})
		assert.Equal(t, "\r"+randomString, out)
	}
}

func TestFprinto(t *testing.T) {
	for _, randomString := range randomStrings {
		out := captureStdout(func(w io.Writer) {
			pterm.Fprinto(w, randomString)
		})
		assert.Equal(t, "\r"+randomString, out)
	}
}

func TestRemoveColors(t *testing.T) {
	for _, randomString := range randomStrings {
		testString := pterm.Cyan(randomString)
		assert.Equal(t, randomString, pterm.RemoveColors(testString))
	}
}
