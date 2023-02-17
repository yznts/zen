package async

/*
New runs a function in a goroutine and returns Future object for it.
*/
func New[T any](fn func() (T, error)) *Future[T] {
	// Create future
	future := Future[T]{value: make(chan T)}
	// Run thread
	go func() {
		// Run function
		value, err := fn()
		// Set error
		future.err = err
		// Set value
		future.value <- value
		// Close value channel
		close(future.value)
	}()
	// Return future
	return &future
}
