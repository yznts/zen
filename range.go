package zen

// Filter returns filtered slice according to the given function.
//
// Usage:
//  Filter([]int{1, 2, 3}, func(v int) bool { return v < 3 }) // []int{1, 2}
func Filter[T any](slice []T, fn func(v T) bool) []T {
	var a []T
	for _, v := range slice {
		if fn(v) {
			a = append(a, v)
		}
	}
	return a
}

// Map returns a new slice with the results of applying the given function to each element in the given slice.
//
// Usage:
//  Map([]string{"asd", "qwe"}, func(v string) int { return len(v) }) // []int{3, 3}
func Map[T1 any, T2 any](slice []T1, fn func(v T1) T2) []T2 {
	a := make([]T2, len(slice))
	for i, v := range slice {
		a[i] = fn(v)
	}
	return a
}

// Range returns a new slice of integers in the given range (from, to).
//
// Usage:
// 	Range(1, 5) // []int{1, 2, 3, 4, 5}
func Range(from, to int) []int {
	a := make([]int, to-from+1)
	for i := 0; i <= to-from; i++ {
		a[i] = i + from
	}
	return a
}

// In returns true if the given value is in the given slice.
//
// Usage:
// 	In(1, []int{1, 2, 3}) // true
func In[T comparable](val T, slice []T) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

// InRuntime is a runtime analogue of In (made to be compatible with templates).
func InRuntime(val any, slice []any) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
