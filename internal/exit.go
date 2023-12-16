package internal

import "os"

// ExitFuncType is the type of function used to exit the program.
type ExitFuncType func(int)

// DefaultExitFunc is the default function used to exit the program.
var DefaultExitFunc ExitFuncType = os.Exit

// Exit calls the current exit function.
func Exit(code int) {
	DefaultExitFunc(code)
}
