package slice

/*
Filter returns filtered slice according to the given function.

Usage:

	Filter([]int{1, 2, 3}, func(v int) bool { return v < 3 }) // []int{1, 2}
*/
func Filter[T any](slice []T, fn func(v T) bool) []T {
	newslice := make([]T, 0)

	for _, v := range slice {
		if fn(v) {
			newslice = append(newslice, v)
		}
	}

	return newslice
}
