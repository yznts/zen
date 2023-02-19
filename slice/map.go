package slice

/*
Map returns a new slice with the results of applying the given function to each element in the given slice.

Usage:

	Map([]string{"asd", "qwe"}, func(v string) int { return len(v) }) // []int{3, 3}
*/
func Map[T1 any, T2 any](slice []T1, fn func(v T1) T2) []T2 {
	a := make([]T2, len(slice))
	for i, v := range slice {
		a[i] = fn(v)
	}

	return a
}
