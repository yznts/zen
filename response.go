/*
	-

	Response

	Zen provides a *http.Response wrapper with an ability to chain response processing.
	It allows to operate with wrapped response in a more convenient way.
	Check status code, dump response to stdout for debug, convert into map or decode directly into value.
	Almost everything in single line of code.

	Example:

		func main() {
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
		}
*/
package zen

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

// ResponseWrapper is a wrapper around http.Response.
// It provides a set of functions for a chained response processing.
type ResponseWrapper struct {
	*http.Response

	err error
}

// Must is a chain closer.
// Ensures that there was no errors in processing chain.
// If not, it panics.
func (r *ResponseWrapper) Must() {
	if r.err != nil {
		panic(r.err)
	}
}

// Ensure is a chain closer.
// Ensures that there was no errors in processing chain.
// If not, returns an error.
func (r *ResponseWrapper) Ensure() error {
	return r.err
}

// Debug prints the response to stdout.
// If something goes wrong during dump, chain execution will be stopped.
// Returns wrapper for chaining.
func (r *ResponseWrapper) Debug() *ResponseWrapper {
	// Check error status
	if r.err != nil {
		return r
	}
	// Dump raw response
	dump, err := httputil.DumpResponse(r.Response, true)
	// If we got an error, prevent further chain execution
	if err != nil {
		r.err = err
		return r
	}
	// Print raw response
	println(string(dump))
	// Return wrapper
	return r
}

// Success ensures that response code is between 200 and 299.
// If not, chain execution will be stopped.
// Returns wrapper for chaining.
func (r *ResponseWrapper) Success() *ResponseWrapper {
	// Check error status
	if r.err != nil {
		return r
	}
	// Check status code
	if r.StatusCode < 200 || r.StatusCode > 299 {
		// Prevent further chain execution
		r.err = fmt.Errorf("status code is not between 200 and 299: %d", r.StatusCode)
	}
	// Return wrapper
	return r
}

// Text reads response body as a text into provided target.
// If something goes wrong during reading, chain execution will be stopped.
// Returns wrapper for chaining.
func (r *ResponseWrapper) Text(target *string) *ResponseWrapper {
	// Check error status
	if r.err != nil {
		return r
	}
	// Read response body
	bts, err := ioutil.ReadAll(r.Body)
	// If we got an error, prevent further chain execution
	if err != nil {
		r.err = err
		return r
	}
	// Set value to the target
	*target = string(bts)
	// Return wrapper
	return r
}

// Decode detects response type and decodes it into target.
// If response type is not supported, or there is an error during decoding,
// chain execution will be stopped.
// Returns wrapper for chaining.
func (r *ResponseWrapper) Decode(target any) *ResponseWrapper {
	// Check error status
	if r.err != nil {
		return r
	}
	// Process response
	if r.Header.Get("Content-Type") == "application/json" { // JSON
		// If we got an error, prevent further chain execution
		if err := json.NewDecoder(r.Body).Decode(&target); err != nil {
			r.err = err
			return r
		}
	} else { // If not detected, prevent further chain execution
		r.err = fmt.Errorf("content type is not detected by library: %s", r.Header.Get("Content-Type"))
		return r
	}
	// Return wrapper
	return r
}

// Response wraps *http.Response with own wrapper,
// providing extra functions.
// See `ResponseWrapper` for details.
func Response(resp *http.Response, err ...error) *ResponseWrapper {
	if len(err) == 0 {
		err = append(err, nil)
	}
	return &ResponseWrapper{resp, err[0]}
}
