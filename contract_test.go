package pterm_test

// This file contains pterm's printer contract tests.
//
// Every printer must follow the conventions documented in printers.go. Instead
// of repeating near-identical tests in every *_printer_test.go file, the
// contracts are verified here once, systematically, for every printer:
//
//   - Builder contract: every With* method uses a value receiver, returns a
//     pointer to a modified copy, never mutates the printer it was called on,
//     and actually sets the field it is named after.
//   - TextPrinter contract: the Print* methods write exactly what the
//     corresponding Sprint* methods return, PrintOnError only prints non-nil
//     errors, custom writers are honored, and styling can be stripped or
//     disabled without changing the visible text layout.
//   - RenderPrinter contract: Render() writes exactly what Srender() returns,
//     and the same styling guarantees hold.
//
// The per-printer test files then only need to test the printer's actual
// rendering behavior (layout, math, alignment), not the plumbing.

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
)

// stripANSI removes all ANSI escape codes (colors, styles, hyperlinks) from s.
func stripANSI(s string) string {
	return pterm.RemoveColorFromString(s)
}

// restoreGlobalStyling makes sure a test that toggles the global styling/color
// state leaves it in the default (fully enabled) state for later tests.
func restoreGlobalStyling(t *testing.T) {
	t.Helper()
	t.Cleanup(func() {
		pterm.EnableStyling()
		pterm.EnableColor()
	})
}

// ---------------------------------------------------------------------------
// Builder contract
// ---------------------------------------------------------------------------

// builderPrinters lists every exported printer (and printer-like type) that
// follows the With* builder convention. Add new printers here.
var builderPrinters = []any{
	// Text printers
	pterm.BasicTextPrinter{},
	pterm.BoxPrinter{},
	pterm.CenterPrinter{},
	pterm.HeaderPrinter{},
	pterm.ParagraphPrinter{},
	pterm.PrefixPrinter{},
	pterm.SectionPrinter{},
	// Render printers
	pterm.BarChartPrinter{},
	pterm.BigTextPrinter{},
	pterm.BulletListPrinter{},
	pterm.HeatmapPrinter{},
	pterm.PanelPrinter{},
	pterm.TablePrinter{},
	pterm.TreePrinter{},
	// Live printers
	pterm.AreaPrinter{},
	pterm.MultiPrinter{},
	pterm.ProgressbarPrinter{},
	pterm.SpinnerPrinter{},
	// Interactive printers
	pterm.InteractiveConfirmPrinter{},
	pterm.InteractiveContinuePrinter{},
	pterm.InteractiveMultiselectPrinter{},
	pterm.InteractiveSelectPrinter{},
	pterm.InteractiveTextInputPrinter{},
	// Other builder-style types
	pterm.Logger{},
}

// builderSeeds provides a non-zero starting instance for printers whose With*
// methods depend on existing state (e.g. options that must already be set).
var builderSeeds = map[string]any{
	"InteractiveContinuePrinter": pterm.InteractiveContinuePrinter{Options: []string{"yes", "no"}},
}

// builderArgs provides explicit arguments for With* methods whose parameters
// must be consistent with the printer's state (index bounds, matching slice
// lengths) and therefore cannot be synthesized blindly.
var builderArgs = map[string][]any{
	"InteractiveContinuePrinter.WithDefaultValueIndex": {1},
	"InteractiveContinuePrinter.WithDefaultValue":      {"no"},
	"InteractiveContinuePrinter.WithHandles":           {[]string{"y", "n"}},
}

