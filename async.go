/*
	-

	Async

	Zen provides a way to define and run asynchronous functions.
	It's based on Go's standard goroutines and channels.
	Future object holds value channel and error.
	It's used as an awaitable object.
	As far as Go is not provides an async/await syntax,
	your function must to return a Future, provided by Async function.

	Example:

		func Foo() zen.Future[string] {
			return zen.Async(func() (string, error) {
				return "Bar", nil
			})
		}

		func main() {
			// Non-blocking calls
			fbar1 := Foo()
			fbar2 := Foo()
			// Await for results (errors are passed to simplify example)
			bar1, _ := zen.Await(fbar1)
			bar2, _ := zen.Await(fbar2)
		}

*/
package zen

// Future is an awaitable object.
// Behavior is similar to JS Promise.
type Future[T any] struct {
	value chan T
	err   error
}

// Await for a future object.
func Await[T any](f Future[T]) (T, error) {
	// Wait for value
	v := <-f.value
	// Return
	return v, f.err
}

// Async runs a function in a goroutine and returns Future object for it.
func Async[T any](f func() (T, error)) Future[T] {
	// Create future
	future := Future[T]{value: make(chan T)}
	// Run thread
	go func() {
		// Run function
		value, err := f()
		// Set error
		future.err = err
		// Set value
		future.value <- value
		// Close value channel
		close(future.value)
	}()
	// Return future
	return future
}
