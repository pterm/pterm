package pterm

var (
	// ThemeDefault is the default theme used by PTerm.
	// If this variable is overwritten, the new value is used as default theme.
	ThemeDefault = Theme{
		PrimaryStyle:            NewStyle(FgCyan),
		SecondaryStyle:          NewStyle(FgLightMagenta),
		HighlightStyle:          NewStyle(Bold, FgYellow),
		InfoMessageStyle:        NewStyle(FgLightCyan),
		InfoPrefixStyle:         NewStyle(FgBlack, BgCyan),
		SuccessMessageStyle:     NewStyle(FgGreen),
		SuccessPrefixStyle:      NewStyle(FgBlack, BgGreen),
		WarningMessageStyle:     NewStyle(FgYellow),
		WarningPrefixStyle:      NewStyle(FgBlack, BgYellow),
		ErrorMessageStyle:       NewStyle(FgLightRed),
		ErrorPrefixStyle:        NewStyle(FgBlack, BgLightRed),
		FatalMessageStyle:       NewStyle(FgLightRed),
		FatalPrefixStyle:        NewStyle(FgBlack, BgLightRed),
		DescriptionMessageStyle: NewStyle(FgWhite),
		DescriptionPrefixStyle:  NewStyle(FgLightWhite, BgDarkGray),
		ScopeStyle:              NewStyle(FgGray),
		ProgressbarBarStyle:     NewStyle(FgLightCyan),
		ProgressbarTitleStyle:   NewStyle(FgLightCyan),
		HeaderTextStyle:         NewStyle(FgLightWhite, Bold),
		HeaderBackgroundStyle:   NewStyle(BgGray),
	}
)

// Theme for PTerm.
// Theme contains every Style used in PTerm. You can create own themes for your application or use one
// of the existing themes.
type Theme struct {
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
