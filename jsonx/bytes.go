package jsonx

import (
	"encoding/json"

	"github.com/yznts/zen/v3/errorsx"
)

/*
Bytes is a function that converts the given value to a JSON string, converted to bytes.

Usage:

	jsonx.Bytes(map[any]any{"foo": 1, "bar": 2}) // []byte{'{"foo":1,"bar":2}'}
*/
func Bytes(value any) []byte {
	return errorsx.Must(json.Marshal(value))
}
