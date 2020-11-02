package pterm

// Bars is used to display multiple Bar.
type Bars []Bar

// Bar is used in bar charts.
type Bar struct {
	Label string
	Value int
	Style *Style
}

// WithLabel returns a new Bar with a specific option.
func (p Bar) WithLabel(s string) *Bar {
	p.Label = s
	return &p
}

// WithValue returns a new Bar with a specific option.
func (p Bar) WithValue(value int) *Bar {
	p.Value = value
	return &p
}

// WithStyle returns a new Bar with a specific option.
func (p Bar) WithStyle(style *Style) *Bar {
	p.Style = style
	return &p
}
