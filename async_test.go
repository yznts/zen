package zen

import (
	"log"
	"testing"
)

func TestAwaitMultiple(t *testing.T) {
	foo := Async(func() (string, error) {
		return "foo", nil
	})

	bar, err := Await(foo)
	if err != nil {
		panic(err)
	}
	baz, err := Await(foo)
	if err != nil {
		panic(err)
	}

	log.Println(bar)
	log.Println(baz)
}
