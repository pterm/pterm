### center/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Center moves text to the middle of the terminal. By default the text
	// moves as one block, so the lines of a box stay aligned.
	box := pterm.Box.WithTitle("pterm").Sprint("Beautiful terminal output\nwith one line of code")
	pterm.Center.Println(box)
}
```

</details>

### center/each-line

![Animation](each-line/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// For loose text, every line can be centered on its own.
	pterm.Center.WithCenterEachLineSeparately().Println("Thanks for trying pterm\nStar it on GitHub if you like it\nhttps://github.com/pterm/pterm")
}
```

</details>

