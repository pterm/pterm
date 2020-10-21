package internal

// WithBoolean helps an option setter (WithXXX(b ...bool) to return true, if no boolean is set, but false if it's explicitly set to false.
func WithBoolean(b []bool) bool {
	if len(b) == 0 {
		b = append(b, true)
	}
	return b[0]
}
