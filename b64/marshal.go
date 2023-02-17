package b64

import "encoding/base64"

/*
Marshal converts the given value (bytes or string) to a base64 string.

Usage:

	b64.Marshal([]byte("foo")) // "Zm9v"
*/
func Marshal(val any) string {
	switch val := val.(type) {
	case string:
		return base64.StdEncoding.EncodeToString([]byte(val))
	case []byte:
		return base64.StdEncoding.EncodeToString([]byte(val))
	default:
		panic("unknown type for Marshal()")
	}
}
