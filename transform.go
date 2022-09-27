/*
-

# Transform

Zen provides a number of functions that can be used to transform
data into different types and forms.
Most of these functions are working with base data types.

Examples:

	func main() {
		// Common data types transformations
		numptr := zen.Ptr(1) // *int{1}  Inline pointer
		boolval := zen.Bool(3) // bool{true}
		intval := zen.Int("5") // int{5}
		floatval := zen.Float64("6.5") // float64{6.5}
		strval := zen.String(7) // string{"7"}

		// Map composition (useful for templates)
		resmap := zen.Compose("foo", 1, "bar", "2") // map[any]any{"foo": 1, "bar": "2"}

		// JSON
		resjson := zen.JSON(resmap) // string{`{"foo": 1, "bar": "2"}`}

		// Base64
		resbase64 := zen.B64Enc(resjson) // string{`eyJmb28iOiAxLCAiYmFyIjogIjIifQ==`}
		resbase64dec := string(zen.B64Dec(resbase64)) // string{`{"foo": 1, "bar": "2"}`}
	}
*/
package zen

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// Ptr makes a pointer for a given value.
//
// Usage:
//
//	zen.Ptr(1) // *int 1
func Ptr[T any](val T) *T {
	return &val
}

// Int converts the given value to boolean.
//
// Usage:
//
//	zen.Bool(4.5) // true
func Bool(val any) bool {
	switch val := val.(type) {
	case bool:
		return val
	case *bool:
		if val == nil {
			return false
		}
		return *val
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return val != 0
	case *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *float32, *float64:
		if val == nil {
			return false
		}
		return reflect.ValueOf(val).Int() != 0
	case string:
		return val != ""
	case *string:
		if val == nil {
			return false
		}
		return reflect.ValueOf(val).String() != ""
	case nil:
		return false
	default:
		panic("unknown type for Bool()")
	}
}

// Int converts the given value to int.
//
// Usage:
//
//	zen.Int("123") // 123
func Int(val any) int {
	switch val := val.(type) {
	case bool:
		if val {
			return 1
		}
		return 0
	case int:
		return val
	case int8:
		return int(val)
	case int16:
		return int(val)
	case int32:
		return int(val)
	case int64:
		return int(val)
	case uint:
		return int(val)
	case uint8:
		return int(val)
	case uint16:
		return int(val)
	case uint32:
		return int(val)
	case uint64:
		return int(val)
	case float32:
		return int(val)
	case float64:
		return int(val)
	case *int:
		if val == nil {
			return 0
		}
		return *val
	case *int8:
		if val == nil {
			return 0
		}
		return int(*val)
	case *int16:
		if val == nil {
			return 0
		}
		return int(*val)
	case *int32:
		if val == nil {
			return 0
		}
		return int(*val)
	case *int64:
		if val == nil {
			return 0
		}
		return int(*val)
	case *uint:
		if val == nil {
			return 0
		}
		return int(*val)
	case *uint8:
		if val == nil {
			return 0
		}
		return int(*val)
	case *uint16:
		if val == nil {
			return 0
		}
		return int(*val)
	case *uint32:
		if val == nil {
			return 0
		}
		return int(*val)
	case *uint64:
		if val == nil {
			return 0
		}
		return int(*val)

	case *float32:
		if val == nil {
			return 0
		}
		return int(*val)

	case *float64:
		if val == nil {
			return 0
		}
		return int(*val)
	case string:
		i, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		return i
	case *string:
		if val == nil {
			return 0
		}
		i, err := strconv.Atoi(*val)
		if err != nil {
			panic(err)
		}
		return i
	case nil:
		return 0
	default:
		panic("unknown type for Int()")
	}
}

// Float64 converts the given value to float64.
//
// Usage:
//
//	zen.Float64("5") // 5.0
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
		return float64(val)
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
		return float64(*val)
	case string:
		i, err := strconv.ParseFloat(val, 64)
		if err != nil {
			panic(err)
		}
		return float64(i)
	case *string:
		if val == nil {
			return 0
		}
		i, err := strconv.ParseFloat(*val, 64)
		if err != nil {
			panic(err)
		}
		return float64(i)
	case nil:
		return 0
	default:
		panic("unknown type for Float64()")
	}
}

// String converts the given value to a string.
//
// Usage:
//
//	zen.String(1) // "1"
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

// Compose makes a map with the given keys and values.
// Useful as a template function to pass multiple values to a template.
// Based on even and odd values.
//
// Usage:
//
//	// Code
//	zen.Compose("foo", 1, "bar", 2) // map[any]any{"foo": 1, "bar": 2}
//	// Template
//	{{ compose "foo" 1 "bar" 2 }}
func Compose(vals ...any) map[any]any {
	m := make(map[any]any)
	for i := 0; i < len(vals); i += 2 {
		m[vals[i]] = vals[i+1]
	}
	return m
}

// JSON is a function that converts the given value to a JSON string.
// Useful as a template function.
//
// Usage:
//
//	// Code
//	zen.JSON(map[any]any{"foo": 1, "bar": 2}) // {"foo":1,"bar":2}
//	// Template
//	{{ json .Value }}
func JSON(val any) string {
	content, err := json.Marshal(val)
	if err != nil {
		panic(err)
	}
	return string(content)
}

// B64Enc converts the given value (bytes or string) to a base64 string.
//
// Usage:
//
//	// Code
//	zen.B64Enc([]byte("foo")) // "Zm9v"
//	Template
//	{{ b64enc "foo" }}
func B64Enc(val any) string {
	switch val := val.(type) {
	case []byte:
		return base64.StdEncoding.EncodeToString(val)
	case string:
		return base64.RawStdEncoding.EncodeToString([]byte(val))
	default:
		panic("unknown type for B64Enc()")
	}
}

// B64Dec converts the given base64 string to a value (bytes)
//
// Usage:
//
//	// Code
//	zen.B64Dec("Zm9v") // []byte("foo")
//	// Template
//	{{ b64dec "Zm9v" }}
func B64Dec(val string) []byte {
	data, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		panic(err)
	}
	return data
}
