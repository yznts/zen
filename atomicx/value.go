package atomicx

import (
	"sync"
)

/*
Value is an atomic value holder,
which have internal rwmutex to avoid concurrent writes.

Usage:

	// Actually, we have a NewValue() constructor,
	// but let's show how to use a Value without it.
	val := &Value[int]{}
	val.Set(1)
	if val.Get() == 1 {
		log.Println("value is 1")
	}
*/
type Value[T any] struct {
	value T
	lock  sync.RWMutex
}

/*
Get locks value mutex and returns current value.
*/
func (a *Value[T]) Get() T {
	a.lock.RLock()
	defer a.lock.RUnlock()
	return a.value
}

/*
Set locks value mutex and sets current value.
*/
func (a *Value[T]) Set(value T) {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value = value
}

/*
Context locks value mutex and executes provided function
with releasing lock in the end.
*/
func (a *Value[T]) Context(c func(value T, set func(value T))) {
	a.lock.Lock()
	defer a.lock.Unlock()
	setter := func(value T) {
		a.value = value
	}
	c(a.value, setter)
}

/*
NewValue is a constructor of Value,
which accepts default value.
*/
func NewValue[T any](defaultval T) *Value[T] {
	return &Value[T]{
		value: defaultval,
	}
}
