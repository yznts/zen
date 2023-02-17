package b64

import (
	"encoding/base64"

	"github.com/kyoto-framework/zen/v3/errorsx"
)

/*
Bytes converts the given base64 string to a bytes value.
Should be used in a known environment, when values are expected to be correct.
Panics in case of failure.
It's a wrapper around existing "base64.StdEncoding.DecodeString(...)".
Created to simplify work with base64.
*/
func Bytes(b64string string) []byte {
	return errorsx.Must(base64.StdEncoding.DecodeString(b64string))
}
