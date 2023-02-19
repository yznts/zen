package async

/*
Filter returns filtered slice according to the given function.

Asynchronous version of slice.Filter. Please note, it's not always faster! Goroutines allocation have own cost.

Usage:

	async.Filter([]int{1, 2, 3}, func(v int) bool { return v > 1 }) // int{2, 3}
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
