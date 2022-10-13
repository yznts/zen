package zen

// Or acts like "||" for values in any other language.
// Unfortuantely, in Go this operator only works for conditions.
//
// Usage:
//
//	zen.Or(0, 1) // 1
func Or[T comparable](a, b T) T {
	var c T
	if a != c {
		return a
	}
	return b
}

// Tr acts like a ternary operator in other languages.
// Unfortuantely, Go doesn't have this operator.
//
// Usage:
//
//	zen.Tr(false, "asd", "qwe") // string{"qwe"}
func Tr[T comparable](condition bool, v1, v2 T) T {
	if condition {
		return v1
	}
	return v2
}
