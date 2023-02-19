package slice

/*
Insert injects a provided value into slice on the given index.

Usage:

	Insert([]string{"b", "c"}, 0, "a") // []string{"a", "b", "c"}
*/
func Insert[T any](slice []T, index int, value T) []T {
	if len(slice) == index {
		return append(slice, value)
	}

	newslice := append(slice[:index+1], slice[index:]...) //nolint:gocritic
	newslice[index] = value

	return newslice
}
