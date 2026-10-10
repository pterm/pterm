# style/inline

![Animation](animation.gif)

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
