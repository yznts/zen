package arithmetic

import (
	"github.com/kyoto-framework/zen/v3/cast"
	"golang.org/x/exp/constraints"
)

/*
Mul returns the multiplication result of the given values.

Usage:

	arithmetic.Mul(1, 2, 3) // 6
	arithmetic.Mul([]float64{3.4, 6.5, 4.3}...) // 95.03
*/
func Mul[T constraints.Integer | constraints.Float](vals ...T) T {
	if len(vals) < 2 {
		panic("Mul() requires at least 2 arguments")
	}
	var mul T = vals[0]
	for _, val := range vals[1:] {
		mul *= val
	}
	return mul
}

/*
MulRuntime is a runtime version of arithmetic.Mul.
*/
func MulRuntime(vals ...any) any {
	switch vals[0].(type) {
	case int:
		return Mul(cast.Slice[int](vals)...)
	case int8:
		return Mul(cast.Slice[int8](vals)...)
	case int16:
		return Mul(cast.Slice[int16](vals)...)
	case int32:
		return Mul(cast.Slice[int32](vals)...)
	case int64:
		return Mul(cast.Slice[int64](vals)...)
	case uint:
		return Mul(cast.Slice[uint](vals)...)
	case uint8:
		return Mul(cast.Slice[uint8](vals)...)
	case uint16:
		return Mul(cast.Slice[uint16](vals)...)
	case uint32:
		return Mul(cast.Slice[uint32](vals)...)
	case uint64:
		return Mul(cast.Slice[uint64](vals)...)
	case float32:
		return Mul(cast.Slice[float32](vals)...)
	case float64:
		return Mul(cast.Slice[float64](vals)...)
	default:
		panic("unknown type for MulRuntime()")
	}
}
