//nolint:dupl
package mathx

import (
	"go.kyoto.codes/zen/v3/cast"
	"golang.org/x/exp/constraints"
)

//nolint:gomnd
/*
Sub returns the subtraction of the given values from the first one.

Usage:

	mathx.Sub(4, 2, 1) // 1
	mathx.Sub([]float64{5.4, 3, 1}...) // 1.4
*/
func Sub[T constraints.Integer | constraints.Float](vals ...T) T {
	if len(vals) < 2 {
		panic("Sub() requires at least 2 arguments")
	}

	sub := vals[0]
	for _, val := range vals[1:] {
		sub -= val
	}

	return sub
}

//nolint:cyclop
/*
SubRuntime is a runtime version of arithmetic.Sub.
*/
func SubRuntime(vals ...any) any {
	switch vals[0].(type) {
	case int:
		return Sub(cast.Slice[int](vals)...)
	case int8:
		return Sub(cast.Slice[int8](vals)...)
	case int16:
		return Sub(cast.Slice[int16](vals)...)
	case int32:
		return Sub(cast.Slice[int32](vals)...)
	case int64:
		return Sub(cast.Slice[int64](vals)...)
	case uint:
		return Sub(cast.Slice[uint](vals)...)
	case uint8:
		return Sub(cast.Slice[uint8](vals)...)
	case uint16:
		return Sub(cast.Slice[uint16](vals)...)
	case uint32:
		return Sub(cast.Slice[uint32](vals)...)
	case uint64:
		return Sub(cast.Slice[uint64](vals)...)
	case float32:
		return Sub(cast.Slice[float32](vals)...)
	case float64:
		return Sub(cast.Slice[float64](vals)...)
	default:
		panic("unknown type for SubRuntime()")
	}
}
