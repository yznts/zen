package conv

/*
Ptr makes a pointer for a given value.

Usage:

	conv.Ptr(1) // *int{1}
*/
func Ptr[T any](val T) *T {
	return &val
}

/*
PtrRuntime is a runtime version of conv.Ptr.

Usage:

	// Please note, you'll receive "any" type which you'll need to cast.
	conv.PtrRuntime(1) // any, which holds *int{1}
*/
func PtrRuntime(val any) any {
	return &val
}
