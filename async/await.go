package async

/*
Interface for runtime await.
*/
type ImplementsAwaitRuntime interface {
	AwaitRuntime() (any, error)
}

/*
Await for a future object results.
*/
func Await[T any](f *Future[T]) (T, error) {
	return f.Await()
}

/*
AwaitAll is the same as Await, but for multiple futures.
*/
func AwaitAll[T any](futures ...*Future[T]) ([]T, error) {
	vals := make([]T, len(futures))
	for i, f := range futures {
		val, err := f.Await()
		if err != nil {
			return nil, err
		}
		vals[i] = val
	}
	return vals, nil
}

/*
AwaitRuntime is a runtime version of async.Await.
*/
func AwaitRuntime(f ImplementsAwaitRuntime) (any, error) {
	return f.AwaitRuntime()
}
