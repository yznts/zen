/*
errx - a package that provides a set of helpers to work with errors in Go.

# Must / Ignore

Must is a helper that wraps a call to a function returning value and error,
and panics if the error is non-nil.

Usage:

	// We can use Must function to panic if an error is not nil
	val := errx.Must(strconv.Atoi("42"))

Ignore is a helper that wraps a call to a function returning error,
and does nothing if the error is nil.

Usage:

	// We can use Ignore function to ignore an error
	val := errx.Ignore(strconv.Atoi("42"))

# Result[T]

This generic type tries to mimic the Result type from other languages.
It wraps a value/error pair,
and provides a set of methods to resolve/unwrap underlying value with a specific value/error handling.

Usage:

	// We can use Wrap function to wrap a value/error pair into Result[T] type
	res := errx.Wrap(strconv.Atoi("42"))

	// Getter methods to get value or error
	val, err := res.Value(), res.Error()

	// Getter for getting value or default value if error is present
	val := res.ValueOr(-1)

	// Getter for getting value or panicking if error is present
	val := res.Must()

	// Call handlers to execute functions based on error presence
	res.Then(func(val int) {
		fmt.Println(val)
	}).Catch(func(err error) {
		fmt.Println("Error:", err)
	})

	// The main point of handlers is using shared functions to handle different cases
	res.Then(processValue).Catch(logError)
	// or
	val := res.Catch(logError).ValueOr(-1)
*/
package errx
