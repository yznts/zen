package slice

/*
Any ensures that at least one value from a given slice satisfies a given condition.

Usage:

	Any([]int{-1, 0, 1}, func(v int) bool { return v < 0 }) // true
*/
func Any[T any](slice []T, fn func(v T) bool) bool {
	for _, v := range slice {
		if fn(v) {
			return true
		}
	}

	return false
}
