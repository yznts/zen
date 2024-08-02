package errx

// Result is a generic type that wraps a value and an error.
// It also provides a set of methods to resolve/unwrap underlying value
// with a specific value/error handling.
type Result[T any] struct {
	value T
	error error
}

// Simple value getters

// Error getter
func (r Result[T]) Error() error {
	return r.error
}

// Value getter
func (r Result[T]) Value() T {
	return r.value
}

// Advanced value getters

// ValueOr returns the value if no error is present,
// otherwise returns the provided default value
func (r Result[T]) ValueOr(def T) T {
	if r.error != nil {
		return def
	}

	return r.value
}

// Must returns the value if no error is present,
// otherwise panics with the error
func (r Result[T]) Must() T {
	return Must(r.value, r.error)
}

// Case handlers

// Then executes the provided function if no error is present
func (r Result[T]) Then(fn func(T)) Result[T] {
	if r.error == nil {
		fn(r.value)
	}

	return r
}

// Catch executes the provided function if an error is present
func (r Result[T]) Catch(fn func(error)) Result[T] {
	if r.error != nil {
		fn(r.error)
	}

	return r
}

// Wrap wraps a value and an error into a Result
func Wrap[T any](val T, err error) Result[T] {
	return Result[T]{value: val, error: err}
}
