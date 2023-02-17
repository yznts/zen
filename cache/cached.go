package cache

import "time"

/*
Cached function type.
*/
type CachedFunc[T any] func() (T, error)

/*
NewCachedFunc is a function wrapper with exprire duration setting.
Consider it as a cached getter builder.
As far as NewCachedFunc doesn't support functions with arguments
(it will require much more effort and complications),
you'll need to create cached getter for each argument set.
At least, until the creation of more advanced cached builder.

Usage:

	getter := NewCachedFunc(1 * time.Minute, func() (string, error) {
		time.Sleep(1 * time.Second) Imitate some work
		return "Function cached result"
	})
	log.Println(getter()) // Takes some time
	log.Println(getter()) // Gets a value from cache
*/
func NewCachedFunc[T any](expire time.Duration, fn CachedFunc[T]) CachedFunc[T] {
	// Define variables
	var (
		err       error
		value     T
		lastfetch time.Time
	)
	// Return getter
	return func() (T, error) {
		// Update value/err cache if expired
		if time.Since(lastfetch) > expire {
			value, err = fn()
			lastfetch = time.Now()
		}
		// Return from cache
		return value, err
	}
}
