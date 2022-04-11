package zen

// Must is a helper that wraps a call to a function returning
// and panics if the error is non-nil.
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}
