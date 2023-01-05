package pterm

// Bars is used to display multiple Bar.
type Bars []Bar

// Bar is used in bar charts.
type Bar struct {
	Label      string
	Value      int
	Style      *Style
	LabelStyle *Style
}

// Checkmark is used in the interactive multiselect printer.
type Checkmark struct {
	Checked   string
	Unchecked string
}

// WithLabel returns a new Bar with a specific option.
func (p Bar) WithLabel(s string) *Bar {
	p.Label = s
	return &p
}

// WithLabelStyle returns a new Bar with a specific option.
func (p Bar) WithLabelStyle(style *Style) *Bar {
	p.LabelStyle = style
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
