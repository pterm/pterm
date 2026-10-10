### badge/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"fmt"

	"github.com/pterm/pterm"
)

func main() {
	// Badges label the state or kind of something. Sprint returns the
	// badge as a string, so it fits into any line of output.
	fmt.Println(pterm.Badge.WithTone(pterm.ToneWarning).Sprint("dry run"), "Nothing was deleted")
	fmt.Println(pterm.Badge.WithTone(pterm.ToneSuccess).Sprint("pass"), "42 tests in 1.3s")
	fmt.Println(pterm.Badge.WithTone(pterm.ToneError).Sprint("fail"), "2 tests in 0.4s")
	fmt.Println(pterm.Badge.Sprint("new"), "Faster startup")
}
```

</details>

### badge/chip

![Animation](chip/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"fmt"

	"github.com/pterm/pterm"
)

func main() {
	// Chips are the quiet variant: the text in the badge's tone, between
	// brackets. They work well for tags next to list items.
	chip := pterm.Badge.Chip()

	fmt.Println("api-server ", chip.WithTone(pterm.ToneSuccess).Sprint("running"))
	fmt.Println("worker     ", chip.WithTone(pterm.ToneWarning).Sprint("restarting"))
	fmt.Println("scheduler  ", chip.WithTone(pterm.ToneMuted).Sprint("stopped"))
}
```

</details>

