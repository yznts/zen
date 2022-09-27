/*
-

# Atomic

Zen provides a generic atomic wrapper, based on RWMutex.
Usually, actions are performed with Get and Set methods.
For complex cases (like simultaneous Get and Set),
Atomic provides a Context method which allows to pass a function
and lock a mutex for the duration of the function.

Example:

	func main() {
		// Initialize atomic value
		value := Atomic[int]{}
		value.Set(1)

		// Get and check value
		if value.Get() == 1 {
			println("It's OK")
		}

		// Pass a context function, which will obtain a lock
		value.Context(func(value int, set func(value int)) {
			if value == 1 {
				set(2)
			}
		})
	}
*/
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
