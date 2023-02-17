package slice

/*
Limit makes a slice subset if bigger than a given limit.

Usage:

	Limit([]string{"a", "b", "c"}, 2) // []string{"a", "b"}
*/
func Limit[T any](slice []T, limit int) []T {
	if len(slice) > limit {
		return slice[:limit]
	}
	return slice
}
