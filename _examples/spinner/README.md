### spinner/demo

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
	spinner, _ := pterm.Spinner.Start("Resolving dependencies")
	time.Sleep(2 * time.Second)
	spinner.UpdateText("Downloading 42 packages")
	time.Sleep(2 * time.Second)

	// Output printed while the spinner runs appears above it.
	pterm.Info.Println("Using the module cache at ~/go/pkg/mod")
	time.Sleep(time.Second)

	spinner.Success("Installed 42 packages")
}
```

</details>

### spinner/context

![Animation](context/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"context"
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// The spinner stops as canceled when the context ends, which makes
	// timeouts a one-liner.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	spinner, _ := pterm.Spinner.WithContext(ctx).Start("Waiting for the database")

	<-ctx.Done()
	time.Sleep(100 * time.Millisecond)

	// Stopping a spinner that already ended does nothing.
	_ = spinner.Stop()
}
```

</details>

### spinner/run-with-spinner

![Animation](run-with-spinner/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"context"
	"errors"
	"time"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	ctx := context.Background()

	// The spinner ends as a success when the function returns nil.
	_ = putils.RunWithSpinner(ctx, "Fetching releases", func(ctx context.Context, s *pterm.SpinnerPrinter) error {
		time.Sleep(time.Second)
		s.UpdateText("Fetched 12 releases")

		return nil
	})

	// An error ends it as a failure that names the task and the error.
	err := putils.RunWithSpinner(ctx, "Uploading report", func(ctx context.Context, _ *pterm.SpinnerPrinter) error {
		time.Sleep(time.Second)
		return errors.New("the bucket does not exist")
	})
	if err != nil {
		pterm.Info.Println("Create the bucket and run the upload again.")
	}
}
```

</details>

### spinner/statuses

![Animation](statuses/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// A spinner ends as a status line. Pick the ending that matches what
	// happened.
	s, _ := pterm.Spinner.Start("Compiling")
	time.Sleep(time.Second)
	s.Success("Compiled 18 packages")

	s, _ = pterm.Spinner.Start("Running tests")
	time.Sleep(time.Second)
	s.Warning("2 tests were skipped")

	s, _ = pterm.Spinner.Start("Uploading coverage")
	time.Sleep(time.Second)
	s.Fail("Coverage upload timed out")

	s, _ = pterm.Spinner.Start("Publishing docs")
	s.Skip("Docs are unchanged")
}
```

</details>

