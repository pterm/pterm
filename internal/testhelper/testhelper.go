// Package testhelper provides minimal assertion helpers used by pterm's tests.
// Predicate logic is delegated to atomicgo.dev/assert; this package only adds
// the *testing.T plumbing that fails the test on a negative result.
package testhelper

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"

	"atomicgo.dev/assert"
)

// msg formats an optional trailing message in the testza style: a single
// string is used as-is, otherwise the args are formatted via fmt.Sprintf when
// the first arg is a format string, falling back to fmt.Sprint.
func msg(args []any) string {
	if len(args) == 0 {
		return ""
	}
	if format, ok := args[0].(string); ok {
		if len(args) == 1 {
			return ": " + format
		}
		return ": " + fmt.Sprintf(format, args[1:]...)
	}
	return ": " + fmt.Sprint(args...)
}

// AssertEqual fails the test if expected != actual (deep equality).
func AssertEqual(t *testing.T, expected, actual any, msgAndArgs ...any) {
	t.Helper()
	if !assert.Equal(expected, actual) {
		t.Errorf("expected %#v, got %#v%s", expected, actual, msg(msgAndArgs))
	}
}

// AssertEqualValues fails the test if expected and actual are not equal after
// converting expected to actual's type. This matches testza's behavior, which
// is needed for comparisons like []float32 vs a named slice type.
func AssertEqualValues(t *testing.T, expected, actual any, msgAndArgs ...any) {
	t.Helper()
	if !equalValues(expected, actual) {
		t.Errorf("expected %#v, got %#v%s", expected, actual, msg(msgAndArgs))
	}
}

func equalValues(expected, actual any) bool {
	if assert.Equal(expected, actual) {
		return true
	}
	actualType := reflect.TypeOf(actual)
	if actualType == nil {
		return false
	}
	expectedValue := reflect.ValueOf(expected)
	if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
		return reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
	}
	return false
}

// AssertTrue fails the test if v is false.
func AssertTrue(t *testing.T, v bool, msgAndArgs ...any) {
	t.Helper()
	if !v {
		t.Errorf("expected true, got false%s", msg(msgAndArgs))
	}
}

// AssertFalse fails the test if v is true.
func AssertFalse(t *testing.T, v bool, msgAndArgs ...any) {
	t.Helper()
	if v {
		t.Errorf("expected false, got true%s", msg(msgAndArgs))
	}
}

// AssertNil fails the test if v is not nil.
func AssertNil(t *testing.T, v any, msgAndArgs ...any) {
	t.Helper()
	if !assert.Nil(v) {
		t.Errorf("expected nil, got %#v%s", v, msg(msgAndArgs))
	}
}

// AssertNotNil fails the test if v is nil.
func AssertNotNil(t *testing.T, v any, msgAndArgs ...any) {
	t.Helper()
	if assert.Nil(v) {
		t.Errorf("expected non-nil value%s", msg(msgAndArgs))
	}
}

// AssertZero fails the test if v is not the zero value of its type.
func AssertZero(t *testing.T, v any, msgAndArgs ...any) {
	t.Helper()
	if !assert.Zero(v) {
		t.Errorf("expected zero value, got %#v%s", v, msg(msgAndArgs))
	}
}

// AssertNotZero fails the test if v equals the zero value of its type.
func AssertNotZero(t *testing.T, v any, msgAndArgs ...any) {
	t.Helper()
	if assert.Zero(v) {
		t.Errorf("expected non-zero value, got %#v%s", v, msg(msgAndArgs))
	}
}

// AssertNoError fails the test if err is non-nil.
func AssertNoError(t *testing.T, err error, msgAndArgs ...any) {
	t.Helper()
	if err != nil {
		t.Errorf("unexpected error: %v%s", err, msg(msgAndArgs))
	}
}

// AssertContains fails the test if container does not contain element.
// Strings, slices, arrays, and maps are supported.
func AssertContains(t *testing.T, container, element any, msgAndArgs ...any) {
	t.Helper()
	if !assert.Contains(container, element) {
		t.Errorf("expected %#v to contain %#v%s", container, element, msg(msgAndArgs))
	}
}

// AssertNotContains fails the test if container contains element.
func AssertNotContains(t *testing.T, container, element any, msgAndArgs ...any) {
	t.Helper()
	if assert.Contains(container, element) {
		t.Errorf("expected %#v to not contain %#v%s", container, element, msg(msgAndArgs))
	}
}

// AssertPanics fails the test if f does not panic.
func AssertPanics(t *testing.T, f func(), msgAndArgs ...any) {
	t.Helper()
	if !assert.Panic(f) {
		t.Errorf("expected function to panic%s", msg(msgAndArgs))
	}
}

// AssertNotPanics fails the test if f panics.
func AssertNotPanics(t *testing.T, f func(), msgAndArgs ...any) {
	t.Helper()
	if assert.Panic(f) {
		t.Errorf("expected function not to panic%s", msg(msgAndArgs))
	}
}

// CaptureStderr redirects os.Stderr for the duration of f and returns
// everything written to it. The signature mirrors testza.CaptureStderr so the
// existing call sites do not need to change.
//
// Reads run on a background goroutine so the pipe buffer cannot fill up while
// f is still running; without that, a writer like ProgressbarPrinter's
// schedule.Every goroutine could deadlock on a write once the OS pipe buffer
// is exhausted.
func CaptureStderr(f func(w io.Writer) error) (string, error) {
	original := os.Stderr
	r, w, err := os.Pipe()
	if err != nil {
		return "", err
	}
	os.Stderr = w

	type readResult struct {
		out string
		err error
	}
	done := make(chan readResult, 1)
	go func() {
		var buf bytes.Buffer
		_, copyErr := io.Copy(&buf, r)
		done <- readResult{buf.String(), copyErr}
	}()

	fErr := f(w)

	os.Stderr = original
	_ = w.Close()
	res := <-done
	_ = r.Close()
	if fErr != nil {
		return res.out, fErr
	}
	return res.out, res.err
}
