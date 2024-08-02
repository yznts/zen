/*
async - a package that provides tools for parallel execution using built-in goroutines.
Please note, because of the way goroutines work, execution is parallel, not concurrent!
On top of existing async/await functionality,
package also provides high-level functions like Map, Pool, etc.

# Async / Await

The most basic functionality of the package is async/await pattern.
It allows you to run functions in parallel and wait for their completion.

Please note, that async functions are not concurrent, they are parallel.
It's important because you still have to be careful with shared resources.

	// Example of async function.
	func mult(n int) *async.Future[int] {
		return async.New(func() (int, error) {
			time.Sleep(time.Second)
			return n * 2, nil
		})
	}

	// We will use this future to demonstrate resolving.
	ftr := mult(2)

	// We have multiple ways to handle future resolving.
	// One of them is to use Then/Catch methods.
	// Then is called on execution completion, Catch is called when future has an error.
	ftr.Then(func(val int) {
		fmt.Println("Resolved:", val)
	}).Catch(func(err error) {
		fmt.Println("Error:", err)
	})

	// Another way is to use Await method.
	// It will block until future is resolved.
	val, err := ftr.Await()

	// We also have an async.Await function if you prefer functional style.
	val, err := async.Await(ftr)

# Map / Filter / Pool

The package provides high-level functions to work with collections,
such as Map, Filter, and Pool.

Behavior of Map and Filter is similar to their counterparts in slice package.
The difference is that they run functions in parallel.
Please note, performance is usually lower than using slice package!
Use them only if you know that cpu/io cost of a logic is higher than goroutines spin up cost.

	// Example of Map function.
	results := async.Map(slice.Range(1, 1000), func(v int) int {
		return Workload(v)
	})

	// Example of Filter function.
	results := async.Filter(slice.Range(1, 1000), func(v int) bool {
		return v < 500
	})

In addition to Map and Filter, the package provides Pool function.
It creates a pool of workers and channels for input and output.

	// Task holds processing data, like id, input, and result.
	type Task struct {
		ID int
		Value int
		Result int
	}

	// Spin up a pool of 10 workers.
	in, out := async.Pool(10, func(v *Task) *Task {
		return Workload(v)
	})

	// Send 1000 tasks to the pool.
	go func() {
		for i := 0; i < 1000; i++ {
			in <- &Task{ID: i, Value: i}
		}
		close(in)
	}()

	// Read results from the pool.
	// Out channel will be closed automatically on input channel close
	// and workers completion.
	for v := range out {
		fmt.Println(v.ID, v.Result)
	}
*/
package async
