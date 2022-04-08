package zen

// CastSice is a function to cast a slice of any values ([]any)
// to a slice of the given type.
//
// Usage:
//  zen.CastSlice[int]([]any{1, 2, 3}) // []int{1, 2, 3}
func CastSlice[T any](slice []any) []T {
	_slice := make([]T, len(slice))
	for _, el := range slice {
		_slice = append(_slice, el.(T))
	}
	return _slice
}
