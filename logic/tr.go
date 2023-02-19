package logic

/*
Tr acts like a ternary operator in other languages.
Unfortunately, Go doesn't have this operator.
Please note, this operator is just a function and not prevents execution.

Usage:

	zen.Tr(false, "asd", "qwe") // string{"qwe"}
*/
func Tr[T comparable](condition bool, v1, v2 T) T {
	if condition {
		return v1
	}

	return v2
}
