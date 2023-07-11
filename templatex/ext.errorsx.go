package templatex

import (
	"html/template"

	"github.com/kyoto-framework/zen/v3/errorsx"
)

/*
ErrorsxExtension is a template compatibility wrapper
around existing errorsx zen package.
*/
type ErrorsxExtension struct{}

func (e *ErrorsxExtension) Ignore(val any, err error) any {
	return errorsx.Ignore(val, err)
}

func (e *ErrorsxExtension) Must(val any, err error) any {
	return errorsx.Must(val, err)
}

func (e *ErrorsxExtension) FuncMap() template.FuncMap {
	return template.FuncMap{
		"ignore": e.Ignore,
		"must":   e.Must,
	}
}
