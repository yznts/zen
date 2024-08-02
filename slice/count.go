package slice

func Count[T any](slice []T, fn func(v T) bool) int {
	count := 0

	for _, v := range slice {
		if fn(v) {
			count++
		}
	}

	return count
}