// TestBuilderContract verifies the With* builder convention for every printer:
//
//  1. Every With* method must be declared on a value receiver (so it can never
//     mutate the printer it is called on) and return *T.
//  2. Calling a With* method must return a copy in which something actually
//     changed; if the struct has a field named after the method (WithFoo ->
//     Foo), that field must now hold the given value.
//  3. The original printer must be left untouched (deep-equal to before).
func TestBuilderContract(t *testing.T) {
	for _, printer := range builderPrinters {
		printerType := reflect.TypeOf(printer)
		t.Run(printerType.Name(), func(t *testing.T) {
			ptrType := reflect.PointerTo(printerType)

			valueMethods := map[string]bool{}
			for method := range printerType.Methods() {
				valueMethods[method.Name] = true
			}

			for method := range ptrType.Methods() {
				if !strings.HasPrefix(method.Name, "With") {
					continue
				}

				t.Run(method.Name, func(t *testing.T) {
					// 1. Value receiver: pointer-receiver With* methods would
					// mutate the original printer and break chaining safety.
					require.True(t, valueMethods[method.Name],
						"%s.%s must use a value receiver (see printers.go)", printerType.Name(), method.Name)

					m, _ := printerType.MethodByName(method.Name)
					requireBuilderSignature(t, printerType, m)

					original := reflect.New(printerType).Elem()
					if seed, ok := builderSeeds[printerType.Name()]; ok {
						original.Set(reflect.ValueOf(seed))
					}

					before := original.Interface()

					var args, synthesized []reflect.Value

					if explicit, ok := builderArgs[printerType.Name()+"."+method.Name]; ok {
						for _, arg := range explicit {
							args = append(args, reflect.ValueOf(arg))
							synthesized = append(synthesized, reflect.ValueOf(arg))
						}
					} else {
						args, synthesized = synthesizeArgs(t, m.Func.Type())
					}

					result := original.Method(m.Index).Call(args)

					// 3. The original value must not have been changed.
					assert.Equal(t, before, original.Interface(),
						"%s.%s mutated its receiver", printerType.Name(), method.Name)

					got := result[0]
					require.False(t, got.IsNil(), "%s.%s returned nil", printerType.Name(), method.Name)

					modified := got.Elem()

					// 2. The named field (if any) must hold the given value.
					fieldName := strings.TrimPrefix(method.Name, "With")
					if !assertFieldSet(t, printerType, modified, fieldName, m.Func.Type(), synthesized) {
						// No matching field: at minimum, the call must have
						// changed *something* on the copy.
						assert.NotEqual(t, before, modified.Interface(),
							"%s.%s changed nothing on the returned copy", printerType.Name(), method.Name)
					}
				})
			}
		})
	}
}

// requireBuilderSignature checks that a With* method returns exactly one
// value: a pointer to its own printer type.
func requireBuilderSignature(t *testing.T, printerType reflect.Type, m reflect.Method) {
	t.Helper()

	funcType := m.Func.Type()
	require.Equal(t, 1, funcType.NumOut(),
		"%s.%s must return exactly one value", printerType.Name(), m.Name)
	require.Equal(t, reflect.PointerTo(printerType), funcType.Out(0),
		"%s.%s must return *%s", printerType.Name(), m.Name, printerType.Name())
}

// synthesizeArgs builds a non-zero argument list for the given method type
// (index 0 is the receiver, which is skipped). For variadic bool methods
// (the `WithFoo(b ...bool)` pattern) no argument is passed, because the
// convention is that the bare call sets the flag to true.
//
// It returns the call arguments and the plain values that were synthesized
// (excluding the receiver), so callers can verify fields against them.
func synthesizeArgs(t *testing.T, funcType reflect.Type) (args []reflect.Value, synthesized []reflect.Value) {
	t.Helper()

	for i := 1; i < funcType.NumIn(); i++ {
		paramType := funcType.In(i)
		if funcType.IsVariadic() && i == funcType.NumIn()-1 {
			if paramType.Elem().Kind() == reflect.Bool {
				return args, synthesized // bare call => flag defaults to true
			}

			value := synthesizeValue(t, paramType.Elem(), 0)
			args = append(args, value)
			synthesized = append(synthesized, value)

			return args, synthesized
		}

		value := synthesizeValue(t, paramType, 0)
		args = append(args, value)
		synthesized = append(synthesized, value)
	}

	return args, synthesized
}

