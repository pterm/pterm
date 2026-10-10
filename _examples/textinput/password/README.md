# textinput/password

![Animation](animation.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// WithMask draws every typed character as the mask. The answered
	// prompt shows a fixed number of masks, so it does not give away the
	// length of the password either.
	password, _ := pterm.TextInput.
		WithMask("*").
		Show("Password")

	pterm.Success.Printfln("Signed in with a %d character password", len(password))
}
```
