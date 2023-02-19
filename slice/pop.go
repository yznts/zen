package slice

/*
Pop takes an last element from a slice (with deletion), or with a given index.

Usage:

	a := []int{1, 2, 3}
	b := Pop(a)     // 3
	fmt.println(a)  // []int{1, 2}
*/
func Pop[T any](slice []T, index ...int) ([]T, T) {
	// If index is not specified, pop last element
	if len(index) == 0 {
		index = append(index, len(slice)-1)
	}
	// If index is out of range, panic
	if index[0] > len(slice)-1 {
		panic("index out of range")
	}
	// Extract value
	value := slice[index[0]]
	// Truncate slice
	newslice := append(slice[:index[0]], slice[index[0]+1:]...) //nolint:gocritic
	// Return value
	return newslice, value
}
