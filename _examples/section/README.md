### section/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Sections split output into parts. Each one starts with a blank line,
	// so the program does not have to add spacing itself.
	pterm.Section.Println("Installation")
	pterm.Text.Println("Run the installer and follow the prompts.")

	pterm.Section.Println("Configuration")
	pterm.Text.Println("Settings live in ~/.config/app/config.toml.")
}
```

</details>

### section/levels

![Animation](levels/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Deeper levels get quieter: level 2 is bold text, level 3 is muted.
	pterm.Section.Println("Packages")
	pterm.Section.WithLevel(2).Println("Runtime")
	pterm.Text.Println("atomicgo.dev/color, atomicgo.dev/cursor")
	pterm.Section.WithLevel(3).Println("Indirect")
	pterm.Text.Println("golang.org/x/term")

	// A prefix makes the depth visible, like headings in Markdown.
	pterm.Section.WithPrefix("#").WithLevel(2).WithBottomPadding(1).Println("Development")
	pterm.Text.Println("github.com/pterm/pterm/internal/vt")
}
```

</details>

