package pterm_test

import (
	"atomicgo.dev/keyboard/keys"
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/MarvinJWendt/testza"
	"github.com/pterm/pterm"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

var printables = []interface{}{"Hello, World!", 1337, true, false, -1337, 'c', 1.5, "\\", "%s"}
var terminalWidth = 80
var terminalHeight = 60

func TestMain(m *testing.M) {
	pterm.SetForcedTerminalSize(terminalWidth, terminalHeight)
	os.Stderr = os.NewFile(0, os.DevNull)
	setupStdoutCapture()
	exitVal := m.Run()
	teardownStdoutCapture()
	os.Exit(exitVal)
}

// testPrintContains can be used to test Print methods.
func testPrintContains(t *testing.T, logic func(w io.Writer, a interface{})) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			s := captureStdout(func(w io.Writer) {
				logic(w, printable)
			})
			testza.AssertContains(t, s, fmt.Sprint(printable))
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			s := captureStdout(func(w io.Writer) {
				logic(w, printable)
			})
			testza.AssertContains(t, s, fmt.Sprint(printable))
		})
		pterm.EnableStyling()
	}
}

// testPrintfContains can be used to test Printf methods.
func testPrintfContains(t *testing.T, logic func(w io.Writer, format string, a interface{})) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			s := captureStdout(func(w io.Writer) {
				logic(w, "Hello, %v!", printable)
			})
			testza.AssertContains(t, s, fmt.Sprintf("Hello, %v!", fmt.Sprint(printable)))
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			s := captureStdout(func(w io.Writer) {
				logic(w, "Hello, %v!", printable)
			})
			testza.AssertContains(t, s, fmt.Sprintf("Hello, %v!", fmt.Sprint(printable)))
		})
		pterm.EnableStyling()
	}
}

// testPrintflnContains can be used to test Printfln methods.
func testPrintflnContains(t *testing.T, logic func(w io.Writer, format string, a interface{})) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testPrintfContains(t, logic)
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testPrintfContains(t, logic)
		})
		pterm.EnableStyling()
	}
}

// testPrintlnContains can be used to test Println methods.
func testPrintlnContains(t *testing.T, logic func(w io.Writer, a interface{})) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testPrintContains(t, logic)
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testPrintContains(t, logic)
		})
		pterm.EnableStyling()
	}
}

// testSprintContains can be used to test Sprint methods.
func testSprintContains(t *testing.T, logic func(a interface{}) string) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testza.AssertContains(t, logic(printable), fmt.Sprint(printable))
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testza.AssertContains(t, logic(printable), fmt.Sprint(printable))
		})
		pterm.EnableStyling()
	}
}

// testSprintContainsWithoutError can be used to test Sprint methods which return an error.
func testSprintContainsWithoutError(t *testing.T, logic func(a interface{}) (string, error)) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			s, err := logic(printable)
			testza.AssertContains(t, s, fmt.Sprint(printable))
			testza.AssertNoError(t, err)
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			s, err := logic(printable)
			testza.AssertContains(t, s, fmt.Sprint(printable))
			testza.AssertNoError(t, err)
		})
		pterm.EnableStyling()
	}
}

// testSprintfContains can be used to test Sprintf methods.
func testSprintfContains(t *testing.T, logic func(format string, a interface{}) string) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testza.AssertContains(t, logic("Hello, %v!", printable), fmt.Sprintf("Hello, %v!", printable))
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testza.AssertContains(t, logic("Hello, %v!", printable), fmt.Sprintf("Hello, %v!", printable))
		})
		pterm.EnableStyling()
	}
}

// testSprintflnContains can be used to test Sprintfln methods.
func testSprintflnContains(t *testing.T, logic func(format string, a interface{}) string) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testSprintfContains(t, logic)
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testSprintfContains(t, logic)
		})
		pterm.EnableStyling()
	}
}

// testSprintlnContains can be used to test Sprintln methods.
func testSprintlnContains(t *testing.T, logic func(a interface{}) string) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testSprintContains(t, logic)
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testSprintContains(t, logic)
		})
		pterm.EnableStyling()
	}
}

// testDoesOutput can be used to test if something is outputted to stdout.
func testDoesOutput(t *testing.T, logic func(w io.Writer)) {
	testza.AssertNotZero(t, captureStdout(logic))
	pterm.DisableStyling()
	testza.AssertNotZero(t, captureStdout(logic))
	pterm.EnableStyling()
}

// testEmpty checks that a function does not return a string.
func testEmpty(t *testing.T, logic func(a interface{}) string) {
	for _, printable := range printables {
		testza.AssertZero(t, logic(printable))
		pterm.DisableStyling()
		testza.AssertZero(t, logic(printable))
		pterm.EnableStyling()
	}
}

