package zen

import (
	"golang.org/x/exp/constraints"
)

// Sum returns the sum of the given values
func Sum[T constraints.Integer | constraints.Float | string](vals ...T) T {
	var sum T
	for _, val := range vals {
		sum += val
	}
	return sum
}

// SumRuntime is a runtime analogue of Sum (made to be compatible with templates)
func SumRuntime(vals ...any) any {
	switch vals[0].(type) {
	case int:
		return Sum(CastSlice[int](vals)...)
	case int8:
		return Sum(CastSlice[int8](vals)...)
	case int16:
		return Sum(CastSlice[int16](vals)...)
	case int32:
		return Sum(CastSlice[int32](vals)...)
	case int64:
		return Sum(CastSlice[int64](vals)...)
	case uint:
		return Sum(CastSlice[uint](vals)...)
	case uint8:
		return Sum(CastSlice[uint8](vals)...)
	case uint16:
		return Sum(CastSlice[uint16](vals)...)
	case uint32:
		return Sum(CastSlice[uint32](vals)...)
	case uint64:
		return Sum(CastSlice[uint64](vals)...)
	case float32:
		return Sum(CastSlice[float32](vals)...)
	case float64:
		return Sum(CastSlice[float64](vals)...)
	case string:
		return Sum(CastSlice[string](vals)...)
	default:
		panic("unknown type for SumRuntime()")
	}
}

// Sub returns the subtraction of the given values from the first one
func Sub[T constraints.Integer | constraints.Float](vals ...T) T {
	var sub T = vals[0]
	for _, val := range vals[1:] {
		sub -= val
	}
	return sub
}

// SubRuntime is a runtime analogue of Sub (made to be compatible with templates)
func SubRuntime(vals ...any) any {
	switch vals[0].(type) {
	case int:
		return Sub(CastSlice[int](vals)...)
	case int8:
		return Sub(CastSlice[int8](vals)...)
	case int16:
		return Sub(CastSlice[int16](vals)...)
	case int32:
		return Sub(CastSlice[int32](vals)...)
	case int64:
		return Sub(CastSlice[int64](vals)...)
	case uint:
		return Sub(CastSlice[uint](vals)...)
	case uint8:
		return Sub(CastSlice[uint8](vals)...)
	case uint16:
		return Sub(CastSlice[uint16](vals)...)
	case uint32:
		return Sub(CastSlice[uint32](vals)...)
	case uint64:
		return Sub(CastSlice[uint64](vals)...)
	case float32:
		return Sub(CastSlice[float32](vals)...)
	case float64:
		return Sub(CastSlice[float64](vals)...)
	default:
		panic("unknown type for SubRuntime()")
	}
}

// Mul returns the multiplication of the given values
func Mul[T constraints.Integer | constraints.Float](vals ...T) T {
	var mul T = vals[0]
	for _, val := range vals[1:] {
		mul *= val
	}
	return mul
}

// MulRuntime is a runtime analogue of Mul (made to be compatible with templates)
func MulRuntime(vals ...any) any {
	switch vals[0].(type) {
	case int:
		return Mul(CastSlice[int](vals)...)
	case int8:
		return Mul(CastSlice[int8](vals)...)
	case int16:
		return Mul(CastSlice[int16](vals)...)
	case int32:
		return Mul(CastSlice[int32](vals)...)
	case int64:
		return Mul(CastSlice[int64](vals)...)
	case uint:
		return Mul(CastSlice[uint](vals)...)
	case uint8:
		return Mul(CastSlice[uint8](vals)...)
	case uint16:
		return Mul(CastSlice[uint16](vals)...)
	case uint32:
		return Mul(CastSlice[uint32](vals)...)
	case uint64:
		return Mul(CastSlice[uint64](vals)...)
	case float32:
		return Mul(CastSlice[float32](vals)...)
	case float64:
		return Mul(CastSlice[float64](vals)...)
	default:
		panic("unknown type for MulRuntime()")
	}
}

// Div returns the division of the given values to the first one
func Div[T constraints.Integer | constraints.Float](vals ...T) T {
	var div T = vals[0]
	for _, val := range vals[1:] {
		div /= val
	}
	return div
}

// DivRuntime is a runtime analogue of Div (made to be compatible with templates)
func DivRuntime(vals ...any) any {
	switch vals[0].(type) {
	case int:
		return Div(CastSlice[int](vals)...)
	case int8:
		return Div(CastSlice[int8](vals)...)
	case int16:
		return Div(CastSlice[int16](vals)...)
	case int32:
		return Div(CastSlice[int32](vals)...)
	case int64:
		return Div(CastSlice[int64](vals)...)
	case uint:
		return Div(CastSlice[uint](vals)...)
	case uint8:
		return Div(CastSlice[uint8](vals)...)
	case uint16:
		return Div(CastSlice[uint16](vals)...)
	case uint32:
		return Div(CastSlice[uint32](vals)...)
	case uint64:
		return Div(CastSlice[uint64](vals)...)
	case float32:
		return Div(CastSlice[float32](vals)...)
	case float64:
		return Div(CastSlice[float64](vals)...)
	default:
		panic("unknown type for DivRuntime()")
	}
}
