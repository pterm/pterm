### confirm/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// "y" and "n" answer at once. The arrow keys move between the answers
	// and enter takes the highlighted one, which starts on the default.
	ok, _ := pterm.Confirm.
		WithDefaultValue(true).
		Show("Delete 23 caches?")

	if !ok {
		pterm.Info.Println("Nothing was deleted")
		return
	}

	pterm.Success.Println("Deleted 23 caches")
}
```

</details>

### confirm/custom-answers

![Animation](custom-answers/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

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

</details>

