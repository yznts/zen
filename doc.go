/*
Package zen is an utility-first package that provides a set of commonly used functions, helpers and extensions.
Most of the functions are adapted to be used with `html/template`.

We had too many cases of copy-pasting the same code between projects over and over again.
So we decided to combine it all in one place.

Aggregative

Library provides basic aggregative functions.

	zen.Min(1, 2, 3) // 1
	zen.Max(1, 2, 3) // 3
	zen.Avg(1, 2, 3) // 2

Arithmetic

Library provides simple arithmetic functions for sets of values: Sum, Sub, Mul, Div.
Main point is to provide runtime template functions (which are missing in the built-in `html/template`).
See Funcmap section for details.

Async

Library provides a way to define and run asynchronous functions.
It's based on Go's standard goroutines and channels.
Future object holds value channel and error.
It's used as an awaitable object.
As far as Go is not provides any async/await syntax,
your function must to return a Future,
provided by Async function.

	// Example asynchronous function
	func Foo() *zen.Future[string] {
		return zen.Async(func() (string, error) {
			return "Bar", nil
		})
	}

	func main() {
		// Non-blocking calls
		fbar1 := Foo()
		fbar2 := Foo()
		// Await for results (errors are passed to simplify example)
		bar1, _ := zen.Await(fbar1)
		bar2, _ := zen.Await(fbar2)
	}

Atomic

Zen provides a generic atomic wrapper, based on RWMutex.
Usually, actions are performed with Get and Set methods.
For complex cases (like simultaneous Get and Set),
Atomic provides a Context method which allows to pass a function
and lock a mutex for the duration of the function.

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

Cast

Library provides a comfortable way to work with slices casting ([]any).

	// Define a sample untyped slice
	values := []any{1, 2, 3}
	// Casting example
	castedValues := zen.CastSlice[int](values) // []int{1, 2, 3}

Errors

Library provides a simple helper functions to handle error cases,
like Must or Ignore.

	// Panic on error
	value := zen.Must(strconv.Atoi("abc"))
	// Empty value on error
	value := zen.Ignore(strconv.Atoi("abc")) // 0

Funcmap

Library provides an utilities funcmap to be attached to template rendering.
See FuncMap variable for details.

Strings

Zen provides some useful non-standard string formatting functions.

	zen.Replace("Hello, 42 World!", "[0-9]+", "<number>") // Hello, <number> World!
	zen.FormatNumber(12345.456, 0, "$", "") // "$12,345"
	zen.FormatNumberP0(12345.456) // "12,345"
	zen.FormatNumberP1(12345.456) // "12,345.4"
	zen.FormatNumberNumeral(12345.456, 0) // "12k"
	zen.FormatNumberNumeralP0(12345.456) // "12k"
	zen.FormatNumberNumeralP1(12345.456) // "12.3k"

Logical

Logical expressions from another languages, but missing in Go.
Unfortunately, it's not syntax-level feature.
That's why you can't use conditional code execution here.

	// Go is not supporting "or" for values, like (0 || 1)
	zen.Or("", "asd") // string{"asd"}
	// Go doesn't have "ternary" operator, like (true ? "asd" : "qwe")
	zen.Tr(false, "asd", "qwe") // string{"qwe"}
	// As far as it's not a syntax feature, code is executing in any case.
	// These lines will panic, because values[0] is executing even if condition is matching first case.
	values := []int{}
	zen.Tr(len(values) == 0, 123, values[0])

Range

Some generic functions for basic slice operations.

	// Creating an integers slice with a Range function
	var slice = zen.Range(1, 5) // []int{1, 2, 3, 4, 5}

	// Filtering
	zen.Filter(slice, func(v int) bool { return v < 3 }) // []int{1, 2}

	// Creating a new slice, based on existing one with a Map function
	zen.Map(slice, func(v int) int { return v * 2 }) // []int{2, 4, 6, 8, 10}

	// Checking if an element is in the slice
	zen.In(1, slice) // true

	// Pop an element at the given index from the slice (returns a new slice and the value)
	zen.Pop(slice, 1) // ([]int{1, 3, 4, 5}, 2)

	// Insert an element at the given index in the slice (returns a new slice)
	zen.Insert(slice, 1, 2) // []int{1, 2, 2, 3, 4, 5}

	// Get the last element from the slice
	zen.Last(slice) // 5

	// Check if any element in the slice matches the given function
	zen.Any(slice, func(v int) bool { return v == 2 }) // true

	// Check if all elements in the slice match the given function
	zen.All(slice, func(v int) bool { return v < 6 }) // true

net/http

Library provides a set of wrappers and builder in addition to default net/http package.

RequestBuilder is a builder with an ability to chain request definition and execution.
You can initialize RequestBuilder with zen.Request function.
It allows to set query, headers, body, almost everything in single line of code.
Supports both setting values with structs/maps and native values (like url.Values for query).
As far as RequestBuilder have an option to return *ResponseWrapper,
you can chain both request building and response processing into single line.

	// This request building chain will:
	// - Parse initial values (method and url)
	// - Set a query value "qval" for a key "qkey"
	// - Set a header "hkey: hval"
	// - Set an url encoded body and a header "Content-Type: application/x-www-form-urlencoded"
	// - Build a resulting request and return
	request := zen.Request("POST", "https://httpbin.org/post").
		Query("qkey", "qval").
		Header("hkey", "hval").
		Form(map[string]string{"fkey": "fval"}).
		Build()

	// Request wrapper also have an inline execution option.
	// This gives an ability to process response inline.
	data := map[string]any{}
	err := zen.Request("GET", "https://httpbin.org/get").
		Do().
		Success().
		Decode(&data).
		Ensure()
	if err != nil {
		fmt.Println("Something went wrong: %v", err)
	}

ResponseWrapper is a wrapper with an ability to chain response processing.
You can initialize ResponseWrapper with a zen.Response function.
It allows to operate with wrapped response in a more convenient way.
Check status code, dump response to stdout for debug, convert body into map or decode directly into value.
Almost everything in single line of code.

	// Data holder
	data := map[string]any{}
	// This execution chain will:
	// - Wrap response and error with zen.ResponseWrapper
	// - Print raw response dump into stdout
	// - Ensure status is between 200-299
	// - Decode data into provided data holder (struct or map)
	// - Panics, if something went wrong during processing
	zen.Response(http.Get("https://httpbin.org/get")).
		Debug().Success().
		Decode(&data).
		Must()

Transform

Library provides a number of functions that can be used to transform data into different types and forms.
Most of these functions are working with base data types.

	// Common data types transformations
	numptr := zen.Ptr(1) // *int{1}  Inline pointer
	boolval := zen.Bool(3) // bool{true}
	intval := zen.Int("5") // int{5}
	floatval := zen.Float64("6.5") // float64{6.5}
	strval := zen.String(7) // string{"7"}

	// Map composition (useful for templates)
	resmap := zen.Compose("foo", 1, "bar", "2") // map[any]any{"foo": 1, "bar": "2"}

	// JSON
	resjson := zen.JSON(resmap) // string{`{"foo": 1, "bar": "2"}`}

	// Base64
	resbase64 := zen.B64Enc(resjson) // string{`eyJmb28iOiAxLCAiYmFyIjogIjIifQ==`}
	resbase64dec := string(zen.B64Dec(resbase64)) // string{`{"foo": 1, "bar": "2"}`}

*/
package zen