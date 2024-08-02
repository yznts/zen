package slice

/*
Contains checks if a value is in a slice.

Usage:

	Contains([]int{1, 2, 3}, 1) // true
	Contains([]int{1, 2, 3}, 4) // false
*/
func Contains[T comparable](slice []T, val T) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}

	return false
}

//nolint:cyclop,forcetypeassert
/*
ContainsRuntime is a runtime version of Contains.
*/
func ContainsRuntime(slice any, val any) bool {
	switch val := val.(type) {
	case int:
		return Contains(slice.([]int), val)
	case int8:
		return Contains(slice.([]int8), val)
	case int16:
		return Contains(slice.([]int16), val)
	case int32:
		return Contains(slice.([]int32), val)
	case int64:
		return Contains(slice.([]int64), val)
	case uint:
		return Contains(slice.([]uint), val)
	case uint8:
		return Contains(slice.([]uint8), val)
	case uint16:
		return Contains(slice.([]uint16), val)
	case uint32:
		return Contains(slice.([]uint32), val)
	case uint64:
		return Contains(slice.([]uint64), val)
	case float32:
		return Contains(slice.([]float32), val)
	case float64:
		return Contains(slice.([]float64), val)
	case string:
		return Contains(slice.([]string), val)
	default:
		panic("unknown type for ContainsRuntime()")
	}
}