// synthesizeValue creates a non-zero value of the given type, so that setter
// effects are distinguishable from the zero value.
func synthesizeValue(t *testing.T, typ reflect.Type, depth int) reflect.Value {
	t.Helper()

	if depth > 4 {
		return reflect.Zero(typ)
	}

	// Types that must be built through their constructor to be usable, or
	// whose meaningful state lives in unexported fields.
	switch typ {
	case reflect.TypeFor[*csv.Reader]():
		return reflect.ValueOf(csv.NewReader(strings.NewReader("a,b\nc,d\n")))
	case reflect.TypeFor[time.Time]():
		return reflect.ValueOf(time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC))
	}

	switch typ.Kind() {
	case reflect.Bool:
		return reflect.ValueOf(true).Convert(typ)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.ValueOf(int64(42)).Convert(typ)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return reflect.ValueOf(uint64(42)).Convert(typ)
	case reflect.Float32, reflect.Float64:
		return reflect.ValueOf(4.2).Convert(typ)
	case reflect.String:
		return reflect.ValueOf("pterm-contract-test").Convert(typ)
	case reflect.Slice:
		slice := reflect.MakeSlice(typ, 1, 1)
		slice.Index(0).Set(synthesizeValue(t, typ.Elem(), depth+1))

		return slice

	case reflect.Map:
		m := reflect.MakeMap(typ)
		m.SetMapIndex(synthesizeValue(t, typ.Key(), depth+1), synthesizeValue(t, typ.Elem(), depth+1))

		return m

	case reflect.Pointer:
		ptr := reflect.New(typ.Elem())
		ptr.Elem().Set(synthesizeValue(t, typ.Elem(), depth+1))

		return ptr

	case reflect.Struct:
		value := reflect.New(typ).Elem()
		// Make the struct non-zero by setting its first settable field.
		for i := range typ.NumField() {
			if field := value.Field(i); field.CanSet() {
				field.Set(synthesizeValue(t, typ.Field(i).Type, depth+1))
				break
			}
		}

		return value

	case reflect.Func:
		return reflect.MakeFunc(typ, func(_ []reflect.Value) []reflect.Value {
			results := make([]reflect.Value, typ.NumOut())
			for i := range results {
				results[i] = reflect.Zero(typ.Out(i))
			}

			return results
		})

	case reflect.Interface:
		if reflect.TypeFor[*os.File]().Implements(typ) {
			return reflect.ValueOf(os.Stderr).Convert(typ)
		}

		if typ == reflect.TypeFor[any]() {
			return reflect.ValueOf("pterm-contract-test")
		}

		t.Fatalf("cannot synthesize value for interface type %s; extend synthesizeValue", typ)

		return reflect.Value{}

	default:
		t.Fatalf("cannot synthesize value for type %s; extend synthesizeValue", typ)

		return reflect.Value{}
	}
}

// assertFieldSet verifies that the field named after a With* method holds the
// value that was passed to it. It reports whether the check was applicable
// (i.e. a matching field exists).
func assertFieldSet(t *testing.T, printerType reflect.Type, modified reflect.Value, fieldName string, funcType reflect.Type, synthesized []reflect.Value) bool {
	t.Helper()

	structField, ok := printerType.FieldByName(fieldName)
	if !ok {
		return false
	}

	field := modified.FieldByName(fieldName)

	// Bare variadic-bool call: the flag must have been set to true.
	if len(synthesized) == 0 {
		if structField.Type.Kind() != reflect.Bool {
			return false
		}

		assert.True(t, field.Bool(), "%s.With%s() must set %s to true", printerType.Name(), fieldName, fieldName)

		return true
	}

	// Variadic non-bool (e.g. WithColors(colors ...Color)): the slice field
	// must hold exactly the given values.
	if funcType.IsVariadic() && structField.Type.Kind() == reflect.Slice && len(synthesized) == 1 &&
		synthesized[0].Type() == structField.Type.Elem() {
		expected := reflect.MakeSlice(structField.Type, 1, 1)
		expected.Index(0).Set(synthesized[0])
		assert.Equal(t, expected.Interface(), field.Interface(),
			"%s.With%s must set the %s field", printerType.Name(), fieldName, fieldName)

		return true
	}

	// Single argument setter: the field must hold the given value.
	if len(synthesized) == 1 && synthesized[0].Type().ConvertibleTo(structField.Type) {
		expected := synthesized[0].Convert(structField.Type)
		if expected.Type().Comparable() || expected.Kind() == reflect.Slice || expected.Kind() == reflect.Map || expected.Kind() == reflect.Struct {
			assert.Equal(t, expected.Interface(), field.Interface(),
				"%s.With%s must set the %s field", printerType.Name(), fieldName, fieldName)

			return true
		}
	}

	return false
}

