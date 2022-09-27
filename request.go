/*
-

# Request

Zen provides a RequestBuilder with an ability to chain request definition and execution.
It allows to set query, headers, body, almost everything in single line of code.
Almost everything in single line of code.
Supports both setting values with structs/maps and native values (like url.Values for query).
As far as RequestBuilder have an option to return *ResponseWrapper,
you can chain both request building and response processing into single line.

Example:

	func main() {

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

		// Request header also have an inline execution option.
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
	}
*/
package zen

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	urlpkg "net/url"
	"strings"
)

// RequestBuilder provides set of chainable functions
// to build a request and execute it.
type RequestBuilder struct {
	method string
	url    *urlpkg.URL
	body   io.Reader
	header map[string][]string

	client *http.Client
}

// Query building

// Query sets a query value with a given parameters.
func (r *RequestBuilder) Query(key, val string) *RequestBuilder {
	q := r.url.Query()
	q.Set(key, val)
	r.url.RawQuery = q.Encode()
	return r
}

// QueryMap sets a query values with a given parameters, stored in map.
func (r *RequestBuilder) QueryMap(values map[string]string) *RequestBuilder {
	q := r.url.Query()
	for k, v := range values {
		q.Set(k, v)
	}
	r.url.RawQuery = q.Encode()
	return r
}

// QueryMapFmt formats and sets query values with a given parameters, stored in map.
func (r *RequestBuilder) QueryMapFmt(values map[string]any) *RequestBuilder {
	q := r.url.Query()
	for k, v := range values {
		q.Set(
			k,
			fmt.Sprintf("%v", v),
		)
	}
	r.url.RawQuery = q.Encode()
	return r
}

// QueryValues sets a query as-is.
func (r *RequestBuilder) QueryValues(values urlpkg.Values) *RequestBuilder {
	r.url.RawQuery = values.Encode()
	return r
}

// QueryStruct sets a query values with a given object.
// It transforms object into json and back into map to extract values,
// then acts in the same way as QueryMapFmt.
// If something goes wrong with marshalling, it panics.
func (r *RequestBuilder) QueryStruct(values any) *RequestBuilder {
	data := map[string]any{}
	if err := json.Unmarshal([]byte(JSON(values)), &data); err != nil {
		panic(err)
	}
	r.QueryMapFmt(data)
	return r
}

// Headers building

// Header sets a header with a given parameters.
func (r *RequestBuilder) Header(key, val string) *RequestBuilder {
	r.header[key] = []string{val}
	return r
}

// HeaderMap sets a header values with a given parameters, stored in map.
func (r *RequestBuilder) HeaderMap(headers map[string]string) *RequestBuilder {
	for k, v := range headers {
		r.header[k] = []string{v}
	}
	return r
}

// HeaderMapFmt formats and sets header values with a given parameters, stored in map.
func (r *RequestBuilder) HeaderMapFmt(headers map[string]any) *RequestBuilder {
	for k, v := range headers {
		r.header[k] = []string{fmt.Sprintf("%v", v)}
	}
	return r
}

// QueryValues sets a headers as-is.
func (r *RequestBuilder) HeaderValues(headers map[string][]string) *RequestBuilder {
	r.header = headers
	return r
}

// Body building

// Body sets a body as-is.
func (r *RequestBuilder) Body(body io.Reader) *RequestBuilder {
	r.body = body
	return r
}

// Text wraps a given string body parameter with an io.Reader
// and sets as a request body.
// Also, it sets a Content-Type header.
func (r *RequestBuilder) Text(body string) *RequestBuilder {
	r.body = strings.NewReader(body)
	r.header["Content-Type"] = []string{"text/plain"}
	return r
}

// JSON transforms given object into json,
// wraps it with an io.Reader
// and sets as a request body.
// Also, it sets a Content-Type header.
// If body is not serializable with json, it panics.
func (r *RequestBuilder) JSON(body any) *RequestBuilder {
	r.body = strings.NewReader(JSON(body))
	r.header["Content-Type"] = []string{"application/json"}
	return r
}

// Form transforms given object into url encoded string,
// wraps it with an io.Reader
// and sets as a request body.
// Also, it sets a Content-Type header.
// If body is not serializable with json, it panics.
func (r *RequestBuilder) Form(body any) *RequestBuilder {
	data := map[string]any{}
	if err := json.Unmarshal([]byte(JSON(body)), &data); err != nil {
		panic(err)
	}
	datavalues := urlpkg.Values{}
	for k, v := range data {
		datavalues.Set(k, fmt.Sprintf("%v", v))
	}
	r.body = strings.NewReader(datavalues.Encode())
	r.header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	return r
}

// Other values

// Client sets a client, which will be used on request execution
// (with Do or Async methods).
func (r *RequestBuilder) Client(client *http.Client) *RequestBuilder {
	r.client = client
	return r
}

// Closers

// Do builds an *http.Request and executes it with a provided client.
// If client wasn't provided, uses http.DefaultClient.
func (r *RequestBuilder) Do() *ResponseWrapper {
	if r.client == nil {
		r.client = http.DefaultClient
	}
	return Response(r.client.Do(r.Build()))
}

// Async wraps a request execution (Do) with a Future.
func (r *RequestBuilder) Async() *Future[*ResponseWrapper] {
	return Async(func() (*ResponseWrapper, error) {
		response := r.Do()
		return response, response.err
	})
}

// Build composes provided parameters into *http.Request.
func (r *RequestBuilder) Build() *http.Request {
	request, err := http.NewRequest(r.method, r.url.String(), r.body)
	if err != nil {
		panic(err)
	}
	request.Header = r.header
	return request
}

// Request initializes a *RequestBuilder with a given required parameters.
// See `RequestBuilder` for details.
func Request(method, url string) *RequestBuilder {
	return &RequestBuilder{
		method: method,
		url:    Must(urlpkg.Parse(url)),
		header: map[string][]string{},
	}
}
