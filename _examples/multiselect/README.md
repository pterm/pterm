### multiselect/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Groups, descriptions and right-aligned details turn a plain list into
	// a small report. Selected options start checked; disabled ones are
	// shown but cannot be toggled.
	items := []pterm.Option{
		{Group: "Package managers", Label: "npm cache", Description: "~/.npm/_cacache", Detail: "1.2 GB", Selected: true},
		{Group: "Package managers", Label: "yarn cache", Description: "~/Library/Caches/Yarn", Detail: "640 MB", Selected: true},
		{Group: "Package managers", Label: "pnpm store", Description: "~/Library/pnpm/store", Detail: "2.1 GB", Selected: true},
		{Group: "Package managers", Label: "go module cache", Description: "~/go/pkg/mod", Detail: "3.4 GB"},
		{Group: "Package managers", Label: "cargo registry", Description: "~/.cargo/registry", Detail: "890 MB", Selected: true},
		{Group: "Build caches", Label: "go build cache", Description: "~/Library/Caches/go-build", Detail: "5.8 GB", Selected: true},
		{Group: "Build caches", Label: "Xcode DerivedData", Description: "~/Library/Developer/Xcode/DerivedData", Detail: "12.4 GB", Selected: true},
		{Group: "Build caches", Label: "Gradle caches", Description: "~/.gradle/caches", Detail: "1.9 GB", Selected: true},
		{Group: "Build caches", Label: "Bazel output base", Description: "in use by a running build", Detail: "7.2 GB", Disabled: true},
		{Group: "Containers", Label: "Docker images", Description: "dangling and unused images", Detail: "18.3 GB", Selected: true},
		{Group: "Containers", Label: "Docker build cache", Description: "buildx layers", Detail: "6.1 GB", Selected: true},
		{Group: "Containers", Label: "Docker volumes", Description: "may hold databases, check first", Detail: "4.7 GB"},
	}

	cleaned, _ := pterm.Multiselect.
		WithItems(items).
		WithMaxHeight(12).
		Show("What should be cleaned?")

	pterm.Info.Printfln("Cleaning %d caches", len(cleaned))
}
```

</details>

### multiselect/custom-keys

![Animation](custom-keys/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"atomicgo.dev/keyboard/keys"

	"github.com/pterm/pterm"
)

func main() {
	// Space toggles and enter confirms by default. Here enter toggles and
	// tab confirms instead, and the checkboxes are drawn in brackets. The
	// key hints at the bottom follow the keys.
	services, _ := pterm.Multiselect.
		WithOptions([]string{"api", "worker", "scheduler", "web"}).
		WithKeySelect(keys.Enter).
		WithKeyConfirm(keys.Tab).
		WithCheckmark("[x]", "[ ]").
		Show("Which services should restart?")

	pterm.Info.Printfln("Restarting %v", services)
}
```

</details>

### multiselect/options

![Animation](options/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Plain strings work as options too. Default options start checked,
	// and are what Show returns when the program does not run in a
	// terminal.
	toppings, _ := pterm.Multiselect.
		WithOptions([]string{"Cheese", "Mushrooms", "Olives", "Onions", "Peppers", "Pineapple"}).
		WithDefaultOptions([]string{"Cheese"}).
		Show("Pick your toppings")

	pterm.Info.Printfln("Ordering a pizza with %d toppings", len(toppings))
}
```

</details>

### multiselect/show-selected-options

![Animation](show-selected-options/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// By default the answered prompt only counts the selection. With
	// WithShowSelectedOptions it also lists what was picked, which reads
	// well in logs and recordings.
	pterm.Multiselect.
		WithOptions([]string{"linux/amd64", "linux/arm64", "darwin/amd64", "darwin/arm64", "windows/amd64"}).
		WithDefaultOptions([]string{"linux/amd64", "darwin/arm64"}).
		WithShowSelectedOptions().
		Show("Build for which platforms?")
}
```

</details>