// ---------------------------------------------------------------------------
// TextPrinter contract
// ---------------------------------------------------------------------------

// textPrinterCase describes one TextPrinter implementation under contract
// test. withWriter is nil for printers that have no Writer field (colors).
type textPrinterCase struct {
	name       string
	printer    pterm.TextPrinter
	withWriter func(io.Writer) pterm.TextPrinter
}

func textPrinterCases() []textPrinterCase {
	return []textPrinterCase{
		{"BasicTextPrinter", &pterm.DefaultBasicText, func(w io.Writer) pterm.TextPrinter { return pterm.DefaultBasicText.WithWriter(w) }},
		{"PrefixPrinter", &pterm.Info, func(w io.Writer) pterm.TextPrinter { return pterm.Info.WithWriter(w) }},
		{"HeaderPrinter", &pterm.DefaultHeader, func(w io.Writer) pterm.TextPrinter { return pterm.DefaultHeader.WithWriter(w) }},
		{"SectionPrinter", &pterm.DefaultSection, func(w io.Writer) pterm.TextPrinter { return pterm.DefaultSection.WithWriter(w) }},
		{"BoxPrinter", &pterm.DefaultBox, func(w io.Writer) pterm.TextPrinter { return pterm.DefaultBox.WithWriter(w) }},
		{"CenterPrinter", &pterm.DefaultCenter, func(w io.Writer) pterm.TextPrinter { return pterm.DefaultCenter.WithWriter(w) }},
		{"ParagraphPrinter", &pterm.DefaultParagraph, func(w io.Writer) pterm.TextPrinter { return pterm.DefaultParagraph.WithWriter(w) }},
		{"Color", pterm.FgCyan, nil},
		{"RGB", pterm.NewRGB(12, 34, 56), nil},
		{"RGBStyle", pterm.NewRGBStyle(pterm.RGB{R: 200, G: 100, B: 50}), nil},
	}
}

// textPrinterSamples are inputs every text printer must handle. They cover
// plain text, multiline text, unicode, format verbs as literals, and empty
// input.
var textPrinterSamples = []struct {
	name string
	args []any
}{
	{"simple", []any{"Hello, PTerm!"}},
	{"multiline", []any{"first line\nsecond line"}},
	{"unicode", []any{"ünïcödé ✓ 汉字"}},
	{"percent literal", []any{"100%s complete"}},
	{"multiple operands", []any{"a", 1, true}},
	{"empty", []any{""}},
}

