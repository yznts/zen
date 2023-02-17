package agg

import "golang.org/x/exp/constraints"

/*
Max returns the maximum value of the given values.

Usage:

	agg.Max(1, 2, 3) // 3
	agg.Max([]float64{2.3, 54.3, 3.5}...) // 54.3
*/
func Max[T constraints.Ordered](vals ...T) T {
	if len(vals) == 0 {
		panic("Max() called with no arguments")
	}
	max := vals[0]
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}
