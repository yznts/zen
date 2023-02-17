package async

/*
Filter returns filtered slice according to the given function.

Asynchronous version of slice.Filter. Please note, it's not always faster! Goroutines allocation have own cost.

Usage:

	async.Filter([]int{1, 2, 3}, func(v int) bool { return v > 1 }) // int{2, 3}
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
