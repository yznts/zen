package errorsx

/*
Must is a helper that wraps a call to a function returning value and error
and panics if the error is non-nil.
*/
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}

	return val
}

/*
MustRuntime is a runtime version of errorsx.Must().
*/
func MustRuntime(val any, err error) any {
	if err != nil {
		panic(err)
	}

	return val
}
