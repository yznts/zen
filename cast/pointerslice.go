package cast

/*
PointerSlice is a function to cast a slice of any values ([]any)
to a slice of the given type pointers.

Usage:

	// Let's assume we have []any{1, 2, nil} in "values" variable.
	cast.PointerSlice(int)([]any{1, 2, nil}) // []*int{1, 2, nil}
*/
func PointerSlice[T any](slice []any) []*T {
	newslice := make([]*T, 0, len(slice))

	for _, v := range slice {
		if v == nil {
			newslice = append(newslice, nil)
		} else {
			cv := v.(T) //nolint:forcetypeassert // We expect a panic, if something is wrong
			newslice = append(newslice, &cv)
		}
	}

	return newslice
}
