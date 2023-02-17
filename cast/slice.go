package cast

/*
Slice is a function to cast a slice of any values ([]any)
to a slice of the given type.

Usage:

	// Let's assume we have []any{1, 2, 3} in "values" variable.
	cast.Slice[int](values) // []int{1, 2, 3}
*/
func Slice[T any](slice []any) []T {
	_slice := make([]T, 0, len(slice))
	for _, el := range slice {
		_slice = append(_slice, el.(T))
	}
	return _slice
}
