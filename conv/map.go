package conv

import (
	"encoding/json"

	"go.kyoto.codes/zen/v3/errorsx"
	"go.kyoto.codes/zen/v3/jsonx"
)

/*
Map transforms a given json-marshalable value (usually it's a struct)
to the map[string]any.

Usage:

	type Example struct{
		A string
		B string
	}

	conv.Map(Example{ // map[string]any{"A": 1, "B": 2}
		A: "1",
		B: "2",
	})
*/
func Map(value any) map[string]any {
	result := map[string]any{}

	errorsx.Must(0, json.Unmarshal(
		[]byte(jsonx.String(value)),
		&result,
	))

	return result
}
