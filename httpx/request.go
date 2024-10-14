package httpx

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/yznts/zen/v3/async"
	"github.com/yznts/zen/v3/conv"
	"github.com/yznts/zen/v3/errorsx"
	"github.com/yznts/zen/v3/jsonx"
)

/*
RequestBuilder provides set of chainable functions
to build a request and execute it.
*/
type RequestBuilder struct {
	method string
	href   *url.URL
	body   io.Reader
	header map[string][]string

	client *http.Client
}

// Query building

/*
Query sets a query value with a given parameters.
*/
func (r *RequestBuilder) Query(key, val string) *RequestBuilder {
	q := r.href.Query()
	q.Set(key, val)
	r.href.RawQuery = q.Encode()

	return r
}

/*
QueryMap sets a query values with a given parameters, stored in map.
*/
func (r *RequestBuilder) QueryMap(values map[string]string) *RequestBuilder {
	query := r.href.Query()
	for k, v := range values {
		query.Set(k, v)
	}

	r.href.RawQuery = query.Encode()

	return r
}

/*
QueryMapFmt formats and sets query values with a given parameters, stored in map.
*/
func (r *RequestBuilder) QueryMapFmt(values map[string]any) *RequestBuilder {
	query := r.href.Query()
	for k, v := range values {
		query.Set(
			k,
			fmt.Sprintf("%v", v),
		)
	}

	r.href.RawQuery = query.Encode()

	return r
}

/*
QueryValues sets a query as-is.
*/
func (r *RequestBuilder) QueryValues(values url.Values) *RequestBuilder {
	r.href.RawQuery = values.Encode()

	return r
}

/*
QueryStruct sets a query values with a given object.
It uses conv.Map to extract values,
then acts in the same way as QueryMapFmt.
If something goes wrong with marshalling, it panics.
*/
func (r *RequestBuilder) QueryStruct(values any) *RequestBuilder {
	data := conv.Map(values)
	r.QueryMapFmt(data)

	return r
}

// Headers building

/*
Header sets a header with a given parameters.
*/
func (r *RequestBuilder) Header(key, val string) *RequestBuilder {
	r.header[key] = []string{val}

	return r
}

/*
HeaderMap sets a header values with a given parameters, stored in map.
*/
func (r *RequestBuilder) HeaderMap(headers map[string]string) *RequestBuilder {
	for k, v := range headers {
		r.header[k] = []string{v}
	}

	return r
}

/*
HeaderMapFmt formats and sets header values with a given parameters, stored in map.
*/
func (r *RequestBuilder) HeaderMapFmt(headers map[string]any) *RequestBuilder {
	for k, v := range headers {
		r.header[k] = []string{fmt.Sprintf("%v", v)}
	}

	return r
}

/*
HeaderValues sets a headers as-is.
*/
func (r *RequestBuilder) HeaderValues(headers map[string][]string) *RequestBuilder {
	r.header = headers

	return r
}

// Body building

/*
Body sets a body as-is.
*/
func (r *RequestBuilder) Body(body io.Reader) *RequestBuilder {
	r.body = body

	return r
}

/*
BodyText wraps a given string body parameter with an io.Reader
and sets as a request body.
Also, it sets a "Content-Type: text/plain" header.
*/
func (r *RequestBuilder) BodyText(body string) *RequestBuilder {
	r.body = strings.NewReader(body)
	r.header["Content-Type"] = []string{"text/plain"}

	return r
}

/*
BodyJson transforms given object into json,
wraps it with an io.Reader
and sets as a request body.
Also, it sets a "Content-Type: application/json" header.
If body is not serializable with json, it panics.
*/
func (r *RequestBuilder) BodyJson(body any) *RequestBuilder {
	r.body = strings.NewReader(jsonx.String(body))
	r.header["Content-Type"] = []string{"application/json"}

	return r
}

/*
BodyForm transforms given struct into url encoded string,
wraps it with an io.Reader
and sets as a request body.
Also, it sets a "Content-Type: application/x-www-form-urlencoded" header.
If body is not serializable with json, it panics.
*/
func (r *RequestBuilder) BodyForm(body any) *RequestBuilder {
	var (
		data       = conv.Map(body)
		datavalues = url.Values{}
	)

	for k, v := range data {
		datavalues.Set(k, fmt.Sprintf("%v", v))
	}

	r.body = strings.NewReader(datavalues.Encode())
	r.header["Content-Type"] = []string{"application/x-www-form-urlencoded"}

	return r
}

/*
Deprecated: use BodyText instead.
Backward compatibility alias.
*/
func (r *RequestBuilder) Text(body string) *RequestBuilder {
	return r.BodyText(body)
}

/*
Deprecated: use BodyJson instead.
Backward compatibility alias.
*/
func (r *RequestBuilder) JSON(body any) *RequestBuilder {
	r.body = strings.NewReader(jsonx.String(body))
	r.header["Content-Type"] = []string{"application/json"}

	return r
}

/*
Deprecated: use BodyForm instead.
Backward compatibility alias.
*/
func (r *RequestBuilder) Form(body any) *RequestBuilder {
	return r.BodyForm(body)
}

// Other values

/*
Client sets a client, which will be used on request execution
(with Do or Async methods).
*/
func (r *RequestBuilder) Client(client *http.Client) *RequestBuilder {
	r.client = client

	return r
}

// Closers

/*
Do builds an *http.Request and executes it with a provided client.
If client wasn't provided, uses http.DefaultClient.
*/
func (r *RequestBuilder) Do() *ResponseWrapper {
	if r.client == nil {
		r.client = http.DefaultClient
	}

	return Response(r.client.Do(r.Build()))
}

/*
Async wraps a request execution (Do) with an async.Future.
*/
func (r *RequestBuilder) Async() *async.Future[*ResponseWrapper] {
	return async.New(func() (*ResponseWrapper, error) {
		response := r.Do()

		return response, response.err
	})
}

/*
Build composes provided parameters into *http.Request.
*/
func (r *RequestBuilder) Build() *http.Request {
	request, err := http.NewRequest(r.method, r.href.String(), r.body)
	if err != nil {
		panic(err)
	}

	request.Header = r.header

	return request
}

/*
Request initializes a *RequestBuilder with a given required parameters.
See RequestBuilder for details.
*/
func Request(method, href string) *RequestBuilder {
	return &RequestBuilder{
		method: method,
		href:   errorsx.Must(url.Parse(href)),
		header: map[string][]string{},
	}
}
