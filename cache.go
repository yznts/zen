package zen

import "time"

type cachedFunc[T any] func() (T, error)

// Cached is a function wrapper with exprire duration setting.
// Consider it as a cached getter builder.
// As far as Cached doesn't support functions with arguments
// (it will require much more effort and complications),
// you'll need to create cached getter for each argument set.
// At least, until the creation of more advanced cached builder.
//
// Usage:
//
//	getter := Cached(1 * time.Minute, func() (string, error) {
//		time.Sleep(1 * time.Second) // Imitate some work
//		return "Function cached result"
//	})
//	log.Println(getter())
//	log.Println(getter())
func Cached[T any](expire time.Duration, fn cachedFunc[T]) cachedFunc[T] {
	var (
		err       error
		value     T
		lastfetch time.Time
	)
	return func() (T, error) {
		if time.Since(lastfetch) > expire {
			value, err = fn()
			lastfetch = time.Now()
		}
		return value, err
	}
}
