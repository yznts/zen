package mapx

/*
Values extracts values from a given map.

Usage:

	example := map[int]int{1: 2, 5: 6, 8: 10}
	mapx.Values(example) // []int{2, 6, 10}
*/
func Values[T1 comparable, T2 any](m map[T1]T2) (values []T2) {
	for _, v := range m {
		values = append(values, v)
	}

	return
}
