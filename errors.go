package pterm

import "errors"

var (
	// ErrTerminalSizeNotDetectable - the terminal size can not be detected and the fallback values are used.
	ErrTerminalSizeNotDetectable = errors.New("terminal size could not be detected - using fallback value")

	// ErrHexCodeIsInvalid - the given HEX code is invalid.
	ErrHexCodeIsInvalid = errors.New("hex code is not valid")

	// ErrKeyWithoutValue - an odd number of arguments was passed to a pterm Logger's Args method.
	ErrKeyWithoutValue = "ERROR: key_without_value"
)
