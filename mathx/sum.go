package mathx

import (
	"go.kyoto.codes/zen/v3/cast"
	"golang.org/x/exp/constraints"
)

//nolint:gomnd
/*
Sum returns the sum of the given values.

Usage:

	mathx.Sum(4, 2, 1) // 7
	mathx.Sum([]float64{5.4, 3, 1}...) // 9.4
*/
func Sum[T constraints.Integer | constraints.Float | string](vals ...T) T {
	if len(vals) < 2 {
		panic("Sum() requires at least 2 arguments")
	}

	var sum T
	for _, val := range vals {
		sum += val
	}

	return sum
}

//nolint:cyclop
/*
SumRuntime is a runtime version of arithmetic.Sum.
*/
func SumRuntime(vals ...any) any {
	switch vals[0].(type) {
	case int:
		return Sum(cast.Slice[int](vals)...)
	case int8:
		return Sum(cast.Slice[int8](vals)...)
	case int16:
		return Sum(cast.Slice[int16](vals)...)
	case int32:
		return Sum(cast.Slice[int32](vals)...)
	case int64:
		return Sum(cast.Slice[int64](vals)...)
	case uint:
		return Sum(cast.Slice[uint](vals)...)
	case uint8:
		return Sum(cast.Slice[uint8](vals)...)
	case uint16:
		return Sum(cast.Slice[uint16](vals)...)
	case uint32:
		return Sum(cast.Slice[uint32](vals)...)
	case uint64:
		return Sum(cast.Slice[uint64](vals)...)
	case float32:
		return Sum(cast.Slice[float32](vals)...)
	case float64:
		return Sum(cast.Slice[float64](vals)...)
	case string:
		return Sum(cast.Slice[string](vals)...)
	default:
		panic("unknown type for SumRuntime()")
	}
}
