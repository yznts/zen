package slice

/*
Last takes a last element of a given slice.
In that way, you don't need to mess with len(slice)-1.

Usage:

	Last([]string{"a", "b", "c"}) // "c"
*/
func Last[T any](slice []T) T {
	return slice[len(slice)-1]
}