// testDoesNotOutput can be used, to test that something does not output anything to stdout.
func testDoesNotOutput(t *testing.T, logic func(w io.Writer)) {
	testza.AssertZero(t, captureStdout(logic))
	pterm.DisableStyling()
	testza.AssertZero(t, captureStdout(logic))
	pterm.EnableStyling()
}

var outBuf bytes.Buffer

// setupStdoutCapture sets up a fake stdout capture.
func setupStdoutCapture() {
	outBuf.Reset()
	pterm.SetDefaultOutput(&outBuf)
}

// teardownStdoutCapture restores the real stdout.
func teardownStdoutCapture() {
	pterm.SetDefaultOutput(os.Stdout)
}

// captureStdout simulates capturing of os.stdout with a buffer and returns what was writted to the screen
func captureStdout(f func(w io.Writer)) string {
	setupStdoutCapture()
	f(&outBuf)
	return readStdout()
}

// readStdout reads the current stdout buffor. Assumes setupStdoutCapture() has been called before.
func readStdout() string {
	content := outBuf.String()
	outBuf.Reset()
	return content
}

func proxyToDevNull() {
	pterm.SetDefaultOutput(os.NewFile(0, os.DevNull))
}

// testWithMethods calls all methods of a struct starting with "With"
// and checks if the corresponding option has changed.
func testWithMethods(t *testing.T, inputStruct any, blacklist ...string) {
	value := reflect.ValueOf(inputStruct)
	typeOf := value.Type()

	for i := 0; i < value.NumMethod(); i++ {
		method := typeOf.Method(i)

		// Check if method is blacklisted
		if strings.Contains(strings.Join(blacklist, ","), method.Name) {
			continue
		}

		if strings.HasPrefix(method.Name, "With") {
			t.Run(method.Name, func(t *testing.T) {
				// Check if method starts with 'With'
				// Check if method has at least one input parameter
				if method.Type.NumIn() >= 2 { // 2 because the first input is always the receiver
					optionName := method.Name[4:]
					var oldOptionValue any
					if reflect.Indirect(value).FieldByName(optionName).IsValid() {
						oldOptionValue = reflect.Indirect(value).FieldByName(optionName).Interface()
					}

					params := make([]reflect.Value, 0, method.Type.NumIn()-1)
					isVariadic := method.Type.IsVariadic()

					for j := 1; j < method.Type.NumIn(); j++ {
						paramType := method.Type.In(j)
						if isVariadic && j == method.Type.NumIn()-1 {
							// If it's a variadic function, the last parameter is a slice
							paramType = paramType.Elem()
						}

						var param reflect.Value
						switch paramType.Kind() {
						case reflect.String:
							param = reflect.ValueOf("helloworld")
						case reflect.Float32:
							param = reflect.ValueOf(1.0)
						case reflect.Float64:
							param = reflect.ValueOf(1.0)
						case reflect.Int64:
							switch paramType {
							case reflect.TypeOf(time.Duration(0)):
								param = reflect.ValueOf(time.Second)
							default:
								param = reflect.ValueOf(1)
							}
						case reflect.Int:
							switch paramType {
							case reflect.TypeOf(pterm.LogFormatter(0)):
								param = reflect.ValueOf(pterm.LogFormatter(1))
							case reflect.TypeOf(pterm.LogLevel(0)):
								param = reflect.ValueOf(pterm.LogLevel(1))
							case reflect.TypeOf(keys.KeyCode(0)):
								param = reflect.ValueOf(keys.Enter)
							default:
								param = reflect.ValueOf(1)
							}
						case reflect.Bool:
							param = reflect.ValueOf(!oldOptionValue.(bool))
						case reflect.Map:
							switch paramType {
							case reflect.TypeOf(map[string]string{}):
								param = reflect.ValueOf(map[string]string{"a": "b", "c": "d"})
							case reflect.TypeOf(map[string]pterm.Style{}):
								param = reflect.ValueOf(map[string]pterm.Style{"a": pterm.Style{pterm.FgRed}})
							}
						case reflect.Slice:
							switch paramType {
							case reflect.TypeOf(pterm.Style{}):
								param = reflect.ValueOf(pterm.Style{pterm.FgRed})
							case reflect.TypeOf([]string{}):
								param = reflect.ValueOf([]string{"helloworld", "helloworld", "helloworld"})
							case reflect.TypeOf([][]string{}):
								param = reflect.ValueOf([][]string{{"helloworld", "helloworld", "helloworld"}, {"helloworld", "helloworld", "helloworld"}})
							case reflect.TypeOf([]bool{}):
								param = reflect.ValueOf([]bool{true})
							case reflect.TypeOf([]pterm.Bar{}):
								param = reflect.ValueOf(pterm.Bars{
									{"a", 1, pterm.NewStyle(), pterm.NewStyle()},
									{"b", 2, pterm.NewStyle(), pterm.NewStyle()},
									{"c", 3, pterm.NewStyle(), pterm.NewStyle()},
									{"d", 4, pterm.NewStyle(), pterm.NewStyle()},
								})
							case reflect.TypeOf(pterm.Bars{}):
								param = reflect.ValueOf(pterm.Bars{
									{"a", 1, pterm.NewStyle(), pterm.NewStyle()},
									{"b", 2, pterm.NewStyle(), pterm.NewStyle()},
									{"c", 3, pterm.NewStyle(), pterm.NewStyle()},
									{"d", 4, pterm.NewStyle(), pterm.NewStyle()},
								})
							case reflect.TypeOf(pterm.Letters{}):
								param = reflect.ValueOf(pterm.Letters{
									{"a", pterm.NewStyle(), pterm.NewRGB(255, 0, 0)},
									{"b", pterm.NewStyle(), pterm.NewRGB(0, 255, 0)},
									{"c", pterm.NewStyle(), pterm.NewRGB(0, 0, 255)},
								})
							case reflect.TypeOf(pterm.Panels{}):
								param = reflect.ValueOf(pterm.Panels{
									[]pterm.Panel{
										{Data: "a"},
										{Data: "b"},
										{Data: "c"},
									},
								})
							case reflect.TypeOf([]pterm.BulletListItem{}):
								param = reflect.ValueOf([]pterm.BulletListItem{
									{1, "a", pterm.NewStyle(), "b", pterm.NewStyle()},
								})
							}
						case reflect.Pointer:
							switch paramType {
							case reflect.TypeOf(&pterm.Style{}):
								param = reflect.ValueOf(pterm.NewStyle(pterm.FgRed))
							case reflect.TypeOf(&csv.Reader{}):
								param = reflect.ValueOf(csv.NewReader(strings.NewReader("a,b,c\n1,2,3\n4,5,6\nx,y,z")))
							case reflect.TypeOf(&pterm.Checkmark{}):
								param = reflect.ValueOf(&pterm.Checkmark{Checked: "yes", Unchecked: "no"})
							}
						case reflect.Interface:
							t.Skipf("Cannot check With method %s, because it has an interface parameter.", method.Name)
						case reflect.Struct:
							switch paramType {
							case reflect.TypeOf(pterm.TableData{}):
								param = reflect.ValueOf(pterm.TableData{
									{"a", "b", "c"},
									{"1", "2", "3"},
									{"4", "5", "6"},
									{"x", "y", "z"},
								})
							case reflect.TypeOf(pterm.Checkmark{}):
								param = reflect.ValueOf(pterm.Checkmark{Checked: "yes", Unchecked: "no"})
							case reflect.TypeOf(pterm.BoxPrinter{}):
								param = reflect.ValueOf(pterm.DefaultBox)
							case reflect.TypeOf(pterm.Prefix{}):
								param = reflect.ValueOf(pterm.Prefix{
									Text:  "asd",
									Style: pterm.NewStyle(),
								})
							case reflect.TypeOf(pterm.Scope{}):
								param = reflect.ValueOf(pterm.Scope{
									Text:  "asd",
									Style: pterm.NewStyle(),
								})
							case reflect.TypeOf(pterm.TreeNode{}):
								param = reflect.ValueOf(pterm.TreeNode{
									Text:     "asd",
									Children: []pterm.TreeNode{},
								})
							}

						default:
							// Unsupported type, log error and skip this method
							t.Logf("Unsupported parameter type %s for method %s, skipping", paramType, method.Name)
							continue
						}

						params = append(params, param)
					}

					func() {
						defer func() {
							if r := recover(); r != nil {
								t.Errorf("Panic occurred while calling method %s: %v", method.Name, r)
							}
						}()

						results := value.Method(i).Call(params)
						if len(results) > 0 {
							newStruct := results[0]
							inputStruct = newStruct.Interface()

							newOptionValue := reflect.Indirect(newStruct).FieldByName(optionName).Interface()
							if reflect.DeepEqual(oldOptionValue, newOptionValue) {
								t.Errorf("Option %s did not change after calling method %s (old: %s, new: %s)", optionName, method.Name, oldOptionValue, newOptionValue)
							}
						}
					}()
				}
			})
		}

	}
}

func printerTest(t *testing.T, f func()) {
	teardownStdoutCapture()
	f()
	setupStdoutCapture()
	f()
	err := testza.SnapshotCreateOrValidate(t, t.Name(), outBuf.String())
	if err != nil {
		panic(err)
	}
}
