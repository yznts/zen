package slice

/*
Chunks generates a chunks with a given size from a given slice.

Usage:

	Chunks([]int{1, 2, 3, 4}, 2) // [][]int{ []int{1, 2}, []int{3, 4} }
*/
func Chunks[T any](slice []T, size int) [][]T {
	var chunks [][]T
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}
