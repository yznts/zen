package cache_test

import (
	"context"
	"testing"
	"time"

	"github.com/kyoto-framework/zen/v3/cache"
	"github.com/kyoto-framework/zen/v3/errorsx"
)

func TestPeriodicFunc(t *testing.T) {
	t.Parallel()

	// Define context with cancel
	ctx, cancel := context.WithCancel(context.Background())

	// Define periodic getter
	var getter cache.PeriodicFunc[int]
	{
		counter := 0
		getter = cache.NewPeriodicFunc(ctx, 10*time.Millisecond, func() (int, error) {
			counter++
			// Return current counter
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

	// Ensure cache periodically updates
	{
		time.Sleep(10 * time.Millisecond)
		if errorsx.Must(getter()) != 2 {
			t.Error("Cache is not updating periodically")
		}
	}

	// Ensure context cancelation works as expected
	{
		value := errorsx.Must(getter())
		cancel()
		time.Sleep(20 * time.Millisecond)
		if errorsx.Must(getter()) != value {
			t.Error("Cache is still updating periodically after context cancelation")
		}
	}
}
