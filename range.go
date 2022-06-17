/*
	-

	Range

	Zen provides some generic functions for basic slice operations.

	Examples:

		func main() {
			// Creating an integers slice with a Range function
			var slice = zen.Range(1, 5) // []int{1, 2, 3, 4, 5}

			// Filtering
			zen.Filter(slice, func(v int) bool { return v < 3 }) // []int{1, 2}

			// Creating a new slice, based on existing one with a Map function
			zen.Map(slice, func(v int) int { return v * 2 }) // []int{2, 4, 6, 8, 10}

			// Checking if an element is in the slice
			zen.In(1, slice) // true

			// Pop an element at the given index from the slice (returns a new slice and the value)
			zen.Pop(slice, 1) // ([]int{1, 3, 4, 5}, 2)

			// Insert an element at the given index in the slice (returns a new slice)
			zen.Insert(slice, 1, 2) // []int{1, 2, 2, 3, 4, 5}

			// Get the last element from the slice
			zen.Last(slice) // 5

			// Check if any element in the slice matches the given function
			zen.Any(slice, func(v int) bool { return v == 2 }) // true

			// Check if all elements in the slice match the given function
			zen.All(slice, func(v int) bool { return v < 6 }) // true
		}
*/
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

func Insert[T any](slice []T, index int, value T) []T {
	if len(slice) == index {
		return append(slice, value)
	}
	_slice := append(slice[:index+1], slice[index:]...)
	_slice[index] = value
	return _slice
}

func Last[T any](slice []T) T {
	return slice[len(slice)-1]
}

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

func Any[T any](slice []T, fn func(v T) bool) bool {
	for _, v := range slice {
		if fn(v) {
			return true
		}
	}
	return false
}

func All[T any](slice []T, fn func(v T) bool) bool {
	for _, v := range slice {
		if !fn(v) {
			return false
		}
	}
	return true
}
