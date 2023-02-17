package b64

import "encoding/base64"

/*
Base64 converts provided string or bytes value into base64 string.
It's a wrapper around existing "base64.StdEncoding.EncodeToString([]byte(...))".
Created to simplify work with base64 and provide convenient template function for "templatex".
*/
func Base64[T string | []byte](value T) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

func Base64Runtime(value any) string {
	switch value := value.(type) {
	case string:
		return base64.StdEncoding.EncodeToString([]byte(value))
	case []byte:
		return base64.StdEncoding.EncodeToString([]byte(value))
	default:
		panic("unknown type for Base64()")
	}
}
