/*
	-

	Errors

	Zen provides a simple  helper functions to handle error cases.
	F.e. Must, that wraps a call to a function returning value and error,
	and panics if the error is non-nil.

	Example:

		func main() {
			zen.Must(strconv.Atoi("asd")) // panic
		}
*/
package zen

// Must is a helper that wraps a call to a function returning value and error
// and panics if the error is non-nil.
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

// Ignore is a helper that wraps a call to a function returning value and error
// and ignores if the error is non-nil.
func Ingore[T any](val T, err error) T {
	return val
}
