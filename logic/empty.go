package logic

/*
Empty returns true if the given value is equal to the default empty value of the given type.
Works only for comparable types.

Usage:

	logic.Empty("") // true
	logic.Empty(0) // true
	logic.Empty(42) // false
*/
func Empty[T comparable](value T) bool {
	// Initialize default empty value with a given type
	var empty T
	// Return comparison result
	return value == empty
}

/*
Empty returns true if the given value is not equal to the default empty value of the given type.
Works only for comparable types.

Usage:

	logic.Empty("") // false
	logic.Empty(0) // false
	logic.Empty(42) // true
	slice.Filter([]int{0, 0, 1, 4, 5}, logic.NotEmpty[int]) // []int{1, 4, 5}
*/
func NotEmpty[T comparable](value T) bool {
	return !Empty(value)
}
