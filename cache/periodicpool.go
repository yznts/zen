package cache

import (
	"context"
	"errors"
	"time"
)

var ErrPeriodicPoolMissingKey = errors.New("referenced key is missing")

/*
PeriodicPool manages multiple PeriodicFunc.
Main idea of pool is having shared context between
PeriodicFunc in the pool.
*/
type PeriodicPool[T any] struct {
	Pool    map[string]PeriodicFunc[T]
	Context context.Context //nolint:containedctx
}

// New creates a new PeriodicFunc under given key.
func (p *PeriodicPool[T]) New(key string, interval time.Duration, fn PeriodicFunc[T]) {
	p.Pool[key] = NewPeriodicFunc(p.Context, interval, fn)
}

// Get allows to access specific PeriodicFunc with a given key.
func (p *PeriodicPool[T]) Get(key string) (T, error) {
	if p.Pool[key] == nil {
		// Initialize empty value
		var v T
		// Return empty value and error
		return v, ErrPeriodicPoolMissingKey
	}
	// Return periodic result
	return p.Pool[key]()
}

/*
NewPeriodicPool is a PeriodicPool builder.
Check PeriodicPool for details.

Usage:

	ctx, cancel := context.WithCancel(context.Background())
	pool := NewPeriodicPool(context.Background())
	pool.New("example", 1 * time.Minute, func() (string, error) {
		time.Sleep(1 * time.Second) Imitate some work
		return "Function cached result"
	})
	log.Println(pool.Get("example")) // Value from getter
*/
func NewPeriodicPool[T any](ctx context.Context) *PeriodicPool[T] {
	return &PeriodicPool[T]{
		Pool:    map[string]PeriodicFunc[T]{},
		Context: ctx,
	}
}
