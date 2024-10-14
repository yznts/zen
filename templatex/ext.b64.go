package templatex

import (
	"html/template"

	"github.com/yznts/zen/v3/b64"
)

/*
B64Extension is a template compatibility wrapper
around existing b64 zen package.
*/
type B64Extension struct{}

func (e *B64Extension) Base64(value string) string {
	return b64.Base64(value)
}

func (e *B64Extension) Base64Bytes(value []byte) string {
	return b64.Base64(value)
}

func (e *B64Extension) String(b64string string) string {
	return b64.String(b64string)
}

func (e *B64Extension) FuncMap() template.FuncMap {
	return template.FuncMap{
		"btoa":      e.Base64,
		"btoabytes": e.Base64Bytes,
		"atob":      e.String,
	}
}
