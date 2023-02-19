package mapx

/*
Keys extracts keys from a given map.

Usage:

	example := map[int]int{1: 2, 5: 6, 8: 10}
	mapx.Keys(example) // []int{1, 5, 8}
*/
func Keys[T1 comparable, T2 any](m map[T1]T2) (keys []T1) {
	for k := range m {
		keys = append(keys, k)
	}

	return
}
