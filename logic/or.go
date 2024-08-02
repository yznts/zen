package logic

/*
Or acts like "||" for values in any other language.
Unfortunately, in Go this operator only works for conditions.
Please note, this operator is just a function and not prevents execution.

Usage:

	zen.Or(0, 1) // 1
*/
func Or[T comparable](vals ...T) T {
	var def T
	for _, v := range vals {
		if v != def {
			return v
		}
	}

	return def
}
