package cache

import "time"

/*
CachedFunc is a cached function wrapper with exprire duration setting.
Consider it as a cached getter.
As far as CachedFunc doesn't support functions with arguments
(it will require much more effort and complications),
you'll need to create cached getter for each argument set.
At least, until the creation of more advanced cached builder.
*/
type CachedFunc[T any] func() (T, error)

/*
NewCachedFunc is a CachedFunc builder.
Check CachedFunc for details.

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
