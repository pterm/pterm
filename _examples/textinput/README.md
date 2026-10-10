### textinput/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// A single-line input. The placeholder shows what is expected and
	// disappears with the first key press; enter submits.
	name, _ := pterm.TextInput.
		WithPlaceholder("lowercase, no spaces").
		Show("Project name")

	pterm.Success.Printfln("Created %s", name)
}
```

</details>

### textinput/default-value

![Animation](default-value/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// The default value is shown muted until the user types. Enter on an
	// empty input returns it, and tab copies it into the input to edit it.
	branch, _ := pterm.TextInput.
		WithDefaultValue("main").
		Show("Branch to deploy")

	pterm.Info.Printfln("Deploying %s", branch)
}
```

</details>

### textinput/multi-line

![Animation](multi-line/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// In multi-line mode enter starts a new line and ctrl+d submits. Long
	// lines wrap at word boundaries, and pasted text keeps its line breaks.
	message, _ := pterm.TextInput.
		WithMultiLine().
		Show("Commit message")

	pterm.Info.Printfln("Committing %d characters", len(message))
}
```

</details>

### textinput/password

![Animation](password/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

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

</details>

