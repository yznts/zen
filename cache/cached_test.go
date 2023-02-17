package cache_test

import (
	"testing"
	"time"

	"github.com/kyoto-framework/zen/v3/cache"
	"github.com/kyoto-framework/zen/v3/errorsx"
)

func TestCachedFunc(t *testing.T) {
	t.Parallel()

	// Define cached getter
	var getter cache.CachedFunc[int]
	{
		counter := 0
		getter = cache.NewCachedFunc(10*time.Millisecond, func() (int, error) {
			counter++
			return counter, nil
		})
	}

	// Wait for a first execution
	time.Sleep(1 * time.Millisecond)

	// Ensure we're getting first result
	{
		if errorsx.Must(getter()) != 1 {
			t.Error("First result is not right")
		}
	}

	// Ensure we're still getting first result on second run
	{
		if errorsx.Must(getter()) != 1 {
			t.Error("Second run is not right, perhaps got non-cache results")
		}
	}

	// Ensure cache expire is working
	{
		time.Sleep(10 * time.Millisecond)
		if errorsx.Must(getter()) != 2 {
			t.Error("Cache expire is not working as expected")
		}
	}
}