// TestTextPrinterContract verifies, for every TextPrinter, that the Print*
// methods write exactly what the Sprint* counterparts return, that error
// helpers behave, that custom writers are honored, and that color/styling can
// be disabled without breaking the output.
func TestTextPrinterContract(t *testing.T) {
	for _, tc := range textPrinterCases() {
		t.Run(tc.name, func(t *testing.T) {
			t.Run("Print writes exactly Sprint", func(t *testing.T) {
				for _, sample := range textPrinterSamples {
					t.Run(sample.name, func(t *testing.T) {
						expected := tc.printer.Sprint(sample.args...)
						got := captureStdout(func(_ io.Writer) { tc.printer.Print(sample.args...) })
						assert.Equal(t, expected, got)
					})
				}
			})

			t.Run("Println writes exactly Sprintln", func(t *testing.T) {
				for _, sample := range textPrinterSamples {
					t.Run(sample.name, func(t *testing.T) {
						expected := tc.printer.Sprintln(sample.args...)
						got := captureStdout(func(_ io.Writer) { tc.printer.Println(sample.args...) })
						assert.Equal(t, expected, got)
					})
				}
			})

			t.Run("Printf writes exactly Sprintf", func(t *testing.T) {
				expected := tc.printer.Sprintf("value: %d/%s", 42, "many")
				got := captureStdout(func(_ io.Writer) { tc.printer.Printf("value: %d/%s", 42, "many") })
				assert.Equal(t, expected, got)
			})

			t.Run("Printfln writes exactly Sprintfln", func(t *testing.T) {
				expected := tc.printer.Sprintfln("value: %d/%s", 42, "many")
				got := captureStdout(func(_ io.Writer) { tc.printer.Printfln("value: %d/%s", 42, "many") })
				assert.Equal(t, expected, got)
			})

			t.Run("Sprintf formats like fmt", func(t *testing.T) {
				assert.Equal(t,
					tc.printer.Sprint(fmt.Sprintf("value: %d/%s", 42, "many")),
					tc.printer.Sprintf("value: %d/%s", 42, "many"),
				)
			})

			t.Run("PrintOnError", func(t *testing.T) {
				t.Run("prints non-nil errors", func(t *testing.T) {
					got := captureStdout(func(_ io.Writer) { tc.printer.PrintOnError(errors.New("something broke")) })
					assert.Contains(t, stripANSI(got), "something broke")
				})
				t.Run("prints nothing for nil errors", func(t *testing.T) {
					got := captureStdout(func(_ io.Writer) { tc.printer.PrintOnError(nil, "not an error", 42) })
					assert.Empty(t, got)
				})
			})

			t.Run("PrintOnErrorf", func(t *testing.T) {
				t.Run("wraps non-nil errors", func(t *testing.T) {
					got := captureStdout(func(_ io.Writer) { tc.printer.PrintOnErrorf("wrapped: %w", errors.New("inner cause")) })
					assert.Contains(t, stripANSI(got), "wrapped: inner cause")
				})
				t.Run("prints nothing for nil errors", func(t *testing.T) {
					got := captureStdout(func(_ io.Writer) { tc.printer.PrintOnErrorf("wrapped: %w", nil) })
					assert.Empty(t, got)
				})
			})

			if tc.withWriter != nil {
				t.Run("honors custom writer", func(t *testing.T) {
					var buf strings.Builder
					p := tc.withWriter(&buf)
					stdout := captureStdout(func(_ io.Writer) { p.Print("to custom writer") })
					assert.Empty(t, stdout, "output must not leak to the default writer")
					assert.Contains(t, stripANSI(buf.String()), "to custom writer")
				})
			}

			t.Run("styling invariants", func(t *testing.T) {
				restoreGlobalStyling(t)

				styled := tc.printer.Sprint("styled invariant check")
				assert.Contains(t, stripANSI(styled), "styled invariant check")

				// Disabling color must yield exactly the styled output minus
				// the escape codes: color must never change the layout.
				pterm.DisableColor()

				uncolored := tc.printer.Sprint("styled invariant check")
				assert.Equal(t, stripANSI(styled), uncolored,
					"disabling color must not change the visible output")
				pterm.EnableColor()

				// Raw output mode must not emit any escape codes at all.
				pterm.DisableStyling()

				raw := tc.printer.Sprint("styled invariant check")
				assert.NotContains(t, raw, "\x1b", "raw output must not contain escape codes")
				assert.Contains(t, raw, "styled invariant check")
				pterm.EnableStyling()
			})
		})
	}
}

