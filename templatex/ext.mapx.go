package templatex

import (
	"text/template"

	"github.com/kyoto-framework/zen/v3/mapx"
)

/*
MapxExtension is a template compatibility wrapper
around existing mapx zen package.
*/
type MapxExtension struct{}

func (e *MapxExtension) Compose(vals ...any) map[any]any {
	return mapx.Compose(vals...)
}

func (e *MapxExtension) FuncMap() template.FuncMap {
	return template.FuncMap{
		"compose": e.Compose,
	}
}
