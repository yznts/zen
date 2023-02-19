package conv

/*
Compose makes a map with the given keys and values.
Useful as a template function to pass multiple values to a template.
Based on even and odd values.

Usage:

	// Code
	zen.Compose("foo", 1, "bar", 2) // map[any]any{"foo": 1, "bar": 2}
*/
func Compose(vals ...any) map[any]any {
	m := make(map[any]any)

	for i := 0; i < len(vals); i += 2 {
		m[vals[i]] = vals[i+1]
	}

	return m
}
