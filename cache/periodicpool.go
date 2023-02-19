package cache

import (
	"context"
	"errors"
	"time"
)

var ErrPeriodicPoolMissingKey = errors.New("referenced key is missing")

type PeriodicPool[T any] struct {
	Pool    map[string]PeriodicFunc[T]
	Context context.Context //nolint:containedctx
}

func (p *PeriodicPool[T]) New(key string, interval time.Duration, fn PeriodicFunc[T]) {
	p.Pool[key] = NewPeriodicFunc(p.Context, interval, fn)
}

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

func NewPeriodicPool[T any](ctx context.Context) *PeriodicPool[T] {
	return &PeriodicPool[T]{
		Pool:    map[string]PeriodicFunc[T]{},
		Context: ctx,
	}
}
