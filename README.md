
# Zen

Zen is an utility-first package that provides a set of commonly used functions, helpers and extensions.
Most of the functions are adapted to be used with `html/template`.

## Motivation

We had too many cases of copy-pasting the same code between projects over and over again.
So I decided to combine it all in one place.

## Documentation

Code is well-documented, so it's pretty comfortable to use documentation provided by [pkg.go.dev](https://pkg.go.dev/github.com/kyoto-framework/zen/v2)

## Examples

Provided examples doesn't cover all the features of the library. Please refer to the documentation for details.

```go
package main

import (
    "time"

    "github.com/kyoto-framework/zen/v2"
)

// Example of async function
func foo() *zen.Future[string] {
    return zen.Async(func() (string, error) {
        // Imitate work
        time.Sleep(time.Second)
        // Return result
        return "bar", nil
    })
}

func main() {
    // Sample slice with zen.Range
    slice := zen.Range(1, 5) // []int{1, 2, 3, 4, 5}

    // Aggregatives
    zen.Min(slice) // 1
    zen.Max(slice) // 5
    zen.Avg(slice) // 3

    // Range operations
    zen.Filter(slice, func(v int) bool { return v < 3 }) // []int{1, 2}
    zen.Map(slice, func(v int) int { return v * 2 }) // []int{2, 4, 6, 8, 10}
    zen.In(1, slice) // true
    zen.Pop(slice, 1) // ([]int{1, 3, 4, 5}, 2)
    zen.Insert(slice, 1, 2) // []int{1, 2, 2, 3, 4, 5}
    zen.Last(slice) // 5
    zen.Any(slice, func(v int) bool { return v == 2 }) // true
    zen.All(slice, func(v int) bool { return v < 6 }) // true

    // Inline transformations
    zen.Ptr(1) // *int{1}  Inline pointer
    zen.Bool(3) // bool{true}
    zen.Int("5") // int{5}
    zen.Float64("6.5") // float64{6.5}
    zen.String(7) // string{"7"}

    // Map composition (useful for templates)
    zen.Compose("foo", 1, "bar", "2") // map[any]any{"foo": 1, "bar": "2"}

    // Must for opinionated panic handling
    zen.Must(strconv.Atoi("1")) // int{1}, without error

    // Await example
    futureFoo := foo() // It's not blocking
    println("foo call is not blocked")
    // Wait for result
    result, _ := zen.Await(futureFoo)
    println(result)
}
```
