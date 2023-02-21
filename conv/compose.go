package conv

// Deprecated: Use mapx.Compose instead
func Compose(vals ...any) map[any]any {
	m := make(map[any]any)

	for i := 0; i < len(vals); i += 2 {
		m[vals[i]] = vals[i+1]
	}

	return m
}
