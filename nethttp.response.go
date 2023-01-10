package zen

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"
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

// Error is a chain closer.
// Ensures that there was no errors in processing chain.
// If not, error is not nil.
func (r *ResponseWrapper) Error() error {
	return r.err
}

// Ensure is an alias for .Error()
func (r *ResponseWrapper) Ensure() error {
	return r.Error()
}

// Debug prints the response to stdout.
// If something goes wrong during dump, chain execution will be stopped.
// Returns wrapper for chaining.
func (r *ResponseWrapper) Debug() *ResponseWrapper {
	// Check error status
	if r.err != nil {
		return r
	}
	// Dump url
	println(r.Request.URL.String())
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

// Text reads response body as a text.
func (r *ResponseWrapper) Text() string {
	return string(Ignore(ioutil.ReadAll(r.Body)))
}

// Unmarshal detects response type and decodes it into target.
// If response type is not supported, or there is an error during decoding,
// chain execution will be stopped.
// Returns wrapper for chaining.
func (r *ResponseWrapper) Unmarshal(target any) *ResponseWrapper {
	// Check error status
	if r.err != nil {
		return r
	}
	// Process response
	switch strings.Split(r.Header.Get("Content-Type"), ";")[0] {
	case "application/json":
		if err := json.NewDecoder(r.Body).Decode(&target); err != nil {
			r.err = err
			return r
		}
	case "text/plain":
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			r.err = err
			return r
		}
		*target.(*string) = string(data)
	default:
		panic("content type is not supported")
	}
	// Return wrapper
	return r
}

// Response wraps *http.Response with own wrapper,
// providing extra functions.
// See ResponseWrapper for details.
func Response(resp *http.Response, err ...error) *ResponseWrapper {
	if len(err) == 0 {
		err = append(err, nil)
	}
	return &ResponseWrapper{resp, err[0]}
}
