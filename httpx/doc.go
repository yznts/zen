/*
httpx - a package with http extensions and wrappers.
Gives ability for advanced path/query processing,
request/response actions, etc.

# Path

For path processing, you can use Path type.
It provides a few useful extra methods to operate with path.
Might be helpful in case of using Go's built-in mux router,
which doesn't provide path parameters extraction.

Usage:

	// httpx.Path is a string sub type, so we are wrapping a string path with it.
	path := httpx.Path("/foo/bar/baz")
	// Tokens returns path tokens, splitted by slash.
	path.Tokens() // []string{"foo", "bar", "baz"}
	// Get returns path token at the given index.
	path.Get(1) // string{"bar"}
	// GetAfter returns path token located after provided token.
	path.GetAfter("bar") // string{"baz"}
	// GetAfterWithIndex returns path token and it's index located after provided token.
	path.GetAfterWithIndex("bar") // string{"baz"}, 2
	// GetBefore returns path token located before provided token.
	path.GetBefore("baz") // string{"bar"}
	// GetBeforeWithIndex returns path token and it's index located before provided token.
	path.GetBeforeWithIndex("baz") // string{"bar"}, 1

# Query

For query processing, you can use Query type.
It provides a few useful extra methods to operate with query.

Usage:

	// httpx.Query is a url.Values sub type, so we are wrapping a url.Values with it.
	query := httpx.Query(r.URL.Query())

	// With a Query wrapper we can use Unmarshal method to unmarshal query values into a struct.
	type QueryParams struct {
		Foo string `query:"foo"`
		Bar string `query:"bar"`
	}
	var params QueryParams
	err := query.Unmarshal(&params)

# Request / Response

This package comes with a chainable request builder and response wrapper.

Usage:

	// You can use httpx.Request to build a new request
	// and chain request building methods.
	req := httpx.Request("GET", "http://example.com").
		Query("foo", "bar"). // or QueryMap, QueryMapFmt, QueryValues, QueryStruct
		Header("X-Foo", "Bar"). // or HeaderMap, HeaderMapFmt, HeaderValues
		BodyJson(map[string]any{"foo": "bar"}) // or Body, BodyText, BodyForm

	// We have multiple ways to finalize request building.
	res := req.Build() // Build resulting *http.Request
	res := req.Do() // Execute request and get *ResponseWrapper
	res := req.Async() // Execute request asynchronously and get *async.Future[*ResponseWrapper]

	// You can use httpx.Response to wrap existing *http.Response with error.
	res := httpx.Response(http.DefaultClient.Get("http://example.com"))

	// ResponseWrapper provides a few useful methods to operate with response.
	// It's chainable, so you can use it in a fluent way.
	// If error occurs on any stage,
	// following methods will not be executed.
	// Error will be stored in ResponseWrapper
	// and you can get it with an Error method.
	res.
		Debug(). // Print response debug info.
		Success(). // Ensure that response status code is in 2xx range.
		Unmarshal(&data). // Unmarshal response body into variable.
		Must() // Ensure that everything went fine, otherwise panic.

	err := res.Error() // Get processing error. If someting went wrong on any chain stage, it will be here.
	txt := res.Text() // Get response body as a string.
*/
package httpx
