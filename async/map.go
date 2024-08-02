package async

import "sync"

/*
Map returns a new slice with the results of applying the given function to each element in the given slice.
Asynchronous version of slice.Map. Please note, it's not always faster!
Goroutines allocation and synchronization have own cost.

Usage:

	// Let's assume we have some workload in Workload function, which returns an integer.

	results := async.Map(slice.Range(1, 1000), func(v int) int {
		return Workload(v)
	})
*/
func Map[T1 any, T2 any](slice []T1, fn func(v T1) T2) []T2 {
	newslice := make([]T2, len(slice))

	wg := sync.WaitGroup{}
	wg.Add(len(slice))

	for i, v := range slice {
		go func(i int, v T1, wg *sync.WaitGroup) {
			defer wg.Done()
			newslice[i] = fn(v)
		}(i, v, &wg)
	}

	wg.Wait()

	return newslice
}
