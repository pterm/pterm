package internal

func WithBoolean(b []bool) bool {
	if len(b) == 0 {
		b = append(b, true)
	}
	return b[0]
}
