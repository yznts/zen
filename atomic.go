package zen

import "sync"

type Atomic[T any] struct {
	value T
	lock  sync.RWMutex
}

func (a *Atomic[T]) Get() T {
	a.lock.RLock()
	defer a.lock.RUnlock()
	return a.value
}

func (a *Atomic[T]) Set(value T) {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value = value
}

func (a *Atomic[T]) Context(c func(value T, set func(value T))) {
	a.lock.Lock()
	defer a.lock.Unlock()
	setter := func(value T) {
		a.value = value
	}
	c(a.value, setter)
}
