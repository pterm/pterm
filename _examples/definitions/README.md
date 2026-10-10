### definitions/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Groups get an uppercase, muted heading. The term column is shared by
	// all groups, so every description starts in the same column.
	pterm.Definitions.
		WithGroup("Commands",
			pterm.Def("doctor", "Check the installation and report what is missing"),
			pterm.Def("mcp start", "Boot an MCP server"),
			pterm.Def("disk clean", "Free space used by package manager caches, old Docker images and build output that nothing references anymore"),
		).
		WithGroup("Flags",
			pterm.Def("-v, --verbose", "Print what is happening"),
			pterm.Def("-o, --output <file>", "Where to write the result"),
			pterm.Def("-h, --help", "Show this help"),
		).
		Render()
}
```

</details>

### definitions/plain

![Animation](plain/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Without groups there is no heading and no indent, which suits short
	// reference lists such as environment variables.
	pterm.Definitions.WithDefinitions(
		pterm.Def("PTERM_THEME", "macchiato, latte, ansi or auto"),
		pterm.Def("PTERM_ASCII", "use ASCII glyphs"),
		pterm.Def("PTERM_THEME_<SLOT>", "override one color of the theme"),
		pterm.Def("NO_COLOR", "turn colors off"),
	).Render()
}
```

</details>

