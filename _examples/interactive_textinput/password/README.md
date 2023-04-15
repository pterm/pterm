# interactive_textinput/password

![Animation](animation.svg)

```go
package main

import "github.com/pterm/pterm"

func main() {
	result, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Enter your password")

	logger := pterm.DefaultLogger
	logger.Info("Password received", logger.Args("password", result))
}

```
