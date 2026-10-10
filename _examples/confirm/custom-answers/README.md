# confirm/custom-answers

![Animation](animation.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// The answers can say what they do. Their first letters become the
	// shortcuts, here "a" for Apply and "c" for Cancel.
	apply, _ := pterm.Confirm.
		WithConfirmText("Apply").
		WithRejectText("Cancel").
		Show("Apply 4 changes to production?")

	if apply {
		pterm.Success.Println("Applied 4 changes")
	}
}
```
