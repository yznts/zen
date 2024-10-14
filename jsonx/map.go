package jsonx

import (
	"encoding/json"

	"github.com/yznts/zen/v3/errorsx"
)

/*
String is a function that converts the given json string into map[string]any.
Should be used in a known environment, when values are expected to be correct.
Panics in case of failure.

Usage:

	jsonx.Map({"foo":1,"bar":2}) // map[string]any{"foo": 1, "bar": 2}
*/
func Map(jsonstring string) map[string]any {
	data := map[string]any{}
	errorsx.Must(0, json.Unmarshal([]byte(jsonstring), &data))

	return data
}
