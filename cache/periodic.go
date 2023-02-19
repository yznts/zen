package cache

import (
	"context"
	"time"
)

type PeriodicFunc[T any] func() (T, error)

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
