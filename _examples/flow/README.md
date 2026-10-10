### flow/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	flow, _ := pterm.Flow.WithSubtitle("23 targets").Start("disk clean")

	step := flow.Step("Measuring what can be freed")
	time.Sleep(2 * time.Second)
	step.Success()

	// Any printer used while the flow runs is drawn on its rail.
	pterm.Callout.Warn("Heads up").Println("Up to 110 GB across 23 targets.\nDocker images are downloaded again on the next build.")

	flow.Section("Package managers")

	for _, cache := range []string{"npm", "pnpm", "go build cache"} {
		step := flow.Step("Cleaning " + cache)
		time.Sleep(700 * time.Millisecond)
		step.Success("Cleaned " + cache)
	}

	flow.Step("Docker images").Skip("Docker is not running")

	flow.Success("Freed 24.7 GB")
}
```

</details>

### flow/live-printers

![Animation](live-printers/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	flow, _ := pterm.Flow.Start("deploy")

	step := flow.Step("Building image")
	time.Sleep(time.Second)
	step.Success()

	// Any printer started while the flow runs joins its rail. A progress bar
	// animates in place and collapses into a status line when it is done.
	bar, _ := pterm.Progressbar.WithTotal(24).Start("Uploading layers")
	for range 24 {
		time.Sleep(60 * time.Millisecond)
		bar.Increment()
	}

	// Rendered output, such as a table, is indented next to the rail.
	_ = pterm.Table.WithData(pterm.TableData{
		{"Pod", "Status", "Restarts"},
		{"api-7f9c", "running", "0"},
		{"worker-2b1d", "running", "1"},
	}).Render()

	flow.Success("Deployed v2.4.1")
}
```

</details>

### flow/sections

![Animation](sections/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	flow, _ := pterm.Flow.Start("release")

	flow.Section("Build")
	build := flow.Step("Compiling for 6 platforms")
	time.Sleep(time.Second)
	build.Success()
	flow.Log("dist/ holds 6 archives, 48 MB")

	flow.Section("Publish")
	upload := flow.Step("Uploading to GitHub")
	time.Sleep(time.Second)
	upload.Fail("GitHub returned 502")
	pterm.Error.Println("The release is a draft until the upload works")

	flow.Fail("Release incomplete")
}
```

</details>

