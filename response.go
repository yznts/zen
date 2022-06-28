/*
	-

	Response

	Zen provides a *http.Response wrapper with a few useful extra methods.
	It allows to operate with wrapped response in a more convenient way.
	Check status code, dump response to stdout for debug, convert into map or decode directly into value.
	Almost everything in single line of code.

	Example:

		func main() {
			// Make a request
			resp, err := http.Get("https://example.com/api/json")
			if err != nil {
				panic(err)
			}
			// Dump, check, convert to map
			data := zen.Response(resp).Debug().Must().Map()
		}
*/
package zen

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

// ResponseWrapper is a wrapper for http.Response.
// It provides a few useful extra methods.
type ResponseWrapper struct {
	*http.Response
}

// Must ensures that response code is between 200 and 299.
// If not, it panics.
// Returns itself for chaining.
func (r *ResponseWrapper) Must() *ResponseWrapper {
	if r.StatusCode < 200 || r.StatusCode > 299 {
		panic("unexpected response code")
	}
	return r
}

// Debug prints the response to stdout.
// Returns itself for chaining.
func (r *ResponseWrapper) Debug() *ResponseWrapper {
	dump, err := httputil.DumpResponse(r.Response, true)
	if err != nil {
		panic(err)
	}
	println(string(dump))
	return r
}

// Map detects response type and decodes it to map[string]any{}.
// If the response type is not supported, it panics.
func (r *ResponseWrapper) Map() map[string]any {
	data := map[string]any{}
	if r.Header.Get("Content-Type") == "application/json" {
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			panic(err)
		}
	} else {
		panic("unsupported response type")
	}
	return data
}

// Text reads response body as a text.
// If something wrong, it panics.
func (r *ResponseWrapper) Text() string {
	bts, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	return string(bts)
}

// Decode detects response type and decodes it to target.
// If the response type is not supported, it panics.
func (r *ResponseWrapper) Decode(target any) {
	if r.Header.Get("Content-Type") == "application/json" {
		if err := json.NewDecoder(r.Body).Decode(&target); err != nil {
			panic(err)
		}
	} else {
		panic("unsupported response type")
	}
}

// Response wraps *http.Response with own wrapper,
// providing extra methods.
func Response(resp *http.Response) *ResponseWrapper {
	return &ResponseWrapper{resp}
}
