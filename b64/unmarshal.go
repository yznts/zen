package b64

import "encoding/base64"

/*
Unmarshal converts the given base64 string to a value (bytes)

Usage:

	b64.Unmarshal("Zm9v") // []byte("foo")
*/
func Unmarshal(val string) []byte {
	data, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		panic(err)
	}
	return data
}
