package slice

//nolint:dupword
/*
Unique returns a new slice with the unique slice values.
Comparable value is defined by a given function.

Usage:

	Unique([]int{1, 2, 2, 3}, func(v int) int { return v }) // []int{1, 2, 3}
*/
func Unique[T1 any, T2 comparable](slice []T1, fn func(v T1) T2) []T1 {
	flags := map[T2]bool{}
	return Filter(slice, func(v T1) bool {
		compareval := fn(v)
		defer func() {
			flags[compareval] = true
		}()
		return !flags[compareval]
	})
}
