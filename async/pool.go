package async

import "sync"

/*
Pool creates i/o channels and spins up a pool of workers for it.
To stop the pool, you have to close input channel.
Output channel will be closed automatically on workers completion.

Please note, output order is not guaranteed!
Use own wrapper if you need to identify output.

Usage:

	// Task holds processing data, like id, input, and result.
	type Task struct {
		ID int
		Value int
		Result int
	}

	in, out := async.Pool(10, func(v *Task) *Task {
		return Workload(v)
	})

	go func() {
		for i := 0; i < 1000; i++ {
			in <- &Task{ID: i, Value: i}
		}
		close(in)
	}()

	for v := range out {
		fmt.Println(v.ID, v.Result)
	}
*/
func Pool[T1 any, T2 any](num int, worker func(v T1) T2) (chan T1, chan T2) {
	in := make(chan T1)
	out := make(chan T2)

	wg := sync.WaitGroup{}

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range in {
				out <- worker(v)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return in, out
}
