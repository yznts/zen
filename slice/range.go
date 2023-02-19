package slice

/*
Range returns a new slice of integers in the given range (from, to).

Usage:

	Range(1, 5) // []int{1, 2, 3, 4, 5}
*/
func Range(from, to int) []int {
	result := make([]int, to-from+1)
	for i := 0; i <= to-from; i++ {
		result[i] = i + from
	}

	return result
}
