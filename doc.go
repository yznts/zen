//nolint:dupword
/*
Zen is a set of small utilities that you probably miss.
It's a common situation when simple things drive you crazy
like missing ternary operator,
atomic operations that take at least 3 lines,
dealing with complex loops due to missing map/filter,
or having to deal with goroutines and synchronization even for simple things.

Zen tries to solve it.
Not solves, but definitely tries.
It provides you with a number of small packages to make your work with Go easier.
Let's look at a fairly common situation where you need to filter a slice.

	data := []int{654, 234, 546, 23, 76, 87, 34, 232, 656, 767, 23, 4, 546, 56}
	newdata := []int{}
	for _, v := range data {
		if v > 50 {
			newdata = append(newdata, v)
		}
	}

It's really annoying, isn't it?
The language has no built-in capabilities for doing such operations inline.
Let's see what it would look like with our package.

	data := []int{654, 234, 546, 23, 76, 87, 34, 232, 656, 767, 23, 4, 546, 56}
	newdata := slice.Filter(data, func(v int) {
		return v > 50
	})

Let's look at another example.
Sometimes you run into situations where the structure takes a pointer to a simple type, like string.
It's understandable, sometimes we need to take nil as one of the possible states.
But if we will try to create a pointer from an inline value, we get an error.

	SomeStruct{
		Value: &"predefined", // invalid operation: cannot take address of "predefined" (untyped string constant)
	}

So you'll end up defining one more variable before creating the struct.
Now you can sleep peacefully.

	SomeStruct{
		Value: conv.Ptr("predefined"), // works!
	}

Need some kind of async/await instead of managing mutexes by hand?
Yep, sure.

	// Let's fetch uuid from httbin for this workload example.
	// Please note, example is highly simplified and not includes error checking.
	func httpbinuuid() *async.Future[string] {
		data := jsonx.Map(
			httpx.Request("GET", "https://httpbin.org/uuid").
				Do().Text(),
		)
		return data["uuid"].(string), nil
	}

	// Spawn 3 futures
	futures := []*async.Future[string]{
		httpbinuuid(),
		httpbinuuid(),
		httpbinuuid(),
	}

	// Print execution results
	for _, future := range futures {
		log.Println(future.Await()) // <uuid>, nil
	}

In addition to the above,
the library contains many interesting things that go beyond a basic overview.
First, check the library index to explore the proposed sub-packages.
Each one has its own mini-documentation, its own index, and consists of
well-documented types and functions.

Zen has been trying to be modular since v3,
so it now consists of sub-packages and does not provide anything from the root package,
except the package overview.
*/
package zen
