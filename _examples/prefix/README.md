### prefix/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	pterm.Info.Println("Fetching the latest release")
	pterm.Success.Println("Deployed v1.4.2 to production")
	pterm.Warning.Println("The build cache is older than 30 days")
	pterm.Error.Println("Could not reach registry.example.com")

	// Debug output only shows up after EnableDebugMessages, for example
	// behind a --verbose flag.
	pterm.EnableDebugMessages()
	pterm.Debug.Println("Resolved 42 packages in 120ms")
}
```

</details>

### prefix/custom

![Animation](custom/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Status printers are values. With* returns a changed copy, so the
	// package defaults stay untouched.
	deploy := pterm.Success.WithGlyph("↑").WithTone(pterm.ToneAccent)
	deploy.Println("Pushed 3 images to the registry")

	// Store a customized printer to reuse it across the program.
	hint := pterm.Info.WithTone(pterm.ToneMuted)
	hint.Println("Run with --verbose to see every request")
}
```

</details>

### prefix/multiline

![Animation](multiline/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// A message with more than one line turns into a callout. The first
	// line becomes the title.
	pterm.Error.Println("Unknown command disk\nRun `app --help` to see available commands.")

	pterm.Println()

	// WithTitle sets the title explicitly and keeps the whole message as
	// the body.
	pterm.Warning.WithTitle("Heads up").Println("Nothing here can be undone.")
}
```

</details>

