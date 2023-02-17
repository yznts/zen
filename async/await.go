package async

/*
Interface for runtime await.
*/
type implementsAwaitRuntime interface {
	AwaitRuntime() (any, error)
}

/*
Await for a future object results.
*/
func Await[T any](f *Future[T]) (T, error) {
	return f.Await()
}

/*
AwaitRuntime is a runtime version of async.Await.
*/
func AwaitRuntime(f implementsAwaitRuntime) (any, error) {
	return f.AwaitRuntime()
}
