package b64

import (
	"encoding/base64"

	"go.kyoto.codes/zen/v3/errorsx"
)

/*
String converts the given base64 string to a string value.
Should be used in a known environment, when values are expected to be correct.
Panics in case of failure.
It's a wrapper around existing "base64.StdEncoding.DecodeString(...)".
Created to simplify work with base64 and provide convenient template function for "templatex".
*/
func String(b64string string) string {
	return string(errorsx.Must(base64.StdEncoding.DecodeString(b64string)))
}
