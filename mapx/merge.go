package mapx

/*
Merge merges given maps into single map.
Each next map overrides previous one keys.

Usage:

	res := mapx.Merge(  // map[string]int{"a": 1, "b": 3, "c": 4}
		map[string]int{"a": 1, "b": 2},
		map[string]int{"b": 3, "c": 4},
	)
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
