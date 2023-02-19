package slice

/*
Index finds element index according to the given function.
If nothing found, returns -1.

Usage:

	Index([]int{1, 2, 3}, func(v int) bool { return v > 2 }) // 2
*/
func Index[T any](slice []T, fn func(v T) bool) int {
	for i, v := range slice {
		if fn(v) {
			return i
		}
	}

	return -1
}
