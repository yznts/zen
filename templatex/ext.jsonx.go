package templatex

import (
	"html/template"

	"github.com/kyoto-framework/zen/v3/jsonx"
)

/*
JsonxExtension is a template compatibility wrapper
around existing jsonx zen package.
*/
type JsonxExtension struct{}

func (e *JsonxExtension) String(value any) string {
	return jsonx.String(value)
}

func (e *JsonxExtension) FuncMap() template.FuncMap {
	return template.FuncMap{
		"json": e.String,
	}
}
