package pterm

var (
	// ThemeDefault is the default theme used by PTerm.
	// If this variable is overwritten, the new value is used as default theme.
	ThemeDefault = Theme{
		DefaultText:             Style{FgDefault, BgDefault},
		PrimaryStyle:            Style{FgLightCyan},
		SecondaryStyle:          Style{FgLightMagenta},
		HighlightStyle:          Style{Bold, FgYellow},
		InfoMessageStyle:        Style{FgLightCyan},
		InfoPrefixStyle:         Style{FgBlack, BgCyan},
		SuccessMessageStyle:     Style{FgGreen},
		SuccessPrefixStyle:      Style{FgBlack, BgGreen},
		WarningMessageStyle:     Style{FgYellow},
		WarningPrefixStyle:      Style{FgBlack, BgYellow},
		ErrorMessageStyle:       Style{FgLightRed},
		ErrorPrefixStyle:        Style{FgBlack, BgLightRed},
		FatalMessageStyle:       Style{FgLightRed},
		FatalPrefixStyle:        Style{FgBlack, BgLightRed},
		DescriptionMessageStyle: Style{FgDefault},
		DescriptionPrefixStyle:  Style{FgLightWhite, BgDarkGray},
		ScopeStyle:              Style{FgGray},
		ProgressbarBarStyle:     Style{FgCyan},
		ProgressbarTitleStyle:   Style{FgLightCyan},
		HeaderTextStyle:         Style{FgLightWhite, Bold},
		HeaderBackgroundStyle:   Style{BgGray},
		SpinnerStyle:            Style{FgLightCyan},
		SpinnerTextStyle:        Style{FgLightWhite},
		TableStyle:              Style{FgDefault},
		TableHeaderStyle:        Style{FgLightCyan},
		TableSeparatorStyle:     Style{FgGray},
		HeatmapStyle:            Style{FgDefault},
		HeatmapHeaderStyle:      Style{FgLightCyan},
		HeatmapSeparatorStyle:   Style{FgDefault},
		SectionStyle:            Style{Bold, FgYellow},
		BulletListTextStyle:     Style{FgDefault},
		BulletListBulletStyle:   Style{FgGray},
		TreeStyle:               Style{FgGray},
		TreeTextStyle:           Style{FgDefault},
		LetterStyle:             Style{FgDefault},
		DebugMessageStyle:       Style{FgGray},
		DebugPrefixStyle:        Style{FgBlack, BgGray},
		BoxStyle:                Style{FgDefault},
		BoxTextStyle:            Style{FgDefault},
		BarLabelStyle:           Style{FgLightCyan},
		BarStyle:                Style{FgCyan},
		TimerStyle:              Style{FgGray},
		Checkmark: Checkmark{
			Checked:   Green("✓"),
			Unchecked: Red("✗"),
		},
	}
)

// Theme for PTerm.
// Theme contains every Style used in PTerm. You can create own themes for your application or use one
// of the existing themes.
type Theme struct {
	DefaultText             Style
	PrimaryStyle            Style
	SecondaryStyle          Style
	HighlightStyle          Style
	InfoMessageStyle        Style
	InfoPrefixStyle         Style
	SuccessMessageStyle     Style
	SuccessPrefixStyle      Style
	WarningMessageStyle     Style
	WarningPrefixStyle      Style
	ErrorMessageStyle       Style
	ErrorPrefixStyle        Style
	FatalMessageStyle       Style
	FatalPrefixStyle        Style
	DescriptionMessageStyle Style
	DescriptionPrefixStyle  Style
	ScopeStyle              Style
	ProgressbarBarStyle     Style
	ProgressbarTitleStyle   Style
	HeaderTextStyle         Style
	HeaderBackgroundStyle   Style
	SpinnerStyle            Style
	SpinnerTextStyle        Style
	TimerStyle              Style
	TableStyle              Style
	TableHeaderStyle        Style
	TableSeparatorStyle     Style
	HeatmapStyle            Style
	HeatmapHeaderStyle      Style
	HeatmapSeparatorStyle   Style
	SectionStyle            Style
	BulletListTextStyle     Style
	BulletListBulletStyle   Style
	TreeStyle               Style
	TreeTextStyle           Style
	LetterStyle             Style
	DebugMessageStyle       Style
	DebugPrefixStyle        Style
	BoxStyle                Style
	BoxTextStyle            Style
	BarLabelStyle           Style
	BarStyle                Style
	Checkmark               Checkmark
}

// WithPrimaryStyle returns a new theme with overridden value.
func (t Theme) WithPrimaryStyle(style Style) Theme {
	t.PrimaryStyle = style
	return t
}

