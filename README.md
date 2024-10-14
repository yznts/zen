
<p align="center">
    <img width="200" src="logo.svg" />
</p>

<h1 align="center">zen</h1>

<p align="center">
    A set of small utilities that you probably miss
</p>

```go
import "github.com/yznts/zen/v3"
```

It's a common situation when simple things drive you crazy like missing ternary operator, atomic operations that take at least 3 lines, dealing with complex loops due to missing map/filter, or having to deal with goroutines and synchronization even for simple things.

Zen tries to solve it. Not solves, but definitely tries. It provides you with a number of small packages to make your work with Go easier. Let's look at a fairly common situation where you need to filter a slice.

```go
data := []int{654, 234, 546, 23, 76, 87, 34, 232, 656, 767, 23, 4, 546, 56}
newdata := []int{}
for _, v := range data {
        if v > 50 {
                newdata = append(newdata, v)
        }
}
```

It's really annoying, isn't it? The language has no built\-in capabilities for doing such operations inline. Let's see what it would look like with our package.

```go
data := []int{654, 234, 546, 23, 76, 87, 34, 232, 656, 767, 23, 4, 546, 56}
newdata := slice.Filter(data, func(v int) {
        return v > 50
})
```

Need to convert a slice of values to something different? Not a big deal. Just give a processing function to "slice.Map".

```go
data := []int{654, 234, 546, 23, 76, 87, 34, 232, 656, 767, 23, 4, 546, 56}
datastr := slice.Map(data, strconv.Itoa) // You'll get []string{...}
```

Let's look at another example. Sometimes you run into situations where the structure takes a pointer to a simple type, like string. It's understandable, sometimes we need to take nil as one of the possible states. But if we will try to create a pointer from an inline value, we get an error.

```go
SomeStruct{
        Value: &"predefined", // invalid operation: cannot take address of "predefined" (untyped string constant)
}
```

So you'll end up defining one more variable before creating the struct. Now you can sleep peacefully.

```go
SomeStruct{
        Value: conv.Ptr("predefined"), // works!
}
```

Let's move on to the next example, which is very similar. They all look alike, don't they? Now, we will try to implement "default value". Of course, without any additional methods or wrappers it would look something like this.

```go
value := source1 // Let's assume it's a string
if value == "" {
        value = source2
}
```

Let's make our lives a little easier with the "logic" mini\-package.

```go
value := logic.Or(source1, source2)
```

Need some kind of async/await instead of managing mutexes by hand? Yep, sure.

```go
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
```

In addition to the above, the library contains many interesting things that go beyond a basic overview. First, check the library index to explore the proposed sub\-packages. Each one has its own mini\-documentation, its own index, and consists of well\-documented types and functions.

Zen has been trying to be modular since v3, so it now consists of sub\-packages and does not provide anything from the root package, except the package overview.
