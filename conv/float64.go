package conv

import "strconv"

/*
Float64 converts the given value to float64.
Should be used in a known environment, when values are expected to be correct.
Panics in case of failure.

Usage:

	conv.Float64("5") // 5.0
	conv.Float64(true) // 1.0
	conv.Float64("qwe") // panic!
*/
func Float64(val any) float64 {
	switch val := val.(type) {
	case bool:
		if val {
			return 1
		}
		return 0
	case int:
		return float64(val)
	case int8:
		return float64(val)
	case int16:
		return float64(val)
	case int32:
		return float64(val)
	case int64:
		return float64(val)
	case uint:
		return float64(val)
	case uint8:
		return float64(val)
	case uint16:
		return float64(val)
	case uint32:
		return float64(val)
	case uint64:
		return float64(val)
	case float32:
		return float64(val)
	case float64:
		return val
	case *int:
		if val == nil {
			return 0
		}
		return float64(*val)
	case *int8:
		if val == nil {
			return 0
		}
		return float64(*val)
	case *int16:
		if val == nil {
			return 0
		}
		return float64(*val)
	case *int32:
		if val == nil {
			return 0
		}
		return float64(*val)
	case *int64:
		if val == nil {
			return 0
		}
		return float64(*val)
	case *uint:
		if val == nil {
			return 0
		}
		return float64(*val)
	case *uint8:
		if val == nil {
			return 0
		}
		return float64(*val)
	case *uint16:
		if val == nil {
			return 0
		}
		return float64(*val)
	case *uint32:
		if val == nil {
			return 0
		}
		return float64(*val)
	case *uint64:
		if val == nil {
			return 0
		}
		return float64(*val)

	case *float32:
		if val == nil {
			return 0
		}
		return float64(*val)

	case *float64:
		if val == nil {
			return 0
		}
		return *val
	case string:
		i, err := strconv.ParseFloat(val, 64)
		if err != nil {
			panic(err)
		}
		return i
	case *string:
		if val == nil {
			return 0
		}
		i, err := strconv.ParseFloat(*val, 64)
		if err != nil {
			panic(err)
		}
		return i
	case nil:
		return 0
	default:
		panic("unknown type for Float64()")
	}
}
