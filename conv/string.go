package conv

import (
	"fmt"
	"strconv"
)

/*
String converts the given value to a string.
Uses fmt.Sprintf("%v", val) as fallback for unknown types.

Usage:

	conv.String(true) // "true"
	conv.String(1) // "1"
*/
func String(val any) string {
	switch val := val.(type) {
	case bool:
		return strconv.FormatBool(val)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", val)
	case float32, float64:
		return fmt.Sprintf("%f", val)
	case string:
		return val
	case nil:
		return ""
	default:
		return fmt.Sprintf("%v", val)
	}
}
