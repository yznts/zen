package agg

import "golang.org/x/exp/constraints"

/*
Min returns the minimum value of the given values.

Usage:

	agg.Min(1, 2, 3) // 1
	agg.Min([]float64{2.3, 54.3, 3.5}...) // 2.3
*/
func Min[T constraints.Ordered](vals ...T) T {
	if len(vals) == 0 {
		panic("Min() called with no arguments")
	}

	min := vals[0]
	for _, v := range vals {
		if v < min {
			min = v
		}
	}

	return min
}
