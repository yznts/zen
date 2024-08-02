package async

/*
Filter returns filtered slice according to the given function.
Same as slice.Filter, but executed in parallel.
Useful for cases when comparison function becomes expensive.

Usage:

	Filter(slice.Range(1, 1000), func(v int) bool { return v < 500 })
*/
func Filter[T any](slice []T, fn func(v T) bool) []T {
	sliceflag := Map(slice, fn)

	slicenew := make([]T, 0, len(slice))
	for i, v := range slice {
		if sliceflag[i] {
			slicenew = append(slicenew, v)
		}
	}

	return slicenew
}
