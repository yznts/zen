package zen

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"
)

// ResponseWrapper is a wrapper for http.Response.
// It provides a few useful extra methods.
type ResponseWrapper struct {
	Response *http.Response
}

// Must ensures that response code is between 200 and 299.
// If not, it panics.
// Returns itself for chaining.
func (r *ResponseWrapper) Must() *ResponseWrapper {
	if r.Response.StatusCode < 200 || r.Response.StatusCode > 299 {
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
	if r.Response.Header.Get("Content-Type") == "application/json" {
		if err := json.NewDecoder(r.Response.Body).Decode(&data); err != nil {
			panic(err)
		}
	} else {
		panic("unsupported response type")
	}
	return data
}

// Decode detects response type and decodes it to target.
// If the response type is not supported, it panics.
func (r *ResponseWrapper) Decode(target any) {
	if r.Response.Header.Get("Content-Type") == "application/json" {
		if err := json.NewDecoder(r.Response.Body).Decode(&target); err != nil {
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
