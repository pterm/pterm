package pterm_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
)

// TestThemeWithMethods verifies every With* method on Theme via reflection:
// it must take a single Style, return a Theme copy in which exactly the field
// named after the method holds the given style, and leave the original
// untouched (value receiver semantics).
func TestThemeWithMethods(t *testing.T) {
	style := pterm.Style{pterm.FgRed, pterm.BgBlue, pterm.Bold}
	themeType := reflect.TypeFor[pterm.Theme]()

	methodsChecked := 0

	for method := range themeType.Methods() {
		if !strings.HasPrefix(method.Name, "With") {
			continue
		}

		methodsChecked++

		t.Run(method.Name, func(t *testing.T) {
			funcType := method.Func.Type()
			require.Equal(t, 2, funcType.NumIn(), "%s must take exactly one argument", method.Name)
			require.Equal(t, reflect.TypeFor[pterm.Style](), funcType.In(1), "%s must take a Style", method.Name)
			require.Equal(t, 1, funcType.NumOut())
			require.Equal(t, themeType, funcType.Out(0), "%s must return a Theme copy", method.Name)

			original := pterm.Theme{}
			result := reflect.ValueOf(original).Method(method.Index).Call([]reflect.Value{reflect.ValueOf(style)})[0]

			assert.Equal(t, pterm.Theme{}, original, "%s must not mutate its receiver", method.Name)

			fieldName := strings.TrimPrefix(method.Name, "With")
			field := result.FieldByName(fieldName)
			require.True(t, field.IsValid(), "Theme has no field %q matching %s", fieldName, method.Name)
			assert.Equal(t, style, field.Interface(), "%s must set the %s field", method.Name, fieldName)

			// Nothing else may have changed on the copy: resetting the named
			// field must yield the zero Theme again.
			resultReset := result.Interface().(pterm.Theme)
			resetField := reflect.ValueOf(&resultReset).Elem().FieldByName(fieldName)
			resetField.Set(reflect.Zero(resetField.Type()))
			assert.Equal(t, pterm.Theme{}, resultReset, "%s must only change the %s field", method.Name, fieldName)
		})
	}

	assert.NotZero(t, methodsChecked, "no With* methods found on Theme")
}

// TestThemeDefaultStylesFlowIntoPrinters verifies that the default printers
// really render with the styles from ThemeDefault: changing a theme style
// changes the escape codes the printer emits.
func TestThemeDefaultStylesFlowIntoPrinters(t *testing.T) {
	restoreGlobalStyling(t)

	original := pterm.ThemeDefault.SectionStyle
	pterm.ThemeDefault.SectionStyle = pterm.Style{pterm.FgRed}

	t.Cleanup(func() { pterm.ThemeDefault.SectionStyle = original })

	out := pterm.DefaultSection.Sprint("Title")
	assert.Contains(t, out, "\x1b[31m", "DefaultSection must render with the (changed) ThemeDefault.SectionStyle")
	assert.Contains(t, stripANSI(out), "Title")
}

// TestThemeLoggerStylesFlowIntoLogger verifies that the logger's level styles
// come from ThemeDefault: changing a logger theme style changes the escape
// codes the logger emits.
func TestThemeLoggerStylesFlowIntoLogger(t *testing.T) {
	restoreGlobalStyling(t)

	original := pterm.ThemeDefault.LoggerInfoStyle
	pterm.ThemeDefault.LoggerInfoStyle = pterm.Style{pterm.FgMagenta}

	t.Cleanup(func() { pterm.ThemeDefault.LoggerInfoStyle = original })

	var buf strings.Builder

	pterm.DefaultLogger.WithWriter(&buf).WithTime(false).Info("hello")
	assert.Contains(t, buf.String(), "\x1b[35m", "the logger must render INFO with the (changed) ThemeDefault.LoggerInfoStyle")
}

// TestThemeHeatmapColorsFlowIntoDefaultHeatmap verifies that the default
// heatmap colors are defined by the theme.
func TestThemeHeatmapColorsFlowIntoDefaultHeatmap(t *testing.T) {
	assert.Equal(t, pterm.ThemeDefault.HeatmapColors, pterm.DefaultHeatmap.Colors)
	assert.Equal(t, pterm.ThemeDefault.HeatmapTextColor, pterm.DefaultHeatmap.TextColor)
	assert.Equal(t, pterm.ThemeDefault.HeatmapRGBRange, pterm.DefaultHeatmap.RGBRange)
	assert.Equal(t, pterm.ThemeDefault.HeatmapTextRGB, pterm.DefaultHeatmap.TextRGB)
}

// TestPrinterWithCustomStyleEmitsItsCodes verifies that a printer constructed
// with an explicit style renders exactly that style's codes, independent of
// the theme.
func TestPrinterWithCustomStyleEmitsItsCodes(t *testing.T) {
	section := pterm.SectionPrinter{
		Level:           1,
		IndentCharacter: "#",
	}

	out := section.WithStyle(pterm.NewStyle(pterm.FgMagenta, pterm.Bold)).Sprint("Custom")
	assert.Contains(t, out, "\x1b[35;1m")
	assert.Contains(t, stripANSI(out), "# Custom")
}
