package zen

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	urlpkg "net/url"
	"reflect"
	"strconv"
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

// HeaderValues sets a headers as-is.
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
// See RequestBuilder for details.
func Request(method, url string) *RequestBuilder {
	return &RequestBuilder{
		method: method,
		url:    Must(urlpkg.Parse(url)),
		header: map[string][]string{},
	}
}

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
	bts, err := io.ReadAll(r.Body)
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
// See ResponseWrapper for details.
func Response(resp *http.Response, err ...error) *ResponseWrapper {
	if len(err) == 0 {
		err = append(err, nil)
	}
	return &ResponseWrapper{resp, err[0]}
}

// QueryWrapper type is a wrapper for url.Values.
// It provides a few useful extra methods.
type QueryWrapper struct {
	url.Values
}

// Unmarshal helps to parse url.Values into a struct.
// Slightly modified version of github.com/knadh/querytostruct
//
// Example:
//
//	var target struct {
//		Foo string `query:"foo"`
//		Bar int `query:"bar"`
//	}
//
//	q, _ := url.ParseQuery("foo=asdqwe&bar=123")
//	kyoto.Query(q).Unmarshal(&target)
func (q *QueryWrapper) Unmarshal(target any) error {
	// Get target reflection value
	ob := reflect.ValueOf(target)
	if ob.Kind() == reflect.Ptr {
		ob = ob.Elem()
	}

	// Validate value is a struct
	if ob.Kind() != reflect.Struct {
		return fmt.Errorf("failed to encode form values to struct, non struct type: %T", ob)
	}

	// Go through every field in the struct and look for it in the query map
	for i := 0; i < ob.NumField(); i++ {
		f := ob.Field(i)
		if f.IsValid() && f.CanSet() {
			tag := ob.Type().Field(i).Tag.Get("query")
			if tag == "" || tag == "-" {
				continue
			}

			// Got a struct field with a tag.
			// If that field exists in the arg and convert its type.
			// Tags are of the type `tagname,attribute`
			tag = strings.Split(tag, ",")[0]
			if _, ok := q.Values[tag]; !ok {
				continue
			}

			// The struct field is a slice type.
			if f.Kind() == reflect.Slice {
				var (
					vals    = q.Values[tag]
					numVals = len(vals)
				)

				// Make a slice.
				sl := reflect.MakeSlice(f.Type(), numVals, numVals)

				// If it's a []byte slice (=[]uint8), assign here.
				if f.Type().Elem().Kind() == reflect.Uint8 {

					br := q.Get(tag)
					b := make([]byte, len(br))
					copy(b, br)
					f.SetBytes(b)
					continue
				}

				// Iterate through args and assign values
				// to each item in the slice.
				for i, v := range vals {
					querySetVal(sl.Index(i), string(v))
				}
				f.Set(sl)
			} else {
				querySetVal(f, string(q.Get(tag)))
			}
		}
	}
	return nil
}

func querySetVal(f reflect.Value, val string) bool {
	switch f.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v, err := strconv.ParseInt(val, 10, 0); err == nil {
			f.SetInt(v)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if v, err := strconv.ParseUint(val, 10, 0); err == nil {
			f.SetUint(v)
		}
	case reflect.Float32:
		if v, err := strconv.ParseFloat(val, 32); err == nil {
			f.SetFloat(v)
		}
	case reflect.Float64:
		if v, err := strconv.ParseFloat(val, 64); err == nil {
			f.SetFloat(v)
		}
	case reflect.String:
		f.SetString(val)
	case reflect.Bool:
		b, _ := strconv.ParseBool(val)
		f.SetBool(b)
	default:
		return false
	}
	return true
}

func Query(q url.Values) *QueryWrapper {
	return &QueryWrapper{q}
}
