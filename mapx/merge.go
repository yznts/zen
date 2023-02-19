package mapx

/*
Merge merges given maps into single map.
Each next map overrides previous one keys.

Usage:

	example := map[int]int{1: 2, 5: 6, 8: 10}
	mapx.Keys(example) // []int{1, 5, 8}
*/
func Merge[T1 comparable, T2 any](maps ...map[T1]T2) map[T1]T2 {
	newmap := map[T1]T2{}

	for _, m := range maps {
		for k, v := range m {
			newmap[k] = v
		}
	}

	return newmap
}
