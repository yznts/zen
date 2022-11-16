package zen

// Index finds element index according to the given function.
// If nothing found, returns -1.
//
// Usage:
//
//	Index([]int{1, 2, 3}, func(v int) bool { return v > 2 }) // 2
func Index[T any](slice []T, fn func(v T) bool) int {
	for i, v := range slice {
		if fn(v) {
			return i
		}
	}
	return -1
}

// Filter returns filtered slice according to the given function.
//
// Usage:
//
//	Filter([]int{1, 2, 3}, func(v int) bool { return v < 3 }) // []int{1, 2}
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
//
//	Map([]string{"asd", "qwe"}, func(v string) int { return len(v) }) // []int{3, 3}
func Map[T1 any, T2 any](slice []T1, fn func(v T1) T2) []T2 {
	a := make([]T2, len(slice))
	for i, v := range slice {
		a[i] = fn(v)
	}
	return a
}

// Unique returns a new slice with the unique slice values.
// Comparable value is defined by a given function.
//
// Usage:
//
//	Unique([]int{1, 2, 2, 3}, func(v int) int { return v }) // []int{1, 2, 3}
func Unique[T1 any, T2 comparable](slice []T1, fn func(v T1) T2) []T1 {
	flags := map[T2]bool{}
	return Filter(slice, func(v T1) bool {
		compareval := fn(v)
		defer func() {
			flags[compareval] = true
		}()
		return !flags[compareval]
	})
}

// Range returns a new slice of integers in the given range (from, to).
//
// Usage:
//
//	Range(1, 5) // []int{1, 2, 3, 4, 5}
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
//
//	In(1, []int{1, 2, 3}) // true
func In[T comparable](val T, slice []T) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

// InRuntime is a runtime analogue of In (made to be compatible with templates).
func InRuntime(val any, slice any) bool {
	switch val := val.(type) {
	case int:
		return In(val, slice.([]int))
	case int8:
		return In(val, slice.([]int8))
	case int16:
		return In(val, slice.([]int16))
	case int32:
		return In(val, slice.([]int32))
	case int64:
		return In(val, slice.([]int64))
	case uint:
		return In(val, slice.([]uint))
	case uint8:
		return In(val, slice.([]uint8))
	case uint16:
		return In(val, slice.([]uint16))
	case uint32:
		return In(val, slice.([]uint32))
	case uint64:
		return In(val, slice.([]uint64))
	case float32:
		return In(val, slice.([]float32))
	case float64:
		return In(val, slice.([]float64))
	case string:
		return In(val, slice.([]string))
	default:
		panic("unknown type for InRuntime()")
	}
}

// Pop takes an last element from a slice (with deletion), or with a given index.
//
// Usage:
//
//	a := []int{1, 2, 3}
//	b := Pop(a)     // 3
//	fmt.println(a)  // []int{1, 2}
func Pop[T any](slice []T, index ...int) ([]T, T) {
	// If index is not specified, pop last element
	if len(index) == 0 {
		index = append(index, len(slice)-1)
	}
	// If index is out of range, panic
	if index[0] > len(slice)-1 {
		panic("index out of range")
	}
	// Extract value
	value := slice[index[0]]
	// Truncate slice
	tslice := append(slice[:index[0]], slice[index[0]+1:]...)
	// Return value
	return tslice, value
}

// Insert injects a provided value into slice on the given index.
//
// Usage:
//
//	Insert([]string{"b", "c"}, 0, "a") // []string{"a", "b", "c"}
func Insert[T any](slice []T, index int, value T) []T {
	if len(slice) == index {
		return append(slice, value)
	}
	_slice := append(slice[:index+1], slice[index:]...)
	_slice[index] = value
	return _slice
}

// Last takes a last element of a given slice.
//
// Usage:
//
//	Last([]string{"a", "b", "c"}) // "c"
func Last[T any](slice []T) T {
	return slice[len(slice)-1]
}

// Chunks generates a chunks with a given size from a given slice.
//
// Usage:
//
//	Chunks([]int{1, 2, 3, 4}, 2) // [][]int{ []int{1, 2}, []int{3, 4} }
func Chunks[T any](slice []T, size int) [][]T {
	var chunks [][]T
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// Any ensures that at least one value from a given slice satisfies a given condition.
//
// Usage:
//
//	Any([]int{-1, 0, 1}, func(v int) bool { return v < 0 }) // true
func Any[T any](slice []T, fn func(v T) bool) bool {
	for _, v := range slice {
		if fn(v) {
			return true
		}
	}
	return false
}

// All ensures that all values from a given slice satisfies a given condition.
//
// Usage:
//
//	All([]int{1, 2, 3}, func(v int) bool { return v > 0 }) // true
func All[T any](slice []T, fn func(v T) bool) bool {
	for _, v := range slice {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Cartesian makes a product of two or more sets.
//
// Usage:
//
//	Cartesian([]int{1, 2}, []int{3, 4}) // [[1 3] [1 4] [2 3] [2 4]]
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
