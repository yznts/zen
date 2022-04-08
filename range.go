package zen

func Filter[T any](slice []T, fn func(any) bool) []T {
	var a []T
	for _, v := range slice {
		if fn(v) {
			a = append(a, v)
		}
	}
	return a
}

func Map[T1 any, T2 any](slice []T1, fn func(T1) T2) []T2 {
	a := make([]T2, len(slice))
	for i, v := range slice {
		a[i] = fn(v)
	}
	return a
}

func Range(from, to int) []int {
	a := make([]int, to-from+1)
	for i := 0; i <= to-from; i++ {
		a[i] = i + from
	}
	return a
}

func In(val any, slice []any) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
