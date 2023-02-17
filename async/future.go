package async

import "encoding/json"

/*
Future is an execution result of an asynchronous function
that returns immediately, without locking execution thread.
To lock execution and wait for result, use .Await() method
or async.Await() function. As an alternative you can use
a syntax similar to JavaScript Promise, using .Then()
and .Catch() methods.

Usage:

	// Let's assume we have a future object in "ftr" variable.
	// We can lock execution and wait for a result with .Await()
	res, err := ftr.Await()
	// Or, we can use async.Await()
	res, err := async.Await(ftr)
	// Or, we can avoid locking execution and provide then/catch
	// functions to handle execution results.
	ftr.Then(func(val string) {
		println(val)
	}).Catch(func(err error) {
		println(err.Error())
	})
*/
type Future[T any] struct {
	value chan T
	cache *T
	err   error
}

/*
Await for a future object results.

Usage:

	// Let's assume we have a future object in "ftr" variable.
	res, err := ftr.Await()
*/
func (c *Future[T]) Await() (T, error) {
	// Return from cache, if exists
	if c.cache != nil {
		return *c.cache, c.err
	}
	// Wait for value
	v := <-c.value
	// Save to cache
	c.cache = &v
	// Return
	return v, c.err
}

/*
AwaitRuntime is a runtime version of .Await()

Usage:

	// Let's assume we have a future object in "ftr" variable.
	// Result will be stored as "any" type, so you'll need to cast it.
	res, err := ftr.AwaitRuntime()
*/
func (c *Future[T]) AwaitRuntime() (any, error) {
	return c.Await()
}

/*
Then accepts a function, that will be executed on
future work completion.

Usage:

	// Let's assume we have a future object of string in "ftr" variable.
	ftr.Then(func(v string) {
		println(v)
	})
*/
func (c *Future[T]) Then(f func(T)) *Future[T] {
	// Await first
	c.Await() //nolint:errcheck
	// If no error, call provided function
	if c.err == nil {
		f(*c.cache)
	}
	// Self-return
	return c
}

/*
Catch accepts a function, that will be executed on
future execution error.

Usage:

	// Let's assume we have a future object of string in "ftr" variable.
	ftr.Catch(func(err error) {
		println(err.Error())
	})
*/
func (c *Future[T]) Catch(f func(error)) *Future[T] {
	// Await first
	c.Await() //nolint:errcheck
	// If error, call provided function
	if c.err != nil {
		f(c.err)
	}
	// Self-return
	return c
}

/*
MarshalJSON implements future marshalling.
*/
func (f *Future[T]) MarshalJSON() ([]byte, error) {
	val, err := Await(f)
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(val)
}

/*
UnmarshalJSON implements future unmarshalling.
*/
func (c *Future[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &c.cache)
}
