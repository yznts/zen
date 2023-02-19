package logic

/*
Or acts like "||" for values in any other language.
Unfortunately, in Go this operator only works for conditions.
Please note, this operator is just a function and not prevents execution.

Usage:

	zen.Or(0, 1) // 1
*/
func Or[T comparable](a, b T) T {
	var c T
	if a != c {
		return a
	}

	return b
}
