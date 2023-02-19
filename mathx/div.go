//nolint:dupl
package mathx

import (
	"github.com/kyoto-framework/zen/v3/cast"
	"golang.org/x/exp/constraints"
)

//nolint:gomnd
/*
Div returns the division of the first value to the second and later.
Makes no sense in unual use-cases, but useful in templates.

Usage:

	arithmetic.Div(5, 2, 2) // 1.25
	arithmetic.Div([]int{5, 2, 2}...) // 1.25
*/
func Div[T constraints.Integer | constraints.Float](vals ...T) T {
	if len(vals) < 2 {
		panic("Div() requires at least 2 arguments")
	}

	div := vals[0]
	for _, val := range vals[1:] {
		div /= val
	}

	return div
}

//nolint:cyclop
/*
DivRuntime is a runtime version of arithmetic.Div.
*/
func DivRuntime(vals ...any) any {
	switch vals[0].(type) {
	case int:
		return Div(cast.Slice[int](vals)...)
	case int8:
		return Div(cast.Slice[int8](vals)...)
	case int16:
		return Div(cast.Slice[int16](vals)...)
	case int32:
		return Div(cast.Slice[int32](vals)...)
	case int64:
		return Div(cast.Slice[int64](vals)...)
	case uint:
		return Div(cast.Slice[uint](vals)...)
	case uint8:
		return Div(cast.Slice[uint8](vals)...)
	case uint16:
		return Div(cast.Slice[uint16](vals)...)
	case uint32:
		return Div(cast.Slice[uint32](vals)...)
	case uint64:
		return Div(cast.Slice[uint64](vals)...)
	case float32:
		return Div(cast.Slice[float32](vals)...)
	case float64:
		return Div(cast.Slice[float64](vals)...)
	default:
		panic("unknown type for DivRuntime()")
	}
}
