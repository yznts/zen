package zen

import "encoding/json"

// Future is an awaitable object.
// Behavior is similar to JS Promise.
type Future[T any] struct {
	value chan T
	cache *T
	err   error
}

// MarshalJSON implements future marshalling.
func (f *Future[T]) MarshalJSON() ([]byte, error) {
	val, err := Await(f)
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(val)
}

// MarshalJSON implements future unmarshalling.
func (c *Future[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &c.cache)
}

// Await for a future object.
func Await[T any](f *Future[T]) (T, error) {
	// Return from cache, if exists
	if f.cache != nil {
		return *f.cache, f.err
	}
	// Wait for value
	v := <-f.value
	// Save to cache
	f.cache = &v
	// Return
	return v, f.err
}

// Async runs a function in a goroutine and returns Future object for it.
func Async[T any](f func() (T, error)) *Future[T] {
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
	return &future
}
