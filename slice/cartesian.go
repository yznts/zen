package slice

/*
Cartesian makes a product of two or more sets.

Usage:

	Cartesian([]int{1, 2}, []int{3, 4}) // [[1 3] [1 4] [2 3] [2 4]]
*/
func Cartesian[T any](slices ...[]T) (res [][]T) {
	if len(slices) == 0 {
		return [][]T{nil}
	}

	r := Cartesian(slices[1:]...)
	for _, e := range slices[0] {
		for _, p := range r {
			res = append(res, append([]T{e}, p...))
		}
	}

	return
}
