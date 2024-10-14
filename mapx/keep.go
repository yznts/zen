package mapx

import "github.com/yznts/zen/v3/slice"

/*
Keep reverses logic of mapx.Delete.
Removes keys from a given map, if key not found in the given keys.

Usage:

	example := map[int]int{1: 2, 5: 6, 8: 10}
	mapx.Keep(example, 5, 8) // map[int]int{5: 6, 8: 10}
*/
func Keep[T1 comparable, T2 any](m map[T1]T2, keys ...T1) map[T1]T2 {
	for k := range m {
		if !slice.Contains(keys, k) {
			delete(m, k)
		}
	}
	return m
}
