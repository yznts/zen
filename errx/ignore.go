package errx

/*
Ignore is a helper that wraps a call to a function returning value and error
and ignores if the error is non-nil.
*/
func Ignore[T any](val T, err error) T {
	return val
}

/*
IgnoreRuntime is a runtime version of errorsx.Ignore().
*/
func IgnoreRuntime(val any, err error) any {
	return val
}
