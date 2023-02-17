package slice

/*
In returns true if the given value is in the given slice.

Usage:

	In(1, []int{1, 2, 3}) // true
*/
func In[T comparable](val T, slice []T) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

/*
InRuntime is a runtime version of In.
*/
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
