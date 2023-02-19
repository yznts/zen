package agg

import (
	"github.com/kyoto-framework/zen/v3/mathx"
	"golang.org/x/exp/constraints"
)

/*
Avg returns the average value of the given values.

Usage:

	agg.Avg(1, 2, 3) // 2
	agg.Avg([]float64{3.3, 6.6, 4.5}...) // 4.8
*/
func Avg[T constraints.Integer | constraints.Float](vals ...T) T {
	if len(vals) == 0 {
		panic("Avg() called with no arguments")
	}

	sum := mathx.Sum(vals...)

	return sum / T(len(vals))
}