// WithSecondaryStyle returns a new theme with overridden value.
func (t Theme) WithSecondaryStyle(style Style) Theme {
	t.SecondaryStyle = style
	return t
}

// WithHighlightStyle returns a new theme with overridden value.
func (t Theme) WithHighlightStyle(style Style) Theme {
	t.HighlightStyle = style
	return t
}

// WithInfoMessageStyle returns a new theme with overridden value.
func (t Theme) WithInfoMessageStyle(style Style) Theme {
	t.InfoMessageStyle = style
	return t
}

// WithInfoPrefixStyle returns a new theme with overridden value.
func (t Theme) WithInfoPrefixStyle(style Style) Theme {
	t.InfoPrefixStyle = style
	return t
}

// WithSuccessMessageStyle returns a new theme with overridden value.
func (t Theme) WithSuccessMessageStyle(style Style) Theme {
	t.SuccessMessageStyle = style
	return t
}

// WithSuccessPrefixStyle returns a new theme with overridden value.
func (t Theme) WithSuccessPrefixStyle(style Style) Theme {
	t.SuccessPrefixStyle = style
	return t
}

// WithWarningMessageStyle returns a new theme with overridden value.
func (t Theme) WithWarningMessageStyle(style Style) Theme {
	t.WarningMessageStyle = style
	return t
}

// WithWarningPrefixStyle returns a new theme with overridden value.
func (t Theme) WithWarningPrefixStyle(style Style) Theme {
	t.WarningPrefixStyle = style
	return t
}

// WithErrorMessageStyle returns a new theme with overridden value.
func (t Theme) WithErrorMessageStyle(style Style) Theme {
	t.ErrorMessageStyle = style
	return t
}

// WithErrorPrefixStyle returns a new theme with overridden value.
func (t Theme) WithErrorPrefixStyle(style Style) Theme {
	t.ErrorPrefixStyle = style
	return t
}

// WithFatalMessageStyle returns a new theme with overridden value.
func (t Theme) WithFatalMessageStyle(style Style) Theme {
	t.FatalMessageStyle = style
	return t
}

// WithFatalPrefixStyle returns a new theme with overridden value.
func (t Theme) WithFatalPrefixStyle(style Style) Theme {
	t.FatalPrefixStyle = style
	return t
}

// WithDescriptionMessageStyle returns a new theme with overridden value.
func (t Theme) WithDescriptionMessageStyle(style Style) Theme {
	t.DescriptionMessageStyle = style
	return t
}

// WithDescriptionPrefixStyle returns a new theme with overridden value.
func (t Theme) WithDescriptionPrefixStyle(style Style) Theme {
	t.DescriptionPrefixStyle = style
	return t
}

// WithBulletListTextStyle returns a new theme with overridden value.
func (t Theme) WithBulletListTextStyle(style Style) Theme {
	t.BulletListTextStyle = style
	return t
}

// WithBulletListBulletStyle returns a new theme with overridden value.
func (t Theme) WithBulletListBulletStyle(style Style) Theme {
	t.BulletListBulletStyle = style
	return t
}

// WithLetterStyle returns a new theme with overridden value.
func (t Theme) WithLetterStyle(style Style) Theme {
	t.LetterStyle = style
	return t
}

// WithDebugMessageStyle returns a new theme with overridden value.
func (t Theme) WithDebugMessageStyle(style Style) Theme {
	t.DebugMessageStyle = style
	return t
}

// WithDebugPrefixStyle returns a new theme with overridden value.
func (t Theme) WithDebugPrefixStyle(style Style) Theme {
	t.DebugPrefixStyle = style
	return t
}

// WithTreeStyle returns a new theme with overridden value.
func (t Theme) WithTreeStyle(style Style) Theme {
	t.TreeStyle = style
	return t
}

// WithTreeTextStyle returns a new theme with overridden value.
func (t Theme) WithTreeTextStyle(style Style) Theme {
	t.TreeTextStyle = style
	return t
}

// WithBoxStyle returns a new theme with overridden value.
func (t Theme) WithBoxStyle(style Style) Theme {
	t.BoxStyle = style
	return t
}

// WithBoxTextStyle returns a new theme with overridden value.
func (t Theme) WithBoxTextStyle(style Style) Theme {
	t.BoxTextStyle = style
	return t
}

// WithBarLabelStyle returns a new theme with overridden value.
func (t Theme) WithBarLabelStyle(style Style) Theme {
	t.BarLabelStyle = style
	return t
}

// WithBarStyle returns a new theme with overridden value.
func (t Theme) WithBarStyle(style Style) Theme {
	t.BarStyle = style
	return t
}
