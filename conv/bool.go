package conv

import "reflect"

//nolint:cyclop
/*
Bool converts the given value to boolean.
Should be used in a known environment, when values are expected to be correct.
Panics in case of failure.

Usage:

	conv.Bool(5.4) // true
	conv.Bool("") // false
	conv.Bool(0) // false
	conv.Bool(struct{}) // panic!
*/
func Bool(val any) bool {
	switch val := val.(type) {
	case bool:
		return val
	case *bool:
		if val == nil {
			return false
		}
		// Return actual value
		return *val
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return val != 0
	case *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *float32, *float64:
		if val == nil {
			return false
		}
		// Return integer boolean representation
		return reflect.ValueOf(val).Int() != 0
	case string:
		return val != ""
	case *string:
		if val == nil {
			return false
		}
		// Return string boolean representation
		return reflect.ValueOf(val).String() != ""
	case nil:
		return false
	default:
		panic("unknown type for Bool()")
	}
}
