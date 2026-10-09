### interactive_confirm/demo

![Animation](https://vhs.charm.sh/vhs-5fNTnAUGfbaRlaCGkWMBk3.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// The confirm prompt accepts "y" and "n" as shortcuts. Pressing enter
	// answers with the default value, which is "no" unless changed with
	// WithDefaultValue.
	result, _ := pterm.DefaultInteractiveConfirm.Show()

	pterm.Println()
	pterm.Info.Printfln("You answered: %s", boolToText(result))
}

// boolToText renders the answer as a colored "Yes" or "No".
func boolToText(b bool) string {
	if b {
		return pterm.Green("Yes")
	}
	return pterm.Red("No")
}
```

</details>

### interactive_confirm/custom-answers

![Animation](https://vhs.charm.sh/vhs-3UyCDeB5UPa3TwxwhnAMjL.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// The answer texts can be customized. The keyboard shortcuts follow
	// along: the prompt accepts the lowercase first letter of each answer,
	// here "a" for Apply and "c" for Cancel.
	result, _ := pterm.DefaultInteractiveConfirm.
		WithDefaultText("Do you want to apply the update?").
		WithConfirmText("Apply").
		WithRejectText("Cancel").
		Show()

	pterm.Println()
	pterm.Info.Printfln("You answered: %s", boolToText(result))
}

// boolToText renders the answer as a colored "Apply" or "Cancel".
func boolToText(b bool) string {
	if b {
		return pterm.Green("Apply")
	}
	return pterm.Red("Cancel")
}
```

</details>

