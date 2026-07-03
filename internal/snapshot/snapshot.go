// Package snapshot implements pterm's snapshot testing.
//
// A snapshot test captures the rendered output of a printer and compares it
// against a snapshot file committed to the repository. If a change to pterm
// alters the output of a printer, the affected snapshot tests fail and the
// diff shows exactly what changed.
//
// Snapshot files live in testdata/snapshots/ (relative to the package under
// test) and are named after the test that created them. They are committed to
// git on purpose: reviewing a pull request that changes rendering behavior
// shows the output change right next to the code change.
//
// Escape sequences are stored in a human-readable form (ESC is written as the
// literal text `\x1b`) so snapshot files are diffable and readable in any
// editor.
//
// Workflow:
//   - Running a snapshot test for the first time creates the snapshot file
//     and passes. Commit the file.
//   - If output changes, the test fails and prints a line diff.
//   - Run tests with UPDATE_SNAPSHOTS=1 to rewrite the snapshots after an
//     intended rendering change; then review and commit the new files.
//   - On CI (CI env variable set), missing snapshots fail the test instead of
//     being created, so forgetting to commit a snapshot cannot go unnoticed.
package snapshot

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// Dir is the directory (relative to the package under test) in which
// snapshot files are stored.
const Dir = "testdata/snapshots"

var unsafeChars = regexp.MustCompile(`[^a-zA-Z0-9._\-]+`)

// Assert compares got against the snapshot stored for the current test and
// fails the test on any difference. See the package documentation for the
// snapshot workflow.
func Assert(t *testing.T, got string) {
	t.Helper()

	path := pathFor(t)
	encoded := encode(got)

	stored, err := os.ReadFile(path) // #nosec G304 -- the path is derived from the test name, not user input
	if errors.Is(err, os.ErrNotExist) {
		if os.Getenv("CI") != "" {
			t.Fatalf("snapshot %s does not exist; run the test locally and commit the created snapshot", path)
		}

		write(t, path, encoded)
		t.Logf("created snapshot %s", path)

		return
	}

	if err != nil {
		t.Fatalf("reading snapshot %s: %v", path, err)
	}

	if string(stored) == encoded {
		return
	}

	if updateEnabled() {
		write(t, path, encoded)
		t.Logf("updated snapshot %s", path)

		return
	}

	t.Errorf("output does not match snapshot %s\n(run tests with UPDATE_SNAPSHOTS=1 if this change is intended)\n\n%s", path, diff(string(stored), encoded))
}

func updateEnabled() bool {
	v := os.Getenv("UPDATE_SNAPSHOTS")
	return v != "" && v != "0" && !strings.EqualFold(v, "false")
}

// pathFor derives the snapshot file path from the test name.
// Subtest separators become double underscores so hierarchies stay visible.
func pathFor(t *testing.T) string {
	t.Helper()
	name := strings.ReplaceAll(t.Name(), "/", "__")
	name = unsafeChars.ReplaceAllString(name, "_")

	return filepath.Join(Dir, name+".snap")
}

func write(t *testing.T, path, content string) {
	t.Helper()

	if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
		t.Fatalf("creating snapshot directory: %v", err)
	}

	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("writing snapshot %s: %v", path, err)
	}
}

// encode makes control sequences visible so snapshot files are readable and
// diffable. ESC bytes are written as the literal text `\x1b`; everything
// else (including newlines) is kept as-is.
func encode(s string) string {
	s = strings.ReplaceAll(s, "\x1b", `\x1b`)
	s = strings.ReplaceAll(s, "\r", `\r`)

	return s
}

// diff returns a simple line-based diff between the stored snapshot and the
// new output, capped to the first 20 differing lines.
func diff(want, got string) string {
	wantLines := strings.Split(want, "\n")
	gotLines := strings.Split(got, "\n")

	var b strings.Builder
	differences := 0

	for i := 0; i < len(wantLines) || i < len(gotLines); i++ {
		w, g := "<no line>", "<no line>"
		if i < len(wantLines) {
			w = wantLines[i]
		}

		if i < len(gotLines) {
			g = gotLines[i]
		}

		if w == g {
			continue
		}

		differences++
		if differences > 20 {
			b.WriteString("... more differences omitted ...\n")
			break
		}

		fmt.Fprintf(&b, "line %d:\n  snapshot: %s\n  actual:   %s\n", i+1, w, g)
	}

	return b.String()
}
