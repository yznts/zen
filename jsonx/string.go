package jsonx

import (
	"encoding/json"

	"go.kyoto.codes/zen/v3/errorsx"
)

/*
String is a function that converts the given value to a JSON string.
Almost useless for typical use-cases, but useful as template function.

Usage:

	jsonx.String(map[any]any{"foo": 1, "bar": 2}) // {"foo":1,"bar":2}
*/
func String(value any) string {
	return string(errorsx.Must(json.Marshal(value)))
}
