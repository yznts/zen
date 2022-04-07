package zen

import "golang.org/x/exp/constraints"

// Min returns the minimum value of the given values
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

// Max returns the maximum value of the given values
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

// Avg returns the average value of the given values
func Avg[T constraints.Integer | constraints.Float](vals ...T) T {
	if len(vals) == 0 {
		panic("Avg() called with no arguments")
	}
	var sum T = Sum(vals...)
	return sum / T(len(vals))
}
