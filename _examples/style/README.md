### style/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"atomicgo.dev/color"

	"github.com/pterm/pterm"
)

func main() {
	// A Style combines a foreground tone, an optional background tone and
	// text attributes. Tones follow the active theme.
	title := pterm.NewStyle(pterm.TonePrimary, color.Bold)
	label := pterm.NewStyle(pterm.ToneText).On(pterm.ToneBorder)
	link := pterm.NewStyle(pterm.ToneInfo, color.Underline)

	pterm.Println(title.Sprint("Release notes"))
	pterm.Println(label.Sprint(" v0.13.0 "), pterm.ToneMuted.Sprint("2026-10-10"))
	pterm.Println("Read more at", link.Sprint("https://pterm.sh"))

	// With adds attributes to a copy, so the base style stays untouched.
	pterm.Println(title.With(color.Italic).Sprint("Thanks to all contributors"))
}
```

</details>

### style/inline

![Animation](inline/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Tones color single words inside your own text. They turn into plain
	// text when the output is piped or colors are disabled.
	pterm.Printfln("Uploaded %s to %s in %s",
		pterm.ToneAccent.Sprint("report.pdf"),
		pterm.TonePrimary.Bold().Sprint("s3://backups"),
		pterm.ToneMuted.Sprint("1.4s"),
	)

	pterm.Printfln("%s passed, %s failed, %s skipped",
		pterm.ToneSuccess.Sprint(41),
		pterm.ToneError.Sprint(2),
		pterm.ToneMuted.Sprint(3),
	)
}
```

</details>

