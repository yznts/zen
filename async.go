package zen

type Future[T any] struct {
	value chan T
	err   error
}

func Await[T any](f Future[T]) (T, error) {
	// Wait for value
	v := <-f.value
	// Return
	return v, f.err
}

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
	}()
	// Return future
	return future
}
