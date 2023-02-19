package slice

/*
All ensures that all values from a given slice satisfies a given condition.

Usage:

	All([]int{1, 2, 3}, func(v int) bool { return v > 0 }) // true
*/
func All[T any](slice []T, fn func(v T) bool) bool {
	for _, v := range slice {
		if !fn(v) {
			return false
		}
	}

	return true
}
