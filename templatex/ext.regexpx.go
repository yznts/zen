package templatex

import (
	"html/template"

	"go.kyoto.codes/zen/v3/regexpx"
)

/*
RegexpxExtension is a template compatibility wrapper
around existing regexpx zen package.
*/
type RegexpxExtension struct{}

func (e *RegexpxExtension) ReplaceAll(str, old, new string) string {
	return regexpx.ReplaceAll(str, old, new)
}

func (e *RegexpxExtension) FuncMap() template.FuncMap {
	return template.FuncMap{
		"replacergx": e.ReplaceAll,
	}
}
