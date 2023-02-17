package slice

/*
Filter returns filtered slice according to the given function.

Usage:

	Filter([]int{1, 2, 3}, func(v int) bool { return v < 3 }) // []int{1, 2}
*/
func Filter[T any](slice []T, fn func(v T) bool) []T {
	var a []T
	for _, v := range slice {
		if fn(v) {
			a = append(a, v)
		}
	}
	return a
}
