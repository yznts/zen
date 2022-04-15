package zen

// Or acts like "||" for values in any other language.
// Unfortuantely, in Go this operator only works for conditions.
//
// Usage:
//  zen.Or(0, 1) // 1
func Or[T comparable](a, b T) T {
	var c T
	if a != c {
		return a
	}
	return b
}
