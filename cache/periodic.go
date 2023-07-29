package cache

import (
	"context"
	"time"
)

/*
PeriodicFunc is a function wrapper that periodically executes function and caches result.
As far as CachedFunc doesn't support functions with arguments
(it will require much more effort and complications),
you'll need to create periodic getter for each argument set.
At least, until the creation of more advanced periodic builder.
*/
type PeriodicFunc[T any] func() (T, error)

/*
NewPeriodicFunc is a PeriodicFunc builder.
Check PeriodicFunc for details.

Usage:

	getter := NewPeriodicFunc(context.Background(), 1 * time.Minute, func() (string, error) {
		time.Sleep(1 * time.Second) Imitate some work
		return "Function cached result"
	})
	time.Sleep(1 * time.Second)
	log.Println(getter()) // Get a value from cache
	log.Println(getter()) // Get a value from cache
*/
func NewPeriodicFunc[T any](ctx context.Context, interval time.Duration, fn PeriodicFunc[T]) PeriodicFunc[T] {
	// Define variables
	var (
		err   error
		value T
	)
	// Schedule periodic update gorutine
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				// Update cache
				value, err = fn()
				// Sleep for interval
				time.Sleep(interval)
			}
		}
	}()
	// Return getter
	return func() (T, error) {
		return value, err
	}
}
