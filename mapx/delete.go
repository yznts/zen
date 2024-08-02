package mapx

/*
Delete removes given keys from a given map.

Usage:

	example := map[int]int{1: 2, 5: 6, 8: 10}
	mapx.Delete(example, 5, 8) // map[int]int{1: 2}
*/
func Delete[T1 comparable, T2 any](m map[T1]T2, keys ...T1) map[T1]T2 {
	for _, k := range keys {
		delete(m, k)
	}
	return m
}
