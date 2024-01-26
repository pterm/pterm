package pterm

import "strings"

type OptionsRenderer interface {
	RenderOptions(options []string) string
}

type InterpolatedOptionsRenderer struct {
	separator string
}

// RenderOptions implements OptionsRenderer.
func (r InterpolatedOptionsRenderer) RenderOptions(options []string) string {
	content := ThemeDefault.SecondaryStyle.Sprint("you have selected: ")
	content += ThemeDefault.SecondaryStyle.Add(*Italic.ToStyle()).
		Sprintln(strings.Join(options, r.separator))
	return content
}

func NewInterpolatedOptions(separator string) OptionsRenderer {
	return &InterpolatedOptionsRenderer{
		separator: separator,
	}
}
