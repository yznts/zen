package slice

//nolint:dupword
/*
FilterIn returns a function that checks if the given value is in the given top-level values slice.
Works only for comparable types.
Useful for filtering slices with slice.Filter.

Usage:

	slice.Filter([]int{1, 2, 3, 3, 4, 5}, slice.FilterIn(1, 2, 3)) // []int{1, 2, 3, 3}
*/
func FilterIn[T comparable](values ...T) func(T) bool {
	return func(t T) bool {
		return In(t, values)
	}
}

//nolint:dupword
/*
FilterNotIn returns a function that checks if the given value is not in the given top-level values slice.
Works only for comparable types.
Useful for filtering slices with slice.Filter.

Usage:

	slice.Filter([]int{1, 2, 3, 3, 4, 5}, slice.FilterNotIn(1, 2, 3)) // []int{4, 5}
*/
func FilterNotIn[T comparable](values ...T) func(T) bool {
	return func(t T) bool {
		return !In(t, values)
	}
}