// ---------------------------------------------------------------------------
// RenderPrinter contract
// ---------------------------------------------------------------------------

// renderPrinterCase describes one configured RenderPrinter under contract
// test. The instances are configured with small but representative data,
// because most render printers render nothing (or error) without data.
type renderPrinterCase struct {
	name    string
	printer pterm.RenderPrinter
	// skipColorInvariant marks printers whose layout legitimately depends on
	// whether colors are available (e.g. the heatmap replaces colored cells
	// with values when colors are disabled).
	skipColorInvariant bool
}

func renderPrinterCases() []renderPrinterCase {
	return []renderPrinterCase{
		{name: "TablePrinter", printer: pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{
			{"Name", "Number"},
			{"Alice", "1"},
			{"Bob", "22"},
		})},
		{name: "TreePrinter", printer: pterm.DefaultTree.WithRoot(pterm.TreeNode{
			Text: "root",
			Children: []pterm.TreeNode{
				{Text: "child a"},
				{Text: "child b", Children: []pterm.TreeNode{{Text: "grandchild"}}},
			},
		})},
		{name: "BulletListPrinter", printer: pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
			{Level: 0, Text: "level zero"},
			{Level: 1, Text: "level one"},
		})},
		{name: "PanelPrinter", printer: pterm.DefaultPanel.WithPanels(pterm.Panels{
			{{Data: "left"}, {Data: "right"}},
		})},
		{name: "BarChartPrinter", printer: pterm.DefaultBarChart.WithShowValue().WithBars(pterm.Bars{
			{Label: "a", Value: 1},
			{Label: "b", Value: 3},
		})},
		{name: "BigTextPrinter", printer: pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("Hi"))},
		{name: "HeatmapPrinter", printer: pterm.DefaultHeatmap.WithAxisData(pterm.HeatmapAxis{
			XAxis: []string{"x1", "x2"},
			YAxis: []string{"y1"},
		}).WithData([][]float32{{1, 2}}), skipColorInvariant: true},
	}
}

// TestRenderPrinterContract verifies, for every RenderPrinter, that Render()
// writes exactly what Srender() returns and that styling can be disabled
// without breaking the output.
func TestRenderPrinterContract(t *testing.T) {
	for _, tc := range renderPrinterCases() {
		t.Run(tc.name, func(t *testing.T) {
			t.Run("Srender returns content without error", func(t *testing.T) {
				out, err := tc.printer.Srender()
				require.NoError(t, err)
				assert.NotEmpty(t, strings.TrimSpace(stripANSI(out)))
			})

			t.Run("Render writes exactly Srender", func(t *testing.T) {
				expected, err := tc.printer.Srender()
				require.NoError(t, err)

				got := captureStdout(func(_ io.Writer) {
					assert.NoError(t, tc.printer.Render())
				})
				// Render appends a trailing newline for printers whose
				// Srender output does not end in one; accept both.
				assert.Contains(t, []string{expected, expected + "\n"}, got)
			})

			if !tc.skipColorInvariant {
				t.Run("disabling color keeps the layout", func(t *testing.T) {
					restoreGlobalStyling(t)

					styled, err := tc.printer.Srender()
					require.NoError(t, err)

					pterm.DisableColor()

					uncolored, err := tc.printer.Srender()

					pterm.EnableColor()
					require.NoError(t, err)

					assert.Equal(t, stripANSI(styled), uncolored,
						"disabling color must not change the visible output")
				})
			}

			t.Run("raw output contains no escape codes", func(t *testing.T) {
				restoreGlobalStyling(t)
				pterm.DisableStyling()

				out, err := tc.printer.Srender()

				pterm.EnableStyling()
				require.NoError(t, err)

				assert.NotContains(t, out, "\x1b", "raw output must not contain escape codes")
			})
		})
	}
}
