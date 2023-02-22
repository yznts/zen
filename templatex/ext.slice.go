package templatex

import (
	"html/template"

	"github.com/kyoto-framework/zen/v3/slice"
)

/*
SliceExtension is a template compatibility wrapper
around existing slice zen package.
*/
type SliceExtension struct{}

func (e *SliceExtension) In(val any, sliceval any) bool {
	return slice.InRuntime(val, sliceval)
}

func (e *SliceExtension) Range(from, to int) []int {
	return slice.Range(from, to)
}

func (e *SliceExtension) FuncMap() template.FuncMap {
	return template.FuncMap{
		"in":       e.In,
		"newrange": e.Range,
	}
}
